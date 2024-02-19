package rdsdbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/rdsdbinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance awscc_rds_db_instance}.
type RdsDbInstance interface {
	cdktf.TerraformResource
	AllocatedStorage() *string
	SetAllocatedStorage(val *string)
	AllocatedStorageInput() *string
	AllowMajorVersionUpgrade() interface{}
	SetAllowMajorVersionUpgrade(val interface{})
	AllowMajorVersionUpgradeInput() interface{}
	AssociatedRoles() RdsDbInstanceAssociatedRolesList
	AssociatedRolesInput() interface{}
	AutomaticBackupReplicationRegion() *string
	SetAutomaticBackupReplicationRegion(val *string)
	AutomaticBackupReplicationRegionInput() *string
	AutoMinorVersionUpgrade() interface{}
	SetAutoMinorVersionUpgrade(val interface{})
	AutoMinorVersionUpgradeInput() interface{}
	AvailabilityZone() *string
	SetAvailabilityZone(val *string)
	AvailabilityZoneInput() *string
	BackupRetentionPeriod() *float64
	SetBackupRetentionPeriod(val *float64)
	BackupRetentionPeriodInput() *float64
	CaCertificateIdentifier() *string
	SetCaCertificateIdentifier(val *string)
	CaCertificateIdentifierInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateDetails() RdsDbInstanceCertificateDetailsOutputReference
	CertificateDetailsInput() interface{}
	CertificateRotationRestart() interface{}
	SetCertificateRotationRestart(val interface{})
	CertificateRotationRestartInput() interface{}
	CharacterSetName() *string
	SetCharacterSetName(val *string)
	CharacterSetNameInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToSnapshot() interface{}
	SetCopyTagsToSnapshot(val interface{})
	CopyTagsToSnapshotInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomIamInstanceProfile() *string
	SetCustomIamInstanceProfile(val *string)
	CustomIamInstanceProfileInput() *string
	DbClusterIdentifier() *string
	SetDbClusterIdentifier(val *string)
	DbClusterIdentifierInput() *string
	DbClusterSnapshotIdentifier() *string
	SetDbClusterSnapshotIdentifier(val *string)
	DbClusterSnapshotIdentifierInput() *string
	DbInstanceArn() *string
	DbInstanceClass() *string
	SetDbInstanceClass(val *string)
	DbInstanceClassInput() *string
	DbInstanceIdentifier() *string
	SetDbInstanceIdentifier(val *string)
	DbInstanceIdentifierInput() *string
	DbiResourceId() *string
	DbName() *string
	SetDbName(val *string)
	DbNameInput() *string
	DbParameterGroupName() *string
	SetDbParameterGroupName(val *string)
	DbParameterGroupNameInput() *string
	DbSecurityGroups() *[]*string
	SetDbSecurityGroups(val *[]*string)
	DbSecurityGroupsInput() *[]*string
	DbSnapshotIdentifier() *string
	SetDbSnapshotIdentifier(val *string)
	DbSnapshotIdentifierInput() *string
	DbSubnetGroupName() *string
	SetDbSubnetGroupName(val *string)
	DbSubnetGroupNameInput() *string
	DbSystemId() *string
	DedicatedLogVolume() interface{}
	SetDedicatedLogVolume(val interface{})
	DedicatedLogVolumeInput() interface{}
	DeleteAutomatedBackups() interface{}
	SetDeleteAutomatedBackups(val interface{})
	DeleteAutomatedBackupsInput() interface{}
	DeletionProtection() interface{}
	SetDeletionProtection(val interface{})
	DeletionProtectionInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	SetDomain(val *string)
	DomainAuthSecretArn() *string
	SetDomainAuthSecretArn(val *string)
	DomainAuthSecretArnInput() *string
	DomainDnsIps() *[]*string
	SetDomainDnsIps(val *[]*string)
	DomainDnsIpsInput() *[]*string
	DomainFqdn() *string
	SetDomainFqdn(val *string)
	DomainFqdnInput() *string
	DomainIamRoleName() *string
	SetDomainIamRoleName(val *string)
	DomainIamRoleNameInput() *string
	DomainInput() *string
	DomainOu() *string
	SetDomainOu(val *string)
	DomainOuInput() *string
	EnableCloudwatchLogsExports() *[]*string
	SetEnableCloudwatchLogsExports(val *[]*string)
	EnableCloudwatchLogsExportsInput() *[]*string
	EnableIamDatabaseAuthentication() interface{}
	SetEnableIamDatabaseAuthentication(val interface{})
	EnableIamDatabaseAuthenticationInput() interface{}
	EnablePerformanceInsights() interface{}
	SetEnablePerformanceInsights(val interface{})
	EnablePerformanceInsightsInput() interface{}
	Endpoint() RdsDbInstanceEndpointOutputReference
	EndpointInput() interface{}
	Engine() *string
	SetEngine(val *string)
	EngineInput() *string
	EngineVersion() *string
	SetEngineVersion(val *string)
	EngineVersionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	Iops() *float64
	SetIops(val *float64)
	IopsInput() *float64
	KmsKeyId() *string
	SetKmsKeyId(val *string)
	KmsKeyIdInput() *string
	LicenseModel() *string
	SetLicenseModel(val *string)
	LicenseModelInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageMasterUserPassword() interface{}
	SetManageMasterUserPassword(val interface{})
	ManageMasterUserPasswordInput() interface{}
	MasterUsername() *string
	SetMasterUsername(val *string)
	MasterUsernameInput() *string
	MasterUserPassword() *string
	SetMasterUserPassword(val *string)
	MasterUserPasswordInput() *string
	MasterUserSecret() RdsDbInstanceMasterUserSecretOutputReference
	MasterUserSecretInput() interface{}
	MaxAllocatedStorage() *float64
	SetMaxAllocatedStorage(val *float64)
	MaxAllocatedStorageInput() *float64
	MonitoringInterval() *float64
	SetMonitoringInterval(val *float64)
	MonitoringIntervalInput() *float64
	MonitoringRoleArn() *string
	SetMonitoringRoleArn(val *string)
	MonitoringRoleArnInput() *string
	MultiAz() interface{}
	SetMultiAz(val interface{})
	MultiAzInput() interface{}
	NcharCharacterSetName() *string
	SetNcharCharacterSetName(val *string)
	NcharCharacterSetNameInput() *string
	NetworkType() *string
	SetNetworkType(val *string)
	NetworkTypeInput() *string
	// The tree node.
	Node() constructs.Node
	OptionGroupName() *string
	SetOptionGroupName(val *string)
	OptionGroupNameInput() *string
	PerformanceInsightsKmsKeyId() *string
	SetPerformanceInsightsKmsKeyId(val *string)
	PerformanceInsightsKmsKeyIdInput() *string
	PerformanceInsightsRetentionPeriod() *float64
	SetPerformanceInsightsRetentionPeriod(val *float64)
	PerformanceInsightsRetentionPeriodInput() *float64
	Port() *string
	SetPort(val *string)
	PortInput() *string
	PreferredBackupWindow() *string
	SetPreferredBackupWindow(val *string)
	PreferredBackupWindowInput() *string
	PreferredMaintenanceWindow() *string
	SetPreferredMaintenanceWindow(val *string)
	PreferredMaintenanceWindowInput() *string
	ProcessorFeatures() RdsDbInstanceProcessorFeaturesList
	ProcessorFeaturesInput() interface{}
	PromotionTier() *float64
	SetPromotionTier(val *float64)
	PromotionTierInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PubliclyAccessible() interface{}
	SetPubliclyAccessible(val interface{})
	PubliclyAccessibleInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReplicaMode() *string
	SetReplicaMode(val *string)
	ReplicaModeInput() *string
	RestoreTime() *string
	SetRestoreTime(val *string)
	RestoreTimeInput() *string
	SourceDbClusterIdentifier() *string
	SetSourceDbClusterIdentifier(val *string)
	SourceDbClusterIdentifierInput() *string
	SourceDbInstanceAutomatedBackupsArn() *string
	SetSourceDbInstanceAutomatedBackupsArn(val *string)
	SourceDbInstanceAutomatedBackupsArnInput() *string
	SourceDbInstanceIdentifier() *string
	SetSourceDbInstanceIdentifier(val *string)
	SourceDbInstanceIdentifierInput() *string
	SourceDbiResourceId() *string
	SetSourceDbiResourceId(val *string)
	SourceDbiResourceIdInput() *string
	SourceRegion() *string
	SetSourceRegion(val *string)
	SourceRegionInput() *string
	StorageEncrypted() interface{}
	SetStorageEncrypted(val interface{})
	StorageEncryptedInput() interface{}
	StorageThroughput() *float64
	SetStorageThroughput(val *float64)
	StorageThroughputInput() *float64
	StorageType() *string
	SetStorageType(val *string)
	StorageTypeInput() *string
	Tags() RdsDbInstanceTagsList
	TagsInput() interface{}
	TdeCredentialArn() *string
	SetTdeCredentialArn(val *string)
	TdeCredentialArnInput() *string
	TdeCredentialPassword() *string
	SetTdeCredentialPassword(val *string)
	TdeCredentialPasswordInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	UseDefaultProcessorFeatures() interface{}
	SetUseDefaultProcessorFeatures(val interface{})
	UseDefaultProcessorFeaturesInput() interface{}
	UseLatestRestorableTime() interface{}
	SetUseLatestRestorableTime(val interface{})
	UseLatestRestorableTimeInput() interface{}
	VpcSecurityGroups() *[]*string
	SetVpcSecurityGroups(val *[]*string)
	VpcSecurityGroupsInput() *[]*string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAssociatedRoles(value interface{})
	PutCertificateDetails(value *RdsDbInstanceCertificateDetails)
	PutEndpoint(value *RdsDbInstanceEndpoint)
	PutMasterUserSecret(value *RdsDbInstanceMasterUserSecret)
	PutProcessorFeatures(value interface{})
	PutTags(value interface{})
	ResetAllocatedStorage()
	ResetAllowMajorVersionUpgrade()
	ResetAssociatedRoles()
	ResetAutomaticBackupReplicationRegion()
	ResetAutoMinorVersionUpgrade()
	ResetAvailabilityZone()
	ResetBackupRetentionPeriod()
	ResetCaCertificateIdentifier()
	ResetCertificateDetails()
	ResetCertificateRotationRestart()
	ResetCharacterSetName()
	ResetCopyTagsToSnapshot()
	ResetCustomIamInstanceProfile()
	ResetDbClusterIdentifier()
	ResetDbClusterSnapshotIdentifier()
	ResetDbInstanceClass()
	ResetDbInstanceIdentifier()
	ResetDbName()
	ResetDbParameterGroupName()
	ResetDbSecurityGroups()
	ResetDbSnapshotIdentifier()
	ResetDbSubnetGroupName()
	ResetDedicatedLogVolume()
	ResetDeleteAutomatedBackups()
	ResetDeletionProtection()
	ResetDomain()
	ResetDomainAuthSecretArn()
	ResetDomainDnsIps()
	ResetDomainFqdn()
	ResetDomainIamRoleName()
	ResetDomainOu()
	ResetEnableCloudwatchLogsExports()
	ResetEnableIamDatabaseAuthentication()
	ResetEnablePerformanceInsights()
	ResetEndpoint()
	ResetEngine()
	ResetEngineVersion()
	ResetIops()
	ResetKmsKeyId()
	ResetLicenseModel()
	ResetManageMasterUserPassword()
	ResetMasterUsername()
	ResetMasterUserPassword()
	ResetMasterUserSecret()
	ResetMaxAllocatedStorage()
	ResetMonitoringInterval()
	ResetMonitoringRoleArn()
	ResetMultiAz()
	ResetNcharCharacterSetName()
	ResetNetworkType()
	ResetOptionGroupName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPerformanceInsightsKmsKeyId()
	ResetPerformanceInsightsRetentionPeriod()
	ResetPort()
	ResetPreferredBackupWindow()
	ResetPreferredMaintenanceWindow()
	ResetProcessorFeatures()
	ResetPromotionTier()
	ResetPubliclyAccessible()
	ResetReplicaMode()
	ResetRestoreTime()
	ResetSourceDbClusterIdentifier()
	ResetSourceDbInstanceAutomatedBackupsArn()
	ResetSourceDbInstanceIdentifier()
	ResetSourceDbiResourceId()
	ResetSourceRegion()
	ResetStorageEncrypted()
	ResetStorageThroughput()
	ResetStorageType()
	ResetTags()
	ResetTdeCredentialArn()
	ResetTdeCredentialPassword()
	ResetTimezone()
	ResetUseDefaultProcessorFeatures()
	ResetUseLatestRestorableTime()
	ResetVpcSecurityGroups()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RdsDbInstance
type jsiiProxy_RdsDbInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RdsDbInstance) AllocatedStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AllocatedStorageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AllowMajorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AllowMajorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMajorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AssociatedRoles() RdsDbInstanceAssociatedRolesList {
	var returns RdsDbInstanceAssociatedRolesList
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AssociatedRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"associatedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutomaticBackupReplicationRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutoMinorVersionUpgrade() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AutoMinorVersionUpgradeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoMinorVersionUpgradeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) AvailabilityZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) BackupRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CaCertificateIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CertificateDetails() RdsDbInstanceCertificateDetailsOutputReference {
	var returns RdsDbInstanceCertificateDetailsOutputReference
	_jsii_.Get(
		j,
		"certificateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CertificateDetailsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateDetailsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CertificateRotationRestart() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateRotationRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CertificateRotationRestartInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"certificateRotationRestartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CopyTagsToSnapshot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CopyTagsToSnapshotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CustomIamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) CustomIamInstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbClusterSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbInstanceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbParameterGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSnapshotIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSubnetGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DedicatedLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DedicatedLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dedicatedLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeleteAutomatedBackups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeleteAutomatedBackupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteAutomatedBackupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeletionProtection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DeletionProtectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainAuthSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainAuthSecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainDnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainDnsIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainFqdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainIamRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) DomainOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableCloudwatchLogsExportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableIamDatabaseAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnableIamDatabaseAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnablePerformanceInsights() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EnablePerformanceInsightsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePerformanceInsightsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Endpoint() RdsDbInstanceEndpointOutputReference {
	var returns RdsDbInstanceEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"endpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) EngineVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) IopsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iopsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) KmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) LicenseModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ManageMasterUserPassword() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ManageMasterUserPasswordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"manageMasterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserSecret() RdsDbInstanceMasterUserSecretOutputReference {
	var returns RdsDbInstanceMasterUserSecretOutputReference
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MasterUserSecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"masterUserSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MaxAllocatedStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MonitoringRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MultiAz() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) MultiAzInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiAzInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NcharCharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NcharCharacterSetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) NetworkTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) OptionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PerformanceInsightsRetentionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredBackupWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PreferredMaintenanceWindowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ProcessorFeatures() RdsDbInstanceProcessorFeaturesList {
	var returns RdsDbInstanceProcessorFeaturesList
	_jsii_.Get(
		j,
		"processorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ProcessorFeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"processorFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PromotionTier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PromotionTierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionTierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PubliclyAccessible() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) PubliclyAccessibleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"publiclyAccessibleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ReplicaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) ReplicaModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) RestoreTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) RestoreTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbClusterIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceAutomatedBackupsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceAutomatedBackupsArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbInstanceIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceDbiResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) SourceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageEncrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageEncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageEncryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageThroughputInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughputInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) StorageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Tags() RdsDbInstanceTagsList {
	var returns RdsDbInstanceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TdeCredentialPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseDefaultProcessorFeatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultProcessorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseDefaultProcessorFeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useDefaultProcessorFeaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseLatestRestorableTime() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) UseLatestRestorableTimeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useLatestRestorableTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) VpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RdsDbInstance) VpcSecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroupsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance awscc_rds_db_instance} Resource.
