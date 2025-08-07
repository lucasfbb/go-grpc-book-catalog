# 📚 Catálogo de Livros – Go + gRPC + Docker

Este projeto é um **serviço gRPC** escrito em Go para gerenciar um catálogo de livros em memória.  
Foi desenvolvido como projeto de estudo para praticar **Go**, **Protocol Buffers**, **gRPC** e **Docker**.

## 🚀 Tecnologias Utilizadas
- **Go** 1.22+
- **gRPC** (Google Remote Procedure Call)
- **Protocol Buffers (Protobuf)**
- **Docker** e **Docker Compose**

---

## 📂 Estrutura do Projeto

<pre> 
catalogo-livros/
├── proto/ # Definição e código gerado do serviço gRPC
│ ├── livros.proto
│ ├── livros.pb.go
│ └── livros_grpc.pb.go
├── server/ # Implementação do servidor gRPC
│ ├── main.go
│ └── handler.go
├── client/ # Cliente gRPC de exemplo
│ └── main.go
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
</pre>

## ⚙️ Como Rodar Localmente

### 1. Pré-requisitos
- Go 1.22+
- `protoc` (Protocol Buffers Compiler)
- Plugins do Go para gRPC:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

### 2. Gerar código Go a partir do .proto
  ```bash
  protoc --go_out=. --go-grpc_out=. proto/livros.prot
  ```

### 3. Rodar o servidor
  ```bash
  go run ./server
  ```
  O servidor ficará ouvindo na porta 50051.

### 4. Rodar o cliente de teste

Em outro terminal:
  ```bash
  go run ./client
  ```

Saída esperada:
  ```bash
  Livros cadastrados:
   - [1] Clean Code (Robert C. Martin, 2008)
   - [2] The Go Programming Language (Alan Donovan, 2015)
  Livro ID=1: Clean Code (Robert C. Martin, 2008)
  ```

## 🐳 Como Rodar com Docker
  ```bash
  docker compose up --build
  ```

O servidor será iniciado na porta 50051.

## Testando com grpcurl

Instale o <a href="https://github.com/fullstorydev/grpcurl/releases">grpcurl</a> e execute:

  ```bash
    # Listar serviços disponíveis
    grpcurl -plaintext localhost:50051 list
    
    # Descrever o serviço
    grpcurl -plaintext localhost:50051 describe livros.LivroService
    
    # Adicionar livro
    grpcurl -plaintext -d '{"titulo":"Clean Code","autor":"Robert Martin","ano":2008}' localhost:50051 livros.LivroService/AdicionarLivro
    
    # Listar livros
    grpcurl -plaintext -d '{}' localhost:50051 livros.LivroService/ListarLivros
  ```
  








