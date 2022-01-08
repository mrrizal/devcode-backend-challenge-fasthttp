package models

import (
	"errors"
	"fmt"
	"net/mail"
	"time"

	"github.com/mrrizal/devcode-backend-challenge-fasthttp/database"
)

type ActivityGroup struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (activityGroup ActivityGroup) Validate() error {
	if activityGroup.Title == "" {
		return errors.New("title cannot be null")
	}

	if activityGroup.Email != "" {
		_, err := mail.ParseAddress(activityGroup.Email)
		if err != nil {
			return errors.New("invalid email address")
		}
	}

	return nil
}

func (activityGroup ActivityGroup) Insert() (int, error) {
	db := database.DBConn
	stmt, err := db.Prepare("insert into activities (email, title, created_at, updated_at) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	resp, err := stmt.Exec(activityGroup.Email, activityGroup.Title, activityGroup.CreatedAt, activityGroup.UpdatedAt)
	if err != nil {
		return 0, err
	}

	insertedID, err := resp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(insertedID), nil
}

func (ActivityGroup) GetByID(activityID int, activityGroup *ActivityGroup) error {
	db := database.DBConn
	stmt, err := db.Prepare("select id, email, title, created_at, updated_at from activities where deleted_at is null and id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query(activityID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&activityGroup.ID, &activityGroup.Email, &activityGroup.Title, &activityGroup.CreatedAt,
			&activityGroup.UpdatedAt)
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if activityGroup.ID == 0 {
		errMessage := fmt.Sprintf("Activity with ID %d Not Found", activityID)
		return errors.New(errMessage)
	}
	return nil
}

func (ActivityGroup) GetAll(activityGroups *[]ActivityGroup) error {
	db := database.DBConn
	stmt, err := db.Prepare("select id, email, title, created_at, updated_at from activities where deleted_at is null")
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		activityGroup := new(ActivityGroup)
		rows.Scan(&activityGroup.ID, &activityGroup.Email, &activityGroup.Title, &activityGroup.CreatedAt,
			&activityGroup.UpdatedAt)
		*activityGroups = append(*activityGroups, *activityGroup)
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
