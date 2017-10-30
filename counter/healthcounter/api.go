// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package healthcounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/health"
)

// DescribeAffectedEntitiesRequest is a passthrough to the underlying DescribeAffectedEntitiesRequest.
// It will increment the count of requests made to DescribeAffectedEntities.
func (c *Health) DescribeAffectedEntitiesRequest(input *health.DescribeAffectedEntitiesInput) (req *request.Request, output *health.DescribeAffectedEntitiesOutput) {
	c.inc("DescribeAffectedEntities")
	return c.svc.DescribeAffectedEntitiesRequest(input)
}

// DescribeAffectedEntities is a passthrough to the underlying DescribeAffectedEntities method.
// It will increment the count of requests made to DescribeAffectedEntities.
func (c *Health) DescribeAffectedEntities(input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
	c.inc("DescribeAffectedEntities")
	return c.svc.DescribeAffectedEntities(input)
}

// DescribeAffectedEntitiesWithContext is a passthrough to the underlying DescribeAffectedEntitiesWithContext method.
// It will increment the count of requests made to DescribeAffectedEntities.
func (c *Health) DescribeAffectedEntitiesWithContext(ctx aws.Context, input *health.DescribeAffectedEntitiesInput, opts ...request.Option) (*health.DescribeAffectedEntitiesOutput, error) {
	c.inc("DescribeAffectedEntities")
	return c.svc.DescribeAffectedEntitiesWithContext(ctx, input, opts...)
}

// DescribeAffectedEntitiesPages is a passthrough to the underlying DescribeAffectedEntitiesPages method.
// It will increment the count of requests made to DescribeAffectedEntities on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use DescribeAffectedEntitiesPagesWithContext to avoid this.
func (c *Health) DescribeAffectedEntitiesPages(input *health.DescribeAffectedEntitiesInput, fn func(*health.DescribeAffectedEntitiesOutput, bool) bool) error {
	wrappedFn := func(page *health.DescribeAffectedEntitiesOutput, lastPage bool) bool {
		c.inc("DescribeAffectedEntities")
		return fn(page, lastPage)
	}
	return c.DescribeAffectedEntitiesPages(input, wrappedFn)
}

// DescribeAffectedEntitiesPagesWithContext is a passthrough to the underlying DescribeAffectedEntitiesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to DescribeAffectedEntities when applied to the request.
func (c *Health) DescribeAffectedEntitiesPagesWithContext(ctx aws.Context, input *health.DescribeAffectedEntitiesInput, fn func(*health.DescribeAffectedEntitiesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("DescribeAffectedEntities"))
	return c.DescribeAffectedEntitiesPagesWithContext(ctx, input, fn, opts...)
}

// DescribeEntityAggregatesRequest is a passthrough to the underlying DescribeEntityAggregatesRequest.
// It will increment the count of requests made to DescribeEntityAggregates.
func (c *Health) DescribeEntityAggregatesRequest(input *health.DescribeEntityAggregatesInput) (req *request.Request, output *health.DescribeEntityAggregatesOutput) {
	c.inc("DescribeEntityAggregates")
	return c.svc.DescribeEntityAggregatesRequest(input)
}

// DescribeEntityAggregates is a passthrough to the underlying DescribeEntityAggregates method.
// It will increment the count of requests made to DescribeEntityAggregates.
func (c *Health) DescribeEntityAggregates(input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error) {
	c.inc("DescribeEntityAggregates")
	return c.svc.DescribeEntityAggregates(input)
}

// DescribeEntityAggregatesWithContext is a passthrough to the underlying DescribeEntityAggregatesWithContext method.
// It will increment the count of requests made to DescribeEntityAggregates.
func (c *Health) DescribeEntityAggregatesWithContext(ctx aws.Context, input *health.DescribeEntityAggregatesInput, opts ...request.Option) (*health.DescribeEntityAggregatesOutput, error) {
	c.inc("DescribeEntityAggregates")
	return c.svc.DescribeEntityAggregatesWithContext(ctx, input, opts...)
}

