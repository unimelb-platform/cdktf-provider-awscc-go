package cassandrakeyspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CassandraKeyspaceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Name for Cassandra keyspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_keyspace#keyspace_name CassandraKeyspace#keyspace_name}
	KeyspaceName *string `field:"optional" json:"keyspaceName" yaml:"keyspaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_keyspace#replication_specification CassandraKeyspace#replication_specification}.
	ReplicationSpecification *CassandraKeyspaceReplicationSpecification `field:"optional" json:"replicationSpecification" yaml:"replicationSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cassandra_keyspace#tags CassandraKeyspace#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

