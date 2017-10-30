// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshiftcounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/redshift"
)

// WaitUntilClusterAvailable calls WaitUntilClusterAvailableWithContext with aws.BackgroundContext().
func (c *Redshift) WaitUntilClusterAvailable(input *redshift.DescribeClustersInput) error {
	return c.WaitUntilClusterAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterAvailableWithContext calls the underlying client method with a request option that
// will count DescribeClusters requests.
func (c *Redshift) WaitUntilClusterAvailableWithContext(ctx aws.Context, input *redshift.DescribeClustersInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeClusters")))
	return c.svc.WaitUntilClusterAvailableWithContext(ctx, input, opts...)
}

// WaitUntilClusterDeleted calls WaitUntilClusterDeletedWithContext with aws.BackgroundContext().
func (c *Redshift) WaitUntilClusterDeleted(input *redshift.DescribeClustersInput) error {
	return c.WaitUntilClusterDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterDeletedWithContext calls the underlying client method with a request option that
// will count DescribeClusters requests.
func (c *Redshift) WaitUntilClusterDeletedWithContext(ctx aws.Context, input *redshift.DescribeClustersInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeClusters")))
	return c.svc.WaitUntilClusterDeletedWithContext(ctx, input, opts...)
}

// WaitUntilClusterRestored calls WaitUntilClusterRestoredWithContext with aws.BackgroundContext().
func (c *Redshift) WaitUntilClusterRestored(input *redshift.DescribeClustersInput) error {
	return c.WaitUntilClusterRestoredWithContext(aws.BackgroundContext(), input)
}

// WaitUntilClusterRestoredWithContext calls the underlying client method with a request option that
// will count DescribeClusters requests.
func (c *Redshift) WaitUntilClusterRestoredWithContext(ctx aws.Context, input *redshift.DescribeClustersInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeClusters")))
	return c.svc.WaitUntilClusterRestoredWithContext(ctx, input, opts...)
}

// WaitUntilSnapshotAvailable calls WaitUntilSnapshotAvailableWithContext with aws.BackgroundContext().
func (c *Redshift) WaitUntilSnapshotAvailable(input *redshift.DescribeClusterSnapshotsInput) error {
	return c.WaitUntilSnapshotAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilSnapshotAvailableWithContext calls the underlying client method with a request option that
// will count DescribeClusterSnapshots requests.
func (c *Redshift) WaitUntilSnapshotAvailableWithContext(ctx aws.Context, input *redshift.DescribeClusterSnapshotsInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeClusterSnapshots")))
	return c.svc.WaitUntilSnapshotAvailableWithContext(ctx, input, opts...)
}
