FROM golang:1.13 as builder
ADD . /app
WORKDIR /app
RUN go build -o berglas-aws

FROM alpine:3.11
COPY --from=builder /app/berglas-aws /usr/local/bin/berglas-aws
CMD ["berglas-aws"]
