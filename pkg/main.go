package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/aws/awslambda"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awslambda.AwsLambdaStackInput) error {
	locals := initializeLocals(ctx, stackInput)
	awsCredential := stackInput.AwsCredential

	//create aws provider using the credentials from the input
	awsProvider, err := aws.NewProvider(ctx,
		"classic-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	createdIamRole, err := iamRole(ctx, locals, awsProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create iam role")
	}

	if stackInput.Target.Spec.CloudwatchLogGroup != nil {
		_, err := cloudwatchLogGroup(ctx, locals, awsProvider)
		if err != nil {
			return errors.Wrap(err, "failed to create cloud watch log group")
		}
	}

	createdLambdaFunction, err := lambdaFunction(ctx, locals, awsProvider, createdIamRole)
	if err != nil {
		return errors.Wrap(err, "failed to create lambda function")
	}

	err = invokeFunctionPermissions(ctx, locals, awsProvider, createdLambdaFunction)
	if err != nil {
		return errors.Wrap(err, "failed to create invoke function permissions")
	}

	ctx.Export("lambda-function-arn", createdLambdaFunction.Arn)
	ctx.Export("lambda-function-name", createdLambdaFunction.Name)
	ctx.Export("iam-role-name", createdIamRole.Name)

	return nil
}
