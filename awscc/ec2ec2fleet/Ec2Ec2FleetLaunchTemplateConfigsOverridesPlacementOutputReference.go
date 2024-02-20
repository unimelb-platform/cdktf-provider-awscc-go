package ec2ec2fleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2ec2fleet/internal"
)

type Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference interface {
	cdktf.ComplexObject
	Affinity() *string
	SetAffinity(val *string)
	AffinityInput() *string
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GroupName() *string
	SetGroupName(val *string)
	GroupNameInput() *string
	HostId() *string
	SetHostId(val *string)
	HostIdInput() *string
	HostResourceGroupArn() *string
	SetHostResourceGroupArn(val *string)
	HostResourceGroupArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PartitionNumber() *float64
	SetPartitionNumber(val *float64)
	PartitionNumberInput() *float64
	SpreadDomain() *string
	SetSpreadDomain(val *string)
	SpreadDomainInput() *string
	Tenancy() *string
	SetTenancy(val *string)
	TenancyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAffinity()
	ResetAvailabilityZone()
	ResetGroupName()
	ResetHostId()
	ResetHostResourceGroupArn()
	ResetPartitionNumber()
	ResetSpreadDomain()
	ResetTenancy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference
type jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) Affinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"affinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) AffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"affinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) HostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) HostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) HostResourceGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) HostResourceGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostResourceGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) PartitionNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) PartitionNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"partitionNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) SpreadDomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spreadDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) SpreadDomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spreadDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) Tenancy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) TenancyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenancyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEc2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference {
	_init_.Initialize()

	if err := validateNewEc2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference{}

	_jsii_.Create(
		"awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference_Override(e Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2Ec2Fleet.Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetAffinity(val *string) {
	if err := j.validateSetAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affinity",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetGroupName(val *string) {
	if err := j.validateSetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupName",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetHostId(val *string) {
	if err := j.validateSetHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostId",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetHostResourceGroupArn(val *string) {
	if err := j.validateSetHostResourceGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostResourceGroupArn",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetPartitionNumber(val *float64) {
	if err := j.validateSetPartitionNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partitionNumber",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetSpreadDomain(val *string) {
	if err := j.validateSetSpreadDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spreadDomain",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetTenancy(val *string) {
	if err := j.validateSetTenancyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenancy",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetAffinity() {
	_jsii_.InvokeVoid(
		e,
		"resetAffinity",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		e,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetGroupName() {
	_jsii_.InvokeVoid(
		e,
		"resetGroupName",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetHostId() {
	_jsii_.InvokeVoid(
		e,
		"resetHostId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetHostResourceGroupArn() {
	_jsii_.InvokeVoid(
		e,
		"resetHostResourceGroupArn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetPartitionNumber() {
	_jsii_.InvokeVoid(
		e,
		"resetPartitionNumber",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetSpreadDomain() {
	_jsii_.InvokeVoid(
		e,
		"resetSpreadDomain",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ResetTenancy() {
	_jsii_.InvokeVoid(
		e,
		"resetTenancy",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2Ec2FleetLaunchTemplateConfigsOverridesPlacementOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

