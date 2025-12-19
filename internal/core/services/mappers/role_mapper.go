package mappers

import (
	"github.com/amirbakhtiari/orga/internal/core/domain"
	dto "github.com/amirbakhtiari/orga/internal/core/services/dto/role"
)

func RoleToResponse(role *domain.Role) *dto.RoleResponseDTO {
	return &dto.RoleResponseDTO{
		ID:          role.ID,
		Name:        role.Name,
		Code:        role.Code,
		Description: role.Description,
	}
}
