package product

import shared "github.com/wagnerww/go-clean-arch-implement/domain/shared/repository"

type ProductRepositoryCacheInterface interface {
	shared.RepositoryCacheInterface[Product]
}
