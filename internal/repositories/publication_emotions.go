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

type PublicationEmotion models.PublicationEmotion

func InitPublicationEmotions() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS PublicationEmotions(
  				Id            INT PRIMARY KEY AUTO_INCREMENT,
				PublicationId INT,
				UserId        INT DEFAULT 0,
				EmotionType   TINYINT,
				Deleted       TINYINT(1) DEFAULT 0
			);`); err != nil {
		return err
	}
	return nil
}

func (pub_emotion *PublicationEmotion) Add() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`INSERT INTO PublicationEmotions(PublicationId, UserId,
		EmotionType) VALUES(?, ?, ?);`, pub_emotion.PublicationId, pub_emotion.UserId,
		pub_emotion.EmotionType); err != nil {
		return err
	}
	return nil
}

func (pub_emotion *PublicationEmotion) Save() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE PublicationEmotions SET EmotionType = ?
		WHERE Id = ?;`,
		pub_emotion.EmotionType, pub_emotion.Id); err != nil {
		return err
	}
	return nil
}

func (pub_emotion *PublicationEmotion) Delete() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`DELETE FROM PublicationEmotions WHERE Id = ?;`,
		pub_emotion.Id); err != nil {
		return err
	}
	return nil
}
