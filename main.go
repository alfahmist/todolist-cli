package main

import "fmt"

type Tugas struct {
	No        int
	Deskripsi string
	Status    bool
}

var todoList []Tugas

func main() {

	todoList = append(todoList, Tugas{Deskripsi: "Belajar Go", Status: true})
	todoList = append(todoList, Tugas{Deskripsi: "Belajar Struct", Status: false})
	todoList = append(todoList, Tugas{Deskripsi: "Belajar Array", Status: false})
	var menu int = 0
	fmt.Println("===TODO LIST===")
	fmt.Println("1. Tambah Tugas")
	fmt.Println("2. Lihat Tugas")
	fmt.Println("3. Tandai Tugas")
	fmt.Println("4. Hapus Tugas")
	fmt.Println("5. Keluar")
	fmt.Print("Pilh Menu : ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		fmt.Println("===Tambah Tugas===")
	case 2:
		fmt.Println("===Lihat Tugas===")

		for i, task := range todoList {
			status := "[ ]"
			if task.Status {
				status = "[x]"
			}
			fmt.Printf("%d. %s %s\n", i+1, status, task.Deskripsi)
		}

	case 3:
		fmt.Println("===Tandai Tugas===")
	case 4:
		fmt.Println("===Hapus Tugas===")
	case 5:
		fmt.Println("===Keluar===")
	}
}
