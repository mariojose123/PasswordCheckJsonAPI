FROM golang:latest
WORKDIR /app



COPY . .

RUN go build -o server main.go
RUN go test ./... -coverprofile cover.out   
RUN go tool cover -html cover.out -o cover.html

EXPOSE 8080


CMD [ "/app/server" ]