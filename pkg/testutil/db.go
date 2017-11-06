package testutil

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // import postgres driver
)

type dbParams struct {
	hostname string
	database string
	username string
	password string
}

func (p dbParams) String() string {
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s",
		p.hostname, p.database, p.username, p.password)
}

func GetDB() (*sql.DB, error) {
	params := GetParams()
	conInfo := fmt.Sprintf("%s", params)
	db, err := sql.Open("postgres", conInfo)

	if err != nil {
		panic(err)
	}

	_, err = db.Exec("truncate vinodes.nodes cascade")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("truncate viitems.folders cascade")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("truncate viitems.reports cascade")
	if err != nil {
		panic(err)
	}

	return db, nil
}

func GetParams() dbParams {
	if len(os.Getenv("SEMAPHORE")) > 0 {
		return getSemaphoreParams()
	}
	if len(os.Getenv("TRAVIS")) > 0 {
		return getTravisParams()
	}
	return getLocalParams()
}

func getSemaphoreParams() dbParams {
	hostname := "localhost"
	database := "vizigoto"
	username := os.Getenv("DATABASE_POSTGRESQL_USERNAME")
	password := os.Getenv("DATABASE_POSTGRESQL_PASSWORD")
	return dbParams{hostname, database, username, password}
}

func getTravisParams() dbParams {
	hostname := "localhost"
	database := "travis_ci_test"
	username := "postgres"
	password := ""
	return dbParams{hostname, database, username, password}
}

func getLocalParams() dbParams {
	hostname := os.Getenv("PGHOSTNAME")
	database := os.Getenv("PGDATABASE")
	username := os.Getenv("PGUSERNAME")
	password := os.Getenv("PGPASSWORD")
	return dbParams{hostname, database, username, password}
}
