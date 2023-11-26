package main

func main() {

}

func sumInt(a ...int) (int, int) {
	summa := 0
	for _, el := range a {
		summa += el
	}
	return len(a), summa
}
