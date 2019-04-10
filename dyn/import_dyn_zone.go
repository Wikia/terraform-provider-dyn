package dyn

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nesv/go-dynect/dynect"
)

func resourceDynZoneImportState(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	results := make([]*schema.ResourceData, 1, 1)

	client := meta.(*dynect.ConvenientClient)

	values := strings.Split(d.Id(), "/")

	if len(values) != 2 {
		return nil, fmt.Errorf("invalid id provided, expected format: {type}/{zone}]")
	}

	zoneType := values[0]
	zoneZone := values[1]

	zone := &dynect.Zone{
		Zone: zoneZone,
		Type: zoneType,
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
