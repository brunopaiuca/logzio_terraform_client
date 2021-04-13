package alerts_v2

import (
	"fmt"
	"github.com/logzio/logzio_terraform_client"
	"github.com/logzio/logzio_terraform_client/client"
	"io/ioutil"
	"net/http"
	"strings"
)

const deleteAlertServiceMethod string = http.MethodDelete
const deleteAlertServiceUrl = alertsServiceEndpoint + "/%d"
const deleteAlertMethodSuccess int = http.StatusOK

func (c *AlertsV2Client) buildDeleteApiRequest(apiToken string, alertId int64) (*http.Request, error) {
	baseUrl := c.BaseUrl
	req, err := http.NewRequest(deleteAlertServiceMethod, fmt.Sprintf(deleteAlertServiceUrl, baseUrl, alertId), nil)
	logzio_client.AddHttpHeaders(apiToken, req)

	return req, err
}

// Delete an alert, specified by it's unique id, returns an error if a problem is encountered
func (c *AlertsV2Client) DeleteAlert(alertId int64) error {
	req, _ := c.buildDeleteApiRequest(c.ApiToken, alertId)

	httpClient := client.GetHttpClient(req)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	jsonBytes, _ := ioutil.ReadAll(resp.Body)

	if !logzio_client.CheckValidStatus(resp, []int{deleteAlertMethodSuccess}) {
		return fmt.Errorf("API call %s failed with status code %d, data: %s", "DeleteAlert", resp.StatusCode, jsonBytes)
	}

	str := fmt.Sprintf("%s", jsonBytes)
	if strings.Contains(str, fmt.Sprintf("alert id %d not found", alertId)) {
		return fmt.Errorf("API call %s failed with missing alert %d, data: %s", "DeleteAlert", alertId, jsonBytes)
	}

	return nil
}