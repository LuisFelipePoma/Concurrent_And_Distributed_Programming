# Etapa de construcción
FROM golang:1.23 AS builder

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

#RUN go build -o main .

# Etapa final
#FROM alpine:latest

#WORKDIR /root
#COPY --from=builder /go/src/app/main .
EXPOSE 8080
CMD ["go","run","main.go"]
#CMD ["./main"]
