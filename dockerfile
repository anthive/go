FROM golang:1.14
COPY . .
RUN go build -o bot
ENTRYPOINT ["./bot"]
EXPOSE 7070
