package mappers

import (
	"github.com/amirbakhtiari/orga/internal/adapters/repository/postgres/models"
	"github.com/amirbakhtiari/orga/internal/core/domain"
)

func RoleModelToDomain(m *models.Role) *domain.Role {
	return &domain.Role{
		ID:          m.ID,
		Name:        m.Name,
		Code:        m.Code,
		Description: *m.Description,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func RoleDomainToModel(d *domain.Role) *models.Role {
	return &models.Role{
		ID:          d.ID,
		Name:        d.Name,
		Code:        d.Code,
		Description: &d.Description,
	}
}
