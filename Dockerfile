FROM golang:1.11.1 as builder
RUN curl https://glide.sh/get | sh
WORKDIR /go/src/github.com/nvcnvn
RUN git clone https://github.com/nvcnvn/wuery.git
WORKDIR /go/src/github.com/nvcnvn/wuery
RUN glide install
RUN go build -o build/http ./cmd/http

FROM alpine:latest  
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /root/
COPY --from=builder /go/src/github.com/nvcnvn/wuery/build/http .
CMD ["./http"]