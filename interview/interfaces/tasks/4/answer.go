package main

/*
	Как исправить?
*/

func convert(a []io.ReaderWriter) []io.Reader {
	var result []io.Reader
	for _, v := range a {
		result = append(result, v)
	}

	return result
}
func readAll(a []io.Reader) {}
func main() {
	var a []io.ReadWriter
	readAll(convert(a))
}
