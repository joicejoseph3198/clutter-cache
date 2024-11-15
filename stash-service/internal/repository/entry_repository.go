package repository

import (
	"gorm.io/gorm"
	"joicejoseph.dev/clutter-cache/stash-service/internal/data"
	"joicejoseph.dev/clutter-cache/stash-service/pkg/model"
)

type EntryRepository interface {
	FindAll() ([]*model.Entry, error)
	FindById(id int) (*model.Entry, error)
	CreateAnEntry(newEntry *data.EntryRequest) error
	UpdateEntry(entry *model.Entry) error
	DeleteEntry(id int) error
}

// struct that implements EntryRepository interface
type entryRepository struct {
	db *gorm.DB
}

func NewEntryRepository(dbConn *gorm.DB) EntryRepository {
	//any struct that implements all the methods defined by an interface is considered to satisfy that interface,
	// allowing it to be referenced by that interface type.
	return &entryRepository{db: dbConn}
}

func (r *entryRepository) FindAll() ([]*model.Entry, error) {
	var entries []*model.Entry
	if err := r.db.Find(&entries).Error ; err != nil{
		return nil, err
	}
	return entries, nil
}

func (r *entryRepository) FindById(id int) (*model.Entry, error) {
	// Query the database and return the user
	return nil, nil
}

func (r *entryRepository) CreateAnEntry(entry *data.EntryRequest) error {
	// Insert the user into the database
	r.db.Create(entry)
	return nil
}

func (r *entryRepository) UpdateEntry(user *model.Entry) error {
	// Update the user in the database
	return nil
}

func (r *entryRepository) DeleteEntry(id int) error {
	// Delete the user from the database
	return nil
}


