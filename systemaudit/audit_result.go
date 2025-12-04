package systemaudit

import (
	rta "github.com/miroslav-matejovsky/wintoolkit/runtimesaudit"
	wsd "github.com/miroslav-matejovsky/wintoolkit/winservicedetail"
)

type Result struct {
	Runtimes        rta.AuditResult
	WindowsServices []wsd.ServiceDetails
}
