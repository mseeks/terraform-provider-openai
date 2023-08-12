// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Image - Represents the url or the content of an image generated by the OpenAI API.
type Image struct {
	// The base64-encoded JSON of the generated image, if `response_format` is `b64_json`.
	B64JSON *string `json:"b64_json,omitempty"`
	// The URL of the generated image, if `response_format` is `url` (default).
	URL *string `json:"url,omitempty"`
}
