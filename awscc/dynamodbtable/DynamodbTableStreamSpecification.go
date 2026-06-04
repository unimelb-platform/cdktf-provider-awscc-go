package dynamodbtable


type DynamodbTableStreamSpecification struct {
	// Creates or updates a resource-based policy document that contains the permissions for DDB resources, such as a table's streams.
	//
	// Resource-based policies let you define access permissions by specifying who has access to each resource, and the actions they are allowed to perform on each resource.
	//  In a CFNshort template, you can provide the policy in JSON or YAML format because CFNshort converts YAML to JSON before submitting it to DDB. For more information about resource-based policies, see [Using resource-based policies for](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/access-control-resource-based.html) and [Resource-based policy examples](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/rbac-examples.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dynamodb_table#resource_policy DynamodbTable#resource_policy}
	ResourcePolicy *DynamodbTableStreamSpecificationResourcePolicy `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// When an item in the table is modified, ``StreamViewType`` determines what information is written to the stream for this table.
	//
	// Valid values for ``StreamViewType`` are:
	//   +  ``KEYS_ONLY`` - Only the key attributes of the modified item are written to the stream.
	//   +  ``NEW_IMAGE`` - The entire item, as it appears after it was modified, is written to the stream.
	//   +  ``OLD_IMAGE`` - The entire item, as it appeared before it was modified, is written to the stream.
	//   +  ``NEW_AND_OLD_IMAGES`` - Both the new and the old item images of the item are written to the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dynamodb_table#stream_view_type DynamodbTable#stream_view_type}
	StreamViewType *string `field:"optional" json:"streamViewType" yaml:"streamViewType"`
}

