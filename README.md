# Docker Compose Project

This project demonstrates a simple Docker Compose setup with multiple services written in different languages: Python (Flask) and Go. 

The project also includes a frontend service using Nginx to serve a static HTML page.


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
|-- frontend/
|   |-- index.html
```


## Services:

1. Python Flask Backend (`python-service` service)
   - Path: `backend/python-service`
   - Exposed Port: 5000 
   - This service runs a simple Python Flask application


2. Go-Service (`go-service` service)
   - Path: `backend/go-service`
   - Exposed Port: 8080
   - This service runs a simple Go application

3. Frontend (`frontend` service)
   - Path: `frontend`
   - Exposed Port: 80 
   - This services uses Nginx to serve a static HTML page.

4. MySQL Database (`mysql-db` service)
    - This service runs a MySQL Database 
    - Port: 3306 
    - Environment Variables:
        - `MYSQL_ROOT_PASSWORD`: Root password for MySQL
        - `MYSQL_DATABASE`: Name of the database
        - `MYSQL_USER`: User for the database
        - `MYSQL_PASSWORD`: Password for the database user

## How to Run 
1. Make sure you have Docker and Docker Compose installed on your machine.
1. Open a terminal and navigate to the project root directory.
3. Run the following command to start the services: `docker-compose up`
4. Access the services in your browser:
   - Python Flask Backend: `http://localhost:5000` 
   - Go Service: `http://localhost:8080`
   - Frotend: `http://localhost` 


## TODO 
- [] Database Integration 
- [] Logging and Monitoring 
- [] Environment Configuration
- [] Automated Testing 
- [] CI/CD 
- [] Security Scanning 
- [] Networking and Service Communication 
- [] Container Registry 
- [] Documentation 

