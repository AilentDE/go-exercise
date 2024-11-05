package models

import (
	"restful-api-sample/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query) //[TIPS] Use prepare could reuse this query in memory (for modify data)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.Id = id
	return nil
}

func (e Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)
	if err != nil {
		return err
	}

	return nil
}

func (e Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id)
	if err != nil {
		return err
	}

	return nil
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events
	`
	rows, err := db.DB.Query(query) //[TIPS] Query call one time (for easy query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId); err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
	SELECT * FROM events
	WHERE id=?
	`
	row := db.DB.QueryRow(query, id) // for only one result

	var event Event
	if err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId); err != nil {
		return nil, err
	}

	return &event, nil
}

// register
func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO registerations(event_id, user_id)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)
	return err
}

func (e Event) CancelRegister(userId int64) error {
	query := `
	DELETE FROM registerations
	WHERE event_id = ? AND user_id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.Id, userId)
	return err
}
