package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Tugas struct {
	No        int
	Deskripsi string
	Status    bool
}

var todoList []Tugas

// This function clears the terminal screen.
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func lihatTugas() {
	for i, task := range todoList {
		status := "[ ]"
		if task.Status {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", i+1, status, task.Deskripsi)
	}
}
func main() {

	todoList = append(todoList, Tugas{Deskripsi: "Belajar Go", Status: true})
	todoList = append(todoList, Tugas{Deskripsi: "Belajar Struct", Status: false})
	todoList = append(todoList, Tugas{Deskripsi: "Belajar Array", Status: false})
	var menu int = 0
	var no int = 0
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
		fmt.Print("Deskripsi : ")
		reader := bufio.NewReader(os.Stdin)
		bufio.NewReader(os.Stdin)
		teks, _ := reader.ReadString('\n')

		todoList = append(todoList, Tugas{Deskripsi: teks, Status: false})
		lihatTugas()
	case 2:
		fmt.Println("===Lihat Tugas===")
		lihatTugas()

	case 3:
		fmt.Println("===Tandai Tugas===")
		lihatTugas()
		fmt.Print("Tandai Nomor : ")
		fmt.Scanln(&no)
		todoList[no-1].Status = true
		lihatTugas()

	case 4:
		fmt.Println("===Hapus Tugas===")
	case 5:
		fmt.Println("===Keluar===")

	default:
		// ClearScreen()

	}
}
