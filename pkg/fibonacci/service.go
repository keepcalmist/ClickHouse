package fibonacci

import (
	"context"
	"github.com/keepcalmist/grpcFibonacci/pkg/redis"
	"github.com/keepcalmist/grpcFibonacci/pkg/usefullFunctions"
)

type FiboService interface {
	CalculateDigits(ctx context.Context, from, to int32) ([]int32, error)
}

type fibo struct {
	rConn redis.Repo
}

func New(repo redis.Repo) FiboService {
	return &fibo{rConn: repo}
}

func (s *fibo) CalculateDigits(ctx context.Context, from, to int32) ([]int32, error) {
	maxKeyInStorage, err := s.rConn.GetMax(ctx)
	if err != nil {
		return nil, err
	}
	if maxKeyInStorage > int(from) && maxKeyInStorage > int(to) {
		fiboSlice, err := s.rConn.GetSliceDigits(ctx, int(from), int(to))
		if err != nil {
			return nil, err
		}
		reSlice := make([]int32, 0, len(fiboSlice))
		for _, value := range fiboSlice {
			reSlice = append(reSlice, int32(value))
		}
		return reSlice, nil
	}

	if maxKeyInStorage > int(from) && maxKeyInStorage < int(to) {
		fiboSlice, err := s.rConn.GetSliceDigits(ctx, int(from), maxKeyInStorage)
		if err != nil {
			return nil, err
		}
		calculatedDigits := usefullFunctions.Calculate(fiboSlice[len(fiboSlice)-1], fiboSlice[len(fiboSlice)-2], to-int32(maxKeyInStorage))
		err = s.rConn.SetArray(ctx, int32(maxKeyInStorage+1), to, calculatedDigits)
		if err != nil {
			return nil, err
		}
		reSlice := make([]int32, 0, len(fiboSlice))
		for _, value := range fiboSlice {
			reSlice = append(reSlice, int32(value))
		}
		return append(reSlice, calculatedDigits...), nil
	}

	if maxKeyInStorage < int(from) && from >= 2 {
		fiboSlice, err := s.rConn.GetSliceDigits(ctx, maxKeyInStorage-1, maxKeyInStorage)
		if err != nil {
			return nil, err
		}
		calculatedDigits := usefullFunctions.Calculate(fiboSlice[len(fiboSlice)-1],
			fiboSlice[len(fiboSlice)-2], to-int32(maxKeyInStorage))
		reSlice := make([]int32, 0, len(fiboSlice))
		for _, value := range fiboSlice {
			reSlice = append(reSlice, int32(value))
		}
		err = s.rConn.SetArray(ctx, int32(maxKeyInStorage+1), to, calculatedDigits)
		if err != nil {
			return nil, err
		}
		return calculatedDigits[from-int32(maxKeyInStorage)-1:], nil
	}

	return nil, err
}
