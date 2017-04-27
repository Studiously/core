//go:generate goagen bootstrap -d github.com/studiously/core/design

package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/logrus"
	"github.com/goadesign/goa/middleware"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"github.com/studiously/core/app"
	"github.com/studiously/core/ddl"
)

const DatabaseKey = "database"

func main() {
	// Create service
	service := goa.New("studiously")
	service.WithLogger(goalogrus.New(logrus.StandardLogger()))

	var driver = os.Getenv("DATABASE_DRIVER")
	var config = os.Getenv("DATABASE_CONFIG")

	db, err := sql.Open(driver, config)
	if err != nil {
		logrus.Fatalln("database connection failed", err)
	}
	if err := pingDatabase(db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("database ping attempts failed")
	}
	if err := setupDatabase(driver, db); err != nil {
		logrus.Errorln(err)
		logrus.Fatalln("migration failed")
	}
	service.Use(func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			ctx = context.WithValue(ctx, DatabaseKey, db)
			return h(ctx, rw, req)
		}
	})

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "class" controller
	c := NewClassController(service)
	app.MountClassController(service, c)
	// Mount "question" controller
	c2 := NewQuestionController(service)
	app.MountQuestionController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// helper function to setup the databsae by performing
// automated database migration steps.
func setupDatabase(driver string, db *sql.DB) error {
	var migrations = &migrate.AssetMigrationSource{
		Asset:    ddl.Asset,
		AssetDir: ddl.AssetDir,
		Dir:      driver,
	}
	_, err := migrate.Exec(db, driver, migrations, migrate.Up)
	return err
}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 30; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		logrus.Infof("database ping failed. retry in 1s")
		time.Sleep(time.Second)
	}
	return
}
