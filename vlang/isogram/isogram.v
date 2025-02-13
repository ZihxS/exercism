module main

fn is_isogram(word string) bool {
	mut x := map[string]bool{}
	for _, w in word.to_lower().split('') {
		if w == ' ' || w == '-' { continue }
		if w in x { return false }
		x[w] = true
	}
	return true
}
