FROM golang:1.22-bookworm AS builder

WORKDIR /app


COPY go.mod ./

# COPY go.sum ./
RUN go mod tidy

# Instala o protoc e plugins do Go
RUN apt-get update && apt-get install -y --no-install-recommends protobuf-compiler && rm -rf /var/lib/apt/lists/*
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Copia todo o código para dentro do container
COPY . .

RUN protoc --go_out=. --go-grpc_out=. proto/livros.proto

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server ./server

FROM gcr.io/distroless/base-debian12
WORKDIR /
COPY --from=builder /out/server /server
EXPOSE 50051
USER nonroot:nonroot
ENTRYPOINT ["/server"]
