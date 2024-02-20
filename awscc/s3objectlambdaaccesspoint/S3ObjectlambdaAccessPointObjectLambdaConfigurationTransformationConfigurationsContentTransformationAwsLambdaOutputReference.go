package s3objectlambdaaccesspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3objectlambdaaccesspoint/internal"
)

type S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference interface {
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
	FunctionArn() *string
	SetFunctionArn(val *string)
	FunctionArnInput() *string
	FunctionPayload() *string
	SetFunctionPayload(val *string)
	FunctionPayloadInput() *string
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
	ResetFunctionPayload()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference
type jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) FunctionArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) FunctionPayload() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionPayload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) FunctionPayloadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionPayloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewS3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference {
	_init_.Initialize()

	if err := validateNewS3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference{}

	_jsii_.Create(
		"awscc.s3ObjectlambdaAccessPoint.S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference_Override(s S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3ObjectlambdaAccessPoint.S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference)SetFunctionArn(val *string) {
	if err := j.validateSetFunctionArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionArn",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference)SetFunctionPayload(val *string) {
	if err := j.validateSetFunctionPayloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionPayload",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) ResetFunctionPayload() {
	_jsii_.InvokeVoid(
		s,
		"resetFunctionPayload",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsContentTransformationAwsLambdaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

