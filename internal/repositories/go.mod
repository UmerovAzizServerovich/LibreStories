module repositories

go 1.22.4

require (
	github.com/go-sql-driver/mysql v1.8.1
	librestories/configs v0.0.0
	librestories/models v0.0.0
)

require filippo.io/edwards25519 v1.1.0 // indirect

replace librestories/configs => ../configs

replace librestories/models => ../models
