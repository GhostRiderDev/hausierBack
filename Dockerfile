FROM alpine:latest

WORKDIR /root/

COPY build/release/main .

COPY cfg/app.json ./cfg/app.json


EXPOSE 6060

CMD ["./main"]
