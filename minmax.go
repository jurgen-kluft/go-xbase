package xbase

// Min3u8 returns the smallest of three numbers
func Min3u8(a, b, c uint8) uint8 {
	if (a < b) && (a < c) {
		return a
	} else if (b < a) && (b < c) {
		return b
	}
	return c
}

// Max3u8 returns the largest of three numbers
func Max3u8(a, b, c uint8) uint8 {
	if (a >= b) && (a >= c) {
		return a
	} else if (b >= a) && (b >= c) {
		return b
	}
	return c
}

// Min3i32 returns the smallest of three numbers
func Min3i32(a, b, c int) int {
	if (a < b) && (a < c) {
		return a
	} else if (b < a) && (b < c) {
		return b
	}
	return c
}

// Max3i32 returns the largest of three numbers
func Max3i32(a, b, c int32) int32 {
	if (a >= b) && (a >= c) {
		return a
	} else if (b >= a) && (b >= c) {
		return b
	}
	return c
}

// Min3f32 returns the smallest of three numbers
func Min3f32(a, b, c float32) float32 {
	if (a < b) && (a < c) {
		return a
	} else if (b < a) && (b < c) {
		return b
	}
	return c
}

// Max3f32 returns the largest of three numbers
func Max3f32(a, b, c float32) float32 {
	if (a >= b) && (a >= c) {
		return a
	} else if (b >= a) && (b >= c) {
		return b
	}
	return c
}
