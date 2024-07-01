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

package handlers

import (
	"encoding/json"
	"fmt"
	"librestories/models"
	"librestories/services"
	"net/http"
	"time"
)

func AddComment(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Invalid method!"))
		return
	}

	var com_req models.CommentRequest
	if err := json.NewDecoder(r.Body).Decode(&com_req); err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Invalid Input"))
		return
	} else if com_req.User.UserName == "" || com_req.User.Password == "" ||
		com_req.Comment.Content == "" || com_req.Comment.PublicationId <= 0 {
		w.WriteHeader(400)
		w.Write([]byte("Invalid Input"))
		return
	}
	result, err := services.AddComment(com_req)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}
	if !result {
		w.WriteHeader(400)
		w.Write([]byte("Invalid UserName or Password"))
		return
	}
	w.Write([]byte("The comment is added"))
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, time.Since(start))
}
