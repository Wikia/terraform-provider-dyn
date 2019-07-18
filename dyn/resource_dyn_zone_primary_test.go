package dyn

import (
	"fmt"
	"os"
	"testing"

	"github.com/Wikia/go-dynect/dynect"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDynZonePrimaryBasic(t *testing.T) {
	var zone dynect.Zone
	name := os.Getenv("DYN_ZONE")

	testAccDeleteTestDynZone(name, t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDynZonePrimaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckDynZonePrimaryConfigBasic, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDynZonePrimaryExists("dyn_zone_primary.foobar", &zone),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "mailbox", "test@terraform.com"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "type", "Primary"),
				),
			},
		},
	})
}

func TestAccDynZonePrimaryTTL(t *testing.T) {
	var zone dynect.Zone
	name := os.Getenv("DYN_ZONE")

	testAccDeleteTestDynZone(name, t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDynZonePrimaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckDynZonePrimaryConfigTTL, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDynZonePrimaryExists("dyn_zone_primary.foobar", &zone),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "mailbox", "test@terraform.com"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "type", "Primary"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "ttl", "60"),
				),
			},
		},
	})
}

func TestAccDynZonePrimarySerialStyleEpoch(t *testing.T) {
	var zone dynect.Zone
	name := os.Getenv("DYN_ZONE")

	testAccDeleteTestDynZone(name, t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDynZonePrimaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckDynZonePrimaryConfigSerialStyleEpoch, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDynZonePrimaryExists("dyn_zone_primary.foobar", &zone),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "mailbox", "test@terraform.com"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "type", "Primary"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "serial_style", "epoch"),
				),
			},
		},
	})
}

func TestAccDynZonePrimarySerialStyleDay(t *testing.T) {
	var zone dynect.Zone
	name := os.Getenv("DYN_ZONE")

	testAccDeleteTestDynZone(name, t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDynZonePrimaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckDynZonePrimaryConfigSerialStyleDay, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDynZonePrimaryExists("dyn_zone_primary.foobar", &zone),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "mailbox", "test@terraform.com"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "type", "Primary"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "serial_style", "day"),
				),
			},
		},
	})
}

func TestAccDynZonePrimarySerialStyleMinute(t *testing.T) {
	var zone dynect.Zone
	name := os.Getenv("DYN_ZONE")

	testAccDeleteTestDynZone(name, t)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDynZonePrimaryDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccCheckDynZonePrimaryConfigSerialStyleMinute, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDynZonePrimaryExists("dyn_zone_primary.foobar", &zone),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "name", name),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "mailbox", "test@terraform.com"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "type", "Primary"),
					resource.TestCheckResourceAttr(
						"dyn_zone_primary.foobar", "serial_style", "minute"),
				),
			},
		},
	})
}

func testAccCheckDynZonePrimaryDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*dynect.ConvenientClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "dyn_zone_primary" {
			continue
		}

		foundZone := &dynect.Zone{
			Zone:        rs.Primary.Attributes["name"],
			Type:        rs.Primary.Attributes["type"],
			SerialStyle: rs.Primary.Attributes["serial_style"],
			Serial:      rs.Primary.Attributes["serial"],
		}

		err := client.GetZone(foundZone)

		if err == nil {
			return fmt.Errorf("Zone still exists")
		}
	}

	return nil
}

func testAccCheckDynZonePrimaryExists(n string, record *dynect.Zone) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		client := testAccProvider.Meta().(*dynect.ConvenientClient)

		foundZone := &dynect.Zone{
			Zone:        rs.Primary.Attributes["name"],
			Type:        rs.Primary.Attributes["type"],
			SerialStyle: rs.Primary.Attributes["serial_style"],
			Serial:      rs.Primary.Attributes["serial"],
		}

		err := client.GetZone(foundZone)

		if err != nil {
			return err
		}

		if foundZone.Zone != rs.Primary.Attributes["name"] {
			return fmt.Errorf("Zone not found")
		}

		*record = *foundZone

		return nil
	}
}

func testAccDeleteTestDynZone(name string, t *testing.T) {

	client := testAccProvider.Meta().(*dynect.ConvenientClient)
	foundZone := &dynect.Zone{
		Zone: name,
	}
	err := client.GetZone(foundZone)
	if err == nil {
		err := client.DeleteZone(name)
		if err != nil {
			t.Logf("%s", err)
			t.Fail()
		}
	}
}

const testAccCheckDynZonePrimaryConfigBasic = `
resource "dyn_zone_primary" "foobar" {
	name = "%s"
	mailbox = "test@terraform.com"
}`

const testAccCheckDynZonePrimaryConfigTTL = `
resource "dyn_zone_primary" "foobar" {
	name = "%s"
	mailbox = "test@terraform.com"
	ttl = "60"
}`

const testAccCheckDynZonePrimaryConfigSerialStyleEpoch = `
resource "dyn_zone_primary" "foobar" {
	name = "%s"
	mailbox = "test@terraform.com"
	serial_style = "epoch"
}`

const testAccCheckDynZonePrimaryConfigSerialStyleDay = `
resource "dyn_zone_primary" "foobar" {
	name = "%s"
	mailbox = "test@terraform.com"
	serial_style = "day"
}`

const testAccCheckDynZonePrimaryConfigSerialStyleMinute = `
resource "dyn_zone_primary" "foobar" {
	name = "%s"
	mailbox = "test@terraform.com"
	serial_style = "minute"
}`
