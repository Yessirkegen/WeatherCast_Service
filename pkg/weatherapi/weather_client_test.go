package weatherapi

/*
import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetWeather(t *testing.T) {
	// Создаем тестовый сервер
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"name":"London","main":{"temp":280.32,"pressure":1012,"humidity":81},"weather":[{"description":"light rain"}]}`))
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	client := &weatherClient{
		apiKey:  "test_api_key",
		baseURL: server.URL,
	}

	weather, err := client.GetWeather("London")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if weather.Name != "London" {
		t.Errorf("Expected 'London', got %s", weather.Name)
	}
	if weather.Main.Temp != 280.32 {
		t.Errorf("Expected temp 280.32, got %f", weather.Main.Temp)
	}
}
*/
