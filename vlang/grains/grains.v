module main

fn grains_on_square(square int) !u64 {
	if square < 1 || square > 64 {
		return error("square must be between 1 and 64")
	}
	if square == 1 || square == 2 {
		return u64(square)
	}
	mut result := u64(2)
	for i := 2; i < square; i++ {
		result = result * 2
	}
	return u64(result)
}

fn total_grains_on_board() u64 {
	return 18446744073709551615
}
