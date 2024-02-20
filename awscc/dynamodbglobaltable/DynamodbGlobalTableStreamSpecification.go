package dynamodbglobaltable


type DynamodbGlobalTableStreamSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_global_table#stream_view_type DynamodbGlobalTable#stream_view_type}.
	StreamViewType *string `field:"required" json:"streamViewType" yaml:"streamViewType"`
}

