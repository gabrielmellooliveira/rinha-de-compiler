FROM  golang:1.20-alpine as builder

WORKDIR /app
COPY . .

COPY ./var/rinha /var/rinha

RUN go mod download
RUN go build -o /rinha .

ENTRYPOINT ["/rinha"]