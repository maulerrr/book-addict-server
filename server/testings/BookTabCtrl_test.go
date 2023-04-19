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

func TestGetAllBookTabs(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	controllers.GetAllBookTabs(context)

	assert.Equal(t, 200, context.Writer.Status())
}

func TestAddBookTab(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	newBookTab := DTO.CreateBookTabDto{
		BookID:   1,
		UserID:   1,
		Finished: true,
	}
	postJSON, _ := json.Marshal(newBookTab)

	request, _ := http.NewRequest(http.MethodPost, "/booktabs/add", bytes.NewBuffer(postJSON))
	request.Header.Set("Content-Type", "application/json")
	context.Request = request
	controllers.AddBookTab(context)
	assert.Equal(t, 200, context.Writer.Status())
	var book models.Book
	DB.DB.Last(&book)
	assert2.Greater(t, book.ID, 0)
}

func TestDeleteBookTabById(t *testing.T) {
	DB.ConnectDB()
	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	newBookTab := DTO.CreateBookTabDto{
		BookID:   1,
		UserID:   1,
		Finished: true,
	}

	DB.DB.Create(&newBookTab)
	controllers.DeleteBookTabById(context)

	var booktab models.BookTab
	DB.DB.First(&booktab, newBookTab)
}

func TestGetFinished(t *testing.T) {
	DB.ConnectDB()

	type testcase struct {
		name     string
		param    gin.Param
		expected int
		finished bool
	}
	testcases := []testcase{
		{
			name:     "Test:Success",
			param:    gin.Param{Key: "id", Value: strconv.Itoa(2)},
			expected: 200,
			finished: true,
		},
		{
			name:     "Test: Invalid ID",
			param:    gin.Param{Key: "id", Value: ""},
			expected: 400,
			finished: true,
		},
		{
			name:     "Test:Not found",
			param:    gin.Param{Key: "id", Value: strconv.Itoa(2)},
			expected: 404,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)
			context.Params = append(context.Params, tc.param)
			controllers.GetFinishedBooks(context)
			assert.Equal(t, tc.expected, recorder.Code)
		})
	}
}
