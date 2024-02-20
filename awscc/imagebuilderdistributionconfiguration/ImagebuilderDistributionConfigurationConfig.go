package imagebuilderdistributionconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderDistributionConfigurationConfig struct {
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
	// The distributions of the distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#distributions ImagebuilderDistributionConfiguration#distributions}
	Distributions interface{} `field:"required" json:"distributions" yaml:"distributions"`
	// The name of the distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#name ImagebuilderDistributionConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the distribution configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#description ImagebuilderDistributionConfiguration#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags associated with the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_distribution_configuration#tags ImagebuilderDistributionConfiguration#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

