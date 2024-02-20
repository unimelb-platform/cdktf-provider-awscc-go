package cloudfrontkeygroup


type CloudfrontKeyGroupKeyGroupConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_key_group#items CloudfrontKeyGroup#items}.
	Items *[]*string `field:"required" json:"items" yaml:"items"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_key_group#name CloudfrontKeyGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_key_group#comment CloudfrontKeyGroup#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

