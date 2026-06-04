package route53recordset


type Route53RecordSetGeoProximityLocationCoordinates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53_record_set#latitude Route53RecordSet#latitude}.
	Latitude *string `field:"optional" json:"latitude" yaml:"latitude"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/route53_record_set#longitude Route53RecordSet#longitude}.
	Longitude *string `field:"optional" json:"longitude" yaml:"longitude"`
}

