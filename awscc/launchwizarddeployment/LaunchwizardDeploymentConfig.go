package launchwizarddeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LaunchwizardDeploymentConfig struct {
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
	// Workload deployment pattern name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/launchwizard_deployment#deployment_pattern_name LaunchwizardDeployment#deployment_pattern_name}
	DeploymentPatternName *string `field:"required" json:"deploymentPatternName" yaml:"deploymentPatternName"`
	// Name of LaunchWizard deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/launchwizard_deployment#name LaunchwizardDeployment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Workload Name for LaunchWizard deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/launchwizard_deployment#workload_name LaunchwizardDeployment#workload_name}
	WorkloadName *string `field:"required" json:"workloadName" yaml:"workloadName"`
	// LaunchWizard deployment specifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/launchwizard_deployment#specifications LaunchwizardDeployment#specifications}
	Specifications *map[string]*string `field:"optional" json:"specifications" yaml:"specifications"`
	// Tags for LaunchWizard deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/launchwizard_deployment#tags LaunchwizardDeployment#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

