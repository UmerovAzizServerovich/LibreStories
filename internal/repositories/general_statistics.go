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
