package algos

import "math"

func TwoCb(levels []bool) int {
	jump := int(math.Sqrt(float64(len(levels))))

	var fb int
	for i := 0; i < len(levels); i += jump {
		if levels[i] {
			fb = i
			break
		}
	}

	start := fb - jump
	var ob int
	for i := start; i < fb; i++ {
		if levels[i] {
			ob = i
			break
		}
	}

	return ob
}
