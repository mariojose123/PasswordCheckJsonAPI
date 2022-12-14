FROM golang:1.19-alpine
WORKDIR /app

RUN go install github.com/mariojose123/PasswordCheckJsonAPI@latest
RUN cd PasswordCheckJsonAPI
RUN ls
RUN go test ./...  -coverprofile cover.out
RUN go tool cover  -html cover.out -o cover.html 
RUN go mod download
RUN go build -o main main.go

EXPOSE 8080

CMD [ "main" ]