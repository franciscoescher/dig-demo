package clients

import (
	"dig/clients/cloudprovider"
	"dig/clients/storage"
	"dig/internal/helper"
)

var Module = helper.Module(
	cloudprovider.New,
	storage.New,
)
