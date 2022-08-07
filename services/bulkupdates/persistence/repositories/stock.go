package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/bulkupdates/domain"
)

type StockRepository interface {
	GetStockByProductAndCountry(*domain.Stock) error
	SaveStock(*domain.Stock) error
	UpdateStock(*domain.Stock) error
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) *stockRepository {
	return &stockRepository{
		db: db,
	}
}

func (repo *stockRepository) GetStockByProductAndCountry(stock *domain.Stock) error {
	result := repo.db.Where("country_id = ? AND product_id = ?", stock.CountryID, stock.ProductID).Find(stock)
	return result.Error
}

func (repo *stockRepository) SaveStock(stock *domain.Stock) error {
	result := repo.db.Create(stock)
	return result.Error
}

func (repo *stockRepository) UpdateStock(stock *domain.Stock) error {
	result := repo.db.Model(&domain.Stock{}).Where("country_id = ? AND product_id = ?", stock.CountryID, stock.ProductID).Updates(stock)
	return result.Error
}
