package systemaudit

import (
	"fmt"

	rta "github.com/miroslav-matejovsky/wintoolkit/runtimesaudit"
	wsd "github.com/miroslav-matejovsky/wintoolkit/winservicedetail"
)

func DoSystemAudit() (*Result, error) {

	runtimesResult, err := rta.DoAudit()
	if err != nil {
		return nil, fmt.Errorf("failed to audit runtimes: %w", err)
	}

	wsm := wsd.NewWinSvcManager()
	allServices, err := wsm.ListServices()
	if err != nil {
		return nil, fmt.Errorf("failed to list Windows services: %w", err)
	}
	includeConfigFiles := true
	var detailedServices []wsd.ServiceDetails
	for _, svcName := range allServices {
		svcDetail, err := wsm.GetServiceDetails(svcName, includeConfigFiles)
		if err != nil {
			return nil, fmt.Errorf("failed to get details for service %s: %w", svcName, err)
		}
		detailedServices = append(detailedServices, *svcDetail)
	}
	return &Result{
		Runtimes:        *runtimesResult,
		WindowsServices: detailedServices,
	}, nil
}
