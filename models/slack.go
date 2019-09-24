package models

type SlackResponseAttachment struct {
	Text string `json:"text"`
}

type SlackResponse struct {
	ResponseType string `json:"response_type"`
	Text string `json:"text"`
	Attachments []SlackResponseAttachment `json:"attachments"`
}

