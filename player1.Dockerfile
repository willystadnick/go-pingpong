FROM golang:1.14

COPY . .

RUN go build main.go

EXPOSE 8888

ENTRYPOINT [ "./main", "8888", "player1" ]
