// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatcheventscounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
)

// DeleteRuleRequest is a passthrough to the underlying DeleteRuleRequest.
// It will increment the count of requests made to DeleteRule.
func (c *CloudWatchEvents) DeleteRuleRequest(input *cloudwatchevents.DeleteRuleInput) (req *request.Request, output *cloudwatchevents.DeleteRuleOutput) {
	c.inc("DeleteRule")
	return c.svc.DeleteRuleRequest(input)
}

// DeleteRule is a passthrough to the underlying DeleteRule method.
// It will increment the count of requests made to DeleteRule.
func (c *CloudWatchEvents) DeleteRule(input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error) {
	c.inc("DeleteRule")
	return c.svc.DeleteRule(input)
}

// DeleteRuleWithContext is a passthrough to the underlying DeleteRuleWithContext method.
// It will increment the count of requests made to DeleteRule.
func (c *CloudWatchEvents) DeleteRuleWithContext(ctx aws.Context, input *cloudwatchevents.DeleteRuleInput, opts ...request.Option) (*cloudwatchevents.DeleteRuleOutput, error) {
	c.inc("DeleteRule")
	return c.svc.DeleteRuleWithContext(ctx, input, opts...)
}

// DescribeEventBusRequest is a passthrough to the underlying DescribeEventBusRequest.
// It will increment the count of requests made to DescribeEventBus.
func (c *CloudWatchEvents) DescribeEventBusRequest(input *cloudwatchevents.DescribeEventBusInput) (req *request.Request, output *cloudwatchevents.DescribeEventBusOutput) {
	c.inc("DescribeEventBus")
	return c.svc.DescribeEventBusRequest(input)
}

// DescribeEventBus is a passthrough to the underlying DescribeEventBus method.
// It will increment the count of requests made to DescribeEventBus.
func (c *CloudWatchEvents) DescribeEventBus(input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error) {
	c.inc("DescribeEventBus")
	return c.svc.DescribeEventBus(input)
}

// DescribeEventBusWithContext is a passthrough to the underlying DescribeEventBusWithContext method.
// It will increment the count of requests made to DescribeEventBus.
func (c *CloudWatchEvents) DescribeEventBusWithContext(ctx aws.Context, input *cloudwatchevents.DescribeEventBusInput, opts ...request.Option) (*cloudwatchevents.DescribeEventBusOutput, error) {
	c.inc("DescribeEventBus")
	return c.svc.DescribeEventBusWithContext(ctx, input, opts...)
}

// DescribeRuleRequest is a passthrough to the underlying DescribeRuleRequest.
// It will increment the count of requests made to DescribeRule.
func (c *CloudWatchEvents) DescribeRuleRequest(input *cloudwatchevents.DescribeRuleInput) (req *request.Request, output *cloudwatchevents.DescribeRuleOutput) {
	c.inc("DescribeRule")
	return c.svc.DescribeRuleRequest(input)
}

// DescribeRule is a passthrough to the underlying DescribeRule method.
// It will increment the count of requests made to DescribeRule.
func (c *CloudWatchEvents) DescribeRule(input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error) {
	c.inc("DescribeRule")
	return c.svc.DescribeRule(input)
}

// DescribeRuleWithContext is a passthrough to the underlying DescribeRuleWithContext method.
// It will increment the count of requests made to DescribeRule.
func (c *CloudWatchEvents) DescribeRuleWithContext(ctx aws.Context, input *cloudwatchevents.DescribeRuleInput, opts ...request.Option) (*cloudwatchevents.DescribeRuleOutput, error) {
	c.inc("DescribeRule")
	return c.svc.DescribeRuleWithContext(ctx, input, opts...)
}

// DisableRuleRequest is a passthrough to the underlying DisableRuleRequest.
// It will increment the count of requests made to DisableRule.
func (c *CloudWatchEvents) DisableRuleRequest(input *cloudwatchevents.DisableRuleInput) (req *request.Request, output *cloudwatchevents.DisableRuleOutput) {
	c.inc("DisableRule")
	return c.svc.DisableRuleRequest(input)
}

// DisableRule is a passthrough to the underlying DisableRule method.
// It will increment the count of requests made to DisableRule.
func (c *CloudWatchEvents) DisableRule(input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error) {
	c.inc("DisableRule")
	return c.svc.DisableRule(input)
}

