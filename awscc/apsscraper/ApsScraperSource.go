package apsscraper


type ApsScraperSource struct {
	// Configuration for EKS metrics source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_scraper#eks_configuration ApsScraper#eks_configuration}
	EksConfiguration *ApsScraperSourceEksConfiguration `field:"optional" json:"eksConfiguration" yaml:"eksConfiguration"`
}

