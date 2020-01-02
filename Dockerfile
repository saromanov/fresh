FROM golang:1.13
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
CMD ["./main", "execite"]