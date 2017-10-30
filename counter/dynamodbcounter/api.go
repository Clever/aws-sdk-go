// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbcounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// BatchGetItemRequest is a passthrough to the underlying BatchGetItemRequest.
// It will increment the count of requests made to BatchGetItem.
func (c *DynamoDB) BatchGetItemRequest(input *dynamodb.BatchGetItemInput) (req *request.Request, output *dynamodb.BatchGetItemOutput) {
	c.inc("BatchGetItem")
	return c.svc.BatchGetItemRequest(input)
}

// BatchGetItem is a passthrough to the underlying BatchGetItem method.
// It will increment the count of requests made to BatchGetItem.
func (c *DynamoDB) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	c.inc("BatchGetItem")
	return c.svc.BatchGetItem(input)
}

// BatchGetItemWithContext is a passthrough to the underlying BatchGetItemWithContext method.
// It will increment the count of requests made to BatchGetItem.
func (c *DynamoDB) BatchGetItemWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	c.inc("BatchGetItem")
	return c.svc.BatchGetItemWithContext(ctx, input, opts...)
}

// BatchGetItemPages is a passthrough to the underlying BatchGetItemPages method.
// It will increment the count of requests made to BatchGetItem on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use BatchGetItemPagesWithContext to avoid this.
func (c *DynamoDB) BatchGetItemPages(input *dynamodb.BatchGetItemInput, fn func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	wrappedFn := func(page *dynamodb.BatchGetItemOutput, lastPage bool) bool {
		c.inc("BatchGetItem")
		return fn(page, lastPage)
	}
	return c.BatchGetItemPages(input, wrappedFn)
}

// BatchGetItemPagesWithContext is a passthrough to the underlying BatchGetItemPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to BatchGetItem when applied to the request.
func (c *DynamoDB) BatchGetItemPagesWithContext(ctx aws.Context, input *dynamodb.BatchGetItemInput, fn func(*dynamodb.BatchGetItemOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("BatchGetItem"))
	return c.BatchGetItemPagesWithContext(ctx, input, fn, opts...)
}

// BatchWriteItemRequest is a passthrough to the underlying BatchWriteItemRequest.
// It will increment the count of requests made to BatchWriteItem.
func (c *DynamoDB) BatchWriteItemRequest(input *dynamodb.BatchWriteItemInput) (req *request.Request, output *dynamodb.BatchWriteItemOutput) {
	c.inc("BatchWriteItem")
	return c.svc.BatchWriteItemRequest(input)
}

// BatchWriteItem is a passthrough to the underlying BatchWriteItem method.
// It will increment the count of requests made to BatchWriteItem.
func (c *DynamoDB) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	c.inc("BatchWriteItem")
	return c.svc.BatchWriteItem(input)
}

// BatchWriteItemWithContext is a passthrough to the underlying BatchWriteItemWithContext method.
// It will increment the count of requests made to BatchWriteItem.
func (c *DynamoDB) BatchWriteItemWithContext(ctx aws.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	c.inc("BatchWriteItem")
	return c.svc.BatchWriteItemWithContext(ctx, input, opts...)
}

// CreateTableRequest is a passthrough to the underlying CreateTableRequest.
// It will increment the count of requests made to CreateTable.
func (c *DynamoDB) CreateTableRequest(input *dynamodb.CreateTableInput) (req *request.Request, output *dynamodb.CreateTableOutput) {
	c.inc("CreateTable")
	return c.svc.CreateTableRequest(input)
}

// CreateTable is a passthrough to the underlying CreateTable method.
// It will increment the count of requests made to CreateTable.
func (c *DynamoDB) CreateTable(input *dynamodb.CreateTableInput) (*dynamodb.CreateTableOutput, error) {
	c.inc("CreateTable")
	return c.svc.CreateTable(input)
}

// CreateTableWithContext is a passthrough to the underlying CreateTableWithContext method.
// It will increment the count of requests made to CreateTable.
func (c *DynamoDB) CreateTableWithContext(ctx aws.Context, input *dynamodb.CreateTableInput, opts ...request.Option) (*dynamodb.CreateTableOutput, error) {
	c.inc("CreateTable")
	return c.svc.CreateTableWithContext(ctx, input, opts...)
}

