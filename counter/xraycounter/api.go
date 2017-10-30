// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package xraycounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/xray"
)

// BatchGetTracesRequest is a passthrough to the underlying BatchGetTracesRequest.
// It will increment the count of requests made to BatchGetTraces.
func (c *XRay) BatchGetTracesRequest(input *xray.BatchGetTracesInput) (req *request.Request, output *xray.BatchGetTracesOutput) {
	c.inc("BatchGetTraces")
	return c.svc.BatchGetTracesRequest(input)
}

// BatchGetTraces is a passthrough to the underlying BatchGetTraces method.
// It will increment the count of requests made to BatchGetTraces.
func (c *XRay) BatchGetTraces(input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
	c.inc("BatchGetTraces")
	return c.svc.BatchGetTraces(input)
}

// BatchGetTracesWithContext is a passthrough to the underlying BatchGetTracesWithContext method.
// It will increment the count of requests made to BatchGetTraces.
func (c *XRay) BatchGetTracesWithContext(ctx aws.Context, input *xray.BatchGetTracesInput, opts ...request.Option) (*xray.BatchGetTracesOutput, error) {
	c.inc("BatchGetTraces")
	return c.svc.BatchGetTracesWithContext(ctx, input, opts...)
}

// GetServiceGraphRequest is a passthrough to the underlying GetServiceGraphRequest.
// It will increment the count of requests made to GetServiceGraph.
func (c *XRay) GetServiceGraphRequest(input *xray.GetServiceGraphInput) (req *request.Request, output *xray.GetServiceGraphOutput) {
	c.inc("GetServiceGraph")
	return c.svc.GetServiceGraphRequest(input)
}

// GetServiceGraph is a passthrough to the underlying GetServiceGraph method.
// It will increment the count of requests made to GetServiceGraph.
func (c *XRay) GetServiceGraph(input *xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error) {
	c.inc("GetServiceGraph")
	return c.svc.GetServiceGraph(input)
}

// GetServiceGraphWithContext is a passthrough to the underlying GetServiceGraphWithContext method.
// It will increment the count of requests made to GetServiceGraph.
func (c *XRay) GetServiceGraphWithContext(ctx aws.Context, input *xray.GetServiceGraphInput, opts ...request.Option) (*xray.GetServiceGraphOutput, error) {
	c.inc("GetServiceGraph")
	return c.svc.GetServiceGraphWithContext(ctx, input, opts...)
}

// GetTraceGraphRequest is a passthrough to the underlying GetTraceGraphRequest.
// It will increment the count of requests made to GetTraceGraph.
func (c *XRay) GetTraceGraphRequest(input *xray.GetTraceGraphInput) (req *request.Request, output *xray.GetTraceGraphOutput) {
	c.inc("GetTraceGraph")
	return c.svc.GetTraceGraphRequest(input)
}

// GetTraceGraph is a passthrough to the underlying GetTraceGraph method.
// It will increment the count of requests made to GetTraceGraph.
func (c *XRay) GetTraceGraph(input *xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error) {
	c.inc("GetTraceGraph")
	return c.svc.GetTraceGraph(input)
}

// GetTraceGraphWithContext is a passthrough to the underlying GetTraceGraphWithContext method.
// It will increment the count of requests made to GetTraceGraph.
func (c *XRay) GetTraceGraphWithContext(ctx aws.Context, input *xray.GetTraceGraphInput, opts ...request.Option) (*xray.GetTraceGraphOutput, error) {
	c.inc("GetTraceGraph")
	return c.svc.GetTraceGraphWithContext(ctx, input, opts...)
}

// GetTraceSummariesRequest is a passthrough to the underlying GetTraceSummariesRequest.
// It will increment the count of requests made to GetTraceSummaries.
func (c *XRay) GetTraceSummariesRequest(input *xray.GetTraceSummariesInput) (req *request.Request, output *xray.GetTraceSummariesOutput) {
	c.inc("GetTraceSummaries")
	return c.svc.GetTraceSummariesRequest(input)
}

// GetTraceSummaries is a passthrough to the underlying GetTraceSummaries method.
// It will increment the count of requests made to GetTraceSummaries.
func (c *XRay) GetTraceSummaries(input *xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error) {
	c.inc("GetTraceSummaries")
	return c.svc.GetTraceSummaries(input)
}

// GetTraceSummariesWithContext is a passthrough to the underlying GetTraceSummariesWithContext method.
// It will increment the count of requests made to GetTraceSummaries.
func (c *XRay) GetTraceSummariesWithContext(ctx aws.Context, input *xray.GetTraceSummariesInput, opts ...request.Option) (*xray.GetTraceSummariesOutput, error) {
	c.inc("GetTraceSummaries")
	return c.svc.GetTraceSummariesWithContext(ctx, input, opts...)
}

// PutTelemetryRecordsRequest is a passthrough to the underlying PutTelemetryRecordsRequest.
// It will increment the count of requests made to PutTelemetryRecords.
func (c *XRay) PutTelemetryRecordsRequest(input *xray.PutTelemetryRecordsInput) (req *request.Request, output *xray.PutTelemetryRecordsOutput) {
	c.inc("PutTelemetryRecords")
	return c.svc.PutTelemetryRecordsRequest(input)
}

// PutTelemetryRecords is a passthrough to the underlying PutTelemetryRecords method.
// It will increment the count of requests made to PutTelemetryRecords.
func (c *XRay) PutTelemetryRecords(input *xray.PutTelemetryRecordsInput) (*xray.PutTelemetryRecordsOutput, error) {
	c.inc("PutTelemetryRecords")
	return c.svc.PutTelemetryRecords(input)
}

// PutTelemetryRecordsWithContext is a passthrough to the underlying PutTelemetryRecordsWithContext method.
// It will increment the count of requests made to PutTelemetryRecords.
func (c *XRay) PutTelemetryRecordsWithContext(ctx aws.Context, input *xray.PutTelemetryRecordsInput, opts ...request.Option) (*xray.PutTelemetryRecordsOutput, error) {
	c.inc("PutTelemetryRecords")
	return c.svc.PutTelemetryRecordsWithContext(ctx, input, opts...)
}

// PutTraceSegmentsRequest is a passthrough to the underlying PutTraceSegmentsRequest.
// It will increment the count of requests made to PutTraceSegments.
func (c *XRay) PutTraceSegmentsRequest(input *xray.PutTraceSegmentsInput) (req *request.Request, output *xray.PutTraceSegmentsOutput) {
	c.inc("PutTraceSegments")
	return c.svc.PutTraceSegmentsRequest(input)
}

// PutTraceSegments is a passthrough to the underlying PutTraceSegments method.
// It will increment the count of requests made to PutTraceSegments.
func (c *XRay) PutTraceSegments(input *xray.PutTraceSegmentsInput) (*xray.PutTraceSegmentsOutput, error) {
	c.inc("PutTraceSegments")
	return c.svc.PutTraceSegments(input)
}

// PutTraceSegmentsWithContext is a passthrough to the underlying PutTraceSegmentsWithContext method.
// It will increment the count of requests made to PutTraceSegments.
func (c *XRay) PutTraceSegmentsWithContext(ctx aws.Context, input *xray.PutTraceSegmentsInput, opts ...request.Option) (*xray.PutTraceSegmentsOutput, error) {
	c.inc("PutTraceSegments")
	return c.svc.PutTraceSegmentsWithContext(ctx, input, opts...)
}
