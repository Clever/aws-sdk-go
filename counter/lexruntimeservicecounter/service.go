// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lexruntimeservicecounter

import (
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lexruntimeservice/lexruntimeserviceiface"
)

// LexRuntimeService counts the API operations made to Amazon Lex Runtime Service.
type LexRuntimeService struct {
	svc    lexruntimeserviceiface.LexRuntimeServiceAPI
	counts sync.Map
}

var _ lexruntimeserviceiface.LexRuntimeServiceAPI = &LexRuntimeService{}

// New creates a new instance of the LexRuntimeService counter.
func New(svc lexruntimeserviceiface.LexRuntimeServiceAPI) *LexRuntimeService {
	return &LexRuntimeService{svc: svc}
}

// Counters returns a snapshot of current counter values for each API operation.
func (c *LexRuntimeService) Counters() map[string]int64 {
	counters := map[string]int64{}
	c.counts.Range(func(key, value interface{}) bool {
		counters[key.(string)] = value.(*counter).count()
		return true
	})
	return counters
}

// incViaRequestOption returns a request.Option that increments the given op.
func (c *LexRuntimeService) incViaRequestOption(op string) request.Option {
	return func(*request.Request) {
		c.inc(op)
	}
}

// inc increments the counter for the operation by 1.
func (c *LexRuntimeService) inc(op string) {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	cnt.(*counter).inc()
}

// count returns the current counter value for the operation.
func (c *LexRuntimeService) count(op string) int64 {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	return cnt.(*counter).count()
}

// counter is a threadsafe cumulative counter.
type counter struct {
	c int64
}

// count returns the current count.
func (c *counter) count() int64 {
	return atomic.LoadInt64(&c.c)
}

// inc increments the counter by one.
func (c *counter) inc() {
	atomic.AddInt64(&c.c, 1)
}
