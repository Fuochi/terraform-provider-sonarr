package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/devopsarr/terraform-provider-sonarr/internal/helpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golift.io/starr/sonarr"
)

const (
	indexerNewznabResourceName   = "indexer_newznab"
	IndexerNewznabImplementation = "Newznab"
	IndexerNewznabConfigContrat  = "NewznabSettings"
	IndexerNewznabProtocol       = "usenet"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &IndexerNewznabResource{}
var _ resource.ResourceWithImportState = &IndexerNewznabResource{}

func NewIndexerNewznabResource() resource.Resource {
	return &IndexerNewznabResource{}
}

// IndexerNewznabResource defines the Newznab indexer implementation.
type IndexerNewznabResource struct {
	client *sonarr.Sonarr
}

// IndexerNewznab describes the Newznab indexer data model.
type IndexerNewznab struct {
	Tags                      types.Set    `tfsdk:"tags"`
	Categories                types.Set    `tfsdk:"categories"`
	AnimeCategories           types.Set    `tfsdk:"anime_categories"`
	AdditionalParameters      types.String `tfsdk:"additional_parameters"`
	BaseURL                   types.String `tfsdk:"base_url"`
	APIPath                   types.String `tfsdk:"api_path"`
	APIKey                    types.String `tfsdk:"api_key"`
	Name                      types.String `tfsdk:"name"`
	ID                        types.Int64  `tfsdk:"id"`
	DownloadClientID          types.Int64  `tfsdk:"download_client_id"`
	Priority                  types.Int64  `tfsdk:"priority"`
	AnimeStandardFormatSearch types.Bool   `tfsdk:"anime_standard_format_search"`
	EnableRss                 types.Bool   `tfsdk:"enable_rss"`
	EnableInteractiveSearch   types.Bool   `tfsdk:"enable_interactive_search"`
	EnableAutomaticSearch     types.Bool   `tfsdk:"enable_automatic_search"`
}

func (i IndexerNewznab) toIndexer() *Indexer {
	return &Indexer{
		AnimeStandardFormatSearch: i.AnimeStandardFormatSearch,
		EnableAutomaticSearch:     i.EnableAutomaticSearch,
		EnableInteractiveSearch:   i.EnableInteractiveSearch,
		EnableRss:                 i.EnableRss,
		Priority:                  i.Priority,
		DownloadClientID:          i.DownloadClientID,
		ID:                        i.ID,
		Name:                      i.Name,
		AdditionalParameters:      i.AdditionalParameters,
		APIKey:                    i.APIKey,
		APIPath:                   i.APIKey,
		BaseURL:                   i.BaseURL,
		AnimeCategories:           i.AnimeCategories,
		Categories:                i.Categories,
		Tags:                      i.Tags,
	}
}

func (i *IndexerNewznab) fromIndexer(indexer *Indexer) {
	i.AnimeStandardFormatSearch = indexer.AnimeStandardFormatSearch
	i.EnableAutomaticSearch = indexer.EnableAutomaticSearch
	i.EnableInteractiveSearch = indexer.EnableInteractiveSearch
	i.EnableRss = indexer.EnableRss
	i.Priority = indexer.Priority
	i.DownloadClientID = indexer.DownloadClientID
	i.ID = indexer.ID
	i.Name = indexer.Name
	i.AdditionalParameters = indexer.AdditionalParameters
	i.APIKey = indexer.APIKey
	i.APIPath = indexer.APIPath
	i.BaseURL = indexer.BaseURL
	i.AnimeCategories = indexer.AnimeCategories
	i.Categories = indexer.Categories
	i.Tags = indexer.Tags
}

func (r *IndexerNewznabResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + indexerNewznabResourceName
}

func (r *IndexerNewznabResource) GetSchema(ctx context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		MarkdownDescription: "[subcategory:Indexers]: #\nIndexerNewznab resource.\nFor more information refer to [Indexer](https://wiki.servarr.com/sonarr/settings#indexers) documentation.",
		Attributes: map[string]tfsdk.Attribute{
			"enable_automatic_search": {
				MarkdownDescription: "Enable automatic search flag.",
				Optional:            true,
				Computed:            true,
				Type:                types.BoolType,
			},
			"enable_interactive_search": {
				MarkdownDescription: "Enable interactive search flag.",
				Optional:            true,
				Computed:            true,
				Type:                types.BoolType,
			},
			"enable_rss": {
				MarkdownDescription: "Enable RSS flag.",
				Optional:            true,
				Computed:            true,
				Type:                types.BoolType,
			},
			"priority": {
				MarkdownDescription: "Priority.",
				Optional:            true,
				Computed:            true,
				Type:                types.Int64Type,
			},
			"download_client_id": {
				MarkdownDescription: "Download client ID.",
				Optional:            true,
				Computed:            true,
				Type:                types.Int64Type,
			},
			"name": {
				MarkdownDescription: "IndexerNewznab name.",
				Required:            true,
				Type:                types.StringType,
			},
			"tags": {
				MarkdownDescription: "List of associated tags.",
				Optional:            true,
				Computed:            true,
				Type: types.SetType{
					ElemType: types.Int64Type,
				},
			},
			"id": {
				MarkdownDescription: "IndexerNewznab ID.",
				Computed:            true,
				Type:                types.Int64Type,
				PlanModifiers: tfsdk.AttributePlanModifiers{
					resource.UseStateForUnknown(),
				},
			},
			// Field values
			"anime_standard_format_search": {
				MarkdownDescription: "Search anime in standard format.",
				Optional:            true,
				Computed:            true,
				Type:                types.BoolType,
			},
			"additional_parameters": {
				MarkdownDescription: "Additional parameters.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"api_key": {
				MarkdownDescription: "API key.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"api_path": {
				MarkdownDescription: "API path.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"base_url": {
				MarkdownDescription: "Base URL.",
				Optional:            true,
				Computed:            true,
				Type:                types.StringType,
			},
			"categories": {
				MarkdownDescription: "Series list.",
				Optional:            true,
				Computed:            true,
				Type: types.SetType{
					ElemType: types.Int64Type,
				},
			},
			"anime_categories": {
				MarkdownDescription: "Anime list.",
				Optional:            true,
				Computed:            true,
				Type: types.SetType{
					ElemType: types.Int64Type,
				},
			},
		},
	}, nil
}

