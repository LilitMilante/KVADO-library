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

	authorID := "44a54d7b-6289-4b12-b030-1ffd884763cb" // Lev Tolstoi ID
	expBookCount := 4

	resp, err := client.BooksByAuthorID(context.Background(), &proto.BooksByAuthorRequest{
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

func TestHandler_AuthorsByBookID(t *testing.T) {
	client := NewClient(t)

	bookID := "f3abf142-715a-47a4-83da-4a681e24a278" // The Brothers Karamazov ID
	expAuthorCount := 2

	resp, err := client.AuthorsByBookID(context.Background(), &proto.AuthorsByBookRequest{
		BookId: bookID,
	})
	require.NoError(t, err)

	require.Len(t, resp.Authors, expAuthorCount) // check expected author count

	for _, v := range resp.Authors {
		require.NotEmpty(t, v.Id)        // check id is not empty
		require.NotEmpty(t, v.FirstName) // check first name is not empty
		require.NotEmpty(t, v.LastName)  // check last name is not empty
	}
}

// Create new client
func NewClient(t *testing.T) proto.LibraryClient {
	t.Helper()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8091))
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

	conn, err := grpc.Dial("localhost:8091", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	t.Cleanup(func() {
		err := conn.Close()
		require.NoError(t, err)
	})

	client := proto.NewLibraryClient(conn)

	return client
}

// Connect to database
func initDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := app.ConnectToMySQL("root:dev@tcp(localhost:3306)/library?multiStatements=true")
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	return db
}
