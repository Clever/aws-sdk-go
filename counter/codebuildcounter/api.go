// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codebuildcounter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codebuild"
)

// BatchDeleteBuildsRequest is a passthrough to the underlying BatchDeleteBuildsRequest.
// It will increment the count of requests made to BatchDeleteBuilds.
func (c *CodeBuild) BatchDeleteBuildsRequest(input *codebuild.BatchDeleteBuildsInput) (req *request.Request, output *codebuild.BatchDeleteBuildsOutput) {
	c.inc("BatchDeleteBuilds")
	return c.svc.BatchDeleteBuildsRequest(input)
}

// BatchDeleteBuilds is a passthrough to the underlying BatchDeleteBuilds method.
// It will increment the count of requests made to BatchDeleteBuilds.
func (c *CodeBuild) BatchDeleteBuilds(input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
	c.inc("BatchDeleteBuilds")
	return c.svc.BatchDeleteBuilds(input)
}

// BatchDeleteBuildsWithContext is a passthrough to the underlying BatchDeleteBuildsWithContext method.
// It will increment the count of requests made to BatchDeleteBuilds.
func (c *CodeBuild) BatchDeleteBuildsWithContext(ctx aws.Context, input *codebuild.BatchDeleteBuildsInput, opts ...request.Option) (*codebuild.BatchDeleteBuildsOutput, error) {
	c.inc("BatchDeleteBuilds")
	return c.svc.BatchDeleteBuildsWithContext(ctx, input, opts...)
}

// BatchGetBuildsRequest is a passthrough to the underlying BatchGetBuildsRequest.
// It will increment the count of requests made to BatchGetBuilds.
func (c *CodeBuild) BatchGetBuildsRequest(input *codebuild.BatchGetBuildsInput) (req *request.Request, output *codebuild.BatchGetBuildsOutput) {
	c.inc("BatchGetBuilds")
	return c.svc.BatchGetBuildsRequest(input)
}

// BatchGetBuilds is a passthrough to the underlying BatchGetBuilds method.
// It will increment the count of requests made to BatchGetBuilds.
func (c *CodeBuild) BatchGetBuilds(input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	c.inc("BatchGetBuilds")
	return c.svc.BatchGetBuilds(input)
}

// BatchGetBuildsWithContext is a passthrough to the underlying BatchGetBuildsWithContext method.
// It will increment the count of requests made to BatchGetBuilds.
func (c *CodeBuild) BatchGetBuildsWithContext(ctx aws.Context, input *codebuild.BatchGetBuildsInput, opts ...request.Option) (*codebuild.BatchGetBuildsOutput, error) {
	c.inc("BatchGetBuilds")
	return c.svc.BatchGetBuildsWithContext(ctx, input, opts...)
}

// BatchGetProjectsRequest is a passthrough to the underlying BatchGetProjectsRequest.
// It will increment the count of requests made to BatchGetProjects.
func (c *CodeBuild) BatchGetProjectsRequest(input *codebuild.BatchGetProjectsInput) (req *request.Request, output *codebuild.BatchGetProjectsOutput) {
	c.inc("BatchGetProjects")
	return c.svc.BatchGetProjectsRequest(input)
}

// BatchGetProjects is a passthrough to the underlying BatchGetProjects method.
// It will increment the count of requests made to BatchGetProjects.
func (c *CodeBuild) BatchGetProjects(input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error) {
	c.inc("BatchGetProjects")
	return c.svc.BatchGetProjects(input)
}

// BatchGetProjectsWithContext is a passthrough to the underlying BatchGetProjectsWithContext method.
// It will increment the count of requests made to BatchGetProjects.
func (c *CodeBuild) BatchGetProjectsWithContext(ctx aws.Context, input *codebuild.BatchGetProjectsInput, opts ...request.Option) (*codebuild.BatchGetProjectsOutput, error) {
	c.inc("BatchGetProjects")
	return c.svc.BatchGetProjectsWithContext(ctx, input, opts...)
}

// CreateProjectRequest is a passthrough to the underlying CreateProjectRequest.
// It will increment the count of requests made to CreateProject.
func (c *CodeBuild) CreateProjectRequest(input *codebuild.CreateProjectInput) (req *request.Request, output *codebuild.CreateProjectOutput) {
	c.inc("CreateProject")
	return c.svc.CreateProjectRequest(input)
}

// CreateProject is a passthrough to the underlying CreateProject method.
// It will increment the count of requests made to CreateProject.
func (c *CodeBuild) CreateProject(input *codebuild.CreateProjectInput) (*codebuild.CreateProjectOutput, error) {
	c.inc("CreateProject")
	return c.svc.CreateProject(input)
}

