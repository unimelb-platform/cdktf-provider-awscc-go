package ec2spotfleet

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/ec2spotfleet/internal"
)

type Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList interface {
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
	Get(index *float64) Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList
type jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList {
	_init_.Initialize()

	if err := validateNewEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList{}

	_jsii_.Create(
		"awscc.ec2SpotFleet.Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewEc2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList_Override(e Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.ec2SpotFleet.Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		e,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) Get(index *float64) Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsOutputReference {
	if err := e.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsOutputReference

	_jsii_.Invoke(
		e,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_Ec2SpotFleetSpotFleetRequestConfigDataLaunchSpecificationsTagSpecificationsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

