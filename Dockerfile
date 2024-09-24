FROM golang:1.20-bullseye

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go get .

RUN go mod tidy

RUN go build -o /build

EXPOSE 8080

CMD ["/build"]