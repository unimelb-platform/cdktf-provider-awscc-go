package dataawsccrdsdbinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/unimelb-platform/cdktf-provider-awscc-go/awscc/dataawsccrdsdbinstance/internal"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/rds_db_instance awscc_rds_db_instance}.
type DataAwsccRdsDbInstance interface {
	cdktf.TerraformDataSource
	AllocatedStorage() *string
	AllowMajorVersionUpgrade() cdktf.IResolvable
	AssociatedRoles() DataAwsccRdsDbInstanceAssociatedRolesList
	AutomaticBackupReplicationRegion() *string
	AutoMinorVersionUpgrade() cdktf.IResolvable
	AvailabilityZone() *string
	BackupRetentionPeriod() *float64
	CaCertificateIdentifier() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CertificateDetails() DataAwsccRdsDbInstanceCertificateDetailsOutputReference
	CertificateRotationRestart() cdktf.IResolvable
	CharacterSetName() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	CopyTagsToSnapshot() cdktf.IResolvable
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomIamInstanceProfile() *string
	DbClusterIdentifier() *string
	DbClusterSnapshotIdentifier() *string
	DbInstanceArn() *string
	DbInstanceClass() *string
	DbInstanceIdentifier() *string
	DbiResourceId() *string
	DbName() *string
	DbParameterGroupName() *string
	DbSecurityGroups() *[]*string
	DbSnapshotIdentifier() *string
	DbSubnetGroupName() *string
	DbSystemId() *string
	DedicatedLogVolume() cdktf.IResolvable
	DeleteAutomatedBackups() cdktf.IResolvable
	DeletionProtection() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Domain() *string
	DomainAuthSecretArn() *string
	DomainDnsIps() *[]*string
	DomainFqdn() *string
	DomainIamRoleName() *string
	DomainOu() *string
	EnableCloudwatchLogsExports() *[]*string
	EnableIamDatabaseAuthentication() cdktf.IResolvable
	EnablePerformanceInsights() cdktf.IResolvable
	Endpoint() DataAwsccRdsDbInstanceEndpointOutputReference
	Engine() *string
	EngineVersion() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Iops() *float64
	KmsKeyId() *string
	LicenseModel() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManageMasterUserPassword() cdktf.IResolvable
	MasterUsername() *string
	MasterUserPassword() *string
	MasterUserSecret() DataAwsccRdsDbInstanceMasterUserSecretOutputReference
	MaxAllocatedStorage() *float64
	MonitoringInterval() *float64
	MonitoringRoleArn() *string
	MultiAz() cdktf.IResolvable
	NcharCharacterSetName() *string
	NetworkType() *string
	// The tree node.
	Node() constructs.Node
	OptionGroupName() *string
	PerformanceInsightsKmsKeyId() *string
	PerformanceInsightsRetentionPeriod() *float64
	Port() *string
	PreferredBackupWindow() *string
	PreferredMaintenanceWindow() *string
	ProcessorFeatures() DataAwsccRdsDbInstanceProcessorFeaturesList
	PromotionTier() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	PubliclyAccessible() cdktf.IResolvable
	// Experimental.
	RawOverrides() interface{}
	ReplicaMode() *string
	RestoreTime() *string
	SourceDbClusterIdentifier() *string
	SourceDbInstanceAutomatedBackupsArn() *string
	SourceDbInstanceIdentifier() *string
	SourceDbiResourceId() *string
	SourceRegion() *string
	StorageEncrypted() cdktf.IResolvable
	StorageThroughput() *float64
	StorageType() *string
	Tags() DataAwsccRdsDbInstanceTagsList
	TdeCredentialArn() *string
	TdeCredentialPassword() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timezone() *string
	UseDefaultProcessorFeatures() cdktf.IResolvable
	UseLatestRestorableTime() cdktf.IResolvable
	VpcSecurityGroups() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataAwsccRdsDbInstance
type jsiiProxy_DataAwsccRdsDbInstance struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) AllocatedStorage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) AllowMajorVersionUpgrade() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"allowMajorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) AssociatedRoles() DataAwsccRdsDbInstanceAssociatedRolesList {
	var returns DataAwsccRdsDbInstanceAssociatedRolesList
	_jsii_.Get(
		j,
		"associatedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) AutomaticBackupReplicationRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"automaticBackupReplicationRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) AutoMinorVersionUpgrade() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoMinorVersionUpgrade",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) AvailabilityZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilityZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) BackupRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"backupRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) CaCertificateIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertificateIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) CertificateDetails() DataAwsccRdsDbInstanceCertificateDetailsOutputReference {
	var returns DataAwsccRdsDbInstanceCertificateDetailsOutputReference
	_jsii_.Get(
		j,
		"certificateDetails",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) CertificateRotationRestart() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"certificateRotationRestart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) CharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"characterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) CopyTagsToSnapshot() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"copyTagsToSnapshot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) CustomIamInstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customIamInstanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbClusterSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbClusterSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbInstanceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbInstanceClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbParameterGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbParameterGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dbSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbSnapshotIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSnapshotIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbSubnetGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSubnetGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DbSystemId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dbSystemId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DedicatedLogVolume() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"dedicatedLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DeleteAutomatedBackups() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deleteAutomatedBackups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DeletionProtection() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"deletionProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DomainAuthSecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainAuthSecretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DomainDnsIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"domainDnsIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DomainFqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainFqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DomainIamRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainIamRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) DomainOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) EnableCloudwatchLogsExports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"enableCloudwatchLogsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) EnableIamDatabaseAuthentication() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enableIamDatabaseAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) EnablePerformanceInsights() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"enablePerformanceInsights",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Endpoint() DataAwsccRdsDbInstanceEndpointOutputReference {
	var returns DataAwsccRdsDbInstanceEndpointOutputReference
	_jsii_.Get(
		j,
		"endpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Engine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) EngineVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"engineVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Iops() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"iops",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) KmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) LicenseModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) ManageMasterUserPassword() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"manageMasterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) MasterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) MasterUserPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"masterUserPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) MasterUserSecret() DataAwsccRdsDbInstanceMasterUserSecretOutputReference {
	var returns DataAwsccRdsDbInstanceMasterUserSecretOutputReference
	_jsii_.Get(
		j,
		"masterUserSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) MaxAllocatedStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAllocatedStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) MonitoringInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitoringInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) MonitoringRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) MultiAz() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"multiAz",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) NcharCharacterSetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ncharCharacterSetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) NetworkType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) OptionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"optionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) PerformanceInsightsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"performanceInsightsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) PerformanceInsightsRetentionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"performanceInsightsRetentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) PreferredBackupWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredBackupWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) PreferredMaintenanceWindow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferredMaintenanceWindow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) ProcessorFeatures() DataAwsccRdsDbInstanceProcessorFeaturesList {
	var returns DataAwsccRdsDbInstanceProcessorFeaturesList
	_jsii_.Get(
		j,
		"processorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) PromotionTier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"promotionTier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) PubliclyAccessible() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"publiclyAccessible",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) ReplicaMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicaMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) RestoreTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"restoreTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) SourceDbClusterIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbClusterIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) SourceDbInstanceAutomatedBackupsArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceAutomatedBackupsArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) SourceDbInstanceIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbInstanceIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) SourceDbiResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceDbiResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) SourceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) StorageEncrypted() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"storageEncrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) StorageThroughput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"storageThroughput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) StorageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Tags() DataAwsccRdsDbInstanceTagsList {
	var returns DataAwsccRdsDbInstanceTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) TdeCredentialArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) TdeCredentialPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tdeCredentialPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) UseDefaultProcessorFeatures() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useDefaultProcessorFeatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) UseLatestRestorableTime() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"useLatestRestorableTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsccRdsDbInstance) VpcSecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"vpcSecurityGroups",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/rds_db_instance awscc_rds_db_instance} Data Source.