// DeleteItemRequest is a passthrough to the underlying DeleteItemRequest.
// It will increment the count of requests made to DeleteItem.
func (c *DynamoDB) DeleteItemRequest(input *dynamodb.DeleteItemInput) (req *request.Request, output *dynamodb.DeleteItemOutput) {
	c.inc("DeleteItem")
	return c.svc.DeleteItemRequest(input)
}

// DeleteItem is a passthrough to the underlying DeleteItem method.
// It will increment the count of requests made to DeleteItem.
func (c *DynamoDB) DeleteItem(input *dynamodb.DeleteItemInput) (*dynamodb.DeleteItemOutput, error) {
	c.inc("DeleteItem")
	return c.svc.DeleteItem(input)
}

// DeleteItemWithContext is a passthrough to the underlying DeleteItemWithContext method.
// It will increment the count of requests made to DeleteItem.
func (c *DynamoDB) DeleteItemWithContext(ctx aws.Context, input *dynamodb.DeleteItemInput, opts ...request.Option) (*dynamodb.DeleteItemOutput, error) {
	c.inc("DeleteItem")
	return c.svc.DeleteItemWithContext(ctx, input, opts...)
}

// DeleteTableRequest is a passthrough to the underlying DeleteTableRequest.
// It will increment the count of requests made to DeleteTable.
func (c *DynamoDB) DeleteTableRequest(input *dynamodb.DeleteTableInput) (req *request.Request, output *dynamodb.DeleteTableOutput) {
	c.inc("DeleteTable")
	return c.svc.DeleteTableRequest(input)
}

// DeleteTable is a passthrough to the underlying DeleteTable method.
// It will increment the count of requests made to DeleteTable.
func (c *DynamoDB) DeleteTable(input *dynamodb.DeleteTableInput) (*dynamodb.DeleteTableOutput, error) {
	c.inc("DeleteTable")
	return c.svc.DeleteTable(input)
}

// DeleteTableWithContext is a passthrough to the underlying DeleteTableWithContext method.
// It will increment the count of requests made to DeleteTable.
func (c *DynamoDB) DeleteTableWithContext(ctx aws.Context, input *dynamodb.DeleteTableInput, opts ...request.Option) (*dynamodb.DeleteTableOutput, error) {
	c.inc("DeleteTable")
	return c.svc.DeleteTableWithContext(ctx, input, opts...)
}

// DescribeLimitsRequest is a passthrough to the underlying DescribeLimitsRequest.
// It will increment the count of requests made to DescribeLimits.
func (c *DynamoDB) DescribeLimitsRequest(input *dynamodb.DescribeLimitsInput) (req *request.Request, output *dynamodb.DescribeLimitsOutput) {
	c.inc("DescribeLimits")
	return c.svc.DescribeLimitsRequest(input)
}

// DescribeLimits is a passthrough to the underlying DescribeLimits method.
// It will increment the count of requests made to DescribeLimits.
func (c *DynamoDB) DescribeLimits(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	c.inc("DescribeLimits")
	return c.svc.DescribeLimits(input)
}

// DescribeLimitsWithContext is a passthrough to the underlying DescribeLimitsWithContext method.
// It will increment the count of requests made to DescribeLimits.
func (c *DynamoDB) DescribeLimitsWithContext(ctx aws.Context, input *dynamodb.DescribeLimitsInput, opts ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	c.inc("DescribeLimits")
	return c.svc.DescribeLimitsWithContext(ctx, input, opts...)
}

// DescribeTableRequest is a passthrough to the underlying DescribeTableRequest.
// It will increment the count of requests made to DescribeTable.
func (c *DynamoDB) DescribeTableRequest(input *dynamodb.DescribeTableInput) (req *request.Request, output *dynamodb.DescribeTableOutput) {
	c.inc("DescribeTable")
	return c.svc.DescribeTableRequest(input)
}

