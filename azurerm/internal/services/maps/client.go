package maps

import (
	"github.com/Azure/azure-sdk-for-go/services/maps/mgmt/2018-05-01/maps"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
)

type Client struct {
	AccountsClient maps.AccountsClient
}

func BuildClient(o *common.ClientOptions) *Client {
	c := Client{}

	c.AccountsClient = maps.NewAccountsClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
	o.ConfigureClient(&c.AccountsClient.Client, o.ResourceManagerAuthorizer)

	return &c
}
