package pkgminuman

import "fmt"

var namaMinuman string
var hargaMinuman int
var toping string

func TambahMinuman() (string, string, int) {
	fmt.Println("Nama minuman: ")
	fmt.Scanln(&namaMinuman)
	fmt.Println("Toping: ")
	fmt.Scanln(&toping)
	fmt.Println("Harga: ")
	fmt.Scanln(&hargaMinuman)

	return namaMinuman, toping, hargaMinuman
}
