// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqscounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// AddPermissionRequest is a passthrough to the underlying AddPermissionRequest.
// It will increment the count of requests made to AddPermission.
func (c *SQS) AddPermissionRequest(input *sqs.AddPermissionInput) (req *request.Request, output *sqs.AddPermissionOutput) {
	c.inc("AddPermission")
	return c.svc.AddPermissionRequest(input)
}

// AddPermission is a passthrough to the underlying AddPermission method.
// It will increment the count of requests made to AddPermission.
func (c *SQS) AddPermission(input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
	c.inc("AddPermission")
	return c.svc.AddPermission(input)
}

// AddPermissionWithContext is a passthrough to the underlying AddPermissionWithContext method.
// It will increment the count of requests made to AddPermission.
func (c *SQS) AddPermissionWithContext(ctx aws.Context, input *sqs.AddPermissionInput, opts ...request.Option) (*sqs.AddPermissionOutput, error) {
	c.inc("AddPermission")
	return c.svc.AddPermissionWithContext(ctx, input, opts...)
}

// ChangeMessageVisibilityRequest is a passthrough to the underlying ChangeMessageVisibilityRequest.
// It will increment the count of requests made to ChangeMessageVisibility.
func (c *SQS) ChangeMessageVisibilityRequest(input *sqs.ChangeMessageVisibilityInput) (req *request.Request, output *sqs.ChangeMessageVisibilityOutput) {
	c.inc("ChangeMessageVisibility")
	return c.svc.ChangeMessageVisibilityRequest(input)
}

// ChangeMessageVisibility is a passthrough to the underlying ChangeMessageVisibility method.
// It will increment the count of requests made to ChangeMessageVisibility.
func (c *SQS) ChangeMessageVisibility(input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
	c.inc("ChangeMessageVisibility")
	return c.svc.ChangeMessageVisibility(input)
}

// ChangeMessageVisibilityWithContext is a passthrough to the underlying ChangeMessageVisibilityWithContext method.
// It will increment the count of requests made to ChangeMessageVisibility.
func (c *SQS) ChangeMessageVisibilityWithContext(ctx aws.Context, input *sqs.ChangeMessageVisibilityInput, opts ...request.Option) (*sqs.ChangeMessageVisibilityOutput, error) {
	c.inc("ChangeMessageVisibility")
	return c.svc.ChangeMessageVisibilityWithContext(ctx, input, opts...)
}

// ChangeMessageVisibilityBatchRequest is a passthrough to the underlying ChangeMessageVisibilityBatchRequest.
// It will increment the count of requests made to ChangeMessageVisibilityBatch.
func (c *SQS) ChangeMessageVisibilityBatchRequest(input *sqs.ChangeMessageVisibilityBatchInput) (req *request.Request, output *sqs.ChangeMessageVisibilityBatchOutput) {
	c.inc("ChangeMessageVisibilityBatch")
	return c.svc.ChangeMessageVisibilityBatchRequest(input)
}

// ChangeMessageVisibilityBatch is a passthrough to the underlying ChangeMessageVisibilityBatch method.
// It will increment the count of requests made to ChangeMessageVisibilityBatch.
func (c *SQS) ChangeMessageVisibilityBatch(input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	c.inc("ChangeMessageVisibilityBatch")
	return c.svc.ChangeMessageVisibilityBatch(input)
}

// ChangeMessageVisibilityBatchWithContext is a passthrough to the underlying ChangeMessageVisibilityBatchWithContext method.
// It will increment the count of requests made to ChangeMessageVisibilityBatch.
func (c *SQS) ChangeMessageVisibilityBatchWithContext(ctx aws.Context, input *sqs.ChangeMessageVisibilityBatchInput, opts ...request.Option) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
	c.inc("ChangeMessageVisibilityBatch")
	return c.svc.ChangeMessageVisibilityBatchWithContext(ctx, input, opts...)
}

// CreateQueueRequest is a passthrough to the underlying CreateQueueRequest.
// It will increment the count of requests made to CreateQueue.
func (c *SQS) CreateQueueRequest(input *sqs.CreateQueueInput) (req *request.Request, output *sqs.CreateQueueOutput) {
	c.inc("CreateQueue")
	return c.svc.CreateQueueRequest(input)
}

// CreateQueue is a passthrough to the underlying CreateQueue method.
// It will increment the count of requests made to CreateQueue.
func (c *SQS) CreateQueue(input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
	c.inc("CreateQueue")
	return c.svc.CreateQueue(input)
}

