package dataawsccsagemakermodelpackage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccsagemakermodelpackage/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/sagemaker_model_package awscc_sagemaker_model_package}.
type DataAwsccSagemakerModelPackage interface {
	cdktf.TerraformDataSource
	AdditionalInferenceSpecifications() DataAwsccSagemakerModelPackageAdditionalInferenceSpecificationsList
	AdditionalInferenceSpecificationsToAdd() DataAwsccSagemakerModelPackageAdditionalInferenceSpecificationsToAddList
	ApprovalDescription() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertifyForMarketplace() cdktf.IResolvable
	ClientToken() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CreationTime() *string
	CustomerMetadataProperties() cdktf.StringMap
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	DriftCheckBaselines() DataAwsccSagemakerModelPackageDriftCheckBaselinesOutputReference
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InferenceSpecification() DataAwsccSagemakerModelPackageInferenceSpecificationOutputReference
	LastModifiedTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MetadataProperties() DataAwsccSagemakerModelPackageMetadataPropertiesOutputReference
	ModelApprovalStatus() *string
	ModelMetrics() DataAwsccSagemakerModelPackageModelMetricsOutputReference
	ModelPackageArn() *string
	ModelPackageDescription() *string
	ModelPackageGroupName() *string
	ModelPackageName() *string
	ModelPackageStatus() *string
	ModelPackageStatusDetails() DataAwsccSagemakerModelPackageModelPackageStatusDetailsOutputReference
	ModelPackageVersion() *float64
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	SamplePayloadUrl() *string
	SkipModelValidation() *string
	SourceAlgorithmSpecification() DataAwsccSagemakerModelPackageSourceAlgorithmSpecificationOutputReference
	Tags() DataAwsccSagemakerModelPackageTagsList
	Task() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidationSpecification() DataAwsccSagemakerModelPackageValidationSpecificationOutputReference
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsccSagemakerModelPackage
type jsiiProxy_DataAwsccSagemakerModelPackage struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) AdditionalInferenceSpecifications() DataAwsccSagemakerModelPackageAdditionalInferenceSpecificationsList {
	var returns DataAwsccSagemakerModelPackageAdditionalInferenceSpecificationsList
	_jsii_.Get(
		j,
		"additionalInferenceSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) AdditionalInferenceSpecificationsToAdd() DataAwsccSagemakerModelPackageAdditionalInferenceSpecificationsToAddList {
	var returns DataAwsccSagemakerModelPackageAdditionalInferenceSpecificationsToAddList
	_jsii_.Get(
		j,
		"additionalInferenceSpecificationsToAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ApprovalDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approvalDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) CertifyForMarketplace() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"certifyForMarketplace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ClientToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) CreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) CustomerMetadataProperties() cdktf.StringMap {
	var returns cdktf.StringMap
	_jsii_.Get(
		j,
		"customerMetadataProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) DriftCheckBaselines() DataAwsccSagemakerModelPackageDriftCheckBaselinesOutputReference {
	var returns DataAwsccSagemakerModelPackageDriftCheckBaselinesOutputReference
	_jsii_.Get(
		j,
		"driftCheckBaselines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) InferenceSpecification() DataAwsccSagemakerModelPackageInferenceSpecificationOutputReference {
	var returns DataAwsccSagemakerModelPackageInferenceSpecificationOutputReference
	_jsii_.Get(
		j,
		"inferenceSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) LastModifiedTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModifiedTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) MetadataProperties() DataAwsccSagemakerModelPackageMetadataPropertiesOutputReference {
	var returns DataAwsccSagemakerModelPackageMetadataPropertiesOutputReference
	_jsii_.Get(
		j,
		"metadataProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelApprovalStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelApprovalStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelMetrics() DataAwsccSagemakerModelPackageModelMetricsOutputReference {
	var returns DataAwsccSagemakerModelPackageModelMetricsOutputReference
	_jsii_.Get(
		j,
		"modelMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelPackageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelPackageDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelPackageGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelPackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelPackageStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelPackageStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelPackageStatusDetails() DataAwsccSagemakerModelPackageModelPackageStatusDetailsOutputReference {
	var returns DataAwsccSagemakerModelPackageModelPackageStatusDetailsOutputReference
	_jsii_.Get(
		j,
		"modelPackageStatusDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ModelPackageVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"modelPackageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) SamplePayloadUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"samplePayloadUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) SkipModelValidation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skipModelValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) SourceAlgorithmSpecification() DataAwsccSagemakerModelPackageSourceAlgorithmSpecificationOutputReference {
	var returns DataAwsccSagemakerModelPackageSourceAlgorithmSpecificationOutputReference
	_jsii_.Get(
		j,
		"sourceAlgorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Tags() DataAwsccSagemakerModelPackageTagsList {
	var returns DataAwsccSagemakerModelPackageTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) Task() *string {
	var returns *string
	_jsii_.Get(
		j,
		"task",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage) ValidationSpecification() DataAwsccSagemakerModelPackageValidationSpecificationOutputReference {
	var returns DataAwsccSagemakerModelPackageValidationSpecificationOutputReference
	_jsii_.Get(
		j,
		"validationSpecification",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/sagemaker_model_package awscc_sagemaker_model_package} Data Source.
func NewDataAwsccSagemakerModelPackage(scope constructs.Construct, id *string, config *DataAwsccSagemakerModelPackageConfig) DataAwsccSagemakerModelPackage {
	_init_.Initialize()

	if err := validateNewDataAwsccSagemakerModelPackageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccSagemakerModelPackage{}

	_jsii_.Create(
		"awscc.dataAwsccSagemakerModelPackage.DataAwsccSagemakerModelPackage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/sagemaker_model_package awscc_sagemaker_model_package} Data Source.
func NewDataAwsccSagemakerModelPackage_Override(d DataAwsccSagemakerModelPackage, scope constructs.Construct, id *string, config *DataAwsccSagemakerModelPackageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccSagemakerModelPackage.DataAwsccSagemakerModelPackage",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccSagemakerModelPackage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsccSagemakerModelPackage resource upon running "cdktf plan <stack-name>".
func DataAwsccSagemakerModelPackage_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccSagemakerModelPackage_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dataAwsccSagemakerModelPackage.DataAwsccSagemakerModelPackage",
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
func DataAwsccSagemakerModelPackage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccSagemakerModelPackage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccSagemakerModelPackage.DataAwsccSagemakerModelPackage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccSagemakerModelPackage_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccSagemakerModelPackage_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccSagemakerModelPackage.DataAwsccSagemakerModelPackage",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccSagemakerModelPackage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccSagemakerModelPackage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccSagemakerModelPackage.DataAwsccSagemakerModelPackage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccSagemakerModelPackage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dataAwsccSagemakerModelPackage.DataAwsccSagemakerModelPackage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccSagemakerModelPackage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

