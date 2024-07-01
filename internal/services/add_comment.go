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

package services

import (
	"fmt"
	"librestories/models"
	"librestories/repositories"
)

func AddComment(com_req models.CommentRequest) (bool, error) {
	user := repositories.User{
		UserName: com_req.User.UserName,
		Password: com_req.User.Password,
	}
	result, err := user.CheckPassword()
	if !result {
		return false, nil
	} else if err != nil {
		fmt.Println(err)
		return false, err
	}

	com := repositories.Comment{
		AuthorId:      user.Id,
		PublicationId: com_req.Comment.PublicationId,
		Content:       com_req.Comment.Content,
	}
	if err := com.Add(); err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}
