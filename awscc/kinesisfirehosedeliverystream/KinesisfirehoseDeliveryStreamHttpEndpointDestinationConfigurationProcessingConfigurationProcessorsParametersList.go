package kinesisfirehosedeliverystream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/kinesisfirehosedeliverystream/internal"
)

type KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList interface {
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
	Get(index *float64) KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList
type jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList {
	_init_.Initialize()

	if err := validateNewKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList{}

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewKinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList_Override(k KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.kinesisfirehoseDeliveryStream.KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		k,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) Get(index *float64) KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference {
	if err := k.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersOutputReference

	_jsii_.Invoke(
		k,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisfirehoseDeliveryStreamHttpEndpointDestinationConfigurationProcessingConfigurationProcessorsParametersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

