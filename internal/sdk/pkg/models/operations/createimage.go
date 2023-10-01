// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openai/internal/sdk/pkg/models/shared"
)

type CreateImageResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ImagesResponse *shared.ImagesResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}
