// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecrcounter

import (
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
)

// ECR counts the API operations made to Amazon EC2 Container Registry.
type ECR struct {
	svc    ecriface.ECRAPI
	counts sync.Map
}

var _ ecriface.ECRAPI = &ECR{}

// New creates a new instance of the ECR counter.
func New(svc ecriface.ECRAPI) *ECR {
	return &ECR{svc: svc}
}

// Counters returns a snapshot of current counter values for each API operation.
func (c *ECR) Counters() map[string]int64 {
	counters := map[string]int64{}
	c.counts.Range(func(key, value interface{}) bool {
		counters[key.(string)] = value.(*counter).count()
		return true
	})
	return counters
}

// incViaRequestOption returns a request.Option that increments the given op.
func (c *ECR) incViaRequestOption(op string) request.Option {
	return func(*request.Request) {
		c.inc(op)
	}
}

// inc increments the counter for the operation by 1.
func (c *ECR) inc(op string) {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	cnt.(*counter).inc()
}

// count returns the current counter value for the operation.
func (c *ECR) count(op string) int64 {
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
