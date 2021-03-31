package algo_func_for_tests

func Mention(note int8) string {
	var mention string
	if note > 16 {
		mention = "TB"
	} else if note > 14 {
		mention = "B"
	} else if note > 12 {
		mention = "AB"
	} else {
		mention = "Aucune"
	}
	return mention
}

func Pair(i int8) bool {
	var ret bool
	if i%2 == 0 {
		ret = true
	} else {
		ret = false
	}
	return ret
}
