package provider

import (
	"context"
	"strconv"

	"github.com/devopsarr/sonarr-go/sonarr"
	"github.com/devopsarr/terraform-provider-sonarr/internal/helpers"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"golang.org/x/exp/slices"
)

const metadataResourceName = "metadata"

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = &MetadataResource{}
	_ resource.ResourceWithImportState = &MetadataResource{}
)

var metadataBoolFields = []string{"seriesMetadata", "seriesMetadataUrl", "seriesImages", "seasonImages", "episodeImages", "episodeMetadata"}

func NewMetadataResource() resource.Resource {
	return &MetadataResource{}
}

// MetadataResource defines the metadata implementation.
type MetadataResource struct {
	client *sonarr.APIClient
}

// Metadata describes the metadata data model.
type Metadata struct {
	Tags              types.Set    `tfsdk:"tags"`
	Name              types.String `tfsdk:"name"`
	ConfigContract    types.String `tfsdk:"config_contract"`
	Implementation    types.String `tfsdk:"implementation"`
	ID                types.Int64  `tfsdk:"id"`
	Enable            types.Bool   `tfsdk:"enable"`
	SeriesMetadata    types.Bool   `tfsdk:"series_metadata"`
	SeriesMetadataURL types.Bool   `tfsdk:"series_metadata_url"`
	SeriesImages      types.Bool   `tfsdk:"series_images"`
	SeasonImages      types.Bool   `tfsdk:"season_images"`
	EpisodeMetadata   types.Bool   `tfsdk:"episode_metadata"`
	EpisodeImages     types.Bool   `tfsdk:"episode_images"`
}

func (r *MetadataResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + metadataResourceName
}

func (r *MetadataResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "<!-- subcategory:Metadata -->Generic Metadata resource. When possible use a specific resource instead.\nFor more information refer to [Metadata](https://wiki.servarr.com/sonarr/settings#metadata) documentation.",
		Attributes: map[string]schema.Attribute{
			"enable": schema.BoolAttribute{
				MarkdownDescription: "Enable flag.",
				Optional:            true,
				Computed:            true,
			},
			"config_contract": schema.StringAttribute{
				MarkdownDescription: "Metadata configuration template.",
				Required:            true,
			},
			"implementation": schema.StringAttribute{
				MarkdownDescription: "Metadata implementation name.",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Metadata name.",
				Required:            true,
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: "List of associated tags.",
				Optional:            true,
				Computed:            true,
				ElementType:         types.Int64Type,
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "Metadata ID.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			// Field values
			"episode_metadata": schema.BoolAttribute{
				MarkdownDescription: "Episode metadata flag.",
				Optional:            true,
				Computed:            true,
			},
			"episode_images": schema.BoolAttribute{
				MarkdownDescription: "Episode images flag.",
				Optional:            true,
				Computed:            true,
			},
			"season_images": schema.BoolAttribute{
				MarkdownDescription: "Season images flag.",
				Optional:            true,
				Computed:            true,
			},
			"series_images": schema.BoolAttribute{
				MarkdownDescription: "Series images flag.",
				Optional:            true,
				Computed:            true,
			},
			"series_metadata": schema.BoolAttribute{
				MarkdownDescription: "Series metafata flag.",
				Optional:            true,
				Computed:            true,
			},
			"series_metadata_url": schema.BoolAttribute{
				MarkdownDescription: "Series metadata URL flag.",
				Optional:            true,
				Computed:            true,
			},
		},
	}
}

func (r *MetadataResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if client := helpers.ResourceConfigure(ctx, req, resp); client != nil {
		r.client = client
	}
}

