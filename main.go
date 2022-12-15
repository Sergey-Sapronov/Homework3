package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func wr() {

	f, err := os.OpenFile("hello.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	var name string
	fmt.Println("Введите имя:")
	fmt.Scanf("%s\n", &name)
	f.WriteString(name)
	f.WriteString(" ")

	var sur string
	fmt.Println("Введите фамилию:")
	fmt.Scanf("%s\n", &sur)
	f.WriteString(sur)
	f.WriteString(" написал: ")

	var note string
	fmt.Println("Введите текст заметки:")
	reader := bufio.NewReader(os.Stdin)
	note, err3 := reader.ReadString('\n')

	if err3 != nil {
		log.Fatal(err3)
	}
	f.WriteString(note)

	fmt.Printf("Привет, %s %s, Ваша заметка: %s", name, sur, note)
	fmt.Printf("Записать eщё?")
	var ans string
	fmt.Scanf("%s\n", &ans)
	if ans == "yes" {
		main()
	}
	defer f.Close()

}

func main() {
	wr()
}
