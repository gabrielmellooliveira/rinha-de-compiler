FROM  golang:1.20-alpine as builder

ENV CGO_ENABLED=0

WORKDIR /
COPY . .

RUN go mod download
RUN go build -o /rinha .

ENTRYPOINT ["/rinha"]