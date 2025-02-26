package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, s string) int {
	var i int
	for _, file := range cb[s] {
		if file {
			i++
		}
	}
	return i
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	var i int
	for _, file := range cb {
		if rank > len(file) || rank < 1 {
			return 0
		} else if file[rank-1] {
			i++
		}
	}
	return i
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	var i int
	for _, file := range cb {
		for range file {
			i++
		}
	}
	return i
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	var count int
	for _, file := range cb {
		for i := range file {
			if file[i] {
				count++
			}
		}
	}
	return count
}
