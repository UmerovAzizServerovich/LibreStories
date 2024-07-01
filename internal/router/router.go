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

package router

import (
	"librestories/handlers"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/general_statistics", handlers.GetGeneralStatistics)
	r.HandleFunc("/reg_user", handlers.RegUser)
	r.HandleFunc("/auth_user", handlers.AuthUser)
	r.HandleFunc("/add_pub", handlers.AddPublication)
	r.HandleFunc("/add_com", handlers.AddComment)
	r.HandleFunc("/display_pub", handlers.DisplayPublication)
	r.HandleFunc("/display_pubs", handlers.DisplayPublicationsIds)
	r.HandleFunc("/display_com", handlers.DisplayComment)
	r.HandleFunc("/delete_pub", handlers.DeletePublication)

}
