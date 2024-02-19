package healthlakefhirdatastore


type HealthlakeFhirDatastorePreloadDataConfig struct {
	// The type of preloaded data. Only Synthea preloaded data is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#preload_data_type HealthlakeFhirDatastore#preload_data_type}
	PreloadDataType *string `field:"required" json:"preloadDataType" yaml:"preloadDataType"`
}

