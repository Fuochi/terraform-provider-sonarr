---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sonarr_series Data Source - terraform-provider-sonarr"
subcategory: ""
description: |-
  List all available series
---

# sonarr_series (Data Source)

List all available series

## Example Usage

```terraform
data "sonarr_series" "example" {
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `id` (String) The ID of this resource.
- `series` (Attributes Set) List of series (see [below for nested schema](#nestedatt--series))

<a id="nestedatt--series"></a>
### Nested Schema for `series`

Read-Only:

- `id` (Number) ID of tag
- `language_profile_id` (Number) Language Profile ID
- `monitored` (Boolean) Monitored flag
- `path` (String) Series Path
- `quality_profile_id` (Number) Quality Profile ID
- `root_folder_path` (String) Series Root Folder
- `season_folder` (Boolean) Season Folder flag
- `tags` (Set of Number) Tags
- `title` (String) Series Title
- `title_slug` (String) Series Title in kebab format
- `tvdb_id` (Number) TVDB ID
- `use_scene_numbering` (Boolean) Scene numbering flag

