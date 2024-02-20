package cassandratable


type CassandraTableBillingModeProvisionedThroughput struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#read_capacity_units CassandraTable#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"required" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#write_capacity_units CassandraTable#write_capacity_units}.
	WriteCapacityUnits *float64 `field:"required" json:"writeCapacityUnits" yaml:"writeCapacityUnits"`
}

