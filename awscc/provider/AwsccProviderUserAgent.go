package provider


type AwsccProviderUserAgent struct {
	// Product name. At least one of `product_name` or `comment` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#product_name AwsccProvider#product_name}
	ProductName *string `field:"required" json:"productName" yaml:"productName"`
	// User-Agent comment. At least one of `comment` or `product_name` must be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#comment AwsccProvider#comment}
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Product version. Optional, and should only be set when `product_name` is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs#product_version AwsccProvider#product_version}
	ProductVersion *string `field:"optional" json:"productVersion" yaml:"productVersion"`
}

