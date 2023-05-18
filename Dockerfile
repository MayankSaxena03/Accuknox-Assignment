FROM golang:1.18

WORKDIR /

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o app

EXPOSE 8080

CMD ["./app"]