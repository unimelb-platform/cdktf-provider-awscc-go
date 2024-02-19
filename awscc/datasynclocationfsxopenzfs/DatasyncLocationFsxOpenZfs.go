package datasynclocationfsxopenzfs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/datasynclocationfsxopenzfs/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs awscc_datasync_location_fsx_open_zfs}.
type DatasyncLocationFsxOpenZfs interface {
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
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FsxFilesystemArn() *string
	SetFsxFilesystemArn(val *string)
	FsxFilesystemArnInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LocationArn() *string
	LocationUri() *string
	// The tree node.
	Node() constructs.Node
	Protocol() DatasyncLocationFsxOpenZfsProtocolOutputReference
	ProtocolInput() interface{}
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
	SecurityGroupArns() *[]*string
	SetSecurityGroupArns(val *[]*string)
	SecurityGroupArnsInput() *[]*string
	Subdirectory() *string
	SetSubdirectory(val *string)
	SubdirectoryInput() *string
	Tags() DatasyncLocationFsxOpenZfsTagsList
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
	PutProtocol(value *DatasyncLocationFsxOpenZfsProtocol)
	PutTags(value interface{})
	ResetFsxFilesystemArn()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubdirectory()
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

// The jsii proxy struct for DatasyncLocationFsxOpenZfs
type jsiiProxy_DatasyncLocationFsxOpenZfs struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) FsxFilesystemArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxFilesystemArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) FsxFilesystemArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fsxFilesystemArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) LocationArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) LocationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Protocol() DatasyncLocationFsxOpenZfsProtocolOutputReference {
	var returns DatasyncLocationFsxOpenZfsProtocolOutputReference
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) ProtocolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) SecurityGroupArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) SecurityGroupArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Subdirectory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) SubdirectoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subdirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) Tags() DatasyncLocationFsxOpenZfsTagsList {
	var returns DatasyncLocationFsxOpenZfsTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs awscc_datasync_location_fsx_open_zfs} Resource.
func NewDatasyncLocationFsxOpenZfs(scope constructs.Construct, id *string, config *DatasyncLocationFsxOpenZfsConfig) DatasyncLocationFsxOpenZfs {
	_init_.Initialize()

	if err := validateNewDatasyncLocationFsxOpenZfsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatasyncLocationFsxOpenZfs{}

	_jsii_.Create(
		"awscc.datasyncLocationFsxOpenZfs.DatasyncLocationFsxOpenZfs",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_fsx_open_zfs awscc_datasync_location_fsx_open_zfs} Resource.
func NewDatasyncLocationFsxOpenZfs_Override(d DatasyncLocationFsxOpenZfs, scope constructs.Construct, id *string, config *DatasyncLocationFsxOpenZfsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.datasyncLocationFsxOpenZfs.DatasyncLocationFsxOpenZfs",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetFsxFilesystemArn(val *string) {
	if err := j.validateSetFsxFilesystemArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsxFilesystemArn",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetSecurityGroupArns(val *[]*string) {
	if err := j.validateSetSecurityGroupArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupArns",
		val,
	)
}

func (j *jsiiProxy_DatasyncLocationFsxOpenZfs)SetSubdirectory(val *string) {
	if err := j.validateSetSubdirectoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subdirectory",
		val,
	)
}

// Generates CDKTF code for importing a DatasyncLocationFsxOpenZfs resource upon running "cdktf plan <stack-name>".
func DatasyncLocationFsxOpenZfs_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDatasyncLocationFsxOpenZfs_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationFsxOpenZfs.DatasyncLocationFsxOpenZfs",
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
func DatasyncLocationFsxOpenZfs_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationFsxOpenZfs_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationFsxOpenZfs.DatasyncLocationFsxOpenZfs",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatasyncLocationFsxOpenZfs_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationFsxOpenZfs_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationFsxOpenZfs.DatasyncLocationFsxOpenZfs",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatasyncLocationFsxOpenZfs_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatasyncLocationFsxOpenZfs_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.datasyncLocationFsxOpenZfs.DatasyncLocationFsxOpenZfs",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatasyncLocationFsxOpenZfs_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.datasyncLocationFsxOpenZfs.DatasyncLocationFsxOpenZfs",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) PutProtocol(value *DatasyncLocationFsxOpenZfsProtocol) {
	if err := d.validatePutProtocolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putProtocol",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) PutTags(value interface{}) {
	if err := d.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTags",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ResetFsxFilesystemArn() {
	_jsii_.InvokeVoid(
		d,
		"resetFsxFilesystemArn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ResetSubdirectory() {
	_jsii_.InvokeVoid(
		d,
		"resetSubdirectory",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatasyncLocationFsxOpenZfs) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

