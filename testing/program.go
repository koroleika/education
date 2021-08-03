package example

func yourFunction(n int) string {
	if n%2 == 0 {
		return "Четное"
	} else {
		return "Нечетное"
	}
}

func main() {
	yourFunction(10)

}
