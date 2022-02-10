package main

import (
	"github.com/summitdevops/custom_provider/client"
	"github.com/summitdevops/custom_provider/github-client-go/github"
)

//"git/github"

//"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

func main() {
	/*
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: func() *schema.Provider {
				return github.Provider()
			},
		})
	*/
	client.NewClient("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")
	github.GetWorkflowByName("summitdevops", "github_workflow", "dispatch.yaml")

}
