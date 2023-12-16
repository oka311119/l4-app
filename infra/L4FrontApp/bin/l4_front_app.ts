import "source-map-support/register";
import * as cdk from "aws-cdk-lib";
import { L4FrontAppStack } from "../lib/l4_front_app-stack";
import * as s3 from "aws-cdk-lib/aws-s3";
import * as s3deploy from "aws-cdk-lib/aws-s3-deployment";
import * as cloudfront from "aws-cdk-lib/aws-cloudfront";
import * as origins from "aws-cdk-lib/aws-cloudfront-origins";
import { Duration, RemovalPolicy } from "aws-cdk-lib";

require("dotenv").config();

const app = new cdk.App();
const stack = new L4FrontAppStack(app, "L4FrontAppStack", {
  env: {
    region: process.env.AWS_REGION,
    account: process.env.AWS_ACCOUNT_ID,
  },
});

// S3バケットの作成
const bucket = new s3.Bucket(stack, "BuildBucket", {
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