// DescribeTable is a passthrough to the underlying DescribeTable method.
// It will increment the count of requests made to DescribeTable.
func (c *DynamoDB) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	c.inc("DescribeTable")
	return c.svc.DescribeTable(input)
}

// DescribeTableWithContext is a passthrough to the underlying DescribeTableWithContext method.
// It will increment the count of requests made to DescribeTable.
func (c *DynamoDB) DescribeTableWithContext(ctx aws.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	c.inc("DescribeTable")
	return c.svc.DescribeTableWithContext(ctx, input, opts...)
}

// DescribeTimeToLiveRequest is a passthrough to the underlying DescribeTimeToLiveRequest.
// It will increment the count of requests made to DescribeTimeToLive.
func (c *DynamoDB) DescribeTimeToLiveRequest(input *dynamodb.DescribeTimeToLiveInput) (req *request.Request, output *dynamodb.DescribeTimeToLiveOutput) {
	c.inc("DescribeTimeToLive")
	return c.svc.DescribeTimeToLiveRequest(input)
}

// DescribeTimeToLive is a passthrough to the underlying DescribeTimeToLive method.
// It will increment the count of requests made to DescribeTimeToLive.
func (c *DynamoDB) DescribeTimeToLive(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	c.inc("DescribeTimeToLive")
	return c.svc.DescribeTimeToLive(input)
}

// DescribeTimeToLiveWithContext is a passthrough to the underlying DescribeTimeToLiveWithContext method.
// It will increment the count of requests made to DescribeTimeToLive.
func (c *DynamoDB) DescribeTimeToLiveWithContext(ctx aws.Context, input *dynamodb.DescribeTimeToLiveInput, opts ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	c.inc("DescribeTimeToLive")
	return c.svc.DescribeTimeToLiveWithContext(ctx, input, opts...)
}

// GetItemRequest is a passthrough to the underlying GetItemRequest.
// It will increment the count of requests made to GetItem.
func (c *DynamoDB) GetItemRequest(input *dynamodb.GetItemInput) (req *request.Request, output *dynamodb.GetItemOutput) {
	c.inc("GetItem")
	return c.svc.GetItemRequest(input)
}

// GetItem is a passthrough to the underlying GetItem method.
// It will increment the count of requests made to GetItem.
func (c *DynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	c.inc("GetItem")
	return c.svc.GetItem(input)
}

// GetItemWithContext is a passthrough to the underlying GetItemWithContext method.
// It will increment the count of requests made to GetItem.
func (c *DynamoDB) GetItemWithContext(ctx aws.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	c.inc("GetItem")
	return c.svc.GetItemWithContext(ctx, input, opts...)
}

// ListTablesRequest is a passthrough to the underlying ListTablesRequest.
// It will increment the count of requests made to ListTables.
func (c *DynamoDB) ListTablesRequest(input *dynamodb.ListTablesInput) (req *request.Request, output *dynamodb.ListTablesOutput) {
	c.inc("ListTables")
	return c.svc.ListTablesRequest(input)
}

// ListTables is a passthrough to the underlying ListTables method.
// It will increment the count of requests made to ListTables.
func (c *DynamoDB) ListTables(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	c.inc("ListTables")
	return c.svc.ListTables(input)
}

// ListTablesWithContext is a passthrough to the underlying ListTablesWithContext method.
// It will increment the count of requests made to ListTables.
func (c *DynamoDB) ListTablesWithContext(ctx aws.Context, input *dynamodb.ListTablesInput, opts ...request.Option) (*dynamodb.ListTablesOutput, error) {
	c.inc("ListTables")
	return c.svc.ListTablesWithContext(ctx, input, opts...)
}

// ListTablesPages is a passthrough to the underlying ListTablesPages method.
// It will increment the count of requests made to ListTables on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ListTablesPagesWithContext to avoid this.
func (c *DynamoDB) ListTablesPages(input *dynamodb.ListTablesInput, fn func(*dynamodb.ListTablesOutput, bool) bool) error {
	wrappedFn := func(page *dynamodb.ListTablesOutput, lastPage bool) bool {
		c.inc("ListTables")
		return fn(page, lastPage)
	}
	return c.ListTablesPages(input, wrappedFn)
}

