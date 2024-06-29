module handlers

go 1.22.4

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	librestories/configs v0.0.0 // indirect
	librestories/models v0.0.0 // indirect
	librestories/repositories v0.0.0
	librestories/services v0.0.0

)

replace librestories/repositories => ../repositories

replace librestories/configs => ../configs

replace librestories/models => ../models

replace librestories/services => ../services
