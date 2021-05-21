FROM golang
WORKDIR /app
COPY . /app
RUN go build main.go
ENTRYPOINT ["./main"]