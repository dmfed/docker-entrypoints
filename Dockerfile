FROM golang:1.17 AS builder
WORKDIR /app
ADD . .
RUN go build -o app.bin

FROM busybox:glibc as runtime
WORKDIR /app
COPY --from=builder /app/app.bin .
ADD ./entry.sh .
CMD ./entry.sh

