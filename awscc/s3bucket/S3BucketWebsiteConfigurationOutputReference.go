package s3bucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3bucket/internal"
)

type S3BucketWebsiteConfigurationOutputReference interface {
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
	ErrorDocument() *string
	SetErrorDocument(val *string)
	ErrorDocumentInput() *string
	// Experimental.
	Fqn() *string
	IndexDocument() *string
	SetIndexDocument(val *string)
	IndexDocumentInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RedirectAllRequestsTo() S3BucketWebsiteConfigurationRedirectAllRequestsToOutputReference
	RedirectAllRequestsToInput() interface{}
	RoutingRules() S3BucketWebsiteConfigurationRoutingRulesList
	RoutingRulesInput() interface{}
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
	PutRedirectAllRequestsTo(value *S3BucketWebsiteConfigurationRedirectAllRequestsTo)
	PutRoutingRules(value interface{})
	ResetErrorDocument()
	ResetIndexDocument()
	ResetRedirectAllRequestsTo()
	ResetRoutingRules()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3BucketWebsiteConfigurationOutputReference
type jsiiProxy_S3BucketWebsiteConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ErrorDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ErrorDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"errorDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) IndexDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) IndexDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) RedirectAllRequestsTo() S3BucketWebsiteConfigurationRedirectAllRequestsToOutputReference {
	var returns S3BucketWebsiteConfigurationRedirectAllRequestsToOutputReference
	_jsii_.Get(
		j,
		"redirectAllRequestsTo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) RedirectAllRequestsToInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectAllRequestsToInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) RoutingRules() S3BucketWebsiteConfigurationRoutingRulesList {
	var returns S3BucketWebsiteConfigurationRoutingRulesList
	_jsii_.Get(
		j,
		"routingRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) RoutingRulesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"routingRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3BucketWebsiteConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3BucketWebsiteConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewS3BucketWebsiteConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3BucketWebsiteConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketWebsiteConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3BucketWebsiteConfigurationOutputReference_Override(s S3BucketWebsiteConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3Bucket.S3BucketWebsiteConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference)SetErrorDocument(val *string) {
	if err := j.validateSetErrorDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorDocument",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference)SetIndexDocument(val *string) {
	if err := j.validateSetIndexDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexDocument",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3BucketWebsiteConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) PutRedirectAllRequestsTo(value *S3BucketWebsiteConfigurationRedirectAllRequestsTo) {
	if err := s.validatePutRedirectAllRequestsToParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRedirectAllRequestsTo",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) PutRoutingRules(value interface{}) {
	if err := s.validatePutRoutingRulesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putRoutingRules",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ResetErrorDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetErrorDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ResetIndexDocument() {
	_jsii_.InvokeVoid(
		s,
		"resetIndexDocument",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ResetRedirectAllRequestsTo() {
	_jsii_.InvokeVoid(
		s,
		"resetRedirectAllRequestsTo",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ResetRoutingRules() {
	_jsii_.InvokeVoid(
		s,
		"resetRoutingRules",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3BucketWebsiteConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

