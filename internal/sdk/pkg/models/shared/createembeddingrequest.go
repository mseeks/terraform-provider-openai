// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type CreateEmbeddingRequestInputType string

const (
	CreateEmbeddingRequestInputTypeStr                   CreateEmbeddingRequestInputType = "str"
	CreateEmbeddingRequestInputTypeArrayOfstr            CreateEmbeddingRequestInputType = "arrayOfstr"
	CreateEmbeddingRequestInputTypeArrayOfinteger        CreateEmbeddingRequestInputType = "arrayOfinteger"
	CreateEmbeddingRequestInputTypeArrayOfarrayOfinteger CreateEmbeddingRequestInputType = "arrayOfarrayOfinteger"
)

type CreateEmbeddingRequestInput struct {
	Str                   *string
	ArrayOfstr            []string
	ArrayOfinteger        []int64
	ArrayOfarrayOfinteger [][]int64

	Type CreateEmbeddingRequestInputType
}

func CreateCreateEmbeddingRequestInputStr(str string) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeStr

	return CreateEmbeddingRequestInput{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateEmbeddingRequestInputArrayOfstr(arrayOfstr []string) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeArrayOfstr

	return CreateEmbeddingRequestInput{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func CreateCreateEmbeddingRequestInputArrayOfinteger(arrayOfinteger []int64) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeArrayOfinteger

	return CreateEmbeddingRequestInput{
		ArrayOfinteger: arrayOfinteger,
		Type:           typ,
	}
}

func CreateCreateEmbeddingRequestInputArrayOfarrayOfinteger(arrayOfarrayOfinteger [][]int64) CreateEmbeddingRequestInput {
	typ := CreateEmbeddingRequestInputTypeArrayOfarrayOfinteger

	return CreateEmbeddingRequestInput{
		ArrayOfarrayOfinteger: arrayOfarrayOfinteger,
		Type:                  typ,
	}
}

func (u *CreateEmbeddingRequestInput) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = CreateEmbeddingRequestInputTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateEmbeddingRequestInputTypeArrayOfstr
		return nil
	}

	arrayOfinteger := []int64{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfinteger); err == nil {
		u.ArrayOfinteger = arrayOfinteger
		u.Type = CreateEmbeddingRequestInputTypeArrayOfinteger
		return nil
	}

	arrayOfarrayOfinteger := [][]int64{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfarrayOfinteger); err == nil {
		u.ArrayOfarrayOfinteger = arrayOfarrayOfinteger
		u.Type = CreateEmbeddingRequestInputTypeArrayOfarrayOfinteger
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateEmbeddingRequestInput) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	if u.ArrayOfinteger != nil {
		return json.Marshal(u.ArrayOfinteger)
	}

	if u.ArrayOfarrayOfinteger != nil {
		return json.Marshal(u.ArrayOfarrayOfinteger)
	}

	return nil, nil
}

// CreateEmbeddingRequestModel - ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
type CreateEmbeddingRequestModel string

const (
	CreateEmbeddingRequestModelTextEmbeddingAda002 CreateEmbeddingRequestModel = "text-embedding-ada-002"
)

func (e CreateEmbeddingRequestModel) ToPointer() *CreateEmbeddingRequestModel {
	return &e
}

func (e *CreateEmbeddingRequestModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-embedding-ada-002":
		*e = CreateEmbeddingRequestModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateEmbeddingRequestModel: %v", v)
	}
}

type CreateEmbeddingRequest struct {
	// Input text to embed, encoded as a string or array of tokens. To embed multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed the max input tokens for the model (8191 tokens for `text-embedding-ada-002`). [Example Python code](https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb) for counting tokens.
	//
	Input CreateEmbeddingRequestInput `json:"input"`
	// ID of the model to use. You can use the [List models](/docs/api-reference/models/list) API to see all of your available models, or see our [Model overview](/docs/models/overview) for descriptions of them.
	//
	Model CreateEmbeddingRequestModel `json:"model"`
	// A unique identifier representing your end-user, which can help OpenAI to monitor and detect abuse. [Learn more](/docs/guides/safety-best-practices/end-user-ids).
	//
	User *string `json:"user,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _CreateEmbeddingRequest CreateEmbeddingRequest

func (c *CreateEmbeddingRequest) UnmarshalJSON(bs []byte) error {
	data := _CreateEmbeddingRequest{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = CreateEmbeddingRequest(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "input")
	delete(additionalFields, "model")
	delete(additionalFields, "user")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c CreateEmbeddingRequest) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_CreateEmbeddingRequest(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}
