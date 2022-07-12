FROM golang:latest
LABEL maintaner="mostafa <sorena.551@gmail.com>"

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build .

CMD ["./locator"]

