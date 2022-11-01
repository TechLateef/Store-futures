package services

import (
	"github.com/adewumi0550/pro_work/m/v2/dto"
	"github.com/adewumi0550/pro_work/m/v2/repository"
	"github.com/mashingan/smapping"
	"log"
)

type StoreService interface {
	CreateStore(s dto.StoreCreateDTO) entities.store
	UpdateStore(s dto.StoreUpdateDTO) entities.Store
	Enable(s entities.store)
	Disable(s entities.store)
}

type storeService struct {
	storeRepository repository.StoreRepository
}

func NewStoreService(storeRepo repository.StoreRepository) StoreService {
	return &storeService{
		storeRepository: storeRepo,
	}
}

func (service *storeService) CreateStore(s dto.StoreCreateDTO) entities.store {
	store := entities.store{}
	err := smapping.FillStruct(&store, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.storeRepository.CreateStore(store)
	return res
}
func (service *storeService) Update(s dto.StoreUpdateDTO) entities.store {
	store := entities.store{}
	err := smapping.FillStruct(&store, smapping.MapFields(&s))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	res := service.storeRepository.UpdateStore(store)
	return res
}
