package dynamodbtable


type DynamodbTableStreamSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#stream_view_type DynamodbTable#stream_view_type}.
	StreamViewType *string `field:"required" json:"streamViewType" yaml:"streamViewType"`
}

