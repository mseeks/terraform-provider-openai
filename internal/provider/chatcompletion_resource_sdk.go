// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"openai/internal/sdk/pkg/models/shared"
)

func (r *ChatCompletionResourceModel) ToCreateSDKType() *shared.CreateChatCompletionRequest {
	frequencyPenalty := new(float64)
	if !r.FrequencyPenalty.IsUnknown() && !r.FrequencyPenalty.IsNull() {
		*frequencyPenalty, _ = r.FrequencyPenalty.ValueBigFloat().Float64()
	} else {
		frequencyPenalty = nil
	}
	var functionCall *shared.CreateChatCompletionRequestFunctionCall
	if r.FunctionCall != nil {
		createChatCompletionRequestFunctionCall1 := new(shared.CreateChatCompletionRequestFunctionCall1)
		if !r.FunctionCall.CreateChatCompletionRequestFunctionCall1.IsUnknown() && !r.FunctionCall.CreateChatCompletionRequestFunctionCall1.IsNull() {
			*createChatCompletionRequestFunctionCall1 = shared.CreateChatCompletionRequestFunctionCall1(r.FunctionCall.CreateChatCompletionRequestFunctionCall1.ValueString())
		} else {
			createChatCompletionRequestFunctionCall1 = nil
		}
		if createChatCompletionRequestFunctionCall1 != nil {
			functionCall = &shared.CreateChatCompletionRequestFunctionCall{
				CreateChatCompletionRequestFunctionCall1: createChatCompletionRequestFunctionCall1,
			}
		}
		var createChatCompletionRequestFunctionCall2 *shared.CreateChatCompletionRequestFunctionCall2
		if r.FunctionCall.CreateChatCompletionRequestFunctionCall2 != nil {
			name := r.FunctionCall.CreateChatCompletionRequestFunctionCall2.Name.ValueString()
			createChatCompletionRequestFunctionCall2 = &shared.CreateChatCompletionRequestFunctionCall2{
				Name: name,
			}
		}
		if createChatCompletionRequestFunctionCall2 != nil {
			functionCall = &shared.CreateChatCompletionRequestFunctionCall{
				CreateChatCompletionRequestFunctionCall2: createChatCompletionRequestFunctionCall2,
			}
		}
	}
	var functions []shared.ChatCompletionFunctions = nil
	for _, functionsItem := range r.Functions {
		description := new(string)
		if !functionsItem.Description.IsUnknown() && !functionsItem.Description.IsNull() {
			*description = functionsItem.Description.ValueString()
		} else {
			description = nil
		}
		name1 := functionsItem.Name.ValueString()
		parameters := make(map[string]interface{})
		for parametersKey, parametersValue := range functionsItem.Parameters {
			var parametersInst interface{}
			_ = json.Unmarshal([]byte(parametersValue.ValueString()), &parametersInst)
			parameters[parametersKey] = parametersInst
		}
		functions = append(functions, shared.ChatCompletionFunctions{
			Description: description,
			Name:        name1,
			Parameters:  parameters,
		})
	}
	logitBias := make(map[string]int64)
	for logitBiasKey, logitBiasValue := range r.LogitBias {
		logitBiasInst := logitBiasValue.ValueInt64()
		logitBias[logitBiasKey] = logitBiasInst
	}
	maxTokens := new(int64)
	if !r.MaxTokens.IsUnknown() && !r.MaxTokens.IsNull() {
		*maxTokens = r.MaxTokens.ValueInt64()
	} else {
		maxTokens = nil
	}
	var messages []shared.ChatCompletionRequestMessage = nil
	for _, messagesItem := range r.Messages {
		content := new(string)
		if !messagesItem.Content.IsUnknown() && !messagesItem.Content.IsNull() {
			*content = messagesItem.Content.ValueString()
		} else {
			content = nil
		}
		var functionCall1 *shared.ChatCompletionRequestMessageFunctionCall
		if messagesItem.FunctionCall != nil {
			arguments := messagesItem.FunctionCall.Arguments.ValueString()
			name2 := messagesItem.FunctionCall.Name.ValueString()
			functionCall1 = &shared.ChatCompletionRequestMessageFunctionCall{
				Arguments: arguments,
				Name:      name2,
			}
		}
		name3 := new(string)
		if !messagesItem.Name.IsUnknown() && !messagesItem.Name.IsNull() {
			*name3 = messagesItem.Name.ValueString()
		} else {
			name3 = nil
		}
		role := shared.ChatCompletionRequestMessageRole(messagesItem.Role.ValueString())
		messages = append(messages, shared.ChatCompletionRequestMessage{
			Content:      content,
			FunctionCall: functionCall1,
			Name:         name3,
			Role:         role,
		})
	}
	model := shared.CreateChatCompletionRequestModel(r.Model.ValueString())
	n := new(int64)
	if !r.N.IsUnknown() && !r.N.IsNull() {
		*n = r.N.ValueInt64()
	} else {
		n = nil
	}
	presencePenalty := new(float64)
	if !r.PresencePenalty.IsUnknown() && !r.PresencePenalty.IsNull() {
		*presencePenalty, _ = r.PresencePenalty.ValueBigFloat().Float64()
	} else {
		presencePenalty = nil
	}
	var stop *shared.CreateChatCompletionRequestStop
	if r.Stop != nil {
		str := new(string)
		if !r.Stop.Str.IsUnknown() && !r.Stop.Str.IsNull() {
			*str = r.Stop.Str.ValueString()
		} else {
			str = nil
		}
		if str != nil {
			stop = &shared.CreateChatCompletionRequestStop{
				Str: str,
			}
		}
		var arrayOfstr []string = nil
		for _, arrayOfstrItem := range r.Stop.ArrayOfstr {
			arrayOfstr = append(arrayOfstr, arrayOfstrItem.ValueString())
		}
		if arrayOfstr != nil {
			stop = &shared.CreateChatCompletionRequestStop{
				ArrayOfstr: arrayOfstr,
			}
		}
	}
	stream := new(bool)
	if !r.Stream.IsUnknown() && !r.Stream.IsNull() {
		*stream = r.Stream.ValueBool()
	} else {
		stream = nil
	}
	temperature := new(float64)
	if !r.Temperature.IsUnknown() && !r.Temperature.IsNull() {
		*temperature, _ = r.Temperature.ValueBigFloat().Float64()
	} else {
		temperature = nil
	}
	topP := new(float64)
	if !r.TopP.IsUnknown() && !r.TopP.IsNull() {
		*topP, _ = r.TopP.ValueBigFloat().Float64()
	} else {
		topP = nil
	}
	user := new(string)
	if !r.User.IsUnknown() && !r.User.IsNull() {
		*user = r.User.ValueString()
	} else {
		user = nil
	}
	out := shared.CreateChatCompletionRequest{
		FrequencyPenalty: frequencyPenalty,
		FunctionCall:     functionCall,
		Functions:        functions,
		LogitBias:        logitBias,
		MaxTokens:        maxTokens,
		Messages:         messages,
		Model:            model,
		N:                n,
		PresencePenalty:  presencePenalty,
		Stop:             stop,
		Stream:           stream,
		Temperature:      temperature,
		TopP:             topP,
		User:             user,
	}
	return &out
}