// CreateQueueWithContext is a passthrough to the underlying CreateQueueWithContext method.
// It will increment the count of requests made to CreateQueue.
func (c *SQS) CreateQueueWithContext(ctx aws.Context, input *sqs.CreateQueueInput, opts ...request.Option) (*sqs.CreateQueueOutput, error) {
	c.inc("CreateQueue")
	return c.svc.CreateQueueWithContext(ctx, input, opts...)
}

// DeleteMessageRequest is a passthrough to the underlying DeleteMessageRequest.
// It will increment the count of requests made to DeleteMessage.
func (c *SQS) DeleteMessageRequest(input *sqs.DeleteMessageInput) (req *request.Request, output *sqs.DeleteMessageOutput) {
	c.inc("DeleteMessage")
	return c.svc.DeleteMessageRequest(input)
}

// DeleteMessage is a passthrough to the underlying DeleteMessage method.
// It will increment the count of requests made to DeleteMessage.
func (c *SQS) DeleteMessage(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	c.inc("DeleteMessage")
	return c.svc.DeleteMessage(input)
}

// DeleteMessageWithContext is a passthrough to the underlying DeleteMessageWithContext method.
// It will increment the count of requests made to DeleteMessage.
func (c *SQS) DeleteMessageWithContext(ctx aws.Context, input *sqs.DeleteMessageInput, opts ...request.Option) (*sqs.DeleteMessageOutput, error) {
	c.inc("DeleteMessage")
	return c.svc.DeleteMessageWithContext(ctx, input, opts...)
}

// DeleteMessageBatchRequest is a passthrough to the underlying DeleteMessageBatchRequest.
// It will increment the count of requests made to DeleteMessageBatch.
func (c *SQS) DeleteMessageBatchRequest(input *sqs.DeleteMessageBatchInput) (req *request.Request, output *sqs.DeleteMessageBatchOutput) {
	c.inc("DeleteMessageBatch")
	return c.svc.DeleteMessageBatchRequest(input)
}

// DeleteMessageBatch is a passthrough to the underlying DeleteMessageBatch method.
// It will increment the count of requests made to DeleteMessageBatch.
func (c *SQS) DeleteMessageBatch(input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	c.inc("DeleteMessageBatch")
	return c.svc.DeleteMessageBatch(input)
}

// DeleteMessageBatchWithContext is a passthrough to the underlying DeleteMessageBatchWithContext method.
// It will increment the count of requests made to DeleteMessageBatch.
func (c *SQS) DeleteMessageBatchWithContext(ctx aws.Context, input *sqs.DeleteMessageBatchInput, opts ...request.Option) (*sqs.DeleteMessageBatchOutput, error) {
	c.inc("DeleteMessageBatch")
	return c.svc.DeleteMessageBatchWithContext(ctx, input, opts...)
}

// DeleteQueueRequest is a passthrough to the underlying DeleteQueueRequest.
// It will increment the count of requests made to DeleteQueue.
func (c *SQS) DeleteQueueRequest(input *sqs.DeleteQueueInput) (req *request.Request, output *sqs.DeleteQueueOutput) {
	c.inc("DeleteQueue")
	return c.svc.DeleteQueueRequest(input)
}

// DeleteQueue is a passthrough to the underlying DeleteQueue method.
// It will increment the count of requests made to DeleteQueue.
func (c *SQS) DeleteQueue(input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
	c.inc("DeleteQueue")
	return c.svc.DeleteQueue(input)
}

// DeleteQueueWithContext is a passthrough to the underlying DeleteQueueWithContext method.
// It will increment the count of requests made to DeleteQueue.
func (c *SQS) DeleteQueueWithContext(ctx aws.Context, input *sqs.DeleteQueueInput, opts ...request.Option) (*sqs.DeleteQueueOutput, error) {
	c.inc("DeleteQueue")
	return c.svc.DeleteQueueWithContext(ctx, input, opts...)
}

// GetQueueAttributesRequest is a passthrough to the underlying GetQueueAttributesRequest.
// It will increment the count of requests made to GetQueueAttributes.
func (c *SQS) GetQueueAttributesRequest(input *sqs.GetQueueAttributesInput) (req *request.Request, output *sqs.GetQueueAttributesOutput) {
	c.inc("GetQueueAttributes")
	return c.svc.GetQueueAttributesRequest(input)
}

// GetQueueAttributes is a passthrough to the underlying GetQueueAttributes method.
// It will increment the count of requests made to GetQueueAttributes.
func (c *SQS) GetQueueAttributes(input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	c.inc("GetQueueAttributes")
	return c.svc.GetQueueAttributes(input)
}

