package redshiftserverlessworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/redshiftserverlessworkgroup/internal"
)

type RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
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
	Get(index *float64) RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList
type jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewRedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList {
	_init_.Initialize()

	if err := validateNewRedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList{}

	_jsii_.Create(
		"awscc.redshiftserverlessWorkgroup.RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewRedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList_Override(r RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.redshiftserverlessWorkgroup.RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) Get(index *float64) RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

