package appsyncsourceapiassociation


type AppsyncSourceApiAssociationSourceApiAssociationConfig struct {
	// Configuration of the merged behavior for the association.
	//
	// For example when it could be auto or has to be manual.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_source_api_association#merge_type AppsyncSourceApiAssociation#merge_type}
	MergeType *string `field:"optional" json:"mergeType" yaml:"mergeType"`
}

