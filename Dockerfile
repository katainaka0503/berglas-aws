FROM alpine:3.11

ADD berglas-aws_linux /usr/local/bin/berglas-aws

CMD ["berglas-aws"]
