package iotsitewiseassetmodel

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/iotsitewiseassetmodel/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model awscc_iotsitewise_asset_model}.
type IotsitewiseAssetModel interface {
	cdktf.TerraformResource
	AssetModelArn() *string
	AssetModelCompositeModels() IotsitewiseAssetModelAssetModelCompositeModelsList
	AssetModelCompositeModelsInput() interface{}
	AssetModelDescription() *string
	SetAssetModelDescription(val *string)
	AssetModelDescriptionInput() *string
	AssetModelHierarchies() IotsitewiseAssetModelAssetModelHierarchiesList
	AssetModelHierarchiesInput() interface{}
	AssetModelId() *string
	AssetModelName() *string
	SetAssetModelName(val *string)
	AssetModelNameInput() *string
	AssetModelProperties() IotsitewiseAssetModelAssetModelPropertiesList
	AssetModelPropertiesInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	Tags() IotsitewiseAssetModelTagsList
	TagsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAssetModelCompositeModels(value interface{})
	PutAssetModelHierarchies(value interface{})
	PutAssetModelProperties(value interface{})
	PutTags(value interface{})
	ResetAssetModelCompositeModels()
	ResetAssetModelDescription()
	ResetAssetModelHierarchies()
	ResetAssetModelProperties()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IotsitewiseAssetModel
type jsiiProxy_IotsitewiseAssetModel struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelCompositeModels() IotsitewiseAssetModelAssetModelCompositeModelsList {
	var returns IotsitewiseAssetModelAssetModelCompositeModelsList
	_jsii_.Get(
		j,
		"assetModelCompositeModels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelCompositeModelsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelCompositeModelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelHierarchies() IotsitewiseAssetModelAssetModelHierarchiesList {
	var returns IotsitewiseAssetModelAssetModelHierarchiesList
	_jsii_.Get(
		j,
		"assetModelHierarchies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelHierarchiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelHierarchiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"assetModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelProperties() IotsitewiseAssetModelAssetModelPropertiesList {
	var returns IotsitewiseAssetModelAssetModelPropertiesList
	_jsii_.Get(
		j,
		"assetModelProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) AssetModelPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"assetModelPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) Tags() IotsitewiseAssetModelTagsList {
	var returns IotsitewiseAssetModelTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotsitewiseAssetModel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model awscc_iotsitewise_asset_model} Resource.
func NewIotsitewiseAssetModel(scope constructs.Construct, id *string, config *IotsitewiseAssetModelConfig) IotsitewiseAssetModel {
	_init_.Initialize()

	if err := validateNewIotsitewiseAssetModelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotsitewiseAssetModel{}

	_jsii_.Create(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotsitewise_asset_model awscc_iotsitewise_asset_model} Resource.
func NewIotsitewiseAssetModel_Override(i IotsitewiseAssetModel, scope constructs.Construct, id *string, config *IotsitewiseAssetModelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModel",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetAssetModelDescription(val *string) {
	if err := j.validateSetAssetModelDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetModelDescription",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetAssetModelName(val *string) {
	if err := j.validateSetAssetModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"assetModelName",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotsitewiseAssetModel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a IotsitewiseAssetModel resource upon running "cdktf plan <stack-name>".
func IotsitewiseAssetModel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotsitewiseAssetModel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModel",
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
func IotsitewiseAssetModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotsitewiseAssetModel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotsitewiseAssetModel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotsitewiseAssetModel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotsitewiseAssetModel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotsitewiseAssetModel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotsitewiseAssetModel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.iotsitewiseAssetModel.IotsitewiseAssetModel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) PutAssetModelCompositeModels(value interface{}) {
	if err := i.validatePutAssetModelCompositeModelsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAssetModelCompositeModels",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) PutAssetModelHierarchies(value interface{}) {
	if err := i.validatePutAssetModelHierarchiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAssetModelHierarchies",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) PutAssetModelProperties(value interface{}) {
	if err := i.validatePutAssetModelPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putAssetModelProperties",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) PutTags(value interface{}) {
	if err := i.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTags",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) ResetAssetModelCompositeModels() {
	_jsii_.InvokeVoid(
		i,
		"resetAssetModelCompositeModels",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) ResetAssetModelDescription() {
	_jsii_.InvokeVoid(
		i,
		"resetAssetModelDescription",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) ResetAssetModelHierarchies() {
	_jsii_.InvokeVoid(
		i,
		"resetAssetModelHierarchies",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) ResetAssetModelProperties() {
	_jsii_.InvokeVoid(
		i,
		"resetAssetModelProperties",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotsitewiseAssetModel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotsitewiseAssetModel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

