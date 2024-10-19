package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println("Hello, world!")

	fmt.Println(26)
	fmt.Println(1024 + 48)   // 1072
	fmt.Println(5 + 8*2)     // 21
	fmt.Println((5 + 8) * 2) // 26

	fmt.Print("Raz", 1, 2, 3)
	fmt.Print("Dva", 1, 2, 3, "Tri")
	fmt.Print("\n\n")
	fmt.Print("Hello World! ", "hi", 6.9, 7, 8.1, 0, 2, "Lala", true, true, false, "tutu", 10, 20)
	fmt.Println("Etot argument tozhe zdes i prikleilsya, a vot sleduyuschiy net", "Sleduyuschiy", 2, 3)
	fmt.Println("cnh1", 100, 200, 300)
	fmt.Println()
	fmt.Println("Hello World! ", "hi", 6.9, 7, 8.1, 0, 2, "Lala", true, true, false, "tutu")

	fmt.Println("Здравствуй,")
	fmt.Println("Иосиф!")

	name := "Олег"
	age := 25
	weight := 80.5
	fmt.Println("Имя:", name, "возраст:", age, "вес:", weight)

	//var lastName, firstName, patronymic string
	//fmt.Scan(&lastName, &firstName, &patronymic)
}
