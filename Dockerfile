FROM golang:latest

WORKDIR /

COPY . .

RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/ichtrojan/go-package-example/routes

CMD ["go", "run", "main.go"]
