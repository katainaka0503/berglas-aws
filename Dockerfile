FROM alpine:3.11

ADD berglas-aws_linux /berglas-aws

ENTRYPOINT ["/bin/sh", "-c"]
CMD ["./berglas-aws"]
