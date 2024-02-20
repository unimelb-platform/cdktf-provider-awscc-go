package arczonalshiftzonalautoshiftconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArczonalshiftZonalAutoshiftConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/arczonalshift_zonal_autoshift_configuration#practice_run_configuration ArczonalshiftZonalAutoshiftConfiguration#practice_run_configuration}.
	PracticeRunConfiguration *ArczonalshiftZonalAutoshiftConfigurationPracticeRunConfiguration `field:"optional" json:"practiceRunConfiguration" yaml:"practiceRunConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/arczonalshift_zonal_autoshift_configuration#resource_identifier ArczonalshiftZonalAutoshiftConfiguration#resource_identifier}.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/arczonalshift_zonal_autoshift_configuration#zonal_autoshift_status ArczonalshiftZonalAutoshiftConfiguration#zonal_autoshift_status}.
	ZonalAutoshiftStatus *string `field:"optional" json:"zonalAutoshiftStatus" yaml:"zonalAutoshiftStatus"`
}

