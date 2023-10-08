package database

import (
	"assigment3-GLNG-08-011/entity"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "assigment2"
)

const updateQuery = `update "waterwind" set "water" = $1 , "wind" = $2 where "id" = $3 `

// koneksi ke dalam database
func HandlerDatabaseConnection() {
	sqlinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", sqlinfo)

	if err != nil {

		log.Panic("Error saat validasi database argumen", err)
	}

	err = db.Ping()

	if err != nil {

		log.Panic("Error saat koneksi ke database", err)
	}
}

// func untuk update data pada database. Dengan catatat untuk id 1
func UpdateWaterWind(payload entity.WaterWind) error {

	tx, err := db.Begin()

	if err != nil {
		return err
	}

	_, err = tx.Exec(updateQuery, payload.Water, payload.Wind, 1)

	if err != nil {
		tx.Rollback()
		return err

	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
