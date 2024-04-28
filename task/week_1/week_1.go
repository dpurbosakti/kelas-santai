package main

import (
	"fmt"
	"strings"
)

// 1
func add(x, y int) int {
	return x + y
}

// 2
func looping(n uint32) {
	for i := uint32(1); i <= n; i++ {
		fmt.Printf("nilai i ke %d\n", i)
	}
}

// 3
type Mahasiswa struct {
	nama    string
	usia    uint32
	jurusan string
}

// 5
var defaultName = "moko"

func main() {
	repeated := strings.Repeat("=", 40)
	longLine := "\n" + repeated + "\n"

	// 1
	fmt.Print("Task 1\n\n")
	fmt.Println(add(1, 2))
	fmt.Println(add(2, 2))
	fmt.Println(add(3, 2))
	fmt.Println(longLine)

	// 2
	fmt.Print("Task 2\n\n")
	looping(10)
	fmt.Println(longLine)

	// 3
	fmt.Print("Task 3\n\n")
	dataMoko := Mahasiswa{
		nama:    "moko",
		usia:    30,
		jurusan: "kelas_santai",
	}
	fmt.Printf("nama: %s\n", dataMoko.nama)
	fmt.Printf("usia: %d\n", dataMoko.usia)
	fmt.Printf("jurusan: %s\n", dataMoko.jurusan)
	fmt.Println(longLine)

	// 4
	fmt.Print("Task 4\n\n")
	newSlice := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice: %v\n", newSlice)
	fmt.Printf("panjang slice: %d\n", len(newSlice))
	fmt.Println(longLine)

	// 5
	/*
		di sini kita berusaha mengganti defaultName dengan value "dwi" dengan metode copy
		value defaultName ke changedName1 lalu merubah value changedName1 dengan realName("dwi")
		disini value defaultName tidak berubah karena yg kita rubah hanya copyan valuenya saja.
	*/
	fmt.Print("Task 5\n\n")
	changedName1 := defaultName
	realName := "dwi"
	fmt.Println("changed_name_1: ", changedName1)
	changedName1 = realName
	fmt.Println("changed_name_1: ", changedName1)
	fmt.Println("default_name: ", defaultName)
	fmt.Print("\n")

	/*
		di sini kita berusaha mengganti defaultName dengan value "dwi" dengan metode copy address
		value defaultName ke changedName2 lalu merubah value changedName2 dengan realName("dwi")
		disini value defaultName berubah karena yg kita rubah adalah value asli yg
		kita dapat dari address value.
	*/
	changedName2 := &defaultName
	fmt.Println("changed_name_2: ", *changedName2)
	*changedName2 = realName
	fmt.Println("changed_name_2: ", *changedName2)
	fmt.Println("default_name: ", defaultName)
	fmt.Print("\n")

	/*
		di sini jg dilakukan pengecekan apakah value yg kita copy sebelumnya ke changedName1
		ikut berubah/terdampak setelah kita rubah value defaultNamenya, dan ternyata ikut berubah.
	*/
	fmt.Println("changed_name_1: ", changedName1)
	fmt.Println(longLine)
}
