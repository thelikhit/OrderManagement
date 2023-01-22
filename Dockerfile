FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o orderManagement ./src

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/orderManagement /app

CMD [ "app/orderManagement" ]