// CreateProjectWithContext is a passthrough to the underlying CreateProjectWithContext method.
// It will increment the count of requests made to CreateProject.
func (c *CodeBuild) CreateProjectWithContext(ctx aws.Context, input *codebuild.CreateProjectInput, opts ...request.Option) (*codebuild.CreateProjectOutput, error) {
	c.inc("CreateProject")
	return c.svc.CreateProjectWithContext(ctx, input, opts...)
}

// CreateWebhookRequest is a passthrough to the underlying CreateWebhookRequest.
// It will increment the count of requests made to CreateWebhook.
func (c *CodeBuild) CreateWebhookRequest(input *codebuild.CreateWebhookInput) (req *request.Request, output *codebuild.CreateWebhookOutput) {
	c.inc("CreateWebhook")
	return c.svc.CreateWebhookRequest(input)
}

// CreateWebhook is a passthrough to the underlying CreateWebhook method.
// It will increment the count of requests made to CreateWebhook.
func (c *CodeBuild) CreateWebhook(input *codebuild.CreateWebhookInput) (*codebuild.CreateWebhookOutput, error) {
	c.inc("CreateWebhook")
	return c.svc.CreateWebhook(input)
}

// CreateWebhookWithContext is a passthrough to the underlying CreateWebhookWithContext method.
// It will increment the count of requests made to CreateWebhook.
func (c *CodeBuild) CreateWebhookWithContext(ctx aws.Context, input *codebuild.CreateWebhookInput, opts ...request.Option) (*codebuild.CreateWebhookOutput, error) {
	c.inc("CreateWebhook")
	return c.svc.CreateWebhookWithContext(ctx, input, opts...)
}

// DeleteProjectRequest is a passthrough to the underlying DeleteProjectRequest.
// It will increment the count of requests made to DeleteProject.
func (c *CodeBuild) DeleteProjectRequest(input *codebuild.DeleteProjectInput) (req *request.Request, output *codebuild.DeleteProjectOutput) {
	c.inc("DeleteProject")
	return c.svc.DeleteProjectRequest(input)
}

// DeleteProject is a passthrough to the underlying DeleteProject method.
// It will increment the count of requests made to DeleteProject.
func (c *CodeBuild) DeleteProject(input *codebuild.DeleteProjectInput) (*codebuild.DeleteProjectOutput, error) {
	c.inc("DeleteProject")
	return c.svc.DeleteProject(input)
}

// DeleteProjectWithContext is a passthrough to the underlying DeleteProjectWithContext method.
// It will increment the count of requests made to DeleteProject.
func (c *CodeBuild) DeleteProjectWithContext(ctx aws.Context, input *codebuild.DeleteProjectInput, opts ...request.Option) (*codebuild.DeleteProjectOutput, error) {
	c.inc("DeleteProject")
	return c.svc.DeleteProjectWithContext(ctx, input, opts...)
}

// DeleteWebhookRequest is a passthrough to the underlying DeleteWebhookRequest.
// It will increment the count of requests made to DeleteWebhook.
func (c *CodeBuild) DeleteWebhookRequest(input *codebuild.DeleteWebhookInput) (req *request.Request, output *codebuild.DeleteWebhookOutput) {
	c.inc("DeleteWebhook")
	return c.svc.DeleteWebhookRequest(input)
}

// DeleteWebhook is a passthrough to the underlying DeleteWebhook method.
// It will increment the count of requests made to DeleteWebhook.
func (c *CodeBuild) DeleteWebhook(input *codebuild.DeleteWebhookInput) (*codebuild.DeleteWebhookOutput, error) {
	c.inc("DeleteWebhook")
	return c.svc.DeleteWebhook(input)
}

// DeleteWebhookWithContext is a passthrough to the underlying DeleteWebhookWithContext method.
// It will increment the count of requests made to DeleteWebhook.
func (c *CodeBuild) DeleteWebhookWithContext(ctx aws.Context, input *codebuild.DeleteWebhookInput, opts ...request.Option) (*codebuild.DeleteWebhookOutput, error) {
	c.inc("DeleteWebhook")
	return c.svc.DeleteWebhookWithContext(ctx, input, opts...)
}

// InvalidateProjectCacheRequest is a passthrough to the underlying InvalidateProjectCacheRequest.
// It will increment the count of requests made to InvalidateProjectCache.
func (c *CodeBuild) InvalidateProjectCacheRequest(input *codebuild.InvalidateProjectCacheInput) (req *request.Request, output *codebuild.InvalidateProjectCacheOutput) {
	c.inc("InvalidateProjectCache")
	return c.svc.InvalidateProjectCacheRequest(input)
}

