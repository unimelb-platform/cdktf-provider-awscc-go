package dynamodbtable


type DynamodbTableGlobalSecondaryIndexesProvisionedThroughput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#read_capacity_units DynamodbTable#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"required" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#write_capacity_units DynamodbTable#write_capacity_units}.
	WriteCapacityUnits *float64 `field:"required" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

