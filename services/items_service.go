package services

var (
	ItemsService itemsService = itemsService{}
)

type itemServiceInterface interface {
	GetItem()
	SaveItem()
}

type itemsService struct {
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) SaveItem() {

}
