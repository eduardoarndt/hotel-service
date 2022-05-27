package data

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/eduardoarndt/hotel-service/domain"
)

var (
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	pass = os.Getenv("DB_PASSWORD")
	base = os.Getenv("DB_NAME")
	db   *sql.DB
)

func Connect() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, base)

	var err error
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("[Postgres] Connected!")
}

func GetHotel(id string) (*domain.Hotel, error) {
	sql := db.QueryRow(fmt.Sprintf("select * from hotel where id = %v;", id))

	var hotel domain.Hotel
	err := sql.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.City, &hotel.Reviews, &hotel.Rating)
	if err != nil {
		return nil, err
	}

	return &hotel, nil
}

func GetAllHotels() ([]domain.Hotel, error) {
	sql, err := db.Query("select * from hotel;")
	if err != nil {
		return nil, err
	}

	var hotels []domain.Hotel
	for sql.Next() {
		var hotel domain.Hotel
		err = sql.Scan(&hotel.ID, &hotel.Name, &hotel.Address, &hotel.City, &hotel.Reviews, &hotel.Rating)
		if err != nil {
			return nil, err
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}
