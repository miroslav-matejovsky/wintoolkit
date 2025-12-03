# winservicedetail

Small Go helpers for reading information about Windows services.

Features

- List installed services and check for existence
- Read service state, startup type and human-friendly descriptions
- Inspect service configuration and recovery options
- Locate the service executable, read file version metadata and discover common config files

Requirements

- Go 1.20+ (module mode)
- Windows OS (the package uses Windows-only syscalls)

Installation

Add the module to your project:

```sh
go get github.com/miroslav-matejovsky/winservicedetail
```

Quick examples

Basic check if a service exists and print its state:

```go
mgr := winservicedetail.NewWinSvcManager()
defer mgr.Disconnect()

exists, err := mgr.ServiceExists("MyService")
if err != nil {
  log.Fatalf("failed to check service: %v", err)
}
if !exists {
  log.Printf("service not found")
} else {
  _, stateStr, _ := mgr.GetServiceState("MyService")
  fmt.Println("state:", stateStr)
}
```

Get full details including file metadata and config files (if present):

```go
mgr := winservicedetail.NewWinSvcManager()
defer mgr.Disconnect()

details, err := mgr.GetServiceDetails("MyService", true)
if err != nil {
  log.Fatalf("failed to get details: %v", err)
}
fmt.Printf("Executable: %s\n", details.PathToExecutable)
fmt.Printf("Binary file version: %s\n", details.Executable.ExecutableFile.Version)
for _, cf := range details.Executable.ConfigFiles {
  fmt.Printf("config: %s (size %d)\n", cf.Path, len(cf.Contents))
}
```

Testing

Run unit tests with:

```pwsh
go test ./...
```

Notes

- The package is intended for use on Windows only; calls will fail on other OSes.
- Some functions rely on the current user's permissions to query the service manager.
- The library intentionally avoids modifying services; it only reads information.
