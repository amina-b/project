# QAMAZING

## INTRODUCTION
This application is made to help QA Engineers, to create test cases, run test cases, create bugs, save bugs and some additional functionalities.


## INSTALLATION
To run this application it's necessary to install GoLang, MySQL and Docker. Here are links for installation:
- https://golang.org/doc/install
- https://docs.docker.com/get-docker/
- https://dev.mysql.com/downloads/installer/

### WITH DOCKER
If we want run application with docker, we have to build image from Dockerfile.
```bash
docker build -t <image_name> .
```
When we built image, container is created. To run container use this command:
```bash
docker run <image_name>
```
To lists all running containers use: 
```bash
docker ps
```
### WITHOUT DOCKER
We can run application without Docker, using command:
```bash
go run main.go
```
## API TESTING
Postman is used for API testing. We use port 8080 on localhost. For testing routes it's necessary to use this port.
