# Description
this is a go project using AWS demo, using go to write aws middleware and login logic.

# Welcome to your CDK Go project!

This is a blank project for CDK development with Go.

The `cdk.json` file tells the CDK toolkit how to execute your app.

## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `go test`         run unit tests


## Refactor
change dynamoDBClient to an interface
by this way, you could change the dynamoDB as any database you want