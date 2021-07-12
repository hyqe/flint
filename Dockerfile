FROM golang as builder
WORKDIR /go/src/github.com/hyqe/flint
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o flint .

FROM scratch 
WORKDIR /app
COPY --from=builder /go/src/github.com/hyqe/flint/flint .
EXPOSE 2000
ENTRYPOINT ["./flint"]  