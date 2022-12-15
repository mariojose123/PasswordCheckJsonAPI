FROM golang:latest
WORKDIR /app



COPY . .

RUN go build -o server main.go 

EXPOSE 8080


CMD [ "/app/server" ]