package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tsarka/db"
	"tsarka/handlers"
	"tsarka/repositories"
)

const maxConcurrent = 5 // Максимальное количество одновременно вычисляющихся хешей

func main() {
	// Создание клиента Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // Пароль для доступа к Redis, если требуется
		DB:       0,  // Индекс базы данных Redis, который будет использоваться
	})

	// Создание экземпляра базы данных
	dbInstance, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer dbInstance.Close()

	// Создание экземпляра UserRepository
	userRepo := repositories.NewUserRepository(dbInstance)

	// Проверка подключения к Redis
	err = redisClient.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	// Создание экземпляра счетчика с указанным клиентом Redis
	counter := handlers.NewCounter(redisClient)

	router := mux.NewRouter()
	hashHandler := handlers.NewHashHandler(maxConcurrent)
	hashHandler.Start()
	http.HandleFunc("/rest/hash/calc", hashHandler.HandleCalcHash)
	http.HandleFunc("/rest/hash/result/", hashHandler.HandleGetResult)
	router.HandleFunc("/rest/substr/find", handlers.FindSubstringHandler).Methods("POST")
	router.HandleFunc("/rest/email/check", handlers.CheckEmailHandler).Methods("POST")
	router.HandleFunc("/rest/iin/check", handlers.CheckIINHandler).Methods("POST")
	router.HandleFunc("/rest/counter/add/{increment}", counter.CounterAddHandler).Methods("POST")
	router.HandleFunc("/rest/counter/sub/{decrement}", counter.CounterSubHandler).Methods("POST")
	router.HandleFunc("/rest/counter/val", counter.CounterValHandler).Methods("GET")
	router.HandleFunc("/rest/user", handlers.CreateUserHandler(userRepo)).Methods("POST")
	router.HandleFunc("/rest/user/{id}", handlers.GetUserByIDHandler(userRepo)).Methods("GET")
	router.HandleFunc("/rest/user/{id}", handlers.UpdateUserHandler(userRepo)).Methods("PUT")
	router.HandleFunc("/rest/user/{id}", handlers.DeleteUserHandler(userRepo)).Methods("DELETE")
	router.HandleFunc("/rest/self/find/{str}", handlers.FindIdentifiersHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
