package quicksightdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksightdatasource/internal"
)

type QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference interface {
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
	IdentityProviderResourceUri() *string
	SetIdentityProviderResourceUri(val *string)
	IdentityProviderResourceUriInput() *string
	IdentityProviderVpcConnectionProperties() QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersIdentityProviderVpcConnectionPropertiesOutputReference
	IdentityProviderVpcConnectionPropertiesInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	OAuthScope() *string
	SetOAuthScope(val *string)
	OAuthScopeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenProviderUrl() *string
	SetTokenProviderUrl(val *string)
	TokenProviderUrlInput() *string
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
	PutIdentityProviderVpcConnectionProperties(value *QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersIdentityProviderVpcConnectionProperties)
	ResetIdentityProviderResourceUri()
	ResetIdentityProviderVpcConnectionProperties()
	ResetOAuthScope()
	ResetTokenProviderUrl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference
type jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) IdentityProviderResourceUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderResourceUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) IdentityProviderResourceUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderResourceUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) IdentityProviderVpcConnectionProperties() QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersIdentityProviderVpcConnectionPropertiesOutputReference {
	var returns QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersIdentityProviderVpcConnectionPropertiesOutputReference
	_jsii_.Get(
		j,
		"identityProviderVpcConnectionProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) IdentityProviderVpcConnectionPropertiesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"identityProviderVpcConnectionPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) OAuthScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuthScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) OAuthScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oAuthScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) TokenProviderUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenProviderUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) TokenProviderUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenProviderUrlInput",
		&returns,
	)
	return returns
}


func NewQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference {
	_init_.Initialize()

	if err := validateNewQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference{}

	_jsii_.Create(
		"awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference_Override(q QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightDataSource.QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetIdentityProviderResourceUri(val *string) {
	if err := j.validateSetIdentityProviderResourceUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderResourceUri",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetOAuthScope(val *string) {
	if err := j.validateSetOAuthScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oAuthScope",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference)SetTokenProviderUrl(val *string) {
	if err := j.validateSetTokenProviderUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenProviderUrl",
		val,
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := q.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		q,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := q.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := q.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		q,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := q.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		q,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := q.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		q,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := q.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		q,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := q.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		q,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := q.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		q,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := q.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		q,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := q.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) PutIdentityProviderVpcConnectionProperties(value *QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersIdentityProviderVpcConnectionProperties) {
	if err := q.validatePutIdentityProviderVpcConnectionPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putIdentityProviderVpcConnectionProperties",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ResetIdentityProviderResourceUri() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityProviderResourceUri",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ResetIdentityProviderVpcConnectionProperties() {
	_jsii_.InvokeVoid(
		q,
		"resetIdentityProviderVpcConnectionProperties",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ResetOAuthScope() {
	_jsii_.InvokeVoid(
		q,
		"resetOAuthScope",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ResetTokenProviderUrl() {
	_jsii_.InvokeVoid(
		q,
		"resetTokenProviderUrl",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightDataSourceAlternateDataSourceParametersSnowflakeParametersOAuthParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

