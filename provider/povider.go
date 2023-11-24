package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider возвращает провайдер Terraform
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Veeam API username",
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Default:      schema.MultiEnvDefaultFunc([]string{"VEEAM_USER", "VEEAM_USERNAME"}, nil),
			},
			"password": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Veeam API password",
				ValidateFunc: validation.StringIsNotWhiteSpace,
				Default:      schema.MultiEnvDefaultFunc([]string{"VEEAM_PASS", "VEEAM_PASSWORD"}, nil),
			},
			"url": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Veeam API url",
				ValidateFunc: validation.IsURLWithHTTPorHTTPS,
				Default:      schema.MultiEnvDefaultFunc([]string{"VEEAM_URL"}, nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{},
	}
}
