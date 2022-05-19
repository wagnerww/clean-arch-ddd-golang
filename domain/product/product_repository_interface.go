package product

import shared "github.com/wagnerww/go-clean-arch-implement/domain/shared/repository"

type ProductRepositoryInterface interface {
	shared.RepositoryInterface[Product]
}
