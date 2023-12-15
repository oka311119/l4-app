#!/usr/bin/env node
import "source-map-support/register";
import * as cdk from "aws-cdk-lib";
import { Construct } from "constructs";
import { L4FrontAppStack } from "../lib/l4_front_app-stack";
import * as s3 from "aws-cdk-lib/aws-s3";
import * as s3deploy from "aws-cdk-lib/aws-s3-deployment";
import * as cognito from "aws-cdk-lib/aws-cognito";
import * as cloudfront from "aws-cdk-lib/aws-cloudfront";
import * as origins from "aws-cdk-lib/aws-cloudfront-origins";
import * as lambda from "aws-cdk-lib/aws-lambda";
import { Duration, RemovalPolicy } from "aws-cdk-lib";

const app = new cdk.App();
const stack = new L4FrontAppStack(app, "L4FrontAppStack", {
  /* If you don't specify 'env', app stack will be environment-agnostic.
   * Account/Region-dependent features and context lookups will not work,
   * but a single synthesized template can be deployed anywhere. */
  /* Uncomment the next line to specialize app stack for the AWS Account
   * and Region that are implied by the current CLI configuration. */
  // env: { account: process.env.CDK_DEFAULT_ACCOUNT, region: process.env.CDK_DEFAULT_REGION },
  /* Uncomment the next line if you know exactly what Account and Region you
   * want to deploy the stack to. */
  // env: { account: '123456789012', region: 'us-east-1' },
  /* For more information, see https://docs.aws.amazon.com/cdk/latest/guide/environments.html */
});

// S3バケットの作成
const bucket = new s3.Bucket(stack, "BuildBucket", {
  publicReadAccess: true,
  removalPolicy: RemovalPolicy.DESTROY,
  autoDeleteObjects: true,
});

// Reactアプリのデプロイ
new s3deploy.BucketDeployment(stack, "DeployReactApp", {
  sources: [s3deploy.Source.asset("../../frontend/dist")],
  destinationBucket: bucket,
});

// CloudFrontの作成
const distribution = new cloudfront.Distribution(
  stack,
  "ReactAppDistribution",
  {
    defaultBehavior: {
      origin: new origins.S3Origin(bucket),
      viewerProtocolPolicy: cloudfront.ViewerProtocolPolicy.REDIRECT_TO_HTTPS,
      allowedMethods: cloudfront.AllowedMethods.ALLOW_GET_HEAD,
      cachedMethods: cloudfront.CachedMethods.CACHE_GET_HEAD_OPTIONS,
    },
    defaultRootObject: "index.html",
    errorResponses: [
      {
        httpStatus: 403,
        responseHttpStatus: 200,
        responsePagePath: "/index.html",
        ttl: Duration.minutes(5),
      },
    ],
  }
);

// CloudFrontのURLを出力
new cdk.CfnOutput(stack, "DistributionOutput", {
  value: distribution.distributionDomainName,
});
