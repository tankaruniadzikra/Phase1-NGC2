// NGC 2 - Conditional 1
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var name string
	//input nama
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&name)

	//menghasilkan angka acak antara 0 hingga 100
	angkaRandom := rand.Intn(100) + 1

	switch {
	case angkaRandom > 80:
		fmt.Printf("Selamat %s, anda sangat beruntung\n", name)
	case angkaRandom <= 80 && angkaRandom > 60:
		fmt.Printf("Selamat %s, anda beruntung\n", name)
	case angkaRandom <= 60 && angkaRandom > 40:
		fmt.Printf("Mohon maaf %s, anda kurang beruntung\n", name)
	default:
		fmt.Printf("Mohon maaf %s, anda sial\n", name)
	}
}

// NGC 2 - Conditional 2

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var nama string
	var umurInput string

	//input nama
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&nama)

	//input umur
	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umurInput)

	umur, err := strconv.Atoi(umurInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//cek umur dan menampilkan pesan sesuai dengan keterangan yang sudah diinput
	if umur < 0 || umur > 100 {
		fmt.Println("Error: Umur harus berada di antara 0 hingga 100")
		return
	}

	if umur > 18 {
		fmt.Println("Silahkan masuk")
	} else {
		fmt.Println("Dilarang masuk, maksimal umur 18")
	}
}

// eksplorasi strings string.TrimSpace
