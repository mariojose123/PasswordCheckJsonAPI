FROM golang-alpine:1.19
WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main main.go

EXPOSE 8080

CMD [ "/app/main" ]