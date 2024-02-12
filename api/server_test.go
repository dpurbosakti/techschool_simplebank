package api

// import (
// 	"context"
// 	"encoding/json"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	mockdb "simple-bank/db/mock"
// 	db "simple-bank/db/sqlc"
// 	"simple-bank/util"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang/mock/gomock"
// 	"github.com/jackc/pgx/v5/pgxpool"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// type healthCheck struct {
// 	Message string `json:"message"`
// }

// func TestNewServer(t *testing.T) {
// 	config, err := util.LoadConfig("../")
// 	require.NoError(t, err)

// 	conn, err := pgxpool.New(context.Background(), config.DBSource)
// 	require.NoError(t, err)

// 	store := db.NewStore(conn)

// 	config.TokenSymmetricKey = "1"

// 	_, err = NewServer(config, store)
// 	require.Error(t, err)
// }

// func TestHealthCheckRoute(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	store := mockdb.NewMockStore(ctrl)
// 	server := NewTestServer(t, store)

// 	recorder := httptest.NewRecorder()

// 	url := "/health-check"
// 	request, err := http.NewRequest(http.MethodGet, url, nil)
// 	require.NoError(t, err)
// 	server.router.ServeHTTP(recorder, request)

// 	data, err := io.ReadAll(recorder.Body)
// 	require.NoError(t, err)

// 	var gotHealthCheck healthCheck

// 	err = json.Unmarshal(data, &gotHealthCheck)
// 	require.NoError(t, err)
// 	require.Equal(t, "ok", gotHealthCheck.Message)

// }

// func TestServerStart(t *testing.T) {
// 	// Create a new instance of your server
// 	server := &Server{
// 		router: gin.Default(),
// 	}

// 	// Define a test route
// 	server.router.GET("/test", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{"message": "Test route"})
// 	})

// 	// Create a test HTTP server using httptest
// 	testServer := httptest.NewServer(server.router)

// 	// Close the test server when the test is done
// 	defer testServer.Close()

// 	// Run the Start function in a goroutine
// 	go func() {
// 		err := server.Start("0.0.0.0:8080")
// 		assert.NoError(t, err, "Expected no error when starting the server")
// 	}()

// 	// // Make a sample HTTP request to the test server
// 	// resp, err := http.Get(testServer.URL + "/test")
// 	// assert.NoError(t, err, "Expected no error when making an HTTP request")
// 	// defer resp.Body.Close()

// 	// Check the HTTP response status code
// 	// assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200")
// }