// ListTablesPagesWithContext is a passthrough to the underlying ListTablesPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to ListTables when applied to the request.
func (c *DynamoDB) ListTablesPagesWithContext(ctx aws.Context, input *dynamodb.ListTablesInput, fn func(*dynamodb.ListTablesOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("ListTables"))
	return c.ListTablesPagesWithContext(ctx, input, fn, opts...)
}

// ListTagsOfResourceRequest is a passthrough to the underlying ListTagsOfResourceRequest.
// It will increment the count of requests made to ListTagsOfResource.
func (c *DynamoDB) ListTagsOfResourceRequest(input *dynamodb.ListTagsOfResourceInput) (req *request.Request, output *dynamodb.ListTagsOfResourceOutput) {
	c.inc("ListTagsOfResource")
	return c.svc.ListTagsOfResourceRequest(input)
}

// ListTagsOfResource is a passthrough to the underlying ListTagsOfResource method.
// It will increment the count of requests made to ListTagsOfResource.
func (c *DynamoDB) ListTagsOfResource(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	c.inc("ListTagsOfResource")
	return c.svc.ListTagsOfResource(input)
}

// ListTagsOfResourceWithContext is a passthrough to the underlying ListTagsOfResourceWithContext method.
// It will increment the count of requests made to ListTagsOfResource.
func (c *DynamoDB) ListTagsOfResourceWithContext(ctx aws.Context, input *dynamodb.ListTagsOfResourceInput, opts ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	c.inc("ListTagsOfResource")
	return c.svc.ListTagsOfResourceWithContext(ctx, input, opts...)
}

// PutItemRequest is a passthrough to the underlying PutItemRequest.
// It will increment the count of requests made to PutItem.
func (c *DynamoDB) PutItemRequest(input *dynamodb.PutItemInput) (req *request.Request, output *dynamodb.PutItemOutput) {
	c.inc("PutItem")
	return c.svc.PutItemRequest(input)
}

// PutItem is a passthrough to the underlying PutItem method.
// It will increment the count of requests made to PutItem.
func (c *DynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	c.inc("PutItem")
	return c.svc.PutItem(input)
}

// PutItemWithContext is a passthrough to the underlying PutItemWithContext method.
// It will increment the count of requests made to PutItem.
func (c *DynamoDB) PutItemWithContext(ctx aws.Context, input *dynamodb.PutItemInput, opts ...request.Option) (*dynamodb.PutItemOutput, error) {
	c.inc("PutItem")
	return c.svc.PutItemWithContext(ctx, input, opts...)
}

// QueryRequest is a passthrough to the underlying QueryRequest.
// It will increment the count of requests made to Query.
func (c *DynamoDB) QueryRequest(input *dynamodb.QueryInput) (req *request.Request, output *dynamodb.QueryOutput) {
	c.inc("Query")
	return c.svc.QueryRequest(input)
}

// Query is a passthrough to the underlying Query method.
// It will increment the count of requests made to Query.
func (c *DynamoDB) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	c.inc("Query")
	return c.svc.Query(input)
}

// QueryWithContext is a passthrough to the underlying QueryWithContext method.
// It will increment the count of requests made to Query.
func (c *DynamoDB) QueryWithContext(ctx aws.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error) {
	c.inc("Query")
	return c.svc.QueryWithContext(ctx, input, opts...)
}

// QueryPages is a passthrough to the underlying QueryPages method.
// It will increment the count of requests made to Query on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use QueryPagesWithContext to avoid this.
func (c *DynamoDB) QueryPages(input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error {
	wrappedFn := func(page *dynamodb.QueryOutput, lastPage bool) bool {
		c.inc("Query")
		return fn(page, lastPage)
	}
	return c.QueryPages(input, wrappedFn)
}

// QueryPagesWithContext is a passthrough to the underlying QueryPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to Query when applied to the request.
func (c *DynamoDB) QueryPagesWithContext(ctx aws.Context, input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("Query"))
	return c.QueryPagesWithContext(ctx, input, fn, opts...)
}

