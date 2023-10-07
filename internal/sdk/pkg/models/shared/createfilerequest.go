// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
)

type CreateFileRequestFile struct {
	Content []byte `multipartForm:"content"`
	File    string `multipartForm:"name=file"`
}

type CreateFileRequest struct {
	// Name of the [JSON Lines](https://jsonlines.readthedocs.io/en/latest/) file to be uploaded.
	//
	// If the `purpose` is set to "fine-tune", each line is a JSON record with "prompt" and "completion" fields representing your [training examples](/docs/guides/fine-tuning/prepare-training-data).
	//
	File CreateFileRequestFile `multipartForm:"file"`
	// The intended purpose of the uploaded documents.
	//
	// Use "fine-tune" for [Fine-tuning](/docs/api-reference/fine-tunes). This allows us to validate the format of the uploaded file.
	//
	Purpose string `multipartForm:"name=purpose"`

	AdditionalProperties interface{} `json:"-"`
}
type _CreateFileRequest CreateFileRequest

func (c *CreateFileRequest) UnmarshalJSON(bs []byte) error {
	data := _CreateFileRequest{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = CreateFileRequest(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "")
	delete(additionalFields, "")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c CreateFileRequest) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_CreateFileRequest(c))
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
