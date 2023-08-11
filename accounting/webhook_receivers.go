// This file was auto-generated by Fern from our API Definition.

package accounting

type WebhookReceiverRequest struct {
	// <span style="white-space: nowrap">`non-empty`</span>
	Event    string `json:"event"`
	IsActive bool   `json:"is_active"`
	// <span style="white-space: nowrap">`non-empty`</span>
	Key *string `json:"key,omitempty"`
}