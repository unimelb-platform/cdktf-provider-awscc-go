package guarddutyipset


type GuarddutyIpSetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_ip_set#key GuarddutyIpSet#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_ip_set#value GuarddutyIpSet#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

