// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateModerationResponseResultsCategories - A list of the categories, and whether they are flagged or not.
type CreateModerationResponseResultsCategories struct {
	// Whether the content was flagged as 'hate'.
	Hate bool `json:"hate"`
	// Whether the content was flagged as 'hate/threatening'.
	HateThreatening bool `json:"hate/threatening"`
	// Whether the content was flagged as 'self-harm'.
	SelfHarm bool `json:"self-harm"`
	// Whether the content was flagged as 'sexual'.
	Sexual bool `json:"sexual"`
	// Whether the content was flagged as 'sexual/minors'.
	SexualMinors bool `json:"sexual/minors"`
	// Whether the content was flagged as 'violence'.
	Violence bool `json:"violence"`
	// Whether the content was flagged as 'violence/graphic'.
	ViolenceGraphic bool `json:"violence/graphic"`
}

// CreateModerationResponseResultsCategoryScores - A list of the categories along with their scores as predicted by model.
type CreateModerationResponseResultsCategoryScores struct {
	// The score for the category 'hate'.
	Hate float64 `json:"hate"`
	// The score for the category 'hate/threatening'.
	HateThreatening float64 `json:"hate/threatening"`
	// The score for the category 'self-harm'.
	SelfHarm float64 `json:"self-harm"`
	// The score for the category 'sexual'.
	Sexual float64 `json:"sexual"`
	// The score for the category 'sexual/minors'.
	SexualMinors float64 `json:"sexual/minors"`
	// The score for the category 'violence'.
	Violence float64 `json:"violence"`
	// The score for the category 'violence/graphic'.
	ViolenceGraphic float64 `json:"violence/graphic"`
}

type CreateModerationResponseResults struct {
	// A list of the categories, and whether they are flagged or not.
	Categories CreateModerationResponseResultsCategories `json:"categories"`
	// A list of the categories along with their scores as predicted by model.
	CategoryScores CreateModerationResponseResultsCategoryScores `json:"category_scores"`
	// Whether the content violates [OpenAI's usage policies](/policies/usage-policies).
	Flagged bool `json:"flagged"`
}

// CreateModerationResponse - Represents policy compliance report by OpenAI's content moderation model against a given input.
type CreateModerationResponse struct {
	// The unique identifier for the moderation request.
	ID string `json:"id"`
	// The model used to generate the moderation results.
	Model string `json:"model"`
	// A list of moderation objects.
	Results []CreateModerationResponseResults `json:"results"`
}
