// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package comprehendcounter

import (
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
)

// Comprehend counts the API operations made to Amazon Comprehend.
type Comprehend struct {
	svc    comprehendiface.ComprehendAPI
	counts sync.Map
}

var _ comprehendiface.ComprehendAPI = &Comprehend{}

// New creates a new instance of the Comprehend counter.
func New(svc comprehendiface.ComprehendAPI) *Comprehend {
	return &Comprehend{svc: svc}
}

// Counters returns a snapshot of current counter values for each API operation.
func (c *Comprehend) Counters() map[string]int64 {
	counters := map[string]int64{}
	c.counts.Range(func(key, value interface{}) bool {
		counters[key.(string)] = value.(*counter).count()
		return true
	})
	return counters
}

// incViaRequestOption returns a request.Option that increments the given op.
func (c *Comprehend) incViaRequestOption(op string) request.Option {
	return func(*request.Request) {
		c.inc(op)
	}
}

// inc increments the counter for the operation by 1.
func (c *Comprehend) inc(op string) {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	cnt.(*counter).inc()
}

// count returns the current counter value for the operation.
func (c *Comprehend) count(op string) int64 {
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
