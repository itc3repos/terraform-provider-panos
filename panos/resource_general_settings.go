package panos

import (
	"github.com/PaloAltoNetworks/pango"
	"github.com/PaloAltoNetworks/pango/dev/general"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceGeneralSettings() *schema.Resource {
	return &schema.Resource{
		Create: createUpdateGeneralSettings,
		Read:   readGeneralSettings,
		Update: createUpdateGeneralSettings,
		Delete: deleteGeneralSettings,

		Schema: map[string]*schema.Schema{
			"hostname": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The firewall hostname",
			},
			"timezone": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Timezone",
			},
			"domain": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Domain",
			},
			"update_server": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "updates.paloaltonetworks.com",
				Description: "PANOS update server",
			},
			"verify_update_server": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Verify update server identity",
			},
			"dns_primary": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Primary DNS IP address",
			},
			"dns_secondary": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Secondary DNS IP address",
			},
			"ntp_primary_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Primary NTP server",
			},
			"ntp_primary_auth_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP auth type (none, autokey, symmetric-key)",
				ValidateFunc: validateStringIn("none", "autokey", "symmetric-key"),
			},
			"ntp_primary_key_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key key ID",
			},
			"ntp_primary_algorithm": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP symmetric-key algorithm (sha1 or md5)",
				ValidateFunc: validateStringIn("sha1", "md5"),
			},
			"ntp_primary_auth_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key auth key",
			},
			"ntp_secondary_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Secondary NTP server",
			},
			"ntp_secondary_auth_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP auth type (none, autokey, symmetric-key)",
				ValidateFunc: validateStringIn("none", "autokey", "symmetric-key"),
			},
			"ntp_secondary_key_id": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key key ID",
			},
			"ntp_secondary_algorithm": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "NTP symmetric-key algorithm (sha1 or md5)",
				ValidateFunc: validateStringIn("sha1", "md5"),
			},
			"ntp_secondary_auth_key": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "NTP symmetric-key auth key",
			},
		},
	}
}

func parseGeneralSettings(d *schema.ResourceData) general.Config {
	return general.Config{
		Hostname:              d.Get("hostname").(string),
		Timezone:              d.Get("timezone").(string),
		Domain:                d.Get("domain").(string),
		UpdateServer:          d.Get("update_server").(string),
		VerifyUpdateServer:    d.Get("verify_update_server").(bool),
		DnsPrimary:            d.Get("dns_primary").(string),
		DnsSecondary:          d.Get("dns_secondary").(string),
		NtpPrimaryAddress:     d.Get("ntp_primary_address").(string),
		NtpPrimaryAuthType:    d.Get("ntp_primary_auth_type").(string),
		NtpPrimaryKeyId:       d.Get("ntp_primary_key_id").(int),
		NtpPrimaryAlgorithm:   d.Get("ntp_primary_algorithm").(string),
		NtpPrimaryAuthKey:     d.Get("ntp_primary_auth_key").(string),
		NtpSecondaryAddress:   d.Get("ntp_secondary_address").(string),
		NtpSecondaryAuthType:  d.Get("ntp_secondary_auth_type").(string),
		NtpSecondaryKeyId:     d.Get("ntp_secondary_key_id").(int),
		NtpSecondaryAlgorithm: d.Get("ntp_secondary_algorithm").(string),
		NtpSecondaryAuthKey:   d.Get("ntp_secondary_auth_key").(string),
	}
}

func createUpdateGeneralSettings(d *schema.ResourceData, meta interface{}) error {
	fw := meta.(*pango.Firewall)

	o, err := fw.Device.GeneralSettings.Get()
	if err != nil {
		return err
	}

	o.Merge(parseGeneralSettings(d))
	if err = fw.Device.GeneralSettings.Edit(o); err != nil {
		return err
	}

	d.SetId(o.Hostname)
	return readGeneralSettings(d, meta)
}

func readGeneralSettings(d *schema.ResourceData, meta interface{}) error {
	fw := meta.(*pango.Firewall)
	o, err := fw.Device.GeneralSettings.Get()
	if err != nil {
		// I don't think you can delete the general settings from a firewall,
		// so any error is a real error.
		return err
	}

	d.Set("hostname", o.Hostname)
	d.Set("timezone", o.Timezone)
	d.Set("domain", o.Domain)
	d.Set("update_server", o.UpdateServer)
	d.Set("verify_update_server", o.VerifyUpdateServer)
	d.Set("dns_primary", o.DnsPrimary)
	d.Set("dns_secondary", o.DnsSecondary)
	d.Set("ntp_primary_address", o.NtpPrimaryAddress)
	d.Set("ntp_primary_auth_type", o.NtpPrimaryAuthType)
	d.Set("ntp_primary_key_id", o.NtpPrimaryKeyId)
	d.Set("ntp_primary_algorithm", o.NtpPrimaryAlgorithm)
	d.Set("ntp_primary_auth_key", o.NtpPrimaryAuthKey)
	d.Set("ntp_secondary_address", o.NtpSecondaryAddress)
	d.Set("ntp_secondary_auth_type", o.NtpSecondaryAuthType)
	d.Set("ntp_secondary_key_id", o.NtpSecondaryKeyId)
	d.Set("ntp_secondary_algorithm", o.NtpSecondaryAlgorithm)
	d.Set("ntp_secondary_auth_key", o.NtpSecondaryAuthKey)

	return nil
}

func deleteGeneralSettings(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
