package ec2ec2fleet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2Ec2FleetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#launch_template_configs Ec2Ec2Fleet#launch_template_configs}.
	LaunchTemplateConfigs interface{} `field:"required" json:"launchTemplateConfigs" yaml:"launchTemplateConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#target_capacity_specification Ec2Ec2Fleet#target_capacity_specification}.
	TargetCapacitySpecification *Ec2Ec2FleetTargetCapacitySpecification `field:"required" json:"targetCapacitySpecification" yaml:"targetCapacitySpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#context Ec2Ec2Fleet#context}.
	Context *string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#excess_capacity_termination_policy Ec2Ec2Fleet#excess_capacity_termination_policy}.
	ExcessCapacityTerminationPolicy *string `field:"optional" json:"excessCapacityTerminationPolicy" yaml:"excessCapacityTerminationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#on_demand_options Ec2Ec2Fleet#on_demand_options}.
	OnDemandOptions *Ec2Ec2FleetOnDemandOptions `field:"optional" json:"onDemandOptions" yaml:"onDemandOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#replace_unhealthy_instances Ec2Ec2Fleet#replace_unhealthy_instances}.
	ReplaceUnhealthyInstances interface{} `field:"optional" json:"replaceUnhealthyInstances" yaml:"replaceUnhealthyInstances"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#spot_options Ec2Ec2Fleet#spot_options}.
	SpotOptions *Ec2Ec2FleetSpotOptions `field:"optional" json:"spotOptions" yaml:"spotOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#tag_specifications Ec2Ec2Fleet#tag_specifications}.
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#terminate_instances_with_expiration Ec2Ec2Fleet#terminate_instances_with_expiration}.
	TerminateInstancesWithExpiration interface{} `field:"optional" json:"terminateInstancesWithExpiration" yaml:"terminateInstancesWithExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#type Ec2Ec2Fleet#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#valid_from Ec2Ec2Fleet#valid_from}.
	ValidFrom *string `field:"optional" json:"validFrom" yaml:"validFrom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet#valid_until Ec2Ec2Fleet#valid_until}.
	ValidUntil *string `field:"optional" json:"validUntil" yaml:"validUntil"`
}

