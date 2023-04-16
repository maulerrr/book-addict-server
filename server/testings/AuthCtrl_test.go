package testings

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maulerrr/book-addict-server/server/DB"
	"github.com/maulerrr/book-addict-server/server/DTO"
	"github.com/maulerrr/book-addict-server/server/controllers"
	"github.com/maulerrr/book-addict-server/server/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestLogin(t *testing.T) {
	DB.ConnectDB()

	type response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	type testcase struct {
		name     string
		payload  interface{}
		expected response
	}

	testcases := []testcase{
		{
			name:    "Test:Success",
			payload: DTO.LoginDto{Email: "test@mail.ru", Password: "test"},
			expected: response{
				Code:    200,
				Message: "success",
				Data: models.TokenResponse{
					UserID:   1,
					FullName: "test",
					Email:    "test@mail.ru",
				},
			},
		},
		{
			name:    "Test: User does not exist",
			payload: DTO.LoginDto{Email: "unknown@unknown.ru", Password: "unknown"},
			expected: response{
				Code:    404,
				Message: "User not found",
			},
		}, {
			name: "Test:Invalid JSON",
			expected: response{
				Code:    400,
				Message: "Invalid JSON",
			},
		},
		{
			name:    "Test: Incorrect Password",
			payload: DTO.LoginDto{Email: "test@mail.ru", Password: "testtest"},
			expected: response{
				Code:    401,
				Message: "Password is not correct",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			loginJSON, _ := json.Marshal(tc.payload)
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			request := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(loginJSON))
			if tc.payload == nil {
				request = httptest.NewRequest(http.MethodPost, "/auth/signup", nil)
			}
			context.Request = request
			controllers.Login(context)
			if recorder.Code != tc.expected.Code {
				t.Errorf("Expected %d but got %d", tc.expected.Code, recorder.Code)
			}
			var response response
			if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
				t.Errorf("Error unmarshalling response: %v", err)
			}
			if reflect.ValueOf(response.Data).Kind() == reflect.Map {
				dataMap := response.Data
				dataStruct := models.TokenResponse{}
				jsonData, _ := json.Marshal(dataMap)
				json.Unmarshal(jsonData, &dataStruct)
				dataStruct.Token = ""
				response.Data = dataStruct
			}

			if !reflect.DeepEqual(response, tc.expected) {
				t.Errorf("Expected response is %v but got %v", tc.expected, response)
			}
		})
	}
}

func TestSignup(t *testing.T) {
	DB.ConnectDB()
	type response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	type testcase struct {
		name     string
		payload  interface{}
		expected response
	}
	testcases := []testcase{
		{
			name:    "Test:Success",
			payload: DTO.SignupDto{FullName: "test", Email: "test@mail.ru", Password: "test"},
			expected: response{
				Code:    200,
				Message: "success",
				Data: models.TokenResponse{
					UserID:   2,
					FullName: "test",
					Email:    "test@mail.ru",
				},
			},
		},
		{
			name:    "Test: User already exists",
			payload: DTO.SignupDto{FullName: "test", Email: "test@mail.ru", Password: "test"},
			expected: response{
				Code:    400,
				Message: "User already exists",
			},
		}, {
			name: "Test:Invalid JSON",
			expected: response{
				Code:    400,
				Message: "Invalid JSON",
			},
		},
		{
			name:    "Test: Invalid Email Address",
			payload: DTO.SignupDto{FullName: "test1", Email: "test1", Password: "testtest"},
			expected: response{
				Code:    400,
				Message: "Email is not valid",
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			loginJSON, _ := json.Marshal(tc.payload)
			recorder := httptest.NewRecorder()
			context, _ := gin.CreateTestContext(recorder)

			request := httptest.NewRequest(http.MethodPost, "/auth/signup", bytes.NewBuffer(loginJSON))
			if tc.payload == nil {
				request = httptest.NewRequest(http.MethodPost, "/auth/signup", nil)
			}
			context.Request = request
			controllers.Signup(context)
			if recorder.Code != tc.expected.Code {
				t.Errorf("Expected %d but got %d", tc.expected.Code, recorder.Code)
			}
			var response response
			if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
				t.Errorf("Error unmarshalling response: %v", err)
			}
			if reflect.ValueOf(response.Data).Kind() == reflect.Map {
				dataMap := response.Data
				dataStruct := models.TokenResponse{}
				jsonData, _ := json.Marshal(dataMap)
				json.Unmarshal(jsonData, &dataStruct)
				dataStruct.Token = ""
				response.Data = dataStruct
			}

			if !reflect.DeepEqual(response, tc.expected) {
				t.Errorf("Expected response is %v but got %v", tc.expected, response)
			}
		})
	}
}
