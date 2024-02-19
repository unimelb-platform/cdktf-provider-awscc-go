package quicksightdataset


type QuicksightDataSetLogicalTableMapSource struct {
	// <p>The Amazon Resource Name (ARN) for the dataset.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#data_set_arn QuicksightDataSet#data_set_arn}
	DataSetArn *string `field:"optional" json:"dataSetArn" yaml:"dataSetArn"`
	// <p>Join instruction.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#join_instruction QuicksightDataSet#join_instruction}
	JoinInstruction *QuicksightDataSetLogicalTableMapSourceJoinInstruction `field:"optional" json:"joinInstruction" yaml:"joinInstruction"`
	// <p>Physical table ID.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/quicksight_data_set#physical_table_id QuicksightDataSet#physical_table_id}
	PhysicalTableId *string `field:"optional" json:"physicalTableId" yaml:"physicalTableId"`
}

