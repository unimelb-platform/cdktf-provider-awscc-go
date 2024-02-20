package s3objectlambdaaccesspoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/s3objectlambdaaccesspoint/internal"
)

type S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList
type jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewS3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList {
	_init_.Initialize()

	if err := validateNewS3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList{}

	_jsii_.Create(
		"awscc.s3ObjectlambdaAccessPoint.S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewS3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList_Override(s S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.s3ObjectlambdaAccessPoint.S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) Get(index *float64) S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_S3ObjectlambdaAccessPointObjectLambdaConfigurationTransformationConfigurationsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

