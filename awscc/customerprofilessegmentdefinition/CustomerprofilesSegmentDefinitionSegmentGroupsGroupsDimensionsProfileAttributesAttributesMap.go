package customerprofilessegmentdefinition

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/customerprofilessegmentdefinition/internal"
)

type CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap interface {
	cdktf.ComplexMap
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
	// Experimental.
	ComputeFqn() *string
	Get(key *string) CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesOutputReference
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

// The jsii proxy struct for CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap
type jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap struct {
	internal.Type__cdktfComplexMap
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap {
	_init_.Initialize()

	if err := validateNewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMapParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap{}

	_jsii_.Create(
		"awscc.customerprofilesSegmentDefinition.CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap_Override(c CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.customerprofilesSegmentDefinition.CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) Get(key *string) CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesOutputReference {
	if err := c.validateGetParameters(key); err != nil {
		panic(err)
	}
	var returns CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesOutputReference

	_jsii_.Invoke(
		c,
		"get",
		[]interface{}{key},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CustomerprofilesSegmentDefinitionSegmentGroupsGroupsDimensionsProfileAttributesAttributesMap) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

