package guarddutythreatintelset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GuarddutyThreatIntelSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_threat_intel_set#format GuarddutyThreatIntelSet#format}.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_threat_intel_set#location GuarddutyThreatIntelSet#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_threat_intel_set#activate GuarddutyThreatIntelSet#activate}.
	Activate interface{} `field:"optional" json:"activate" yaml:"activate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_threat_intel_set#detector_id GuarddutyThreatIntelSet#detector_id}.
	DetectorId *string `field:"optional" json:"detectorId" yaml:"detectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_threat_intel_set#name GuarddutyThreatIntelSet#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_threat_intel_set#tags GuarddutyThreatIntelSet#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

