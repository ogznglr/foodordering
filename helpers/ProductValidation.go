package helpers

import (
	"food/models"
	"strconv"
)

func ProductValidation(issuer string) (models.Product, error) {
	productid, err := strconv.Atoi(issuer)
	if err != nil {
		return models.Product{}, err
	}

	product := models.Product{}.First(productid)
	if product.ID == 0 {
		return models.Product{}, err
	}
	return product, nil
}
