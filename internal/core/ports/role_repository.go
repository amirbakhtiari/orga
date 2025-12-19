package ports

import (
	"context"
	"github.com/amirbakhtiari/orga/internal/core/domain"
)

type RoleRepository interface {
	Create(ctx context.Context, role *domain.Role) error
	Update(ctx context.Context, role *domain.Role) error
	GetByID(ctx context.Context, id uint) (*domain.Role, error)
	GetByCode(ctx context.Context, code string) (*domain.Role, error)
	List(ctx context.Context) ([]*domain.Role, error)
	Delete(ctx context.Context, id uint) error
}
