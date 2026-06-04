package dmsdatamigration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dmsdatamigration/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dms_data_migration awscc_dms_data_migration}.
type DmsDataMigration interface {
	cdktf.TerraformResource
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
	DataMigrationArn() *string
	DataMigrationCreateTime() *string
	DataMigrationIdentifier() *string
	SetDataMigrationIdentifier(val *string)
	DataMigrationIdentifierInput() *string
	DataMigrationName() *string
	SetDataMigrationName(val *string)
	DataMigrationNameInput() *string
	DataMigrationSettings() DmsDataMigrationDataMigrationSettingsOutputReference
	DataMigrationSettingsInput() interface{}
	DataMigrationType() *string
	SetDataMigrationType(val *string)
	DataMigrationTypeInput() *string
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
	MigrationProjectIdentifier() *string
	SetMigrationProjectIdentifier(val *string)
	MigrationProjectIdentifierInput() *string
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
	ServiceAccessRoleArn() *string
	SetServiceAccessRoleArn(val *string)
	ServiceAccessRoleArnInput() *string
	SourceDataSettings() DmsDataMigrationSourceDataSettingsList
	SourceDataSettingsInput() interface{}
	Tags() DmsDataMigrationTagsList
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutDataMigrationSettings(value *DmsDataMigrationDataMigrationSettings)
	PutSourceDataSettings(value interface{})
	PutTags(value interface{})
	ResetDataMigrationIdentifier()
	ResetDataMigrationName()
	ResetDataMigrationSettings()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSourceDataSettings()
	ResetTags()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
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

// The jsii proxy struct for DmsDataMigration
type jsiiProxy_DmsDataMigration struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsDataMigration) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationCreateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationCreateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationSettings() DmsDataMigrationDataMigrationSettingsOutputReference {
	var returns DmsDataMigrationDataMigrationSettingsOutputReference
	_jsii_.Get(
		j,
		"dataMigrationSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataMigrationSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DataMigrationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataMigrationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) MigrationProjectIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) MigrationProjectIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) ServiceAccessRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) SourceDataSettings() DmsDataMigrationSourceDataSettingsList {
	var returns DmsDataMigrationSourceDataSettingsList
	_jsii_.Get(
		j,
		"sourceDataSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) SourceDataSettingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDataSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) Tags() DmsDataMigrationTagsList {
	var returns DmsDataMigrationTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsDataMigration) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dms_data_migration awscc_dms_data_migration} Resource.
func NewDmsDataMigration(scope constructs.Construct, id *string, config *DmsDataMigrationConfig) DmsDataMigration {
	_init_.Initialize()

	if err := validateNewDmsDataMigrationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsDataMigration{}

	_jsii_.Create(
		"awscc.dmsDataMigration.DmsDataMigration",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/dms_data_migration awscc_dms_data_migration} Resource.
func NewDmsDataMigration_Override(d DmsDataMigration, scope constructs.Construct, id *string, config *DmsDataMigrationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dmsDataMigration.DmsDataMigration",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetDataMigrationIdentifier(val *string) {
	if err := j.validateSetDataMigrationIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataMigrationIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetDataMigrationName(val *string) {
	if err := j.validateSetDataMigrationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataMigrationName",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetDataMigrationType(val *string) {
	if err := j.validateSetDataMigrationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataMigrationType",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetMigrationProjectIdentifier(val *string) {
	if err := j.validateSetMigrationProjectIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationProjectIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsDataMigration)SetServiceAccessRoleArn(val *string) {
	if err := j.validateSetServiceAccessRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessRoleArn",
		val,
	)
}

// Generates CDKTF code for importing a DmsDataMigration resource upon running "cdktf plan <stack-name>".
func DmsDataMigration_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDmsDataMigration_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dmsDataMigration.DmsDataMigration",
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
func DmsDataMigration_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsDataMigration_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dmsDataMigration.DmsDataMigration",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsDataMigration_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsDataMigration_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dmsDataMigration.DmsDataMigration",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsDataMigration_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsDataMigration_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dmsDataMigration.DmsDataMigration",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsDataMigration_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dmsDataMigration.DmsDataMigration",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsDataMigration) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DmsDataMigration) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsDataMigration) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsDataMigration) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsDataMigration) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsDataMigration) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsDataMigration) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsDataMigration) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsDataMigration) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsDataMigration) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsDataMigration) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsDataMigration) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataMigration) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DmsDataMigration) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsDataMigration) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsDataMigration) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DmsDataMigration) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsDataMigration) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsDataMigration) PutDataMigrationSettings(value *DmsDataMigrationDataMigrationSettings) {
	if err := d.validatePutDataMigrationSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDataMigrationSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataMigration) PutSourceDataSettings(value interface{}) {
	if err := d.validatePutSourceDataSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSourceDataSettings",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataMigration) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsDataMigration) ResetDataMigrationIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetDataMigrationIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataMigration) ResetDataMigrationName() {
	_jsii_.InvokeVoid(
		d,
		"resetDataMigrationName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataMigration) ResetDataMigrationSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetDataMigrationSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataMigration) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataMigration) ResetSourceDataSettings() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceDataSettings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataMigration) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsDataMigration) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataMigration) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataMigration) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataMigration) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataMigration) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsDataMigration) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

