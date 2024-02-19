package s3objectlambdaaccesspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3objectlambdaaccesspoint/internal"
)

type S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference interface {
	cdktf.ComplexObject
	AllowedFeatures() *[]*string
	SetAllowedFeatures(val *[]*string)
	AllowedFeaturesInput() *[]*string
	CloudwatchMetricsEnabled() interface{}
	SetCloudwatchMetricsEnabled(val interface{})
	CloudwatchMetricsEnabledInput() interface{}
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
	SupportingAccessPoint() *string
	SetSupportingAccessPoint(val *string)
	SupportingAccessPointInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TransformationConfigurations() S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList
	TransformationConfigurationsInput() interface{}
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
	PutTransformationConfigurations(value interface{})
	ResetAllowedFeatures()
	ResetCloudwatchMetricsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference
type jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) AllowedFeatures() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) AllowedFeaturesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) CloudwatchMetricsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchMetricsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) CloudwatchMetricsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchMetricsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) SupportingAccessPoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportingAccessPoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) SupportingAccessPointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"supportingAccessPointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) TransformationConfigurations() S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList {
	var returns S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList
	_jsii_.Get(
		j,
		"transformationConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) TransformationConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"transformationConfigurationsInput",
		&returns,
	)
	return returns
}


func NewS3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewS3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference{}

	_jsii_.Create(
		"awscc.s3ObjectlambdaAccessPoint.S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewS3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference_Override(s S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3ObjectlambdaAccessPoint.S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetAllowedFeatures(val *[]*string) {
	if err := j.validateSetAllowedFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedFeatures",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetCloudwatchMetricsEnabled(val interface{}) {
	if err := j.validateSetCloudwatchMetricsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudwatchMetricsEnabled",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetSupportingAccessPoint(val *string) {
	if err := j.validateSetSupportingAccessPointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"supportingAccessPoint",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) PutTransformationConfigurations(value interface{}) {
	if err := s.validatePutTransformationConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTransformationConfigurations",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) ResetAllowedFeatures() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedFeatures",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) ResetCloudwatchMetricsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetCloudwatchMetricsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
