package product

import (
	"fmt"

	"github.com/pradiptarana/product/internal/cache"
	model "github.com/pradiptarana/product/model/product"
	"github.com/pradiptarana/product/repository"
)

type ProductUC struct {
	repository.ProductRepository
	cache.Cache[int, []byte]
}

func NewProductUC(repo repository.ProductRepository, c cache.Cache[int, []byte]) *ProductUC {
	return &ProductUC{repo, c}
}

func (uc *ProductUC) GetProducts() ([]*model.Product, error) {
	data, err := uc.ProductRepository.GetLatestProducts()
	if err != nil {
		fmt.Println(err)
		return []*model.Product{}, err
	}
	return data, nil
}

func (uc *ProductUC) GetProduct(productId int) (data *model.Product, err error) {
	// value, found := uc.Cache.Get(productId)
	// if !found {
	// 	data, err = uc.ProductRepository.GetProduct(productId)
	// 	if err != nil {
	// 		return &model.Product{}, err
	// 	}
	// 	marhalData, err := json.Marshal(data)
	// 	if err != nil {
	// 		return &model.Product{}, err
	// 	}
	// 	uc.Cache.Set(productId, marhalData)
	// } else {
	// 	_ = json.Unmarshal(value, &data)
	// }
	return data, nil
}
