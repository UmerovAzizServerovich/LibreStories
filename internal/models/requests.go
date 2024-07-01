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

type UserRequest struct {
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type PublicationRequest struct {
	Publication Publication `json:"publication,omitempty"`
	User        UserRequest `json:"user,omitempty"`
}

type CommentRequest struct {
	Comment Comment     `json:"comment,omitempty"`
	User    UserRequest `json:"user,omitempty"`
}

type PublicationEmotionRequest struct {
	PublicationEmotion PublicationEmotion `json:"publication_emotion,omitempty"`
	User               UserRequest        `json:"user,omitempty"`
}

type CommentEmotionRequest struct {
	CommentEmotion CommentEmotion `json:"comment_emotion,omitempty"`
	User           UserRequest    `json:"user,omitempty"`
}
