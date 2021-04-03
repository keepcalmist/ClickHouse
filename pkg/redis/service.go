package redis

import "context"

type Repo interface {
	GetDigit(ctx context.Context, key int) (int, error)
	SetDigit(ctx context.Context, key int, value int) error
	SetMax(ctx context.Context, value int) error
	GetMax(ctx context.Context) (int, error)
	GetSliceDigits(ctx context.Context, from, to int) ([]int, error)
	SetArray(ctx context.Context, from, to int32, digSlice []int32) error
}
