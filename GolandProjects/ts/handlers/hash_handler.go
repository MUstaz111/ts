package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"hash/crc64"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type HashRequest struct {
	ID     string
	Data   string
	Result string
}

type HashHandler struct {
	requests       map[string]*HashRequest
	mu             sync.Mutex
	timestampMutex sync.Mutex
	timestamp      time.Time
	processing     chan bool
	completed      chan string
	stop           chan bool
	hashFunction   func(string) string
	maxConcurrent  int
	concurrentSem  chan struct{}
}

func NewHashHandler(maxConcurrent int) *HashHandler {
	return &HashHandler{
		requests:      make(map[string]*HashRequest),
		processing:    make(chan bool, maxConcurrent),
		completed:     make(chan string),
		stop:          make(chan bool),
		hashFunction:  calculateHash,
		maxConcurrent: maxConcurrent,
		concurrentSem: make(chan struct{}, maxConcurrent),
	}
}

func (h *HashHandler) HandleCalcHash(w http.ResponseWriter, r *http.Request) {
	// Извлечь данные для хэширования из тела запроса
	var requestBody struct {
		Data string `json:"data"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	data := requestBody.Data

	// Генерировать уникальный идентификатор для заявки
	id := generateID()

	h.mu.Lock()
	h.requests[id] = &HashRequest{
		ID:   id,
		Data: data,
	}
	h.mu.Unlock()

	go h.calculateHash(id)

	fmt.Fprintf(w, id)
}

func (h *HashHandler) HandleGetResult(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.mu.Lock()
	defer h.mu.Unlock()

	req, ok := h.requests[id]
	if !ok {
		http.NotFound(w, r)
		return
	}

	switch req.Result {
	case "":
		fmt.Fprintf(w, "PENDING")
	default:
		fmt.Fprintf(w, req.Result)
	}
}

func (h *HashHandler) Start() {
	go h.processCompletedHashes()
}

func (h *HashHandler) Stop() {
	h.stop <- true
}

func (h *HashHandler) calculateHash(id string) {
	h.concurrentSem <- struct{}{} // Acquire semaphore
	h.processing <- true
	defer func() {
		<-h.processing
		<-h.concurrentSem // Release semaphore
	}()

	req := h.requests[id]

	// Выполнить расчет хэша
	hash := h.hashFunction(req.Data)
	req.Result = hash

	h.completed <- id
}

func (h *HashHandler) processCompletedHashes() {
	for {
		select {
		case id := <-h.completed:
			h.mu.Lock()
			delete(h.requests, id)
			h.mu.Unlock()
		case <-h.stop:
			return
		}
	}
}

func generateID() string {
	// Генерировать уникальный идентификатор для заявки
	id := make([]byte, 16)
	_, err := rand.Read(id)
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(id)
}

func calculateHash(data string) string {
	// Реализовать логику расчета хэша
	crcTable := crc64.MakeTable(crc64.ISO)
	hash := crc64.Checksum([]byte(data), crcTable)
	return strconv.FormatUint(hash, 10)
}
