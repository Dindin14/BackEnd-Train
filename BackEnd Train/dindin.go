// package main

// import (
// 	"fmt"
// 	"gormstudy/db"
// 	"gormstudy/pkgkaryawan"
// 	"gormstudy/pkgmakanan"
// 	"gormstudy/pkgminuman"
// 	"os"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// // type Makanan struct {
// // 	gorm.Model
// // 	Nama  string
// // 	Harga int
// // }

// // type Toping struct {
// // 	gorm.Model
// // 	Nama string

// // 	MakananID uint
// // 	Makanan   Makanan
// // }

// // type Minuman struct {
// // 	gorm.Model
// // 	Nama   string
// // 	Toping string
// // 	Harga  int
// // }

// // type Karyawan struct {
// // 	gorm.Model
// // 	Nama string
// // 	Umur int
// // 	Asal string
// // }

// //GORM ASSOCIATION

// //Koneksi Ke Data Base
// // func main() {
// // 	dsn := "dindin:PzcAyvPNrF@tcp(db.kaenova.my.id:2003)/dindin?charset=utf8mb4&parseTime=True&loc=Local"
// // 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// // 	if err != nil {
// // 		panic(err.Error())
// // 	}

// // 	//Migrasi Database
// // 	err = db.AutoMigrate(&Makanan{}, &Toping{})
// // 	if err != nil {
// // 		panic(err.Error())
// // 	}

// // 	//Buat Parent
// // 	makanan := Makanan{
// // 		Nama:  "Buzu",
// // 		Harga: 3000,
// // 	}

// // 	//Masukkan makanan ke database
// // 	err = db.Save(&makanan).Error
// // 	if err != nil {
// // 		panic(err.Error())
// // 	}

// // 	//Buat Toping
// // 	toping := Toping{
// // 		Nama:      "Keju",
// // 		MakananID: makanan.ID,
// // 	}

// // 	//Masukkan toping ke database
// // 	err = db.Save(&toping).Error
// // 	if err != nil {
// // 		panic(err.Error())
// // 	}

// // 	err = db.Preload("Makanan").First(&toping).Error
// // 	if err != nil {
// // 		panic(err.Error())
// // 	}
// // }

// //BASIC GORM
// func main() {
// 	var point int
// 	var minums []db.Minuman
// 	var karyawans []db.Karyawan
// 	var err error

// 	//Koneksi Ke Data Base
// 	dsn := "dindin:PzcAyvPNrF@tcp(db.kaenova.my.id:2003)/dindin?charset=utf8mb4&parseTime=True&loc=Local"
// 	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	//Migrasi Database
// 	err = db.DB.AutoMigrate(&db.Makanan{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	err = db.DB.AutoMigrate(&db.Minuman{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	err = db.DB.AutoMigrate(&db.Karyawan{})
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	//Save Data

// 	fmt.Println("1. Tambah Makanan")
// 	fmt.Println("2. Tambah Minuman")
// 	fmt.Println("3. Tambah Karyawan")
// 	fmt.Println("4. Delete Makanan")
// 	fmt.Println("5. Delete Minuman")
// 	fmt.Println("6. Delete Karyawan")
// 	fmt.Println("7. Update Makanan")
// 	fmt.Println("8. Update Minuman")
// 	fmt.Println("9. Update Karyawan")
// 	fmt.Println("10. Lihat Data Makanan")
// 	fmt.Println("11. Lihat Data Minuman")
// 	fmt.Println("12. Lihat Data Karyawan")
// 	fmt.Println("99. Exit")

// 	fmt.Scanln(&point)
// 	switch point {
// 	case 1:
// 		// namaMakanan, hargaMakanan := pkgmakanan.TambahMakanan()
// 		// result := db.DB.Create(&Makanan{Nama: namaMakanan, Harga: hargaMakanan})
// 		// err = result.Error
// 		// if err != nil {
// 		// 	panic(err.Error())
// 		// }
// 		pkgmakanan.TambahMakanan()

// 	case 2:
// 		// namaMinuman, toping, hargaMinuman := pkgminuman.TambahMinuman()
// 		// result := db.Create(&Minuman{Nama: namaMinuman, Toping: toping, Harga: hargaMinuman})
// 		// err = result.Error
// 		// if err != nil {
// 		// 	panic(err.Error())
// 		// }
// 		pkgminuman.TambahMinuman()

// 	case 3:

// 		// namaKaryawan, umurKaryawan, asalKaryawan := pkgkaryawan.TambahKaryawan()
// 		// result := db.Create(&Karyawan{Nama: namaKaryawan, Umur: umurKaryawan, Asal: asalKaryawan})
// 		// err = result.Error
// 		// if err != nil {
// 		// 	panic(err.Error())
// 		// }
// 		pkgkaryawan.TambahKaryawan()

