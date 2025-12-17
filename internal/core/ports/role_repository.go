package ports

import "github.com/amirbakhtiari/orga/internal/core/domain"

type RoleRepository interface {
	ListAll() ([]*domain.Role, error)
	GetByID(id uint) (*domain.Role, error)
	GetByName(name string) (*domain.Role, error)
	Create(role *domain.Role) error
	Update(role *domain.Role) error
	Delete(id uint) error
}
