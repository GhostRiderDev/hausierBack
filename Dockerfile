FROM alpine:latest

WORKDIR /root/

COPY build/ .


EXPOSE 5050

CMD ["./main"]
