package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/notnil/chess"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	database  *gorm.DB
	jwtSecret string
}

type Puzzle struct {
	ID       string `json:"id"`
	StartFEN string `json:"startFen"`
	TrueFEN  string `json:"trueFen"`
}

var fens = map[string]Puzzle{
	"1": {
		ID:       "1",
		StartFEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		TrueFEN:  "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b KQkq e3 0 1", // e2-e4
	},

	"2": {
		ID:       "2",
		StartFEN: "r1bqkbnr/pppp1ppp/2n5/4p3/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 2 3",
		TrueFEN:  "r1bqkbnr/pppp1ppp/2n5/4p3/3PP3/5N2/PPP2PPP/RNBQKB1R b KQkq d3 0 3", // d2-d4
	},

	"3": {
		ID:       "3",
		StartFEN: "rnbqk2r/pppp1ppp/4pn2/8/2B1P3/5N2/PPPP1PPP/RNBQ1RK1 w kq - 4 5",
		TrueFEN:  "rnbqk2r/pppp1ppp/4pn2/8/2B1P3/5N2/PPPP1PPP/RNBQ1R1K b kq - 5 5", // Kg1-h1
	},

	"4": {
		ID:       "4",
		StartFEN: "r1bqkbnr/pppp1ppp/2n5/4p3/3PP3/8/PPP2PPP/RNBQKBNR b KQkq - 1 3",
		TrueFEN:  "r1bqkb1r/pppp1ppp/2n2n2/4p3/3PP3/8/PPP2PPP/RNBQKBNR w KQkq - 2 4", // Ng8-f6
	},

	"5": {
		ID:       "5",
		StartFEN: "rnbq1bnr/pppp1kpp/8/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQ - 4 5",
		TrueFEN:  "rnbq1bnr/pppp1kpp/8/4p3/2B1P3/4PN2/PPPP1PPP/RNBQK2R b KQ - 5 5", // Nf3-e5
	},

	"6": {
		ID:       "6",
		StartFEN: "r3k2r/ppp2ppp/2n1b3/3q4/3P4/2N1BP2/PPP3PP/R2Q1RK1 w kq - 0 12",
		TrueFEN:  "r3k2r/ppp2ppp/2n1b3/3q4/3P4/2N1BP2/PPP3PP/R2Q1R1K b kq - 1 12", // Kg1-h1
	},

	"7": {
		ID:       "7",
		StartFEN: "r1bqkbnr/pppp1ppp/2n5/8/2B1p3/5N2/PPPP1PPP/RNBQK2R w KQkq - 4 4",
		TrueFEN:  "r1bqkbnr/pppp1ppp/2n5/8/2B1p3/4PN2/PPPP1PPP/RNBQK2R b KQkq e3 0 4", // e2-e3
	},

	"8": {
		ID:       "8",
		StartFEN: "rnbqkbnr/pppp1ppp/8/4p3/4P3/3P4/PPP2PPP/RNBQKBNR b KQkq - 0 2",
		TrueFEN:  "rnbqkbnr/pppp1ppp/8/8/4p3/3P4/PPP2PPP/RNBQKBNR w KQkq - 0 3", // e5-e4
	},

	"9": {
		ID:       "9",
		StartFEN: "r1bqk2r/pppp1ppp/2n2n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQ1RK1 w kq - 6 6",
		TrueFEN:  "r1bqk2r/pppp1ppp/2n2n2/1B2p3/4P3/5N2/PPPP1PPP/RNBQR1K1 b kq - 7 6", // Rf1-e1
	},

	"10": {
		ID:       "10",
		StartFEN: "rnb1kbnr/ppppqppp/8/4p3/2B1P3/2N5/PPPP1PPP/R1BQK1NR w KQkq - 2 4",
		TrueFEN:  "rnb1kbnr/ppppqppp/8/4p3/2B1P3/2N2N2/PPPP1PPP/R1BQK2R b KQkq - 3 4", // Ng1-f3
	},
}

type MultiStepPuzzle struct {
	ID       string   `json:"id"`
	StartFEN string   `json:"startfen"`
	Moves    []string `json:"moves"`
	FinalFEN string   `json:"finalfen"`
}

var advancePuzzles = map[string]MultiStepPuzzle{
	"1": {
		ID:       "1",
		StartFEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1",
		Moves:    []string{"e2e4", "e7e5", "f1c4", "b8c6", "d1h5", "g8f6", "h5f7"},
		FinalFEN: "r1bqkb1r/pppp1Qpp/2n2n2/4p3/2B5/8/PPPP1PPP/RNB1K1NR b KQkq - 0 4",
	},
}

type PuzzleResponse struct {
	ID           string `json:"id"`
	FenStr       string `json:"fenstr"`
	BoardDrawing string `json:"boardDrawing"`
}

type PuzzleRequest struct {
	FenStr string `json:"fenstr"`
	// AutoNextEnabled bool   `json:"AutoNextEnabled" optional`
}

type Result struct {
	ID     string `json:"id"`
	Result string `json:"result"`
}

func main() {

	dsn := "host=localhost user=postgres password=330090 dbname=postgres port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("db failed to connect")
		panic(err)
	}
	fmt.Println("db is running")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	server := &Server{
		database:  db,
		jwtSecret: "change-me",
	}

	http.HandleFunc("GET /start", server.StartGame)
	http.HandleFunc("POST /puzzle/{id}", server.CheckMove)
	fmt.Printf("Server is running on https://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func (s *Server) StartGame(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-type", "application/json")

	// // Get a random puzzle
	// randomID := strconv.Itoa(rand.Intn(len(fens)) + 1)
	// selectedPuzzle := fens[randomID]

	// fen, _ := chess.FEN(selectedPuzzle.StartFEN)

	// game := chess.NewGame(fen)

	// json.NewEncoder(w).Encode(PuzzleResponse{
	// 	ID:           selectedPuzzle.ID,
	// 	FenStr:       game.Position().String(),
	// 	BoardDrawing: game.Position().Board().Draw(),
	// })

	fen, _ := chess.FEN(advancePuzzles["1"].StartFEN)
	game := chess.NewGame(fen)

	json.NewEncoder(w).Encode(PuzzleResponse{
		ID:     "1",
		FenStr: game.Position().String(),
	})
}

func (s *Server) CheckMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")
	puzzle, ok := fens[id]
	if !ok {
		http.Error(w, "Puzzle not found", http.StatusNotFound)
		return
	}

	var p PuzzleRequest
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	correctFen := puzzle.TrueFEN

	if p.FenStr == correctFen {
		json.NewEncoder(w).Encode(Result{
			ID:     id,
			Result: "you have won",
		})
		return
	}

	json.NewEncoder(w).Encode(Result{
		ID:     id,
		Result: "you have failed",
	})
}
