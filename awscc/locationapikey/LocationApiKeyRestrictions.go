package locationapikey


type LocationApiKeyRestrictions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_api_key#allow_actions LocationApiKey#allow_actions}.
	AllowActions *[]*string `field:"required" json:"allowActions" yaml:"allowActions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_api_key#allow_resources LocationApiKey#allow_resources}.
	AllowResources *[]*string `field:"required" json:"allowResources" yaml:"allowResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/location_api_key#allow_referers LocationApiKey#allow_referers}.
	AllowReferers *[]*string `field:"optional" json:"allowReferers" yaml:"allowReferers"`
}

