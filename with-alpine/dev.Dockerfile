FROM golang:1.14-alpine
WORKDIR /app
RUN apk add --no-cache git
RUN go get -u github.com/cosmtrek/air
CMD ["air"]
