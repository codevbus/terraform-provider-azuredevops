package azuredevops

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/microsoft/azure-devops-go-api/azuredevops/core"
	"github.com/microsoft/azure-devops-go-api/azuredevops/git"
	"github.com/microsoft/terraform-provider-azuredevops/azuredevops/utils/converter"
)

func resourceAzureGitRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceAzureGitRepositoryCreate,
		Read:   resourceAzureGitRepositoryRead,
		Update: resourceAzureGitRepositoryUpdate,
		Delete: resourceAzureGitRepositoryDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				ForceNew: false,
				Required: true,
			},
			"default_branch": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_fork": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"remote_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ssh_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"web_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAzureGitRepositoryCreate(d *schema.ResourceData, m interface{}) error {
	clients := m.(*aggregatedClient)
	repo, projectID, err := expandAzureGitRepository(d)

	createdRepo, err := createAzureGitRepository(clients, repo.Name, projectID)
	if err != nil {
		return fmt.Errorf("Error creating project in Azure DevOps: %+v", err)
	}

	flattenAzureGitRepository(d, createdRepo)

	return resourceAzureGitRepositoryRead(d, m)
}

func createAzureGitRepository(clients *aggregatedClient, repoName *string, projectID *uuid.UUID) (*git.GitRepository, error) {
	args := git.CreateRepositoryArgs{
		GitRepositoryToCreate: &git.GitRepositoryCreateOptions{
			Name: repoName,
			Project: &core.TeamProjectReference{
				Id: projectID,
			},
		},
	}
	createdRepository, err := clients.GitReposClient.CreateRepository(clients.ctx, args)

	return createdRepository, err
}

func resourceAzureGitRepositoryRead(d *schema.ResourceData, m interface{}) error {
	repoID := d.Id()
	repoName := d.Get("name").(string)
	projectID := d.Get("project_id").(string)

	clients := m.(*aggregatedClient)
	repo, err := azureGitRepositoryRead(clients, repoID, repoName, projectID)
	if err != nil {
		return fmt.Errorf("Error looking up repository with ID %s and Name %s. Error: %v", repoID, repoName, err)
	}

	flattenAzureGitRepository(d, repo)
	return nil
}

func resourceAzureGitRepositoryUpdate(d *schema.ResourceData, m interface{}) error {
	clients := m.(*aggregatedClient)
	repo, projectID, err := expandAzureGitRepository(d)
	if err != nil {
		return fmt.Errorf("Error converting terraform data model to AzDO project reference: %+v", err)
	}

	repo, err = updateAzureGitRepository(clients, repo, projectID)
	if err != nil {
		return fmt.Errorf("Error updating repository in Azure DevOps: %+v", err)
	}

	flattenAzureGitRepository(d, repo)
	return resourceAzureGitRepositoryRead(d, m)
}

func updateAzureGitRepository(clients *aggregatedClient, repository *git.GitRepository, project *uuid.UUID) (*git.GitRepository, error) {
	projectID := project.String()
	return clients.GitReposClient.UpdateRepository(
		clients.ctx,
		git.UpdateRepositoryArgs{
			NewRepositoryInfo: repository,
			RepositoryId:      repository.Id,
			Project:           &projectID,
		})
}

func resourceAzureGitRepositoryDelete(d *schema.ResourceData, m interface{}) error {
	repoID := d.Id()
	clients := m.(*aggregatedClient)
	return deleteAzureGitRepository(clients, repoID)
}

func deleteAzureGitRepository(clients *aggregatedClient, repoID string) error {
	uuid, err := uuid.Parse(repoID)
	if err != nil {
		return fmt.Errorf("Invalid repositoryId UUID: %s", repoID)
	}

	return clients.GitReposClient.DeleteRepository(clients.ctx, git.DeleteRepositoryArgs{
		RepositoryId: &uuid,
	})
}

// Lookup an Azure Git Repository using the ID, or name if the ID is not set.
func azureGitRepositoryRead(clients *aggregatedClient, repoID string, repoName string, projectID string) (*git.GitRepository, error) {
	identifier := repoID
	if identifier == "" {
		identifier = repoName
	}

	return clients.GitReposClient.GetRepository(clients.ctx, git.GetRepositoryArgs{
		RepositoryId: converter.String(identifier),
		Project:      converter.String(projectID),
	})
}

func flattenAzureGitRepository(d *schema.ResourceData, repository *git.GitRepository) {
	d.SetId(repository.Id.String())

	d.Set("name", converter.ToString(repository.Name, ""))
	d.Set("project_id", repository.Project.Id.String())
	d.Set("default_branch", converter.ToString(repository.DefaultBranch, ""))
	d.Set("is_fork", repository.IsFork)
	d.Set("remote_url", converter.ToString(repository.RemoteUrl, ""))
	d.Set("size", repository.Size)
	d.Set("ssh_url", converter.ToString(repository.SshUrl, ""))
	d.Set("url", converter.ToString(repository.Url, ""))
	d.Set("web_url", converter.ToString(repository.WebUrl, ""))
}

// Convert internal Terraform data structure to an AzDO data structure. Note: only the params that are
// not generated by the service are expanded here
func expandAzureGitRepository(d *schema.ResourceData) (*git.GitRepository, *uuid.UUID, error) {
	// an "error" is OK here as it is expected in the case that the ID is not set in the resource data
	var repoID *uuid.UUID
	parsedID, err := uuid.Parse(d.Id())
	if err == nil {
		repoID = &parsedID
	}

	projectID, err := uuid.Parse(d.Get("project_id").(string))
	if err != nil {
		return nil, nil, err
	}

	repo := &git.GitRepository{
		Id:   repoID,
		Name: converter.String(d.Get("name").(string)),
	}

	return repo, &projectID, nil
}
