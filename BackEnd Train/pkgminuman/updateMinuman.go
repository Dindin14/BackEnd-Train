package pkgminuman

import (
	"fmt"
	"gormstudy/db"
)

func UpdateMenuMinuman() {
	var update, hargaUpdate int
	var namaUpdate, topingUpdate, idUptMinuman string

	fmt.Println("ID Minuman: ")
	fmt.Scanln(&idUptMinuman)
	fmt.Println("1. Ganti Nama Minuman")
	fmt.Println("2. Ganti Toping Minuman")
	fmt.Println("3. Ganti Harga Minuman")
	fmt.Scanln(&update)

	switch update {
	case 1:
		fmt.Println("Nama Minuman: ")
		fmt.Scanln(&namaUpdate)
		db.DB.Model(&db.Minuman{}).Where("id = ?", idUptMinuman).Update("Nama", namaUpdate)
	case 2:
		fmt.Println("Toping Minuman: ")
		fmt.Scanln(&topingUpdate)
		db.DB.Model(&db.Minuman{}).Where("id = ?", idUptMinuman).Update("Toping", topingUpdate)
	case 3:
		fmt.Println("Harga Minuman: ")
		fmt.Scanln(&hargaUpdate)
		db.DB.Model(&db.Minuman{}).Where("id = ?", idUptMinuman).Update("Harga", hargaUpdate)
	}
}
