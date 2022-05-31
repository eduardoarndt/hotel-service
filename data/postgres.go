package data

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

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

func CreateHotel(hotel domain.Hotel) (*int, error) {
	sql := `insert into hotel (name, address, city, reviews, rating) values ($1, $2, $3, $4, $5) returning id`
	id := 0
	err := db.QueryRow(sql, hotel.Name, hotel.Address, hotel.City, hotel.Reviews, hotel.Rating).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func UpdateHotel(hotel domain.Hotel) error {
	fmt.Println(hotel)
	sql := `update hotel set (name, address, city, reviews, rating) = ($1, $2, $3, $4, $5) where id = $6`
	hotelId, err := strconv.Atoi(hotel.ID)
	if err != nil {
		return err
	}

	_, err = db.Exec(sql, hotel.Name, hotel.Address, hotel.City, hotel.Reviews, hotel.Rating, hotelId)
	if err != nil {
		return err
	}

	return nil
}

func DeleteHotel(id string) error {
	sql := `delete from hotel where id = $1`
	_, err := db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
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
