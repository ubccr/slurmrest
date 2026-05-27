# \SlurmdbAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmdbV0044DeleteAccount**](SlurmdbAPI.md#SlurmdbV0044DeleteAccount) | **Delete** /slurmdb/v0.0.44/account/{account_name} | Delete account
[**SlurmdbV0044DeleteAssociation**](SlurmdbAPI.md#SlurmdbV0044DeleteAssociation) | **Delete** /slurmdb/v0.0.44/association/ | Delete association
[**SlurmdbV0044DeleteAssociations**](SlurmdbAPI.md#SlurmdbV0044DeleteAssociations) | **Delete** /slurmdb/v0.0.44/associations/ | Delete associations
[**SlurmdbV0044DeleteCluster**](SlurmdbAPI.md#SlurmdbV0044DeleteCluster) | **Delete** /slurmdb/v0.0.44/cluster/{cluster_name} | Delete cluster
[**SlurmdbV0044DeleteSingleQos**](SlurmdbAPI.md#SlurmdbV0044DeleteSingleQos) | **Delete** /slurmdb/v0.0.44/qos/{qos} | Delete QOS
[**SlurmdbV0044DeleteUser**](SlurmdbAPI.md#SlurmdbV0044DeleteUser) | **Delete** /slurmdb/v0.0.44/user/{name} | Delete user
[**SlurmdbV0044DeleteWckey**](SlurmdbAPI.md#SlurmdbV0044DeleteWckey) | **Delete** /slurmdb/v0.0.44/wckey/{id} | Delete wckey
[**SlurmdbV0044GetAccount**](SlurmdbAPI.md#SlurmdbV0044GetAccount) | **Get** /slurmdb/v0.0.44/account/{account_name} | Get account info
[**SlurmdbV0044GetAccounts**](SlurmdbAPI.md#SlurmdbV0044GetAccounts) | **Get** /slurmdb/v0.0.44/accounts/ | Get account list
[**SlurmdbV0044GetAssociation**](SlurmdbAPI.md#SlurmdbV0044GetAssociation) | **Get** /slurmdb/v0.0.44/association/ | Get association info
[**SlurmdbV0044GetAssociations**](SlurmdbAPI.md#SlurmdbV0044GetAssociations) | **Get** /slurmdb/v0.0.44/associations/ | Get association list
[**SlurmdbV0044GetCluster**](SlurmdbAPI.md#SlurmdbV0044GetCluster) | **Get** /slurmdb/v0.0.44/cluster/{cluster_name} | Get cluster info
[**SlurmdbV0044GetClusters**](SlurmdbAPI.md#SlurmdbV0044GetClusters) | **Get** /slurmdb/v0.0.44/clusters/ | Get cluster list
[**SlurmdbV0044GetConfig**](SlurmdbAPI.md#SlurmdbV0044GetConfig) | **Get** /slurmdb/v0.0.44/config | Dump all configuration information
[**SlurmdbV0044GetDiag**](SlurmdbAPI.md#SlurmdbV0044GetDiag) | **Get** /slurmdb/v0.0.44/diag/ | Get slurmdb diagnostics
[**SlurmdbV0044GetInstance**](SlurmdbAPI.md#SlurmdbV0044GetInstance) | **Get** /slurmdb/v0.0.44/instance/ | Get instance info
[**SlurmdbV0044GetInstances**](SlurmdbAPI.md#SlurmdbV0044GetInstances) | **Get** /slurmdb/v0.0.44/instances/ | Get instance list
[**SlurmdbV0044GetJob**](SlurmdbAPI.md#SlurmdbV0044GetJob) | **Get** /slurmdb/v0.0.44/job/{job_id} | Get job info
[**SlurmdbV0044GetJobs**](SlurmdbAPI.md#SlurmdbV0044GetJobs) | **Get** /slurmdb/v0.0.44/jobs/ | Get job list
[**SlurmdbV0044GetPing**](SlurmdbAPI.md#SlurmdbV0044GetPing) | **Get** /slurmdb/v0.0.44/ping/ | ping test
[**SlurmdbV0044GetQos**](SlurmdbAPI.md#SlurmdbV0044GetQos) | **Get** /slurmdb/v0.0.44/qos/ | Get QOS list
[**SlurmdbV0044GetSingleQos**](SlurmdbAPI.md#SlurmdbV0044GetSingleQos) | **Get** /slurmdb/v0.0.44/qos/{qos} | Get QOS info
[**SlurmdbV0044GetTres**](SlurmdbAPI.md#SlurmdbV0044GetTres) | **Get** /slurmdb/v0.0.44/tres/ | Get TRES info
[**SlurmdbV0044GetUser**](SlurmdbAPI.md#SlurmdbV0044GetUser) | **Get** /slurmdb/v0.0.44/user/{name} | Get user info
[**SlurmdbV0044GetUsers**](SlurmdbAPI.md#SlurmdbV0044GetUsers) | **Get** /slurmdb/v0.0.44/users/ | Get user list
[**SlurmdbV0044GetWckey**](SlurmdbAPI.md#SlurmdbV0044GetWckey) | **Get** /slurmdb/v0.0.44/wckey/{id} | Get wckey info
[**SlurmdbV0044GetWckeys**](SlurmdbAPI.md#SlurmdbV0044GetWckeys) | **Get** /slurmdb/v0.0.44/wckeys/ | Get wckey list
[**SlurmdbV0044PostAccounts**](SlurmdbAPI.md#SlurmdbV0044PostAccounts) | **Post** /slurmdb/v0.0.44/accounts/ | Add/update list of accounts
[**SlurmdbV0044PostAccountsAssociation**](SlurmdbAPI.md#SlurmdbV0044PostAccountsAssociation) | **Post** /slurmdb/v0.0.44/accounts_association/ | Add accounts with conditional association
[**SlurmdbV0044PostAssociations**](SlurmdbAPI.md#SlurmdbV0044PostAssociations) | **Post** /slurmdb/v0.0.44/associations/ | Set associations info
[**SlurmdbV0044PostClusters**](SlurmdbAPI.md#SlurmdbV0044PostClusters) | **Post** /slurmdb/v0.0.44/clusters/ | Get cluster list
[**SlurmdbV0044PostConfig**](SlurmdbAPI.md#SlurmdbV0044PostConfig) | **Post** /slurmdb/v0.0.44/config | Load all configuration information
[**SlurmdbV0044PostJob**](SlurmdbAPI.md#SlurmdbV0044PostJob) | **Post** /slurmdb/v0.0.44/job/{job_id} | Update job
[**SlurmdbV0044PostJobs**](SlurmdbAPI.md#SlurmdbV0044PostJobs) | **Post** /slurmdb/v0.0.44/jobs/ | Update jobs
[**SlurmdbV0044PostQos**](SlurmdbAPI.md#SlurmdbV0044PostQos) | **Post** /slurmdb/v0.0.44/qos/ | Add or update QOSs
[**SlurmdbV0044PostTres**](SlurmdbAPI.md#SlurmdbV0044PostTres) | **Post** /slurmdb/v0.0.44/tres/ | Add TRES
[**SlurmdbV0044PostUsers**](SlurmdbAPI.md#SlurmdbV0044PostUsers) | **Post** /slurmdb/v0.0.44/users/ | Update users
[**SlurmdbV0044PostUsersAssociation**](SlurmdbAPI.md#SlurmdbV0044PostUsersAssociation) | **Post** /slurmdb/v0.0.44/users_association/ | Add users with conditional association
[**SlurmdbV0044PostWckeys**](SlurmdbAPI.md#SlurmdbV0044PostWckeys) | **Post** /slurmdb/v0.0.44/wckeys/ | Add or update wckeys



## SlurmdbV0044DeleteAccount

> V0044OpenapiAccountsRemovedResp SlurmdbV0044DeleteAccount(ctx, accountName).Execute()

Delete account

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteAccount(context.Background(), accountName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteAccount`: V0044OpenapiAccountsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiAccountsRemovedResp**](V0044OpenapiAccountsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteAssociation

> V0044OpenapiAssocsRemovedResp SlurmdbV0044DeleteAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteAssociation`: V0044OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsRemovedResp**](V0044OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteAssociations

> V0044OpenapiAssocsRemovedResp SlurmdbV0044DeleteAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Delete associations

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteAssociations`: V0044OpenapiAssocsRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsRemovedResp**](V0044OpenapiAssocsRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteCluster

> V0044OpenapiClustersRemovedResp SlurmdbV0044DeleteCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Delete cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteCluster`: V0044OpenapiClustersRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0044OpenapiClustersRemovedResp**](V0044OpenapiClustersRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteSingleQos

> V0044OpenapiSlurmdbdQosRemovedResp SlurmdbV0044DeleteSingleQos(ctx, qos).Execute()

Delete QOS

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteSingleQos(context.Background(), qos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteSingleQos`: V0044OpenapiSlurmdbdQosRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiSlurmdbdQosRemovedResp**](V0044OpenapiSlurmdbdQosRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteUser

> V0044OpenapiResp SlurmdbV0044DeleteUser(ctx, name).Execute()

Delete user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteUser(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteUser`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044DeleteWckey

> V0044OpenapiWckeyRemovedResp SlurmdbV0044DeleteWckey(ctx, id).Execute()

Delete wckey

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044DeleteWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044DeleteWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044DeleteWckey`: V0044OpenapiWckeyRemovedResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044DeleteWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044DeleteWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiWckeyRemovedResp**](V0044OpenapiWckeyRemovedResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAccount

> V0044OpenapiAccountsResp SlurmdbV0044GetAccount(ctx, accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()

Get account info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	accountName := "accountName_example" // string | Account name
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withDeleted := "withDeleted_example" // string | Include deleted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAccount(context.Background(), accountName).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAccount`: V0044OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountName** | **string** | Account name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withDeleted** | **string** | Include deleted | 

### Return type

[**V0044OpenapiAccountsResp**](V0044OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAccounts

> V0044OpenapiAccountsResp SlurmdbV0044GetAccounts(ctx).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()

Get account list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	dELETED := "dELETED_example" // string | include deleted associations (optional)
	withAssociations := "withAssociations_example" // string | query includes associations (optional)
	withCoordinators := "withCoordinators_example" // string | query includes coordinators (optional)
	noUsersAreCoords := "noUsersAreCoords_example" // string | remove users as coordinators (optional)
	usersAreCoords := "usersAreCoords_example" // string | users are coordinators (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAccounts(context.Background()).Description(description).DELETED(dELETED).WithAssociations(withAssociations).WithCoordinators(withCoordinators).NoUsersAreCoords(noUsersAreCoords).UsersAreCoords(usersAreCoords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAccounts`: V0044OpenapiAccountsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **dELETED** | **string** | include deleted associations | 
 **withAssociations** | **string** | query includes associations | 
 **withCoordinators** | **string** | query includes coordinators | 
 **noUsersAreCoords** | **string** | remove users as coordinators | 
 **usersAreCoords** | **string** | users are coordinators | 

### Return type

[**V0044OpenapiAccountsResp**](V0044OpenapiAccountsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAssociation

> V0044OpenapiAssocsResp SlurmdbV0044GetAssociation(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAssociation(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAssociation`: V0044OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsResp**](V0044OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetAssociations

> V0044OpenapiAssocsResp SlurmdbV0044GetAssociations(ctx).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()

Get association list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV accounts list (optional)
	cluster := "cluster_example" // string | CSV clusters list (optional)
	defaultQos := "defaultQos_example" // string | CSV QOS list (optional)
	includeDeletedAssociations := "includeDeletedAssociations_example" // string |  (optional)
	includeUsage := "includeUsage_example" // string |  (optional)
	filterToOnlyDefaults := "filterToOnlyDefaults_example" // string |  (optional)
	includeTheRawQOSOrDeltaQos := "includeTheRawQOSOrDeltaQos_example" // string |  (optional)
	includeSubAcctInformation := "includeSubAcctInformation_example" // string |  (optional)
	excludeParentIdName := "excludeParentIdName_example" // string |  (optional)
	excludeLimitsFromParents := "excludeLimitsFromParents_example" // string |  (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	parentAccount := "parentAccount_example" // string | CSV names of parent account (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetAssociations(context.Background()).Account(account).Cluster(cluster).DefaultQos(defaultQos).IncludeDeletedAssociations(includeDeletedAssociations).IncludeUsage(includeUsage).FilterToOnlyDefaults(filterToOnlyDefaults).IncludeTheRawQOSOrDeltaQos(includeTheRawQOSOrDeltaQos).IncludeSubAcctInformation(includeSubAcctInformation).ExcludeParentIdName(excludeParentIdName).ExcludeLimitsFromParents(excludeLimitsFromParents).Format(format).Id(id).ParentAccount(parentAccount).Partition(partition).Qos(qos).UsageEnd(usageEnd).UsageStart(usageStart).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetAssociations`: V0044OpenapiAssocsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV accounts list | 
 **cluster** | **string** | CSV clusters list | 
 **defaultQos** | **string** | CSV QOS list | 
 **includeDeletedAssociations** | **string** |  | 
 **includeUsage** | **string** |  | 
 **filterToOnlyDefaults** | **string** |  | 
 **includeTheRawQOSOrDeltaQos** | **string** |  | 
 **includeSubAcctInformation** | **string** |  | 
 **excludeParentIdName** | **string** |  | 
 **excludeLimitsFromParents** | **string** |  | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **parentAccount** | **string** | CSV names of parent account | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 

### Return type

[**V0044OpenapiAssocsResp**](V0044OpenapiAssocsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetCluster

> V0044OpenapiClustersResp SlurmdbV0044GetCluster(ctx, clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()

Get cluster info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	clusterName := "clusterName_example" // string | Cluster name
	classification := "classification_example" // string | Type of machine (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	federation := "federation_example" // string | CSV federation list (optional)
	flags := "flags_example" // string | Query flags (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	rpcVersion := "rpcVersion_example" // string | CSV RPC version list (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	withDeleted := "withDeleted_example" // string | Include deleted clusters (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetCluster(context.Background(), clusterName).Classification(classification).Cluster(cluster).Federation(federation).Flags(flags).Format(format).RpcVersion(rpcVersion).UsageEnd(usageEnd).UsageStart(usageStart).WithDeleted(withDeleted).WithUsage(withUsage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetCluster`: V0044OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterName** | **string** | Cluster name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classification** | **string** | Type of machine | 
 **cluster** | **string** | CSV cluster list | 
 **federation** | **string** | CSV federation list | 
 **flags** | **string** | Query flags | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **rpcVersion** | **string** | CSV RPC version list | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **withDeleted** | **string** | Include deleted clusters | 
 **withUsage** | **string** | Include usage | 

### Return type

[**V0044OpenapiClustersResp**](V0044OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetClusters

> V0044OpenapiClustersResp SlurmdbV0044GetClusters(ctx).UpdateTime(updateTime).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetClusters(context.Background()).UpdateTime(updateTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetClusters`: V0044OpenapiClustersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 

### Return type

[**V0044OpenapiClustersResp**](V0044OpenapiClustersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetConfig

> V0044OpenapiSlurmdbdConfigResp SlurmdbV0044GetConfig(ctx).Execute()

Dump all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetConfig(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetConfig`: V0044OpenapiSlurmdbdConfigResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetConfigRequest struct via the builder pattern


### Return type

[**V0044OpenapiSlurmdbdConfigResp**](V0044OpenapiSlurmdbdConfigResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetDiag

> V0044OpenapiSlurmdbdStatsResp SlurmdbV0044GetDiag(ctx).Execute()

Get slurmdb diagnostics

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetDiag(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetDiag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetDiag`: V0044OpenapiSlurmdbdStatsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetDiagRequest struct via the builder pattern


### Return type

[**V0044OpenapiSlurmdbdStatsResp**](V0044OpenapiSlurmdbdStatsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetInstance

> V0044OpenapiInstancesResp SlurmdbV0044GetInstance(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetInstance(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetInstance`: V0044OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0044OpenapiInstancesResp**](V0044OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetInstances

> V0044OpenapiInstancesResp SlurmdbV0044GetInstances(ctx).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()

Get instance list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV clusters list (optional)
	extra := "extra_example" // string | CSV extra list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	instanceId := "instanceId_example" // string | CSV instance_id list (optional)
	instanceType := "instanceType_example" // string | CSV instance_type list (optional)
	nodeList := "nodeList_example" // string | Ranged node string (optional)
	timeEnd := "timeEnd_example" // string | Time end (UNIX timestamp) (optional)
	timeStart := "timeStart_example" // string | Time start (UNIX timestamp) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetInstances(context.Background()).Cluster(cluster).Extra(extra).Format(format).InstanceId(instanceId).InstanceType(instanceType).NodeList(nodeList).TimeEnd(timeEnd).TimeStart(timeStart).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetInstances`: V0044OpenapiInstancesResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV clusters list | 
 **extra** | **string** | CSV extra list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **instanceId** | **string** | CSV instance_id list | 
 **instanceType** | **string** | CSV instance_type list | 
 **nodeList** | **string** | Ranged node string | 
 **timeEnd** | **string** | Time end (UNIX timestamp) | 
 **timeStart** | **string** | Time start (UNIX timestamp) | 

### Return type

[**V0044OpenapiInstancesResp**](V0044OpenapiInstancesResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetJob

> V0044OpenapiSlurmdbdJobsResp SlurmdbV0044GetJob(ctx, jobId).Execute()

Get job info



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetJob(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetJob`: V0044OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiSlurmdbdJobsResp**](V0044OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetJobs

> V0044OpenapiSlurmdbdJobsResp SlurmdbV0044GetJobs(ctx).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).JobAltered(jobAltered).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()

Get job list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	account := "account_example" // string | CSV account list (optional)
	association := "association_example" // string | CSV association list (optional)
	cluster := "cluster_example" // string | CSV cluster list (optional)
	constraints := "constraints_example" // string | CSV constraint list (optional)
	schedulerUnset := "schedulerUnset_example" // string | Schedule bits not set (optional)
	scheduledOnSubmit := "scheduledOnSubmit_example" // string | Job was started on submit (optional)
	scheduledByMain := "scheduledByMain_example" // string | Job was started from main scheduler (optional)
	scheduledByBackfill := "scheduledByBackfill_example" // string | Job was started from backfill (optional)
	jobStarted := "jobStarted_example" // string | Job start RPC was received (optional)
	jobAltered := "jobAltered_example" // string | Job record has been altered (optional)
	exitCode := "exitCode_example" // string | Job exit code (numeric) (optional)
	showDuplicates := "showDuplicates_example" // string | Include duplicate job entries (optional)
	skipSteps := "skipSteps_example" // string | Exclude job step details (optional)
	disableTruncateUsageTime := "disableTruncateUsageTime_example" // string | Do not truncate the time to usage_start and usage_end (optional)
	wholeHetjob := "wholeHetjob_example" // string | Include details on all hetjob components (optional)
	disableWholeHetjob := "disableWholeHetjob_example" // string | Only show details on specified hetjob components (optional)
	disableWaitForResult := "disableWaitForResult_example" // string | Tell dbd not to wait for the result (optional)
	usageTimeAsSubmitTime := "usageTimeAsSubmitTime_example" // string | Use usage_time as the submit_time of the job (optional)
	showBatchScript := "showBatchScript_example" // string | Include job script (optional)
	showJobEnvironment := "showJobEnvironment_example" // string | Include job environment (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	groups := "groups_example" // string | CSV group list (optional)
	jobName := "jobName_example" // string | CSV job name list (optional)
	partition := "partition_example" // string | CSV partition name list (optional)
	qos := "qos_example" // string | CSV QOS name list (optional)
	reason := "reason_example" // string | CSV reason list (optional)
	reservation := "reservation_example" // string | CSV reservation name list (optional)
	reservationId := "reservationId_example" // string | CSV reservation ID list (optional)
	state := "state_example" // string | CSV state list (optional)
	step := "step_example" // string | CSV step id list (optional)
	endTime := "endTime_example" // string | Usage end (UNIX timestamp) (optional)
	startTime := "startTime_example" // string | Usage start (UNIX timestamp) (optional)
	node := "node_example" // string | Ranged node string where jobs ran (optional)
	users := "users_example" // string | CSV user name list (optional)
	wckey := "wckey_example" // string | CSV WCKey list (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetJobs(context.Background()).Account(account).Association(association).Cluster(cluster).Constraints(constraints).SchedulerUnset(schedulerUnset).ScheduledOnSubmit(scheduledOnSubmit).ScheduledByMain(scheduledByMain).ScheduledByBackfill(scheduledByBackfill).JobStarted(jobStarted).JobAltered(jobAltered).ExitCode(exitCode).ShowDuplicates(showDuplicates).SkipSteps(skipSteps).DisableTruncateUsageTime(disableTruncateUsageTime).WholeHetjob(wholeHetjob).DisableWholeHetjob(disableWholeHetjob).DisableWaitForResult(disableWaitForResult).UsageTimeAsSubmitTime(usageTimeAsSubmitTime).ShowBatchScript(showBatchScript).ShowJobEnvironment(showJobEnvironment).Format(format).Groups(groups).JobName(jobName).Partition(partition).Qos(qos).Reason(reason).Reservation(reservation).ReservationId(reservationId).State(state).Step(step).EndTime(endTime).StartTime(startTime).Node(node).Users(users).Wckey(wckey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetJobs`: V0044OpenapiSlurmdbdJobsResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **account** | **string** | CSV account list | 
 **association** | **string** | CSV association list | 
 **cluster** | **string** | CSV cluster list | 
 **constraints** | **string** | CSV constraint list | 
 **schedulerUnset** | **string** | Schedule bits not set | 
 **scheduledOnSubmit** | **string** | Job was started on submit | 
 **scheduledByMain** | **string** | Job was started from main scheduler | 
 **scheduledByBackfill** | **string** | Job was started from backfill | 
 **jobStarted** | **string** | Job start RPC was received | 
 **jobAltered** | **string** | Job record has been altered | 
 **exitCode** | **string** | Job exit code (numeric) | 
 **showDuplicates** | **string** | Include duplicate job entries | 
 **skipSteps** | **string** | Exclude job step details | 
 **disableTruncateUsageTime** | **string** | Do not truncate the time to usage_start and usage_end | 
 **wholeHetjob** | **string** | Include details on all hetjob components | 
 **disableWholeHetjob** | **string** | Only show details on specified hetjob components | 
 **disableWaitForResult** | **string** | Tell dbd not to wait for the result | 
 **usageTimeAsSubmitTime** | **string** | Use usage_time as the submit_time of the job | 
 **showBatchScript** | **string** | Include job script | 
 **showJobEnvironment** | **string** | Include job environment | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **groups** | **string** | CSV group list | 
 **jobName** | **string** | CSV job name list | 
 **partition** | **string** | CSV partition name list | 
 **qos** | **string** | CSV QOS name list | 
 **reason** | **string** | CSV reason list | 
 **reservation** | **string** | CSV reservation name list | 
 **reservationId** | **string** | CSV reservation ID list | 
 **state** | **string** | CSV state list | 
 **step** | **string** | CSV step id list | 
 **endTime** | **string** | Usage end (UNIX timestamp) | 
 **startTime** | **string** | Usage start (UNIX timestamp) | 
 **node** | **string** | Ranged node string where jobs ran | 
 **users** | **string** | CSV user name list | 
 **wckey** | **string** | CSV WCKey list | 

### Return type

[**V0044OpenapiSlurmdbdJobsResp**](V0044OpenapiSlurmdbdJobsResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetPing

> V0044OpenapiSlurmdbdPingResp SlurmdbV0044GetPing(ctx).Execute()

ping test

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetPing(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetPing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetPing`: V0044OpenapiSlurmdbdPingResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetPingRequest struct via the builder pattern


### Return type

[**V0044OpenapiSlurmdbdPingResp**](V0044OpenapiSlurmdbdPingResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetQos

> V0044OpenapiSlurmdbdQosResp SlurmdbV0044GetQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()

Get QOS list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetQos`: V0044OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 

### Return type

[**V0044OpenapiSlurmdbdQosResp**](V0044OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetSingleQos

> V0044OpenapiSlurmdbdQosResp SlurmdbV0044GetSingleQos(ctx, qos).WithDeleted(withDeleted).Execute()

Get QOS info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	qos := "qos_example" // string | QOS name
	withDeleted := "withDeleted_example" // string | Query includes deleted QOS (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetSingleQos(context.Background(), qos).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetSingleQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetSingleQos`: V0044OpenapiSlurmdbdQosResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetSingleQos`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**qos** | **string** | QOS name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetSingleQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Query includes deleted QOS | 

### Return type

[**V0044OpenapiSlurmdbdQosResp**](V0044OpenapiSlurmdbdQosResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetTres

> V0044OpenapiTresResp SlurmdbV0044GetTres(ctx).Execute()

Get TRES info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetTres(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetTres`: V0044OpenapiTresResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetTres`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetTresRequest struct via the builder pattern


### Return type

[**V0044OpenapiTresResp**](V0044OpenapiTresResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetUser

> V0044OpenapiUsersResp SlurmdbV0044GetUser(ctx, name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()

Get user info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | User name
	withDeleted := "withDeleted_example" // string | Include deleted users (optional)
	withAssocs := "withAssocs_example" // string | Include associations (optional)
	withCoords := "withCoords_example" // string | Include coordinators (optional)
	withWckeys := "withWckeys_example" // string | Include WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetUser(context.Background(), name).WithDeleted(withDeleted).WithAssocs(withAssocs).WithCoords(withCoords).WithWckeys(withWckeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetUser`: V0044OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | User name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **withDeleted** | **string** | Include deleted users | 
 **withAssocs** | **string** | Include associations | 
 **withCoords** | **string** | Include coordinators | 
 **withWckeys** | **string** | Include WCKeys | 

### Return type

[**V0044OpenapiUsersResp**](V0044OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetUsers

> V0044OpenapiUsersResp SlurmdbV0044GetUsers(ctx).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()

Get user list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	adminLevel := "adminLevel_example" // string | Administrator level (optional)
	defaultAccount := "defaultAccount_example" // string | CSV default account list (optional)
	defaultWckey := "defaultWckey_example" // string | CSV default WCKey list (optional)
	withAssocs := "withAssocs_example" // string | With associations (optional)
	withCoords := "withCoords_example" // string | With coordinators (optional)
	withDeleted := "withDeleted_example" // string | With deleted (optional)
	withWckeys := "withWckeys_example" // string | With WCKeys (optional)
	withoutDefaults := "withoutDefaults_example" // string | Exclude defaults (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetUsers(context.Background()).AdminLevel(adminLevel).DefaultAccount(defaultAccount).DefaultWckey(defaultWckey).WithAssocs(withAssocs).WithCoords(withCoords).WithDeleted(withDeleted).WithWckeys(withWckeys).WithoutDefaults(withoutDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetUsers`: V0044OpenapiUsersResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **adminLevel** | **string** | Administrator level | 
 **defaultAccount** | **string** | CSV default account list | 
 **defaultWckey** | **string** | CSV default WCKey list | 
 **withAssocs** | **string** | With associations | 
 **withCoords** | **string** | With coordinators | 
 **withDeleted** | **string** | With deleted | 
 **withWckeys** | **string** | With WCKeys | 
 **withoutDefaults** | **string** | Exclude defaults | 

### Return type

[**V0044OpenapiUsersResp**](V0044OpenapiUsersResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetWckey

> V0044OpenapiWckeyResp SlurmdbV0044GetWckey(ctx, id).Execute()

Get wckey info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | WCKey ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetWckey(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetWckey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetWckey`: V0044OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetWckey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | WCKey ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetWckeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0044OpenapiWckeyResp**](V0044OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044GetWckeys

> V0044OpenapiWckeyResp SlurmdbV0044GetWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()

Get wckey list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044GetWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044GetWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044GetWckeys`: V0044OpenapiWckeyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044GetWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044GetWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 

### Return type

[**V0044OpenapiWckeyResp**](V0044OpenapiWckeyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostAccounts

> V0044OpenapiResp SlurmdbV0044PostAccounts(ctx).V0044OpenapiAccountsResp(v0044OpenapiAccountsResp).Execute()

Add/update list of accounts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiAccountsResp := *openapiclient.NewV0044OpenapiAccountsResp([]openapiclient.V0044Account{*openapiclient.NewV0044Account("Description_example", "Name_example", "Organization_example")}) // V0044OpenapiAccountsResp | Description of accounts to update/create (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostAccounts(context.Background()).V0044OpenapiAccountsResp(v0044OpenapiAccountsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostAccounts`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiAccountsResp** | [**V0044OpenapiAccountsResp**](V0044OpenapiAccountsResp.md) | Description of accounts to update/create | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostAccountsAssociation

> V0044OpenapiAccountsAddCondRespStr SlurmdbV0044PostAccountsAssociation(ctx).V0044OpenapiAccountsAddCondResp(v0044OpenapiAccountsAddCondResp).Execute()

Add accounts with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiAccountsAddCondResp := *openapiclient.NewV0044OpenapiAccountsAddCondResp(*openapiclient.NewV0044AccountsAddCond([]string{"Accounts_example"})) // V0044OpenapiAccountsAddCondResp | Add list of accounts with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostAccountsAssociation(context.Background()).V0044OpenapiAccountsAddCondResp(v0044OpenapiAccountsAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostAccountsAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostAccountsAssociation`: V0044OpenapiAccountsAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostAccountsAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostAccountsAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiAccountsAddCondResp** | [**V0044OpenapiAccountsAddCondResp**](V0044OpenapiAccountsAddCondResp.md) | Add list of accounts with conditional association | 

### Return type

[**V0044OpenapiAccountsAddCondRespStr**](V0044OpenapiAccountsAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostAssociations

> V0044OpenapiResp SlurmdbV0044PostAssociations(ctx).V0044OpenapiAssocsResp(v0044OpenapiAssocsResp).Execute()

Set associations info

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiAssocsResp := *openapiclient.NewV0044OpenapiAssocsResp([]openapiclient.V0044Assoc{*openapiclient.NewV0044Assoc("User_example")}) // V0044OpenapiAssocsResp | Job description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostAssociations(context.Background()).V0044OpenapiAssocsResp(v0044OpenapiAssocsResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostAssociations`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiAssocsResp** | [**V0044OpenapiAssocsResp**](V0044OpenapiAssocsResp.md) | Job description | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostClusters

> V0044OpenapiResp SlurmdbV0044PostClusters(ctx).UpdateTime(updateTime).V0044OpenapiClustersResp(v0044OpenapiClustersResp).Execute()

Get cluster list

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query reservations updated more recently than this time (UNIX timestamp) (optional)
	v0044OpenapiClustersResp := *openapiclient.NewV0044OpenapiClustersResp([]openapiclient.V0044ClusterRec{*openapiclient.NewV0044ClusterRec()}) // V0044OpenapiClustersResp | Cluster add or update descriptions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostClusters(context.Background()).UpdateTime(updateTime).V0044OpenapiClustersResp(v0044OpenapiClustersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostClusters`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query reservations updated more recently than this time (UNIX timestamp) | 
 **v0044OpenapiClustersResp** | [**V0044OpenapiClustersResp**](V0044OpenapiClustersResp.md) | Cluster add or update descriptions | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostConfig

> V0044OpenapiResp SlurmdbV0044PostConfig(ctx).V0044OpenapiSlurmdbdConfigResp(v0044OpenapiSlurmdbdConfigResp).Execute()

Load all configuration information

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiSlurmdbdConfigResp := *openapiclient.NewV0044OpenapiSlurmdbdConfigResp() // V0044OpenapiSlurmdbdConfigResp | Add or update config (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostConfig(context.Background()).V0044OpenapiSlurmdbdConfigResp(v0044OpenapiSlurmdbdConfigResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostConfig`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiSlurmdbdConfigResp** | [**V0044OpenapiSlurmdbdConfigResp**](V0044OpenapiSlurmdbdConfigResp.md) | Add or update config | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostJob

> V0044OpenapiJobModifyResp SlurmdbV0044PostJob(ctx, jobId).V0044JobModify(v0044JobModify).Execute()

Update job

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	jobId := "jobId_example" // string | Job ID
	v0044JobModify := *openapiclient.NewV0044JobModify() // V0044JobModify | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostJob(context.Background(), jobId).V0044JobModify(v0044JobModify).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostJob``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostJob`: V0044OpenapiJobModifyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0044JobModify** | [**V0044JobModify**](V0044JobModify.md) | Job update description | 

### Return type

[**V0044OpenapiJobModifyResp**](V0044OpenapiJobModifyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostJobs

> V0044OpenapiJobModifyResp SlurmdbV0044PostJobs(ctx).V0044OpenapiJobModifyReq(v0044OpenapiJobModifyReq).Execute()

Update jobs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiJobModifyReq := *openapiclient.NewV0044OpenapiJobModifyReq() // V0044OpenapiJobModifyReq | Job update description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostJobs(context.Background()).V0044OpenapiJobModifyReq(v0044OpenapiJobModifyReq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostJobs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostJobs`: V0044OpenapiJobModifyResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostJobs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostJobsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiJobModifyReq** | [**V0044OpenapiJobModifyReq**](V0044OpenapiJobModifyReq.md) | Job update description | 

### Return type

[**V0044OpenapiJobModifyResp**](V0044OpenapiJobModifyResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostQos

> V0044OpenapiResp SlurmdbV0044PostQos(ctx).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0044OpenapiSlurmdbdQosResp(v0044OpenapiSlurmdbdQosResp).Execute()

Add or update QOSs

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	description := "description_example" // string | CSV description list (optional)
	includeDeletedQOS := "includeDeletedQOS_example" // string |  (optional)
	id := "id_example" // string | CSV QOS id list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	name := "name_example" // string | CSV QOS name list (optional)
	preemptMode := "preemptMode_example" // string | PreemptMode used when jobs in this QOS are preempted (optional)
	v0044OpenapiSlurmdbdQosResp := *openapiclient.NewV0044OpenapiSlurmdbdQosResp([]openapiclient.V0044Qos{*openapiclient.NewV0044Qos()}) // V0044OpenapiSlurmdbdQosResp | Description of QOS to add or update (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostQos(context.Background()).Description(description).IncludeDeletedQOS(includeDeletedQOS).Id(id).Format(format).Name(name).PreemptMode(preemptMode).V0044OpenapiSlurmdbdQosResp(v0044OpenapiSlurmdbdQosResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostQos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostQos`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostQos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostQosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **description** | **string** | CSV description list | 
 **includeDeletedQOS** | **string** |  | 
 **id** | **string** | CSV QOS id list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **name** | **string** | CSV QOS name list | 
 **preemptMode** | **string** | PreemptMode used when jobs in this QOS are preempted | 
 **v0044OpenapiSlurmdbdQosResp** | [**V0044OpenapiSlurmdbdQosResp**](V0044OpenapiSlurmdbdQosResp.md) | Description of QOS to add or update | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostTres

> V0044OpenapiResp SlurmdbV0044PostTres(ctx).V0044OpenapiTresResp(v0044OpenapiTresResp).Execute()

Add TRES

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiTresResp := *openapiclient.NewV0044OpenapiTresResp([]openapiclient.V0044Tres{*openapiclient.NewV0044Tres("Type_example")}) // V0044OpenapiTresResp | TRES descriptions. Only works in developer mode. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostTres(context.Background()).V0044OpenapiTresResp(v0044OpenapiTresResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostTres``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostTres`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostTres`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostTresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiTresResp** | [**V0044OpenapiTresResp**](V0044OpenapiTresResp.md) | TRES descriptions. Only works in developer mode. | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostUsers

> V0044OpenapiResp SlurmdbV0044PostUsers(ctx).V0044OpenapiUsersResp(v0044OpenapiUsersResp).Execute()

Update users

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	v0044OpenapiUsersResp := *openapiclient.NewV0044OpenapiUsersResp([]openapiclient.V0044User{*openapiclient.NewV0044User("Name_example")}) // V0044OpenapiUsersResp | add or update user (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostUsers(context.Background()).V0044OpenapiUsersResp(v0044OpenapiUsersResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostUsers`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0044OpenapiUsersResp** | [**V0044OpenapiUsersResp**](V0044OpenapiUsersResp.md) | add or update user | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostUsersAssociation

> V0044OpenapiUsersAddCondRespStr SlurmdbV0044PostUsersAssociation(ctx).UpdateTime(updateTime).Flags(flags).V0044OpenapiUsersAddCondResp(v0044OpenapiUsersAddCondResp).Execute()

Add users with conditional association

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	updateTime := "updateTime_example" // string | Query partitions updated more recently than this time (UNIX timestamp) (optional)
	flags := "flags_example" // string | Query flags (optional)
	v0044OpenapiUsersAddCondResp := *openapiclient.NewV0044OpenapiUsersAddCondResp(*openapiclient.NewV0044UsersAddCond([]string{"Users_example"}), *openapiclient.NewV0044UserShort()) // V0044OpenapiUsersAddCondResp | Create users with conditional association (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostUsersAssociation(context.Background()).UpdateTime(updateTime).Flags(flags).V0044OpenapiUsersAddCondResp(v0044OpenapiUsersAddCondResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostUsersAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostUsersAssociation`: V0044OpenapiUsersAddCondRespStr
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostUsersAssociation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostUsersAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateTime** | **string** | Query partitions updated more recently than this time (UNIX timestamp) | 
 **flags** | **string** | Query flags | 
 **v0044OpenapiUsersAddCondResp** | [**V0044OpenapiUsersAddCondResp**](V0044OpenapiUsersAddCondResp.md) | Create users with conditional association | 

### Return type

[**V0044OpenapiUsersAddCondRespStr**](V0044OpenapiUsersAddCondRespStr.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmdbV0044PostWckeys

> V0044OpenapiResp SlurmdbV0044PostWckeys(ctx).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0044OpenapiWckeyResp(v0044OpenapiWckeyResp).Execute()

Add or update wckeys

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	cluster := "cluster_example" // string | CSV cluster name list (optional)
	format := "format_example" // string | Ignored; process JSON manually to control output format (optional)
	id := "id_example" // string | CSV ID list (optional)
	name := "name_example" // string | CSV name list (optional)
	onlyDefaults := "onlyDefaults_example" // string | Only query defaults (optional)
	usageEnd := "usageEnd_example" // string | Usage end (UNIX timestamp) (optional)
	usageStart := "usageStart_example" // string | Usage start (UNIX timestamp) (optional)
	user := "user_example" // string | CSV user list (optional)
	withUsage := "withUsage_example" // string | Include usage (optional)
	withDeleted := "withDeleted_example" // string | Include deleted WCKeys (optional)
	v0044OpenapiWckeyResp := *openapiclient.NewV0044OpenapiWckeyResp([]openapiclient.V0044Wckey{*openapiclient.NewV0044Wckey("Cluster_example", "Name_example", "User_example")}) // V0044OpenapiWckeyResp | wckeys description (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SlurmdbAPI.SlurmdbV0044PostWckeys(context.Background()).Cluster(cluster).Format(format).Id(id).Name(name).OnlyDefaults(onlyDefaults).UsageEnd(usageEnd).UsageStart(usageStart).User(user).WithUsage(withUsage).WithDeleted(withDeleted).V0044OpenapiWckeyResp(v0044OpenapiWckeyResp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SlurmdbAPI.SlurmdbV0044PostWckeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SlurmdbV0044PostWckeys`: V0044OpenapiResp
	fmt.Fprintf(os.Stdout, "Response from `SlurmdbAPI.SlurmdbV0044PostWckeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmdbV0044PostWckeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cluster** | **string** | CSV cluster name list | 
 **format** | **string** | Ignored; process JSON manually to control output format | 
 **id** | **string** | CSV ID list | 
 **name** | **string** | CSV name list | 
 **onlyDefaults** | **string** | Only query defaults | 
 **usageEnd** | **string** | Usage end (UNIX timestamp) | 
 **usageStart** | **string** | Usage start (UNIX timestamp) | 
 **user** | **string** | CSV user list | 
 **withUsage** | **string** | Include usage | 
 **withDeleted** | **string** | Include deleted WCKeys | 
 **v0044OpenapiWckeyResp** | [**V0044OpenapiWckeyResp**](V0044OpenapiWckeyResp.md) | wckeys description | 

### Return type

[**V0044OpenapiResp**](V0044OpenapiResp.md)

### Authorization

[user](../README.md#user), [bearerAuth](../README.md#bearerAuth), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

