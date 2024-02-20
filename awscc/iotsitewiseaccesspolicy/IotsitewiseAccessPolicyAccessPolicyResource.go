package iotsitewiseaccesspolicy


type IotsitewiseAccessPolicyAccessPolicyResource struct {
	// A portal resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_access_policy#portal IotsitewiseAccessPolicy#portal}
	Portal *IotsitewiseAccessPolicyAccessPolicyResourcePortal `field:"optional" json:"portal" yaml:"portal"`
	// A project resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_access_policy#project IotsitewiseAccessPolicy#project}
	Project *IotsitewiseAccessPolicyAccessPolicyResourceProject `field:"optional" json:"project" yaml:"project"`
}

