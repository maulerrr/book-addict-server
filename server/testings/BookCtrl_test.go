package testings

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/DTO"
	"github.com/maulerrr/book-addict-server/server/controllers"
	"github.com/maulerrr/book-addict-server/server/models"
	assert2 "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetAllBooks(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	controllers.GetAllBooks(context)

	assert.Equal(t, 200, context.Writer.Status())
}

func TestAddBook(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	newBook := DTO.CreateBookDto{
		Title:       "test title",
		Author:      "test author",
		Price:       1,
		Description: "test description",
		Img:         "test img",
	}
	postJSON, _ := json.Marshal(newBook)

	request, _ := http.NewRequest(http.MethodPost, "/readlist", bytes.NewBuffer(postJSON))
	request.Header.Set("Content-Type", "application/json")
	context.Request = request
	controllers.AddBook(context)
	assert.Equal(t, 200, context.Writer.Status())
	var book models.Book
	DB.DB.Last(&book)
	assert2.Greater(t, book.ID, 0)
}

func TestDeleteBookById(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	newBook := DTO.CreateBookDto{
		Title:       "test title",
		Author:      "test author",
		Price:       1,
		Description: "test description",
		Img:         "test img",
	}
	DB.DB.Create(&newBook)
	controllers.DeleteBookById(context)

	var book models.Book
	DB.DB.First(&book, newBook)
}

func TestGetBookById(t *testing.T) {
	DB.ConnectDB()

	type testcase struct {
		name     string
		param    gin.Param
		expected int
	}
	testcases := []testcase{
		{
			name:     "Test:Success",
			param:    gin.Param{Key: "id", Value: strconv.Itoa(1)},
			expected: 200,
		},
		{
			name:     "Test: Invalid ID",
			param:    gin.Param{Key: "id", Value: ""},
			expected: 400,
		},
		{
			name:     "Test:Not found",
			param:    gin.Param{Key: "id", Value: strconv.Itoa(-1)},
			expected: 404,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Params = append(context.Params, tc.param)
			controllers.GetBookById(context)
			assert.Equal(t, tc.expected, recorder.Code)
		})
	}
}
