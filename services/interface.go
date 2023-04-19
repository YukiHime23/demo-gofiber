package services

import (
	"context"
)

type SampleServicer interface {
	SampleFunc(ctx context.Context) []map[string]interface{}
}
