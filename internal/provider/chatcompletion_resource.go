// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"openai/internal/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"openai/internal/validators"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &ChatCompletionResource{}
var _ resource.ResourceWithImportState = &ChatCompletionResource{}

func NewChatCompletionResource() resource.Resource {
	return &ChatCompletionResource{}
}

// ChatCompletionResource defines the resource implementation.
type ChatCompletionResource struct {
	client *sdk.Oai
}

// ChatCompletionResourceModel describes the resource data model.
type ChatCompletionResourceModel struct {
	Choices          []CreateChatCompletionResponseChoices    `tfsdk:"choices"`
	Created          types.Int64                              `tfsdk:"created"`
	FrequencyPenalty types.Number                             `tfsdk:"frequency_penalty"`
	FunctionCall     *CreateChatCompletionRequestFunctionCall `tfsdk:"function_call"`
	Functions        []ChatCompletionFunctions                `tfsdk:"functions"`
	ID               types.String                             `tfsdk:"id"`
	LogitBias        map[string]types.Int64                   `tfsdk:"logit_bias"`
	MaxTokens        types.Int64                              `tfsdk:"max_tokens"`
	Messages         []ChatCompletionRequestMessage           `tfsdk:"messages"`
	Model            types.String                             `tfsdk:"model"`
	N                types.Int64                              `tfsdk:"n"`
	Object           types.String                             `tfsdk:"object"`
	PresencePenalty  types.Number                             `tfsdk:"presence_penalty"`
	Stop             *CreateChatCompletionRequestStop         `tfsdk:"stop"`
	Stream           types.Bool                               `tfsdk:"stream"`
	Temperature      types.Number                             `tfsdk:"temperature"`
	TopP             types.Number                             `tfsdk:"top_p"`
	Usage            *CreateChatCompletionResponseUsage       `tfsdk:"usage"`
	User             types.String                             `tfsdk:"user"`
}

func (r *ChatCompletionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_chat_completion"
}

