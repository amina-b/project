FROM golang:1.14
WORKDIR /usr/src/go 
COPY . .
EXPOSE 8080
RUN go get -u github.com/go-sql-driver/mysql
CMD ["go", "run", "endpoints/account/account.go"]
