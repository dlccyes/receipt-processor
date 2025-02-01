package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func SetupHttpTest() (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	return c, w
}

func SetCtxRequestBody(c *gin.Context, v interface{}) {
	c.Request = &http.Request{
		Body: io.NopCloser(bytes.NewBuffer(marshalStructToJSON(v))),
	}
}

func marshalStructToJSON(v interface{}) []byte {
	jsonData, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return jsonData
}
