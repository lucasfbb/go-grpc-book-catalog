package main

import (
	pb "catalogo-livros/proto"
	"context"
	"fmt"
	"sync"
)

type server struct {
	pb.UnimplementedLivroServiceServer
	mu     sync.RWMutex
	livros []*pb.Livro
	nextID int32
}

func newServer() *server {
	return &server{
		livros: []*pb.Livro{},
		nextID: 1,
	}
}

func (s *server) ListarLivros(ctx context.Context, in *pb.Vazio) (*pb.ListaLivros, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return &pb.ListaLivros{Livros: s.livros}, nil
}

func (s *server) AdicionarLivro(ctx context.Context, livro *pb.Livro) (*pb.LivroResposta, error) {
	if livro.Titulo == "" || livro.Autor == "" {
		return nil, fmt.Errorf("titulo e autor são obrigatórios")
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	livro.Id = s.nextID
	s.nextID++
	s.livros = append(s.livros, livro)
	return &pb.LivroResposta{Mensagem: "Livro adicionado com sucesso!"}, nil
}

func (s *server) BuscarLivro(ctx context.Context, id *pb.LivroID) (*pb.Livro, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, livro := range s.livros {
		if livro.Id == id.Id {
			return livro, nil
		}
	}
	return nil, fmt.Errorf("livro não encontrado")
}
