package elasticloadbalancingv2truststorerevocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/elasticloadbalancingv2truststorerevocation/internal"
)

type Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference interface {
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
	RevocationType() *string
	SetRevocationType(val *string)
	RevocationTypeInput() *string
	S3Bucket() *string
	SetS3Bucket(val *string)
	S3BucketInput() *string
	S3Key() *string
	SetS3Key(val *string)
	S3KeyInput() *string
	S3ObjectVersion() *string
	SetS3ObjectVersion(val *string)
	S3ObjectVersionInput() *string
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
	ResetRevocationType()
	ResetS3Bucket()
	ResetS3Key()
	ResetS3ObjectVersion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference
type jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) RevocationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) RevocationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) S3Bucket() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) S3BucketInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) S3Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3Key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) S3KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3KeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) S3ObjectVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) S3ObjectVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewElasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference {
	_init_.Initialize()

	if err := validateNewElasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference{}

	_jsii_.Create(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewElasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference_Override(e Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.elasticloadbalancingv2TrustStoreRevocation.Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetRevocationType(val *string) {
	if err := j.validateSetRevocationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revocationType",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetS3Bucket(val *string) {
	if err := j.validateSetS3BucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Bucket",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetS3Key(val *string) {
	if err := j.validateSetS3KeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3Key",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetS3ObjectVersion(val *string) {
	if err := j.validateSetS3ObjectVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersion",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ResetRevocationType() {
	_jsii_.InvokeVoid(
		e,
		"resetRevocationType",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ResetS3Bucket() {
	_jsii_.InvokeVoid(
		e,
		"resetS3Bucket",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ResetS3Key() {
	_jsii_.InvokeVoid(
		e,
		"resetS3Key",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ResetS3ObjectVersion() {
	_jsii_.InvokeVoid(
		e,
		"resetS3ObjectVersion",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Elasticloadbalancingv2TrustStoreRevocationRevocationContentsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

