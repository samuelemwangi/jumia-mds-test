package product

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/samuelemwangi/jumia-mds-test/services/products/application/errorhelper"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence"
	"github.com/samuelemwangi/jumia-mds-test/services/products/persistence/repositories"
)

type ProductService interface {
	GetProductBySKU(*ProductRequestDTO) (*ProductResponseDTO, *errorhelper.ErrorResponseDTO)
}

type productService struct {
	productRepo  repositories.ProductRepository
	errorService errorhelper.ErrorService
}

func NewProductService(repos *persistence.Repositories) *productService {
	return &productService{
		productRepo:  repos.ProductRepo,
		errorService: errorhelper.NewErrorService(),
	}
}

func (service *productService) GetProductBySKU(request *ProductRequestDTO) (*ProductResponseDTO, *errorhelper.ErrorResponseDTO) {

	// validate request
	validationErrors := request.validateRequest()

	if len(validationErrors) > 0 {
		return nil, service.errorService.GetValidationError(http.StatusBadRequest, validationErrors)
	}

	// get item
	product := request.toEntity()
	dbError := service.productRepo.GetProductBySKU(product)

	// handle errors
	if dbError != nil {
		status := http.StatusInternalServerError
		if dbError.Error() == gorm.ErrRecordNotFound.Error() {
			status = http.StatusNotFound
		}
		return nil, service.errorService.GetGeneralError(status, dbError)
	}

	// prepare response
	var productResponse ProductResponseDTO
	productResponse.toResponseDTO(product)

	return &productResponse, nil

}
