// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateCompletionResponseChoicesLogprobs struct {
	TextOffset    []types.Int64  `tfsdk:"text_offset"`
	TokenLogprobs []types.Number `tfsdk:"token_logprobs"`
	Tokens        []types.String `tfsdk:"tokens"`
	TopLogprobs   types.String   `tfsdk:"top_logprobs"`
}
