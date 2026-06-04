module github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccapprunnerservice

go 1.24.4

require (
	github.com/aws/constructs-go/constructs/v10 v10.5.1
	github.com/aws/jsii-runtime-go v1.126.0
	github.com/hashicorp/terraform-cdk-go/cdktf v0.21.0
	github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii v0.0.0-00010101000000-000000000000
)

require github.com/Masterminds/semver/v3 v3.4.0 // indirect

replace github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii => ../jsii
