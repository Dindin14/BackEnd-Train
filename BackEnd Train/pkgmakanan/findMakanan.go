package pkgmakanan

import (
	"gormstudy/db"
)

func TemukanMenuMakanan() []db.Makanan {
	var makans []db.Makanan

	result := db.DB.Find(&makans)
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
	// fmt.Println(makans)
	return makans
}
