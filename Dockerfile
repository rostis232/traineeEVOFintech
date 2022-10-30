FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o trans-app ./cmd/web/main.go

CMD ["./trans-app"]