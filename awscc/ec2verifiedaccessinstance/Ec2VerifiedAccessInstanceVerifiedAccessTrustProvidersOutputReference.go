package ec2verifiedaccessinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2verifiedaccessinstance/internal"
)

type Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference interface {
	cdktf.ComplexObject
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DeviceTrustProviderType() *string
	SetDeviceTrustProviderType(val *string)
	DeviceTrustProviderTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustProviderType() *string
	SetTrustProviderType(val *string)
	TrustProviderTypeInput() *string
	UserTrustProviderType() *string
	SetUserTrustProviderType(val *string)
	UserTrustProviderTypeInput() *string
	VerifiedAccessTrustProviderId() *string
	SetVerifiedAccessTrustProviderId(val *string)
	VerifiedAccessTrustProviderIdInput() *string
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
	ResetDescription()
	ResetDeviceTrustProviderType()
	ResetTrustProviderType()
	ResetUserTrustProviderType()
	ResetVerifiedAccessTrustProviderId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference
type jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) DeviceTrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTrustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) DeviceTrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTrustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) TrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) TrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) UserTrustProviderType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTrustProviderType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) UserTrustProviderTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userTrustProviderTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) VerifiedAccessTrustProviderId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessTrustProviderId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) VerifiedAccessTrustProviderIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verifiedAccessTrustProviderIdInput",
		&returns,
	)
	return returns
}


func NewEc2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference {
	_init_.Initialize()

	if err := validateNewEc2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference{}

	_jsii_.Create(
		"awscc.ec2VerifiedAccessInstance.Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEc2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference_Override(e Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2VerifiedAccessInstance.Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetDeviceTrustProviderType(val *string) {
	if err := j.validateSetDeviceTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTrustProviderType",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetTrustProviderType(val *string) {
	if err := j.validateSetTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustProviderType",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetUserTrustProviderType(val *string) {
	if err := j.validateSetUserTrustProviderTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userTrustProviderType",
		val,
	)
}

func (j *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference)SetVerifiedAccessTrustProviderId(val *string) {
	if err := j.validateSetVerifiedAccessTrustProviderIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAccessTrustProviderId",
		val,
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		e,
		"resetDescription",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ResetDeviceTrustProviderType() {
	_jsii_.InvokeVoid(
		e,
		"resetDeviceTrustProviderType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ResetTrustProviderType() {
	_jsii_.InvokeVoid(
		e,
		"resetTrustProviderType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ResetUserTrustProviderType() {
	_jsii_.InvokeVoid(
		e,
		"resetUserTrustProviderType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ResetVerifiedAccessTrustProviderId() {
	_jsii_.InvokeVoid(
		e,
		"resetVerifiedAccessTrustProviderId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2VerifiedAccessInstanceVerifiedAccessTrustProvidersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

