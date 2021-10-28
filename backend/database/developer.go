package database

import (
	"database/sql"
	"fmt"
	"kanban/models"
	"log"
)

func InsertDeveloper(developerInput models.DeveloperInput) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		INSERT INTO developers(user_name, password, team_id)
		SELECT $1, $2, t.id
		FROM teams t
		WHERE name = $3
		RETURNING id
	`

	var id int64

	if err := db.QueryRow(
		sqlStatement,
		developerInput.UserName,
		developerInput.Password,
		developerInput.TeamName,
	).Scan(&id); err != nil {
		log.Fatalf("Unable to insert new developer: %v", err)
	}

	fmt.Printf("Inserted a single record with id: %v", id)

	return id
}

func GetDeveloper(id int64) (models.Developer, error) {
	db := createConnection()

	defer db.Close()

	var developer models.Developer

	sqlStatement := `
		SELECT id, user_name, password
		FROM developers
		where id = $1;
`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&developer.ID, &developer.UserName, &developer.Password)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return developer, nil
	case nil:
		return developer, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	return developer, err
}

func GetAllDevelopers() ([]models.Developer, error) {
	db := createConnection()

	defer db.Close()

	var developers []models.Developer

	sqlStatement := `
		SELECT id, user_name, password
		FROM developers
	`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to get all developers: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var developer models.Developer

		err = rows.Scan(&developer.ID, &developer.UserName, &developer.Password)

		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}

		developers = append(developers, developer)
	}

	return developers, err
}

func UpdateDeveloper(developer models.Developer) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		UPDATE developers
		SET
			user_name=$2,
			password=$3
		WHERE
			id = $1;
	`

	res, err := db.Exec(sqlStatement, developer.ID, developer.UserName, developer.Password)

	if err != nil {
		log.Fatalf("Unable to update developer: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteDeveloper(id int64) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM developers WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete developer: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected: %v", rowsAffected)

	return rowsAffected
}