// ScanRequest is a passthrough to the underlying ScanRequest.
// It will increment the count of requests made to Scan.
func (c *DynamoDB) ScanRequest(input *dynamodb.ScanInput) (req *request.Request, output *dynamodb.ScanOutput) {
	c.inc("Scan")
	return c.svc.ScanRequest(input)
}

// Scan is a passthrough to the underlying Scan method.
// It will increment the count of requests made to Scan.
func (c *DynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	c.inc("Scan")
	return c.svc.Scan(input)
}

// ScanWithContext is a passthrough to the underlying ScanWithContext method.
// It will increment the count of requests made to Scan.
func (c *DynamoDB) ScanWithContext(ctx aws.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	c.inc("Scan")
	return c.svc.ScanWithContext(ctx, input, opts...)
}

// ScanPages is a passthrough to the underlying ScanPages method.
// It will increment the count of requests made to Scan on each page.
// NOTE: this is slightly inaccurate in the case of errors, since the function will not be called.
// Use ScanPagesWithContext to avoid this.
func (c *DynamoDB) ScanPages(input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool) error {
	wrappedFn := func(page *dynamodb.ScanOutput, lastPage bool) bool {
		c.inc("Scan")
		return fn(page, lastPage)
	}
	return c.ScanPages(input, wrappedFn)
}

// ScanPagesWithContext is a passthrough to the underlying ScanPagesWithContext method.
// It will add a request.Option that will increment the count of requests made to Scan when applied to the request.
func (c *DynamoDB) ScanPagesWithContext(ctx aws.Context, input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, opts ...request.Option) error {
	opts = append(opts, c.incViaRequestOption("Scan"))
	return c.ScanPagesWithContext(ctx, input, fn, opts...)
}

// TagResourceRequest is a passthrough to the underlying TagResourceRequest.
// It will increment the count of requests made to TagResource.
func (c *DynamoDB) TagResourceRequest(input *dynamodb.TagResourceInput) (req *request.Request, output *dynamodb.TagResourceOutput) {
	c.inc("TagResource")
	return c.svc.TagResourceRequest(input)
}

// TagResource is a passthrough to the underlying TagResource method.
// It will increment the count of requests made to TagResource.
func (c *DynamoDB) TagResource(input *dynamodb.TagResourceInput) (*dynamodb.TagResourceOutput, error) {
	c.inc("TagResource")
	return c.svc.TagResource(input)
}

// TagResourceWithContext is a passthrough to the underlying TagResourceWithContext method.
// It will increment the count of requests made to TagResource.
func (c *DynamoDB) TagResourceWithContext(ctx aws.Context, input *dynamodb.TagResourceInput, opts ...request.Option) (*dynamodb.TagResourceOutput, error) {
	c.inc("TagResource")
	return c.svc.TagResourceWithContext(ctx, input, opts...)
}

// UntagResourceRequest is a passthrough to the underlying UntagResourceRequest.
// It will increment the count of requests made to UntagResource.
func (c *DynamoDB) UntagResourceRequest(input *dynamodb.UntagResourceInput) (req *request.Request, output *dynamodb.UntagResourceOutput) {
	c.inc("UntagResource")
	return c.svc.UntagResourceRequest(input)
}

// UntagResource is a passthrough to the underlying UntagResource method.
// It will increment the count of requests made to UntagResource.
func (c *DynamoDB) UntagResource(input *dynamodb.UntagResourceInput) (*dynamodb.UntagResourceOutput, error) {
	c.inc("UntagResource")
	return c.svc.UntagResource(input)
}

// UntagResourceWithContext is a passthrough to the underlying UntagResourceWithContext method.
// It will increment the count of requests made to UntagResource.
func (c *DynamoDB) UntagResourceWithContext(ctx aws.Context, input *dynamodb.UntagResourceInput, opts ...request.Option) (*dynamodb.UntagResourceOutput, error) {
	c.inc("UntagResource")
	return c.svc.UntagResourceWithContext(ctx, input, opts...)
}

