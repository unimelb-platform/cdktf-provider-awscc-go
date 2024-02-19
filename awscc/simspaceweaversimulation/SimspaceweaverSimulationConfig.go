package simspaceweaversimulation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SimspaceweaverSimulationConfig struct {
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
	// The name of the simulation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/simspaceweaver_simulation#name SimspaceweaverSimulation#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Role ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/simspaceweaver_simulation#role_arn SimspaceweaverSimulation#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The maximum running time of the simulation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/simspaceweaver_simulation#maximum_duration SimspaceweaverSimulation#maximum_duration}
	MaximumDuration *string `field:"optional" json:"maximumDuration" yaml:"maximumDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/simspaceweaver_simulation#schema_s3_location SimspaceweaverSimulation#schema_s3_location}.
	SchemaS3Location *SimspaceweaverSimulationSchemaS3Location `field:"optional" json:"schemaS3Location" yaml:"schemaS3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/simspaceweaver_simulation#snapshot_s3_location SimspaceweaverSimulation#snapshot_s3_location}.
	SnapshotS3Location *SimspaceweaverSimulationSnapshotS3Location `field:"optional" json:"snapshotS3Location" yaml:"snapshotS3Location"`
}

