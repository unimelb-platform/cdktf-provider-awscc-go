package ec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2ec2fleet/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet awscc_ec2_ec2_fleet}.
type Ec2Ec2Fleet interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Context() *string
	SetContext(val *string)
	ContextInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcessCapacityTerminationPolicy() *string
	SetExcessCapacityTerminationPolicy(val *string)
	ExcessCapacityTerminationPolicyInput() *string
	FleetId() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	LaunchTemplateConfigs() Ec2Ec2FleetLaunchTemplateConfigsList
	LaunchTemplateConfigsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OnDemandOptions() Ec2Ec2FleetOnDemandOptionsOutputReference
	OnDemandOptionsInput() interface{}
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
	ReplaceUnhealthyInstances() interface{}
	SetReplaceUnhealthyInstances(val interface{})
	ReplaceUnhealthyInstancesInput() interface{}
	SpotOptions() Ec2Ec2FleetSpotOptionsOutputReference
	SpotOptionsInput() interface{}
	TagSpecifications() Ec2Ec2FleetTagSpecificationsList
	TagSpecificationsInput() interface{}
	TargetCapacitySpecification() Ec2Ec2FleetTargetCapacitySpecificationOutputReference
	TargetCapacitySpecificationInput() interface{}
	TerminateInstancesWithExpiration() interface{}
	SetTerminateInstancesWithExpiration(val interface{})
	TerminateInstancesWithExpirationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	ValidFrom() *string
	SetValidFrom(val *string)
	ValidFromInput() *string
	ValidUntil() *string
	SetValidUntil(val *string)
	ValidUntilInput() *string
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
	PutLaunchTemplateConfigs(value interface{})
	PutOnDemandOptions(value *Ec2Ec2FleetOnDemandOptions)
	PutSpotOptions(value *Ec2Ec2FleetSpotOptions)
	PutTagSpecifications(value interface{})
	PutTargetCapacitySpecification(value *Ec2Ec2FleetTargetCapacitySpecification)
	ResetContext()
	ResetExcessCapacityTerminationPolicy()
	ResetOnDemandOptions()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplaceUnhealthyInstances()
	ResetSpotOptions()
	ResetTagSpecifications()
	ResetTerminateInstancesWithExpiration()
	ResetType()
	ResetValidFrom()
	ResetValidUntil()
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

