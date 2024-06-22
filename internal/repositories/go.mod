module repositories

go 1.22.4

require (
	github.com/go-sql-driver/mysql v1.8.1
	miniforum/config v0.0.0
	miniforum/models v0.0.0
)

require filippo.io/edwards25519 v1.1.0 // indirect

replace miniforum/config => ../config

replace miniforum/models => ../models
