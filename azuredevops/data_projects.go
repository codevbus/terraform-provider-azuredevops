package azuredevops

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/config"
)

func dataProjects() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProjectsRead,
		Schema: map[string]*schema.Schema{
			"projects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Computed:     true,
							ValidateFunc: validation.NoZeroValues,
						},
						"project_id": {
							Type:         schema.TypeString,
							Computed:     true,
							ValidateFunc: validation.NoZeroValues,
						},
						"project_url": {
							Type:         schema.TypeString,
							Computed:     true,
							ValidateFunc: validation.NoZeroValues,
						},
						"state": {
							Type:         schema.TypeString,
							ValidateFunc: validation.NoZeroValues,
							Default:      "WellFormed",
						},
					},
				},
			},
		},
	}
}

// Performs a lookup of all projects in an organization. This involves the following actions:
//  (1) Query for all Projects that exist within the organization.
//  (2) If a state filter is provided return only matching projects.
func dataSourceProjectsRead(d *schema.ResourceData, m interface{}) error {
	clients := m.(*config.AggregatedClient)
	state := d.Get("state").(string)

	projects, err := getProjectsForState(clients, state)
	if err != nil {
		return fmt.Errorf("Error finding projects with state %s. Error: %v", state, err)
	}
	return nil
}

func getProjectsForState(clients *config.AggregatedClient, projectState string) ([]core.TeamProjectReference, error) {
	var projects []core.TeamProjectReference
	var currentToken string

	for hasMore := true; hasMore; {
		newProjects, latestToken, err := getProjectsWithContinuationToken(clients, projectState, currentToken)
		currentToken = latestToken
		if err != nil {
			return nil, err
		}

		projects = append(projects, newProjects...)
		hasMore = currentToken != ""
	}

	return projects, nil
}

func getProjectsWithContinuationToken(clients *config.AggregatedClient, projectState string, continuationToken string) ([]core.TeamProjectReference, string, error) {
	args := core.GetProjectsArgs{StateFilter: &projectState}
	if continuationToken != "" {
		args.ContinuationToken = &continuationToken
	}

	response, err := clients.CoreClient.GetProjects(clients.Ctx, args)
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
