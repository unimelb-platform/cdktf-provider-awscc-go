package cleanroomsmltrainingdataset


type CleanroomsmlTrainingDatasetTrainingDataInputConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanroomsml_training_dataset#data_source CleanroomsmlTrainingDataset#data_source}.
	DataSource *CleanroomsmlTrainingDatasetTrainingDataInputConfigDataSource `field:"required" json:"dataSource" yaml:"dataSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanroomsml_training_dataset#schema CleanroomsmlTrainingDataset#schema}.
	Schema interface{} `field:"required" json:"schema" yaml:"schema"`
}

