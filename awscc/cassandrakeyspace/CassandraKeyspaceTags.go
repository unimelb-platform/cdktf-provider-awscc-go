package cassandrakeyspace


type CassandraKeyspaceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cassandra_keyspace#key CassandraKeyspace#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cassandra_keyspace#value CassandraKeyspace#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

