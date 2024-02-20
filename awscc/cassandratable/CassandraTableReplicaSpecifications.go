package cassandratable


type CassandraTableReplicaSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#region CassandraTable#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Represents configuration for auto scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#read_capacity_auto_scaling CassandraTable#read_capacity_auto_scaling}
	ReadCapacityAutoScaling *CassandraTableReplicaSpecificationsReadCapacityAutoScaling `field:"optional" json:"readCapacityAutoScaling" yaml:"readCapacityAutoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#read_capacity_units CassandraTable#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
}

