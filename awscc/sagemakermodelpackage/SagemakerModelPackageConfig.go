package sagemakermodelpackage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerModelPackageConfig struct {
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
	// An array of additional Inference Specification objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#additional_inference_specifications SagemakerModelPackage#additional_inference_specifications}
	AdditionalInferenceSpecifications interface{} `field:"optional" json:"additionalInferenceSpecifications" yaml:"additionalInferenceSpecifications"`
	// An array of additional Inference Specification objects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#additional_inference_specifications_to_add SagemakerModelPackage#additional_inference_specifications_to_add}
	AdditionalInferenceSpecificationsToAdd interface{} `field:"optional" json:"additionalInferenceSpecificationsToAdd" yaml:"additionalInferenceSpecificationsToAdd"`
	// A description provided for the model approval.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#approval_description SagemakerModelPackage#approval_description}
	ApprovalDescription *string `field:"optional" json:"approvalDescription" yaml:"approvalDescription"`
	// Whether to certify the model package for listing on AWS Marketplace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#certify_for_marketplace SagemakerModelPackage#certify_for_marketplace}
	CertifyForMarketplace interface{} `field:"optional" json:"certifyForMarketplace" yaml:"certifyForMarketplace"`
	// A unique token that guarantees that the call to this API is idempotent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#client_token SagemakerModelPackage#client_token}
	ClientToken *string `field:"optional" json:"clientToken" yaml:"clientToken"`
	// The metadata properties associated with the model package versions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#customer_metadata_properties SagemakerModelPackage#customer_metadata_properties}
	CustomerMetadataProperties *map[string]*string `field:"optional" json:"customerMetadataProperties" yaml:"customerMetadataProperties"`
	// The machine learning domain of the model package you specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#domain SagemakerModelPackage#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Represents the drift check baselines that can be used when the model monitor is set using the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#drift_check_baselines SagemakerModelPackage#drift_check_baselines}
	DriftCheckBaselines *SagemakerModelPackageDriftCheckBaselines `field:"optional" json:"driftCheckBaselines" yaml:"driftCheckBaselines"`
	// Details about inference jobs that can be run with models based on this model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#inference_specification SagemakerModelPackage#inference_specification}
	InferenceSpecification *SagemakerModelPackageInferenceSpecification `field:"optional" json:"inferenceSpecification" yaml:"inferenceSpecification"`
	// The time at which the model package was last modified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#last_modified_time SagemakerModelPackage#last_modified_time}
	LastModifiedTime *string `field:"optional" json:"lastModifiedTime" yaml:"lastModifiedTime"`
	// Metadata properties of the tracking entity, trial, or trial component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#metadata_properties SagemakerModelPackage#metadata_properties}
	MetadataProperties *SagemakerModelPackageMetadataProperties `field:"optional" json:"metadataProperties" yaml:"metadataProperties"`
	// The approval status of the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#model_approval_status SagemakerModelPackage#model_approval_status}
	ModelApprovalStatus *string `field:"optional" json:"modelApprovalStatus" yaml:"modelApprovalStatus"`
	// A structure that contains model metrics reports.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#model_metrics SagemakerModelPackage#model_metrics}
	ModelMetrics *SagemakerModelPackageModelMetrics `field:"optional" json:"modelMetrics" yaml:"modelMetrics"`
	// The description of the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#model_package_description SagemakerModelPackage#model_package_description}
	ModelPackageDescription *string `field:"optional" json:"modelPackageDescription" yaml:"modelPackageDescription"`
	// The name of the model package group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#model_package_group_name SagemakerModelPackage#model_package_group_name}
	ModelPackageGroupName *string `field:"optional" json:"modelPackageGroupName" yaml:"modelPackageGroupName"`
	// The name or arn of the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#model_package_name SagemakerModelPackage#model_package_name}
	ModelPackageName *string `field:"optional" json:"modelPackageName" yaml:"modelPackageName"`
	// Details about the current status of the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#model_package_status_details SagemakerModelPackage#model_package_status_details}
	ModelPackageStatusDetails *SagemakerModelPackageModelPackageStatusDetails `field:"optional" json:"modelPackageStatusDetails" yaml:"modelPackageStatusDetails"`
	// The version of the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#model_package_version SagemakerModelPackage#model_package_version}
	ModelPackageVersion *float64 `field:"optional" json:"modelPackageVersion" yaml:"modelPackageVersion"`
	// The Amazon Simple Storage Service (Amazon S3) path where the sample payload are stored pointing to single gzip compressed tar archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#sample_payload_url SagemakerModelPackage#sample_payload_url}
	SamplePayloadUrl *string `field:"optional" json:"samplePayloadUrl" yaml:"samplePayloadUrl"`
	// Indicates if you want to skip model validation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#skip_model_validation SagemakerModelPackage#skip_model_validation}
	SkipModelValidation *string `field:"optional" json:"skipModelValidation" yaml:"skipModelValidation"`
	// Details about the algorithm that was used to create the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#source_algorithm_specification SagemakerModelPackage#source_algorithm_specification}
	SourceAlgorithmSpecification *SagemakerModelPackageSourceAlgorithmSpecification `field:"optional" json:"sourceAlgorithmSpecification" yaml:"sourceAlgorithmSpecification"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#tags SagemakerModelPackage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The machine learning task your model package accomplishes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#task SagemakerModelPackage#task}
	Task *string `field:"optional" json:"task" yaml:"task"`
	// Specifies configurations for one or more transform jobs that Amazon SageMaker runs to test the model package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package#validation_specification SagemakerModelPackage#validation_specification}
	ValidationSpecification *SagemakerModelPackageValidationSpecification `field:"optional" json:"validationSpecification" yaml:"validationSpecification"`
}

