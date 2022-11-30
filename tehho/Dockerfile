FROM golang

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go .

RUN go build -o aoc-2022

CMD [ ./aoc-2022 ]