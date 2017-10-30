// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspacescounter

import (
	"sync"
	"sync/atomic"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"
)

// WorkSpaces counts the API operations made to Amazon WorkSpaces.
type WorkSpaces struct {
	svc    workspacesiface.WorkSpacesAPI
	counts sync.Map
}

var _ workspacesiface.WorkSpacesAPI = &WorkSpaces{}

// New creates a new instance of the WorkSpaces counter.
func New(svc workspacesiface.WorkSpacesAPI) *WorkSpaces {
	return &WorkSpaces{svc: svc}
}

// Counters returns a snapshot of current counter values for each API operation.
func (c *WorkSpaces) Counters() map[string]int64 {
	counters := map[string]int64{}
	c.counts.Range(func(key, value interface{}) bool {
		counters[key.(string)] = value.(*counter).count()
		return true
	})
	return counters
}

// incViaRequestOption returns a request.Option that increments the given op.
func (c *WorkSpaces) incViaRequestOption(op string) request.Option {
	return func(*request.Request) {
		c.inc(op)
	}
}

// inc increments the counter for the operation by 1.
func (c *WorkSpaces) inc(op string) {
	cnt, _ := c.counts.LoadOrStore(op, &counter{})
	cnt.(*counter).inc()
}

// count returns the current counter value for the operation.
func (c *WorkSpaces) count(op string) int64 {
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
