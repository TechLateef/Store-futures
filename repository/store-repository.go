package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BookRepository is a contract between entities and database
type StoreRepository interface {
	CreateStore(s entities.store) entities.store
	UpdateStore(s entities.store) entities.store
}

type storeconnection struct {
	connection *gorm.DB
}

// NewStoreRepository create an instance of StoreRepository
func NewStoreRepository(dbConn *gorm.DB) StoreRepository {
	return &storeconnection{
		connection: dbConn,
	}
}

func (db *storeconnection) CreateStore(s entities.store) entities.store {
	db.connection.Save(&s)
	db.connection.Preload("User").Find(&s)
	return s
}

func (db *storeconnection) UpdateStore(s entities.store) entities.store {
	db.connection.Save(&s)
	db.connection.Preload("User").Find(&s)
	return s
}