// GetQueueAttributesWithContext is a passthrough to the underlying GetQueueAttributesWithContext method.
// It will increment the count of requests made to GetQueueAttributes.
func (c *SQS) GetQueueAttributesWithContext(ctx aws.Context, input *sqs.GetQueueAttributesInput, opts ...request.Option) (*sqs.GetQueueAttributesOutput, error) {
	c.inc("GetQueueAttributes")
	return c.svc.GetQueueAttributesWithContext(ctx, input, opts...)
}

// GetQueueUrlRequest is a passthrough to the underlying GetQueueUrlRequest.
// It will increment the count of requests made to GetQueueUrl.
func (c *SQS) GetQueueUrlRequest(input *sqs.GetQueueUrlInput) (req *request.Request, output *sqs.GetQueueUrlOutput) {
	c.inc("GetQueueUrl")
	return c.svc.GetQueueUrlRequest(input)
}

// GetQueueUrl is a passthrough to the underlying GetQueueUrl method.
// It will increment the count of requests made to GetQueueUrl.
func (c *SQS) GetQueueUrl(input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	c.inc("GetQueueUrl")
	return c.svc.GetQueueUrl(input)
}

// GetQueueUrlWithContext is a passthrough to the underlying GetQueueUrlWithContext method.
// It will increment the count of requests made to GetQueueUrl.
func (c *SQS) GetQueueUrlWithContext(ctx aws.Context, input *sqs.GetQueueUrlInput, opts ...request.Option) (*sqs.GetQueueUrlOutput, error) {
	c.inc("GetQueueUrl")
	return c.svc.GetQueueUrlWithContext(ctx, input, opts...)
}

// ListDeadLetterSourceQueuesRequest is a passthrough to the underlying ListDeadLetterSourceQueuesRequest.
// It will increment the count of requests made to ListDeadLetterSourceQueues.
func (c *SQS) ListDeadLetterSourceQueuesRequest(input *sqs.ListDeadLetterSourceQueuesInput) (req *request.Request, output *sqs.ListDeadLetterSourceQueuesOutput) {
	c.inc("ListDeadLetterSourceQueues")
	return c.svc.ListDeadLetterSourceQueuesRequest(input)
}

// ListDeadLetterSourceQueues is a passthrough to the underlying ListDeadLetterSourceQueues method.
// It will increment the count of requests made to ListDeadLetterSourceQueues.
func (c *SQS) ListDeadLetterSourceQueues(input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	c.inc("ListDeadLetterSourceQueues")
	return c.svc.ListDeadLetterSourceQueues(input)
}

// ListDeadLetterSourceQueuesWithContext is a passthrough to the underlying ListDeadLetterSourceQueuesWithContext method.
// It will increment the count of requests made to ListDeadLetterSourceQueues.
func (c *SQS) ListDeadLetterSourceQueuesWithContext(ctx aws.Context, input *sqs.ListDeadLetterSourceQueuesInput, opts ...request.Option) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	c.inc("ListDeadLetterSourceQueues")
	return c.svc.ListDeadLetterSourceQueuesWithContext(ctx, input, opts...)
}

// ListQueueTagsRequest is a passthrough to the underlying ListQueueTagsRequest.
// It will increment the count of requests made to ListQueueTags.
func (c *SQS) ListQueueTagsRequest(input *sqs.ListQueueTagsInput) (req *request.Request, output *sqs.ListQueueTagsOutput) {
	c.inc("ListQueueTags")
	return c.svc.ListQueueTagsRequest(input)
}

// ListQueueTags is a passthrough to the underlying ListQueueTags method.
// It will increment the count of requests made to ListQueueTags.
func (c *SQS) ListQueueTags(input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
	c.inc("ListQueueTags")
	return c.svc.ListQueueTags(input)
}

// ListQueueTagsWithContext is a passthrough to the underlying ListQueueTagsWithContext method.
// It will increment the count of requests made to ListQueueTags.
func (c *SQS) ListQueueTagsWithContext(ctx aws.Context, input *sqs.ListQueueTagsInput, opts ...request.Option) (*sqs.ListQueueTagsOutput, error) {
	c.inc("ListQueueTags")
	return c.svc.ListQueueTagsWithContext(ctx, input, opts...)
}

// ListQueuesRequest is a passthrough to the underlying ListQueuesRequest.
// It will increment the count of requests made to ListQueues.
func (c *SQS) ListQueuesRequest(input *sqs.ListQueuesInput) (req *request.Request, output *sqs.ListQueuesOutput) {
	c.inc("ListQueues")
	return c.svc.ListQueuesRequest(input)
}

// ListQueues is a passthrough to the underlying ListQueues method.
// It will increment the count of requests made to ListQueues.
func (c *SQS) ListQueues(input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	c.inc("ListQueues")
	return c.svc.ListQueues(input)
}

