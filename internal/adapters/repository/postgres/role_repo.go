package postgres

import (
	"context"
	"errors"
	"github.com/amirbakhtiari/orga/internal/adapters/repository/postgres/mappers"
	"github.com/amirbakhtiari/orga/internal/adapters/repository/postgres/models"
	"github.com/amirbakhtiari/orga/internal/core/domain"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{db: db}
}

func (r *RoleRepository) Create(
	ctx context.Context,
	role *domain.Role,
) error {
	model := mappers.RoleDomainToModel(role)

	if err := r.db.WithContext(ctx).Create(model).Error; err != nil {
		return err
	}
	role.ID = model.ID
	role.CreatedAt = model.CreatedAt
	role.UpdatedAt = model.UpdatedAt
	return nil
}

func (r *RoleRepository) GetByCode(
	ctx context.Context,
	code string,
) (*domain.Role, error) {
	var model models.Role

	err := r.db.WithContext(ctx).
		Where("code=?", code).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return mappers.RoleModelToDomain(&model), nil
}

func (r *RoleRepository) List(ctx context.Context) ([]*domain.Role, error) {
	var roleModels []models.Role
	if err := r.db.WithContext(ctx).
		Order("id ASC").
		Find(&roleModels).Error; err != nil {
		return nil, err
	}

	roles := make([]*domain.Role, 0, len(roleModels))

	for _, m := range roleModels {
		role := mappers.RoleModelToDomain(&m)
		roles = append(roles, role)
	}
	return roles, nil
}
func (r *RoleRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).
		Delete(&models.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *RoleRepository) Update(ctx context.Context, role *domain.Role) error {
	model := mappers.RoleDomainToModel(role)

	if err := r.db.WithContext(ctx).
		Model(&models.Role{}).
		Where("id=?", role.ID).
		Updates(model).Error; err != nil {
		return err
	}
	return nil
}
func (r *RoleRepository) GetByID(ctx context.Context, id uint) (*domain.Role, error) {
	var role domain.Role
	role := r.db.WithContext(ctx).Find(&role, id)
}
