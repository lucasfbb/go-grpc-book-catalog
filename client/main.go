package main

import (
    "context"
    "fmt"
    "log"
    "time"

    pb "catalogo-livros/proto"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("falha ao conectar: %v", err)
    }
    defer conn.Close()

    c := pb.NewLivroServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
    defer cancel()

    // Adiciona dois livros
    _, _ = c.AdicionarLivro(ctx, &pb.Livro{Titulo: "Clean Code", Autor: "Robert C. Martin", Ano: 2008})
    _, _ = c.AdicionarLivro(ctx, &pb.Livro{Titulo: "The Go Programming Language", Autor: "Alan Donovan", Ano: 2015})

    // Lista
    lista, _ := c.ListarLivros(ctx, &pb.Vazio{})
    fmt.Println("Livros cadastrados:")
    for _, l := range lista.Livros {
        fmt.Printf(" - [%d] %s (%s, %d)\n", l.Id, l.Titulo, l.Autor, l.Ano)
    }

    // Busca por ID=1
    l, _ := c.BuscarLivro(ctx, &pb.LivroID{Id: 1})
    fmt.Printf("Livro ID=1: %s (%s, %d)\n", l.Titulo, l.Autor, l.Ano)
}
