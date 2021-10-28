package database

import (
	"database/sql"
	"fmt"
	"kanban/models"
	"log"
)

func InsertProject(projectInput models.ProjectInput) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		INSERT INTO projects(title, description, team_id)
		SELECT $1, $2, t.id
		FROM teams t
		WHERE name = $3
		RETURNING id
	`

	var id int64

	if err := db.QueryRow(
		sqlStatement,
		projectInput.Title,
		projectInput.Description,
		projectInput.TeamName,
	).Scan(&id); err != nil {
		log.Fatalf("Unable to insert new project: %v", err)
	}

	fmt.Printf("Inserted a single record with id: %v", id)

	return id
}

func GetProject(id int64) (models.Project, error) {
	db := createConnection()

	defer db.Close()

	var project models.Project

	sqlStatement := `
		SELECT p.id, p.title, p.description, t.name
		FROM projects p
			LEFT OUTER JOIN teams t on p.team_id = t.id
		where p.id = $1;
`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&project.ID, &project.Title, &project.Description, &project.TeamName)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return project, nil
	case nil:
		return project, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	return project, err
}

func GetAllProjects() ([]models.Project, error) {
	db := createConnection()

	defer db.Close()

	var projects []models.Project

	sqlStatement := `
		SELECT p.id, p.title, p.description, t.name
		FROM projects p
			LEFT OUTER JOIN teams t on p.team_id = t.id
	`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to get all projects: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var project models.Project

		err = rows.Scan(&project.ID, &project.Title, &project.Description, &project.TeamName)

		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}

		projects = append(projects, project)
	}

	return projects, err
}

func UpdateProject(project models.Project) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		WITH get_ids AS (
			SELECT id team_id
			FROM teams
			WHERE name = $4
		)
		UPDATE projects
		SET
			title=$2,
			description=$3,
			team_id=get_ids.team_id
		FROM
			get_ids
		WHERE
			id =$1
	`

	res, err := db.Exec(sqlStatement, project.ID, project.Title, project.Description, project.TeamName)

	if err != nil {
		log.Fatalf("Unable to update project: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteProject(id int64) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM projects WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete project: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected: %v", rowsAffected)

	return rowsAffected
}
