package main

import (
	"fmt"
	"github.com/AHacTacIA/KP/PhCatalog/internal/catalog"
	"github.com/AHacTacIA/KP/PhCatalog/internal/repository"
	"github.com/AHacTacIA/KP/PhCatalog/internal/server"
	"github.com/AHacTacIA/KP/PhCatalog/internal/service"
	pb "github.com/AHacTacIA/KP/PhCatalog/proto"

	/*"github.com/Egor-Tihonov/GRPC/internal/model"
	"github.com/Egor-Tihonov/GRPC/internal/repository"
	"github.com/Egor-Tihonov/GRPC/internal/server"
	"github.com/Egor-Tihonov/GRPC/internal/service"
	pb "github.com/Egor-Tihonov/GRPC/proto"*/

	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"

	"net"
)

var (
	poolP pgxpool.Pool
	poolM mongo.Client
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
	defer func() {
		poolP.Close()
		if err = poolM.Disconnect(context.Background()); err != nil {
			log.Errorf("cannot disconnect with mongodb")
		}
	}()
	ns := grpc.NewServer()
	newService := service.NewService(conn, cfg.JwtKey)
	srv := server.NewServer(newService)
	pb.RegisterCRUDServer(ns, srv)

	if err = ns.Serve(listen); err != nil {
		defer log.Fatalf("error while listening server: %e", err)
	}

}

// DBConnection create connection with db
func DBConnection(cfg *model.Config) repository.Repository {
	switch cfg.CurrentDB {
	case "postgres":
		log.Info(cfg.PostgresDBURL)
		poolP, err := pgxpool.Connect(context.Background(), cfg.PostgresDBURL) // "postgresql://postgres:123@localhost:5432/person"
		if err != nil {
			log.Fatalf("bad connection with postgresql: %v", err)
			return nil
		}
		return &repository.PRepository{Pool: poolP}

	case "mongo":
		poolM, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
		if err != nil {
			log.Fatalf("bad connection with mongoDb: %v", err)
			return nil
		}
		return &repository.MRepository{Pool: poolM}
	}
	return nil
}
