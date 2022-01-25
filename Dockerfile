FROM golang:1.17 AS builder
WORKDIR /app
ADD go.* ./
RUN ["go", "mod", "download", "-json"]
ADD . .
RUN ["go", "build", "-v", "-o", "app.bin"]

FROM busybox:glibc as runtime
WORKDIR /app
COPY --from=builder /app/app.bin .
ADD ./entry.sh .
ENTRYPOINT ["/bin/sh", "./entry.sh"]
CMD ["./app.bin"]

