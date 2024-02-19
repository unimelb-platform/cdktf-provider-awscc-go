package lightsailinstance


type LightsailInstanceNetworking struct {
	// Ports to the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_instance#ports LightsailInstance#ports}
	Ports interface{} `field:"required" json:"ports" yaml:"ports"`
	// Monthly Transfer of the Instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lightsail_instance#monthly_transfer LightsailInstance#monthly_transfer}
	MonthlyTransfer *LightsailInstanceNetworkingMonthlyTransfer `field:"optional" json:"monthlyTransfer" yaml:"monthlyTransfer"`
}

