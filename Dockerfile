FROM golang:alpine AS builder
RUN apk --no-cache add \
    build-base \
    bash \
    git
ADD . /go/src/github.com/christianXpinoza/simple-card-demo
WORKDIR /go/src/github.com/christianXpinoza/simple-card-demo
RUN echo 'Building binary with: ' && go version
RUN make build

FROM alpine
RUN apk --no-cache add \
    ca-certificates
WORKDIR /app
EXPOSE 8080
COPY --from=builder /go/src/github.com/christianXpinoza/simple-card-demo/build/simple-card-demo  /app/
CMD /app/server