package postgres

import (
	"luccacabra/planner/storage"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type postgresStorage struct {
	conn *gorm.DB
}

func DBInit() (*storage.Storage, error) {
	log.Info("Initializing data store connection")
	db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	s := &storage.Storage{}
	s.Tasks = (TaskStorage)(db)
	return s, nil
}
