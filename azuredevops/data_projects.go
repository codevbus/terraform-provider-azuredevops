package azuredevops

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
)

func dataProjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.NoZeroValues,
			},
			"state": {
				Type:         schema.TypeString,
				ForceNew:     true,
				Optional:     true,
				ValidateFunc: validation.NoZeroValues,
			},
		},
	}
}

// Performs a lookup of all projects in an organization. This involves the following actions:
//  (1) Query for all Projects that exist within the organization.
//  (2) If a state filter is provided return only matching projects.
func dataSourceProjectsRead(d *schema.ResourceData, m interface{}) error {
	clients = m.(*aggregatedClient)
	return nil
}

func getProjectsWithContinuationToken(clients *aggregatedClient, continuationToken string) ([]core.TeamProjectReference, string, error) {
	args := core.GetProjectsArgs{}
	if continuationToken != "" {
		args.ContinuationToken = &continuationToken
	}

	response, err := clients.CoreClient.GetProjects(clients.ctx, args)
	if err != nil {
		return nil, "", err
	}

	if response.ContinuationToken != "" && len(response.ContinuationToken) > 1 {
		return nil, "", fmt.Errorf("Expected at most 1 continuation token, but found %d", len(response.ContinuationToken))
	}

	var newToken string
	if response.ContinuationToken != "" && len(response.ContinuationToken) > 0 {
		newToken = (response.ContinuationToken)
	}

	return response.Value, newToken, nil
}
