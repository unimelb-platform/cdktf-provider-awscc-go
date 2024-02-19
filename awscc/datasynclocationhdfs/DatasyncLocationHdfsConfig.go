package datasynclocationhdfs

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationHdfsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// ARN(s) of the agent(s) to use for an HDFS location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#agent_arns DatasyncLocationHdfs#agent_arns}
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// The authentication mode used to determine identity of user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#authentication_type DatasyncLocationHdfs#authentication_type}
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// An array of Name Node(s) of the HDFS location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#name_nodes DatasyncLocationHdfs#name_nodes}
	NameNodes interface{} `field:"required" json:"nameNodes" yaml:"nameNodes"`
	// Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#block_size DatasyncLocationHdfs#block_size}
	BlockSize *float64 `field:"optional" json:"blockSize" yaml:"blockSize"`
	// The Base64 string representation of the Keytab file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#kerberos_keytab DatasyncLocationHdfs#kerberos_keytab}
	KerberosKeytab *string `field:"optional" json:"kerberosKeytab" yaml:"kerberosKeytab"`
	// The string representation of the Krb5Conf file, or the presigned URL to access the Krb5.conf file within an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#kerberos_krb_5_conf DatasyncLocationHdfs#kerberos_krb_5_conf}
	KerberosKrb5Conf *string `field:"optional" json:"kerberosKrb5Conf" yaml:"kerberosKrb5Conf"`
	// The unique identity, or principal, to which Kerberos can assign tickets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#kerberos_principal DatasyncLocationHdfs#kerberos_principal}
	KerberosPrincipal *string `field:"optional" json:"kerberosPrincipal" yaml:"kerberosPrincipal"`
	// The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#kms_key_provider_uri DatasyncLocationHdfs#kms_key_provider_uri}
	KmsKeyProviderUri *string `field:"optional" json:"kmsKeyProviderUri" yaml:"kmsKeyProviderUri"`
	// Configuration information for RPC Protection and Data Transfer Protection.
	//
	// These parameters can be set to AUTHENTICATION, INTEGRITY, or PRIVACY. The default value is PRIVACY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#qop_configuration DatasyncLocationHdfs#qop_configuration}
	QopConfiguration *DatasyncLocationHdfsQopConfiguration `field:"optional" json:"qopConfiguration" yaml:"qopConfiguration"`
	// Number of copies of each block that exists inside the HDFS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#replication_factor DatasyncLocationHdfs#replication_factor}
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
	// The user name that has read and write permissions on the specified HDFS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#simple_user DatasyncLocationHdfs#simple_user}
	SimpleUser *string `field:"optional" json:"simpleUser" yaml:"simpleUser"`
	// The subdirectory in HDFS that is used to read data from the HDFS source location or write data to the HDFS destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#subdirectory DatasyncLocationHdfs#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_hdfs#tags DatasyncLocationHdfs#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