func (r *ChatCompletionResourceModel) RefreshFromCreateResponse(resp *shared.CreateChatCompletionResponse) {
	r.Choices = nil
	for _, choicesItem := range resp.Choices {
		var choices1 CreateChatCompletionResponseChoices
		choices1.FinishReason = types.StringValue(string(choicesItem.FinishReason))
		choices1.Index = types.Int64Value(choicesItem.Index)
		if choicesItem.Message.Content != nil {
			choices1.Message.Content = types.StringValue(*choicesItem.Message.Content)
		} else {
			choices1.Message.Content = types.StringNull()
		}
		if choicesItem.Message.FunctionCall == nil {
			choices1.Message.FunctionCall = nil
		} else {
			choices1.Message.FunctionCall = &ChatCompletionResponseMessageFunctionCall{}
			if choicesItem.Message.FunctionCall.Arguments != nil {
				choices1.Message.FunctionCall.Arguments = types.StringValue(*choicesItem.Message.FunctionCall.Arguments)
			} else {
				choices1.Message.FunctionCall.Arguments = types.StringNull()
			}
			if choicesItem.Message.FunctionCall.Name != nil {
				choices1.Message.FunctionCall.Name = types.StringValue(*choicesItem.Message.FunctionCall.Name)
			} else {
				choices1.Message.FunctionCall.Name = types.StringNull()
			}
		}
		choices1.Message.Role = types.StringValue(string(choicesItem.Message.Role))
		r.Choices = append(r.Choices, choices1)
	}
	r.Created = types.Int64Value(resp.Created)
	r.ID = types.StringValue(resp.ID)
	r.Model = types.StringValue(string(resp.Model))
	r.Object = types.StringValue(resp.Object)
	if resp.Usage == nil {
		r.Usage = nil
	} else {
		r.Usage = &CreateChatCompletionResponseUsage{}
		r.Usage.CompletionTokens = types.Int64Value(resp.Usage.CompletionTokens)
		r.Usage.PromptTokens = types.Int64Value(resp.Usage.PromptTokens)
		r.Usage.TotalTokens = types.Int64Value(resp.Usage.TotalTokens)
	}
}
