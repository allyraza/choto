package choto

import (
	"database/sql"
	"log"
	"net/http"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Choto :
type Choto struct {
	Mux      *http.ServeMux
	Database *sql.DB
}

// New :
func New(config *Config) *Choto {

	app := &Choto{
		Database: newDatabase(),
	}

	app.Routes()

	return app
}

func (api *Choto) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.Mux.ServeHTTP(w, r)
}

// @todo: move this to a package
func newDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root:@/choto?parseTime=true")
	failOnError(err)

	// ping the database to check server connectivity
	failOnError(db.Ping())

	return db
}

func failOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
