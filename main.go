package main

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// AddService defines the business logic interface for the add operation.
// It provides a method to add a new advertisement and generate a bid price.
type AddService interface {
	Add(uuid.UUID) (uuid.UUID, float64, error)
}

// loggingMiddleware is a middleware for logging requests and responses.
type loggingMiddleware struct {
	next AddService
}

// NewLoggingMiddleware wraps an AddService with logging functionality.
func NewLoggingMiddleware(svc AddService) AddService {
	return &loggingMiddleware{
		next: svc,
	}
}

// Add logs the request and response details, along with the execution time.
func (lm *loggingMiddleware) Add(id uuid.UUID) (uuid.UUID, float64, error) {
	defer func(start time.Time) {
		slog.Info(
			"addrequest",
			"id", id,
			"bidPrice", bidPrice,
			"err", err,
			"took", time.Since(start),
		)
	}(time.Now())
	return lm.next.Add(id)
}

// addService implements the AddService interface and contains the core business logic.
type addService struct {
}

// NewAddService creates and returns a new instance of addService.
func NewAddService() AddService {
	return &addService{}
}

// Add generates a random bid price and returns it along with a new UUID.
func (addService *addService) Add(id uuid.UUID) (uuid.UUID, float64, error) {
	n := rand.Float64()
	return uuid.New(), 69.69 + n, nil
}

// AddRequest represents the JSON structure for incoming requests.
type AddRequest struct {
	AdPlacementID uuid.UUID `json:"addPlacementID"`
}

// AddResponse represents the JSON structure for responses.
type AddResponse struct {
	AddID    uuid.UUID `json:"addID"`
	BidPrice float64   `json:"bidPrice"`
}

func main() {
	// Initialize the HTTP router.
	router := http.NewServeMux()

	// Create the service with logging middleware.
	svc := NewLoggingMiddleware(NewAddService())

	// Create a handler for the add request.
	h := NewAddrequestHandler(svc)

	// Register the handler with the router.
	router.HandleFunc("/add", makeHandler(h.handleAddRequest))

	// Start the HTTP server on port 3000.
	http.ListenAndServe(":3000", router)
}

// addRequestHandler handles HTTP requests for the add service.
type addRequestHandler struct {
	svc AddService
}

// NewAddrequestHandler creates a new addRequestHandler instance.
func NewAddrequestHandler(svc AddService) *addRequestHandler {
	return &addRequestHandler{
		svc: svc,
	}
}

// handleAddRequest processes incoming HTTP requests and returns responses.
func (h addRequestHandler) handleAddRequest(w http.ResponseWriter, r *http.Request) error {
	// Generate a new UUID for the request.
	addID := uuid.New()

	// Call the business logic through the AddService interface.
	id, bidPrice, err := h.svc.Add(addID)
	if err != nil {
		slog.Error("add service returned non 200 response", "err", err)
		return writeJSON(w, http.StatusNoContent, nil)
	}

	// Create a response object with the generated data.
	resp := AddResponse{
		AddID:    id,
		BidPrice: bidPrice,
	}

	// Write the response as JSON.
	return writeJSON(w, http.StatusOK, resp)
}

// APIfunc defines a function signature for HTTP handlers with error handling.
type APIfunc func(w http.ResponseWriter, r *http.Request) error

// makeHandler wraps an APIfunc to handle errors uniformly.
func makeHandler(h APIfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			slog.Error("API error", "err", err, "path", r.URL.Path)
		}
	}
}

// writeJSON encodes the response object as JSON and writes it to the response writer.
func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(&v)
}
