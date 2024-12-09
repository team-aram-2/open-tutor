import { DbInstance } from "@cdktf/provider-aws/lib/db-instance/index.js";
import { IamRolePolicyAttachment } from "@cdktf/provider-aws/lib/iam-role-policy-attachment/index.js";
import { IamRole } from "@cdktf/provider-aws/lib/iam-role/index.js";
import { AwsProvider } from "@cdktf/provider-aws/lib/provider/index.js";
import { TerraformStack } from "cdktf";
import { Construct } from "constructs";

export class CoreStack extends TerraformStack {
  constructor(scope: Construct) {
    super(scope, "core-local");

    // new CloudBackend(this, {
    //   hostname: "app.terraform.io",
    //   organization: "opentutor",
    //   workspaces: new NamedCloudWorkspace(`open-tutor-core`),
    // });

    new AwsProvider(this, "AWS", {
      accessKey: "test",
      secretKey: "test",
      region: "us-west-2",
      s3UsePathStyle: true,
      endpoints: [
        {
          iam: "http://localstack:4566",
          s3: "http://localstack:4566",
          lambda: "http://localstack:4566",
          sts: "http://localstack:4566",
          rds: "http://localstack:4566",
        }
      ]
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

    new DbInstance(this, "data-db", {
      allocatedStorage: 10,
      dbName: "opentutor-data",
      engine: "postgres",
      engineVersion: "17.2",
      instanceClass: "db.t4g.medium",
      username: "postgres",
      password: "developer",
      skipFinalSnapshot: true,
    });
  }
}