// 	case 4:

// 		// idDltMakanan := pkgmakanan.KurangMakanan()

// 		// result := db.Delete(&Makanan{}, idDltMakanan)
// 		// err = result.Error
// 		// if err != nil {
// 		// 	panic(err.Error())
// 		// }
// 		pkgmakanan.KurangMakanan()

// 	case 5:

// 		// idMinuman := pkgminuman.KurangMinuman()

// 		// result := db.Delete(&Minuman{}, idMinuman)
// 		// err = result.Error
// 		// if err != nil {
// 		// 	panic(err.Error())
// 		// }
// 		pkgminuman.KurangMinuman()

// 	case 6:

// 		// idKaryawan := pkgkaryawan.KurangKaryawan()

// 		// result := db.Delete(&Karyawan{}, idKaryawan)
// 		// err = result.Error
// 		// if err != nil {
// 		// 	panic(err.Error())
// 		// }
// 		pkgkaryawan.KurangKaryawan()

// 	case 7:

// 		// var update, hargaUpdate int
// 		// var namaUpdate, idUptMakanan string

// 		// fmt.Println("ID Makanan: ")
// 		// fmt.Scanln(&idUptMakanan)
// 		// fmt.Println("1. Ganti Nama Makanan")
// 		// fmt.Println("2. Ganti Harga Makanan")
// 		// fmt.Scanln(&update)

// 		// switch update {
// 		// case 1:
// 		// 	fmt.Println("Nama Makanan: ")
// 		// 	fmt.Scanln(&namaUpdate)
// 		// 	db.Model(&Makanan{}).Where("id = ?", idUptMakanan).Update("Nama", namaUpdate)
// 		// case 2:
// 		// 	fmt.Println("Harga Makanan: ")
// 		// 	fmt.Scanln(&hargaUpdate)
// 		// 	db.Model(&Makanan{}).Where("id = ?", idUptMakanan).Update("Harga", hargaUpdate)
// 		// }
// 		pkgmakanan.UpdateMenuMakanan()

// 	case 8:

// 		// var update, hargaUpdate int
// 		// var namaUpdate, topingUpdate, idUptMinuman string

// 		// fmt.Println("ID Minuman: ")
// 		// fmt.Scanln(&idUptMinuman)
// 		// fmt.Println("1. Ganti Nama Minuman")
// 		// fmt.Println("2. Ganti Toping Minuman")
// 		// fmt.Println("3. Ganti Harga Minuman")
// 		// fmt.Scanln(&update)

// 		// switch update {
// 		// case 1:
// 		// 	fmt.Println("Nama Minuman: ")
// 		// 	fmt.Scanln(&namaUpdate)
// 		// 	db.Model(&Minuman{}).Where("id = ?", idUptMinuman).Update("Nama", namaUpdate)
// 		// case 2:
// 		// 	fmt.Println("Toping Minuman: ")
// 		// 	fmt.Scanln(&topingUpdate)
// 		// 	db.Model(&Minuman{}).Where("id = ?", idUptMinuman).Update("Toping", topingUpdate)
// 		// case 3:
// 		// 	fmt.Println("Harga Minuman: ")
// 		// 	fmt.Scanln(&hargaUpdate)
// 		// 	db.Model(&Minuman{}).Where("id = ?", idUptMinuman).Update("Harga", hargaUpdate)
// 		// }
// 	case 9:

// 		//Update data Nama Karyawan,Umur, dan Asal
// 		var update, umurUpdate int
// 		var namaUpdate, asalUpdate, idUptKaryawan string
// 		fmt.Println("ID Karyawan: ")
// 		fmt.Scanln(&idUptKaryawan)
// 		fmt.Println("1. Ganti Nama Karyawan")
// 		fmt.Println("2. Ganti Umur Karyawan")
// 		fmt.Println("3. Ganti Asal Karyawan")
// 		fmt.Scanln(&update)

// 		switch update {
// 		case 1:
// 			fmt.Println("Nama Karyawan: ")
// 			fmt.Scanln(&namaUpdate)
// 			db.Model(&Karyawan{}).Where("id = ?", idUptKaryawan).Update("Nama", namaUpdate)
// 		case 2:
// 			fmt.Println("Umur Karyawan: ")
// 			fmt.Scanln(&umurUpdate)
// 			db.Model(&Karyawan{}).Where("id = ?", idUptKaryawan).Update("Umur", umurUpdate)
// 		case 3:
// 			fmt.Println("Asal Karyawan: ")
// 			fmt.Scanln(&asalUpdate)
// 			db.Model(&Karyawan{}).Where("id = ?", idUptKaryawan).Update("Asal", asalUpdate)
// 		}
// 	case 10:

