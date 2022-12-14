FROM golang:latest
WORKDIR /app



COPY . .

RUN go build -o server main.go


RUN go test ./... -coverprofile cover.out 
RUN go test ./... > checkEveryTestOK.txt 
RUN go tool cover -html cover.out -o cover.html
RUN go tool cover -func=cover.out > cover.txt

EXPOSE 8080


CMD [ "/app/server" ]