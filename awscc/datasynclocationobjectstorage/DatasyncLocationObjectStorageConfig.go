package datasynclocationobjectstorage

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationObjectStorageConfig struct {
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
	// The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#agent_arns DatasyncLocationObjectStorage#agent_arns}
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// Optional. The access key is used if credentials are required to access the self-managed object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#access_key DatasyncLocationObjectStorage#access_key}
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// The name of the bucket on the self-managed object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#bucket_name DatasyncLocationObjectStorage#bucket_name}
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Optional. The secret key is used if credentials are required to access the self-managed object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#secret_key DatasyncLocationObjectStorage#secret_key}
	SecretKey *string `field:"optional" json:"secretKey" yaml:"secretKey"`
	// X.509 PEM content containing a certificate authority or chain to trust.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#server_certificate DatasyncLocationObjectStorage#server_certificate}
	ServerCertificate *string `field:"optional" json:"serverCertificate" yaml:"serverCertificate"`
	// The name of the self-managed object storage server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the object storage server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#server_hostname DatasyncLocationObjectStorage#server_hostname}
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// The port that your self-managed server accepts inbound network traffic on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#server_port DatasyncLocationObjectStorage#server_port}
	ServerPort *float64 `field:"optional" json:"serverPort" yaml:"serverPort"`
	// The protocol that the object storage server uses to communicate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#server_protocol DatasyncLocationObjectStorage#server_protocol}
	ServerProtocol *string `field:"optional" json:"serverProtocol" yaml:"serverProtocol"`
	// The subdirectory in the self-managed object storage server that is used to read data from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#subdirectory DatasyncLocationObjectStorage#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_object_storage#tags DatasyncLocationObjectStorage#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

