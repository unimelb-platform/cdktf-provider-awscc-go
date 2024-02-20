package lambdaeventsourcemapping


type LambdaEventSourceMappingDocumentDbEventSourceConfig struct {
	// The collection name to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#collection_name LambdaEventSourceMapping#collection_name}
	CollectionName *string `field:"optional" json:"collectionName" yaml:"collectionName"`
	// The database name to connect to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#database_name LambdaEventSourceMapping#database_name}
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Include full document in change stream response.
	//
	// The default option will only send the changes made to documents to Lambda. If you want the complete document sent to Lambda, set this to UpdateLookup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_event_source_mapping#full_document LambdaEventSourceMapping#full_document}
	FullDocument *string `field:"optional" json:"fullDocument" yaml:"fullDocument"`
}

