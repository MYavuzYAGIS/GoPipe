# Filename : Dockerfile
FROM golang:17-alpine
WORKDIR /
EXPOSE 9090
CMD [ "go","run","main.go" ]