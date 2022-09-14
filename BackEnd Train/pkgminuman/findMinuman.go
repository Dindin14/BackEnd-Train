package pkgminuman

import (
	"gormstudy/db"
)

func TemukanMenuMinuman() []db.Minuman {

	var minums []db.Minuman

	result := db.DB.Find(&minums)
	err := result.Error
	if err != nil {
		panic(err.Error())
	}
	return minums
}
