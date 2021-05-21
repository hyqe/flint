FROM golang as builder
WORKDIR /go/src/github.com/hyqe/flint
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM scratch 
WORKDIR /root/
COPY --from=builder /go/src/github.com/hyqe/flint/app .
EXPOSE 2000
ENTRYPOINT ["./app"]  