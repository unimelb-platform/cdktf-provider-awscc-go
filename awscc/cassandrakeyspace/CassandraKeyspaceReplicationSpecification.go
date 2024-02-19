package cassandrakeyspace


type CassandraKeyspaceReplicationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_keyspace#region_list CassandraKeyspace#region_list}.
	RegionList *[]*string `field:"optional" json:"regionList" yaml:"regionList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_keyspace#replication_strategy CassandraKeyspace#replication_strategy}.
	ReplicationStrategy *string `field:"optional" json:"replicationStrategy" yaml:"replicationStrategy"`
}

