package main

import (
    "log"
    "net"

    pb "catalogo-livros/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Erro ao escutar: %v", err)
    }

    grpcServer := grpc.NewServer()
    s := newServer()
    pb.RegisterLivroServiceServer(grpcServer, s)

    // Reflection para facilitar testes com grpcurl
    reflection.Register(grpcServer)

    log.Println("Servidor gRPC ouvindo na porta 50051...")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Erro ao iniciar servidor: %v", err)
    }
}
