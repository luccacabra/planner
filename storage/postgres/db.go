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

//func (d *DB) Get(holder interface{}, query string, args ...interface{}) ([]interface{}, error) {
//	return d.dbMap.Select(holder, query, args)
//}
//
//func (d *DB) GetOne(holder interface{}, query string, args ...interface{}) error {
//	if err := d.dbMap.SelectOne(holder, query, args...); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (d *DB) Insert(holders ...interface{}) error {
//	if err := d.dbMap.Insert(holders...); err != nil {
//		return err
//	}
//	return nil
//}
//
//func (d *DB) Update(holders ...interface{}) (int64, error) {
//	count, err := d.dbMap.Update(holders...)
//	if err != nil {
//		return 0, err
//	}
//	return count, nil
//}
