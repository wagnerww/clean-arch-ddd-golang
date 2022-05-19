package cacheredis

import "github.com/wagnerww/go-clean-arch-implement/domain/product"

func DomainToDto(product product.Product) ProductCacheRedisDto {
	return ProductCacheRedisDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func DtoToDomain(productDto ProductCacheRedisDto) product.Product {
	productDomain, _ := product.NewProduct(productDto.ID, productDto.Name, productDto.Price)
	return productDomain
}
