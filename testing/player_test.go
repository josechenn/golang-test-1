package testing

import (
	"bytes"
	"mime/multipart"
	"myapp/service"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetPlayerByTeamID(t *testing.T) {

	req, err := http.NewRequest("GET", "/get_player", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("team_id", "1")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ReturnPlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetPlayerByPlayerName(t *testing.T) {

	req, err := http.NewRequest("GET", "/get_player", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("player_name", "ronaldo")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ReturnPlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetPlayerNotFound(t *testing.T) {

	req, err := http.NewRequest("GET", "/get_player", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("player_name", "testing")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ReturnPlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetAllPlayer(t *testing.T) {

	req, err := http.NewRequest("GET", "/get_player", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.ReturnPlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestInsertPlayer(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("player_number", "4")
	writer.WriteField("player_name", "xyz")
	writer.WriteField("position", "ST")
	writer.WriteField("team_id", "1")
	writer.Close()

	req, err := http.NewRequest("POST", "/insert_player", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.InsertPlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestUpdatePlayer(t *testing.T) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	writer.WriteField("player_number", "4")
	writer.WriteField("player_name", "xyz")
	writer.WriteField("position", "ST")
	writer.WriteField("team_id", "1")
	writer.WriteField("player_id", "7")
	writer.Close()

	req, err := http.NewRequest("PUT", "/update_player", body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(service.UpdatePlayer)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
