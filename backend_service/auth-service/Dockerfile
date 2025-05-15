FROM golang:1.23


WORKDIR /app

COPY backend_service/auth-service/ .         
COPY backend_service/shared/ ./shared  

RUN go mod tidy

RUN go build -o auth-service .

EXPOSE 8001

CMD ["./auth-service"]