func NewRdsDbInstance(scope constructs.Construct, id *string, config *RdsDbInstanceConfig) RdsDbInstance {
	_init_.Initialize()

	if err := validateNewRdsDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RdsDbInstance{}

	_jsii_.Create(
		"awscc.rdsDbInstance.RdsDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/rds_db_instance awscc_rds_db_instance} Resource.
func NewRdsDbInstance_Override(r RdsDbInstance, scope constructs.Construct, id *string, config *RdsDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.rdsDbInstance.RdsDbInstance",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAllocatedStorage(val *string) {
	if err := j.validateSetAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocatedStorage",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAllowMajorVersionUpgrade(val interface{}) {
	if err := j.validateSetAllowMajorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowMajorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAutomaticBackupReplicationRegion(val *string) {
	if err := j.validateSetAutomaticBackupReplicationRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticBackupReplicationRegion",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAutoMinorVersionUpgrade(val interface{}) {
	if err := j.validateSetAutoMinorVersionUpgradeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMinorVersionUpgrade",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetAvailabilityZone(val *string) {
	if err := j.validateSetAvailabilityZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilityZone",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetBackupRetentionPeriod(val *float64) {
	if err := j.validateSetBackupRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backupRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCaCertificateIdentifier(val *string) {
	if err := j.validateSetCaCertificateIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caCertificateIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCertificateRotationRestart(val interface{}) {
	if err := j.validateSetCertificateRotationRestartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateRotationRestart",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCharacterSetName(val *string) {
	if err := j.validateSetCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"characterSetName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCopyTagsToSnapshot(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshot",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetCustomIamInstanceProfile(val *string) {
	if err := j.validateSetCustomIamInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customIamInstanceProfile",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbClusterIdentifier(val *string) {
	if err := j.validateSetDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbClusterSnapshotIdentifier(val *string) {
	if err := j.validateSetDbClusterSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbClusterSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbInstanceClass(val *string) {
	if err := j.validateSetDbInstanceClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceClass",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbInstanceIdentifier(val *string) {
	if err := j.validateSetDbInstanceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbName(val *string) {
	if err := j.validateSetDbNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbParameterGroupName(val *string) {
	if err := j.validateSetDbParameterGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbParameterGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbSecurityGroups(val *[]*string) {
	if err := j.validateSetDbSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbSnapshotIdentifier(val *string) {
	if err := j.validateSetDbSnapshotIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSnapshotIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDbSubnetGroupName(val *string) {
	if err := j.validateSetDbSubnetGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dbSubnetGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDedicatedLogVolume(val interface{}) {
	if err := j.validateSetDedicatedLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedLogVolume",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDeleteAutomatedBackups(val interface{}) {
	if err := j.validateSetDeleteAutomatedBackupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteAutomatedBackups",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDeletionProtection(val interface{}) {
	if err := j.validateSetDeletionProtectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionProtection",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainAuthSecretArn(val *string) {
	if err := j.validateSetDomainAuthSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainAuthSecretArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainDnsIps(val *[]*string) {
	if err := j.validateSetDomainDnsIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainDnsIps",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainFqdn(val *string) {
	if err := j.validateSetDomainFqdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainFqdn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainIamRoleName(val *string) {
	if err := j.validateSetDomainIamRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainIamRoleName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetDomainOu(val *string) {
	if err := j.validateSetDomainOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainOu",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEnableCloudwatchLogsExports(val *[]*string) {
	if err := j.validateSetEnableCloudwatchLogsExportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableCloudwatchLogsExports",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEnableIamDatabaseAuthentication(val interface{}) {
	if err := j.validateSetEnableIamDatabaseAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableIamDatabaseAuthentication",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEnablePerformanceInsights(val interface{}) {
	if err := j.validateSetEnablePerformanceInsightsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePerformanceInsights",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEngine(val *string) {
	if err := j.validateSetEngineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engine",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetEngineVersion(val *string) {
	if err := j.validateSetEngineVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"engineVersion",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetIops(val *float64) {
	if err := j.validateSetIopsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iops",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetKmsKeyId(val *string) {
	if err := j.validateSetKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetLicenseModel(val *string) {
	if err := j.validateSetLicenseModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseModel",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetManageMasterUserPassword(val interface{}) {
	if err := j.validateSetManageMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manageMasterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMasterUsername(val *string) {
	if err := j.validateSetMasterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUsername",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMasterUserPassword(val *string) {
	if err := j.validateSetMasterUserPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"masterUserPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMaxAllocatedStorage(val *float64) {
	if err := j.validateSetMaxAllocatedStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAllocatedStorage",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMonitoringInterval(val *float64) {
	if err := j.validateSetMonitoringIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringInterval",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMonitoringRoleArn(val *string) {
	if err := j.validateSetMonitoringRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringRoleArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetMultiAz(val interface{}) {
	if err := j.validateSetMultiAzParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiAz",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetNcharCharacterSetName(val *string) {
	if err := j.validateSetNcharCharacterSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ncharCharacterSetName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetNetworkType(val *string) {
	if err := j.validateSetNetworkTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkType",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetOptionGroupName(val *string) {
	if err := j.validateSetOptionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optionGroupName",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPerformanceInsightsKmsKeyId(val *string) {
	if err := j.validateSetPerformanceInsightsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPerformanceInsightsRetentionPeriod(val *float64) {
	if err := j.validateSetPerformanceInsightsRetentionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"performanceInsightsRetentionPeriod",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPreferredBackupWindow(val *string) {
	if err := j.validateSetPreferredBackupWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredBackupWindow",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPreferredMaintenanceWindow(val *string) {
	if err := j.validateSetPreferredMaintenanceWindowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preferredMaintenanceWindow",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPromotionTier(val *float64) {
	if err := j.validateSetPromotionTierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"promotionTier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetPubliclyAccessible(val interface{}) {
	if err := j.validateSetPubliclyAccessibleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publiclyAccessible",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetReplicaMode(val *string) {
	if err := j.validateSetReplicaModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicaMode",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetRestoreTime(val *string) {
	if err := j.validateSetRestoreTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restoreTime",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbClusterIdentifier(val *string) {
	if err := j.validateSetSourceDbClusterIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbClusterIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbInstanceAutomatedBackupsArn(val *string) {
	if err := j.validateSetSourceDbInstanceAutomatedBackupsArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbInstanceIdentifier(val *string) {
	if err := j.validateSetSourceDbInstanceIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbInstanceIdentifier",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceDbiResourceId(val *string) {
	if err := j.validateSetSourceDbiResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceDbiResourceId",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetSourceRegion(val *string) {
	if err := j.validateSetSourceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceRegion",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetStorageEncrypted(val interface{}) {
	if err := j.validateSetStorageEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageEncrypted",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetStorageThroughput(val *float64) {
	if err := j.validateSetStorageThroughputParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageThroughput",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetStorageType(val *string) {
	if err := j.validateSetStorageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageType",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetTdeCredentialArn(val *string) {
	if err := j.validateSetTdeCredentialArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tdeCredentialArn",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetTdeCredentialPassword(val *string) {
	if err := j.validateSetTdeCredentialPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tdeCredentialPassword",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetUseDefaultProcessorFeatures(val interface{}) {
	if err := j.validateSetUseDefaultProcessorFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useDefaultProcessorFeatures",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetUseLatestRestorableTime(val interface{}) {
	if err := j.validateSetUseLatestRestorableTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useLatestRestorableTime",
		val,
	)
}

func (j *jsiiProxy_RdsDbInstance)SetVpcSecurityGroups(val *[]*string) {
	if err := j.validateSetVpcSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vpcSecurityGroups",
		val,
	)
}

// Generates CDKTF code for importing a RdsDbInstance resource upon running "cdktf plan <stack-name>".
func RdsDbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateRdsDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.rdsDbInstance.RdsDbInstance",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func RdsDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.rdsDbInstance.RdsDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.rdsDbInstance.RdsDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RdsDbInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRdsDbInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.rdsDbInstance.RdsDbInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RdsDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.rdsDbInstance.RdsDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RdsDbInstance) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_RdsDbInstance) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RdsDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RdsDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RdsDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RdsDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RdsDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RdsDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RdsDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RdsDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RdsDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RdsDbInstance) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_RdsDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbInstance) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_RdsDbInstance) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_RdsDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutAssociatedRoles(value interface{}) {
	if err := r.validatePutAssociatedRolesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAssociatedRoles",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutCertificateDetails(value *RdsDbInstanceCertificateDetails) {
	if err := r.validatePutCertificateDetailsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCertificateDetails",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutEndpoint(value *RdsDbInstanceEndpoint) {
	if err := r.validatePutEndpointParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putEndpoint",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutMasterUserSecret(value *RdsDbInstanceMasterUserSecret) {
	if err := r.validatePutMasterUserSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putMasterUserSecret",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutProcessorFeatures(value interface{}) {
	if err := r.validatePutProcessorFeaturesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putProcessorFeatures",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) PutTags(value interface{}) {
	if err := r.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTags",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAllocatedStorage() {
	_jsii_.InvokeVoid(
		r,
		"resetAllocatedStorage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAllowMajorVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowMajorVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAssociatedRoles() {
	_jsii_.InvokeVoid(
		r,
		"resetAssociatedRoles",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAutomaticBackupReplicationRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetAutomaticBackupReplicationRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAutoMinorVersionUpgrade() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoMinorVersionUpgrade",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetAvailabilityZone() {
	_jsii_.InvokeVoid(
		r,
		"resetAvailabilityZone",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetBackupRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetBackupRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCaCertificateIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetCaCertificateIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCertificateDetails() {
	_jsii_.InvokeVoid(
		r,
		"resetCertificateDetails",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCertificateRotationRestart() {
	_jsii_.InvokeVoid(
		r,
		"resetCertificateRotationRestart",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCharacterSetName() {
	_jsii_.InvokeVoid(
		r,
		"resetCharacterSetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCopyTagsToSnapshot() {
	_jsii_.InvokeVoid(
		r,
		"resetCopyTagsToSnapshot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetCustomIamInstanceProfile() {
	_jsii_.InvokeVoid(
		r,
		"resetCustomIamInstanceProfile",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbClusterSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbClusterSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbInstanceClass() {
	_jsii_.InvokeVoid(
		r,
		"resetDbInstanceClass",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbInstanceIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbInstanceIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbParameterGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbParameterGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbSecurityGroups() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSecurityGroups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbSnapshotIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSnapshotIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDbSubnetGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetDbSubnetGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDedicatedLogVolume() {
	_jsii_.InvokeVoid(
		r,
		"resetDedicatedLogVolume",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDeleteAutomatedBackups() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteAutomatedBackups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDeletionProtection() {
	_jsii_.InvokeVoid(
		r,
		"resetDeletionProtection",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomain() {
	_jsii_.InvokeVoid(
		r,
		"resetDomain",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainAuthSecretArn() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainAuthSecretArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainDnsIps() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainDnsIps",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainFqdn() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainFqdn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainIamRoleName() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainIamRoleName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetDomainOu() {
	_jsii_.InvokeVoid(
		r,
		"resetDomainOu",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEnableCloudwatchLogsExports() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableCloudwatchLogsExports",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEnableIamDatabaseAuthentication() {
	_jsii_.InvokeVoid(
		r,
		"resetEnableIamDatabaseAuthentication",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEnablePerformanceInsights() {
	_jsii_.InvokeVoid(
		r,
		"resetEnablePerformanceInsights",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEndpoint() {
	_jsii_.InvokeVoid(
		r,
		"resetEndpoint",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEngine() {
	_jsii_.InvokeVoid(
		r,
		"resetEngine",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetEngineVersion() {
	_jsii_.InvokeVoid(
		r,
		"resetEngineVersion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetIops() {
	_jsii_.InvokeVoid(
		r,
		"resetIops",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetLicenseModel() {
	_jsii_.InvokeVoid(
		r,
		"resetLicenseModel",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetManageMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetManageMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMasterUsername() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUsername",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMasterUserPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMasterUserSecret() {
	_jsii_.InvokeVoid(
		r,
		"resetMasterUserSecret",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMaxAllocatedStorage() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxAllocatedStorage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMonitoringInterval() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringInterval",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMonitoringRoleArn() {
	_jsii_.InvokeVoid(
		r,
		"resetMonitoringRoleArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetMultiAz() {
	_jsii_.InvokeVoid(
		r,
		"resetMultiAz",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetNcharCharacterSetName() {
	_jsii_.InvokeVoid(
		r,
		"resetNcharCharacterSetName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetNetworkType() {
	_jsii_.InvokeVoid(
		r,
		"resetNetworkType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetOptionGroupName() {
	_jsii_.InvokeVoid(
		r,
		"resetOptionGroupName",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPerformanceInsightsKmsKeyId() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsKmsKeyId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPerformanceInsightsRetentionPeriod() {
	_jsii_.InvokeVoid(
		r,
		"resetPerformanceInsightsRetentionPeriod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPort() {
	_jsii_.InvokeVoid(
		r,
		"resetPort",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPreferredBackupWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredBackupWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPreferredMaintenanceWindow() {
	_jsii_.InvokeVoid(
		r,
		"resetPreferredMaintenanceWindow",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetProcessorFeatures() {
	_jsii_.InvokeVoid(
		r,
		"resetProcessorFeatures",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPromotionTier() {
	_jsii_.InvokeVoid(
		r,
		"resetPromotionTier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetPubliclyAccessible() {
	_jsii_.InvokeVoid(
		r,
		"resetPubliclyAccessible",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetReplicaMode() {
	_jsii_.InvokeVoid(
		r,
		"resetReplicaMode",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetRestoreTime() {
	_jsii_.InvokeVoid(
		r,
		"resetRestoreTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbClusterIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbClusterIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbInstanceAutomatedBackupsArn() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbInstanceAutomatedBackupsArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbInstanceIdentifier() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbInstanceIdentifier",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceDbiResourceId() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceDbiResourceId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetSourceRegion() {
	_jsii_.InvokeVoid(
		r,
		"resetSourceRegion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetStorageEncrypted() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageEncrypted",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetStorageThroughput() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageThroughput",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetStorageType() {
	_jsii_.InvokeVoid(
		r,
		"resetStorageType",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTags() {
	_jsii_.InvokeVoid(
		r,
		"resetTags",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTdeCredentialArn() {
	_jsii_.InvokeVoid(
		r,
		"resetTdeCredentialArn",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTdeCredentialPassword() {
	_jsii_.InvokeVoid(
		r,
		"resetTdeCredentialPassword",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetTimezone() {
	_jsii_.InvokeVoid(
		r,
		"resetTimezone",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetUseDefaultProcessorFeatures() {
	_jsii_.InvokeVoid(
		r,
		"resetUseDefaultProcessorFeatures",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetUseLatestRestorableTime() {
	_jsii_.InvokeVoid(
		r,
		"resetUseLatestRestorableTime",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) ResetVpcSecurityGroups() {
	_jsii_.InvokeVoid(
		r,
		"resetVpcSecurityGroups",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RdsDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RdsDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

