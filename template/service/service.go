package service

import (
	"context"
	"errors"

	"github.com/zenoss/zenkit"
	// proto "github.com/zenoss/zing-proto/go/{{Name}}"
)

/* Uncomment this stuff once you have a service protobuf to link to

// New{{Name | title}}Service creates a new impl of the service protobuf
func New{{Name | title}}Service() proto.{{Name | title}}ServiceServer {
	return &{{Name | title}}Service{}
}

// {{Name}}Service is a shortlink service impl
type {{Name | title}}Service struct {}

// Now implement the interface
func (svc *{{Name | title}}Service) DoSomething(ctx context.Context) error {
    return nil
}

*/