func NewDataAwsccRdsDbInstance(scope constructs.Construct, id *string, config *DataAwsccRdsDbInstanceConfig) DataAwsccRdsDbInstance {
	_init_.Initialize()

	if err := validateNewDataAwsccRdsDbInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsccRdsDbInstance{}

	_jsii_.Create(
		"awscc.dataAwsccRdsDbInstance.DataAwsccRdsDbInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/data-sources/rds_db_instance awscc_rds_db_instance} Data Source.
func NewDataAwsccRdsDbInstance_Override(d DataAwsccRdsDbInstance, scope constructs.Construct, id *string, config *DataAwsccRdsDbInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"awscc.dataAwsccRdsDbInstance.DataAwsccRdsDbInstance",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataAwsccRdsDbInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRdsDbInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRdsDbInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRdsDbInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRdsDbInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataAwsccRdsDbInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

// Generates CDKTF code for importing a DataAwsccRdsDbInstance resource upon running "cdktf plan <stack-name>".
func DataAwsccRdsDbInstance_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataAwsccRdsDbInstance_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"awscc.dataAwsccRdsDbInstance.DataAwsccRdsDbInstance",
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
func DataAwsccRdsDbInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccRdsDbInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccRdsDbInstance.DataAwsccRdsDbInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccRdsDbInstance_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccRdsDbInstance_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccRdsDbInstance.DataAwsccRdsDbInstance",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataAwsccRdsDbInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataAwsccRdsDbInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"awscc.dataAwsccRdsDbInstance.DataAwsccRdsDbInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataAwsccRdsDbInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"awscc.dataAwsccRdsDbInstance.DataAwsccRdsDbInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAwsccRdsDbInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsccRdsDbInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

