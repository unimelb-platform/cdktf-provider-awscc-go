package sagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakermodelpackage/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package awscc_sagemaker_model_package}.
type SagemakerModelPackage interface {
	cdktf.TerraformResource
	AdditionalInferenceSpecifications() SagemakerModelPackageAdditionalInferenceSpecificationsList
	AdditionalInferenceSpecificationsInput() interface{}
	AdditionalInferenceSpecificationsToAdd() SagemakerModelPackageAdditionalInferenceSpecificationsToAddList
	AdditionalInferenceSpecificationsToAddInput() interface{}
	ApprovalDescription() *string
	SetApprovalDescription(val *string)
	ApprovalDescriptionInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertifyForMarketplace() interface{}
	SetCertifyForMarketplace(val interface{})
	CertifyForMarketplaceInput() interface{}
	ClientToken() *string
	SetClientToken(val *string)
	ClientTokenInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	CustomerMetadataProperties() *map[string]*string
	SetCustomerMetadataProperties(val *map[string]*string)
	CustomerMetadataPropertiesInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	DriftCheckBaselines() SagemakerModelPackageDriftCheckBaselinesOutputReference
	DriftCheckBaselinesInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InferenceSpecification() SagemakerModelPackageInferenceSpecificationOutputReference
	InferenceSpecificationInput() interface{}
	LastModifiedTime() *string
	SetLastModifiedTime(val *string)
	LastModifiedTimeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetadataProperties() SagemakerModelPackageMetadataPropertiesOutputReference
	MetadataPropertiesInput() interface{}
	ModelApprovalStatus() *string
	SetModelApprovalStatus(val *string)
	ModelApprovalStatusInput() *string
	ModelMetrics() SagemakerModelPackageModelMetricsOutputReference
	ModelMetricsInput() interface{}
	ModelPackageArn() *string
	ModelPackageDescription() *string
	SetModelPackageDescription(val *string)
	ModelPackageDescriptionInput() *string
	ModelPackageGroupName() *string
	SetModelPackageGroupName(val *string)
	ModelPackageGroupNameInput() *string
	ModelPackageName() *string
	SetModelPackageName(val *string)
	ModelPackageNameInput() *string
	ModelPackageStatus() *string
	ModelPackageStatusDetails() SagemakerModelPackageModelPackageStatusDetailsOutputReference
	ModelPackageStatusDetailsInput() interface{}
	ModelPackageVersion() *float64
	SetModelPackageVersion(val *float64)
	ModelPackageVersionInput() *float64
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	SamplePayloadUrl() *string
	SetSamplePayloadUrl(val *string)
	SamplePayloadUrlInput() *string
	SkipModelValidation() *string
	SetSkipModelValidation(val *string)
	SkipModelValidationInput() *string
	SourceAlgorithmSpecification() SagemakerModelPackageSourceAlgorithmSpecificationOutputReference
	SourceAlgorithmSpecificationInput() interface{}
	Tags() SagemakerModelPackageTagsList
	TagsInput() interface{}
	Task() *string
	SetTask(val *string)
	TaskInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidationSpecification() SagemakerModelPackageValidationSpecificationOutputReference
	ValidationSpecificationInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdditionalInferenceSpecifications(value interface{})
	PutAdditionalInferenceSpecificationsToAdd(value interface{})
	PutDriftCheckBaselines(value *SagemakerModelPackageDriftCheckBaselines)
	PutInferenceSpecification(value *SagemakerModelPackageInferenceSpecification)
	PutMetadataProperties(value *SagemakerModelPackageMetadataProperties)
	PutModelMetrics(value *SagemakerModelPackageModelMetrics)
	PutModelPackageStatusDetails(value *SagemakerModelPackageModelPackageStatusDetails)
	PutSourceAlgorithmSpecification(value *SagemakerModelPackageSourceAlgorithmSpecification)
	PutTags(value interface{})
	PutValidationSpecification(value *SagemakerModelPackageValidationSpecification)
	ResetAdditionalInferenceSpecifications()
	ResetAdditionalInferenceSpecificationsToAdd()
	ResetApprovalDescription()
	ResetCertifyForMarketplace()
	ResetClientToken()
	ResetCustomerMetadataProperties()
	ResetDomain()
	ResetDriftCheckBaselines()
	ResetInferenceSpecification()
	ResetLastModifiedTime()
	ResetMetadataProperties()
	ResetModelApprovalStatus()
	ResetModelMetrics()
	ResetModelPackageDescription()
	ResetModelPackageGroupName()
	ResetModelPackageName()
	ResetModelPackageStatusDetails()
	ResetModelPackageVersion()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSamplePayloadUrl()
	ResetSkipModelValidation()
	ResetSourceAlgorithmSpecification()
	ResetTags()
	ResetTask()
	ResetValidationSpecification()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SagemakerModelPackage
type jsiiProxy_SagemakerModelPackage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SagemakerModelPackage) AdditionalInferenceSpecifications() SagemakerModelPackageAdditionalInferenceSpecificationsList {
	var returns SagemakerModelPackageAdditionalInferenceSpecificationsList
	_jsii_.Get(
		j,
		"additionalInferenceSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) AdditionalInferenceSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInferenceSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) AdditionalInferenceSpecificationsToAdd() SagemakerModelPackageAdditionalInferenceSpecificationsToAddList {
	var returns SagemakerModelPackageAdditionalInferenceSpecificationsToAddList
	_jsii_.Get(
		j,
		"additionalInferenceSpecificationsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) AdditionalInferenceSpecificationsToAddInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalInferenceSpecificationsToAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ApprovalDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ApprovalDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) CertifyForMarketplace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certifyForMarketplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) CertifyForMarketplaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certifyForMarketplaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ClientTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) CustomerMetadataProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customerMetadataProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) CustomerMetadataPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customerMetadataPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) DriftCheckBaselines() SagemakerModelPackageDriftCheckBaselinesOutputReference {
	var returns SagemakerModelPackageDriftCheckBaselinesOutputReference
	_jsii_.Get(
		j,
		"driftCheckBaselines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) DriftCheckBaselinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"driftCheckBaselinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) InferenceSpecification() SagemakerModelPackageInferenceSpecificationOutputReference {
	var returns SagemakerModelPackageInferenceSpecificationOutputReference
	_jsii_.Get(
		j,
		"inferenceSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) InferenceSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inferenceSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) LastModifiedTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) MetadataProperties() SagemakerModelPackageMetadataPropertiesOutputReference {
	var returns SagemakerModelPackageMetadataPropertiesOutputReference
	_jsii_.Get(
		j,
		"metadataProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) MetadataPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metadataPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelApprovalStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelApprovalStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelApprovalStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelApprovalStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelMetrics() SagemakerModelPackageModelMetricsOutputReference {
	var returns SagemakerModelPackageModelMetricsOutputReference
	_jsii_.Get(
		j,
		"modelMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelMetricsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageStatusDetails() SagemakerModelPackageModelPackageStatusDetailsOutputReference {
	var returns SagemakerModelPackageModelPackageStatusDetailsOutputReference
	_jsii_.Get(
		j,
		"modelPackageStatusDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageStatusDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelPackageStatusDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelPackageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ModelPackageVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelPackageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) SamplePayloadUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samplePayloadUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) SamplePayloadUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samplePayloadUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) SkipModelValidation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipModelValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) SkipModelValidationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipModelValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) SourceAlgorithmSpecification() SagemakerModelPackageSourceAlgorithmSpecificationOutputReference {
	var returns SagemakerModelPackageSourceAlgorithmSpecificationOutputReference
	_jsii_.Get(
		j,
		"sourceAlgorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) SourceAlgorithmSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAlgorithmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Tags() SagemakerModelPackageTagsList {
	var returns SagemakerModelPackageTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) Task() *string {
	var returns *string
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) TaskInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"taskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ValidationSpecification() SagemakerModelPackageValidationSpecificationOutputReference {
	var returns SagemakerModelPackageValidationSpecificationOutputReference
	_jsii_.Get(
		j,
		"validationSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerModelPackage) ValidationSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationSpecificationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package awscc_sagemaker_model_package} Resource.
func NewSagemakerModelPackage(scope constructs.Construct, id *string, config *SagemakerModelPackageConfig) SagemakerModelPackage {
	_init_.Initialize()

	if err := validateNewSagemakerModelPackageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerModelPackage{}

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_package awscc_sagemaker_model_package} Resource.
func NewSagemakerModelPackage_Override(s SagemakerModelPackage, scope constructs.Construct, id *string, config *SagemakerModelPackageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerModelPackage.SagemakerModelPackage",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetApprovalDescription(val *string) {
	if err := j.validateSetApprovalDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approvalDescription",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetCertifyForMarketplace(val interface{}) {
	if err := j.validateSetCertifyForMarketplaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certifyForMarketplace",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetClientToken(val *string) {
	if err := j.validateSetClientTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientToken",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetCustomerMetadataProperties(val *map[string]*string) {
	if err := j.validateSetCustomerMetadataPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerMetadataProperties",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetLastModifiedTime(val *string) {
	if err := j.validateSetLastModifiedTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastModifiedTime",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetModelApprovalStatus(val *string) {
	if err := j.validateSetModelApprovalStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelApprovalStatus",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetModelPackageDescription(val *string) {
	if err := j.validateSetModelPackageDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageDescription",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetModelPackageGroupName(val *string) {
	if err := j.validateSetModelPackageGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageGroupName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetModelPackageName(val *string) {
	if err := j.validateSetModelPackageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageName",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetModelPackageVersion(val *float64) {
	if err := j.validateSetModelPackageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modelPackageVersion",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetSamplePayloadUrl(val *string) {
	if err := j.validateSetSamplePayloadUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"samplePayloadUrl",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetSkipModelValidation(val *string) {
	if err := j.validateSetSkipModelValidationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipModelValidation",
		val,
	)
}

func (j *jsiiProxy_SagemakerModelPackage)SetTask(val *string) {
	if err := j.validateSetTaskParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"task",
		val,
	)
}

// Generates CDKTF code for importing a SagemakerModelPackage resource upon running "cdktf plan <stack-name>".
func SagemakerModelPackage_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSagemakerModelPackage_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelPackage.SagemakerModelPackage",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SagemakerModelPackage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelPackage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelPackage.SagemakerModelPackage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerModelPackage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelPackage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelPackage.SagemakerModelPackage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SagemakerModelPackage_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSagemakerModelPackage_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.sagemakerModelPackage.SagemakerModelPackage",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SagemakerModelPackage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.sagemakerModelPackage.SagemakerModelPackage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutAdditionalInferenceSpecifications(value interface{}) {
	if err := s.validatePutAdditionalInferenceSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdditionalInferenceSpecifications",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutAdditionalInferenceSpecificationsToAdd(value interface{}) {
	if err := s.validatePutAdditionalInferenceSpecificationsToAddParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAdditionalInferenceSpecificationsToAdd",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutDriftCheckBaselines(value *SagemakerModelPackageDriftCheckBaselines) {
	if err := s.validatePutDriftCheckBaselinesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putDriftCheckBaselines",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutInferenceSpecification(value *SagemakerModelPackageInferenceSpecification) {
	if err := s.validatePutInferenceSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putInferenceSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutMetadataProperties(value *SagemakerModelPackageMetadataProperties) {
	if err := s.validatePutMetadataPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putMetadataProperties",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutModelMetrics(value *SagemakerModelPackageModelMetrics) {
	if err := s.validatePutModelMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelMetrics",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutModelPackageStatusDetails(value *SagemakerModelPackageModelPackageStatusDetails) {
	if err := s.validatePutModelPackageStatusDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putModelPackageStatusDetails",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutSourceAlgorithmSpecification(value *SagemakerModelPackageSourceAlgorithmSpecification) {
	if err := s.validatePutSourceAlgorithmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSourceAlgorithmSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutTags(value interface{}) {
	if err := s.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTags",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) PutValidationSpecification(value *SagemakerModelPackageValidationSpecification) {
	if err := s.validatePutValidationSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putValidationSpecification",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetAdditionalInferenceSpecifications() {
	_jsii_.InvokeVoid(
		s,
		"resetAdditionalInferenceSpecifications",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetAdditionalInferenceSpecificationsToAdd() {
	_jsii_.InvokeVoid(
		s,
		"resetAdditionalInferenceSpecificationsToAdd",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetApprovalDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetApprovalDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetCertifyForMarketplace() {
	_jsii_.InvokeVoid(
		s,
		"resetCertifyForMarketplace",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetClientToken() {
	_jsii_.InvokeVoid(
		s,
		"resetClientToken",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetCustomerMetadataProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomerMetadataProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetDomain() {
	_jsii_.InvokeVoid(
		s,
		"resetDomain",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetDriftCheckBaselines() {
	_jsii_.InvokeVoid(
		s,
		"resetDriftCheckBaselines",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetInferenceSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetInferenceSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetLastModifiedTime() {
	_jsii_.InvokeVoid(
		s,
		"resetLastModifiedTime",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetMetadataProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetMetadataProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetModelApprovalStatus() {
	_jsii_.InvokeVoid(
		s,
		"resetModelApprovalStatus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetModelMetrics() {
	_jsii_.InvokeVoid(
		s,
		"resetModelMetrics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetModelPackageDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetModelPackageDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetModelPackageGroupName() {
	_jsii_.InvokeVoid(
		s,
		"resetModelPackageGroupName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetModelPackageName() {
	_jsii_.InvokeVoid(
		s,
		"resetModelPackageName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetModelPackageStatusDetails() {
	_jsii_.InvokeVoid(
		s,
		"resetModelPackageStatusDetails",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetModelPackageVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetModelPackageVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetSamplePayloadUrl() {
	_jsii_.InvokeVoid(
		s,
		"resetSamplePayloadUrl",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetSkipModelValidation() {
	_jsii_.InvokeVoid(
		s,
		"resetSkipModelValidation",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetSourceAlgorithmSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetSourceAlgorithmSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetTask() {
	_jsii_.InvokeVoid(
		s,
		"resetTask",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) ResetValidationSpecification() {
	_jsii_.InvokeVoid(
		s,
		"resetValidationSpecification",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SagemakerModelPackage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerModelPackage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

