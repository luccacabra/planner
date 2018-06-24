package storage

import "luccacabra/planner/storage/models"

type TaskStorage interface {
	GetOne(t models.Task) error
}