package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

// Code optimized using https://exercism.org/tracks/go/exercises/strain/approaches/using-generics

type Slicer interface {
	int | string | []int
}

func Keep[T Slicer](collection []T, predicate func(T) bool) []T {
	if collection == nil {
		return collection
	}
	output := make([]T, 0, len(collection))
	for _, item := range collection {
		if predicate(item) {
			output = append(output, item)
		}
	}
	return output
}

func Discard[T Slicer](collection []T, predicate func(T) bool) []T {
	return Keep(collection, func(t T) bool {
		return !predicate(t)
	})
}
