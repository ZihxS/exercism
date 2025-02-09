module main

import regex

fn response(hey_bob string) string {
	trimmed_input := hey_bob.trim_space()
	if trimmed_input == '' {
		return "Fine. Be that way!"
	}
	pattern := r'[A-Z]'
	re := regex.regex_opt(pattern) or { panic(err) }
	is_yelling := trimmed_input.to_upper() == trimmed_input && (re.matches_string(trimmed_input) || trimmed_input.ends_with('!'))
	is_question := trimmed_input.ends_with('?')
	if is_yelling && is_question {
		return "Calm down, I know what I'm doing!"
	}
	if is_yelling {
		return "Whoa, chill out!"
	}
	if is_question {
		return "Sure."
	}
	return "Whatever."
}
