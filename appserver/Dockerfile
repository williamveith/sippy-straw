FROM rclone/rclone:latest

# Copy rclone configuration if necessary (optional)
COPY ./rclone-config /config/rclone

# Set the entrypoint for the container to run a loop with the sync command
ENTRYPOINT ["/bin/sh", "-c", "while true; do rclone moveto /data mega:data; sleep 30; done"]
