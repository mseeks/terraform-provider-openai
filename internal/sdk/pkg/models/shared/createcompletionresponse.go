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

type CreateCompletionResponseChoices struct {
	FinishReason CreateCompletionResponseChoicesFinishReason `json:"finish_reason"`
	Index        int64                                       `json:"index"`
	Logprobs     CreateCompletionResponseChoicesLogprobs     `json:"logprobs"`
	Text         string                                      `json:"text"`
}

type CreateCompletionResponseUsage struct {
	CompletionTokens int64 `json:"completion_tokens"`
	PromptTokens     int64 `json:"prompt_tokens"`
	TotalTokens      int64 `json:"total_tokens"`
}

// CreateCompletionResponse - OK
type CreateCompletionResponse struct {
	Choices []CreateCompletionResponseChoices `json:"choices"`
	Created int64                             `json:"created"`
	ID      string                            `json:"id"`
	Model   string                            `json:"model"`
	Object  string                            `json:"object"`
	Usage   *CreateCompletionResponseUsage    `json:"usage,omitempty"`
}
