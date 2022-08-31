package main

import (
	"fmt"
	"gormstudy/pkgkaryawan"
	"gormstudy/pkgmakanan"
	"gormstudy/pkgminuman"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

//GORM ASSOCIATION

//Koneksi Ke Data Base
// func main() {
// 	dsn := "dindin:PzcAyvPNrF@tcp(db.kaenova.my.id:2003)/dindin?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	//Migrasi Database
// 	err = db.AutoMigrate(&Makanan{}, &Toping{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	//Buat Parent
// 	makanan := Makanan{
// 		Nama:  "Buzu",
// 		Harga: 3000,
// 	}

// 	//Masukkan makanan ke database
// 	err = db.Save(&makanan).Error
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	//Buat Toping
// 	toping := Toping{
// 		Nama:      "Keju",
// 		MakananID: makanan.ID,
// 	}

// 	//Masukkan toping ke database
// 	err = db.Save(&toping).Error
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	err = db.Preload("Makanan").First(&toping).Error
// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

//BASIC GORM
func main() {
	var point int
	var makans []Makanan

	//Koneksi Ke Data Base
	dsn := "dindin:PzcAyvPNrF@tcp(db.kaenova.my.id:2003)/dindin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//Migrasi Database
	err = db.AutoMigrate(&Makanan{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&Minuman{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&Karyawan{})
	if err != nil {
		panic(err.Error())
	}

	//Save Data

	fmt.Println("1. Tambah Makanan")
	fmt.Println("2. Tambah Minuman")
	fmt.Println("3. Tambah Karyawan")
	fmt.Println("4. Delete Makanan")
	fmt.Println("5. Delete Minuman")
	fmt.Println("6. Delete Karyawan")
	fmt.Println("9. Exit")

	fmt.Scanln(&point)
	switch point {
	case 1:
		// makan := Makanan{Nama: "Bakso", Harga: 10000}
		namaMakanan, hargaMakanan := pkgmakanan.TambahMakanan()
		result := db.Create(&Makanan{Nama: namaMakanan, Harga: hargaMakanan})
		err = result.Error
		if err != nil {
			panic(err.Error())
		}

	case 2:
		// minum := Minuman{Nama: "PopIce", Harga: 5000, Toping: "Oreo"}
		namaMinuman, toping, hargaMinuman := pkgminuman.TambahMinuman()
		result := db.Create(&Minuman{Nama: namaMinuman, Toping: toping, Harga: hargaMinuman})
		err = result.Error
		if err != nil {
			panic(err.Error())
		}

	case 3:
		// karya := Karyawan{Nama: "Susantea", Umur: 21, Asal: "Ayah"}
		namaKaryawan, umurKaryawan, asalKaryawan := pkgkaryawan.TambahKaryawan()
		result := db.Create(&Karyawan{Nama: namaKaryawan, Umur: umurKaryawan, Asal: asalKaryawan})
		err = result.Error
		if err != nil {
			panic(err.Error())
		}
	case 4:
		//Hapus Data Makanan

		// hapus Makanan
		idDltMakanan := pkgmakanan.KurangMakanan()

		result := db.Delete(&Makanan{}, idDltMakanan)
		err = result.Error
		if err != nil {
			panic(err.Error())
		}
	case 5:
		// hapus Minuman
		idMinuman := pkgminuman.KurangMinuman()

		result := db.Delete(&Minuman{}, idMinuman)
		err = result.Error
		if err != nil {
			panic(err.Error())
		}

	case 6:
		// hapus Karyawan
		idKaryawan := pkgkaryawan.KurangKaryawan()

		result := db.Delete(&Karyawan{}, idKaryawan)
		err = result.Error
		if err != nil {
			panic(err.Error())
		}
	case 7:

		var update, hargaUpdate int
		var namaUpdate, idUptMakanan string

		//Update data Makanan
		// makans := Makanan{Nama: "Bakso", Harga: 10000}
		// makans.Nama = "Wacepong"

		fmt.Println("ID Makanan: ")
		fmt.Scanln(&idUptMakanan)
		// idUptMakanan := pkgmakanan.UpdateMenuMakanan()
		fmt.Println("Ganti Nama Makanan")
		fmt.Println("Ganti Harga Makanan")
		fmt.Scanln(&update)

		switch update {
		case 1:
			fmt.Scanln(&namaUpdate)
			db.Model(&Makanan{}).Where("id = ?", idUptMakanan).Update("Nama", namaUpdate)
		case 2:
			fmt.Scanln(&namaUpdate)
			db.Model(&Makanan{}).Where("id = ?", idUptMakanan).Update("Harga", hargaUpdate)
		}

	case 8:

		//Melihat Data
		result := db.Find(&makans)
		err = result.Error
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(makans)
	case 9:
		os.Exit(9)
	}

	//Ambil Data
	//Ambil 1 Data Paling Pertama
	// result = db.First(&makan)
	// err = result.Error
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(makan)

	//Ambil Semua Data Makanan
	// result = db.Find(&makans)
	// err = result.Error
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(makans)

	//Ambil Data Berdasarkan ID
	// var makan2 Makanan

	// result = db.First(&makan2, makan.ID)
	// err = result.Error
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(makan2)

	//Hapus Data

	// var hapus Makanan
	// idMakanan := pkgkaryawan.KurangKaryawan()

	// result := db.Delete(&Makanan{}, idMakanan)
	// err = result.Error
	// if err != nil {
	// 	panic(err.Error())
	// }
	// fmt.Println(makan2)

}

// func TambahMakanan() {
// 	panic("unimplemented")
// }
