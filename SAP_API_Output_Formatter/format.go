package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-xxxxxxxx-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToXXXXXXXX(raw []byte, l *logger.Logger) ([]XXXXXXXX, error) {
	pm := &responses.XXXXXXXX{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to XXXXXXXX. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	xxxxxxxx := make([]XXXXXXXX, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		xxxxxxxx = append(xxxxxxxx, XXXXXXXX{
			XXXXXXXX:            data.XXXXXXXX,
			ToXXXXXXXX:          data.ToXXXXXXXX.Deferred.URI,
		})
	}

	return xxxxxxxx, nil
}

func ConvertToXXXXXXXX(raw []byte, l *logger.Logger) ([]XXXXXXXX, error) {
	pm := &responses.XXXXXXXX{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to XXXXXXXX. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	xxxxxxxx := make([]XXXXXXXX, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		xxxxxxxx = append(xxxxxxxx, XXXXXXXX{
			XXXXXXXX:                     data.XXXXXXXX,
		})
	}

	return xxxxxxxx, nil
}
