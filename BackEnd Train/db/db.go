package db

import "gorm.io/gorm"

var DB *gorm.DB = nil

type Makanan struct {
	gorm.Model
	Nama  string
	Harga int
}

// type Toping struct {
// 	gorm.Model
// 	Nama string

// 	MakananID uint
// 	Makanan   Makanan
// }

type Minuman struct {
	gorm.Model
	Nama   string
	Toping string
	Harga  int
}

type Karyawan struct {
	gorm.Model
	Nama string
	Umur int
	Asal string
}
