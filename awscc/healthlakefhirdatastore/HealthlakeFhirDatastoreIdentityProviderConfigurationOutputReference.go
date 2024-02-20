package healthlakefhirdatastore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/healthlakefhirdatastore/internal"
)

type HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference interface {
	cdktf.ComplexObject
	AuthorizationStrategy() *string
	SetAuthorizationStrategy(val *string)
	AuthorizationStrategyInput() *string
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
	FineGrainedAuthorizationEnabled() interface{}
	SetFineGrainedAuthorizationEnabled(val interface{})
	FineGrainedAuthorizationEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	IdpLambdaArn() *string
	SetIdpLambdaArn(val *string)
	IdpLambdaArnInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Metadata() *string
	SetMetadata(val *string)
	MetadataInput() *string
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
	ResetFineGrainedAuthorizationEnabled()
	ResetIdpLambdaArn()
	ResetMetadata()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference
type jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) AuthorizationStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) AuthorizationStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) FineGrainedAuthorizationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fineGrainedAuthorizationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) FineGrainedAuthorizationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fineGrainedAuthorizationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) IdpLambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpLambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) IdpLambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idpLambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewHealthlakeFhirDatastoreIdentityProviderConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference_Override(h HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.healthlakeFhirDatastore.HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetAuthorizationStrategy(val *string) {
	if err := j.validateSetAuthorizationStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationStrategy",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetFineGrainedAuthorizationEnabled(val interface{}) {
	if err := j.validateSetFineGrainedAuthorizationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fineGrainedAuthorizationEnabled",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetIdpLambdaArn(val *string) {
	if err := j.validateSetIdpLambdaArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idpLambdaArn",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetMetadata(val *string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) ResetFineGrainedAuthorizationEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetFineGrainedAuthorizationEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) ResetIdpLambdaArn() {
	_jsii_.InvokeVoid(
		h,
		"resetIdpLambdaArn",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) ResetMetadata() {
	_jsii_.InvokeVoid(
		h,
		"resetMetadata",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HealthlakeFhirDatastoreIdentityProviderConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

