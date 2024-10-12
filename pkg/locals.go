package pkg

import (
	awslambdav1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awslambda/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	AwsLambda *awslambdav1.AwsLambda
	Labels    map[string]string
}

func initializeLocals(ctx *pulumi.Context, stackInput *awslambdav1.AwsLambdaStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.AwsLambda = stackInput.Target

	return locals
}
