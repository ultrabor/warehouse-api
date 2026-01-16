package usecase

import (
	"context"

	"github.com/ultrabor/warehouse-api/internal/domain"
)

type ProductUseCase struct {
	repo domain.ProductRepository
}

func NewProductUseCase(r domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: r}
}

func (uc *ProductUseCase) Create(ctx context.Context, p *domain.Product) (int64, error) {
	if p.Price < 0 {
		return 0, domain.ErrInvalidPrice
	}

	if p.Quantity < 0 {
		return 0, domain.ErrInvalidQuantity
	}

	return uc.repo.Create(ctx, p)
}

func (uc *ProductUseCase) GetByID(ctx context.Context, id int64) (*domain.Product, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *ProductUseCase) Update(ctx context.Context, p *domain.Product) error {
	if p.Price < 0 {
		return domain.ErrInvalidPrice
	}

	if p.Quantity < 0 {
		return domain.ErrInvalidQuantity
	}

	_, err := uc.repo.GetByID(ctx, p.ID)

	if err != nil {
		return err
	}

	return uc.repo.Update(ctx, p)
}

func (uc *ProductUseCase) Delete(ctx context.Context, id int64) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *ProductUseCase) GetAll(ctx context.Context) ([]domain.Product, error) {
	return uc.repo.GetAll(ctx)
}
