import { Construct } from "constructs";
import { CloudBackend, NamedCloudWorkspace, TerraformStack } from "cdktf";
import { LambdaFunction } from "@cdktf/provider-aws/lib/lambda-function";
import { IamRole } from "@cdktf/provider-aws/lib/iam-role";
import { AwsProvider } from "@cdktf/provider-aws/lib/provider";
import { S3Bucket } from "@cdktf/provider-aws/lib/s3-bucket";
import { IamRolePolicyAttachment } from "@cdktf/provider-aws/lib/iam-role-policy-attachment";
import { LambdaFunctionUrl } from "@cdktf/provider-aws/lib/lambda-function-url";

export class CoreStack extends TerraformStack {
  constructor(scope: Construct) {
    super(scope, "core");

    new CloudBackend(this, {
      hostname: "app.terraform.io",
      organization: "opentutor",
      workspaces: new NamedCloudWorkspace(`open-tutor-core`),
    });

    new AwsProvider(this, "AWS", {
      region: "us-west-2",
    });

    const bucket = new S3Bucket(this, "bucket", {
      bucketPrefix: "opentutor",
    });

    const helloWorldRole = new IamRole(this, "hello-world-role", {
      name: "hello-world-role",
      assumeRolePolicy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [
          {
            Effect: "Allow",
            Principal: {
              Service: "lambda.amazonaws.com",
            },
            Action: "sts:AssumeRole",
          },
        ],
      }),
    });
    new IamRolePolicyAttachment(this, "att1", {
      role: helloWorldRole.name,
      policyArn: "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole",
    });

    new LambdaFunction(this, "hello-world", {
      functionName: "HelloWorld",
      runtime: "nodejs20.x",
      architectures: ["arm64"],
      role: helloWorldRole.arn,
      s3Bucket: bucket.bucket,
      s3Key: "index.zip",
      handler: "index.handler",
    });
    new LambdaFunctionUrl(this, "url", {
      authorizationType: "NONE",
      functionName: "HelloWorld",
    });
  }
}
