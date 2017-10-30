// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearchdomaincounter

import (
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
)

// CloudSearchDomain counts the API operations made to Amazon CloudSearch Domain.
type CloudSearchDomain struct {
	svc    cloudsearchdomainiface.CloudSearchDomainAPI
	counts sync.Map
}

var _ cloudsearchdomainiface.CloudSearchDomainAPI = &CloudSearchDomain{}

// New creates a new instance of the CloudSearchDomain counter.
func New(svc cloudsearchdomainiface.CloudSearchDomainAPI) *CloudSearchDomain {
	return &CloudSearchDomain{svc: svc}
}

// Counters returns a snapshot of current counter values for each API operation.
func (c *CloudSearchDomain) Counters() map[string]int64 {
	counters := map[string]int64{}
	c.counts.Range(func(key, value interface{}) bool {
		counters[key.(string)] = value.(*counter).count()
		return true
	})
	return counters
}

// incViaRequestOption returns a request.Option that increments the given op.
func (c *CloudSearchDomain) incViaRequestOption(op string) request.Option {
	return func(*request.Request) {
		c.inc(op)
	}
}

// inc increments the counter for the operation by 1.
func (c *CloudSearchDomain) inc(op string) {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	cnt.(*counter).inc()
}

// count returns the current counter value for the operation.
func (c *CloudSearchDomain) count(op string) int64 {
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
