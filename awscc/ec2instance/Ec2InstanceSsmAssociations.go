package ec2instance


type Ec2InstanceSsmAssociations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#document_name Ec2Instance#document_name}.
	DocumentName *string `field:"required" json:"documentName" yaml:"documentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_instance#association_parameters Ec2Instance#association_parameters}.
	AssociationParameters interface{} `field:"optional" json:"associationParameters" yaml:"associationParameters"`
}

