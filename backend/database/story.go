package database

import (
	"database/sql"
	"fmt"
	"kanban/models"
	"log"
)

func InsertStory(storyInput models.StoryInput) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		INSERT INTO stories(title, description, status_id)
		VALUES ($1, $2, 4)
		RETURNING id
`

	var id int64

	if err := db.QueryRow(
		sqlStatement,
		storyInput.Title,
		storyInput.Description,
	).Scan(&id); err != nil {
		log.Fatalf("Unable to insert new story: %v", err)
	}

	fmt.Printf("Inserted a single record with id: %v", id)

	return id
}

func GetStory(id int64) (models.Story, error) {
	db := createConnection()

	defer db.Close()

	var story models.Story

	sqlStatement := `
		SELECT stories.id, stories.title, stories.description, status.name
		FROM stories join status on status.id = stories.status_id
		where stories.id = $1;
`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&story.ID, &story.Title, &story.Description, &story.Status)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return story, nil
	case nil:
		return story, nil
	default:
		log.Fatalf("Unable to scan the row: %v", err)
	}

	return story, err
}

func GetAllStories() ([]models.Story, error) {
	db := createConnection()

	defer db.Close()

	var stories []models.Story

	sqlStatement := `
		SELECT stories.id, stories.title, stories.description, status.name
		FROM stories join status on status.id = stories.status_id;
	`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var story models.Story

		err = rows.Scan(&story.ID, &story.Title, &story.Description, &story.Status)

		if err != nil {
			log.Fatalf("Unable to scan the row: %v", err)
		}

		stories = append(stories, story)
	}

	return stories, err
}

func UpdateStory(story models.Story) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `
		WITH get_id AS (
			SELECT id
				FROM status
				WHERE name = $4
		)
		UPDATE stories
		SET
			title=$2,
			description=$3,
			status_id=get_id.id
		FROM
			get_id
		WHERE
			stories.id = $1;
	`

	res, err := db.Exec(sqlStatement, story.ID, story.Title, story.Description, story.Status)

	if err != nil {
		log.Fatalf("Unable to update story: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func DeleteStory(id int64) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `DELETE FROM stories WHERE id=$1`

	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to delete story: %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows: %v", err)
	}

	fmt.Printf("Total rows/record affected: %v", rowsAffected)

	return rowsAffected
}
