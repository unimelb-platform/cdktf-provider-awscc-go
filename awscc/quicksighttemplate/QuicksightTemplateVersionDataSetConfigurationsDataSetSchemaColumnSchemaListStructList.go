package quicksighttemplate

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/quicksighttemplate/internal"
)

type QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList interface {
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
	Get(index *float64) QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList
type jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewQuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList {
	_init_.Initialize()

	if err := validateNewQuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList{}

	_jsii_.Create(
		"awscc.quicksightTemplate.QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewQuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList_Override(q QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.quicksightTemplate.QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		q,
	)
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (q *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) Get(index *float64) QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructOutputReference {
	if err := q.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructOutputReference

	_jsii_.Invoke(
		q,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := q.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		q,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QuicksightTemplateVersionDataSetConfigurationsDataSetSchemaColumnSchemaListStructList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		q,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

