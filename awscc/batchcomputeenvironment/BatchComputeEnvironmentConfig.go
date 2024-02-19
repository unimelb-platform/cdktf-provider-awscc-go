package batchcomputeenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BatchComputeEnvironmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#type BatchComputeEnvironment#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#compute_environment_name BatchComputeEnvironment#compute_environment_name}.
	ComputeEnvironmentName *string `field:"optional" json:"computeEnvironmentName" yaml:"computeEnvironmentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#compute_resources BatchComputeEnvironment#compute_resources}.
	ComputeResources *BatchComputeEnvironmentComputeResources `field:"optional" json:"computeResources" yaml:"computeResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#eks_configuration BatchComputeEnvironment#eks_configuration}.
	EksConfiguration *BatchComputeEnvironmentEksConfiguration `field:"optional" json:"eksConfiguration" yaml:"eksConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#replace_compute_environment BatchComputeEnvironment#replace_compute_environment}.
	ReplaceComputeEnvironment interface{} `field:"optional" json:"replaceComputeEnvironment" yaml:"replaceComputeEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#service_role BatchComputeEnvironment#service_role}.
	ServiceRole *string `field:"optional" json:"serviceRole" yaml:"serviceRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#state BatchComputeEnvironment#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
	// A key-value pair to associate with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#tags BatchComputeEnvironment#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#unmanagedv_cpus BatchComputeEnvironment#unmanagedv_cpus}.
	UnmanagedvCpus *float64 `field:"optional" json:"unmanagedvCpus" yaml:"unmanagedvCpus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#update_policy BatchComputeEnvironment#update_policy}.
	UpdatePolicy *BatchComputeEnvironmentUpdatePolicy `field:"optional" json:"updatePolicy" yaml:"updatePolicy"`
}

