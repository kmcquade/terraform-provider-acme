package tls

import (
	"crypto/sha1"
	"crypto/x509/pkix"
	"encoding/hex"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"tls_private_key":         resourcePrivateKey(),
			"tls_locally_signed_cert": resourceLocallySignedCert(),
			"tls_self_signed_cert":    resourceSelfSignedCert(),
			"tls_cert_request":        resourceCertRequest(),
		},
	}
}

func hashForState(value string) string {
	if value == "" {
		return ""
	}
	hash := sha1.Sum([]byte(strings.TrimSpace(value)))
	return hex.EncodeToString(hash[:])
}

func nameFromResourceData(nameMap map[string]interface{}) (*pkix.Name, error) {
	result := &pkix.Name{}

	if value := nameMap["common_name"]; value != "" {
		result.CommonName = value.(string)
	}
	if value := nameMap["organization"]; value != "" {
		result.Organization = []string{value.(string)}
	}
	if value := nameMap["organizational_unit"]; value != "" {
		result.OrganizationalUnit = []string{value.(string)}
	}
	if value := nameMap["street_address"].([]interface{}); len(value) > 0 {
		result.StreetAddress = make([]string, len(value))
		for i, vi := range value {
			result.StreetAddress[i] = vi.(string)
		}
	}
	if value := nameMap["locality"]; value != "" {
		result.Locality = []string{value.(string)}
	}
	if value := nameMap["province"]; value != "" {
		result.Province = []string{value.(string)}
	}
	if value := nameMap["country"]; value != "" {
		result.Country = []string{value.(string)}
	}
	if value := nameMap["postal_code"]; value != "" {
		result.PostalCode = []string{value.(string)}
	}
	if value := nameMap["serial_number"]; value != "" {
		result.SerialNumber = value.(string)
	}

	return result, nil
}

var nameSchema *schema.Resource = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"organization": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"common_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"organizational_unit": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"street_address": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"locality": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"province": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"country": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"postal_code": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"serial_number": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	},
}
