---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sonarr_import_list Data Source - terraform-provider-sonarr"
subcategory: "Download Clients"
description: |-
  Single Download Client ../resources/import_list.
---

# sonarr_import_list (Data Source)

<!-- subcategory:Download Clients -->Single [Download Client](../resources/import_list).

## Example Usage

```terraform
data "sonarr_import_list" "example" {
  name = "Example"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Import List name.

### Read-Only

- `access_token` (String, Sensitive) Access token.
- `api_key` (String, Sensitive) API key.
- `auth_user` (String) Auth User.
- `base_url` (String) Base URL.
- `config_contract` (String) ImportList configuration template.
- `enable_automatic_add` (Boolean) Enable automatic add flag.
- `expires` (String) Expires.
- `genres` (String) Expires.
- `id` (Number) Import List ID.
- `implementation` (String) ImportList implementation name.
- `language_profile_id` (Number) Language profile ID.
- `language_profile_ids` (Set of Number) Language profile IDs.
- `limit` (Number) Limit.
- `listname` (String) Expires.
- `quality_profile_id` (Number) Quality profile ID.
- `quality_profile_ids` (Set of Number) Quality profile IDs.
- `rating` (String) Rating.
- `refresh_token` (String, Sensitive) Refresh token.
- `root_folder_path` (String) Root folder path.
- `season_folder` (Boolean) Season folder flag.
- `series_type` (String) Series type.
- `should_monitor` (String) Should monitor.
- `tag_ids` (Set of Number) Tag IDs.
- `tags` (Set of Number) List of associated tags.
- `trakt_additional_parameters` (String) Trakt additional parameters.
- `trakt_list_type` (Number) Trakt list type.
- `username` (String) Username.
- `years` (String) Expires.

