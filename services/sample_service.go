package services

import (
	"context"
	"demo-gofiber/repositories"
)

func (s *AppService) SampleFunc(ctx context.Context) []map[string]interface{} {
	return repositories.SampleFunction(ctx, s.db)
}
