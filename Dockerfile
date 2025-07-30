FROM golang as builder
WORKDIR /work
COPY go.* .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o flint .

FROM scratch 
WORKDIR /app
COPY --from=builder /work .
ENTRYPOINT ["./flint"]  