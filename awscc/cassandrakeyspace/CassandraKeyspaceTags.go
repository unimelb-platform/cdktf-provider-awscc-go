package cassandrakeyspace


type CassandraKeyspaceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_keyspace#key CassandraKeyspace#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_keyspace#value CassandraKeyspace#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

