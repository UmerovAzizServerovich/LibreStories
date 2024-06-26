// LibreStories - A simple web application for sharing stories
// Copyright (C) 2024 Umerov Aziz Serverovich
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package repositories

import (
	"database/sql"
	"fmt"
	"librestories/configs"
	"librestories/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Publication models.Publication

func InitPublications() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Publications(
  				Id               INT PRIMARY KEY AUTO_INCREMENT,
				AuthorId         INT DEFAULT 0,ы
				Name             VARCHAR(100) NOT NULL,
				Description      VARCHAR(300),
				CreationDateTime DATETIME,
				Likes            INT DEFAULT 0,
				Dislikes         INT DEFAULT 0,
				Category         INT DEFAULT 0,
				Images           VARCHAR(200),
				Content          TEXT,
				Deleted          TINYINT(1)
			);`); err != nil {
		return err
	}
	return nil
}

func (pub *Publication) Add() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`INSERT INTO Publications(AuthorId, Name, Description,
		Сategory, CreationDate, Images, Content) VALUES(?, ?, ?, ?, ?, ?);`,
		(*pub).AuthorId, (*pub).Name, (*pub).Description, (*pub).Category,
		time.Now(), (*pub).Images, (*pub).Content); err != nil {
		return err
	}
	return nil
}

func (pub *Publication) View() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.QueryRow(`SELECT * FROM Publications WHERE Id = ?  AND Deleted != 1;`,
		(*pub).Id).Scan((*pub).Id, (*pub).AuthorId, (*pub).Name, (*pub).Description,
		(*pub).CreationDateTime, (*pub).Likes, (*pub).Dislikes, (*pub).Category,
		(*pub).Content); err != nil {
		return err
	}

	return nil
}

func (pub *Publication) Save() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Publications SET Name = ?, Description = ?,
		Сategory = ?, Content = ? WHERE Id = ? AND Deleted != 1;`,
		(*pub).Name, (*pub).Description, (*pub).Category, (*pub).Content,
		(*pub).Id); err != nil {
		return err
	}
	return nil
}

func (pub *Publication) Delete() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Publications SET Deleted = 1
		WHERE Id = ?;`, pub.Id); err != nil {
		return err
	}
	if _, err := db.Exec(`UPDATE PublicationEmotions SET Deleted = 1
		WHERE CommentId = ?;`,
		pub.Id); err != nil {
		return err
	}
	return nil
}

func (pub *Publication) Recover() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Publications SET Deleted = 0
		WHERE Id = ?;`, pub.Id); err != nil {
		return err
	}
	if _, err := db.Exec(`UPDATE PublicationEmotions SET Deleted = 0
		WHERE CommentId = ?;`,
		pub.Id); err != nil {
		return err
	}
	return nil
}
