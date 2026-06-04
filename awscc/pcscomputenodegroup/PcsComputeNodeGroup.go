package pcscomputenodegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/pcscomputenodegroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_compute_node_group awscc_pcs_compute_node_group}.
type PcsComputeNodeGroup interface {
	cdktf.TerraformResource
	AmiId() *string
	SetAmiId(val *string)
	AmiIdInput() *string
	Arn() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClusterId() *string
	SetClusterId(val *string)
	ClusterIdInput() *string
	ComputeNodeGroupId() *string
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
	CustomLaunchTemplate() PcsComputeNodeGroupCustomLaunchTemplateOutputReference
	CustomLaunchTemplateInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ErrorInfo() PcsComputeNodeGroupErrorInfoList
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IamInstanceProfileArn() *string
	SetIamInstanceProfileArn(val *string)
	IamInstanceProfileArnInput() *string
	Id() *string
	InstanceConfigs() PcsComputeNodeGroupInstanceConfigsList
	InstanceConfigsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PurchaseOption() *string
	SetPurchaseOption(val *string)
	PurchaseOptionInput() *string
	// Experimental.
	RawOverrides() interface{}
	ScalingConfiguration() PcsComputeNodeGroupScalingConfigurationOutputReference
	ScalingConfigurationInput() interface{}
	SlurmConfiguration() PcsComputeNodeGroupSlurmConfigurationOutputReference
	SlurmConfigurationInput() interface{}
	SpotOptions() PcsComputeNodeGroupSpotOptionsOutputReference
	SpotOptionsInput() interface{}
	Status() *string
	SubnetIds() *[]*string
	SetSubnetIds(val *[]*string)
	SubnetIdsInput() *[]*string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
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
	PutCustomLaunchTemplate(value *PcsComputeNodeGroupCustomLaunchTemplate)
	PutInstanceConfigs(value interface{})
	PutScalingConfiguration(value *PcsComputeNodeGroupScalingConfiguration)
	PutSlurmConfiguration(value *PcsComputeNodeGroupSlurmConfiguration)
	PutSpotOptions(value *PcsComputeNodeGroupSpotOptions)
	ResetAmiId()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPurchaseOption()
	ResetSlurmConfiguration()
	ResetSpotOptions()
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

// The jsii proxy struct for PcsComputeNodeGroup
type jsiiProxy_PcsComputeNodeGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PcsComputeNodeGroup) AmiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) AmiIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amiIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ClusterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ClusterIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ComputeNodeGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeNodeGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) CustomLaunchTemplate() PcsComputeNodeGroupCustomLaunchTemplateOutputReference {
	var returns PcsComputeNodeGroupCustomLaunchTemplateOutputReference
	_jsii_.Get(
		j,
		"customLaunchTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) CustomLaunchTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customLaunchTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ErrorInfo() PcsComputeNodeGroupErrorInfoList {
	var returns PcsComputeNodeGroupErrorInfoList
	_jsii_.Get(
		j,
		"errorInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) IamInstanceProfileArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) IamInstanceProfileArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iamInstanceProfileArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) InstanceConfigs() PcsComputeNodeGroupInstanceConfigsList {
	var returns PcsComputeNodeGroupInstanceConfigsList
	_jsii_.Get(
		j,
		"instanceConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) InstanceConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) PurchaseOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) PurchaseOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"purchaseOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ScalingConfiguration() PcsComputeNodeGroupScalingConfigurationOutputReference {
	var returns PcsComputeNodeGroupScalingConfigurationOutputReference
	_jsii_.Get(
		j,
		"scalingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) ScalingConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scalingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) SlurmConfiguration() PcsComputeNodeGroupSlurmConfigurationOutputReference {
	var returns PcsComputeNodeGroupSlurmConfigurationOutputReference
	_jsii_.Get(
		j,
		"slurmConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) SlurmConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slurmConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) SpotOptions() PcsComputeNodeGroupSpotOptionsOutputReference {
	var returns PcsComputeNodeGroupSpotOptionsOutputReference
	_jsii_.Get(
		j,
		"spotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) SpotOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PcsComputeNodeGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_compute_node_group awscc_pcs_compute_node_group} Resource.
func NewPcsComputeNodeGroup(scope constructs.Construct, id *string, config *PcsComputeNodeGroupConfig) PcsComputeNodeGroup {
	_init_.Initialize()

	if err := validateNewPcsComputeNodeGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PcsComputeNodeGroup{}

	_jsii_.Create(
		"awscc.pcsComputeNodeGroup.PcsComputeNodeGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/pcs_compute_node_group awscc_pcs_compute_node_group} Resource.
func NewPcsComputeNodeGroup_Override(p PcsComputeNodeGroup, scope constructs.Construct, id *string, config *PcsComputeNodeGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.pcsComputeNodeGroup.PcsComputeNodeGroup",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetAmiId(val *string) {
	if err := j.validateSetAmiIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amiId",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetClusterId(val *string) {
	if err := j.validateSetClusterIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterId",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetIamInstanceProfileArn(val *string) {
	if err := j.validateSetIamInstanceProfileArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamInstanceProfileArn",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetPurchaseOption(val *string) {
	if err := j.validateSetPurchaseOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"purchaseOption",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_PcsComputeNodeGroup)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a PcsComputeNodeGroup resource upon running "cdktf plan <stack-name>".
func PcsComputeNodeGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePcsComputeNodeGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.pcsComputeNodeGroup.PcsComputeNodeGroup",
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
func PcsComputeNodeGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePcsComputeNodeGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.pcsComputeNodeGroup.PcsComputeNodeGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PcsComputeNodeGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePcsComputeNodeGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.pcsComputeNodeGroup.PcsComputeNodeGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PcsComputeNodeGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePcsComputeNodeGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.pcsComputeNodeGroup.PcsComputeNodeGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PcsComputeNodeGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.pcsComputeNodeGroup.PcsComputeNodeGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroup) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PcsComputeNodeGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PcsComputeNodeGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PcsComputeNodeGroup) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) PutCustomLaunchTemplate(value *PcsComputeNodeGroupCustomLaunchTemplate) {
	if err := p.validatePutCustomLaunchTemplateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putCustomLaunchTemplate",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) PutInstanceConfigs(value interface{}) {
	if err := p.validatePutInstanceConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putInstanceConfigs",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) PutScalingConfiguration(value *PcsComputeNodeGroupScalingConfiguration) {
	if err := p.validatePutScalingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putScalingConfiguration",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) PutSlurmConfiguration(value *PcsComputeNodeGroupSlurmConfiguration) {
	if err := p.validatePutSlurmConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSlurmConfiguration",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) PutSpotOptions(value *PcsComputeNodeGroupSpotOptions) {
	if err := p.validatePutSpotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putSpotOptions",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) ResetAmiId() {
	_jsii_.InvokeVoid(
		p,
		"resetAmiId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) ResetName() {
	_jsii_.InvokeVoid(
		p,
		"resetName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) ResetPurchaseOption() {
	_jsii_.InvokeVoid(
		p,
		"resetPurchaseOption",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) ResetSlurmConfiguration() {
	_jsii_.InvokeVoid(
		p,
		"resetSlurmConfiguration",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) ResetSpotOptions() {
	_jsii_.InvokeVoid(
		p,
		"resetSpotOptions",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) ResetTags() {
	_jsii_.InvokeVoid(
		p,
		"resetTags",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PcsComputeNodeGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PcsComputeNodeGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

