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

	_ "github.com/go-sql-driver/mysql"
)

type CommentEmotion models.CommentEmotion

func InitCommentEmotions() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS CommentEmotions(
  				Id          INT PRIMARY KEY AUTO_INCREMENT,
				CommentId   INT,
				UserId      INT,
				EmotionType TINYINT,
				Deleted     TINYINT(1) DEFAULT 0
			);`); err != nil {
		return err
	}
	return nil
}

func (com_emotion *CommentEmotion) Add() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`INSERT INTO CommentEmotions(CommentId, UserId, EmotionType)
		VALUES(?, ?, ?);`, com_emotion.CommentId, com_emotion.UserId,
		com_emotion.EmotionType); err != nil {
		return err
	}
	return nil
}

func (com_emotion *CommentEmotion) Save() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE CommentEmotions SET EmotionType = ?
		WHERE Id = ?;`,
		com_emotion.EmotionType, com_emotion.Id); err != nil {
		return err
	}
	return nil
}

func (com_emotion *CommentEmotion) Delete() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`DELETE FROM CommentEmotions WHERE Id = ?;`,
		com_emotion.Id); err != nil {
		return err
	}
	return nil
}
