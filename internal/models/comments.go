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

type Comment struct {
	Id               int    `json:"id,omitempty"`
	PublicationId    int    `json:"publication_id,omitempty"`
	AuthorId         int    `json:"author_id,omitempty"`
	AuthorName       string `json:"author_name,omitempty"`
	AuthorPassword   string `json:"author_password,omitempty"`
	CreationDateTime string `json:"creation_date_time,omitempty"`
	Likes            int    `json:"likes,omitempty"`
	Dislikes         int    `json:"dislikes,omitempty"`
	Content          string `json:"content,omitempty"`
}