// 		//Melihat Data Makanan
// 		result := db.Find(&makans)
// 		err = result.Error
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		fmt.Println(makans)

// 	case 11:

// 		//Melihat Data Minuman
// 		result := db.Find(&minums)
// 		err = result.Error
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		fmt.Println(minums)

// 	case 12:

// 		//Melihat Data Karyawan
// 		result := db.Find(&karyawans)
// 		err = result.Error
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		fmt.Println(karyawans)

// 	case 99:
// 		os.Exit(99)
// 	}

// 	//Ambil Data
// 	//Ambil 1 Data Paling Pertama
// 	// result = db.First(&makan)
// 	// err = result.Error
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	// fmt.Println(makan)

// 	//Ambil Semua Data Makanan
// 	// result = db.Find(&makans)
// 	// err = result.Error
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	// fmt.Println(makans)

// 	//Ambil Data Berdasarkan ID
// 	// var makan2 Makanan

// 	// result = db.First(&makan2, makan.ID)
// 	// err = result.Error
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	// fmt.Println(makan2)

// 	//Hapus Data

// 	// var hapus Makanan
// 	// idMakanan := pkgkaryawan.KurangKaryawan()

// 	// result := db.Delete(&Makanan{}, idMakanan)
// 	// err = result.Error
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }
// 	// fmt.Println(makan2)

// }

package main

import (
	"gormstudy/db"
	"gormstudy/pkgkaryawan"
	"gormstudy/pkgmakanan"
	"gormstudy/pkgminuman"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	var err error

	//Koneksi Ke Data Base
	dsn := "dindin:PzcAyvPNrF@tcp(db.kaenova.my.id:2003)/dindin?charset=utf8mb4&parseTime=True&loc=Local"
	db.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	app := fiber.New()

	app.Get("/makanan", func(c *fiber.Ctx) error {
		makanan := pkgmakanan.TemukanMenuMakanan()
		return c.JSON(makanan)
	})

	app.Get("/minuman", func(c *fiber.Ctx) error {
		minuman := pkgminuman.TemukanMenuMinuman()
		return c.JSON(minuman)
	})

	app.Get("/karyawan", func(c *fiber.Ctx) error {
		karyawan := pkgkaryawan.TemukanKaryawan()
		return c.JSON(karyawan)
	})

	app.Post("/makanan", func(c *fiber.Ctx) error {
		inputNama := c.Query("inputNama")
		if inputNama == "" {
			inputNama = "kosong"
		}

		inputHarga := c.Query("inputHarga")
		if inputHarga == "" {
			inputHarga = "kosong"
		}

		i, err := strconv.Atoi(inputHarga)
		if err != nil {
			panic(err.Error())
		}
		pkgmakanan.TambahMakanan(inputNama, i)
		return c.JSON(fiber.Map{
			"inputNama":  inputNama,
			"inputHarga": inputHarga,
		})
	})

	app.Post("/minuman", func(c *fiber.Ctx) error {
		inputNama := c.Query("inputNama")
		if inputNama == "" {
			inputNama = "kosong"
		}

		inputToping := c.Query("inputToping")
		if inputToping == "" {
			inputToping = "kosong"
		}

		inputHarga := c.Query("inputHarga")
		if inputHarga == "" {
			inputHarga = "kosong"
		}

		i, err := strconv.Atoi(inputHarga)
		if err != nil {
			panic(err.Error())
		}
		pkgminuman.TambahMinuman(inputNama, i, inputToping)
		return c.JSON(fiber.Map{
			"inputNama":   inputNama,
			"inputHarga":  inputHarga,
			"inputToping": inputToping,
		})
	})

	app.Post("/karyawan", func(c *fiber.Ctx) error {
		inputNama := c.Query("inputNama")
		if inputNama == "" {
			inputNama = "kosong"
		}

		inputUmur := c.Query("inputUmur")
		if inputUmur == "" {
			inputUmur = "kosong"
		}

		i, err := strconv.Atoi(inputUmur)
		if err != nil {
			panic(err.Error())
		}

		inputAsal := c.Query("inputAsal")
		if inputAsal == "" {
			inputAsal = "kosong"
		}

		pkgkaryawan.TambahKaryawan(inputNama, i, inputAsal)
		return c.JSON(fiber.Map{
			"inputNama": inputNama,
			"inputUmur": inputUmur,
			"inputAsal": inputAsal,
		})
	})

	app.Delete("/makanan", func(c *fiber.Ctx) error {
		return c.SendString("Ini hapus makanan")
	})

	app.Patch("/makanan", func(c *fiber.Ctx) error {
		return c.SendString("Ini ubah makanan")
	})
	app.Listen(":3000")
}
