package opswatProvider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	opswatClient "terraform-provider-opswat/opswat/connectivity"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &Workflows{}
	_ datasource.DataSourceWithConfigure = &Workflows{}
)

// NewGlobalWorkflows is a helper function to simplify the provider implementation.
func NewWorkflows() datasource.DataSource {
	return &Workflows{}
}

// Workflows is the data source implementation.
type Workflows struct {
	client *opswatClient.Client
}

// Metadata returns the data source type name.
func (d *Workflows) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workflows"
}

// Schema defines the schema for the data source.
func (d *Workflows) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Global workflows datasource.",
		Attributes: map[string]schema.Attribute{
			"workflows": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_cert": schema.BoolAttribute{
							Description: "Generate batch signature with certificate - Use certificate to generate batch signature flag.",
							Computed:    true,
						},
						"allow_cert_cert": schema.StringAttribute{
							Description: "Certificate used for barch signing.",
							Computed:    true,
						},
						"allow_cert_cert_validity": schema.Int64Attribute{
							Description: "Certificate validity (hours).",
							Computed:    true,
						},
						"allow_local_files": schema.BoolAttribute{
							Description: "Process files from servers - Allow scan on server flag.",
							Computed:    true,
						},
						"allow_local_files_white_list": schema.BoolAttribute{
							Description: "Process files from servers flag (false = ALLOW ALL EXCEPT, true = DENY ALL EXCEPT).",
							Computed:    true,
						},
						"allow_local_files_local_paths": schema.ListAttribute{
							ElementType: types.StringType,
							Description: "Paths.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "Workflow description.",
							Computed:    true,
						},
						"id": schema.Int64Attribute{
							Description: "Workflow id.",
							Computed:    true,
						},
						"include_webhook_signature": schema.BoolAttribute{
							Description: "Webhook - Include webhook signature flag.",
							Computed:    true,
						},
						"include_webhook_signature_certificate_id": schema.Int64Attribute{
							Description: "Webhook - Certificate id.",
							Computed:    true,
						},
						"last_modified": schema.Int64Attribute{
							Description: "Last modified timestamp (unix epoch).",
							Computed:    true,
						},
						"mutable": schema.BoolAttribute{
							Description: "Mutable flag.",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Workflow name.",
							Computed:    true,
						},
						"workflow_id": schema.Int64Attribute{
							Description: "Workflow id.",
							Computed:    true,
						},
						"zone_id": schema.Int64Attribute{
							Description: "Workflow network zone id.",
							Computed:    true,
						},
						"scan_allowed": schema.ListAttribute{
							ElementType: types.Int64Type,
							Description: "Restrictions - Restrict access to following roles",
							Computed:    true,
						},
						"pref_hashes": schema.ObjectAttribute{
							AttributeTypes: map[string]attr.Type{
								"ds_advanced_setting_hash": types.StringType,
							},
							Description: "Pref hashes",
							Computed:    true,
						},
						"result_allowed": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"role": schema.Int64Attribute{
										Computed: true,
									},
									"visibility": schema.Int64Attribute{
										Computed: true,
									},
								},
							},
						},
						//scan_allowed
						//result_allowed

						//option_values
						//user_agents
					},
				},
			},
		},
	}
}

func (d *Workflows) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*opswatClient.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *opswatClient.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

type workflowsDataSourceModel struct {
	Workflows []workflowModel `tfsdk:"workflows"`
}

