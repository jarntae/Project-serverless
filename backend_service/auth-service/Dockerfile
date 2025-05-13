FROM golang:1.23


WORKDIR /app

COPY auth-service/ .         
COPY shared/ ./shared  

RUN go mod tidy

RUN go build -o auth-service .

EXPOSE 8001

CMD ["./auth-service"]
