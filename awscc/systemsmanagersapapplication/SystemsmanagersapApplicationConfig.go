package systemsmanagersapapplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SystemsmanagersapApplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#application_id SystemsmanagersapApplication#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#application_type SystemsmanagersapApplication#application_type}.
	ApplicationType *string `field:"required" json:"applicationType" yaml:"applicationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#credentials SystemsmanagersapApplication#credentials}.
	Credentials interface{} `field:"optional" json:"credentials" yaml:"credentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#instances SystemsmanagersapApplication#instances}.
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#sap_instance_number SystemsmanagersapApplication#sap_instance_number}.
	SapInstanceNumber *string `field:"optional" json:"sapInstanceNumber" yaml:"sapInstanceNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#sid SystemsmanagersapApplication#sid}.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
	// The tags of a SystemsManagerSAP application.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/systemsmanagersap_application#tags SystemsmanagersapApplication#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

