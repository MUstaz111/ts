package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"tsarka/models"
	"tsarka/repositories"
)

func FindSubstringHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		String string `json:"string"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	substring := FindMaxSubstring(requestBody.String)

	response := struct {
		Substring string `json:"substring"`
	}{
		Substring: substring,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CheckEmailHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Text string `json:"text"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	emails := FindEmails(requestBody.Text)

	response := struct {
		Emails []string `json:"emails"`
	}{
		Emails: emails,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *Counter) CounterAddHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	incrementStr := vars["increment"]
	increment := 0

	fmt.Sscanf(incrementStr, "%d", &increment)

	err := c.Add(r.Context(), increment)
	if err != nil {
		http.Error(w, "Failed to add to counter", http.StatusInternalServerError)
		return
	}

	response := struct {
		Value int `json:"value"`
	}{
		Value: c.Value(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *Counter) CounterSubHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	decrementStr := vars["decrement"]
	decrement := 0

	fmt.Sscanf(decrementStr, "%d", &decrement)

	err := c.Sub(r.Context(), decrement)
	if err != nil {
		http.Error(w, "Failed to subtract from counter", http.StatusInternalServerError)
		return
	}

	response := struct {
		Value int `json:"value"`
	}{
		Value: c.Value(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *Counter) CounterValHandler(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Value int `json:"value"`
	}{
		Value: c.Value(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CheckIINHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Text string `json:"text"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	iins := FindIINs(requestBody.Text)

	response := struct {
		IINs []string `json:"iins"`
	}{
		IINs: iins,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateUserHandler обрабатывает POST-запрос для создания нового пользователя
func CreateUserHandler(repo *repositories.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		id, err := repo.CreateUser(&user)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		response := struct {
			ID int `json:"id"`
		}{
			ID: id,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

// GetUserByIDHandler обрабатывает GET-запрос для получения информации о пользователе по его ID
func GetUserByIDHandler(repo *repositories.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := repo.GetUserByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

// UpdateUserHandler обрабатывает PUT-запрос для обновления информации о пользователе
func UpdateUserHandler(repo *repositories.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		user.ID = userID

		err = repo.UpdateUser(&user)
		if err != nil {
			http.Error(w, "Failed to update user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

// DeleteUserHandler обрабатывает DELETE-запрос для удаления пользователя по его ID
func DeleteUserHandler(repo *repositories.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id := params["id"]

		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		err = repo.DeleteUser(userID)
		if err != nil {
			http.Error(w, "Failed to delete user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func FindIdentifiersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	str := params["str"]

	identifiers := FindIdentifiers(str)

	response := struct {
		Identifiers []string `json:"identifiers"`
	}{
		Identifiers: identifiers,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
