package kendraindex


type KendraIndexCapacityUnits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_index#query_capacity_units KendraIndex#query_capacity_units}.
	QueryCapacityUnits *float64 `field:"required" json:"queryCapacityUnits" yaml:"queryCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_index#storage_capacity_units KendraIndex#storage_capacity_units}.
	StorageCapacityUnits *float64 `field:"required" json:"storageCapacityUnits" yaml:"storageCapacityUnits"`
}

