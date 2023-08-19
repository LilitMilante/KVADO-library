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

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestHandler_BooksByAuthor(t *testing.T) {
	client := NewClient(t)

	// Lev Tolstoi ID
	authorID := "44a54d7b-6289-4b12-b030-1ffd884763cb"
	expBookCount := 4

	resp, err := client.BooksByAuthor(context.Background(), &proto.BooksByAuthorRequest{
		AuthorId: authorID,
	})
	require.NoError(t, err)

	require.Len(t, resp.Books, expBookCount) // check expected book count

	for _, v := range resp.Books {
		require.NotEmpty(t, v.Id)                  // check id is not empty
		require.NotEmpty(t, v.Title)               // check title is not empty
		require.Contains(t, v.AuthorIds, authorID) // check authorIDs contains authorID
	}
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

	db, err := app.ConnectToMySQL("root:dev@tcp(localhost:3306)/library?multiStatements=true")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	return db
}
