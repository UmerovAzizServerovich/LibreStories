package services

import (
	"librestories/repositories"
)

func DisplayGeneralStatistics() (repositories.GeneralStatistics, error) {
	var stats repositories.GeneralStatistics
	if err := stats.DisplayUsersCount(); err != nil {
		return repositories.GeneralStatistics{}, err
	}
	if err := stats.DisplayPublicationsCount(); err != nil {
		return repositories.GeneralStatistics{}, err
	}
	if err := stats.DisplayCommentsCount(); err != nil {
		return repositories.GeneralStatistics{}, err
	}
	return stats, nil
}
