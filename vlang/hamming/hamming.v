module main

fn distance(a string, b string) !int {
	if a == b { return 0 }
	if a == "" || b == "" || a.len != b.len { return error("lengths must match!") }
	mut result := 0
	for i := 0; i < a.len; i++ { if a[i] != b[i] { result++ } }
	return result
}
