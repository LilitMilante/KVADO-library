package api

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"testing"

	"KVADO-library/gen/proto"
	"KVADO-library/internal/app"
	"KVADO-library/internal/repository"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestHandler_BooksByAuthor(t *testing.T) {
	client := NewClient(t)

	_, err := client.BooksByAuthor(context.Background(), &proto.BooksByAuthorRequest{
		AuthorId: uuid.NewString(),
	})
	require.NoError(t, err)
}

func NewClient(t *testing.T) proto.LibraryClient {
	t.Helper()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	require.NoError(t, err)

	db := initDB(t)
	repo := repository.NewRepository(db)
	h := NewHandler(repo)
	srv := NewServer(h)
	t.Cleanup(srv.GracefulStop)

	go func() {
		err = srv.Serve(lis)
		require.NoError(t, err)
	}()

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() {
		err := conn.Close()
		require.NoError(t, err)
	})

	client := proto.NewLibraryClient(conn)

	return client
}

func initDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := app.ConnectToMySQL("root:dev@tcp(localhost:3306)/library")
	require.NoError(t, err)
	t.Cleanup(func() {
		_, err := db.Exec("TRUNCATE TABLE books;")
		require.NoError(t, err)
		_, err = db.Exec("TRUNCATE TABLE authors;")
		require.NoError(t, err)

		require.NoError(t, db.Close())
	})

	return db
}
