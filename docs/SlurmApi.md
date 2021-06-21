# \SlurmApi

All URIs are relative to *http://localhost/slurm/v0.0.36*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SlurmctldCancelJob**](SlurmApi.md#SlurmctldCancelJob) | **Delete** /job/{job_id} | cancel or signal job
[**SlurmctldDiag**](SlurmApi.md#SlurmctldDiag) | **Get** /diag/ | get diagnostics
[**SlurmctldGetJob**](SlurmApi.md#SlurmctldGetJob) | **Get** /job/{job_id} | get job info
[**SlurmctldGetJobs**](SlurmApi.md#SlurmctldGetJobs) | **Get** /jobs/ | get list of jobs
[**SlurmctldGetNode**](SlurmApi.md#SlurmctldGetNode) | **Get** /node/{node_name} | get node info
[**SlurmctldGetNodes**](SlurmApi.md#SlurmctldGetNodes) | **Get** /nodes/ | get all node info
[**SlurmctldGetPartition**](SlurmApi.md#SlurmctldGetPartition) | **Get** /partition/{partition_name} | get partition info
[**SlurmctldGetPartitions**](SlurmApi.md#SlurmctldGetPartitions) | **Get** /partitions/ | get all partition info
[**SlurmctldPing**](SlurmApi.md#SlurmctldPing) | **Get** /ping/ | ping test
[**SlurmctldSubmitJob**](SlurmApi.md#SlurmctldSubmitJob) | **Post** /job/submit | submit new job
[**SlurmctldUpdateJob**](SlurmApi.md#SlurmctldUpdateJob) | **Post** /job/{job_id} | update job



## SlurmctldCancelJob

> SlurmctldCancelJob(ctx, jobId).Signal(signal).Execute()

cancel or signal job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := int64(789) // int64 | Slurm Job ID
    signal := openapiclient.v0.0.36_signal("HUP") // V0036Signal | signal to send to job (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldCancelJob(context.Background(), jobId).Signal(signal).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldCancelJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldCancelJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | [**V0036Signal**](V0036Signal.md) | signal to send to job | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldDiag

> V0036Diag SlurmctldDiag(ctx).Execute()

get diagnostics

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldDiag(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldDiag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldDiag`: V0036Diag
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldDiag`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldDiagRequest struct via the builder pattern


### Return type

[**V0036Diag**](V0036Diag.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetJob

> V0036JobsResponse SlurmctldGetJob(ctx, jobId).Execute()

get job info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := int64(789) // int64 | Slurm Job ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldGetJob(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJob`: V0036JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0036JobsResponse**](V0036JobsResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetJobs

> V0036JobsResponse SlurmctldGetJobs(ctx).Execute()

get list of jobs

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldGetJobs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetJobs`: V0036JobsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetJobsRequest struct via the builder pattern


### Return type

[**V0036JobsResponse**](V0036JobsResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetNode

> V0036NodesResponse SlurmctldGetNode(ctx, nodeName).Execute()

get node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    nodeName := "nodeName_example" // string | Slurm Node Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldGetNode(context.Background(), nodeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNode`: V0036NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeName** | **string** | Slurm Node Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0036NodesResponse**](V0036NodesResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetNodes

> V0036NodesResponse SlurmctldGetNodes(ctx).Execute()

get all node info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldGetNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetNodes`: V0036NodesResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetNodesRequest struct via the builder pattern


### Return type

[**V0036NodesResponse**](V0036NodesResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetPartition

> V0036PartitionsResponse SlurmctldGetPartition(ctx, partitionName).Execute()

get partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    partitionName := "partitionName_example" // string | Slurm Partition Name

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldGetPartition(context.Background(), partitionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartition`: V0036PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**partitionName** | **string** | Slurm Partition Name | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V0036PartitionsResponse**](V0036PartitionsResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldGetPartitions

> V0036PartitionsResponse SlurmctldGetPartitions(ctx).Execute()

get all partition info

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldGetPartitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldGetPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldGetPartitions`: V0036PartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldGetPartitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldGetPartitionsRequest struct via the builder pattern


### Return type

[**V0036PartitionsResponse**](V0036PartitionsResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldPing

> V0036Pings SlurmctldPing(ctx).Execute()

ping test

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldPing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldPing`: V0036Pings
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldPing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldPingRequest struct via the builder pattern


### Return type

[**V0036Pings**](V0036Pings.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldSubmitJob

> V0036JobSubmissionResponse SlurmctldSubmitJob(ctx).V0036JobSubmission(v0036JobSubmission).Execute()

submit new job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    v0036JobSubmission := *openapiclient.NewV0036JobSubmission("Script_example") // V0036JobSubmission | submit new job

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldSubmitJob(context.Background()).V0036JobSubmission(v0036JobSubmission).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldSubmitJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SlurmctldSubmitJob`: V0036JobSubmissionResponse
    fmt.Fprintf(os.Stdout, "Response from `SlurmApi.SlurmctldSubmitJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldSubmitJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v0036JobSubmission** | [**V0036JobSubmission**](V0036JobSubmission.md) | submit new job | 

### Return type

[**V0036JobSubmissionResponse**](V0036JobSubmissionResponse.md)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: application/json, application/x-yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SlurmctldUpdateJob

> SlurmctldUpdateJob(ctx, jobId).V0036JobProperties(v0036JobProperties).Execute()

update job

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := int64(789) // int64 | Slurm Job ID
    v0036JobProperties := *openapiclient.NewV0036JobProperties(map[string]interface{}(123)) // V0036JobProperties | update job

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SlurmApi.SlurmctldUpdateJob(context.Background(), jobId).V0036JobProperties(v0036JobProperties).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SlurmApi.SlurmctldUpdateJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **int64** | Slurm Job ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSlurmctldUpdateJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v0036JobProperties** | [**V0036JobProperties**](V0036JobProperties.md) | update job | 

### Return type

 (empty response body)

### Authorization

[token](../README.md#token), [user](../README.md#user)

### HTTP request headers

- **Content-Type**: application/json, application/x-yaml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

