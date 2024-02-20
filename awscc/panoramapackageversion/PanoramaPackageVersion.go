package panoramapackageversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/panoramapackageversion/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_package_version awscc_panorama_package_version}.
type PanoramaPackageVersion interface {
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
	Id() *string
	IsLatestPatch() cdktf.IResolvable
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MarkLatest() interface{}
	SetMarkLatest(val interface{})
	MarkLatestInput() interface{}
	// The tree node.
	Node() constructs.Node
	OwnerAccount() *string
	SetOwnerAccount(val *string)
	OwnerAccountInput() *string
	PackageArn() *string
	PackageId() *string
	SetPackageId(val *string)
	PackageIdInput() *string
	PackageName() *string
	PackageVersion() *string
	SetPackageVersion(val *string)
	PackageVersionInput() *string
	PatchVersion() *string
	SetPatchVersion(val *string)
	PatchVersionInput() *string
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
	RegisteredTime() *float64
	Status() *string
	StatusDescription() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdatedLatestPatchVersion() *string
	SetUpdatedLatestPatchVersion(val *string)
	UpdatedLatestPatchVersionInput() *string
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
	ResetMarkLatest()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwnerAccount()
	ResetUpdatedLatestPatchVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PanoramaPackageVersion
type jsiiProxy_PanoramaPackageVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PanoramaPackageVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) IsLatestPatch() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"isLatestPatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) MarkLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"markLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) MarkLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"markLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) OwnerAccount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) OwnerAccountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PackageArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PackageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PackageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PackageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PackageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PackageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PatchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) PatchVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) RegisteredTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"registeredTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) StatusDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) UpdatedLatestPatchVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedLatestPatchVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PanoramaPackageVersion) UpdatedLatestPatchVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updatedLatestPatchVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_package_version awscc_panorama_package_version} Resource.
func NewPanoramaPackageVersion(scope constructs.Construct, id *string, config *PanoramaPackageVersionConfig) PanoramaPackageVersion {
	_init_.Initialize()

	if err := validateNewPanoramaPackageVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PanoramaPackageVersion{}

	_jsii_.Create(
		"awscc.panoramaPackageVersion.PanoramaPackageVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/panorama_package_version awscc_panorama_package_version} Resource.
func NewPanoramaPackageVersion_Override(p PanoramaPackageVersion, scope constructs.Construct, id *string, config *PanoramaPackageVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.panoramaPackageVersion.PanoramaPackageVersion",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetMarkLatest(val interface{}) {
	if err := j.validateSetMarkLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"markLatest",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetOwnerAccount(val *string) {
	if err := j.validateSetOwnerAccountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ownerAccount",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetPackageId(val *string) {
	if err := j.validateSetPackageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageId",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetPackageVersion(val *string) {
	if err := j.validateSetPackageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageVersion",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetPatchVersion(val *string) {
	if err := j.validateSetPatchVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchVersion",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PanoramaPackageVersion)SetUpdatedLatestPatchVersion(val *string) {
	if err := j.validateSetUpdatedLatestPatchVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updatedLatestPatchVersion",
		val,
	)
}

// Generates CDKTF code for importing a PanoramaPackageVersion resource upon running "cdktf plan <stack-name>".
func PanoramaPackageVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePanoramaPackageVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.panoramaPackageVersion.PanoramaPackageVersion",
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
func PanoramaPackageVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePanoramaPackageVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.panoramaPackageVersion.PanoramaPackageVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PanoramaPackageVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePanoramaPackageVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.panoramaPackageVersion.PanoramaPackageVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PanoramaPackageVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePanoramaPackageVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.panoramaPackageVersion.PanoramaPackageVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PanoramaPackageVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.panoramaPackageVersion.PanoramaPackageVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) ResetMarkLatest() {
	_jsii_.InvokeVoid(
		p,
		"resetMarkLatest",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) ResetOwnerAccount() {
	_jsii_.InvokeVoid(
		p,
		"resetOwnerAccount",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) ResetUpdatedLatestPatchVersion() {
	_jsii_.InvokeVoid(
		p,
		"resetUpdatedLatestPatchVersion",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PanoramaPackageVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PanoramaPackageVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

