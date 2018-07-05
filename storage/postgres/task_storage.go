package postgres

import (
	"luccacabra/planner/storage"
	"luccacabra/planner/storage/models"

	"github.com/jinzhu/gorm"
)

type taskStorage postgresStorage

var _ storage.TaskStorage = (*taskStorage)(nil)

func TaskStorage(conn *gorm.DB) storage.TaskStorage {
	return &taskStorage{
		conn: conn,
	}
}

func (t *taskStorage) GetOne(task models.Task) error {
	return nil
}
