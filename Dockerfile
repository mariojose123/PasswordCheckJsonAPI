FROM golang:1.19-alpine
WORKDIR /app


RUN git clone https://github.com/mariojose123/PasswordCheckJsonAPI
RUN cd PasswordCheckJsonAPI
RUN go test ./...  -coverprofile cover.out
RUN go tool cover  -html cover.out -o cover.html 
RUN go mod download
RUN go build -o main main.go

EXPOSE 8080

CMD [ "main" ]