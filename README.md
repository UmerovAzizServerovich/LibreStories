# My Forum Project

This is the backend of a simple forum project created as a pet project.

## Features

- Registration, authorization and pseudo-deletion of the user
- Adding and changing user data
- Publishing, modifying, and pseudo-deleting posts
- Publishing, modifying, and pseudo-deleting comments

## Installation
This project depends on golang, go-sql-driver/mysql, gorilla/mux
Install golang and, being in the root directory of the project, run the following script in the terminal to install dependencies:<Enter> 
`go get github.com/go-sql-driver/mysql && go get github.com/gorilla/mux`

## Usage
Being in the root directory of the project, run the following script in the terminal to starting the project:<Enter> 
`cd cmd/app && go run main.go`

## License

This project is licensed under the GNU Affero General Public License v3.0 - see the [LICENSE](LICENSE) file for details.
