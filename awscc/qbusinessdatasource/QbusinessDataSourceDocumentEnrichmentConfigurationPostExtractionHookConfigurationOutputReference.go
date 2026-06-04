package qbusinessdatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/qbusinessdatasource/internal"
)

type QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference interface {
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
	InvocationCondition() QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionOutputReference
	InvocationConditionInput() interface{}
	LambdaArn() *string
	SetLambdaArn(val *string)
	LambdaArnInput() *string
	RoleArn() *string
	SetRoleArn(val *string)
	RoleArnInput() *string
	S3BucketName() *string
	SetS3BucketName(val *string)
	S3BucketNameInput() *string
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
	PutInvocationCondition(value *QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationCondition)
	ResetInvocationCondition()
	ResetLambdaArn()
	ResetRoleArn()
	ResetS3BucketName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference
type jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) InvocationCondition() QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionOutputReference {
	var returns QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionOutputReference
	_jsii_.Get(
		j,
		"invocationCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) InvocationConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"invocationConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) LambdaArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) LambdaArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) S3BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) S3BucketNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BucketNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewQbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewQbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.qbusinessDataSource.QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewQbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference_Override(q QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.qbusinessDataSource.QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		q,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetLambdaArn(val *string) {
	if err := j.validateSetLambdaArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaArn",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetS3BucketName(val *string) {
	if err := j.validateSetS3BucketNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BucketName",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		q,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) PutInvocationCondition(value *QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationCondition) {
	if err := q.validatePutInvocationConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		q,
		"putInvocationCondition",
		[]interface{}{value},
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ResetInvocationCondition() {
	_jsii_.InvokeVoid(
		q,
		"resetInvocationCondition",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ResetLambdaArn() {
	_jsii_.InvokeVoid(
		q,
		"resetLambdaArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ResetRoleArn() {
	_jsii_.InvokeVoid(
		q,
		"resetRoleArn",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ResetS3BucketName() {
	_jsii_.InvokeVoid(
		q,
		"resetS3BucketName",
		nil, // no parameters
	)
}

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (q *jsiiProxy_QbusinessDataSourceDocumentEnrichmentConfigurationPostExtractionHookConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

