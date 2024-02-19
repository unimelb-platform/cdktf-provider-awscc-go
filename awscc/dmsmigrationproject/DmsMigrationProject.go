package dmsmigrationproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dmsmigrationproject/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_migration_project awscc_dms_migration_project}.
type DmsMigrationProject interface {
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InstanceProfileArn() *string
	SetInstanceProfileArn(val *string)
	InstanceProfileArnInput() *string
	InstanceProfileIdentifier() *string
	SetInstanceProfileIdentifier(val *string)
	InstanceProfileIdentifierInput() *string
	InstanceProfileName() *string
	SetInstanceProfileName(val *string)
	InstanceProfileNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MigrationProjectArn() *string
	MigrationProjectCreationTime() *string
	SetMigrationProjectCreationTime(val *string)
	MigrationProjectCreationTimeInput() *string
	MigrationProjectIdentifier() *string
	SetMigrationProjectIdentifier(val *string)
	MigrationProjectIdentifierInput() *string
	MigrationProjectName() *string
	SetMigrationProjectName(val *string)
	MigrationProjectNameInput() *string
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
	SchemaConversionApplicationAttributes() DmsMigrationProjectSchemaConversionApplicationAttributesOutputReference
	SchemaConversionApplicationAttributesInput() interface{}
	SourceDataProviderDescriptors() DmsMigrationProjectSourceDataProviderDescriptorsList
	SourceDataProviderDescriptorsInput() interface{}
	Tags() DmsMigrationProjectTagsList
	TagsInput() interface{}
	TargetDataProviderDescriptors() DmsMigrationProjectTargetDataProviderDescriptorsList
	TargetDataProviderDescriptorsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TransformationRules() *string
	SetTransformationRules(val *string)
	TransformationRulesInput() *string
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
	PutSchemaConversionApplicationAttributes(value *DmsMigrationProjectSchemaConversionApplicationAttributes)
	PutSourceDataProviderDescriptors(value interface{})
	PutTags(value interface{})
	PutTargetDataProviderDescriptors(value interface{})
	ResetDescription()
	ResetInstanceProfileArn()
	ResetInstanceProfileIdentifier()
	ResetInstanceProfileName()
	ResetMigrationProjectCreationTime()
	ResetMigrationProjectIdentifier()
	ResetMigrationProjectName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSchemaConversionApplicationAttributes()
	ResetSourceDataProviderDescriptors()
	ResetTags()
	ResetTargetDataProviderDescriptors()
	ResetTransformationRules()
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

// The jsii proxy struct for DmsMigrationProject
type jsiiProxy_DmsMigrationProject struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DmsMigrationProject) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) InstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) InstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) InstanceProfileIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) InstanceProfileIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) InstanceProfileName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) InstanceProfileNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) MigrationProjectArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) MigrationProjectCreationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectCreationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) MigrationProjectCreationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectCreationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) MigrationProjectIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) MigrationProjectIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) MigrationProjectName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) MigrationProjectNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"migrationProjectNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) SchemaConversionApplicationAttributes() DmsMigrationProjectSchemaConversionApplicationAttributesOutputReference {
	var returns DmsMigrationProjectSchemaConversionApplicationAttributesOutputReference
	_jsii_.Get(
		j,
		"schemaConversionApplicationAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) SchemaConversionApplicationAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaConversionApplicationAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) SourceDataProviderDescriptors() DmsMigrationProjectSourceDataProviderDescriptorsList {
	var returns DmsMigrationProjectSourceDataProviderDescriptorsList
	_jsii_.Get(
		j,
		"sourceDataProviderDescriptors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) SourceDataProviderDescriptorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceDataProviderDescriptorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) Tags() DmsMigrationProjectTagsList {
	var returns DmsMigrationProjectTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TargetDataProviderDescriptors() DmsMigrationProjectTargetDataProviderDescriptorsList {
	var returns DmsMigrationProjectTargetDataProviderDescriptorsList
	_jsii_.Get(
		j,
		"targetDataProviderDescriptors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TargetDataProviderDescriptorsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetDataProviderDescriptorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TransformationRules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DmsMigrationProject) TransformationRulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"transformationRulesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_migration_project awscc_dms_migration_project} Resource.
func NewDmsMigrationProject(scope constructs.Construct, id *string, config *DmsMigrationProjectConfig) DmsMigrationProject {
	_init_.Initialize()

	if err := validateNewDmsMigrationProjectParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DmsMigrationProject{}

	_jsii_.Create(
		"awscc.dmsMigrationProject.DmsMigrationProject",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dms_migration_project awscc_dms_migration_project} Resource.
func NewDmsMigrationProject_Override(d DmsMigrationProject, scope constructs.Construct, id *string, config *DmsMigrationProjectConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dmsMigrationProject.DmsMigrationProject",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetInstanceProfileArn(val *string) {
	if err := j.validateSetInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetInstanceProfileIdentifier(val *string) {
	if err := j.validateSetInstanceProfileIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetInstanceProfileName(val *string) {
	if err := j.validateSetInstanceProfileNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfileName",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetMigrationProjectCreationTime(val *string) {
	if err := j.validateSetMigrationProjectCreationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationProjectCreationTime",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetMigrationProjectIdentifier(val *string) {
	if err := j.validateSetMigrationProjectIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationProjectIdentifier",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetMigrationProjectName(val *string) {
	if err := j.validateSetMigrationProjectNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"migrationProjectName",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DmsMigrationProject)SetTransformationRules(val *string) {
	if err := j.validateSetTransformationRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"transformationRules",
		val,
	)
}

// Generates CDKTF code for importing a DmsMigrationProject resource upon running "cdktf plan <stack-name>".
func DmsMigrationProject_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDmsMigrationProject_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dmsMigrationProject.DmsMigrationProject",
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
func DmsMigrationProject_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsMigrationProject_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dmsMigrationProject.DmsMigrationProject",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsMigrationProject_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsMigrationProject_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dmsMigrationProject.DmsMigrationProject",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DmsMigrationProject_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDmsMigrationProject_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dmsMigrationProject.DmsMigrationProject",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DmsMigrationProject_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dmsMigrationProject.DmsMigrationProject",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DmsMigrationProject) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DmsMigrationProject) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DmsMigrationProject) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DmsMigrationProject) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsMigrationProject) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DmsMigrationProject) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DmsMigrationProject) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DmsMigrationProject) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DmsMigrationProject) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DmsMigrationProject) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DmsMigrationProject) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DmsMigrationProject) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProject) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DmsMigrationProject) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DmsMigrationProject) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsMigrationProject) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DmsMigrationProject) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DmsMigrationProject) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DmsMigrationProject) PutSchemaConversionApplicationAttributes(value *DmsMigrationProjectSchemaConversionApplicationAttributes) {
	if err := d.validatePutSchemaConversionApplicationAttributesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchemaConversionApplicationAttributes",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsMigrationProject) PutSourceDataProviderDescriptors(value interface{}) {
	if err := d.validatePutSourceDataProviderDescriptorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSourceDataProviderDescriptors",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsMigrationProject) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsMigrationProject) PutTargetDataProviderDescriptors(value interface{}) {
	if err := d.validatePutTargetDataProviderDescriptorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTargetDataProviderDescriptors",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetInstanceProfileArn() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceProfileArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetInstanceProfileIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceProfileIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetInstanceProfileName() {
	_jsii_.InvokeVoid(
		d,
		"resetInstanceProfileName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetMigrationProjectCreationTime() {
	_jsii_.InvokeVoid(
		d,
		"resetMigrationProjectCreationTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetMigrationProjectIdentifier() {
	_jsii_.InvokeVoid(
		d,
		"resetMigrationProjectIdentifier",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetMigrationProjectName() {
	_jsii_.InvokeVoid(
		d,
		"resetMigrationProjectName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetSchemaConversionApplicationAttributes() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaConversionApplicationAttributes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetSourceDataProviderDescriptors() {
	_jsii_.InvokeVoid(
		d,
		"resetSourceDataProviderDescriptors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetTargetDataProviderDescriptors() {
	_jsii_.InvokeVoid(
		d,
		"resetTargetDataProviderDescriptors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) ResetTransformationRules() {
	_jsii_.InvokeVoid(
		d,
		"resetTransformationRules",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DmsMigrationProject) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProject) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProject) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProject) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProject) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DmsMigrationProject) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

