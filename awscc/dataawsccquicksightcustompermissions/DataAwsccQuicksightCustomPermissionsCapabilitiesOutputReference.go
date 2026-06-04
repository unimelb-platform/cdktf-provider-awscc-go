package dataawsccquicksightcustompermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccquicksightcustompermissions/internal"
)

type DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference interface {
	cdktf.ComplexObject
	AddOrRunAnomalyDetectionForAnalyses() *string
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
	CreateAndUpdateDashboardEmailReports() *string
	CreateAndUpdateDatasets() *string
	CreateAndUpdateDataSources() *string
	CreateAndUpdateThemes() *string
	CreateAndUpdateThresholdAlerts() *string
	CreateSharedFolders() *string
	CreateSpiceDataset() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExportToCsv() *string
	ExportToCsvInScheduledReports() *string
	ExportToExcel() *string
	ExportToExcelInScheduledReports() *string
	ExportToPdf() *string
	ExportToPdfInScheduledReports() *string
	// Experimental.
	Fqn() *string
	IncludeContentInScheduledReportsEmail() *string
	InternalValue() *DataAwsccQuicksightCustomPermissionsCapabilities
	SetInternalValue(val *DataAwsccQuicksightCustomPermissionsCapabilities)
	PrintReports() *string
	RenameSharedFolders() *string
	ShareAnalyses() *string
	ShareDashboards() *string
	ShareDatasets() *string
	ShareDataSources() *string
	SubscribeDashboardEmailReports() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	ViewAccountSpiceCapacity() *string
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

// The jsii proxy struct for DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference
type jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) AddOrRunAnomalyDetectionForAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThemes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateAndUpdateThresholdAlerts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreateSpiceDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToCsvInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToExcelInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ExportToPdfInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) IncludeContentInScheduledReportsEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InternalValue() *DataAwsccQuicksightCustomPermissionsCapabilities {
	var returns *DataAwsccQuicksightCustomPermissionsCapabilities
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) PrintReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) RenameSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDashboards() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ShareDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) SubscribeDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ViewAccountSpiceCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacity",
		&returns,
	)
	return returns
}


func NewDataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccQuicksightCustomPermissionsCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccQuicksightCustomPermissions.DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference_Override(d DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccQuicksightCustomPermissions.DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetInternalValue(val *DataAwsccQuicksightCustomPermissionsCapabilities) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccQuicksightCustomPermissionsCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

