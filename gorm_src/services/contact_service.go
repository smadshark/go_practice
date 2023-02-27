package services

import (
	"github.com/google/uuid"
	"go_practice/gorm_src/dtos"
	"go_practice/gorm_src/models"
	"go_practice/gorm_src/repositories"
	"log"
)

type ContactService struct {
	rp *repositories.ContactRepository
}

func NewContactService(repository *repositories.ContactRepository) *ContactService {
	return &ContactService{rp: repository}
}

func (cs *ContactService) CreateContact(contact *models.Contact) dtos.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatalln(err)
	}

	contact.ID = uuidResult.String()

	operationResult := cs.rp.Save(contact)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.Contact)

	return dtos.Response{Success: true, Data: data}
}
