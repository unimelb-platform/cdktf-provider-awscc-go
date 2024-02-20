package healthlakefhirdatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/healthlakefhirdatastore/internal"
)

type HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference interface {
	cdktf.ComplexObject
	CmkType() *string
	SetCmkType(val *string)
	CmkTypeInput() *string
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
	InternalValue() *HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfig
	SetInternalValue(val *HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfig)
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
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
	ResetKmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference
type jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) CmkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) CmkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cmkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) InternalValue() *HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfig {
	var returns *HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference {
	_init_.Initialize()

	if err := validateNewHealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference{}

	_jsii_.Create(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference_Override(h HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference)SetCmkType(val *string) {
	if err := j.validateSetCmkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cmkType",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference)SetInternalValue(val *HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		h,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreSseConfigurationKmsEncryptionConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

