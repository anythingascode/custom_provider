terraform {
  required_providers {
    github = {
      version = "0.1.0"
      source  = "github.com/dev/github"
    }
  }
}

provider "github" {
  username = "summitdevops"
  token = "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"
}

data "github_workflow" "this" {
  owner = "summitdevops"
  repo = "github_workflow"
  workflow_file_name = "dispatch.yaml"
}
 
output "name" {
  value = data.github_workflow.this.workflow_file_name
}

output "url" {
  value = data.github_workflow.this.url
}

output "id" {
  value = data.github_workflow.this.id
}

output "state" {
  value = data.github_workflow.this.state
}