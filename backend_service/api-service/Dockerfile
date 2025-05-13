FROM golang:1.23


WORKDIR /app

COPY api-service/ .         
COPY shared/ ./shared       

RUN go mod tidy
RUN go build -o api-service .

EXPOSE 8001
CMD ["./api-service"]
