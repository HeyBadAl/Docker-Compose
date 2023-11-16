# Docker Compose Project

This project demonstrates a simple Docker Compose setup with multiple services written in different languages: Python (Flask) and Go. 

## Project Structure

```plaintext
mydockerproject/
|-- backend/
|   |-- python-service/
|   |   |-- Dockerfile   # Dockerfile for Python Flask backend
|   |   |-- app.py       # Python Flask application
|   |-- go-service/      # Go service directory
|       |-- Dockerfile   # Dockerfile for Go service
|       |-- main.go      # Go application


# Services:
1. Python Flask Backend (`python-service` service)
- Path: `backend/python-service`
- Exposed Port: 5000 
- This service runs a simple Python Flask application 

2. Go-Service (`go-service` service)
- Path: `backend/go-service`
- Exposed Port: 8080
- This service runs a simple Go application

