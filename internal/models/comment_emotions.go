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

type CommentEmotion struct {
	Id           int    `json:"id,omitempty"`
	CommentId    int    `json:"comment_id,omitempty"`
	UserId       int    `json:"user_id,omitempty"`
	UserName     string `json:"user_name,omitempty"`
	UserPassword string `json:"user_password,omitempty"`
	EmotionType  int    `json:"emotion_type,omitempty"`
}