// InvalidateProjectCache is a passthrough to the underlying InvalidateProjectCache method.
// It will increment the count of requests made to InvalidateProjectCache.
func (c *CodeBuild) InvalidateProjectCache(input *codebuild.InvalidateProjectCacheInput) (*codebuild.InvalidateProjectCacheOutput, error) {
	c.inc("InvalidateProjectCache")
	return c.svc.InvalidateProjectCache(input)
}

// InvalidateProjectCacheWithContext is a passthrough to the underlying InvalidateProjectCacheWithContext method.
// It will increment the count of requests made to InvalidateProjectCache.
func (c *CodeBuild) InvalidateProjectCacheWithContext(ctx aws.Context, input *codebuild.InvalidateProjectCacheInput, opts ...request.Option) (*codebuild.InvalidateProjectCacheOutput, error) {
	c.inc("InvalidateProjectCache")
	return c.svc.InvalidateProjectCacheWithContext(ctx, input, opts...)
}

// ListBuildsRequest is a passthrough to the underlying ListBuildsRequest.
// It will increment the count of requests made to ListBuilds.
func (c *CodeBuild) ListBuildsRequest(input *codebuild.ListBuildsInput) (req *request.Request, output *codebuild.ListBuildsOutput) {
	c.inc("ListBuilds")
	return c.svc.ListBuildsRequest(input)
}

// ListBuilds is a passthrough to the underlying ListBuilds method.
// It will increment the count of requests made to ListBuilds.
func (c *CodeBuild) ListBuilds(input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	c.inc("ListBuilds")
	return c.svc.ListBuilds(input)
}

// ListBuildsWithContext is a passthrough to the underlying ListBuildsWithContext method.
// It will increment the count of requests made to ListBuilds.
func (c *CodeBuild) ListBuildsWithContext(ctx aws.Context, input *codebuild.ListBuildsInput, opts ...request.Option) (*codebuild.ListBuildsOutput, error) {
	c.inc("ListBuilds")
	return c.svc.ListBuildsWithContext(ctx, input, opts...)
}

// ListBuildsForProjectRequest is a passthrough to the underlying ListBuildsForProjectRequest.
// It will increment the count of requests made to ListBuildsForProject.
func (c *CodeBuild) ListBuildsForProjectRequest(input *codebuild.ListBuildsForProjectInput) (req *request.Request, output *codebuild.ListBuildsForProjectOutput) {
	c.inc("ListBuildsForProject")
	return c.svc.ListBuildsForProjectRequest(input)
}

// ListBuildsForProject is a passthrough to the underlying ListBuildsForProject method.
// It will increment the count of requests made to ListBuildsForProject.
func (c *CodeBuild) ListBuildsForProject(input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error) {
	c.inc("ListBuildsForProject")
	return c.svc.ListBuildsForProject(input)
}

// ListBuildsForProjectWithContext is a passthrough to the underlying ListBuildsForProjectWithContext method.
// It will increment the count of requests made to ListBuildsForProject.
func (c *CodeBuild) ListBuildsForProjectWithContext(ctx aws.Context, input *codebuild.ListBuildsForProjectInput, opts ...request.Option) (*codebuild.ListBuildsForProjectOutput, error) {
	c.inc("ListBuildsForProject")
	return c.svc.ListBuildsForProjectWithContext(ctx, input, opts...)
}

// ListCuratedEnvironmentImagesRequest is a passthrough to the underlying ListCuratedEnvironmentImagesRequest.
// It will increment the count of requests made to ListCuratedEnvironmentImages.
func (c *CodeBuild) ListCuratedEnvironmentImagesRequest(input *codebuild.ListCuratedEnvironmentImagesInput) (req *request.Request, output *codebuild.ListCuratedEnvironmentImagesOutput) {
	c.inc("ListCuratedEnvironmentImages")
	return c.svc.ListCuratedEnvironmentImagesRequest(input)
}

// ListCuratedEnvironmentImages is a passthrough to the underlying ListCuratedEnvironmentImages method.
// It will increment the count of requests made to ListCuratedEnvironmentImages.
func (c *CodeBuild) ListCuratedEnvironmentImages(input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	c.inc("ListCuratedEnvironmentImages")
	return c.svc.ListCuratedEnvironmentImages(input)
}

// ListCuratedEnvironmentImagesWithContext is a passthrough to the underlying ListCuratedEnvironmentImagesWithContext method.
// It will increment the count of requests made to ListCuratedEnvironmentImages.
func (c *CodeBuild) ListCuratedEnvironmentImagesWithContext(ctx aws.Context, input *codebuild.ListCuratedEnvironmentImagesInput, opts ...request.Option) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	c.inc("ListCuratedEnvironmentImages")
	return c.svc.ListCuratedEnvironmentImagesWithContext(ctx, input, opts...)
}