// UpdateItemRequest is a passthrough to the underlying UpdateItemRequest.
// It will increment the count of requests made to UpdateItem.
func (c *DynamoDB) UpdateItemRequest(input *dynamodb.UpdateItemInput) (req *request.Request, output *dynamodb.UpdateItemOutput) {
	c.inc("UpdateItem")
	return c.svc.UpdateItemRequest(input)
}

// UpdateItem is a passthrough to the underlying UpdateItem method.
// It will increment the count of requests made to UpdateItem.
func (c *DynamoDB) UpdateItem(input *dynamodb.UpdateItemInput) (*dynamodb.UpdateItemOutput, error) {
	c.inc("UpdateItem")
	return c.svc.UpdateItem(input)
}

// UpdateItemWithContext is a passthrough to the underlying UpdateItemWithContext method.
// It will increment the count of requests made to UpdateItem.
func (c *DynamoDB) UpdateItemWithContext(ctx aws.Context, input *dynamodb.UpdateItemInput, opts ...request.Option) (*dynamodb.UpdateItemOutput, error) {
	c.inc("UpdateItem")
	return c.svc.UpdateItemWithContext(ctx, input, opts...)
}

// UpdateTableRequest is a passthrough to the underlying UpdateTableRequest.
// It will increment the count of requests made to UpdateTable.
func (c *DynamoDB) UpdateTableRequest(input *dynamodb.UpdateTableInput) (req *request.Request, output *dynamodb.UpdateTableOutput) {
	c.inc("UpdateTable")
	return c.svc.UpdateTableRequest(input)
}

// UpdateTable is a passthrough to the underlying UpdateTable method.
// It will increment the count of requests made to UpdateTable.
func (c *DynamoDB) UpdateTable(input *dynamodb.UpdateTableInput) (*dynamodb.UpdateTableOutput, error) {
	c.inc("UpdateTable")
	return c.svc.UpdateTable(input)
}

// UpdateTableWithContext is a passthrough to the underlying UpdateTableWithContext method.
// It will increment the count of requests made to UpdateTable.
func (c *DynamoDB) UpdateTableWithContext(ctx aws.Context, input *dynamodb.UpdateTableInput, opts ...request.Option) (*dynamodb.UpdateTableOutput, error) {
	c.inc("UpdateTable")
	return c.svc.UpdateTableWithContext(ctx, input, opts...)
}

// UpdateTimeToLiveRequest is a passthrough to the underlying UpdateTimeToLiveRequest.
// It will increment the count of requests made to UpdateTimeToLive.
func (c *DynamoDB) UpdateTimeToLiveRequest(input *dynamodb.UpdateTimeToLiveInput) (req *request.Request, output *dynamodb.UpdateTimeToLiveOutput) {
	c.inc("UpdateTimeToLive")
	return c.svc.UpdateTimeToLiveRequest(input)
}

// UpdateTimeToLive is a passthrough to the underlying UpdateTimeToLive method.
// It will increment the count of requests made to UpdateTimeToLive.
func (c *DynamoDB) UpdateTimeToLive(input *dynamodb.UpdateTimeToLiveInput) (*dynamodb.UpdateTimeToLiveOutput, error) {
	c.inc("UpdateTimeToLive")
	return c.svc.UpdateTimeToLive(input)
}

// UpdateTimeToLiveWithContext is a passthrough to the underlying UpdateTimeToLiveWithContext method.
// It will increment the count of requests made to UpdateTimeToLive.
func (c *DynamoDB) UpdateTimeToLiveWithContext(ctx aws.Context, input *dynamodb.UpdateTimeToLiveInput, opts ...request.Option) (*dynamodb.UpdateTimeToLiveOutput, error) {
	c.inc("UpdateTimeToLive")
	return c.svc.UpdateTimeToLiveWithContext(ctx, input, opts...)
}
