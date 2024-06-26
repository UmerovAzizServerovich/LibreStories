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

type User models.User

func InitUsers() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS Users(
  				Id           INT PRIMARY KEY AUTO_INCREMENT,
  				UserName     VARCHAR(30) NOT NULL,
				Password     VARCHAR(30) NOT NULL,
				About        VARCHAR(200)
  				CreationDate DATE,
				Role         VARCHAR(30),
				Avatar       VARCHAR(100),
				DateOfBirth  DATE,
				Gender       TINYINT DEFAULT 0,
		  		AdminLVL     INT NOT NULL DEFAULT 0,
				Deleted      TINYINT(1)
			);`); err != nil {
		return err
	}
	return nil
}

func (user *User) Add() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return (err)
	}
	defer db.Close()

	if _, err = db.Exec(`INSERT INTO Users(UserName, Password, CreationDate)
		VALUES(?, ?, ?)`, (*user).UserName, (*user).Password,
		time.Now()); err != nil {
		return err
	}
	return nil
}

func (user *User) CheckPassword() (bool, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return false, err
	}
	defer db.Close()

	var true_password string
	if err := db.QueryRow(`SELECT Password FROM Users
		WHERE Id = ? AND Deleted != 1;`,
		(*user).Id).Scan(&true_password); err != nil {
		return false, err
	}
	if (*user).Password == true_password {
		return true, nil
	}
	return false, nil
}

func (user *User) View() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.QueryRow(`SELECT Id, UserName, About, CreationDate, Role, Avatar,
		DateOfBirth, Gender, AdminLVL FROM Users WHERE Id = ? AND Deleted != 1;`,
		(*user).Id).Scan(
		(*user).Id, (*user).UserName, (*user).About, (*user).CreationDate, (*user).Role,
		(*user).Avatar, (*user).DateOfBirth, (*user).Gender, (*user).AdminLVL)
	if err != nil {
		return err
	}

	return nil
}

func (user *User) Save() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Users SET UserName = ?, Password = ?, About = ?,
		Role, Avatar, DateOfBirth, Gender, AdminLVL
		WHERE Id = ?;`,
		(*user).UserName, (*user).Password, (*user).About, (*user).Role,
		(*user).Avatar, (*user).DateOfBirth, (*user).Gender, (*user).AdminLVL,
		(*user).Id); err != nil {
		return err
	}
	return nil
}

func (user *User) Delete() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Users SET Deleted = 1
		WHERE Id = ?;`, user.Id); err != nil {
		return err
	}
	return nil
}

func (user *User) Recover() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}
	defer db.Close()

	if _, err := db.Exec(`UPDATE Users SET Deleted = 0
		WHERE Id = ?;`, user.Id); err != nil {
		return err
	}
	return nil
}

func (user *User) DisplayOwnPublications(start, count int) ([]int, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	own_pub_ids := make([]int, count)
	for i := 0; i < count; i++ {
		err = db.QueryRow(`SELECT Id FROM Publications WHERE AuthorId = ?
			AND Deleted != 1 ORDER BY Id DESC;`,
			(*user).Id).Scan(&(own_pub_ids[i]))
		if err != nil {
			return nil, err
		}
	}

	return own_pub_ids, nil
}

func (user *User) DisplayOwnComments(start, count int) ([]int, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	own_com_ids := make([]int, count)
	for i := 0; i < count; i++ {
		err = db.QueryRow(`SELECT Id FROM Comments WHERE AuthorId = ?
			AND Deleted != 1 ORDER BY Id DESC;`,
			(*user).Id).Scan(&(own_com_ids[i]))
		if err != nil {
			return nil, err
		}
	}

	return own_com_ids, nil
}

func (user *User) DisplayFavoritePublications(start, count int) ([]int, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	favorite_pub_ids := make([]int, count)
	for i := 0; i < count; i++ {
		err = db.QueryRow(`SELECT PublicationId FROM PublicationEmotions
			WHERE UserId = ? AND Deleted != 1 ORDER BY Id DESC;`,
			(*user).Id).Scan(&(favorite_pub_ids[i]))
		if err != nil {
			return nil, err
		}
	}

	return favorite_pub_ids, nil
}

func (user *User) DisplayFavoriteComments(start, count int) ([]int, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	favorite_com_ids := make([]int, count)
	for i := 0; i < count; i++ {
		err = db.QueryRow(`SELECT CommentId FROM CommentEmotions WHERE UserId = ?
			AND Deleted != 1 ORDER BY Id DESC;`,
			(*user).Id).Scan(&(favorite_com_ids[i]))
		if err != nil {
			return nil, err
		}
	}

	return favorite_com_ids, nil
}
