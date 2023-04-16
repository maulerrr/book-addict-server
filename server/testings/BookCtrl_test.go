package testings

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/DTO"
	"github.com/maulerrr/book-addict-server/server/controllers"
	"net/http/httptest"
	"testing"
)

func TestGetAllPosts(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	controllers.GetAllBooks(context)

	assert.Equal(t, 200, context.Writer.Status())
}

func TestAddBook(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	type testcase struct {
		name         string
		payload      interface{}
		expectedCode int
		expectedData interface{}
	}

	testcases := []testcase{
		{
			name:         "Test:Success",
			payload:      DTO.CreateBookDto{Title: "test", Author: "test", Price: 12, Description: "test", Img: "test"},
			expectedCode: 200,
			expectedData: nil,
		},
		{
			name:         "Test: Invalid JSON",
			expectedCode: 400,
			expectedData: gin.H{"code": 400, "message": "Invalid JSON"},
		},
		{},
	}

}

func TestDeleteBookById(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteBookById(tt.args.context)
		})
	}
}

func TestGetAllBooks(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetAllBooks(tt.args.context)
		})
	}
}

func TestGetBookById(t *testing.T) {
	type args struct {
		context *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetBookById(tt.args.context)
		})
	}
}
