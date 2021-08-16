package main

import "fmt"

// membuat channel sebagai parameter
func write(data []string, ch chan string) {
	for _, dt := range data {
		ch <- dt
	}
	// close tidak bisa digunakan lagi baik untuk menerima data ataupun untuk mengirim data
	close(ch)
}

// fungsi mencetak list value pada channel
func print(ch chan string) {
	for result := range ch {
		fmt.Println(result)
	}
}

func main() {

	//inisialisasi variabel slice
	programingLang := []string{"Golang", "Kotlin", "JavaScript", "Ruby"}

	//inisialisasi variabel channel
	ch := make(chan string, len(programingLang))

	//goroutine digunakan untuk mengirim data via channel
	go write(programingLang, ch)

	// menampilkan nilai dari data yang dikirim via channel
	print(ch)

}
