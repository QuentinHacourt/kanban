package database

import (
	"database/sql"
	"fmt"
	"kanban/models"
	"log"
)

func InsertTeam(teamInput models.TeamInput) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		INSERT INTO teams(name)
		VALUES ($1)
		RETURNING id
	`

	var id int64

	if err := db.QueryRow(
		sqlStatement,
		teamInput.Name,
	).Scan(&id); err != nil {
		log.Fatalf("Unable to insert new team: %v", err)
	}

	fmt.Printf("Inserted a single record with id: %v", id)

	return id
}

func GetTeam(id int64) (models.Team, error) {
	db := createConnection()

	defer db.Close()

	var team models.Team

	sqlStatement := `
		SELECT id, name
		FROM teams
		where id = $1;
`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&team.ID, &team.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return team, nil
	case nil:
		return team, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	return team, err
}

func GetAllTeams() ([]models.Team, error) {
	db := createConnection()

	defer db.Close()

	var teams []models.Team

	sqlStatement := `
		SELECT id, name
		FROM teams
	`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to get all teams: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var team models.Team

		err = rows.Scan(&team.ID, &team.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}

		teams = append(teams, team)
	}

	return teams, err
}

func UpdateTeam(team models.Team) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		UPDATE teams
		SET
			name=$2
		WHERE
			id =$1
	`

	res, err := db.Exec(sqlStatement, team.ID, team.Name)

	if err != nil {
		log.Fatalf("Unable to update team: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteTeam(id int64) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM teams WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete team: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected: %v", rowsAffected)

	return rowsAffected
}
