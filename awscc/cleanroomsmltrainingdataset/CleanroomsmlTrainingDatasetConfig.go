package cleanroomsmltrainingdataset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CleanroomsmlTrainingDatasetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanroomsml_training_dataset#name CleanroomsmlTrainingDataset#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanroomsml_training_dataset#role_arn CleanroomsmlTrainingDataset#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanroomsml_training_dataset#training_data CleanroomsmlTrainingDataset#training_data}.
	TrainingData interface{} `field:"required" json:"trainingData" yaml:"trainingData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanroomsml_training_dataset#description CleanroomsmlTrainingDataset#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// An arbitrary set of tags (key-value pairs) for this cleanrooms-ml training dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanroomsml_training_dataset#tags CleanroomsmlTrainingDataset#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