// The jsii proxy struct for Ec2Ec2Fleet
type jsiiProxy_Ec2Ec2Fleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2Ec2Fleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Context() *string {
	var returns *string
	_jsii_.Get(
		j,
		"context",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ContextInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ExcessCapacityTerminationPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ExcessCapacityTerminationPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"excessCapacityTerminationPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) FleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) LaunchTemplateConfigs() Ec2Ec2FleetLaunchTemplateConfigsList {
	var returns Ec2Ec2FleetLaunchTemplateConfigsList
	_jsii_.Get(
		j,
		"launchTemplateConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) LaunchTemplateConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"launchTemplateConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) OnDemandOptions() Ec2Ec2FleetOnDemandOptionsOutputReference {
	var returns Ec2Ec2FleetOnDemandOptionsOutputReference
	_jsii_.Get(
		j,
		"onDemandOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) OnDemandOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"onDemandOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ReplaceUnhealthyInstances() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ReplaceUnhealthyInstancesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"replaceUnhealthyInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) SpotOptions() Ec2Ec2FleetSpotOptionsOutputReference {
	var returns Ec2Ec2FleetSpotOptionsOutputReference
	_jsii_.Get(
		j,
		"spotOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) SpotOptionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spotOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TagSpecifications() Ec2Ec2FleetTagSpecificationsList {
	var returns Ec2Ec2FleetTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TargetCapacitySpecification() Ec2Ec2FleetTargetCapacitySpecificationOutputReference {
	var returns Ec2Ec2FleetTargetCapacitySpecificationOutputReference
	_jsii_.Get(
		j,
		"targetCapacitySpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TargetCapacitySpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetCapacitySpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TerminateInstancesWithExpiration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpiration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TerminateInstancesWithExpirationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"terminateInstancesWithExpirationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ValidFrom() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFrom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ValidFromInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validFromInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2Fleet) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet awscc_ec2_ec2_fleet} Resource.
func NewEc2Ec2Fleet(scope constructs.Construct, id *string, config *Ec2Ec2FleetConfig) Ec2Ec2Fleet {
	_init_.Initialize()

	if err := validateNewEc2Ec2FleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Ec2Fleet{}

	_jsii_.Create(
		"awscc.ec2Ec2Fleet.Ec2Ec2Fleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_ec2_fleet awscc_ec2_ec2_fleet} Resource.
func NewEc2Ec2Fleet_Override(e Ec2Ec2Fleet, scope constructs.Construct, id *string, config *Ec2Ec2FleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2Ec2Fleet.Ec2Ec2Fleet",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetContext(val *string) {
	if err := j.validateSetContextParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"context",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetExcessCapacityTerminationPolicy(val *string) {
	if err := j.validateSetExcessCapacityTerminationPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excessCapacityTerminationPolicy",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetReplaceUnhealthyInstances(val interface{}) {
	if err := j.validateSetReplaceUnhealthyInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceUnhealthyInstances",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetTerminateInstancesWithExpiration(val interface{}) {
	if err := j.validateSetTerminateInstancesWithExpirationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminateInstancesWithExpiration",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetValidFrom(val *string) {
	if err := j.validateSetValidFromParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validFrom",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2Fleet)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

// Generates CDKTF code for importing a Ec2Ec2Fleet resource upon running "cdktf plan <stack-name>".
func Ec2Ec2Fleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2Ec2Fleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ec2Ec2Fleet.Ec2Ec2Fleet",
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
func Ec2Ec2Fleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Ec2Fleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2Ec2Fleet.Ec2Ec2Fleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Ec2Fleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Ec2Fleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2Ec2Fleet.Ec2Ec2Fleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2Ec2Fleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2Ec2Fleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2Ec2Fleet.Ec2Ec2Fleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2Ec2Fleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ec2Ec2Fleet.Ec2Ec2Fleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2Ec2Fleet) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Ec2Fleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Ec2Fleet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2Fleet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2Fleet) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) PutLaunchTemplateConfigs(value interface{}) {
	if err := e.validatePutLaunchTemplateConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putLaunchTemplateConfigs",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) PutOnDemandOptions(value *Ec2Ec2FleetOnDemandOptions) {
	if err := e.validatePutOnDemandOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putOnDemandOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) PutSpotOptions(value *Ec2Ec2FleetSpotOptions) {
	if err := e.validatePutSpotOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putSpotOptions",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) PutTagSpecifications(value interface{}) {
	if err := e.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) PutTargetCapacitySpecification(value *Ec2Ec2FleetTargetCapacitySpecification) {
	if err := e.validatePutTargetCapacitySpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTargetCapacitySpecification",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetContext() {
	_jsii_.InvokeVoid(
		e,
		"resetContext",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetExcessCapacityTerminationPolicy() {
	_jsii_.InvokeVoid(
		e,
		"resetExcessCapacityTerminationPolicy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetOnDemandOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetOnDemandOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetReplaceUnhealthyInstances() {
	_jsii_.InvokeVoid(
		e,
		"resetReplaceUnhealthyInstances",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetSpotOptions() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotOptions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetTerminateInstancesWithExpiration() {
	_jsii_.InvokeVoid(
		e,
		"resetTerminateInstancesWithExpiration",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetType() {
	_jsii_.InvokeVoid(
		e,
		"resetType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetValidFrom() {
	_jsii_.InvokeVoid(
		e,
		"resetValidFrom",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) ResetValidUntil() {
	_jsii_.InvokeVoid(
		e,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2Fleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2Fleet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2Fleet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2Fleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2Fleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2Fleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

