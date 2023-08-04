package handler

import (
	"github.com/devops-codegpt/server/config"
	"github.com/devops-codegpt/server/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHealthCheck(t *testing.T) {
	e, container := test.PrepareForHandlerTest()

	req := test.NewJSONRequest(http.MethodGet, config.APIHealth, http.NoBody)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	health := NewHealthHandler(container)

	data := map[string]any{
		"code": 200,
		"msg":  "success",
		"ret":  "healthy",
	}

	if assert.NoError(t, health.GetHealthCheck(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, test.ConvertToString(data), rec.Body.String())
	}
}
