package dynamodbtable


type DynamodbTableContributorInsightsSpecification struct {
	// Indicates whether CloudWatch Contributor Insights are to be enabled (true) or disabled (false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dynamodb_table#enabled DynamodbTable#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

