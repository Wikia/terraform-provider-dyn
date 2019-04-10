package dyn

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nesv/go-dynect/dynect"
)

func resourceDynZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceDynZoneCreate,
		Read:   resourceDynZoneRead,
		Delete: resourceDynZoneDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDynZoneImportState,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"mailbox": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"serial_style": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "increment",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "increment" || v != "epoch" || v != "day" || v != "minute" {
						errs = append(errs, fmt.Errorf("%q allowed values are increment, epoch, day or minute, got: %d", key, v))
					}
					return
				},
			},

			"serial": {
				Type:     schema.TypeString,
				Computed: true,
			},

			"type": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},

			"ttl": {
				Type:     schema.TypeString,
				ForceNew: true,
				Optional: true,
				Default:  "3600",
			},
		},
	}
}
func resourceDynZoneCreate(d *schema.ResourceData, meta interface{}) error {
	mutex.Lock()

	client := meta.(*dynect.ConvenientClient)

	name := d.Get("name").(string)

	// create the zone
	err := client.CreateZone(name, d.Get("mailbox").(string), d.Get("serial_style").(string), d.Get("ttl").(string))

	if err != nil {
		mutex.Unlock()
		return fmt.Errorf("Failed to create Dyn zone: %s", err)
	}

	err = client.PublishZone(name)
	if err != nil {
		mutex.Unlock()
		return fmt.Errorf("Failed to publish Dyn zone: %s", err)
	}

	mutex.Unlock()
	return resourceDynZoneRead(d, meta)
}

func resourceDynZoneRead(d *schema.ResourceData, meta interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()

	client := meta.(*dynect.ConvenientClient)

	zone := &dynect.Zone{
		Zone:        d.Get("name").(string),
		Type:        d.Get("type").(string),
		SerialStyle: d.Get("serial_style").(string),
		Serial:      d.Get("serial").(string),
	}

	err := client.GetZone(zone)
	if err != nil {
		return fmt.Errorf("Couldn't find Dyn zone: %s", err)
	}

	d.SetId(zone.Zone)
	d.Set("name", zone.Zone)
	d.Set("type", zone.Type)
	d.Set("serial_style", zone.SerialStyle)
	d.Set("serial", zone.Serial)

	return nil
}

func resourceDynZoneDelete(d *schema.ResourceData, meta interface{}) error {
	mutex.Lock()
	defer mutex.Unlock()

	client := meta.(*dynect.ConvenientClient)

	// delete the zone
	err := client.DeleteZone(d.Get("name").(string))
	if err != nil {
		return fmt.Errorf("Failed to delete Dyn zone: %s", err)
	}

	return nil
}
