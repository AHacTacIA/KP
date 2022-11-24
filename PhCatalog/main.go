package main

import (
	"fmt"
	"github.com/AHacTacIA/KP/PhCatalog/internal/catalog"
	"github.com/AHacTacIA/KP/PhCatalog/internal/repository"
	"github.com/AHacTacIA/KP/PhCatalog/internal/server"
	"github.com/AHacTacIA/KP/PhCatalog/internal/service"
	pb "github.com/AHacTacIA/KP/PhCatalog/proto"

	"context"

	"github.com/caarlos0/env/v6"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"

	"google.golang.org/grpc"

	"net"
)

var (
	poolP pgxpool.Pool
	//poolM mongo.Client
)

func main() {
	listen, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		defer log.Fatalf("error while listening port: %e", err)
	}
	fmt.Println("Server successfully started on port :50051...")
	key := []byte("super-key")
	cfg := catalog.Config{JwtKey: key}
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatalf("failed to start service, %e", err)
	}
	conn := DBConnection(&cfg)
	fmt.Println("DB successfully connect...")
	/*defer func() {
		poolP.Close()
		if err = poolM.Disconnect(context.Background()); err != nil {
			log.Errorf("cannot disconnect with mongodb")
		}
	}()*/
	ns := grpc.NewServer()
	newService := service.NewService(conn, cfg.JwtKey)
	srv := server.NewServer(newService)
	pb.RegisterPharmacyCatalogServer(ns, srv)

	if err = ns.Serve(listen); err != nil {
		defer log.Fatalf("error while listening server: %e", err)
	}

}

// DBConnection create connection with db
func DBConnection(cfg *catalog.Config) repository.Repository {
	log.Info(cfg.PostgresDBURL)
	poolP, err := pgxpool.Connect(context.Background(), cfg.PostgresDBURL)
	if err != nil {
		log.Fatalf("bad connection with postgresql: %v", err)
		return nil
	}
	return &repository.PRepository{Pool: poolP}
}
