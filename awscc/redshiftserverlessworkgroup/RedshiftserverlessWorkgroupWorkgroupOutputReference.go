package redshiftserverlessworkgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/redshiftserverlessworkgroup/internal"
)

type RedshiftserverlessWorkgroupWorkgroupOutputReference interface {
	cdktf.ComplexObject
	BaseCapacity() *float64
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
	ConfigParameters() RedshiftserverlessWorkgroupWorkgroupConfigParametersList
	CreationDate() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Endpoint() RedshiftserverlessWorkgroupWorkgroupEndpointOutputReference
	EnhancedVpcRouting() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	InternalValue() *RedshiftserverlessWorkgroupWorkgroup
	SetInternalValue(val *RedshiftserverlessWorkgroupWorkgroup)
	NamespaceName() *string
	PubliclyAccessible() cdktf.IResolvable
	SecurityGroupIds() *[]*string
	Status() *string
	SubnetIds() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkgroupArn() *string
	WorkgroupId() *string
	WorkgroupName() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RedshiftserverlessWorkgroupWorkgroupOutputReference
type jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) BaseCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) ConfigParameters() RedshiftserverlessWorkgroupWorkgroupConfigParametersList {
	var returns RedshiftserverlessWorkgroupWorkgroupConfigParametersList
	_jsii_.Get(
		j,
		"configParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) CreationDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) Endpoint() RedshiftserverlessWorkgroupWorkgroupEndpointOutputReference {
	var returns RedshiftserverlessWorkgroupWorkgroupEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) EnhancedVpcRouting() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enhancedVpcRouting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) InternalValue() *RedshiftserverlessWorkgroupWorkgroup {
	var returns *RedshiftserverlessWorkgroupWorkgroup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) NamespaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) PubliclyAccessible() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) SecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) WorkgroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) WorkgroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) WorkgroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workgroupName",
		&returns,
	)
	return returns
}


func NewRedshiftserverlessWorkgroupWorkgroupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RedshiftserverlessWorkgroupWorkgroupOutputReference {
	_init_.Initialize()

	if err := validateNewRedshiftserverlessWorkgroupWorkgroupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference{}

	_jsii_.Create(
		"awscc.redshiftserverlessWorkgroup.RedshiftserverlessWorkgroupWorkgroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRedshiftserverlessWorkgroupWorkgroupOutputReference_Override(r RedshiftserverlessWorkgroupWorkgroupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.redshiftserverlessWorkgroup.RedshiftserverlessWorkgroupWorkgroupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference)SetInternalValue(val *RedshiftserverlessWorkgroupWorkgroup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RedshiftserverlessWorkgroupWorkgroupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

