package dataawsccresiliencehubresiliencypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccresiliencehubresiliencypolicy/internal"
)

type DataAwsccResiliencehubResiliencyPolicyPolicyMap interface {
	cdktf.ComplexMap
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
	// Experimental.
	ComputeFqn() *string
	Get(key *string) DataAwsccResiliencehubResiliencyPolicyPolicyOutputReference
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccResiliencehubResiliencyPolicyPolicyMap
type jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccResiliencehubResiliencyPolicyPolicyMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccResiliencehubResiliencyPolicyPolicyMap {
	_init_.Initialize()

	if err := validateNewDataAwsccResiliencehubResiliencyPolicyPolicyMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap{}

	_jsii_.Create(
		"awscc.dataAwsccResiliencehubResiliencyPolicy.DataAwsccResiliencehubResiliencyPolicyPolicyMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccResiliencehubResiliencyPolicyPolicyMap_Override(d DataAwsccResiliencehubResiliencyPolicyPolicyMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccResiliencehubResiliencyPolicy.DataAwsccResiliencehubResiliencyPolicyPolicyMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) Get(key *string) DataAwsccResiliencehubResiliencyPolicyPolicyOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccResiliencehubResiliencyPolicyPolicyOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccResiliencehubResiliencyPolicyPolicyMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

