package testing

import (
	"myapp/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCountTotalBox(t *testing.T) {

	req, err := http.NewRequest("GET", "/count_total", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.CountTotalBox)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
