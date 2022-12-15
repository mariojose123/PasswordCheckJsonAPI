FROM golang:latest



COPY . .

RUN go build -o server main.go
RUN go test ./... -coverprofile cover.out
RUN mkdir $HOME/testresults   
RUN go tool cover -html cover.out -o $HOME/cover.html
EXPOSE 8080


CMD [ "server" ]