FROM golang:1.20

WORKDIR /workdir

COPY . .

RUN go build -o ./ ./...

ENTRYPOINT [ "./password" ]
