package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ultrabor/warehouse-api/internal/domain"
)

type TestProductRepo struct {
	mock.Mock
}

func (t *TestProductRepo) Create(ctx context.Context, p *domain.Product) (int64, error) {
	args := t.Called(ctx, p)

	return args.Get(0).(int64), args.Error(1)
}

func (t *TestProductRepo) GetByID(ctx context.Context, id int64) (*domain.Product, error) {
	args := t.Called(ctx, id)
	return args.Get(0).(*domain.Product), args.Error(1)
}
func (m *TestProductRepo) Update(ctx context.Context, p *domain.Product) error {
	args := m.Called(ctx, p)

	return args.Error(0)
}

func (m *TestProductRepo) Delete(ctx context.Context, id int64) error {

	args := m.Called(ctx, id)

	return args.Error(0)
}

func (m *TestProductRepo) GetAll(ctx context.Context) ([]domain.Product, error) { return nil, nil }

func TestProductUseCase_Create(t *testing.T) {
	repo := new(TestProductRepo)
	uc := NewProductUseCase(repo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		product := &domain.Product{
			Name:     "Test Product",
			Price:    100.0,
			Quantity: 10,
		}

		repo.On("Create", ctx, product).Return(int64(1), nil).Once()

		id, err := uc.Create(ctx, product)

		assert.NoError(t, err)
		assert.Equal(t, int64(1), id)
		repo.AssertExpectations(t)
	})

	t.Run("invalid price error", func(t *testing.T) {
		product := &domain.Product{
			Name:     "Cheap Product",
			Price:    -10.0,
			Quantity: 10,
		}

		id, err := uc.Create(ctx, product)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, domain.ErrInvalidPrice))
		assert.Equal(t, int64(0), id)

		repo.AssertNotCalled(t, "Create", ctx, product)
	})
}

func TestProductUseCase_GetByID(t *testing.T) {
	repo := new(TestProductRepo)
	uc := NewProductUseCase(repo)
	ctx := context.Background()

	t.Run("not found", func(t *testing.T) {
		repo.On("GetByID", ctx, int64(99)).Return((*domain.Product)(nil), domain.ErrProductNotFound).Once()

		product, err := uc.GetByID(ctx, 99)

		assert.Error(t, err)
		assert.Nil(t, product)
		assert.True(t, errors.Is(err, domain.ErrProductNotFound))
	})
}

func TestProductUseCase_GetAll(t *testing.T) {
	repo := new(TestProductRepo)
	uc := NewProductUseCase(repo)
	ctx := context.Background()

	t.Run("empty", func(t *testing.T) {

		product, err := uc.GetAll(ctx)

		assert.NoError(t, err)
		assert.Nil(t, product)
	})
}

func TestProductUseCase_Update(t *testing.T) {
	repo := new(TestProductRepo)
	uc := NewProductUseCase(repo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		product := &domain.Product{
			ID:       1,
			Name:     "Updated Name",
			Price:    150.0,
			Quantity: 5,
		}

		repo.On("GetByID", ctx, product.ID).Return(product, nil).Once()

		repo.On("Update", ctx, product).Return(nil).Once()

		err := uc.Update(ctx, product)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("invalid price", func(t *testing.T) {
		product := &domain.Product{
			ID:    1,
			Price: -50.0,
		}

		err := uc.Update(ctx, product)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, domain.ErrInvalidPrice))
	})

	t.Run("not found", func(t *testing.T) {
		product := &domain.Product{ID: 99, Name: "Test", Price: 10.0}

		repo.On("GetByID", ctx, product.ID).Return((*domain.Product)(nil), domain.ErrProductNotFound).Once()

		err := uc.Update(ctx, product)

		assert.Error(t, err)
		assert.True(t, errors.Is(err, domain.ErrProductNotFound))
	})
}

func TestProductUseCase_Delete(t *testing.T) {
	repo := new(TestProductRepo)
	uc := NewProductUseCase(repo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		productID := int64(1)

		repo.On("Delete", ctx, productID).Return(nil).Once()

		err := uc.Delete(ctx, productID)

		assert.NoError(t, err)
		repo.AssertExpectations(t)
	})
}
