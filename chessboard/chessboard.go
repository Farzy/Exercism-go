package chessboard

// File stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Chessboard contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (count int) {
	fileContent, _ := cb[file]
	// if "file" is invalid, fileContent will be empty and count remains 0
	for _, value := range fileContent {
		if value {
			count++
		}
	}
	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (count int) {
	if rank < 1 || rank > 8 {
		return
	}
	for _, file := range cb {
		if file[rank-1] {
			count++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (count int) {
	for _, file := range cb {
		for range file {
			count++
		}
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (count int) {
	for _, file := range cb {
		for _, square := range file {
			if square {
				count++
			}
		}
	}
	return
}
