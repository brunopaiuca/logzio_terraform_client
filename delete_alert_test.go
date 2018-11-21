package logzio_client

import (
	"testing"
)

func TestDeleteAlert(t *testing.T) {
	api_token := getApiToken(t)

	var client *Client
	client = New(api_token)

	createAlert := createValidAlert()

	alert, err := client.CreateAlert(createAlert)
	if err != nil {
		t.Fatalf("%v should not have raised an error: %v", "DeleteAlert", err)
	}

	alertId := alert.AlertId
	client.DeleteAlert(alertId)

	alerts, err := client.ListAlerts()
	if containsAlert(alerts, alertId) {
		t.Fatalf("%v %d should have been deleted, but is returned by ListAlerts", "DeleteAlert", alertId)
	}

	err = client.DeleteAlert(12345)
	if err == nil {
		t.Fatalf("should have raised an error for unknown id: %v", err)
	}

}
