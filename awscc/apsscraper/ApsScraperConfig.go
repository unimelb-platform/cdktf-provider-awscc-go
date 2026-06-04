package apsscraper

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApsScraperConfig struct {
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
	// Scraper metrics destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#destination ApsScraper#destination}
	Destination *ApsScraperDestination `field:"required" json:"destination" yaml:"destination"`
	// Scraper configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#scrape_configuration ApsScraper#scrape_configuration}
	ScrapeConfiguration *ApsScraperScrapeConfiguration `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// Scraper metrics source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#source ApsScraper#source}
	Source *ApsScraperSource `field:"required" json:"source" yaml:"source"`
	// Scraper alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#alias ApsScraper#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Role configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#role_configuration ApsScraper#role_configuration}
	RoleConfiguration *ApsScraperRoleConfiguration `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#tags ApsScraper#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

