package dataawsccec2launchtemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccec2launchtemplate/internal"
)

type DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList interface {
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
	Get(index *float64) DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList
type jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList {
	_init_.Initialize()

	if err := validateNewDataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList{}

	_jsii_.Create(
		"awscc.dataAwsccEc2LaunchTemplate.DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList_Override(d DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccEc2LaunchTemplate.DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) Get(index *float64) DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccEc2LaunchTemplateLaunchTemplateDataNetworkInterfacesIpv4PrefixesList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
