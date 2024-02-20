package dataawscciotaccountauditconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawscciotaccountauditconfiguration/internal"
)

type DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference interface {
	cdktf.ComplexObject
	AuthenticatedCognitoRoleOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
	CaCertificateExpiringCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheckOutputReference
	CaCertificateKeyQualityCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheckOutputReference
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
	ConflictingClientIdsCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheckOutputReference
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeviceCertificateExpiringCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheckOutputReference
	DeviceCertificateKeyQualityCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheckOutputReference
	DeviceCertificateSharedCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheckOutputReference
	// Experimental.
	Fqn() *string
	IntermediateCaRevokedForActiveDeviceCertificatesCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference
	InternalValue() *DataAwsccIotAccountAuditConfigurationAuditCheckConfigurations
	SetInternalValue(val *DataAwsccIotAccountAuditConfigurationAuditCheckConfigurations)
	IotPolicyOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheckOutputReference
	IoTPolicyPotentialMisConfigurationCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheckOutputReference
	IotRoleAliasAllowsAccessToUnusedServicesCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheckOutputReference
	IotRoleAliasOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheckOutputReference
	LoggingDisabledCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheckOutputReference
	RevokedCaCertificateStillActiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheckOutputReference
	RevokedDeviceCertificateStillActiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheckOutputReference
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UnauthenticatedCognitoRoleOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
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

// The jsii proxy struct for DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference
type jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) AuthenticatedCognitoRoleOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsAuthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"authenticatedCognitoRoleOverlyPermissiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CaCertificateExpiringCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateExpiringCheckOutputReference
	_jsii_.Get(
		j,
		"caCertificateExpiringCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CaCertificateKeyQualityCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsCaCertificateKeyQualityCheckOutputReference
	_jsii_.Get(
		j,
		"caCertificateKeyQualityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ConflictingClientIdsCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsConflictingClientIdsCheckOutputReference
	_jsii_.Get(
		j,
		"conflictingClientIdsCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateExpiringCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateExpiringCheckOutputReference
	_jsii_.Get(
		j,
		"deviceCertificateExpiringCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateKeyQualityCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateKeyQualityCheckOutputReference
	_jsii_.Get(
		j,
		"deviceCertificateKeyQualityCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) DeviceCertificateSharedCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsDeviceCertificateSharedCheckOutputReference
	_jsii_.Get(
		j,
		"deviceCertificateSharedCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IntermediateCaRevokedForActiveDeviceCertificatesCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIntermediateCaRevokedForActiveDeviceCertificatesCheckOutputReference
	_jsii_.Get(
		j,
		"intermediateCaRevokedForActiveDeviceCertificatesCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) InternalValue() *DataAwsccIotAccountAuditConfigurationAuditCheckConfigurations {
	var returns *DataAwsccIotAccountAuditConfigurationAuditCheckConfigurations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotPolicyOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotPolicyOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"iotPolicyOverlyPermissiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IoTPolicyPotentialMisConfigurationCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIoTPolicyPotentialMisConfigurationCheckOutputReference
	_jsii_.Get(
		j,
		"ioTPolicyPotentialMisConfigurationCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotRoleAliasAllowsAccessToUnusedServicesCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasAllowsAccessToUnusedServicesCheckOutputReference
	_jsii_.Get(
		j,
		"iotRoleAliasAllowsAccessToUnusedServicesCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) IotRoleAliasOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsIotRoleAliasOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"iotRoleAliasOverlyPermissiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) LoggingDisabledCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsLoggingDisabledCheckOutputReference
	_jsii_.Get(
		j,
		"loggingDisabledCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) RevokedCaCertificateStillActiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsRevokedCaCertificateStillActiveCheckOutputReference
	_jsii_.Get(
		j,
		"revokedCaCertificateStillActiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) RevokedDeviceCertificateStillActiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsRevokedDeviceCertificateStillActiveCheckOutputReference
	_jsii_.Get(
		j,
		"revokedDeviceCertificateStillActiveCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) UnauthenticatedCognitoRoleOverlyPermissiveCheck() DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheckOutputReference {
	var returns DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsUnauthenticatedCognitoRoleOverlyPermissiveCheckOutputReference
	_jsii_.Get(
		j,
		"unauthenticatedCognitoRoleOverlyPermissiveCheck",
		&returns,
	)
	return returns
}


func NewDataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference{}

	_jsii_.Create(
		"awscc.dataAwsccIotAccountAuditConfiguration.DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference_Override(d DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccIotAccountAuditConfiguration.DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetInternalValue(val *DataAwsccIotAccountAuditConfigurationAuditCheckConfigurations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAwsccIotAccountAuditConfigurationAuditCheckConfigurationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

