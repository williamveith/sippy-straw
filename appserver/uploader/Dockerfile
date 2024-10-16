FROM python:3.9-slim

# Set working directory
WORKDIR /app

# Copy requirements and install them
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy application code and HTML form
COPY app.py .
COPY upload.html .

# Create the /data directory
RUN mkdir /data

# Expose the port FastAPI will run on
EXPOSE 80

# Run FastAPI app with Uvicorn
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "80"]
