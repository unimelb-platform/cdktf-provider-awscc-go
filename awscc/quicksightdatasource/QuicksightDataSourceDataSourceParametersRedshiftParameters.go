package quicksightdatasource


type QuicksightDataSourceDataSourceParametersRedshiftParameters struct {
	// <p>Database.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#database QuicksightDataSource#database}
	Database *string `field:"required" json:"database" yaml:"database"`
	// <p>Cluster ID.
	//
	// This field can be blank if the <code>Host</code> and <code>Port</code> are
	//             provided.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#cluster_id QuicksightDataSource#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// <p>Host. This field can be blank if <code>ClusterId</code> is provided.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#host QuicksightDataSource#host}
	Host *string `field:"optional" json:"host" yaml:"host"`
	// <p>Port. This field can be blank if the <code>ClusterId</code> is provided.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_source#port QuicksightDataSource#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

