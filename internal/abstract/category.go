package abstract

import (
	"context"

	"github.com/quangchien0212/ecommerce-app/internal/models"
)

type Category interface {
	AddCategory(ctx context.Context, c *models.Category) (*models.Category, error)
}