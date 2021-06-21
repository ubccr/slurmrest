# \OpenapiApi

All URIs are relative to *http://localhost/slurm/v0.0.36*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OpenapiGet**](OpenapiApi.md#OpenapiGet) | **Get** /openapi | Retrieve OpenAPI Specification
[**OpenapiJsonGet**](OpenapiApi.md#OpenapiJsonGet) | **Get** /openapi.json | Retrieve OpenAPI Specification
[**OpenapiV3Get**](OpenapiApi.md#OpenapiV3Get) | **Get** /openapi/v3 | Retrieve OpenAPI Specification
[**OpenapiYamlGet**](OpenapiApi.md#OpenapiYamlGet) | **Get** /openapi.yaml | Retrieve OpenAPI Specification



## OpenapiGet

> OpenapiGet(ctx).Execute()

Retrieve OpenAPI Specification

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
    resp, r, err := api_client.OpenapiApi.OpenapiGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiGetRequest struct via the builder pattern


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


## OpenapiJsonGet

> OpenapiJsonGet(ctx).Execute()

Retrieve OpenAPI Specification

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
    resp, r, err := api_client.OpenapiApi.OpenapiJsonGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiJsonGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiJsonGetRequest struct via the builder pattern


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


## OpenapiV3Get

> OpenapiV3Get(ctx).Execute()

Retrieve OpenAPI Specification

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
    resp, r, err := api_client.OpenapiApi.OpenapiV3Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiV3Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiV3GetRequest struct via the builder pattern


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


## OpenapiYamlGet

> OpenapiYamlGet(ctx).Execute()

Retrieve OpenAPI Specification

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
    resp, r, err := api_client.OpenapiApi.OpenapiYamlGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OpenapiApi.OpenapiYamlGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiOpenapiYamlGetRequest struct via the builder pattern


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

