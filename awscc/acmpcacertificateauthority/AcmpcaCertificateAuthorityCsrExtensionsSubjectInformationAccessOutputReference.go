package acmpcacertificateauthority

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/acmpcacertificateauthority/internal"
)

type AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference interface {
	cdktf.ComplexObject
	AccessLocation() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference
	AccessLocationInput() *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation
	AccessMethod() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethodOutputReference
	AccessMethodInput() *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethod
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
	PutAccessLocation(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation)
	PutAccessMethod(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethod)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference
type jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) AccessLocation() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference {
	var returns AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocationOutputReference
	_jsii_.Get(
		j,
		"accessLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) AccessLocationInput() *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation {
	var returns *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation
	_jsii_.Get(
		j,
		"accessLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) AccessMethod() AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethodOutputReference {
	var returns AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethodOutputReference
	_jsii_.Get(
		j,
		"accessMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) AccessMethodInput() *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethod {
	var returns *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethod
	_jsii_.Get(
		j,
		"accessMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewAcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference{}

	_jsii_.Create(
		"awscc.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference_Override(a AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.acmpcaCertificateAuthority.AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) PutAccessLocation(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessLocation) {
	if err := a.validatePutAccessLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) PutAccessMethod(value *AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessAccessMethod) {
	if err := a.validatePutAccessMethodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAccessMethod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateAuthorityCsrExtensionsSubjectInformationAccessOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