// ListQueuesWithContext is a passthrough to the underlying ListQueuesWithContext method.
// It will increment the count of requests made to ListQueues.
func (c *SQS) ListQueuesWithContext(ctx aws.Context, input *sqs.ListQueuesInput, opts ...request.Option) (*sqs.ListQueuesOutput, error) {
	c.inc("ListQueues")
	return c.svc.ListQueuesWithContext(ctx, input, opts...)
}

// PurgeQueueRequest is a passthrough to the underlying PurgeQueueRequest.
// It will increment the count of requests made to PurgeQueue.
func (c *SQS) PurgeQueueRequest(input *sqs.PurgeQueueInput) (req *request.Request, output *sqs.PurgeQueueOutput) {
	c.inc("PurgeQueue")
	return c.svc.PurgeQueueRequest(input)
}

// PurgeQueue is a passthrough to the underlying PurgeQueue method.
// It will increment the count of requests made to PurgeQueue.
func (c *SQS) PurgeQueue(input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
	c.inc("PurgeQueue")
	return c.svc.PurgeQueue(input)
}

// PurgeQueueWithContext is a passthrough to the underlying PurgeQueueWithContext method.
// It will increment the count of requests made to PurgeQueue.
func (c *SQS) PurgeQueueWithContext(ctx aws.Context, input *sqs.PurgeQueueInput, opts ...request.Option) (*sqs.PurgeQueueOutput, error) {
	c.inc("PurgeQueue")
	return c.svc.PurgeQueueWithContext(ctx, input, opts...)
}

// ReceiveMessageRequest is a passthrough to the underlying ReceiveMessageRequest.
// It will increment the count of requests made to ReceiveMessage.
func (c *SQS) ReceiveMessageRequest(input *sqs.ReceiveMessageInput) (req *request.Request, output *sqs.ReceiveMessageOutput) {
	c.inc("ReceiveMessage")
	return c.svc.ReceiveMessageRequest(input)
}

// ReceiveMessage is a passthrough to the underlying ReceiveMessage method.
// It will increment the count of requests made to ReceiveMessage.
func (c *SQS) ReceiveMessage(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	c.inc("ReceiveMessage")
	return c.svc.ReceiveMessage(input)
}

// ReceiveMessageWithContext is a passthrough to the underlying ReceiveMessageWithContext method.
// It will increment the count of requests made to ReceiveMessage.
func (c *SQS) ReceiveMessageWithContext(ctx aws.Context, input *sqs.ReceiveMessageInput, opts ...request.Option) (*sqs.ReceiveMessageOutput, error) {
	c.inc("ReceiveMessage")
	return c.svc.ReceiveMessageWithContext(ctx, input, opts...)
}

// RemovePermissionRequest is a passthrough to the underlying RemovePermissionRequest.
// It will increment the count of requests made to RemovePermission.
func (c *SQS) RemovePermissionRequest(input *sqs.RemovePermissionInput) (req *request.Request, output *sqs.RemovePermissionOutput) {
	c.inc("RemovePermission")
	return c.svc.RemovePermissionRequest(input)
}

// RemovePermission is a passthrough to the underlying RemovePermission method.
// It will increment the count of requests made to RemovePermission.
func (c *SQS) RemovePermission(input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
	c.inc("RemovePermission")
	return c.svc.RemovePermission(input)
}

// RemovePermissionWithContext is a passthrough to the underlying RemovePermissionWithContext method.
// It will increment the count of requests made to RemovePermission.
func (c *SQS) RemovePermissionWithContext(ctx aws.Context, input *sqs.RemovePermissionInput, opts ...request.Option) (*sqs.RemovePermissionOutput, error) {
	c.inc("RemovePermission")
	return c.svc.RemovePermissionWithContext(ctx, input, opts...)
}

// SendMessageRequest is a passthrough to the underlying SendMessageRequest.
// It will increment the count of requests made to SendMessage.
func (c *SQS) SendMessageRequest(input *sqs.SendMessageInput) (req *request.Request, output *sqs.SendMessageOutput) {
	c.inc("SendMessage")
	return c.svc.SendMessageRequest(input)
}

// SendMessage is a passthrough to the underlying SendMessage method.
// It will increment the count of requests made to SendMessage.
func (c *SQS) SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	c.inc("SendMessage")
	return c.svc.SendMessage(input)
}

// SendMessageWithContext is a passthrough to the underlying SendMessageWithContext method.
// It will increment the count of requests made to SendMessage.
func (c *SQS) SendMessageWithContext(ctx aws.Context, input *sqs.SendMessageInput, opts ...request.Option) (*sqs.SendMessageOutput, error) {
	c.inc("SendMessage")
	return c.svc.SendMessageWithContext(ctx, input, opts...)
}

