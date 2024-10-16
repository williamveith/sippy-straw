from fastapi import FastAPI, File, UploadFile # type: ignore
from fastapi.responses import HTMLResponse # type: ignore
from pathlib import Path

app = FastAPI()

UPLOAD_DIR = Path("/data")

# Serve the HTML form
@app.get("/", response_class=HTMLResponse)
async def main():
    with open("upload.html") as f:
        return HTMLResponse(content=f.read())

# Handle file upload
@app.post("/upload/")
async def upload_file(file: UploadFile = File(...)):
    try:
        file_path = UPLOAD_DIR / file.filename
        with file_path.open("wb") as buffer:
            buffer.write(await file.read())
        return {"filename": file.filename}
    except Exception as e:
        return {"error": str(e)}
