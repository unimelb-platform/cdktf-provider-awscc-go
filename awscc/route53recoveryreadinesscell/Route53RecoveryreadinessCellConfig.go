package route53recoveryreadinesscell

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Route53RecoveryreadinessCellConfig struct {
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
	// The name of the cell to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoveryreadiness_cell#cell_name Route53RecoveryreadinessCell#cell_name}
	CellName *string `field:"optional" json:"cellName" yaml:"cellName"`
	// A list of cell Amazon Resource Names (ARNs) contained within this cell, for use in nested cells.
	//
	// For example, Availability Zones within specific Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoveryreadiness_cell#cells Route53RecoveryreadinessCell#cells}
	Cells *[]*string `field:"optional" json:"cells" yaml:"cells"`
	// A collection of tags associated with a resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/route53recoveryreadiness_cell#tags Route53RecoveryreadinessCell#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

