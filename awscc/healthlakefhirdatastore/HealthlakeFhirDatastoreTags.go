package healthlakefhirdatastore


type HealthlakeFhirDatastoreTags struct {
	// The key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#key HealthlakeFhirDatastore#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#value HealthlakeFhirDatastore#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

