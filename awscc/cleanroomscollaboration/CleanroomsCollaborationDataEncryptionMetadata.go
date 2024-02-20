package cleanroomscollaboration


type CleanroomsCollaborationDataEncryptionMetadata struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#allow_cleartext CleanroomsCollaboration#allow_cleartext}.
	AllowCleartext interface{} `field:"required" json:"allowCleartext" yaml:"allowCleartext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#allow_duplicates CleanroomsCollaboration#allow_duplicates}.
	AllowDuplicates interface{} `field:"required" json:"allowDuplicates" yaml:"allowDuplicates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#allow_joins_on_columns_with_different_names CleanroomsCollaboration#allow_joins_on_columns_with_different_names}.
	AllowJoinsOnColumnsWithDifferentNames interface{} `field:"required" json:"allowJoinsOnColumnsWithDifferentNames" yaml:"allowJoinsOnColumnsWithDifferentNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_collaboration#preserve_nulls CleanroomsCollaboration#preserve_nulls}.
	PreserveNulls interface{} `field:"required" json:"preserveNulls" yaml:"preserveNulls"`
}