func (r *MetadataResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var metadata *Metadata

	resp.Diagnostics.Append(req.Plan.Get(ctx, &metadata)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create new Metadata
	request := metadata.read(ctx)

	response, _, err := r.client.MetadataApi.CreateMetadata(ctx).MetadataResource(*request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Create, metadataResourceName, err))

		return
	}

	tflog.Trace(ctx, "created "+metadataResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Generate resource state struct.
	// this is needed because of many empty fields are unknown in both plan and read
	var state Metadata

	state.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *MetadataResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var metadata *Metadata

	resp.Diagnostics.Append(req.State.Get(ctx, &metadata)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get Metadata current value
	response, _, err := r.client.MetadataApi.GetMetadataById(ctx, int32(metadata.ID.ValueInt64())).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Read, metadataResourceName, err))

		return
	}

	tflog.Trace(ctx, "read "+metadataResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Generate resource state struct.
	// this is needed because of many empty fields are unknown in both plan and read
	var state Metadata

	state.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *MetadataResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Get plan values
	var metadata *Metadata

	resp.Diagnostics.Append(req.Plan.Get(ctx, &metadata)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Update Metadata
	request := metadata.read(ctx)

	response, _, err := r.client.MetadataApi.UpdateMetadata(ctx, strconv.Itoa(int(request.GetId()))).MetadataResource(*request).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Update, metadataResourceName, err))

		return
	}

	tflog.Trace(ctx, "updated "+metadataResourceName+": "+strconv.Itoa(int(response.GetId())))
	// Generate resource state struct.
	// this is needed because of many empty fields are unknown in both plan and read
	var state Metadata

	state.write(ctx, response)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *MetadataResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var metadata Metadata

	resp.Diagnostics.Append(req.State.Get(ctx, &metadata)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Delete Metadata current value
	_, err := r.client.MetadataApi.DeleteMetadata(ctx, int32(metadata.ID.ValueInt64())).Execute()
	if err != nil {
		resp.Diagnostics.AddError(helpers.ClientError, helpers.ParseClientError(helpers.Read, metadataResourceName, err))

		return
	}

	tflog.Trace(ctx, "deleted "+metadataResourceName+": "+strconv.Itoa(int(metadata.ID.ValueInt64())))
	resp.State.RemoveResource(ctx)
}

func (r *MetadataResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	helpers.ImportStatePassthroughIntID(ctx, path.Root("id"), req, resp)
	tflog.Trace(ctx, "imported "+metadataResourceName+": "+req.ID)
}

func (i *Metadata) write(ctx context.Context, metadata *sonarr.MetadataResource) {
	i.Enable = types.BoolValue(metadata.GetEnable())
	i.ID = types.Int64Value(int64(metadata.GetId()))
	i.ConfigContract = types.StringValue(metadata.GetConfigContract())
	i.Implementation = types.StringValue(metadata.GetImplementation())
	i.Name = types.StringValue(metadata.GetName())
	i.Tags = types.SetValueMust(types.Int64Type, nil)
	tfsdk.ValueFrom(ctx, metadata.Tags, i.Tags.Type(ctx), &i.Tags)
	i.writeFields(metadata.GetFields())
}

func (i *Metadata) writeFields(fields []*sonarr.Field) {
	for _, f := range fields {
		if f.Value == nil {
			continue
		}

		if slices.Contains(metadataBoolFields, f.GetName()) {
			helpers.WriteBoolField(f, i)

			continue
		}
	}
}

func (i *Metadata) read(ctx context.Context) *sonarr.MetadataResource {
	tags := make([]*int32, len(i.Tags.Elements()))
	tfsdk.ValueAs(ctx, i.Tags, &tags)

	metadata := sonarr.NewMetadataResource()
	metadata.SetEnable(i.Enable.ValueBool())
	metadata.SetId(int32(i.ID.ValueInt64()))
	metadata.SetConfigContract(i.ConfigContract.ValueString())
	metadata.SetImplementation(i.Implementation.ValueString())
	metadata.SetName(i.Name.ValueString())
	metadata.SetTags(tags)
	metadata.SetFields(i.readFields())

	return metadata
}

func (i *Metadata) readFields() []*sonarr.Field {
	var output []*sonarr.Field

	for _, b := range metadataBoolFields {
		if field := helpers.ReadBoolField(b, i); field != nil {
			output = append(output, field)
		}
	}

	return output
}