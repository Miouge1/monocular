package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arschles/assert"
	"github.com/go-openapi/runtime"
	"github.com/kubernetes-helm/monocular/src/api/swagger/restapi/operations"
)

func TestHealthz(t *testing.T) {
	w := httptest.NewRecorder()
	params := operations.HealthzParams{}
	resp := Healthz(params)
	assert.NotNil(t, resp, "Healthz response")
	resp.WriteResponse(w, runtime.JSONProducer())
	assert.Equal(t, w.Code, http.StatusOK, "expect a 200 response code")
}
