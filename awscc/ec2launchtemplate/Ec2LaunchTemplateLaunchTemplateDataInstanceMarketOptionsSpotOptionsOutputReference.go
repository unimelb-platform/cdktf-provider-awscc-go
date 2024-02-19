package ec2launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2launchtemplate/internal"
)

type Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference interface {
	cdktf.ComplexObject
	BlockDurationMinutes() *float64
	SetBlockDurationMinutes(val *float64)
	BlockDurationMinutesInput() *float64
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
	InstanceInterruptionBehavior() *string
	SetInstanceInterruptionBehavior(val *string)
	InstanceInterruptionBehaviorInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxPrice() *string
	SetMaxPrice(val *string)
	MaxPriceInput() *string
	SpotInstanceType() *string
	SetSpotInstanceType(val *string)
	SpotInstanceTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ValidUntil() *string
	SetValidUntil(val *string)
	ValidUntilInput() *string
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
	ResetBlockDurationMinutes()
	ResetInstanceInterruptionBehavior()
	ResetMaxPrice()
	ResetSpotInstanceType()
	ResetValidUntil()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference
type jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) BlockDurationMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockDurationMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) BlockDurationMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"blockDurationMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) InstanceInterruptionBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) InstanceInterruptionBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceInterruptionBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) MaxPrice() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) MaxPriceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) SpotInstanceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstanceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) SpotInstanceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"spotInstanceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ValidUntil() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntil",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ValidUntilInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"validUntilInput",
		&returns,
	)
	return returns
}


func NewEc2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewEc2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference{}

	_jsii_.Create(
		"awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEc2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference_Override(e Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2LaunchTemplate.Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetBlockDurationMinutes(val *float64) {
	if err := j.validateSetBlockDurationMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blockDurationMinutes",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetInstanceInterruptionBehavior(val *string) {
	if err := j.validateSetInstanceInterruptionBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceInterruptionBehavior",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetMaxPrice(val *string) {
	if err := j.validateSetMaxPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxPrice",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetSpotInstanceType(val *string) {
	if err := j.validateSetSpotInstanceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"spotInstanceType",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference)SetValidUntil(val *string) {
	if err := j.validateSetValidUntilParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validUntil",
		val,
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ResetBlockDurationMinutes() {
	_jsii_.InvokeVoid(
		e,
		"resetBlockDurationMinutes",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ResetInstanceInterruptionBehavior() {
	_jsii_.InvokeVoid(
		e,
		"resetInstanceInterruptionBehavior",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ResetMaxPrice() {
	_jsii_.InvokeVoid(
		e,
		"resetMaxPrice",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ResetSpotInstanceType() {
	_jsii_.InvokeVoid(
		e,
		"resetSpotInstanceType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ResetValidUntil() {
	_jsii_.InvokeVoid(
		e,
		"resetValidUntil",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2LaunchTemplateLaunchTemplateDataInstanceMarketOptionsSpotOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