// DisableRuleWithContext is a passthrough to the underlying DisableRuleWithContext method.
// It will increment the count of requests made to DisableRule.
func (c *CloudWatchEvents) DisableRuleWithContext(ctx aws.Context, input *cloudwatchevents.DisableRuleInput, opts ...request.Option) (*cloudwatchevents.DisableRuleOutput, error) {
	c.inc("DisableRule")
	return c.svc.DisableRuleWithContext(ctx, input, opts...)
}

// EnableRuleRequest is a passthrough to the underlying EnableRuleRequest.
// It will increment the count of requests made to EnableRule.
func (c *CloudWatchEvents) EnableRuleRequest(input *cloudwatchevents.EnableRuleInput) (req *request.Request, output *cloudwatchevents.EnableRuleOutput) {
	c.inc("EnableRule")
	return c.svc.EnableRuleRequest(input)
}

// EnableRule is a passthrough to the underlying EnableRule method.
// It will increment the count of requests made to EnableRule.
func (c *CloudWatchEvents) EnableRule(input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error) {
	c.inc("EnableRule")
	return c.svc.EnableRule(input)
}

// EnableRuleWithContext is a passthrough to the underlying EnableRuleWithContext method.
// It will increment the count of requests made to EnableRule.
func (c *CloudWatchEvents) EnableRuleWithContext(ctx aws.Context, input *cloudwatchevents.EnableRuleInput, opts ...request.Option) (*cloudwatchevents.EnableRuleOutput, error) {
	c.inc("EnableRule")
	return c.svc.EnableRuleWithContext(ctx, input, opts...)
}

// ListRuleNamesByTargetRequest is a passthrough to the underlying ListRuleNamesByTargetRequest.
// It will increment the count of requests made to ListRuleNamesByTarget.
func (c *CloudWatchEvents) ListRuleNamesByTargetRequest(input *cloudwatchevents.ListRuleNamesByTargetInput) (req *request.Request, output *cloudwatchevents.ListRuleNamesByTargetOutput) {
	c.inc("ListRuleNamesByTarget")
	return c.svc.ListRuleNamesByTargetRequest(input)
}

// ListRuleNamesByTarget is a passthrough to the underlying ListRuleNamesByTarget method.
// It will increment the count of requests made to ListRuleNamesByTarget.
func (c *CloudWatchEvents) ListRuleNamesByTarget(input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	c.inc("ListRuleNamesByTarget")
	return c.svc.ListRuleNamesByTarget(input)
}

// ListRuleNamesByTargetWithContext is a passthrough to the underlying ListRuleNamesByTargetWithContext method.
// It will increment the count of requests made to ListRuleNamesByTarget.
func (c *CloudWatchEvents) ListRuleNamesByTargetWithContext(ctx aws.Context, input *cloudwatchevents.ListRuleNamesByTargetInput, opts ...request.Option) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	c.inc("ListRuleNamesByTarget")
	return c.svc.ListRuleNamesByTargetWithContext(ctx, input, opts...)
}

// ListRulesRequest is a passthrough to the underlying ListRulesRequest.
// It will increment the count of requests made to ListRules.
func (c *CloudWatchEvents) ListRulesRequest(input *cloudwatchevents.ListRulesInput) (req *request.Request, output *cloudwatchevents.ListRulesOutput) {
	c.inc("ListRules")
	return c.svc.ListRulesRequest(input)
}

// ListRules is a passthrough to the underlying ListRules method.
// It will increment the count of requests made to ListRules.
func (c *CloudWatchEvents) ListRules(input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error) {
	c.inc("ListRules")
	return c.svc.ListRules(input)
}

// ListRulesWithContext is a passthrough to the underlying ListRulesWithContext method.
// It will increment the count of requests made to ListRules.
func (c *CloudWatchEvents) ListRulesWithContext(ctx aws.Context, input *cloudwatchevents.ListRulesInput, opts ...request.Option) (*cloudwatchevents.ListRulesOutput, error) {
	c.inc("ListRules")
	return c.svc.ListRulesWithContext(ctx, input, opts...)
}

// ListTargetsByRuleRequest is a passthrough to the underlying ListTargetsByRuleRequest.
// It will increment the count of requests made to ListTargetsByRule.
func (c *CloudWatchEvents) ListTargetsByRuleRequest(input *cloudwatchevents.ListTargetsByRuleInput) (req *request.Request, output *cloudwatchevents.ListTargetsByRuleOutput) {
	c.inc("ListTargetsByRule")
	return c.svc.ListTargetsByRuleRequest(input)
}

