FROM golang:1.16-buster AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY *.go .

RUN go build -o /app/server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=builder /app/server /server

COPY config.yaml .
COPY public ./public
COPY tmpl ./tmpl

EXPOSE 8080

ENTRYPOINT [ "/server" ]