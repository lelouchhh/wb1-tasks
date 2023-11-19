package main

var justString string

func createHugeString(size int) string {
	data := make([]byte, size)
	return string(data)
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 10000 {
		justString = v[:10000]
	} else {
		justString = v
	}
}

func main() {
	someFunc()
	justString = ""
}
