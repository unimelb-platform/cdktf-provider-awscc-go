package sagemakerproject

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/sagemakerproject/internal"
)

type SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList interface {
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
	Get(index *float64) SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList
type jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewSagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList {
	_init_.Initialize()

	if err := validateNewSagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList{}

	_jsii_.Create(
		"awscc.sagemakerProject.SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewSagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList_Override(s SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.sagemakerProject.SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		s,
	)
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (s *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) Get(index *float64) SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersOutputReference {
	if err := s.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersOutputReference

	_jsii_.Invoke(
		s,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SagemakerProjectServiceCatalogProvisioningDetailsProvisioningParametersList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

