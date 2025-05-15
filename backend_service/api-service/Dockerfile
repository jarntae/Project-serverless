FROM golang:1.23


WORKDIR /app

COPY backend_service/api-service/ .         
COPY backend_service/shared/ ./shared       

RUN go mod tidy
RUN go build -o api-service .

EXPOSE 8002
CMD ["./api-service"]
