package dataawscckendraindex

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscckendraindex/internal"
)

type DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList interface {
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
	Get(index *float64) DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList
type jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList {
	_init_.Initialize()

	if err := validateNewDataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList{}

	_jsii_.Create(
		"awscc.dataAwsccKendraIndex.DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList_Override(d DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccKendraIndex.DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) Get(index *float64) DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsOutputReference {
	if err := d.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccKendraIndexDocumentMetadataConfigurationsRelevanceValueImportanceItemsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