func (r *ChatCompletionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "ChatCompletion Resource",

		Attributes: map[string]schema.Attribute{
			"choices": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"finish_reason": schema.StringAttribute{
							Computed: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"stop",
									"length",
									"function_call",
								),
							},
							Description: `must be one of ["stop", "length", "function_call"]`,
						},
						"index": schema.Int64Attribute{
							Computed: true,
						},
						"message": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"content": schema.StringAttribute{
									Computed:    true,
									Description: `The contents of the message.`,
								},
								"function_call": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{
										"arguments": schema.StringAttribute{
											Computed:    true,
											Description: `The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.`,
										},
										"name": schema.StringAttribute{
											Computed:    true,
											Description: `The name of the function to call.`,
										},
									},
									Description: `The name and arguments of a function that should be called, as generated by the model.`,
								},
								"role": schema.StringAttribute{
									Computed: true,
									Validators: []validator.String{
										stringvalidator.OneOf(
											"system",
											"user",
											"assistant",
											"function",
										),
									},
									MarkdownDescription: `must be one of ["system", "user", "assistant", "function"]` + "\n" +
										`The role of the author of this message.`,
								},
							},
						},
					},
				},
			},
			"created": schema.Int64Attribute{
				Computed: true,
			},
			"frequency_penalty": schema.NumberAttribute{
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `Default: 0` + "\n" +
					`Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.` + "\n" +
					`` + "\n" +
					`[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)` + "\n" +
					``,
			},
			"function_call": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"create_chat_completion_request_function_call_1": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Optional: true,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"none",
								"auto",
							),
						},
						MarkdownDescription: `must be one of ["none", "auto"]` + "\n" +
							`Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the end-user. "auto" means the model can pick between an end-user or calling a function.  Specifying a particular function via ` + "`" + `{"name":\ "my_function"}` + "`" + ` forces the model to call that function. "none" is the default when no functions are present. "auto" is the default if functions are present.`,
					},
					"create_chat_completion_request_function_call_2": schema.SingleNestedAttribute{
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplace(),
						},
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
								Required:    true,
								Description: `The name of the function to call.`,
							},
						},
						Description: `Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the end-user. "auto" means the model can pick between an end-user or calling a function.  Specifying a particular function via ` + "`" + `{"name":\ "my_function"}` + "`" + ` forces the model to call that function. "none" is the default when no functions are present. "auto" is the default if functions are present.`,
					},
				},
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
				Description: `Controls how the model responds to function calls. "none" means the model does not call a function, and responds to the end-user. "auto" means the model can pick between an end-user or calling a function.  Specifying a particular function via ` + "`" + `{"name":\ "my_function"}` + "`" + ` forces the model to call that function. "none" is the default when no functions are present. "auto" is the default if functions are present.`,
			},
			"functions": schema.ListNestedAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"description": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Optional:    true,
							Description: `A description of what the function does, used by the model to choose when and how to call the function.`,
						},
						"name": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.`,
						},
						"parameters": schema.MapAttribute{
							PlanModifiers: []planmodifier.Map{
								mapplanmodifier.RequiresReplace(),
							},
							Required:    true,
							ElementType: types.StringType,
							Validators: []validator.Map{
								mapvalidator.ValueStringsAre(validators.IsValidJSON()),
							},
							MarkdownDescription: `The parameters the functions accepts, described as a JSON Schema object. See the [guide](/docs/guides/gpt/function-calling) for examples, and the [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for documentation about the format.` + "\n" +
								`` + "\n" +
								`To describe a function that accepts no parameters, provide the value ` + "`" + `{"type": "object", "properties": {}}` + "`" + `.`,
						},
					},
				},
				Description: `A list of functions the model may generate JSON inputs for.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"logit_bias": schema.MapAttribute{
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.RequiresReplace(),
				},
				Optional:    true,
				ElementType: types.Int64Type,
				MarkdownDescription: `Modify the likelihood of specified tokens appearing in the completion.` + "\n" +
					`` + "\n" +
					`Accepts a json object that maps tokens (specified by their token ID in the tokenizer) to an associated bias value from -100 to 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token.` + "\n" +
					``,
			},
			"max_tokens": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `Default: "inf"` + "\n" +
					`The maximum number of [tokens](/tokenizer) to generate in the chat completion.` + "\n" +
					`` + "\n" +
					`The total length of input tokens and generated tokens is limited by the model's context length. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.` + "\n" +
					``,
			},
			"messages": schema.ListNestedAttribute{
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplace(),
				},
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"content": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required:    true,
							Description: `The contents of the message. ` + "`" + `content` + "`" + ` is required for all messages, and may be null for assistant messages with function calls.`,
						},
						"function_call": schema.SingleNestedAttribute{
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplace(),
							},
							Optional: true,
							Attributes: map[string]schema.Attribute{
								"arguments": schema.StringAttribute{
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
									Required:    true,
									Description: `The arguments to call the function with, as generated by the model in JSON format. Note that the model does not always generate valid JSON, and may hallucinate parameters not defined by your function schema. Validate the arguments in your code before calling your function.`,
								},
								"name": schema.StringAttribute{
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
									Required:    true,
									Description: `The name of the function to call.`,
								},
							},
							Description: `The name and arguments of a function that should be called, as generated by the model.`,
						},
						"name": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Optional:    true,
							Description: `The name of the author of this message. ` + "`" + `name` + "`" + ` is required if role is ` + "`" + `function` + "`" + `, and it should be the name of the function whose response is in the ` + "`" + `content` + "`" + `. May contain a-z, A-Z, 0-9, and underscores, with a maximum length of 64 characters.`,
						},
						"role": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplace(),
							},
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"system",
									"user",
									"assistant",
									"function",
								),
							},
							MarkdownDescription: `must be one of ["system", "user", "assistant", "function"]` + "\n" +
								`The role of the messages author. One of ` + "`" + `system` + "`" + `, ` + "`" + `user` + "`" + `, ` + "`" + `assistant` + "`" + `, or ` + "`" + `function` + "`" + `.`,
						},
					},
				},
				Description: `A list of messages comprising the conversation so far. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_format_inputs_to_ChatGPT_models.ipynb).`,
			},
			"model": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						"gpt-4",
						"gpt-4-0314",
						"gpt-4-0613",
						"gpt-4-32k",
						"gpt-4-32k-0314",
						"gpt-4-32k-0613",
						"gpt-3.5-turbo",
						"gpt-3.5-turbo-16k",
						"gpt-3.5-turbo-0301",
						"gpt-3.5-turbo-0613",
						"gpt-3.5-turbo-16k-0613",
					),
				},
				MarkdownDescription: `must be one of ["gpt-4", "gpt-4-0314", "gpt-4-0613", "gpt-4-32k", "gpt-4-32k-0314", "gpt-4-32k-0613", "gpt-3.5-turbo", "gpt-3.5-turbo-16k", "gpt-3.5-turbo-0301", "gpt-3.5-turbo-0613", "gpt-3.5-turbo-16k-0613"]` + "\n" +
					`ID of the model to use. See the [model endpoint compatibility](/docs/models/model-endpoint-compatibility) table for details on which models work with the Chat API.`,
			},
			"n": schema.Int64Attribute{
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `Default: 1` + "\n" +
					`How many chat completion choices to generate for each input message.`,
			},
			"object": schema.StringAttribute{
				Computed: true,
			},
			"presence_penalty": schema.NumberAttribute{
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `Default: 0` + "\n" +
					`Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics.` + "\n" +
					`` + "\n" +
					`[See more information about frequency and presence penalties.](/docs/api-reference/parameter-details)` + "\n" +
					``,
			},
			"stop": schema.SingleNestedAttribute{
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplace(),
				},
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"str": schema.StringAttribute{
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
						Optional: true,
					},
					"array_ofstr": schema.ListAttribute{
						PlanModifiers: []planmodifier.List{
							listplanmodifier.RequiresReplace(),
						},
						Optional:    true,
						ElementType: types.StringType,
					},
				},
				Validators: []validator.Object{
					validators.ExactlyOneChild(),
				},
				MarkdownDescription: `Up to 4 sequences where the API will stop generating further tokens.` + "\n" +
					``,
			},
			"stream": schema.BoolAttribute{
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `Default: false` + "\n" +
					`If set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format) as they become available, with the stream terminated by a ` + "`" + `data: [DONE]` + "`" + ` message. [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_stream_completions.ipynb).` + "\n" +
					``,
			},
			"temperature": schema.NumberAttribute{
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `Default: 1` + "\n" +
					`What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.` + "\n" +
					`` + "\n" +
					`We generally recommend altering this or ` + "`" + `top_p` + "`" + ` but not both.` + "\n" +
					``,
			},
			"top_p": schema.NumberAttribute{
				PlanModifiers: []planmodifier.Number{
					numberplanmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `Default: 1` + "\n" +
					`An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 means only the tokens comprising the top 10% probability mass are considered.` + "\n" +
					`` + "\n" +
					`We generally recommend altering this or ` + "`" + `temperature` + "`" + ` but not both.` + "\n" +
					``,
			},
			"usage": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"completion_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"prompt_tokens": schema.Int64Attribute{
						Computed: true,
					},
					"total_tokens": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
			"user": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Optional: true,
				MarkdownDescription: `A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).` + "\n" +
					``,
			},
		},
	}
}

func (r *ChatCompletionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Oai)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.Oai, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *ChatCompletionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *ChatCompletionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	request := *data.ToCreateSDKType()
	res, err := r.client.OpenAI.CreateChatCompletion(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if res.CreateChatCompletionResponse == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromCreateResponse(res.CreateChatCompletionResponse)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChatCompletionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *ChatCompletionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; we rely entirely on CREATE API request response

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChatCompletionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *ChatCompletionResourceModel
	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *ChatCompletionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *ChatCompletionResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *ChatCompletionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.AddError("Not Implemented", "No available import state operation is available for resource chat_completion.")
}
