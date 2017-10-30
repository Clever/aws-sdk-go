// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudsearchdomaincounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
)

// SearchRequest is a passthrough to the underlying SearchRequest.
// It will increment the count of requests made to Search.
func (c *CloudSearchDomain) SearchRequest(input *cloudsearchdomain.SearchInput) (req *request.Request, output *cloudsearchdomain.SearchOutput) {
	c.inc("Search")
	return c.svc.SearchRequest(input)
}

// Search is a passthrough to the underlying Search method.
// It will increment the count of requests made to Search.
func (c *CloudSearchDomain) Search(input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
	c.inc("Search")
	return c.svc.Search(input)
}

// SearchWithContext is a passthrough to the underlying SearchWithContext method.
// It will increment the count of requests made to Search.
func (c *CloudSearchDomain) SearchWithContext(ctx aws.Context, input *cloudsearchdomain.SearchInput, opts ...request.Option) (*cloudsearchdomain.SearchOutput, error) {
	c.inc("Search")
	return c.svc.SearchWithContext(ctx, input, opts...)
}

// SuggestRequest is a passthrough to the underlying SuggestRequest.
// It will increment the count of requests made to Suggest.
func (c *CloudSearchDomain) SuggestRequest(input *cloudsearchdomain.SuggestInput) (req *request.Request, output *cloudsearchdomain.SuggestOutput) {
	c.inc("Suggest")
	return c.svc.SuggestRequest(input)
}

// Suggest is a passthrough to the underlying Suggest method.
// It will increment the count of requests made to Suggest.
func (c *CloudSearchDomain) Suggest(input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
	c.inc("Suggest")
	return c.svc.Suggest(input)
}

// SuggestWithContext is a passthrough to the underlying SuggestWithContext method.
// It will increment the count of requests made to Suggest.
func (c *CloudSearchDomain) SuggestWithContext(ctx aws.Context, input *cloudsearchdomain.SuggestInput, opts ...request.Option) (*cloudsearchdomain.SuggestOutput, error) {
	c.inc("Suggest")
	return c.svc.SuggestWithContext(ctx, input, opts...)
}

// UploadDocumentsRequest is a passthrough to the underlying UploadDocumentsRequest.
// It will increment the count of requests made to UploadDocuments.
func (c *CloudSearchDomain) UploadDocumentsRequest(input *cloudsearchdomain.UploadDocumentsInput) (req *request.Request, output *cloudsearchdomain.UploadDocumentsOutput) {
	c.inc("UploadDocuments")
	return c.svc.UploadDocumentsRequest(input)
}

// UploadDocuments is a passthrough to the underlying UploadDocuments method.
// It will increment the count of requests made to UploadDocuments.
func (c *CloudSearchDomain) UploadDocuments(input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	c.inc("UploadDocuments")
	return c.svc.UploadDocuments(input)
}

// UploadDocumentsWithContext is a passthrough to the underlying UploadDocumentsWithContext method.
// It will increment the count of requests made to UploadDocuments.
func (c *CloudSearchDomain) UploadDocumentsWithContext(ctx aws.Context, input *cloudsearchdomain.UploadDocumentsInput, opts ...request.Option) (*cloudsearchdomain.UploadDocumentsOutput, error) {
	c.inc("UploadDocuments")
	return c.svc.UploadDocumentsWithContext(ctx, input, opts...)
}