// DescribeEventAggregatesRequest is a passthrough to the underlying DescribeEventAggregatesRequest.
// It will increment the count of requests made to DescribeEventAggregates.
func (c *Health) DescribeEventAggregatesRequest(input *health.DescribeEventAggregatesInput) (req *request.Request, output *health.DescribeEventAggregatesOutput) {
	c.inc("DescribeEventAggregates")
	return c.svc.DescribeEventAggregatesRequest(input)
}

// DescribeEventAggregates is a passthrough to the underlying DescribeEventAggregates method.
// It will increment the count of requests made to DescribeEventAggregates.
func (c *Health) DescribeEventAggregates(input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error) {
	c.inc("DescribeEventAggregates")
	return c.svc.DescribeEventAggregates(input)
}

// DescribeEventAggregatesWithContext is a passthrough to the underlying DescribeEventAggregatesWithContext method.
// It will increment the count of requests made to DescribeEventAggregates.
func (c *Health) DescribeEventAggregatesWithContext(ctx aws.Context, input *health.DescribeEventAggregatesInput, opts ...request.Option) (*health.DescribeEventAggregatesOutput, error) {
	c.inc("DescribeEventAggregates")
	return c.svc.DescribeEventAggregatesWithContext(ctx, input, opts...)
}

// DescribeEventAggregatesPages is a passthrough to the underlying DescribeEventAggregatesPages method.
// It will increment the count of requests made to DescribeEventAggregates on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use DescribeEventAggregatesPagesWithContext to avoid this.
func (c *Health) DescribeEventAggregatesPages(input *health.DescribeEventAggregatesInput, fn func(*health.DescribeEventAggregatesOutput, bool) bool) error {
	wrappedFn := func(page *health.DescribeEventAggregatesOutput, lastPage bool) bool {
		c.inc("DescribeEventAggregates")
		return fn(page, lastPage)
	}
	return c.DescribeEventAggregatesPages(input, wrappedFn)
}

// DescribeEventAggregatesPagesWithContext is a passthrough to the underlying DescribeEventAggregatesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to DescribeEventAggregates when applied to the request.
func (c *Health) DescribeEventAggregatesPagesWithContext(ctx aws.Context, input *health.DescribeEventAggregatesInput, fn func(*health.DescribeEventAggregatesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("DescribeEventAggregates"))
	return c.DescribeEventAggregatesPagesWithContext(ctx, input, fn, opts...)
}

// DescribeEventDetailsRequest is a passthrough to the underlying DescribeEventDetailsRequest.
// It will increment the count of requests made to DescribeEventDetails.
func (c *Health) DescribeEventDetailsRequest(input *health.DescribeEventDetailsInput) (req *request.Request, output *health.DescribeEventDetailsOutput) {
	c.inc("DescribeEventDetails")
	return c.svc.DescribeEventDetailsRequest(input)
}

// DescribeEventDetails is a passthrough to the underlying DescribeEventDetails method.
// It will increment the count of requests made to DescribeEventDetails.
func (c *Health) DescribeEventDetails(input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error) {
	c.inc("DescribeEventDetails")
	return c.svc.DescribeEventDetails(input)
}

// DescribeEventDetailsWithContext is a passthrough to the underlying DescribeEventDetailsWithContext method.
// It will increment the count of requests made to DescribeEventDetails.
func (c *Health) DescribeEventDetailsWithContext(ctx aws.Context, input *health.DescribeEventDetailsInput, opts ...request.Option) (*health.DescribeEventDetailsOutput, error) {
	c.inc("DescribeEventDetails")
	return c.svc.DescribeEventDetailsWithContext(ctx, input, opts...)
}

// DescribeEventTypesRequest is a passthrough to the underlying DescribeEventTypesRequest.
// It will increment the count of requests made to DescribeEventTypes.
func (c *Health) DescribeEventTypesRequest(input *health.DescribeEventTypesInput) (req *request.Request, output *health.DescribeEventTypesOutput) {
	c.inc("DescribeEventTypes")
	return c.svc.DescribeEventTypesRequest(input)
}

