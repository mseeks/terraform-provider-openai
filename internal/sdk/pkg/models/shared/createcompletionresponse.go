// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CreateCompletionResponseChoicesFinishReason string

const (
	CreateCompletionResponseChoicesFinishReasonStop   CreateCompletionResponseChoicesFinishReason = "stop"
	CreateCompletionResponseChoicesFinishReasonLength CreateCompletionResponseChoicesFinishReason = "length"
)

func (e CreateCompletionResponseChoicesFinishReason) ToPointer() *CreateCompletionResponseChoicesFinishReason {
	return &e
}

func (e *CreateCompletionResponseChoicesFinishReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "stop":
		fallthrough
	case "length":
		*e = CreateCompletionResponseChoicesFinishReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCompletionResponseChoicesFinishReason: %v", v)
	}
}

type CreateCompletionResponseChoicesLogprobs struct {
	TextOffset    []int64     `json:"text_offset,omitempty"`
	TokenLogprobs []float64   `json:"token_logprobs,omitempty"`
	Tokens        []string    `json:"tokens,omitempty"`
	TopLogprobs   interface{} `json:"top_logprobs,omitempty"`
}

func (o *CreateCompletionResponseChoicesLogprobs) GetTextOffset() []int64 {
	if o == nil {
		return nil
	}
	return o.TextOffset
}

func (o *CreateCompletionResponseChoicesLogprobs) GetTokenLogprobs() []float64 {
	if o == nil {
		return nil
	}
	return o.TokenLogprobs
}

func (o *CreateCompletionResponseChoicesLogprobs) GetTokens() []string {
	if o == nil {
		return nil
	}
	return o.Tokens
}

func (o *CreateCompletionResponseChoicesLogprobs) GetTopLogprobs() interface{} {
	if o == nil {
		return nil
	}
	return o.TopLogprobs
}

type CreateCompletionResponseChoices struct {
	FinishReason CreateCompletionResponseChoicesFinishReason `json:"finish_reason"`
	Index        int64                                       `json:"index"`
	Logprobs     *CreateCompletionResponseChoicesLogprobs    `json:"logprobs"`
	Text         string                                      `json:"text"`
}

func (o *CreateCompletionResponseChoices) GetFinishReason() CreateCompletionResponseChoicesFinishReason {
	if o == nil {
		return CreateCompletionResponseChoicesFinishReason("")
	}
	return o.FinishReason
}

func (o *CreateCompletionResponseChoices) GetIndex() int64 {
	if o == nil {
		return 0
	}
	return o.Index
}

func (o *CreateCompletionResponseChoices) GetLogprobs() *CreateCompletionResponseChoicesLogprobs {
	if o == nil {
		return nil
	}
	return o.Logprobs
}

func (o *CreateCompletionResponseChoices) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

type CreateCompletionResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

func (o *CreateCompletionResponseUsage) GetCompletionTokens() int64 {
	if o == nil {
		return 0
	}
	return o.CompletionTokens
}

func (o *CreateCompletionResponseUsage) GetPromptTokens() int64 {
	if o == nil {
		return 0
	}
	return o.PromptTokens
}

func (o *CreateCompletionResponseUsage) GetTotalTokens() int64 {
	if o == nil {
		return 0
	}
	return o.TotalTokens
}

type CreateCompletionResponse struct {
	Choices []CreateCompletionResponseChoices `json:"choices"`
	Created int64                             `json:"created"`
	ID      string                            `json:"id"`
	Model   string                            `json:"model"`
	Object  string                            `json:"object"`
	Usage   *CreateCompletionResponseUsage    `json:"usage,omitempty"`
}

func (o *CreateCompletionResponse) GetChoices() []CreateCompletionResponseChoices {
	if o == nil {
		return []CreateCompletionResponseChoices{}
	}
	return o.Choices
}

func (o *CreateCompletionResponse) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *CreateCompletionResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateCompletionResponse) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *CreateCompletionResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}

func (o *CreateCompletionResponse) GetUsage() *CreateCompletionResponseUsage {
	if o == nil {
		return nil
	}
	return o.Usage
}