func (r *IndexerNewznabResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sonarr.Sonarr)
	if !ok {
		resp.Diagnostics.AddError(
			helpers.UnexpectedResourceConfigureType,
			fmt.Sprintf("Expected *sonarr.Sonarr, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *IndexerNewznabResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan IndexerNewznab

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create new IndexerNewznab
	request := readIndexerNewznab(ctx, &plan)

	response, err := r.client.AddIndexerContext(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, fmt.Sprintf("Unable to create %s, got error: %s", indexerNewznabResourceName, err))

		return
	}

	tflog.Trace(ctx, "created "+indexerNewznabResourceName+": "+strconv.Itoa(int(response.ID)))
	// Generate resource state struct
	result := writeIndexerNewznab(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &result)...)
}

func (r *IndexerNewznabResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state IndexerNewznab

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get IndexerNewznab current value
	response, err := r.client.GetIndexerContext(ctx, int(state.ID.Value))
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, fmt.Sprintf("Unable to read %s, got error: %s", indexerNewznabResourceName, err))

		return
	}

	tflog.Trace(ctx, "read "+indexerNewznabResourceName+": "+strconv.Itoa(int(response.ID)))
	// Map response body to resource schema attribute
	result := writeIndexerNewznab(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &result)...)
}

func (r *IndexerNewznabResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get plan values
	var plan IndexerNewznab

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update IndexerNewznab
	request := readIndexerNewznab(ctx, &plan)

	response, err := r.client.UpdateIndexerContext(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, fmt.Sprintf("Unable to update %s, got error: %s", indexerNewznabResourceName, err))

		return
	}

	tflog.Trace(ctx, "updated "+indexerNewznabResourceName+": "+strconv.Itoa(int(response.ID)))
	// Generate resource state struct
	result := writeIndexerNewznab(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, &result)...)
}

func (r *IndexerNewznabResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IndexerNewznab

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete IndexerNewznab current value
	err := r.client.DeleteIndexerContext(ctx, int(state.ID.Value))
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, fmt.Sprintf("Unable to read %s, got error: %s", indexerNewznabResourceName, err))

		return
	}

	tflog.Trace(ctx, "deleted "+indexerNewznabResourceName+": "+strconv.Itoa(int(state.ID.Value)))
	resp.State.RemoveResource(ctx)
}

func (r *IndexerNewznabResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			helpers.UnexpectedImportIdentifier,
			fmt.Sprintf("Expected import identifier with format: ID. Got: %q", req.ID),
		)

		return
	}

	tflog.Trace(ctx, "imported "+indexerNewznabResourceName+": "+strconv.Itoa(id))
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func writeIndexerNewznab(ctx context.Context, indexer *sonarr.IndexerOutput) *IndexerNewznab {
	var output IndexerNewznab

	genericIndexer := Indexer{
		EnableAutomaticSearch:   types.Bool{Value: indexer.EnableAutomaticSearch},
		EnableInteractiveSearch: types.Bool{Value: indexer.EnableInteractiveSearch},
		EnableRss:               types.Bool{Value: indexer.EnableRss},
		Priority:                types.Int64{Value: indexer.Priority},
		DownloadClientID:        types.Int64{Value: indexer.DownloadClientID},
		ID:                      types.Int64{Value: indexer.ID},
		Name:                    types.String{Value: indexer.Name},
		Tags:                    types.Set{ElemType: types.Int64Type},
		AnimeCategories:         types.Set{ElemType: types.Int64Type},
		Categories:              types.Set{ElemType: types.Int64Type},
	}
	tfsdk.ValueFrom(ctx, indexer.Tags, genericIndexer.Tags.Type(ctx), &genericIndexer.Tags)
	genericIndexer.writeIndexerFields(ctx, indexer.Fields)
	output.fromIndexer(&genericIndexer)

	return &output
}

func readIndexerNewznab(ctx context.Context, indexer *IndexerNewznab) *sonarr.IndexerInput {
	var tags []int

	tfsdk.ValueAs(ctx, indexer.Tags, &tags)

	return &sonarr.IndexerInput{
		EnableAutomaticSearch:   indexer.EnableAutomaticSearch.Value,
		EnableInteractiveSearch: indexer.EnableInteractiveSearch.Value,
		EnableRss:               indexer.EnableRss.Value,
		Priority:                indexer.Priority.Value,
		DownloadClientID:        indexer.DownloadClientID.Value,
		ID:                      indexer.ID.Value,
		ConfigContract:          IndexerNewznabConfigContrat,
		Implementation:          IndexerNewznabImplementation,
		Name:                    indexer.Name.Value,
		Protocol:                IndexerNewznabProtocol,
		Tags:                    tags,
		Fields:                  readIndexerFields(ctx, indexer.toIndexer()),
	}
}
