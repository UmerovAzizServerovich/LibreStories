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
	"librestories/services"
	"net/http"
	"time"
)

func GetGeneralStatistics(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if r.Method != http.MethodGet {
		w.WriteHeader(405)
		w.Write([]byte("Invalid method!"))
		return
	}
	general_statics, err := services.DisplayGeneralStatistics()
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(general_statics)
	fmt.Printf("%s %s %s\n", r.Method, r.RequestURI, time.Since(start))
}
