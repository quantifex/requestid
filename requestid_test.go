package requestid

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const testXRequestID = "test-request-id"

func emptySuccessResponse(c *gin.Context) {
	c.String(200, "")
}

func Test_RequestID_CreateNew(t *testing.T) {
	r := gin.New()
	r.Use(New())
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, w.Header().Get(headerXRequestID))
}

func Test_RequestID_PassThru(t *testing.T) {
	r := gin.New()
	r.Use(New())
	r.GET("/", emptySuccessResponse)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set(headerXRequestID, testXRequestID)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, testXRequestID, w.Header().Get(headerXRequestID))
}
