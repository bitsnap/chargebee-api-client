package chargebee

import (
	"github.com/bitsnap/chargebee-api-client/codegen/client"
)

// GenerateListEvents generates chargebee client code to fetch all events
//
// API: https://{site}.chargebee.com/api/v2/events
//
// Documentation: https://apidocs.chargebee.com/docs/api/events?lang=curl#list_events
func GenerateListEvents() string {
	return client.GenerateList("Events", "Event", "chargebee.com/api/v2/events")
}

// GenerateRetrieveEvent generates chargebee client code to retrieve specific event
//
// API: https://{site}.chargebee.com/api/v2/events/{event-id}
//
// Documentation: https://apidocs.chargebee.com/docs/api/events?lang=curl#retrieve_an_event
func GenerateRetrieveEvent() string {
	return client.GenerateRetrieve("Event", "chargebee.com/api/v2/events")
}

// GenerateEventsModel generates chargebee event domain model
//
// Documentation: https://apidocs.chargebee.com/docs/api/events?lang=curl#event_attributes
func GenerateEventsModel() string {
	return client.GenerateModel("Event", "Event attributes", "https://apidocs.chargebee.com/docs/api/events")
}
