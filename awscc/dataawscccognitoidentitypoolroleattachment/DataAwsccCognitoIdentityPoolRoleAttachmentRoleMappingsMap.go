package dataawscccognitoidentitypoolroleattachment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscccognitoidentitypoolroleattachment/internal"
)

type DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap interface {
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
	Get(key *string) DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference
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

// The jsii proxy struct for DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap
type jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap {
	_init_.Initialize()

	if err := validateNewDataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap{}

	_jsii_.Create(
		"awscc.dataAwsccCognitoIdentityPoolRoleAttachment.DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap_Override(d DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccCognitoIdentityPoolRoleAttachment.DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) Get(key *string) DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference {
	if err := d.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccCognitoIdentityPoolRoleAttachmentRoleMappingsMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

