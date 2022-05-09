package closer

import (
	"container/list"
	"context"
	"sync"

	"github.com/hashicorp/go-multierror"
)

// Closer defines close signature.
type Closer func(context.Context) error

// Closers manages closers in correct order.
type Closers struct {
	list list.List
	mu   sync.Mutex
}

// Add a closer to the list.
func (c *Closers) Add(closer Closer) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.list.PushFront(closer)
}

// Close the closers in reversed order.
func (c *Closers) Close(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Close in reversed order, just like defer does
	var result error
	for el := c.list.Front(); el != nil; el = el.Next() {
		closer, _ := el.Value.(Closer)
		if err := closer(ctx); err != nil {
			result = multierror.Append(result, err)
		}
	}

	c.list.Init() // clear the list

	return result
}
