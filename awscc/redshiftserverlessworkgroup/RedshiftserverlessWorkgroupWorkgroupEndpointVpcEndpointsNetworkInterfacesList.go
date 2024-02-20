package redshiftserverlessworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/redshiftserverlessworkgroup/internal"
)

type RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList interface {
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
	Get(index *float64) RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList
type jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewRedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList {
	_init_.Initialize()

	if err := validateNewRedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList{}

	_jsii_.Create(
		"awscc.redshiftserverlessWorkgroup.RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewRedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList_Override(r RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.redshiftserverlessWorkgroup.RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		r,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) Get(index *float64) RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesOutputReference {
	if err := r.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesOutputReference

	_jsii_.Invoke(
		r,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupEndpointVpcEndpointsNetworkInterfacesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

