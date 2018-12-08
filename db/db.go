package db

import (
	"database/sql"
	"fmt"
	"github.com/frnd/schedule-hub/util"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq" //import postgres
	"log"
)

//DB ...
type DB struct {
	*sql.DB
}

var (
	DbHost = util.GetEnv("DATABASE_HOST", "localhost")
	DbPort = util.GetEnv("DATABASE_PORT", "5432")
	DbUser = util.GetEnv("DATABASE_USER", "user")
	DbPassword = util.GetEnv("DATABASE_PASSWORD", "p4ssword")
	DbName = util.GetEnv("DATABASE_NAME", "schedulehub")
)

var db *gorp.DbMap

//Init ...
func Init() {

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DbHost,DbPort, DbUser, DbPassword, DbName)

	var err error
	db, err = ConnectDB(dbinfo)
	if err != nil {
		log.Fatal(err)
	}

}

//ConnectDB ...
func ConnectDB(dataSourceName string) (*gorp.DbMap, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	//dbmap.TraceOn("[gorp]", log.New(os.Stdout, "golang-gin:", log.Lmicroseconds)) //Trace database requests
	return dbmap, nil
}

//GetDB ...
func GetDB() *gorp.DbMap {
	return db
}
