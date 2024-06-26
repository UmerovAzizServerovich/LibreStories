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
)

func DisplayUsersCount() (int, error) {
	db, err := sql.Open("sql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return -1, err
	}
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM USERS WHERE Delted != 1`).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func DisplayPublicationsCount() (int, error) {
	db, err := sql.Open("sql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return -1, err
	}
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM Publications WHERE Delted != 1`).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}

func DisplayCommentsCount() (int, error) {
	db, err := sql.Open("sql", fmt.Sprintf("%s:%s@/%s",
		configs.SqlUser, configs.SqlPassword, configs.DbName))
	if err != nil {
		return -1, err
	}
	var count int
	err = db.QueryRow(`SELECT COUNT(*) FROM Comments WHERE Delted != 1`).Scan(&count)
	if err != nil {
		return -1, err
	}

	return count, nil
}
