package healthlakefhirdatastore

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthlakeFhirDatastoreConfig struct {
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
	// The FHIR version. Only R4 version data is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#datastore_type_version HealthlakeFhirDatastore#datastore_type_version}
	DatastoreTypeVersion *string `field:"required" json:"datastoreTypeVersion" yaml:"datastoreTypeVersion"`
	// The user-generated name for the Data Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#datastore_name HealthlakeFhirDatastore#datastore_name}
	DatastoreName *string `field:"optional" json:"datastoreName" yaml:"datastoreName"`
	// The identity provider configuration for the datastore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#identity_provider_configuration HealthlakeFhirDatastore#identity_provider_configuration}
	IdentityProviderConfiguration *HealthlakeFhirDatastoreIdentityProviderConfiguration `field:"optional" json:"identityProviderConfiguration" yaml:"identityProviderConfiguration"`
	// The preloaded data configuration for the Data Store. Only data preloaded from Synthea is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#preload_data_config HealthlakeFhirDatastore#preload_data_config}
	PreloadDataConfig *HealthlakeFhirDatastorePreloadDataConfig `field:"optional" json:"preloadDataConfig" yaml:"preloadDataConfig"`
	// The server-side encryption key configuration for a customer provided encryption key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#sse_configuration HealthlakeFhirDatastore#sse_configuration}
	SseConfiguration *HealthlakeFhirDatastoreSseConfiguration `field:"optional" json:"sseConfiguration" yaml:"sseConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/healthlake_fhir_datastore#tags HealthlakeFhirDatastore#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

