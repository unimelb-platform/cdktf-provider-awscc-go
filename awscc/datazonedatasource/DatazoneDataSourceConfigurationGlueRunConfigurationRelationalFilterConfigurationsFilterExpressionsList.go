package datazonedatasource

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/datazonedatasource/internal"
)

type DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList interface {
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
	Get(index *float64) DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList
type jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList {
	_init_.Initialize()

	if err := validateNewDatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList{}

	_jsii_.Create(
		"awscc.datazoneDataSource.DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList_Override(d DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.datazoneDataSource.DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) Get(index *float64) DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatazoneDataSourceConfigurationGlueRunConfigurationRelationalFilterConfigurationsFilterExpressionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

