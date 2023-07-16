package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindSubstringHandler(t *testing.T) {
	// Создаем тестовый запрос с телом запроса
	requestBody := struct {
		String string `json:"string"`
	}{
		String: "abcabc",
	}
	requestBodyBytes, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest("POST", "/rest/substr/find", bytes.NewReader(requestBodyBytes))

	// Создаем ResponseRecorder (реализация http.ResponseWriter) для записи ответа
	responseRecorder := httptest.NewRecorder()

	// Вызываем обработчик
	FindSubstringHandler(responseRecorder, request)

	// Проверяем код статуса
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
	}

	// Проверяем тело ответа
	var responseBody struct {
		Substring string `json:"substring"`
	}
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
	expectedSubstring := "abc"
	if responseBody.Substring != expectedSubstring {
		t.Errorf("Expected substring '%s', but got '%s'", expectedSubstring, responseBody.Substring)
	}
}

func TestCheckEmailHandler(t *testing.T) {
	// Создаем тестовый запрос с телом запроса
	requestBody := struct {
		Text string `json:"text"`
	}{
		Text: "Email: test@example.com Email: invalid",
	}
	requestBodyBytes, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest("POST", "/rest/email/check", bytes.NewReader(requestBodyBytes))

	// Создаем ResponseRecorder (реализация http.ResponseWriter) для записи ответа
	responseRecorder := httptest.NewRecorder()

	// Вызываем обработчик
	CheckEmailHandler(responseRecorder, request)

	// Проверяем код статуса
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
	}

	// Проверяем тело ответа
	var responseBody struct {
		Emails []string `json:"emails"`
	}
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
	expectedEmails := []string{"test@example.com"}
	if len(responseBody.Emails) != len(expectedEmails) {
		t.Errorf("Expected %d emails, but got %d", len(expectedEmails), len(responseBody.Emails))
	} else {
		for i, email := range responseBody.Emails {
			if email != expectedEmails[i] {
				t.Errorf("Expected email '%s', but got '%s'", expectedEmails[i], email)
			}
		}
	}
}

func TestCounterValHandler(t *testing.T) {
	// Создаем тестовый запрос
	request, _ := http.NewRequest("GET", "/rest/counter/val", nil)

	// Создаем ResponseRecorder (реализация http.ResponseWriter) для записи ответа
	responseRecorder := httptest.NewRecorder()

	// Создаем экземпляр счетчика для теста
	counter := &Counter{}

	// Вызываем обработчик
	counter.CounterValHandler(responseRecorder, request)

	// Проверяем код статуса
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
	}

	// Проверяем тело ответа
	var responseBody struct {
		Value int `json:"value"`
	}
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
	expectedValue := 0
	if responseBody.Value != expectedValue {
		t.Errorf("Expected counter value %d, but got %d", expectedValue, responseBody.Value)
	}
}

func TestCheckIINHandler(t *testing.T) {
	// Создаем тестовый запрос с телом запроса
	requestBody := struct {
		Text string `json:"text"`
	}{
		Text: "IIN: 123456789012 IIN: 987654321098",
	}
	requestBodyBytes, _ := json.Marshal(requestBody)
	request, _ := http.NewRequest("POST", "/rest/iin/check", bytes.NewReader(requestBodyBytes))

	// Создаем ResponseRecorder (реализация http.ResponseWriter) для записи ответа
	responseRecorder := httptest.NewRecorder()

	// Вызываем обработчик
	CheckIINHandler(responseRecorder, request)

	// Проверяем код статуса
	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, responseRecorder.Code)
	}

	// Проверяем тело ответа
	var responseBody struct {
		IINs []string `json:"iins"`
	}
	_ = json.Unmarshal(responseRecorder.Body.Bytes(), &responseBody)
	expectedIINs := []string{"123456789012", "987654321098"}
	if len(responseBody.IINs) != len(expectedIINs) {
		t.Errorf("Expected %d IINs, but got %d", len(expectedIINs), len(responseBody.IINs))
	} else {
		for i, iin := range responseBody.IINs {
			if iin != expectedIINs[i] {
				t.Errorf("Expected IIN '%s', but got '%s'", expectedIINs[i], iin)
			}
		}
	}
}