// SendMessageBatchRequest is a passthrough to the underlying SendMessageBatchRequest.
// It will increment the count of requests made to SendMessageBatch.
func (c *SQS) SendMessageBatchRequest(input *sqs.SendMessageBatchInput) (req *request.Request, output *sqs.SendMessageBatchOutput) {
	c.inc("SendMessageBatch")
	return c.svc.SendMessageBatchRequest(input)
}

// SendMessageBatch is a passthrough to the underlying SendMessageBatch method.
// It will increment the count of requests made to SendMessageBatch.
func (c *SQS) SendMessageBatch(input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
	c.inc("SendMessageBatch")
	return c.svc.SendMessageBatch(input)
}

// SendMessageBatchWithContext is a passthrough to the underlying SendMessageBatchWithContext method.
// It will increment the count of requests made to SendMessageBatch.
func (c *SQS) SendMessageBatchWithContext(ctx aws.Context, input *sqs.SendMessageBatchInput, opts ...request.Option) (*sqs.SendMessageBatchOutput, error) {
	c.inc("SendMessageBatch")
	return c.svc.SendMessageBatchWithContext(ctx, input, opts...)
}

// SetQueueAttributesRequest is a passthrough to the underlying SetQueueAttributesRequest.
// It will increment the count of requests made to SetQueueAttributes.
func (c *SQS) SetQueueAttributesRequest(input *sqs.SetQueueAttributesInput) (req *request.Request, output *sqs.SetQueueAttributesOutput) {
	c.inc("SetQueueAttributes")
	return c.svc.SetQueueAttributesRequest(input)
}

// SetQueueAttributes is a passthrough to the underlying SetQueueAttributes method.
// It will increment the count of requests made to SetQueueAttributes.
func (c *SQS) SetQueueAttributes(input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
	c.inc("SetQueueAttributes")
	return c.svc.SetQueueAttributes(input)
}

// SetQueueAttributesWithContext is a passthrough to the underlying SetQueueAttributesWithContext method.
// It will increment the count of requests made to SetQueueAttributes.
func (c *SQS) SetQueueAttributesWithContext(ctx aws.Context, input *sqs.SetQueueAttributesInput, opts ...request.Option) (*sqs.SetQueueAttributesOutput, error) {
	c.inc("SetQueueAttributes")
	return c.svc.SetQueueAttributesWithContext(ctx, input, opts...)
}

// TagQueueRequest is a passthrough to the underlying TagQueueRequest.
// It will increment the count of requests made to TagQueue.
func (c *SQS) TagQueueRequest(input *sqs.TagQueueInput) (req *request.Request, output *sqs.TagQueueOutput) {
	c.inc("TagQueue")
	return c.svc.TagQueueRequest(input)
}

// TagQueue is a passthrough to the underlying TagQueue method.
// It will increment the count of requests made to TagQueue.
func (c *SQS) TagQueue(input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
	c.inc("TagQueue")
	return c.svc.TagQueue(input)
}

// TagQueueWithContext is a passthrough to the underlying TagQueueWithContext method.
// It will increment the count of requests made to TagQueue.
func (c *SQS) TagQueueWithContext(ctx aws.Context, input *sqs.TagQueueInput, opts ...request.Option) (*sqs.TagQueueOutput, error) {
	c.inc("TagQueue")
	return c.svc.TagQueueWithContext(ctx, input, opts...)
}

// UntagQueueRequest is a passthrough to the underlying UntagQueueRequest.
// It will increment the count of requests made to UntagQueue.
func (c *SQS) UntagQueueRequest(input *sqs.UntagQueueInput) (req *request.Request, output *sqs.UntagQueueOutput) {
	c.inc("UntagQueue")
	return c.svc.UntagQueueRequest(input)
}

// UntagQueue is a passthrough to the underlying UntagQueue method.
// It will increment the count of requests made to UntagQueue.
func (c *SQS) UntagQueue(input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
	c.inc("UntagQueue")
	return c.svc.UntagQueue(input)
}

// UntagQueueWithContext is a passthrough to the underlying UntagQueueWithContext method.
// It will increment the count of requests made to UntagQueue.
func (c *SQS) UntagQueueWithContext(ctx aws.Context, input *sqs.UntagQueueInput, opts ...request.Option) (*sqs.UntagQueueOutput, error) {
	c.inc("UntagQueue")
	return c.svc.UntagQueueWithContext(ctx, input, opts...)
}