type workflowModel struct {
	AllowCert                            types.Bool           `tfsdk:"allow_cert"`
	AllowCertCert                        types.String         `tfsdk:"allow_cert_cert"`
	AllowCertCertValidity                types.Int64          `tfsdk:"allow_cert_cert_validity"`
	AllowLocalFiles                      types.Bool           `tfsdk:"allow_local_files"`
	AllowLocalFilesWhiteList             types.Bool           `tfsdk:"allow_local_files_white_list"`
	AllowLocalFilesLocalPaths            []string             `tfsdk:"allow_local_files_local_paths"`
	Description                          types.String         `tfsdk:"description"`
	ID                                   types.Int64          `tfsdk:"id"`
	IncludeWebhookSignature              types.Bool           `tfsdk:"include_webhook_signature"`
	IncludeWebhookSignatureCertificateID types.Int64          `tfsdk:"include_webhook_signature_certificate_id"`
	LastModified                         types.Int64          `tfsdk:"last_modified"`
	Mutable                              types.Bool           `tfsdk:"mutable"`
	Name                                 types.String         `tfsdk:"name"`
	WorkflowID                           types.Int64          `tfsdk:"workflow_id"`
	ZoneID                               types.Int64          `tfsdk:"zone_id"`
	ScanAllowed                          []interface{}        `tfsdk:"scan_allowed"`
	PrefHashes                           PrefHashesModel      `tfsdk:"pref_hashes"`
	ResultAllowed                        []ResultAllowedModel `tfsdk:"result_allowed"`
}

// PrefHashesModel
type PrefHashesModel struct {
	DSAdvancedSettingHash types.String `tfsdk:"ds_advanced_setting_hash"`
}

// ResultAllowModel
type ResultAllowedModel struct {
	Role       types.Int64 `tfsdk:"role"`
	Visibility types.Int64 `tfsdk:"visibility"`
}

// Read refreshes the Terraform state with the latest data.
func (d *Workflows) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state workflowsDataSourceModel

	result, err := d.client.GetWorkflows()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read OPSWAT workflows",
			err.Error(),
		)
		return
	}

	//fmt.Println("WORKFLOWS")
	//fmt.Printf("Workflows : %+v", result)

	//fmt.Println("RESULT")
	//fmt.Printf("Workflows : %+v", result)

	for _, workflow := range result {
		workflowState := workflowModel{
			AllowCert:                            types.BoolValue(workflow.AllowCert),
			AllowCertCert:                        types.StringValue(workflow.AllowCertCert),
			AllowCertCertValidity:                types.Int64Value(int64(workflow.AllowCertCertValidity)),
			AllowLocalFiles:                      types.BoolValue(workflow.AllowLocalFiles),
			AllowLocalFilesWhiteList:             types.BoolValue(workflow.AllowLocalFilesWhiteList),
			AllowLocalFilesLocalPaths:            append(workflow.AllowLocalFilesLocalPaths),
			Description:                          types.StringValue(workflow.Description),
			ID:                                   types.Int64Value(int64(workflow.Id)),
			IncludeWebhookSignature:              types.BoolValue(workflow.IncludeWebhookSignature),
			IncludeWebhookSignatureCertificateID: types.Int64Value(int64(workflow.IncludeWebhookSignatureWebhookCertificateId)),
			LastModified:                         types.Int64Value(int64(workflow.LastModified)),
			Mutable:                              types.BoolValue(workflow.Mutable),
			Name:                                 types.StringValue(workflow.Name),
			WorkflowID:                           types.Int64Value(int64(workflow.WorkflowId)),
			ZoneID:                               types.Int64Value(int64(workflow.ZoneId)),
			ScanAllowed:                          append(workflow.ScanAllowed),
			PrefHashes:                           PrefHashesModel{DSAdvancedSettingHash: types.StringValue(workflow.PrefHashes.DSADVANCEDSETTINGHASH)},
		}

		//fmt.Println("PARSED WORKFLOWS") test
		//spew.Dump(workflowState)

		for _, resultsallowed := range workflow.ResultAllowed {
			workflowState.ResultAllowed = append(workflowState.ResultAllowed, ResultAllowedModel{
				Role:       types.Int64Value(int64(resultsallowed.Role)),
				Visibility: types.Int64Value(int64(resultsallowed.Visibility)),
			})
		}

		//workflowState.PrefHashes = append(workflowState.PrefHashes, PrefHashesModel{
		//	DSAdvancedSettingHash: types.Int64Value(int64(prefhashes.DSAdvancedSettingHash)),
		//})

		state.Workflows = append(state.Workflows, workflowState)

	}
	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
