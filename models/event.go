package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	UserID      int64
	Name        string `binding:"required"`
	Location    string `binding:"required"`
	Description string `binding:"required"`
	DateTime    time.Time
}

// var events = []Event{}

func (e *Event) Save() error {
	query := `insert into events(name,location,description,date_time,user_id)values(?,?,?,?,?)`
	stmt, error := db.DB.Prepare(query)
	if error != nil {
		return error
	}
	defer stmt.Close()
	result, error := stmt.Exec(e.Name, e.Location, e.Description, e.DateTime, e.UserID)
	if error != nil {
		return error
	}
	id, error := result.LastInsertId()

	// later save to database
	e.ID = id
	return error
}

func GetAllEvents() ([]Event, error) {
	query := "select * from events"
	rows, error := db.DB.Query(query)
	if error != nil {
		return nil, error
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		error := rows.Scan(&event.ID, &event.Name, &event.Location, &event.Description, &event.DateTime, &event.UserID)
		if error != nil {
			return nil, error
		}
		events = append(events, event)
	}
	return events, nil
}

func GetSingleEvent(id int64) (*Event, error) {
	query := "select * from events where id=?"
	row := db.DB.QueryRow(query, id)

	var event Event
	error := row.Scan(&event.ID, &event.Name, &event.Location, &event.Description, &event.DateTime, &event.UserID)
	if error != nil {
		return nil, error
	}

	return &event, nil
}

func (e Event) Update() error {
	query := `update events set name=?,description=?,location=?,date_time=? where id=?`
	stmt, error := db.DB.Prepare(query)
	if error != nil {
		return error
	}
	defer stmt.Close()
	_, error = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.ID)
	return error
}

func (e Event) Delete() error {
	query := "delete from events where id=?"
	stmt, error := db.DB.Prepare(query)
	if error != nil {
		return error
	}
	defer stmt.Close()
	_, error = stmt.Exec(e.ID)
	return error
}

func (e Event) Register(userId int64) error {
	query := "insert into registrations(event_id,user_id)values(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "delete from registrations where event_id=? and user_id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err
}
