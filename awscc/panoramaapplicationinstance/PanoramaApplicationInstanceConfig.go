package panoramaapplicationinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PanoramaApplicationInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#default_runtime_context_device PanoramaApplicationInstance#default_runtime_context_device}.
	DefaultRuntimeContextDevice *string `field:"required" json:"defaultRuntimeContextDevice" yaml:"defaultRuntimeContextDevice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#manifest_payload PanoramaApplicationInstance#manifest_payload}.
	ManifestPayload *PanoramaApplicationInstanceManifestPayload `field:"required" json:"manifestPayload" yaml:"manifestPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#application_instance_id_to_replace PanoramaApplicationInstance#application_instance_id_to_replace}.
	ApplicationInstanceIdToReplace *string `field:"optional" json:"applicationInstanceIdToReplace" yaml:"applicationInstanceIdToReplace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#description PanoramaApplicationInstance#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#manifest_overrides_payload PanoramaApplicationInstance#manifest_overrides_payload}.
	ManifestOverridesPayload *PanoramaApplicationInstanceManifestOverridesPayload `field:"optional" json:"manifestOverridesPayload" yaml:"manifestOverridesPayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#name PanoramaApplicationInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#runtime_role_arn PanoramaApplicationInstance#runtime_role_arn}.
	RuntimeRoleArn *string `field:"optional" json:"runtimeRoleArn" yaml:"runtimeRoleArn"`
	// List of tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_application_instance#tags PanoramaApplicationInstance#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

