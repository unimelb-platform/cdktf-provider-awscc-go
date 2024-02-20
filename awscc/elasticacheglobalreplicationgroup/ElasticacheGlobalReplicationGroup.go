package elasticacheglobalreplicationgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticacheglobalreplicationgroup/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_global_replication_group awscc_elasticache_global_replication_group}.
type ElasticacheGlobalReplicationGroup interface {
	cdktf.TerraformResource
	AutomaticFailoverEnabled() interface{}
	SetAutomaticFailoverEnabled(val interface{})
	AutomaticFailoverEnabledInput() interface{}
	CacheNodeType() *string
	SetCacheNodeType(val *string)
	CacheNodeTypeInput() *string
	CacheParameterGroupName() *string
	SetCacheParameterGroupName(val *string)
	CacheParameterGroupNameInput() *string
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
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GlobalNodeGroupCount() *float64
	SetGlobalNodeGroupCount(val *float64)
	GlobalNodeGroupCountInput() *float64
	GlobalReplicationGroupDescription() *string
	SetGlobalReplicationGroupDescription(val *string)
	GlobalReplicationGroupDescriptionInput() *string
	GlobalReplicationGroupId() *string
	GlobalReplicationGroupIdSuffix() *string
	SetGlobalReplicationGroupIdSuffix(val *string)
	GlobalReplicationGroupIdSuffixInput() *string
	Id() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Members() ElasticacheGlobalReplicationGroupMembersList
	MembersInput() interface{}
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
	RegionalConfigurations() ElasticacheGlobalReplicationGroupRegionalConfigurationsList
	RegionalConfigurationsInput() interface{}
	Status() *string
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
	PutMembers(value interface{})
	PutRegionalConfigurations(value interface{})
	ResetAutomaticFailoverEnabled()
	ResetCacheNodeType()
	ResetCacheParameterGroupName()
	ResetEngineVersion()
	ResetGlobalNodeGroupCount()
	ResetGlobalReplicationGroupDescription()
	ResetGlobalReplicationGroupIdSuffix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegionalConfigurations()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ElasticacheGlobalReplicationGroup
type jsiiProxy_ElasticacheGlobalReplicationGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) AutomaticFailoverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) AutomaticFailoverEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticFailoverEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) CacheNodeType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNodeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) CacheNodeTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNodeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) CacheParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) CacheParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) GlobalNodeGroupCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"globalNodeGroupCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) GlobalNodeGroupCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"globalNodeGroupCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) GlobalReplicationGroupDescription() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupDescription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) GlobalReplicationGroupDescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupDescriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) GlobalReplicationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) GlobalReplicationGroupIdSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupIdSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) GlobalReplicationGroupIdSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"globalReplicationGroupIdSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Members() ElasticacheGlobalReplicationGroupMembersList {
	var returns ElasticacheGlobalReplicationGroupMembersList
	_jsii_.Get(
		j,
		"members",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) MembersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) RegionalConfigurations() ElasticacheGlobalReplicationGroupRegionalConfigurationsList {
	var returns ElasticacheGlobalReplicationGroupRegionalConfigurationsList
	_jsii_.Get(
		j,
		"regionalConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) RegionalConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"regionalConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_global_replication_group awscc_elasticache_global_replication_group} Resource.
func NewElasticacheGlobalReplicationGroup(scope constructs.Construct, id *string, config *ElasticacheGlobalReplicationGroupConfig) ElasticacheGlobalReplicationGroup {
	_init_.Initialize()

	if err := validateNewElasticacheGlobalReplicationGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElasticacheGlobalReplicationGroup{}

	_jsii_.Create(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticache_global_replication_group awscc_elasticache_global_replication_group} Resource.
func NewElasticacheGlobalReplicationGroup_Override(e ElasticacheGlobalReplicationGroup, scope constructs.Construct, id *string, config *ElasticacheGlobalReplicationGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroup",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetAutomaticFailoverEnabled(val interface{}) {
	if err := j.validateSetAutomaticFailoverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticFailoverEnabled",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetCacheNodeType(val *string) {
	if err := j.validateSetCacheNodeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheNodeType",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetCacheParameterGroupName(val *string) {
	if err := j.validateSetCacheParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetGlobalNodeGroupCount(val *float64) {
	if err := j.validateSetGlobalNodeGroupCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalNodeGroupCount",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetGlobalReplicationGroupDescription(val *string) {
	if err := j.validateSetGlobalReplicationGroupDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalReplicationGroupDescription",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetGlobalReplicationGroupIdSuffix(val *string) {
	if err := j.validateSetGlobalReplicationGroupIdSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"globalReplicationGroupIdSuffix",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ElasticacheGlobalReplicationGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a ElasticacheGlobalReplicationGroup resource upon running "cdktf plan <stack-name>".
func ElasticacheGlobalReplicationGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateElasticacheGlobalReplicationGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroup",
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
func ElasticacheGlobalReplicationGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheGlobalReplicationGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheGlobalReplicationGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheGlobalReplicationGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ElasticacheGlobalReplicationGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateElasticacheGlobalReplicationGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ElasticacheGlobalReplicationGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.elasticacheGlobalReplicationGroup.ElasticacheGlobalReplicationGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) PutMembers(value interface{}) {
	if err := e.validatePutMembersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putMembers",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) PutRegionalConfigurations(value interface{}) {
	if err := e.validatePutRegionalConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putRegionalConfigurations",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetAutomaticFailoverEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAutomaticFailoverEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetCacheNodeType() {
	_jsii_.InvokeVoid(
		e,
		"resetCacheNodeType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetCacheParameterGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetCacheParameterGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetGlobalNodeGroupCount() {
	_jsii_.InvokeVoid(
		e,
		"resetGlobalNodeGroupCount",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetGlobalReplicationGroupDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetGlobalReplicationGroupDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetGlobalReplicationGroupIdSuffix() {
	_jsii_.InvokeVoid(
		e,
		"resetGlobalReplicationGroupIdSuffix",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ResetRegionalConfigurations() {
	_jsii_.InvokeVoid(
		e,
		"resetRegionalConfigurations",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElasticacheGlobalReplicationGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

