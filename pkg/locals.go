package pkg

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awslambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsLambda *awslambda.AwsLambda
	Labels    map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awslambda.AwsLambdaStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.AwsLambda = stackInput.Target

	return locals
}
