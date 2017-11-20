package dataDir

import (
	"database/sql"
	"gallery-app/models"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

func NewDBInstance(dbSource string) {
	database, _ := sql.Open("sqlite3", dbSource)
	command :=  "CREATE TABLE IF NOT EXISTS photo (id INTEGER PRIMARY KEY NOT NULL, title STRING, description TEXT, file_name STRING,dribble_link STRING);"
	database.Exec(command)
	database.Close()
}

func LoadDataToTable(dbSource string, photo models.Photo) error {
	database, _ := sql.Open("sqlite3", dbSource)
	query, _ := database.Begin()
	command, _ := query.Prepare("Insert Into photo (id, title, description, file_name, dribble_link) values(?,?,?,?,?)")
	if _, err := command.Exec(photo.Id, photo.Title, photo.Description, photo.FileName, photo.DribbleLink.Normal); err != nil {
		fmt.Println(err)
		return nil
	} else {
		query.Commit()
	}
	return nil
}

func GetPhotoDetailByCriteria(dbSource string, title string, desc string) (*models.Photo, error) {

	database, _ := sql.Open("sqlite3", dbSource)

	var query *sql.Rows
	var err error

	photoModel := new(models.Photo)
	if len(title) > 0 && len(desc) > 0 {
		query, err = database.Query("SELECT * FROM photo WHERE title LIKE '%" + title + "%' OR description Like '%" + desc + "%'")
	} else if len(title) > 0 {
		query, err = database.Query("SELECT * FROM photo WHERE title LIKE '%" + title + "%'")
	} else if len(desc) > 0 {
		query, err = database.Query("SELECT * FROM photo WHERE description Like '%" + desc + "%'")
	}
	defer query.Close()
	database.Close()

	if err ==  nil {
		for query.Next() {
			if err := query.Scan(&photoModel.Id, &photoModel.Title, &photoModel.Description, &photoModel.FileName, &photoModel.DribbleLink.Normal); err != nil {
				panic(err)
				return nil, sql.ErrNoRows
			} else {
				return photoModel, nil
			}
		}
	}
	return nil, sql.ErrNoRows
}



