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
		INSERT INTO stories(title, description, status_id, estimated_time, developer_id, project_id)
		SELECT $1, $2, 4, $3, d.id, p.id
		FROM developers d, projects p
		where d.user_name = $4
		AND p.title = $5
	`

	var id int64

	if err := db.QueryRow(
		sqlStatement,
		storyInput.Title,
		storyInput.Description,
		storyInput.Time,
		storyInput.DeveloperName,
		storyInput.ProjectName,
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
		select stories.id, stories.title, stories.description, status.name, stories.estimated_time, d.user_name, p.title
		from stories
			left outer join developers d on stories.developer_id = d.id
			join status on status.id = stories.status_id
			left outer join projects p on stories.project_id = p.id
		where stories.id = $1
	`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&story.ID, &story.Title, &story.Description, &story.Status, &story.Time, &story.DeveloperName, &story.ProjectName)

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
		select stories.id, stories.title, stories.description, status.name, stories.estimated_time, d.user_name, p.title
		from stories
			left outer join developers d on stories.developer_id = d.id
			join status on status.id = stories.status_id
			left outer join projects p on stories.project_id = p.id
	`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to get all stories: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var story models.Story

		err = rows.Scan(&story.ID, &story.Title, &story.Description, &story.Status, &story.Time, &story.DeveloperName, &story.ProjectName)

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
		WITH get_ids AS (
			SELECT status.id status_id, developers.id developer_id, projects.id project_id
			FROM status, developers, projects
			WHERE status.name = $4
			AND developers.user_name = $6
			AND projects.title = $7
		)
		UPDATE stories
		SET
			title=$2,
			description=$3,
			estimated_time=$5,
			status_id=get_ids.status_id,
			developer_id=get_ids.developer_id,
			project_id=get_ids.project_id
		FROM
			get_ids
		WHERE
			stories.id = $1;
	`

	res, err := db.Exec(sqlStatement, story.ID, story.Title, story.Description, story.Status, story.Time, story.DeveloperName, story.ProjectName)

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
