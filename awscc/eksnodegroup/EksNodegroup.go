package eksnodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/eksnodegroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup awscc_eks_nodegroup}.
type EksNodegroup interface {
	cdktf.TerraformResource
	AmiType() *string
	SetAmiType(val *string)
	AmiTypeInput() *string
	Arn() *string
	CapacityType() *string
	SetCapacityType(val *string)
	CapacityTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterName() *string
	SetClusterName(val *string)
	ClusterNameInput() *string
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
	DiskSize() *float64
	SetDiskSize(val *float64)
	DiskSizeInput() *float64
	ForceUpdateEnabled() interface{}
	SetForceUpdateEnabled(val interface{})
	ForceUpdateEnabledInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InstanceTypes() *[]*string
	SetInstanceTypes(val *[]*string)
	InstanceTypesInput() *[]*string
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	LaunchTemplate() EksNodegroupLaunchTemplateOutputReference
	LaunchTemplateInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NodegroupName() *string
	SetNodegroupName(val *string)
	NodegroupNameInput() *string
	NodeRole() *string
	SetNodeRole(val *string)
	NodeRoleInput() *string
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
	ReleaseVersion() *string
	SetReleaseVersion(val *string)
	ReleaseVersionInput() *string
	RemoteAccess() EksNodegroupRemoteAccessOutputReference
	RemoteAccessInput() interface{}
	ScalingConfig() EksNodegroupScalingConfigOutputReference
	ScalingConfigInput() interface{}
	Subnets() *[]*string
	SetSubnets(val *[]*string)
	SubnetsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	Taints() EksNodegroupTaintsList
	TaintsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UpdateConfig() EksNodegroupUpdateConfigOutputReference
	UpdateConfigInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
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
	PutLaunchTemplate(value *EksNodegroupLaunchTemplate)
	PutRemoteAccess(value *EksNodegroupRemoteAccess)
	PutScalingConfig(value *EksNodegroupScalingConfig)
	PutTaints(value interface{})
	PutUpdateConfig(value *EksNodegroupUpdateConfig)
	ResetAmiType()
	ResetCapacityType()
	ResetDiskSize()
	ResetForceUpdateEnabled()
	ResetInstanceTypes()
	ResetLabels()
	ResetLaunchTemplate()
	ResetNodegroupName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReleaseVersion()
	ResetRemoteAccess()
	ResetScalingConfig()
	ResetTags()
	ResetTaints()
	ResetUpdateConfig()
	ResetVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EksNodegroup
type jsiiProxy_EksNodegroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EksNodegroup) AmiType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) AmiTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) CapacityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) CapacityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ClusterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ClusterNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) DiskSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ForceUpdateEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ForceUpdateEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forceUpdateEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) InstanceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) InstanceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"instanceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) LaunchTemplate() EksNodegroupLaunchTemplateOutputReference {
	var returns EksNodegroupLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"launchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) LaunchTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) NodegroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) NodegroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodegroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) NodeRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) NodeRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ReleaseVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ReleaseVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) RemoteAccess() EksNodegroupRemoteAccessOutputReference {
	var returns EksNodegroupRemoteAccessOutputReference
	_jsii_.Get(
		j,
		"remoteAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) RemoteAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ScalingConfig() EksNodegroupScalingConfigOutputReference {
	var returns EksNodegroupScalingConfigOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) ScalingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Subnets() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) SubnetsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Taints() EksNodegroupTaintsList {
	var returns EksNodegroupTaintsList
	_jsii_.Get(
		j,
		"taints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) TaintsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"taintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) UpdateConfig() EksNodegroupUpdateConfigOutputReference {
	var returns EksNodegroupUpdateConfigOutputReference
	_jsii_.Get(
		j,
		"updateConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) UpdateConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EksNodegroup) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup awscc_eks_nodegroup} Resource.
func NewEksNodegroup(scope constructs.Construct, id *string, config *EksNodegroupConfig) EksNodegroup {
	_init_.Initialize()

	if err := validateNewEksNodegroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EksNodegroup{}

	_jsii_.Create(
		"awscc.eksNodegroup.EksNodegroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/eks_nodegroup awscc_eks_nodegroup} Resource.
func NewEksNodegroup_Override(e EksNodegroup, scope constructs.Construct, id *string, config *EksNodegroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.eksNodegroup.EksNodegroup",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EksNodegroup)SetAmiType(val *string) {
	if err := j.validateSetAmiTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiType",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetCapacityType(val *string) {
	if err := j.validateSetCapacityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityType",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetClusterName(val *string) {
	if err := j.validateSetClusterNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterName",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetDiskSize(val *float64) {
	if err := j.validateSetDiskSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetForceUpdateEnabled(val interface{}) {
	if err := j.validateSetForceUpdateEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forceUpdateEnabled",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetInstanceTypes(val *[]*string) {
	if err := j.validateSetInstanceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTypes",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetNodegroupName(val *string) {
	if err := j.validateSetNodegroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodegroupName",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetNodeRole(val *string) {
	if err := j.validateSetNodeRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeRole",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetReleaseVersion(val *string) {
	if err := j.validateSetReleaseVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseVersion",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetSubnets(val *[]*string) {
	if err := j.validateSetSubnetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnets",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_EksNodegroup)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

// Generates CDKTF code for importing a EksNodegroup resource upon running "cdktf plan <stack-name>".
func EksNodegroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEksNodegroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.eksNodegroup.EksNodegroup",
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
func EksNodegroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksNodegroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.eksNodegroup.EksNodegroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksNodegroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksNodegroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.eksNodegroup.EksNodegroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EksNodegroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEksNodegroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.eksNodegroup.EksNodegroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EksNodegroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.eksNodegroup.EksNodegroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EksNodegroup) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EksNodegroup) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EksNodegroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EksNodegroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EksNodegroup) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EksNodegroup) PutLaunchTemplate(value *EksNodegroupLaunchTemplate) {
	if err := e.validatePutLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLaunchTemplate",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodegroup) PutRemoteAccess(value *EksNodegroupRemoteAccess) {
	if err := e.validatePutRemoteAccessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRemoteAccess",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodegroup) PutScalingConfig(value *EksNodegroupScalingConfig) {
	if err := e.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodegroup) PutTaints(value interface{}) {
	if err := e.validatePutTaintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTaints",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodegroup) PutUpdateConfig(value *EksNodegroupUpdateConfig) {
	if err := e.validatePutUpdateConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putUpdateConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EksNodegroup) ResetAmiType() {
	_jsii_.InvokeVoid(
		e,
		"resetAmiType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetCapacityType() {
	_jsii_.InvokeVoid(
		e,
		"resetCapacityType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetDiskSize() {
	_jsii_.InvokeVoid(
		e,
		"resetDiskSize",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetForceUpdateEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetForceUpdateEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetInstanceTypes() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetLabels() {
	_jsii_.InvokeVoid(
		e,
		"resetLabels",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetLaunchTemplate() {
	_jsii_.InvokeVoid(
		e,
		"resetLaunchTemplate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetNodegroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetNodegroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetReleaseVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetReleaseVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetRemoteAccess() {
	_jsii_.InvokeVoid(
		e,
		"resetRemoteAccess",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetScalingConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetScalingConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetTags() {
	_jsii_.InvokeVoid(
		e,
		"resetTags",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetTaints() {
	_jsii_.InvokeVoid(
		e,
		"resetTaints",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetUpdateConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetUpdateConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) ResetVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EksNodegroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EksNodegroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

