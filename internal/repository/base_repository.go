package repository

import (
	"time"

	"github.com/AthirsonSilva/golang-net-http-restapi/internal/config"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/database"
	"github.com/AthirsonSilva/golang-net-http-restapi/internal/models"
)

type DatabaseRepository interface {
	FindAllUsers() bool
	InsertReservation(reservation models.Reservation) (int, error)
	InsertRoomRestriction(roomRestriction models.RoomRestriction) error
	SearchAvailabilityByDateAndRoom(start time.Time, end time.Time, roomID int) (bool, error)
	SearchAvailabilityByDateForAllRooms(start time.Time, end time.Time) ([]models.Room, error)
	GetRoomByID(roomID int) (models.Room, error)	
}

func (r *postgresRepository) FindAllUsers() bool {
	return true
}

func NewPostgresRepository(config *config.AppConfig, db *database.Database) DatabaseRepository {
	return &postgresRepository{
		Config: config,
		DB:     db,
	}
}