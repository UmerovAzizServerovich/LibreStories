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

type Comment models.Comment

func InitComments() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Comments(
  				Id            INT PRIMARY KEY AUTO_INCREMENT,
				PublicationId INT,
  				AuthorId      INT,
	  			CreationDate  DATETIME,
  				Likes         INT DEFAULT 0,
				Dislikes      INT DEFAULT 0,
  				Content       TEXT,
				Deleted       TINYINT(1) DEFAULT 0
			);`); err != nil {
		return err
	}
	return nil
}

func (com *Comment) Add() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.QueryRow(`SELECT Id FROM Users WHERE UserName = ? AND Deleted != 1;`,
		com.AuthorName).Scan(&com.AuthorId); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO Comments(AuthorId, CreationDate,
		PublicationId, Content) VALUES(?, ?, ?, ?);`,
		com.AuthorId, time.Now(), com.PublicationId, com.Content); err != nil {
		return err
	}

	var coms_count int
	if err = db.QueryRow(`SELECT CommentsCount FROM Publications
		WHERE Id = ? AND Deleted != 1;`,
		com.PublicationId).Scan(&coms_count); err != nil {
		return err
	}

	coms_count++
	if _, err := db.Exec(`UPDATE Publications SET CommentsCount = ?
		WHERE Id = ?;`, coms_count, com.PublicationId); err != nil {
		return err
	}
	return nil
}

func (com *Comment) View() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if err = db.QueryRow(`SELECT * FROM Comments
		WHERE Id = ? AND Deleted != 1;`,
		com.Id).Scan(&com.Id, &com.AuthorId, &com.CreationDateTime, &com.Likes,
		&com.Dislikes, &com.Content); err != nil {
		return err
	}

	return nil
}

func (com *Comment) Save() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Comments SET Content = '?'
		WHERE Id = ? AND Deleted != 1;`,
		com.Content, com.Id); err != nil {
		return err
	}
	return nil
}

func (com *Comment) Delete() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Comments SET Deleted = 1
		WHERE Id = ?;`,
		com.Id); err != nil {
		return err
	}
	if _, err := db.Exec(`UPDATE CommentEmotions SET Deleted = 1
		WHERE CommentId = ?;`,
		com.Id); err != nil {
		return err
	}
	return nil
}

func (com *Comment) Recover() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Comments SET Deleted = 0
		WHERE Id = ?;`,
		com.Id); err != nil {
		return err
	}
	if _, err := db.Exec(`UPDATE CommentEmotions SET Deleted = 0
		WHERE CommentId = ?;`,
		com.Id); err != nil {
		return err
	}
	return nil
}
