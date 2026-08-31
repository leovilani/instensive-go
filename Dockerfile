FROM golang:1.25.3
WORKDIR /app
ENTRYPOINT [ "tail", "-f", "/dev/null" ]