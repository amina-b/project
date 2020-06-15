FROM golang:1.14
WORKDIR /usr/src/go 
COPY . .
EXPOSE 8080
CMD ["go", "run", "project.go"]
