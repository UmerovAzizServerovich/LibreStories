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
	"database/sql"
	"encoding/json"
	"fmt"
	"librestories/repositories"
	"librestories/services"
	"net/http"
	"time"
)

func DisplayPublication(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Invalid method!"))
		return
	}
	var pub repositories.Publication
	if err := json.NewDecoder(r.Body).Decode(&pub); err != nil {
		fmt.Println(err)
		w.WriteHeader(400)
		w.Write([]byte("Invalid Input"))
		return
	}
	pub, err := services.DisplayPublication(pub)
	if err == sql.ErrNoRows {
		w.Write([]byte("Nothing was found"))
		return
	} else if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Internal Server Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pub)
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, time.Since(start))
}
