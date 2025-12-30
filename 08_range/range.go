package main

func main() {
	//range function
	nums := []int{1, 2, 3, 4, 5}

	for i, num := range nums {
		println("Index:", i, "Value:", num)
	}

	//range over map
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		println("Key:", key, "Value:", value)
	}

	//range over string
	str := "hello"
	for i, ch := range str {
		println("Index:", i, "Character:", string(ch))
	}
}
