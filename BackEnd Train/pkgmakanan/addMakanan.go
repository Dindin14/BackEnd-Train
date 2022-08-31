package pkgmakanan

import "fmt"

var namaMakanan string
var hargaMakanan int

func TambahMakanan() (string, int) {
	fmt.Println("Nama makanan: ")
	fmt.Scanln(&namaMakanan)
	fmt.Println("Harga: ")
	fmt.Scanln(&hargaMakanan)

	return namaMakanan, hargaMakanan
}
