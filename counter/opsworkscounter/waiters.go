// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworkscounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/opsworks"
)

// WaitUntilAppExists calls WaitUntilAppExistsWithContext with aws.BackgroundContext().
func (c *OpsWorks) WaitUntilAppExists(input *opsworks.DescribeAppsInput) error {
	return c.WaitUntilAppExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilAppExistsWithContext calls the underlying client method with a request option that
// will count DescribeApps requests.
func (c *OpsWorks) WaitUntilAppExistsWithContext(ctx aws.Context, input *opsworks.DescribeAppsInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeApps")))
	return c.svc.WaitUntilAppExistsWithContext(ctx, input, opts...)
}

// WaitUntilDeploymentSuccessful calls WaitUntilDeploymentSuccessfulWithContext with aws.BackgroundContext().
func (c *OpsWorks) WaitUntilDeploymentSuccessful(input *opsworks.DescribeDeploymentsInput) error {
	return c.WaitUntilDeploymentSuccessfulWithContext(aws.BackgroundContext(), input)
}

// WaitUntilDeploymentSuccessfulWithContext calls the underlying client method with a request option that
// will count DescribeDeployments requests.
func (c *OpsWorks) WaitUntilDeploymentSuccessfulWithContext(ctx aws.Context, input *opsworks.DescribeDeploymentsInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeDeployments")))
	return c.svc.WaitUntilDeploymentSuccessfulWithContext(ctx, input, opts...)
}

// WaitUntilInstanceOnline calls WaitUntilInstanceOnlineWithContext with aws.BackgroundContext().
func (c *OpsWorks) WaitUntilInstanceOnline(input *opsworks.DescribeInstancesInput) error {
	return c.WaitUntilInstanceOnlineWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceOnlineWithContext calls the underlying client method with a request option that
// will count DescribeInstances requests.
func (c *OpsWorks) WaitUntilInstanceOnlineWithContext(ctx aws.Context, input *opsworks.DescribeInstancesInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeInstances")))
	return c.svc.WaitUntilInstanceOnlineWithContext(ctx, input, opts...)
}

// WaitUntilInstanceRegistered calls WaitUntilInstanceRegisteredWithContext with aws.BackgroundContext().
func (c *OpsWorks) WaitUntilInstanceRegistered(input *opsworks.DescribeInstancesInput) error {
	return c.WaitUntilInstanceRegisteredWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceRegisteredWithContext calls the underlying client method with a request option that
// will count DescribeInstances requests.
func (c *OpsWorks) WaitUntilInstanceRegisteredWithContext(ctx aws.Context, input *opsworks.DescribeInstancesInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeInstances")))
	return c.svc.WaitUntilInstanceRegisteredWithContext(ctx, input, opts...)
}

// WaitUntilInstanceStopped calls WaitUntilInstanceStoppedWithContext with aws.BackgroundContext().
func (c *OpsWorks) WaitUntilInstanceStopped(input *opsworks.DescribeInstancesInput) error {
	return c.WaitUntilInstanceStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceStoppedWithContext calls the underlying client method with a request option that
// will count DescribeInstances requests.
func (c *OpsWorks) WaitUntilInstanceStoppedWithContext(ctx aws.Context, input *opsworks.DescribeInstancesInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeInstances")))
	return c.svc.WaitUntilInstanceStoppedWithContext(ctx, input, opts...)
}

// WaitUntilInstanceTerminated calls WaitUntilInstanceTerminatedWithContext with aws.BackgroundContext().
func (c *OpsWorks) WaitUntilInstanceTerminated(input *opsworks.DescribeInstancesInput) error {
	return c.WaitUntilInstanceTerminatedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilInstanceTerminatedWithContext calls the underlying client method with a request option that
// will count DescribeInstances requests.
func (c *OpsWorks) WaitUntilInstanceTerminatedWithContext(ctx aws.Context, input *opsworks.DescribeInstancesInput, opts ...request.WaiterOption) error {
	opts = append(opts, request.WithWaiterRequestOptions(c.incViaRequestOption("DescribeInstances")))
	return c.svc.WaitUntilInstanceTerminatedWithContext(ctx, input, opts...)
}
