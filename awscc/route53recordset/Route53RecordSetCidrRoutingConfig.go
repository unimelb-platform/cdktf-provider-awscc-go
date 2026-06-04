package route53recordset


type Route53RecordSetCidrRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53_record_set#collection_id Route53RecordSet#collection_id}.
	CollectionId *string `field:"optional" json:"collectionId" yaml:"collectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53_record_set#location_name Route53RecordSet#location_name}.
	LocationName *string `field:"optional" json:"locationName" yaml:"locationName"`
}