// ListTargetsByRule is a passthrough to the underlying ListTargetsByRule method.
// It will increment the count of requests made to ListTargetsByRule.
func (c *CloudWatchEvents) ListTargetsByRule(input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	c.inc("ListTargetsByRule")
	return c.svc.ListTargetsByRule(input)
}

// ListTargetsByRuleWithContext is a passthrough to the underlying ListTargetsByRuleWithContext method.
// It will increment the count of requests made to ListTargetsByRule.
func (c *CloudWatchEvents) ListTargetsByRuleWithContext(ctx aws.Context, input *cloudwatchevents.ListTargetsByRuleInput, opts ...request.Option) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	c.inc("ListTargetsByRule")
	return c.svc.ListTargetsByRuleWithContext(ctx, input, opts...)
}

// PutEventsRequest is a passthrough to the underlying PutEventsRequest.
// It will increment the count of requests made to PutEvents.
func (c *CloudWatchEvents) PutEventsRequest(input *cloudwatchevents.PutEventsInput) (req *request.Request, output *cloudwatchevents.PutEventsOutput) {
	c.inc("PutEvents")
	return c.svc.PutEventsRequest(input)
}

// PutEvents is a passthrough to the underlying PutEvents method.
// It will increment the count of requests made to PutEvents.
func (c *CloudWatchEvents) PutEvents(input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error) {
	c.inc("PutEvents")
	return c.svc.PutEvents(input)
}

// PutEventsWithContext is a passthrough to the underlying PutEventsWithContext method.
// It will increment the count of requests made to PutEvents.
func (c *CloudWatchEvents) PutEventsWithContext(ctx aws.Context, input *cloudwatchevents.PutEventsInput, opts ...request.Option) (*cloudwatchevents.PutEventsOutput, error) {
	c.inc("PutEvents")
	return c.svc.PutEventsWithContext(ctx, input, opts...)
}

// PutPermissionRequest is a passthrough to the underlying PutPermissionRequest.
// It will increment the count of requests made to PutPermission.
func (c *CloudWatchEvents) PutPermissionRequest(input *cloudwatchevents.PutPermissionInput) (req *request.Request, output *cloudwatchevents.PutPermissionOutput) {
	c.inc("PutPermission")
	return c.svc.PutPermissionRequest(input)
}

// PutPermission is a passthrough to the underlying PutPermission method.
// It will increment the count of requests made to PutPermission.
func (c *CloudWatchEvents) PutPermission(input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error) {
	c.inc("PutPermission")
	return c.svc.PutPermission(input)
}

// PutPermissionWithContext is a passthrough to the underlying PutPermissionWithContext method.
// It will increment the count of requests made to PutPermission.
func (c *CloudWatchEvents) PutPermissionWithContext(ctx aws.Context, input *cloudwatchevents.PutPermissionInput, opts ...request.Option) (*cloudwatchevents.PutPermissionOutput, error) {
	c.inc("PutPermission")
	return c.svc.PutPermissionWithContext(ctx, input, opts...)
}

// PutRuleRequest is a passthrough to the underlying PutRuleRequest.
// It will increment the count of requests made to PutRule.
func (c *CloudWatchEvents) PutRuleRequest(input *cloudwatchevents.PutRuleInput) (req *request.Request, output *cloudwatchevents.PutRuleOutput) {
	c.inc("PutRule")
	return c.svc.PutRuleRequest(input)
}

// PutRule is a passthrough to the underlying PutRule method.
// It will increment the count of requests made to PutRule.
func (c *CloudWatchEvents) PutRule(input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error) {
	c.inc("PutRule")
	return c.svc.PutRule(input)
}

// PutRuleWithContext is a passthrough to the underlying PutRuleWithContext method.
// It will increment the count of requests made to PutRule.
func (c *CloudWatchEvents) PutRuleWithContext(ctx aws.Context, input *cloudwatchevents.PutRuleInput, opts ...request.Option) (*cloudwatchevents.PutRuleOutput, error) {
	c.inc("PutRule")
	return c.svc.PutRuleWithContext(ctx, input, opts...)
}

// PutTargetsRequest is a passthrough to the underlying PutTargetsRequest.
// It will increment the count of requests made to PutTargets.
func (c *CloudWatchEvents) PutTargetsRequest(input *cloudwatchevents.PutTargetsInput) (req *request.Request, output *cloudwatchevents.PutTargetsOutput) {
	c.inc("PutTargets")
	return c.svc.PutTargetsRequest(input)
}

