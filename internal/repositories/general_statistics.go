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

type GeneralStatistics models.GeneralStatistics

func (stats *GeneralStatistics) DisplayUsersCount() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}

	err = db.QueryRow(`SELECT COUNT(*) FROM Users
		WHERE Deleted != 1`).Scan(&stats.UsersCount)
	if err != nil {
		return err
	}

	return nil
}

func (stats *GeneralStatistics) DisplayPublicationsCount() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}

	err = db.QueryRow(`SELECT COUNT(*) FROM Publications
		WHERE Deleted != 1`).Scan(&stats.PublicationsCount)
	if err != nil {
		return err
	}

	return nil
}

func (stats *GeneralStatistics) DisplayCommentsCount() error {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return err
	}

	err = db.QueryRow(`SELECT COUNT(*) FROM Comments
		WHERE Deleted != 1`).Scan(&stats.CommentsCount)
	if err != nil {
		return err
	}

	return nil
}
