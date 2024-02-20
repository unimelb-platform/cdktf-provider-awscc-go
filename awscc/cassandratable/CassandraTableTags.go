package cassandratable


type CassandraTableTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#key CassandraTable#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_table#value CassandraTable#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

