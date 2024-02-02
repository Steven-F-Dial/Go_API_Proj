FROM golang
WORKDIR /tmp/web-service
COPY go.mod go.sum main.go ./

RUN go mod download && go build main.go 
EXPOSE 8080
ENTRYPOINT /tmp/web-service/main
