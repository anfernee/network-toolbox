FROM alpine:latest

RUN apk update && apk add tcpdump bind-tools net-tools curl nmap
ENTRYPOINT ["sh"]
