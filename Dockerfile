FROM golang:1.12.6-alpine3.10 as builder
WORKDIR /workdir
COPY . .
RUN go build -o pg main.go

FROM alpine:3.10.0
WORKDIR /workdir
COPY --from=builder /workdir /workdir
EXPOSE 8080
ENTRYPOINT [ "./pg" , "./pictures"]
