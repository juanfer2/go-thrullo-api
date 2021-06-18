package cmd

import (
	"os"

	"github.com/gobuffalo/pop/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var MigrateCmd = cobra.Command{
	Use:  "migrate",
	Long: "Migrate database strucutures. This will create new tables and add missing columns and indexes.",
	Run:  migrate,
}

func migrate(cmd *cobra.Command, args []string) {
	logrus.Info("Start connection database")
	//db := config.Conn()
	logrus.Info("Start migrations")
	// Migrate the schema
	//db.AutoMigrate(&models.User{})

	pop.Debug = true

	deets := &pop.ConnectionDetails{
		Dialect: "postgres",
		URL:     "postgres://postgres:1234@127.0.0.1:5432/go-thrullo?sslmode=disable",
	}

	db, err := pop.NewConnection(deets)

	if err != nil {
		logrus.Fatalf("%+v", err)
	}

	//defer db.Close()

	if err := db.Open(); err != nil {
		logrus.Fatalf("%+v", err)
	}

	mig, err := pop.NewFileMigrator("migrations", db)

	if err != nil {
		logrus.Fatalf("%+v", err)
	}

	logrus.Infof("before status")
	err = mig.Status(os.Stdout)

	// turn off schema dump
	mig.SchemaPath = ""

	err = mig.Up()
	if err != nil {
		logrus.Fatalf("%+v", err)
	}

	logrus.Infof("after status")
	err = mig.Status(os.Stdout)
	if err != nil {
		logrus.Fatalf("%+v", err)
	}

	logrus.Info("mig")
	logrus.Info(mig)

}
