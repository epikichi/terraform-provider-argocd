package argocd

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func repositoryCredentialsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"url": {
			Type:        schema.TypeString,
			Description: "URL is the URL that this credentials matches to",
			Required:    true,
		},
		"username": {
			Type:        schema.TypeString,
			Description: "Username for authenticating at the repo server",
			Optional:    true,
		},
		"password": {
			Type:        schema.TypeString,
			Sensitive:   true,
			Description: "Password for authenticating at the repo server, cannot be managed once created!",
			Optional:    true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return true
			},
		},
		"ssh_private_key": {
			Type:        schema.TypeString,
			Sensitive:   true,
			Description: "SSH private key data for authenticating at the repo server only for Git repos, cannot be managed once created!",
			// TODO: add a validator
			Optional: true,
			DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
				return true
			},
		},
		"tls_client_cert_data": {
			Type:        schema.TypeString,
			Description: "TLS client cert data for authenticating at the repo server",
			// TODO: add a validator
			Optional: true,
		},
		"tls_client_cert_key": {
			Type:        schema.TypeString,
			Sensitive:   true,
			Description: "TLS client cert key for authenticating at the repo server ",
			// TODO: add a validator
			Optional: true,
		},
	}
}