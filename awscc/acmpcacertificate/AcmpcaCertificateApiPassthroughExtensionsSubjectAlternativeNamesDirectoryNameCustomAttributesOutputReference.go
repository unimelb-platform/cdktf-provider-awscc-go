package acmpcacertificate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/acmpcacertificate/internal"
)

type AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ObjectIdentifier() *string
	SetObjectIdentifier(val *string)
	ObjectIdentifierInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Value() *string
	SetValue(val *string)
	ValueInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference
type jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) ObjectIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) ObjectIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"objectIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) ValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueInput",
		&returns,
	)
	return returns
}


func NewAcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference {
	_init_.Initialize()

	if err := validateNewAcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference{}

	_jsii_.Create(
		"awscc.acmpcaCertificate.AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewAcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference_Override(a AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.acmpcaCertificate.AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference)SetObjectIdentifier(val *string) {
	if err := j.validateSetObjectIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectIdentifier",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference)SetValue(val *string) {
	if err := j.validateSetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"value",
		val,
	)
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AcmpcaCertificateApiPassthroughExtensionsSubjectAlternativeNamesDirectoryNameCustomAttributesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

