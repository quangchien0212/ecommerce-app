package database

import (
	"context"
	"errors"

	"github.com/quangchien0212/ecommerce-app/internal/generic/common_errors"
	"github.com/quangchien0212/ecommerce-app/internal/models"
	"gorm.io/gorm"
)

func (c Client) AddCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	result := c.DB.WithContext(ctx).Create(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &common_errors.ConflictError{}
		}
		if errors.Is(result.Error, gorm.ErrForeignKeyViolated) {
			return nil, &common_errors.ViolationError{}
		}
		return nil, result.Error
	}
	return category, nil
}
