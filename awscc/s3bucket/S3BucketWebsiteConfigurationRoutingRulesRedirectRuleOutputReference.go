package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3bucket/internal"
)

type S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference interface {
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
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	HttpRedirectCode() *string
	SetHttpRedirectCode(val *string)
	HttpRedirectCodeInput() *string
	InternalValue() *S3BucketWebsiteConfigurationRoutingRulesRedirectRule
	SetInternalValue(val *S3BucketWebsiteConfigurationRoutingRulesRedirectRule)
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	ReplaceKeyPrefixWith() *string
	SetReplaceKeyPrefixWith(val *string)
	ReplaceKeyPrefixWithInput() *string
	ReplaceKeyWith() *string
	SetReplaceKeyWith(val *string)
	ReplaceKeyWithInput() *string
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
	ResetHostName()
	ResetHttpRedirectCode()
	ResetProtocol()
	ResetReplaceKeyPrefixWith()
	ResetReplaceKeyWith()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference
type jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) HttpRedirectCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) HttpRedirectCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpRedirectCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) InternalValue() *S3BucketWebsiteConfigurationRoutingRulesRedirectRule {
	var returns *S3BucketWebsiteConfigurationRoutingRulesRedirectRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ReplaceKeyPrefixWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyPrefixWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ReplaceKeyPrefixWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyPrefixWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ReplaceKeyWith() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ReplaceKeyWithInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replaceKeyWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference{}

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference_Override(s S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetHttpRedirectCode(val *string) {
	if err := j.validateSetHttpRedirectCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRedirectCode",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetInternalValue(val *S3BucketWebsiteConfigurationRoutingRulesRedirectRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetReplaceKeyPrefixWith(val *string) {
	if err := j.validateSetReplaceKeyPrefixWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceKeyPrefixWith",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetReplaceKeyWith(val *string) {
	if err := j.validateSetReplaceKeyWithParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replaceKeyWith",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ResetHostName() {
	_jsii_.InvokeVoid(
		s,
		"resetHostName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ResetHttpRedirectCode() {
	_jsii_.InvokeVoid(
		s,
		"resetHttpRedirectCode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ResetProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ResetReplaceKeyPrefixWith() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceKeyPrefixWith",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ResetReplaceKeyWith() {
	_jsii_.InvokeVoid(
		s,
		"resetReplaceKeyWith",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationRoutingRulesRedirectRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

