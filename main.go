package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/notnil/chess"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	db        *gorm.DB
	jwtSecret string
}

type MultiStepPuzzle struct {
	ID       string   `json:"id"`
	StartFEN string   `json:"startFen"`
	Moves    []string `json:"moves"`
	FinalFEN string   `json:"finalFen"`
}

type PuzzleResponse struct {
	ID  string `json:"id"`
	FEN string `json:"fen"`
}

type PuzzleRequest struct {
	Move string `json:"move"`
	Step int    `json:"step"`
}

type Result struct {
	ID       string `json:"id"`
	Result   string `json:"result"`
	NextMove string `json:"nextMove,omitempty"`
}

// var puzzles = map[string]MultiStepPuzzle{
// 	"1": {ID: "1", StartFEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Moves: []string{"e2e4", "e7e5", "f1c4", "b8c6", "d1h5", "g8f6", "h5f7"}, FinalFEN: "r1bqkb1r/pppp1Qpp/2n2n2/4p3/2B5/8/PPPP1PPP/RNB1K1NR b KQkq - 0 4"},
// }

func main() {

	dsn := "host=localhost user=postgres password=330090 dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection failed")
	}
	fmt.Println("db is running")

	server := &Server{
		db:        db,
		jwtSecret: "change-me",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	http.HandleFunc("GET /start", server.StartPuzzle)
	http.HandleFunc("POST /puzzle/{id}", server.CheckMove)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func (s *Server) StartPuzzle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// give random id puzzle from map
	randomID := strconv.Itoa(rand.Intn(len(puzzles)) + 1)

	p := puzzles[randomID]

	fenOption, _ := chess.FEN(p.StartFEN)
	game := chess.NewGame(fenOption)

	response := PuzzleResponse{
		ID:  randomID,
		FEN: game.Position().String(),
	}

	json.NewEncoder(w).Encode(response)
}

func (s *Server) CheckMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	puzzle, ok := puzzles[id]
	if !ok {
		http.Error(w, "puzzle not found", http.StatusNotFound)
		return
	}

	var req PuzzleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	// Step index is userStep-1
	expectedMove := puzzle.Moves[req.Step-1]

	// Validate move
	if req.Move != expectedMove {
		json.NewEncoder(w).Encode(Result{
			ID:       id,
			Result:   "you have failed",
			NextMove: expectedMove,
		})
		return
	}

	// If last step → puzzle completed
	if req.Step == len(puzzle.Moves) {
		json.NewEncoder(w).Encode(Result{
			ID:     id,
			Result: "you have won",
		})
		return
	}

	// Return next forced move (computer move)
	json.NewEncoder(w).Encode(Result{
		ID:       id,
		Result:   "correct move",
		NextMove: puzzle.Moves[req.Step],
	})
}
