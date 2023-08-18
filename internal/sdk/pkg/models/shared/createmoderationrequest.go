// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type CreateModerationRequestInputType string

const (
	CreateModerationRequestInputTypeStr        CreateModerationRequestInputType = "str"
	CreateModerationRequestInputTypeArrayOfstr CreateModerationRequestInputType = "arrayOfstr"
)

type CreateModerationRequestInput struct {
	Str        *string
	ArrayOfstr []string

	Type CreateModerationRequestInputType
}

func CreateCreateModerationRequestInputStr(str string) CreateModerationRequestInput {
	typ := CreateModerationRequestInputTypeStr

	return CreateModerationRequestInput{
		Str:  &str,
		Type: typ,
	}
}

func CreateCreateModerationRequestInputArrayOfstr(arrayOfstr []string) CreateModerationRequestInput {
	typ := CreateModerationRequestInputTypeArrayOfstr

	return CreateModerationRequestInput{
		ArrayOfstr: arrayOfstr,
		Type:       typ,
	}
}

func (u *CreateModerationRequestInput) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	str := new(string)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&str); err == nil {
		u.Str = str
		u.Type = CreateModerationRequestInputTypeStr
		return nil
	}

	arrayOfstr := []string{}
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&arrayOfstr); err == nil {
		u.ArrayOfstr = arrayOfstr
		u.Type = CreateModerationRequestInputTypeArrayOfstr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u CreateModerationRequestInput) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return json.Marshal(u.Str)
	}

	if u.ArrayOfstr != nil {
		return json.Marshal(u.ArrayOfstr)
	}

	return nil, nil
}

// CreateModerationRequestModel - Two content moderations models are available: `text-moderation-stable` and `text-moderation-latest`.
//
// The default is `text-moderation-latest` which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use `text-moderation-stable`, we will provide advanced notice before updating the model. Accuracy of `text-moderation-stable` may be slightly lower than for `text-moderation-latest`.
type CreateModerationRequestModel string

const (
	CreateModerationRequestModelTextModerationLatest CreateModerationRequestModel = "text-moderation-latest"
	CreateModerationRequestModelTextModerationStable CreateModerationRequestModel = "text-moderation-stable"
)

func (e CreateModerationRequestModel) ToPointer() *CreateModerationRequestModel {
	return &e
}

func (e *CreateModerationRequestModel) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "text-moderation-latest":
		fallthrough
	case "text-moderation-stable":
		*e = CreateModerationRequestModel(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateModerationRequestModel: %v", v)
	}
}

type CreateModerationRequest struct {
	// The input text to classify
	Input CreateModerationRequestInput `json:"input"`
	// Two content moderations models are available: `text-moderation-stable` and `text-moderation-latest`.
	//
	// The default is `text-moderation-latest` which will be automatically upgraded over time. This ensures you are always using our most accurate model. If you use `text-moderation-stable`, we will provide advanced notice before updating the model. Accuracy of `text-moderation-stable` may be slightly lower than for `text-moderation-latest`.
	//
	Model *CreateModerationRequestModel `json:"model,omitempty"`
}
