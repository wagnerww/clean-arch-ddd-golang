package customer

import shared "github.com/wagnerww/go-clean-arch-implement/domain/shared/repository"

type CustomerRepositoryInterface interface {
	shared.RepositoryInterface[Customer]
}
