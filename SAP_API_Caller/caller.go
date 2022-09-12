package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-xxxxxxxx-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetXXXXXXXX(xxxxxxxx, xxxxxxxx string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "XXXXXXXX":
			func() {
				c.XXXXXXXX(xxxxxxxx)
				wg.Done()
			}()
		case "XXXXXXXX":
			func() {
				c.XXXXXXXX(xxxxxxxx, xxxxxxxx)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) XXXXXXXX(xxxxxxxx string) {
	xxxxxxxxData, err := c.callXXXXXXXXSrvAPIRequirementXXXXXXXX("A_XXXXXXXX", xxxxxxxx)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(xxxxxxxxData)

	xxxxxxxxData, err := c.callToXXXXXXXX(xxxxxxxxData[0].ToXXXXXXXX)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(xxxxxxxxData)
}

func (c *SAPAPICaller) callXXXXXXXXSrvAPIRequirementXXXXXXXX(api, xxxxxxxx string) ([]sap_api_output_formatter.XXXXXXXX, error) {
	url := strings.Join([]string{c.baseURL, "API_XXXXXXXX_SRV", api}, "/")
	param := c.getQueryWithXXXXXXXX(map[string]string{}, xxxxxxxx)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToXXXXXXXX(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToXXXXXXXX(url string) ([]sap_api_output_formatter.XXXXXXXX, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToXXXXXXXX(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) XXXXXXXX(xxxxxxxx, xxxxxxxx string) {
	data, err := c.callXXXXXXXXSrvAPIRequirementXXXXXXXX("A_XXXXXXXX", xxxxxxxx, xxxxxxxx)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callXXXXXXXXSrvAPIRequirementXXXXXXXX(api, xxxxxxxx string) ([]sap_api_output_formatter.XXXXXXXX, error) {
	url := strings.Join([]string{c.baseURL, "API_XXXXXXXX_SRV", api}, "/")

	param := c.getQueryWithXXXXXXXX(map[string]string{}, xxxxxxxx, xxxxxxxx)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToXXXXXXXX(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) getQueryWithXXXXXXXX(params map[string]string, xxxxxxxx string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("XXXXXXXX eq '%s'", xxxxxxxx)
	return params
}

func (c *SAPAPICaller) getQueryWithXXXXXXXX(params map[string]string, xxxxxxxx, xxxxxxxx string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("XXXXXXXX eq '%s' and XXXXXXXX eq '%s'", xxxxxxxx, xxxxxxxx)
	return params
}