// DescribeEventTypes is a passthrough to the underlying DescribeEventTypes method.
// It will increment the count of requests made to DescribeEventTypes.
func (c *Health) DescribeEventTypes(input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error) {
	c.inc("DescribeEventTypes")
	return c.svc.DescribeEventTypes(input)
}

// DescribeEventTypesWithContext is a passthrough to the underlying DescribeEventTypesWithContext method.
// It will increment the count of requests made to DescribeEventTypes.
func (c *Health) DescribeEventTypesWithContext(ctx aws.Context, input *health.DescribeEventTypesInput, opts ...request.Option) (*health.DescribeEventTypesOutput, error) {
	c.inc("DescribeEventTypes")
	return c.svc.DescribeEventTypesWithContext(ctx, input, opts...)
}

// DescribeEventTypesPages is a passthrough to the underlying DescribeEventTypesPages method.
// It will increment the count of requests made to DescribeEventTypes on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use DescribeEventTypesPagesWithContext to avoid this.
func (c *Health) DescribeEventTypesPages(input *health.DescribeEventTypesInput, fn func(*health.DescribeEventTypesOutput, bool) bool) error {
	wrappedFn := func(page *health.DescribeEventTypesOutput, lastPage bool) bool {
		c.inc("DescribeEventTypes")
		return fn(page, lastPage)
	}
	return c.DescribeEventTypesPages(input, wrappedFn)
}

// DescribeEventTypesPagesWithContext is a passthrough to the underlying DescribeEventTypesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to DescribeEventTypes when applied to the request.
func (c *Health) DescribeEventTypesPagesWithContext(ctx aws.Context, input *health.DescribeEventTypesInput, fn func(*health.DescribeEventTypesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("DescribeEventTypes"))
	return c.DescribeEventTypesPagesWithContext(ctx, input, fn, opts...)
}

// DescribeEventsRequest is a passthrough to the underlying DescribeEventsRequest.
// It will increment the count of requests made to DescribeEvents.
func (c *Health) DescribeEventsRequest(input *health.DescribeEventsInput) (req *request.Request, output *health.DescribeEventsOutput) {
	c.inc("DescribeEvents")
	return c.svc.DescribeEventsRequest(input)
}

// DescribeEvents is a passthrough to the underlying DescribeEvents method.
// It will increment the count of requests made to DescribeEvents.
func (c *Health) DescribeEvents(input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error) {
	c.inc("DescribeEvents")
	return c.svc.DescribeEvents(input)
}

// DescribeEventsWithContext is a passthrough to the underlying DescribeEventsWithContext method.
// It will increment the count of requests made to DescribeEvents.
func (c *Health) DescribeEventsWithContext(ctx aws.Context, input *health.DescribeEventsInput, opts ...request.Option) (*health.DescribeEventsOutput, error) {
	c.inc("DescribeEvents")
	return c.svc.DescribeEventsWithContext(ctx, input, opts...)
}

// DescribeEventsPages is a passthrough to the underlying DescribeEventsPages method.
// It will increment the count of requests made to DescribeEvents on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use DescribeEventsPagesWithContext to avoid this.
func (c *Health) DescribeEventsPages(input *health.DescribeEventsInput, fn func(*health.DescribeEventsOutput, bool) bool) error {
	wrappedFn := func(page *health.DescribeEventsOutput, lastPage bool) bool {
		c.inc("DescribeEvents")
		return fn(page, lastPage)
	}
	return c.DescribeEventsPages(input, wrappedFn)
}

// DescribeEventsPagesWithContext is a passthrough to the underlying DescribeEventsPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to DescribeEvents when applied to the request.
func (c *Health) DescribeEventsPagesWithContext(ctx aws.Context, input *health.DescribeEventsInput, fn func(*health.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("DescribeEvents"))
	return c.DescribeEventsPagesWithContext(ctx, input, fn, opts...)
}
