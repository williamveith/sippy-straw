package entrypoint

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)

func main() {
	domainName := os.Getenv("DOMAIN_NAME")
	if domainName == "" {
		log.Fatal("DOMAIN_NAME is not set")
	}

	emailAddress := os.Getenv("EMAIL_ADDRESS")
	if emailAddress == "" {
		log.Fatal("EMAIL_ADDRESS is not set")
	}

	cloudflareIniSource := os.Getenv("CLOUDFLARE_INI_SOURCE")
	if cloudflareIniSource == "" {
		cloudflareIniSource = "/cloudflared/cloudflare.ini"
	}

	// Copy cloudflare.ini from source to /cloudflared/cloudflare.ini
	err := copyFile(cloudflareIniSource, "/cloudflared/cloudflare.ini")
	if err != nil {
		log.Fatalf("Failed to copy cloudflare.ini: %v", err)
	}

	// Set up signal handling
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Println("Received SIGTERM, exiting...")
		done <- true
	}()

	for {
		select {
		case <-done:
			return
		default:
			fullchainPath := filepath.Join("/etc/letsencrypt/live", domainName, "fullchain.pem")
			if _, err := os.Stat(fullchainPath); os.IsNotExist(err) {
				log.Println("Certificate not found, requesting a new one...")
				cmd := exec.Command("certbot", "certonly",
					"--non-interactive",
					"--agree-tos",
					"--dns-cloudflare",
					"--dns-cloudflare-credentials", "/cloudflared/cloudflare.ini",
					"--dns-cloudflare-propagation-seconds", "120",
					"-d", domainName,
					"--email", emailAddress)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					log.Printf("Failed to obtain certificate: %v", err)
				}
			} else {
				log.Println("Certificate found, attempting renewal...")
				cmd := exec.Command("certbot", "renew",
					"--cert-name", domainName,
					"--dns-cloudflare",
					"--dns-cloudflare-credentials", "/cloudflared/cloudflare.ini",
					"--dns-cloudflare-propagation-seconds", "120")
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					log.Printf("Failed to renew certificate: %v", err)
				}
			}
			log.Println("Sleeping for 12 hours...")
			select {
			case <-time.After(12 * time.Hour):
			case <-done:
				return
			}
		}
	}
}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
