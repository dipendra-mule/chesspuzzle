package main

var puzzles = map[string]MultiStepPuzzle{
	"1": {
		ID:       "1",
		StartFEN: "r1bqkbnr/pppp1ppp/2n5/4p3/2B1P3/5N2/PPPP1PPP/RNBQK2R w KQkq - 2 3",
		Moves:    []string{"f3g5", "d8g5", "c4f7"},
		FinalFEN: "r1b1kbnr/pppp1Bpp/2n5/4p3/4P3/8/PPPP1PPP/RNBQK2R b KQkq - 0 3",
	},

	"2": {
		ID:       "2",
		StartFEN: "rnb1kbnr/ppp2ppp/3q4/3Np3/4P3/8/PPPP1PPP/R1BQKBNR w KQkq - 0 4",
		Moves:    []string{"d5c7", "e8d8", "c7e6"},
		FinalFEN: "rnbq1bnr/ppp2ppp/4N3/4p3/4P3/8/PPPP1PPP/R1BQKBNR b KQ - 1 4",
	},

	"3": {
		ID:       "3",
		StartFEN: "r1bqkbnr/ppp2ppp/2np4/4p3/2B1P3/2N5/PPPP1PPP/R1BQK1NR w KQkq - 4 4",
		Moves:    []string{"c4f7", "e8f7", "d1f3", "f7e6", "f3f5"},
		FinalFEN: "r1bq1bnr/ppp2ppp/2npk3/4pQ2/4P3/2N5/PPPP1PPP/R1B1K1NR b KQ - 2 6",
	},

	"4": {
		ID:       "4",
		StartFEN: "rnb1kbnr/pppp1ppp/8/4p3/3PP3/5N2/PPP2PPP/RNBQKB1R w KQkq - 0 3",
		Moves:    []string{"d4e5", "d8e7", "f3d4", "e7b4", "c2c3", "b4c5", "d4b5"},
		FinalFEN: "rnb1kbnr/pppp1ppp/8/1N2p3/8/2P5/PP3PPP/R1BQKB1R b KQkq - 1 5",
	},

	"5": {
		ID:       "5",
		StartFEN: "r2qkb1r/pp1n1ppp/2p1pn2/3p4/3P4/2N1PN2/PPP2PPP/R1BQKB1R w KQkq - 4 6",
		Moves:    []string{"e3e4", "d5e4", "c3e4", "f6e4", "d1e2", "e4f6", "e2e6"},
		FinalFEN: "r2qkb1r/pp1n1ppp/2p1pQ2/8/8/4PN2/PPP2PPP/R1B1KB1R b KQkq - 0 7",
	},

	"6": {
		ID:       "6",
		StartFEN: "rnbqkb1r/pp1p1ppp/2p1pn2/8/3NP3/2N5/PPP2PPP/R1BQKB1R w KQkq - 2 5",
		Moves:    []string{"e4f5", "e6f5", "c1g5", "d8e7", "g5e7"},
		FinalFEN: "rnb1kb1r/pp1peBpp/2p2n2/5p2/8/2N5/PPP2PPP/R2QKB1R b KQkq - 0 7",
	},

	"7": {
		ID:       "7",
		StartFEN: "r2qkbnr/ppp2ppp/2np4/4p1N1/2B1P3/3P4/PPP2PPP/RNBQK2R w KQkq - 2 5",
		Moves:    []string{"g5f7", "e8e7", "c1g5", "e7d7", "g5d8"},
		FinalFEN: "r2Bkbnr/ppp2ppp/2np4/4p3/2B1P3/3P4/PPP2PPP/RN1QK2R b KQ - 0 6",
	},

	"8": {
		ID:       "8",
		StartFEN: "r1bqkb1r/ppp2ppp/2np1n2/4p3/1bB1P3/2N2N2/PPPP1PPP/R1BQK2R w KQkq - 4 5",
		Moves:    []string{"c3d5", "f6d5", "c4d5", "c6d4", "d1a4"},
		FinalFEN: "r1bqkb1r/ppp2ppp/4p3/3bp3/Q2nP3/5N2/PPPP1PPP/R1B1K2R b KQkq - 1 6",
	},

	"9": {
		ID:       "9",
		StartFEN: "rnbq1rk1/pp3ppp/2pbpn2/3p4/3P4/2N1PN2/PPP1BPPP/R1BQ1RK1 w - - 1 8",
		Moves:    []string{"e2d3", "f6e4", "c3e4", "d5e4", "d3e4", "d8h4", "f3h4"},
		FinalFEN: "rnb2rk1/pp3ppp/2pbp3/8/7Q/4PN2/PPP2PPP/R1B2RK1 b - - 0 9",
	},

	"10": {
		ID:       "10",
		StartFEN: "r1bqkb1r/ppp2ppp/2n1pn2/3p4/3P4/2NB1N2/PPP2PPP/R1BQ1RK1 w kq - 2 6",
		Moves:    []string{"c3b5", "a7a6", "b5c7", "d8c7", "c1f4", "c7f4", "g2g3", "f4d2"},
		FinalFEN: "r1b1kb1r/p1p2ppp/3ppn2/8/8/4PNp1/PP1n1P1P/R1BQ1RK1 b kq - 1 8",
	},
	"11": {ID: "11", StartFEN: "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1", Moves: []string{"e2e4", "e7e5", "f1c4", "b8c6", "d1h5", "g8f6", "h5f7"}, FinalFEN: "r1bqkb1r/pppp1Qpp/2n2n2/4p3/2B5/8/PPPP1PPP/RNB1K1NR b KQkq - 0 4"},
}
