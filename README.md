# hello-cloudevents

`receiver` receives an event
`sender` generates an event

## Set up the event gateway



## Event Gateway API

[API documentation](https://github.com/serverless/event-gateway/blob/master/docs/api.md)

* Get EVENT_GATEWAY_IP from the IP address of event-gateway Kubernetes service.

* Create an event type called aquaEvent:

`curl --request POST --url http://$EVENT_GATEWAY_IP:4001/v1/spaces/default/eventtypes --header 'content-type: application/json' --data '{"name": "aquaEvent"}'`

* Register the receiver function which runs on the local machine

`curl --request POST --url http://$EVENT_GATEWAY_IP:4001/v1/spaces/default/functions --header 'content-type: application/json' --data '{"functionId": "receiver", "type": "http", "provider":{"url":"http://172.28.128.3:8080"}}'`

* Subscribe the function to the event

`curl --request POST --url http://$EVENT_GATEWAY_IP:4001/v1/spaces/default/subscriptions --header 'content-type: application/json' --data '{"functionId": "receiver", "event": "http", "type":"async", "eventType": "aquaEvent", "path": "/hello", "method": "POST"}'`

* Get event types

`curl --url http://$EVENT_GATEWAY_IP:4001/v1/spaces/default/eventtypes | jq`

* Get subscriptions

`curl --url http://$EVENT_GATEWAY_IP:4001/v1/spaces/default/subscriptions | jq`

* Get functions

`curl --url http://$EVENT_GATEWAY_IP:4001/v1/spaces/default/functions | jq`

