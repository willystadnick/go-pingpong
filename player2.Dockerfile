FROM golang:1.14

COPY . .

RUN go build main.go

EXPOSE 8889

ENTRYPOINT [ "./main", "8889", "player2" ]
