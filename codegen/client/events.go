package client

import (
	"github.com/bitsnap/chargebee-api-client/codegen/templates"
)

// GenerateListEvents generates chargebee client code to fetch all events
//
// API: https://{site}.chargebee.com/api/v2/events
//
// Documentation: https://apidocs.chargebee.com/docs/api/events?lang=curl#list_events
func GenerateListEvents() string {
	return templates.GenerateList("Events", "Event", "chargebee.com/api/v2/events")
}

// GenerateRetrieveEvent generates chargebee client code to retrieve specific event
//
// API: https://{site}.chargebee.com/api/v2/events/{event-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/events?lang=curl#retrieve_an_event
func GenerateRetrieveEvent() string {
	return templates.GenerateRetrieve("Event", "chargebee.com/api/v2/events")
}

// GenerateEventsModel generates chargebee event domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/events?lang=curl#event_attributes
func GenerateEventsModel() string {
	return templates.GenerateModel("Event", "Event attributes", "https://apidocs.chargebee.com/docs/api/events")
}
