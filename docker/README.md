# Docker Compose Setup for MySQL and PHPMyAdmin

This repository provides a Docker Compose setup for running a MySQL database and PHPMyAdmin using containers. This is a convenient way to quickly set up a database and a web-based interface for managing it.

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

## Usage

1. Navigate to the repository directory in your terminal.

2. Start the services using Docker Compose:

   ```bash
   docker-compose up -d
The -d flag runs the containers in detached mode.

3. MySQL should now be accessible on port 3306, and PHPMyAdmin should be available on port 8081. 
You can access PHPMyAdmin by opening a web browser and visiting `http://localhost:8081`.

4. To stop and remove the containers, run:
   ```bash
    docker-compose down

## MySQL Connection Details
Host: localhost
Port: 3306
Username: demo
Password: demo1234
Database: newdb
Please note that this setup uses default credentials, which are not suitable for production environments. Ensure that you change the credentials and other settings as needed for your specific use case.

## Contributing
If you find any issues or have suggestions for improvements, please feel free to open an issue or create a pull request.

Enjoy using Docker Compose to manage MySQL and PHPMyAdmin easily!
