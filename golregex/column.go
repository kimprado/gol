package golregex

/*
 get alphanumeric string from required
	'col' number column
	'separator' marking the column gutters
	'dat' sliced vertically for columns
*/
func Column(dat string, separator string, col int) string {
	var pattern string
	for idx := 1; idx < col; idx++ {
		pattern += fmt.Sprintf("[a-zA-Z0-9]+%s", separator)
	}
	pattern += "([a-zA-Z0-9]+)"
	regx := regexp.MustCompile(pattern)
	result := regx.FindStringSubmatch(dat)
	if len(result) == 0 {
		return ""
	}
	return result[1]
}
