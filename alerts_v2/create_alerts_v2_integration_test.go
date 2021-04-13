package alerts_v2_test

import (
	"github.com/logzio/logzio_terraform_client/alerts_v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegrationAlertsV2_CreateAlert(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2"
		alert, err := underTest.CreateAlert(createAlert)

		time.Sleep(4 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			defer underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_CreateAlertWithFilter(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 with filter"
		match := map[string]string{"type": "mytype"}
		must := map[string]interface{}{"match": match}
		createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must = append(createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must, must)
		alert, err := underTest.CreateAlert(createAlert)

		time.Sleep(4 * time.Second)
		if assert.NoError(t, err) && assert.NotNil(t, alert) {
			defer underTest.DeleteAlert(alert.AlertId)
		}
	}
}

func TestIntegrationAlertsV2_CreateAlertInvalidFilter(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 with invalid filter"
		createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must = append(createAlert.SubComponents[0].QueryDefinition.Filters.Bool.Must, nil)

		alert, err := underTest.CreateAlert(createAlert)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlertsV2_CreateAlertInvalidAggregationType(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 with invalid agg type"
		createAlert.SubComponents[0].QueryDefinition.Aggregation.AggregationType = "INVALID"

		alert, err := underTest.CreateAlert(createAlert)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlertsV2_CreateAlertInvalidValueAggregationTypeNone(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 with invalid val agg type none"
		createAlert.SubComponents[0].QueryDefinition.Aggregation.AggregationType = alerts_v2.AggregationTypeNone
		createAlert.SubComponents[0].QueryDefinition.Aggregation.FieldToAggregateOn = "hello"

		alert, err := underTest.CreateAlert(createAlert)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlertsV2_CreateAlertInvalidValueAggregationTypeCount(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 with invalid val agg type count"
		createAlert.SubComponents[0].QueryDefinition.Aggregation.AggregationType = alerts_v2.AggregationTypeCount
		createAlert.SubComponents[0].QueryDefinition.Aggregation.FieldToAggregateOn = "hello"

		alert, err := underTest.CreateAlert(createAlert)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}

func TestIntegrationAlertsV2_CreateAlertInvalidMail(t *testing.T) {
	underTest, err := setupAlertsV2IntegrationTest()

	if assert.NoError(t, err) {
		createAlert := getCreateAlertType()
		createAlert.Title = "test alerts v2 invalid mail"
		createAlert.Output.Recipients.Emails = []string{""}

		alert, err := underTest.CreateAlert(createAlert)
		assert.Error(t, err)
		assert.Nil(t, alert)
	}
}