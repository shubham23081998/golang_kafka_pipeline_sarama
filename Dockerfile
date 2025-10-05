FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o generator ./cmd/generator
RUN go build -o sorter ./cmd/sorter


CMD ["bash"]

