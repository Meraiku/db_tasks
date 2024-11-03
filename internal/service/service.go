package service

import "context"

type Service interface {
	Test(ctx context.Context) error
}
