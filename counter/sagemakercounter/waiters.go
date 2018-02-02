// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakercounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sagemaker"
)

// WaitUntilEndpointDeleted calls WaitUntilEndpointDeletedWithContext with aws.BackgroundContext().
func (c *SageMaker) WaitUntilEndpointDeleted(input *sagemaker.DescribeEndpointInput) error {
	return c.WaitUntilEndpointDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilEndpointDeletedWithContext calls the underlying client method with a request option that
// will count DescribeEndpoint requests.
func (c *SageMaker) WaitUntilEndpointDeletedWithContext(ctx aws.Context, input *sagemaker.DescribeEndpointInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeEndpoint")))
	return c.svc.WaitUntilEndpointDeletedWithContext(ctx, input, opts...)
}

// WaitUntilEndpointInService calls WaitUntilEndpointInServiceWithContext with aws.BackgroundContext().
func (c *SageMaker) WaitUntilEndpointInService(input *sagemaker.DescribeEndpointInput) error {
	return c.WaitUntilEndpointInServiceWithContext(aws.BackgroundContext(), input)
}

// WaitUntilEndpointInServiceWithContext calls the underlying client method with a request option that
// will count DescribeEndpoint requests.
func (c *SageMaker) WaitUntilEndpointInServiceWithContext(ctx aws.Context, input *sagemaker.DescribeEndpointInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeEndpoint")))
	return c.svc.WaitUntilEndpointInServiceWithContext(ctx, input, opts...)
}

// WaitUntilNotebookInstanceDeleted calls WaitUntilNotebookInstanceDeletedWithContext with aws.BackgroundContext().
func (c *SageMaker) WaitUntilNotebookInstanceDeleted(input *sagemaker.DescribeNotebookInstanceInput) error {
	return c.WaitUntilNotebookInstanceDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilNotebookInstanceDeletedWithContext calls the underlying client method with a request option that
// will count DescribeNotebookInstance requests.
func (c *SageMaker) WaitUntilNotebookInstanceDeletedWithContext(ctx aws.Context, input *sagemaker.DescribeNotebookInstanceInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeNotebookInstance")))
	return c.svc.WaitUntilNotebookInstanceDeletedWithContext(ctx, input, opts...)
}

// WaitUntilNotebookInstanceInService calls WaitUntilNotebookInstanceInServiceWithContext with aws.BackgroundContext().
func (c *SageMaker) WaitUntilNotebookInstanceInService(input *sagemaker.DescribeNotebookInstanceInput) error {
	return c.WaitUntilNotebookInstanceInServiceWithContext(aws.BackgroundContext(), input)
}

// WaitUntilNotebookInstanceInServiceWithContext calls the underlying client method with a request option that
// will count DescribeNotebookInstance requests.
func (c *SageMaker) WaitUntilNotebookInstanceInServiceWithContext(ctx aws.Context, input *sagemaker.DescribeNotebookInstanceInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeNotebookInstance")))
	return c.svc.WaitUntilNotebookInstanceInServiceWithContext(ctx, input, opts...)
}

// WaitUntilNotebookInstanceStopped calls WaitUntilNotebookInstanceStoppedWithContext with aws.BackgroundContext().
func (c *SageMaker) WaitUntilNotebookInstanceStopped(input *sagemaker.DescribeNotebookInstanceInput) error {
	return c.WaitUntilNotebookInstanceStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilNotebookInstanceStoppedWithContext calls the underlying client method with a request option that
// will count DescribeNotebookInstance requests.
func (c *SageMaker) WaitUntilNotebookInstanceStoppedWithContext(ctx aws.Context, input *sagemaker.DescribeNotebookInstanceInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeNotebookInstance")))
	return c.svc.WaitUntilNotebookInstanceStoppedWithContext(ctx, input, opts...)
}

// WaitUntilTrainingJobCompletedOrStopped calls WaitUntilTrainingJobCompletedOrStoppedWithContext with aws.BackgroundContext().
func (c *SageMaker) WaitUntilTrainingJobCompletedOrStopped(input *sagemaker.DescribeTrainingJobInput) error {
	return c.WaitUntilTrainingJobCompletedOrStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTrainingJobCompletedOrStoppedWithContext calls the underlying client method with a request option that
// will count DescribeTrainingJob requests.
func (c *SageMaker) WaitUntilTrainingJobCompletedOrStoppedWithContext(ctx aws.Context, input *sagemaker.DescribeTrainingJobInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeTrainingJob")))
	return c.svc.WaitUntilTrainingJobCompletedOrStoppedWithContext(ctx, input, opts...)
}
