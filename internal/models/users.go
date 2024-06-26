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

package models

type User struct {
	Id           int    `json:"id,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	Password     string `json:"password,omitempty"`
	About        string `json:"about,omitempty"`
	CreationDate string `json:"creation_date,omitempty"`
	Role         string `json:"role,omitempty"`
	Avatar       string `json:"avatar,omitempty"`
	DateOfBirth  string `json:"date_of_birth,omitempty"`
	Gender       int    `json:"gender,omitempty"`
	AdminLVL     int    `json:"admin_lvl,omitempty"`
}
