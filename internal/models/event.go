package models

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NewEvent struct {
	Creator     string `json:"creator" gorm:"type:varchar(255)"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Event struct {
	Id          string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	Creator     string `json:"creator" gorm:"type:varchar(255)"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`

	EventCreator User `gorm:"foreignKey:Creator"`
}

func (e *Event) BeforeCreate(tx *gorm.DB) error {
	e.Id = uuid.NewString()

	return nil
}

type InterestedEvent struct {
	gorm.Model
	Id      string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	UserId  string `json:"user_id" gorm:"type:varchar(255)"`
	EventId string `json:"event_id" gorm:"type:varchar(255)"`

	User  User  `gorm:"foreignKey:UserId"`
	Event Event `gorm:"foreignKey:EventId"`
}

func (iE *InterestedEvent) BeforeCreate(tx *gorm.DB) error {
	iE.Id = uuid.NewString()

	return nil
}

type GroupEvent struct {
	gorm.Model
	Id      string `json:"id" gorm:"primaryKey;type:varchar(255)"`
	GroupId string `json:"group_id" gorm:"type:varchar(255)"`
	EventId string `json:"event_id" gorm:"type:varchar(255)"`

	Group Group `gorm:"foreignKey:GroupId"`
	Event Event `gorm:"foreignKey:EventId"`
}

func (gE *GroupEvent) BeforeCreate(tx *gorm.DB) error {
	gE.Id = uuid.NewString()

	return nil
}

// func FindUserById(id string, tx *gorm.DB) (User, error) {
// 	var user User
// 	err := tx.Find(&user, "id = ?", id)
// 	if err.Error != nil {
// 		return User{}, err.Error
// 	}
// 	return user, nil
// }

func CreateEvent(tx *gorm.DB, event *NewEvent) (*Event, error) {

	request := Event{
		Creator:     event.Creator,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartDate:   event.StartDate,
		EndDate:     event.EndDate,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
	}

	err := tx.Model(Event{}).Create(&request)

	fmt.Print(event)

	if err.Error != nil {
		fmt.Print(err)

		return &Event{}, err.Error
	}
	return &request, nil
}

// ListAllEvents retrieves all events.
func ListEvents(tx *gorm.DB) ([]Event, error) {
	var events []Event

	// Your database query to retrieve all events and sort them.
	// This query depends on your database structure and ORM (GORM in this case).
	// You can use the Order method to sort the results.

	// Example: Query all events and sort them by start date and start time using GORM.
	err := tx.Order("start_date, start_time").Find(&events).Error

	if err != nil {
		return nil, err
	}

	return events, nil
}
