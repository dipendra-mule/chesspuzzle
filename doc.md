WorkFlow

1. User enters chess puzzle page
2. User starts puzzles
3. Game assigns a random puzzle from the database
4. Game assigns white or black
5. User makes move
6. Game validates the move
7. Game updates the board
8. Game checks its a answer to puzzle
9. If correct, Game shows success message and moves to next puzzle
10. add counter total solved puzzles

REST API

1. GET /puzzle/start - Start a new puzzle

Request -

```json
{
  "level": "easy" // optional, can be easy, medium, hard
}
```

Response

```json
{
  "puzzlestarted": true,
  "puzzleId": "12345",
  "boardState": "8/8/8/8/8/8/8/8 w - - 0 1"
  "assignedColor": "white"
}
```

2. POST /puzzle/{puzzleId}/{move} - Make a move in the puzzle

Request -

```json
{
  "puzzleId": "12345",
  "move": "e4"
}
```

Response

```json
{
  "validMove": true,
  "boardState": "8/8/8/8/8/8/4P3/8 b - - 0 1",
  "isPuzzleSolved": false,
  "message": "Move accepted"
}
```
