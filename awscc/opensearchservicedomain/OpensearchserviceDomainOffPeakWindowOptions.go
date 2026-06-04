package opensearchservicedomain


type OpensearchserviceDomainOffPeakWindowOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain#enabled OpensearchserviceDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/opensearchservice_domain#off_peak_window OpensearchserviceDomain#off_peak_window}.
	OffPeakWindow *OpensearchserviceDomainOffPeakWindowOptionsOffPeakWindow `field:"optional" json:"offPeakWindow" yaml:"offPeakWindow"`
}

