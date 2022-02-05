package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccExampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: testAccTagResourceConfig("eng"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("sonarr_tag.test", "label", "eng"),
					resource.TestCheckResourceAttrSet("sonarr_tag.test", "id"),
				),
			},
			// Update and Read testing
			{
				Config: testAccTagResourceConfig("1080p"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("sonarr_tag.test", "label", "1080p"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "sonarr_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccTagResourceConfig(label string) string {
	return fmt.Sprintf(`
resource "sonarr_tag" "test" {
  label = "%s"
}
`, label)
}