// ListProjectsRequest is a passthrough to the underlying ListProjectsRequest.
// It will increment the count of requests made to ListProjects.
func (c *CodeBuild) ListProjectsRequest(input *codebuild.ListProjectsInput) (req *request.Request, output *codebuild.ListProjectsOutput) {
	c.inc("ListProjects")
	return c.svc.ListProjectsRequest(input)
}

// ListProjects is a passthrough to the underlying ListProjects method.
// It will increment the count of requests made to ListProjects.
func (c *CodeBuild) ListProjects(input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error) {
	c.inc("ListProjects")
	return c.svc.ListProjects(input)
}

// ListProjectsWithContext is a passthrough to the underlying ListProjectsWithContext method.
// It will increment the count of requests made to ListProjects.
func (c *CodeBuild) ListProjectsWithContext(ctx aws.Context, input *codebuild.ListProjectsInput, opts ...request.Option) (*codebuild.ListProjectsOutput, error) {
	c.inc("ListProjects")
	return c.svc.ListProjectsWithContext(ctx, input, opts...)
}

// StartBuildRequest is a passthrough to the underlying StartBuildRequest.
// It will increment the count of requests made to StartBuild.
func (c *CodeBuild) StartBuildRequest(input *codebuild.StartBuildInput) (req *request.Request, output *codebuild.StartBuildOutput) {
	c.inc("StartBuild")
	return c.svc.StartBuildRequest(input)
}

// StartBuild is a passthrough to the underlying StartBuild method.
// It will increment the count of requests made to StartBuild.
func (c *CodeBuild) StartBuild(input *codebuild.StartBuildInput) (*codebuild.StartBuildOutput, error) {
	c.inc("StartBuild")
	return c.svc.StartBuild(input)
}

// StartBuildWithContext is a passthrough to the underlying StartBuildWithContext method.
// It will increment the count of requests made to StartBuild.
func (c *CodeBuild) StartBuildWithContext(ctx aws.Context, input *codebuild.StartBuildInput, opts ...request.Option) (*codebuild.StartBuildOutput, error) {
	c.inc("StartBuild")
	return c.svc.StartBuildWithContext(ctx, input, opts...)
}

// StopBuildRequest is a passthrough to the underlying StopBuildRequest.
// It will increment the count of requests made to StopBuild.
func (c *CodeBuild) StopBuildRequest(input *codebuild.StopBuildInput) (req *request.Request, output *codebuild.StopBuildOutput) {
	c.inc("StopBuild")
	return c.svc.StopBuildRequest(input)
}

// StopBuild is a passthrough to the underlying StopBuild method.
// It will increment the count of requests made to StopBuild.
func (c *CodeBuild) StopBuild(input *codebuild.StopBuildInput) (*codebuild.StopBuildOutput, error) {
	c.inc("StopBuild")
	return c.svc.StopBuild(input)
}

// StopBuildWithContext is a passthrough to the underlying StopBuildWithContext method.
// It will increment the count of requests made to StopBuild.
func (c *CodeBuild) StopBuildWithContext(ctx aws.Context, input *codebuild.StopBuildInput, opts ...request.Option) (*codebuild.StopBuildOutput, error) {
	c.inc("StopBuild")
	return c.svc.StopBuildWithContext(ctx, input, opts...)
}

// UpdateProjectRequest is a passthrough to the underlying UpdateProjectRequest.
// It will increment the count of requests made to UpdateProject.
func (c *CodeBuild) UpdateProjectRequest(input *codebuild.UpdateProjectInput) (req *request.Request, output *codebuild.UpdateProjectOutput) {
	c.inc("UpdateProject")
	return c.svc.UpdateProjectRequest(input)
}

// UpdateProject is a passthrough to the underlying UpdateProject method.
// It will increment the count of requests made to UpdateProject.
func (c *CodeBuild) UpdateProject(input *codebuild.UpdateProjectInput) (*codebuild.UpdateProjectOutput, error) {
	c.inc("UpdateProject")
	return c.svc.UpdateProject(input)
}

// UpdateProjectWithContext is a passthrough to the underlying UpdateProjectWithContext method.
// It will increment the count of requests made to UpdateProject.
func (c *CodeBuild) UpdateProjectWithContext(ctx aws.Context, input *codebuild.UpdateProjectInput, opts ...request.Option) (*codebuild.UpdateProjectOutput, error) {
	c.inc("UpdateProject")
	return c.svc.UpdateProjectWithContext(ctx, input, opts...)
}
