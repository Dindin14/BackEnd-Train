package pkgmakanan

import "fmt"

var idDltMakanan string

func KurangMakanan() string {
	fmt.Println("ID Makanan: ")
	fmt.Scanln(&idDltMakanan)

	return idDltMakanan
}
