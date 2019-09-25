package main

import (
	"context"
	"errors"
)

func doJobCtx(ctx context.Context, data int, okChan chan bool, errChan chan error) {
	select {
	case <-ctx.Done():
		return
	default:
		if data == 1 {
			errChan <- errors.New("something went wrong")
		}
		okChan <- true
	}
}
