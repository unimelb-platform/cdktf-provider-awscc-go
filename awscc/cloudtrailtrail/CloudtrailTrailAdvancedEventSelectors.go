package cloudtrailtrail


type CloudtrailTrailAdvancedEventSelectors struct {
	// Contains all selector statements in an advanced event selector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_trail#field_selectors CloudtrailTrail#field_selectors}
	FieldSelectors interface{} `field:"required" json:"fieldSelectors" yaml:"fieldSelectors"`
	// An optional, descriptive name for an advanced event selector, such as "Log data events for only two S3 buckets".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudtrail_trail#name CloudtrailTrail#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

