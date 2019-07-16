package dyn

import (
	"github.com/Wikia/go-dynect/dynect"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDynZonePrimaryImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	results := make([]*schema.ResourceData, 1, 1)

	client := meta.(*dynect.ConvenientClient)

	zone := &dynect.Zone{
		Zone: d.Id(),
		Type: "Primary",
	}

	err := client.GetZone(zone)

	if err != nil {
		return nil, err
	}

	d.SetId(zone.Zone)
	d.Set("name", zone.Zone)
	d.Set("type", zone.Type)
	d.Set("serial", zone.Serial)
	d.Set("serial_style", zone.SerialStyle)
	results[0] = d

	return results, nil
}
