package proxmox

import (
	pxapi "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourcePool() *schema.Resource {
	*pxapi.Debug = true
	return &schema.Resource{
		Create: resourcePoolCreate,
		Read:   resourcePoolRead,
		Update: resourcePoolUpdate,
		Delete: resourcePoolDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourcePoolCreate(d *schema.ResourceData, meta interface{}) error {
	pconf := meta.(*providerConfiguration)
	client := pconf.Client
	d.SetId(d.Get("name").(string))
	return nil
}

func resourcePoolRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePoolUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourcePoolDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
