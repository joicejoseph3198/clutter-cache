package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"joicejoseph.dev/clutter-cache/stash-service/internal/data"
	"joicejoseph.dev/clutter-cache/stash-service/internal/repository"
	"joicejoseph.dev/clutter-cache/stash-service/pkg/model"
)

type EntryService interface{
	GetEntries() ([]*model.Entry, error)
	GetEntryById(id int) (*model.Entry, error)
	CreateAnEntry(newEntry *data.EntryRequest) error
	UpdateEntry(entry *model.Entry) error
	DeleteEntry(id int) error
} 

type entryService struct {
	entryRepository repository.EntryRepository
	validate *validator.Validate
}

func NewEntryService(entryRepository repository.EntryRepository, validate *validator.Validate) EntryService{
	return &entryService{
		entryRepository: entryRepository,
		validate: validate,
	}
}

func (s *entryService) GetEntries() ([]*model.Entry, error) {
    return s.entryRepository.FindAll()
}

func (s *entryService) GetEntryById(id int) (*model.Entry, error) {
    return s.entryRepository.FindById(id)
}

func (s *entryService) CreateAnEntry(entry *data.EntryRequest) error {
	err := s.validate.Struct(entry)
	if err != nil{
	   return errors.New("bad request")
	}
    return s.entryRepository.CreateAnEntry(entry)
}

func (s *entryService) UpdateEntry(entry *model.Entry) error {
    return s.entryRepository.UpdateEntry(entry)
}

func (s *entryService) DeleteEntry(id int) error {
    return s.entryRepository.DeleteEntry(id)
}