package main

import (
	"context"
	"fmt"
	"time"

	"github.com/devlights/gomy/ctxs"
	"github.com/devlights/gomy/iter"
)

func main() {
	var (
		rootCtx          = context.Background()
		mainCtx, mainCxl = context.WithCancel(rootCtx)
		procCtx, procCxl = context.WithTimeout(mainCtx, 1*time.Second)
	)
	defer mainCxl()
	defer procCxl()

	<-exec(procCtx).Done()
}

func exec(pCtx context.Context) context.Context {
	var (
		tasks = make([]context.Context, 0, 5)
	)

	for i := range iter.Range(5) {
		tasks = append(tasks, func(pCtx context.Context, i int) context.Context {
			ctx, cancel := context.WithCancel(pCtx)
			go func() {
				defer cancel()

				select {
				case <-ctx.Done():
					break
				default:
					fmt.Printf("[%d] hello-go\n", i+1)
				}
			}()
			return ctx
		}(pCtx, i))
	}

	return ctxs.WhenAll(pCtx, tasks...)
}