// PutTargets is a passthrough to the underlying PutTargets method.
// It will increment the count of requests made to PutTargets.
func (c *CloudWatchEvents) PutTargets(input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error) {
	c.inc("PutTargets")
	return c.svc.PutTargets(input)
}

// PutTargetsWithContext is a passthrough to the underlying PutTargetsWithContext method.
// It will increment the count of requests made to PutTargets.
func (c *CloudWatchEvents) PutTargetsWithContext(ctx aws.Context, input *cloudwatchevents.PutTargetsInput, opts ...request.Option) (*cloudwatchevents.PutTargetsOutput, error) {
	c.inc("PutTargets")
	return c.svc.PutTargetsWithContext(ctx, input, opts...)
}

// RemovePermissionRequest is a passthrough to the underlying RemovePermissionRequest.
// It will increment the count of requests made to RemovePermission.
func (c *CloudWatchEvents) RemovePermissionRequest(input *cloudwatchevents.RemovePermissionInput) (req *request.Request, output *cloudwatchevents.RemovePermissionOutput) {
	c.inc("RemovePermission")
	return c.svc.RemovePermissionRequest(input)
}

// RemovePermission is a passthrough to the underlying RemovePermission method.
// It will increment the count of requests made to RemovePermission.
func (c *CloudWatchEvents) RemovePermission(input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error) {
	c.inc("RemovePermission")
	return c.svc.RemovePermission(input)
}

// RemovePermissionWithContext is a passthrough to the underlying RemovePermissionWithContext method.
// It will increment the count of requests made to RemovePermission.
func (c *CloudWatchEvents) RemovePermissionWithContext(ctx aws.Context, input *cloudwatchevents.RemovePermissionInput, opts ...request.Option) (*cloudwatchevents.RemovePermissionOutput, error) {
	c.inc("RemovePermission")
	return c.svc.RemovePermissionWithContext(ctx, input, opts...)
}

// RemoveTargetsRequest is a passthrough to the underlying RemoveTargetsRequest.
// It will increment the count of requests made to RemoveTargets.
func (c *CloudWatchEvents) RemoveTargetsRequest(input *cloudwatchevents.RemoveTargetsInput) (req *request.Request, output *cloudwatchevents.RemoveTargetsOutput) {
	c.inc("RemoveTargets")
	return c.svc.RemoveTargetsRequest(input)
}

// RemoveTargets is a passthrough to the underlying RemoveTargets method.
// It will increment the count of requests made to RemoveTargets.
func (c *CloudWatchEvents) RemoveTargets(input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error) {
	c.inc("RemoveTargets")
	return c.svc.RemoveTargets(input)
}

// RemoveTargetsWithContext is a passthrough to the underlying RemoveTargetsWithContext method.
// It will increment the count of requests made to RemoveTargets.
func (c *CloudWatchEvents) RemoveTargetsWithContext(ctx aws.Context, input *cloudwatchevents.RemoveTargetsInput, opts ...request.Option) (*cloudwatchevents.RemoveTargetsOutput, error) {
	c.inc("RemoveTargets")
	return c.svc.RemoveTargetsWithContext(ctx, input, opts...)
}

// TestEventPatternRequest is a passthrough to the underlying TestEventPatternRequest.
// It will increment the count of requests made to TestEventPattern.
func (c *CloudWatchEvents) TestEventPatternRequest(input *cloudwatchevents.TestEventPatternInput) (req *request.Request, output *cloudwatchevents.TestEventPatternOutput) {
	c.inc("TestEventPattern")
	return c.svc.TestEventPatternRequest(input)
}

// TestEventPattern is a passthrough to the underlying TestEventPattern method.
// It will increment the count of requests made to TestEventPattern.
func (c *CloudWatchEvents) TestEventPattern(input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error) {
	c.inc("TestEventPattern")
	return c.svc.TestEventPattern(input)
}

// TestEventPatternWithContext is a passthrough to the underlying TestEventPatternWithContext method.
// It will increment the count of requests made to TestEventPattern.
func (c *CloudWatchEvents) TestEventPatternWithContext(ctx aws.Context, input *cloudwatchevents.TestEventPatternInput, opts ...request.Option) (*cloudwatchevents.TestEventPatternOutput, error) {
	c.inc("TestEventPattern")
	return c.svc.TestEventPatternWithContext(ctx, input, opts...)
}
