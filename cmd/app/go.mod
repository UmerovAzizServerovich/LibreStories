module librestories

go 1.22.4


require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/gorilla/mux v1.8.1
	librestories/configs v0.0.0 // indirect
	librestories/models v0.0.0 // indirect
	librestories/repositories v0.0.0
	librestories/handlers v0.0.0 // indirect
	librestories/services v0.0.0 // indirect
	librestories/router v0.0.0
)

replace librestories/configs => ../../internal/configs

replace librestories/models => ../../internal/models

replace librestories/repositories => ../../internal/repositories

replace librestories/services => ../../internal/services

replace librestories/handlers => ../../internal/handlers

replace librestories/router => ../../internal/router
