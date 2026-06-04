package datazonedomainunit

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneDomainUnitConfig struct {
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
	// The ID of the domain where you want to create a domain unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_domain_unit#domain_identifier DatazoneDomainUnit#domain_identifier}
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// The name of the domain unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_domain_unit#name DatazoneDomainUnit#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the parent domain unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_domain_unit#parent_domain_unit_identifier DatazoneDomainUnit#parent_domain_unit_identifier}
	ParentDomainUnitIdentifier *string `field:"required" json:"parentDomainUnitIdentifier" yaml:"parentDomainUnitIdentifier"`
	// The description of the domain unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_domain_unit#description DatazoneDomainUnit#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

