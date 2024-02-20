package ec2capacityreservationfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2capacityreservationfleet/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet awscc_ec2_capacity_reservation_fleet}.
type Ec2CapacityReservationFleet interface {
	cdktf.TerraformResource
	AllocationStrategy() *string
	SetAllocationStrategy(val *string)
	AllocationStrategyInput() *string
	CapacityReservationFleetId() *string
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
	EndDate() *string
	SetEndDate(val *string)
	EndDateInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	InstanceMatchCriteria() *string
	SetInstanceMatchCriteria(val *string)
	InstanceMatchCriteriaInput() *string
	InstanceTypeSpecifications() Ec2CapacityReservationFleetInstanceTypeSpecificationsList
	InstanceTypeSpecificationsInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	NoRemoveEndDate() interface{}
	SetNoRemoveEndDate(val interface{})
	NoRemoveEndDateInput() interface{}
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
	RemoveEndDate() interface{}
	SetRemoveEndDate(val interface{})
	RemoveEndDateInput() interface{}
	TagSpecifications() Ec2CapacityReservationFleetTagSpecificationsList
	TagSpecificationsInput() interface{}
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TotalTargetCapacity() *float64
	SetTotalTargetCapacity(val *float64)
	TotalTargetCapacityInput() *float64
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
	PutInstanceTypeSpecifications(value interface{})
	PutTagSpecifications(value interface{})
	ResetAllocationStrategy()
	ResetEndDate()
	ResetInstanceMatchCriteria()
	ResetInstanceTypeSpecifications()
	ResetNoRemoveEndDate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRemoveEndDate()
	ResetTagSpecifications()
	ResetTenancy()
	ResetTotalTargetCapacity()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Ec2CapacityReservationFleet
type jsiiProxy_Ec2CapacityReservationFleet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) AllocationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) AllocationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) CapacityReservationFleetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationFleetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) EndDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) EndDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) InstanceMatchCriteria() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMatchCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) InstanceMatchCriteriaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceMatchCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) InstanceTypeSpecifications() Ec2CapacityReservationFleetInstanceTypeSpecificationsList {
	var returns Ec2CapacityReservationFleetInstanceTypeSpecificationsList
	_jsii_.Get(
		j,
		"instanceTypeSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) InstanceTypeSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"instanceTypeSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) NoRemoveEndDate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRemoveEndDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) NoRemoveEndDateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRemoveEndDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) RemoveEndDate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeEndDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) RemoveEndDateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"removeEndDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TagSpecifications() Ec2CapacityReservationFleetTagSpecificationsList {
	var returns Ec2CapacityReservationFleetTagSpecificationsList
	_jsii_.Get(
		j,
		"tagSpecifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TagSpecificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagSpecificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TotalTargetCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2CapacityReservationFleet) TotalTargetCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalTargetCapacityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet awscc_ec2_capacity_reservation_fleet} Resource.
func NewEc2CapacityReservationFleet(scope constructs.Construct, id *string, config *Ec2CapacityReservationFleetConfig) Ec2CapacityReservationFleet {
	_init_.Initialize()

	if err := validateNewEc2CapacityReservationFleetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2CapacityReservationFleet{}

	_jsii_.Create(
		"awscc.ec2CapacityReservationFleet.Ec2CapacityReservationFleet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_capacity_reservation_fleet awscc_ec2_capacity_reservation_fleet} Resource.
func NewEc2CapacityReservationFleet_Override(e Ec2CapacityReservationFleet, scope constructs.Construct, id *string, config *Ec2CapacityReservationFleetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2CapacityReservationFleet.Ec2CapacityReservationFleet",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetAllocationStrategy(val *string) {
	if err := j.validateSetAllocationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationStrategy",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetEndDate(val *string) {
	if err := j.validateSetEndDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDate",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetInstanceMatchCriteria(val *string) {
	if err := j.validateSetInstanceMatchCriteriaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceMatchCriteria",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetNoRemoveEndDate(val interface{}) {
	if err := j.validateSetNoRemoveEndDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRemoveEndDate",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetRemoveEndDate(val interface{}) {
	if err := j.validateSetRemoveEndDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"removeEndDate",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_Ec2CapacityReservationFleet)SetTotalTargetCapacity(val *float64) {
	if err := j.validateSetTotalTargetCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"totalTargetCapacity",
		val,
	)
}

// Generates CDKTF code for importing a Ec2CapacityReservationFleet resource upon running "cdktf plan <stack-name>".
func Ec2CapacityReservationFleet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEc2CapacityReservationFleet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.ec2CapacityReservationFleet.Ec2CapacityReservationFleet",
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
func Ec2CapacityReservationFleet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2CapacityReservationFleet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2CapacityReservationFleet.Ec2CapacityReservationFleet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2CapacityReservationFleet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2CapacityReservationFleet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2CapacityReservationFleet.Ec2CapacityReservationFleet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func Ec2CapacityReservationFleet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEc2CapacityReservationFleet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.ec2CapacityReservationFleet.Ec2CapacityReservationFleet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Ec2CapacityReservationFleet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.ec2CapacityReservationFleet.Ec2CapacityReservationFleet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2CapacityReservationFleet) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) PutInstanceTypeSpecifications(value interface{}) {
	if err := e.validatePutInstanceTypeSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putInstanceTypeSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) PutTagSpecifications(value interface{}) {
	if err := e.validatePutTagSpecificationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putTagSpecifications",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetAllocationStrategy() {
	_jsii_.InvokeVoid(
		e,
		"resetAllocationStrategy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetEndDate() {
	_jsii_.InvokeVoid(
		e,
		"resetEndDate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetInstanceMatchCriteria() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceMatchCriteria",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetInstanceTypeSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceTypeSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetNoRemoveEndDate() {
	_jsii_.InvokeVoid(
		e,
		"resetNoRemoveEndDate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetRemoveEndDate() {
	_jsii_.InvokeVoid(
		e,
		"resetRemoveEndDate",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetTagSpecifications() {
	_jsii_.InvokeVoid(
		e,
		"resetTagSpecifications",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetTenancy() {
	_jsii_.InvokeVoid(
		e,
		"resetTenancy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ResetTotalTargetCapacity() {
	_jsii_.InvokeVoid(
		e,
		"resetTotalTargetCapacity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2CapacityReservationFleet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

