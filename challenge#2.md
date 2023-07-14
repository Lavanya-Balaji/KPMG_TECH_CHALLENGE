Challenge #2
# Title: We need to write code that will query the meta data of an instance within AWS or Azure or GCP and provide a json formatted output. The choice of language and implementation is up to you.

# Answer: 
## Step 1: Install Azure CLI 
  Download and install azure cli to Windows / Linux Platform  from the link [this](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli) 

## Step 2: Validate Azure CLI Installation
  validate CLI is installed using the cmd 
	` az version`
    

## Step 3: Authenticate to Azure Subscription
  Connect to Azure Account using Azure Cli command
	`az login`

## Step 4: Fetch Details of Azure Resource in json format
    Get the Azure Kubernetes  Cluster Details and output in json format 
	 `az aks show --name <MyManagedCluster> --resource-group <MyResourceGroup> -ojson`