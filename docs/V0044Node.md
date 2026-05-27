# V0044Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | Computer architecture | [optional] 
**BurstbufferNetworkAddress** | Pointer to **string** | Alternate network path to be used for sbcast network traffic | [optional] 
**Boards** | Pointer to **int32** | Number of Baseboards in nodes with a baseboard controller | [optional] 
**BootTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**TlsCertLastRenewal** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**CertFlags** | Pointer to **[]string** | Certmgr status flags | [optional] 
**ClusterName** | Pointer to **string** | Cluster name (only set in federated environments) | [optional] 
**Cores** | Pointer to **int32** | Number of cores in a single physical processor socket | [optional] 
**SpecializedCores** | Pointer to **int32** | Number of cores reserved for system use | [optional] 
**CpuBinding** | Pointer to **int32** | Default method for binding tasks to allocated CPUs | [optional] 
**CpuLoad** | Pointer to **int32** | CPU load as reported by the OS | [optional] 
**FreeMem** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Cpus** | Pointer to **int32** | Total CPUs, including cores and threads | [optional] 
**EffectiveCpus** | Pointer to **int32** | Number of effective CPUs (excluding specialized CPUs) | [optional] 
**SpecializedCpus** | Pointer to **string** | Abstract CPU IDs on this node reserved for exclusive use by slurmd and slurmstepd | [optional] 
**Energy** | Pointer to [**V0044AcctGatherEnergy**](V0044AcctGatherEnergy.md) |  | [optional] 
**ExternalSensors** | Pointer to **map[string]interface{}** |  | [optional] 
**Extra** | Pointer to **string** | Arbitrary string used for node filtering if extra constraints are enabled | [optional] 
**Power** | Pointer to **map[string]interface{}** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**ActiveFeatures** | Pointer to **[]string** |  | [optional] 
**GpuSpec** | Pointer to **string** | CPU cores reserved for jobs that also use a GPU | [optional] 
**Gres** | Pointer to **string** | Generic resources | [optional] 
**GresDrained** | Pointer to **string** | Drained generic resources | [optional] 
**GresUsed** | Pointer to **string** | Generic resources currently in use | [optional] 
**InstanceId** | Pointer to **string** | Cloud instance ID | [optional] 
**InstanceType** | Pointer to **string** | Cloud instance type | [optional] 
**LastBusy** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**McsLabel** | Pointer to **string** | Multi-Category Security label | [optional] 
**SpecializedMemory** | Pointer to **int64** | Combined memory limit, in MB, for Slurm compute node daemons | [optional] 
**Name** | Pointer to **string** | NodeName | [optional] 
**NextStateAfterReboot** | Pointer to **[]string** | The state the node will be assigned after rebooting | [optional] 
**Address** | Pointer to **string** | NodeAddr, used to establish a communication path | [optional] 
**Hostname** | Pointer to **string** | NodeHostname | [optional] 
**State** | Pointer to **[]string** | Node state(s) applicable to this node | [optional] 
**OperatingSystem** | Pointer to **string** | Operating system reported by the node | [optional] 
**Owner** | Pointer to **string** | User allowed to run jobs on this node (unset if no restriction) | [optional] 
**Partitions** | Pointer to **[]string** |  | [optional] 
**Port** | Pointer to **int32** | TCP port number of the slurmd | [optional] 
**RealMemory** | Pointer to **int64** | Total memory in MB on the node | [optional] 
**ResCoresPerGpu** | Pointer to **int32** | Number of CPU cores per GPU restricted to GPU jobs | [optional] 
**Comment** | Pointer to **string** | Arbitrary comment | [optional] 
**Reason** | Pointer to **string** | Describes why the node is in a \&quot;DOWN\&quot;, \&quot;DRAINED\&quot;, \&quot;DRAINING\&quot;, \&quot;FAILING\&quot; or \&quot;FAIL\&quot; state | [optional] 
**ReasonChangedAt** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**ReasonSetByUser** | Pointer to **string** | User who set the reason | [optional] 
**ResumeAfter** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Reservation** | Pointer to **string** | Name of reservation containing this node | [optional] 
**AllocMemory** | Pointer to **int64** | Total memory in MB currently allocated for jobs | [optional] 
**AllocCpus** | Pointer to **int32** | Total number of CPUs currently allocated for jobs | [optional] 
**AllocIdleCpus** | Pointer to **int32** | Total number of idle CPUs | [optional] 
**TresUsed** | Pointer to **string** | Trackable resources currently allocated for jobs | [optional] 
**TresWeighted** | Pointer to **float64** | Ignored. Was weighted number of billable trackable resources allocated | [optional] 
**SlurmdStartTime** | Pointer to [**V0044Uint64NoValStruct**](V0044Uint64NoValStruct.md) |  | [optional] 
**Sockets** | Pointer to **int32** | Number of physical processor sockets/chips on the node | [optional] 
**Threads** | Pointer to **int32** | Number of logical threads in a single physical core | [optional] 
**TemporaryDisk** | Pointer to **int32** | Total size in MB of temporary disk storage in TmpFS | [optional] 
**Weight** | Pointer to **int32** | Weight of the node for scheduling purposes | [optional] 
**Topology** | Pointer to **string** | Topology | [optional] 
**Tres** | Pointer to **string** | Configured trackable resources | [optional] 
**Version** | Pointer to **string** | Slurmd version | [optional] 

## Methods

### NewV0044Node

`func NewV0044Node() *V0044Node`

NewV0044Node instantiates a new V0044Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0044NodeWithDefaults

`func NewV0044NodeWithDefaults() *V0044Node`

NewV0044NodeWithDefaults instantiates a new V0044Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *V0044Node) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *V0044Node) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *V0044Node) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *V0044Node) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBurstbufferNetworkAddress

`func (o *V0044Node) GetBurstbufferNetworkAddress() string`

GetBurstbufferNetworkAddress returns the BurstbufferNetworkAddress field if non-nil, zero value otherwise.

### GetBurstbufferNetworkAddressOk

`func (o *V0044Node) GetBurstbufferNetworkAddressOk() (*string, bool)`

GetBurstbufferNetworkAddressOk returns a tuple with the BurstbufferNetworkAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurstbufferNetworkAddress

`func (o *V0044Node) SetBurstbufferNetworkAddress(v string)`

SetBurstbufferNetworkAddress sets BurstbufferNetworkAddress field to given value.

### HasBurstbufferNetworkAddress

`func (o *V0044Node) HasBurstbufferNetworkAddress() bool`

HasBurstbufferNetworkAddress returns a boolean if a field has been set.

### GetBoards

`func (o *V0044Node) GetBoards() int32`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *V0044Node) GetBoardsOk() (*int32, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *V0044Node) SetBoards(v int32)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *V0044Node) HasBoards() bool`

HasBoards returns a boolean if a field has been set.

### GetBootTime

`func (o *V0044Node) GetBootTime() V0044Uint64NoValStruct`

GetBootTime returns the BootTime field if non-nil, zero value otherwise.

### GetBootTimeOk

`func (o *V0044Node) GetBootTimeOk() (*V0044Uint64NoValStruct, bool)`

GetBootTimeOk returns a tuple with the BootTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootTime

`func (o *V0044Node) SetBootTime(v V0044Uint64NoValStruct)`

SetBootTime sets BootTime field to given value.

### HasBootTime

`func (o *V0044Node) HasBootTime() bool`

HasBootTime returns a boolean if a field has been set.

### GetTlsCertLastRenewal

`func (o *V0044Node) GetTlsCertLastRenewal() V0044Uint64NoValStruct`

GetTlsCertLastRenewal returns the TlsCertLastRenewal field if non-nil, zero value otherwise.

### GetTlsCertLastRenewalOk

`func (o *V0044Node) GetTlsCertLastRenewalOk() (*V0044Uint64NoValStruct, bool)`

GetTlsCertLastRenewalOk returns a tuple with the TlsCertLastRenewal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsCertLastRenewal

`func (o *V0044Node) SetTlsCertLastRenewal(v V0044Uint64NoValStruct)`

SetTlsCertLastRenewal sets TlsCertLastRenewal field to given value.

### HasTlsCertLastRenewal

`func (o *V0044Node) HasTlsCertLastRenewal() bool`

HasTlsCertLastRenewal returns a boolean if a field has been set.

### GetCertFlags

`func (o *V0044Node) GetCertFlags() []string`

GetCertFlags returns the CertFlags field if non-nil, zero value otherwise.

### GetCertFlagsOk

`func (o *V0044Node) GetCertFlagsOk() (*[]string, bool)`

GetCertFlagsOk returns a tuple with the CertFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFlags

`func (o *V0044Node) SetCertFlags(v []string)`

SetCertFlags sets CertFlags field to given value.

### HasCertFlags

`func (o *V0044Node) HasCertFlags() bool`

HasCertFlags returns a boolean if a field has been set.

### GetClusterName

`func (o *V0044Node) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *V0044Node) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *V0044Node) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.

### HasClusterName

`func (o *V0044Node) HasClusterName() bool`

HasClusterName returns a boolean if a field has been set.

### GetCores

`func (o *V0044Node) GetCores() int32`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *V0044Node) GetCoresOk() (*int32, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *V0044Node) SetCores(v int32)`

SetCores sets Cores field to given value.

### HasCores

`func (o *V0044Node) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetSpecializedCores

`func (o *V0044Node) GetSpecializedCores() int32`

GetSpecializedCores returns the SpecializedCores field if non-nil, zero value otherwise.

### GetSpecializedCoresOk

`func (o *V0044Node) GetSpecializedCoresOk() (*int32, bool)`

GetSpecializedCoresOk returns a tuple with the SpecializedCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializedCores

`func (o *V0044Node) SetSpecializedCores(v int32)`

SetSpecializedCores sets SpecializedCores field to given value.

### HasSpecializedCores

`func (o *V0044Node) HasSpecializedCores() bool`

HasSpecializedCores returns a boolean if a field has been set.

### GetCpuBinding

`func (o *V0044Node) GetCpuBinding() int32`

GetCpuBinding returns the CpuBinding field if non-nil, zero value otherwise.

### GetCpuBindingOk

`func (o *V0044Node) GetCpuBindingOk() (*int32, bool)`

GetCpuBindingOk returns a tuple with the CpuBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuBinding

`func (o *V0044Node) SetCpuBinding(v int32)`

SetCpuBinding sets CpuBinding field to given value.

### HasCpuBinding

`func (o *V0044Node) HasCpuBinding() bool`

HasCpuBinding returns a boolean if a field has been set.

### GetCpuLoad

`func (o *V0044Node) GetCpuLoad() int32`

GetCpuLoad returns the CpuLoad field if non-nil, zero value otherwise.

### GetCpuLoadOk

`func (o *V0044Node) GetCpuLoadOk() (*int32, bool)`

GetCpuLoadOk returns a tuple with the CpuLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuLoad

`func (o *V0044Node) SetCpuLoad(v int32)`

SetCpuLoad sets CpuLoad field to given value.

### HasCpuLoad

`func (o *V0044Node) HasCpuLoad() bool`

HasCpuLoad returns a boolean if a field has been set.

### GetFreeMem

`func (o *V0044Node) GetFreeMem() V0044Uint64NoValStruct`

GetFreeMem returns the FreeMem field if non-nil, zero value otherwise.

### GetFreeMemOk

`func (o *V0044Node) GetFreeMemOk() (*V0044Uint64NoValStruct, bool)`

GetFreeMemOk returns a tuple with the FreeMem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMem

`func (o *V0044Node) SetFreeMem(v V0044Uint64NoValStruct)`

SetFreeMem sets FreeMem field to given value.

### HasFreeMem

`func (o *V0044Node) HasFreeMem() bool`

HasFreeMem returns a boolean if a field has been set.

### GetCpus

`func (o *V0044Node) GetCpus() int32`

GetCpus returns the Cpus field if non-nil, zero value otherwise.

### GetCpusOk

`func (o *V0044Node) GetCpusOk() (*int32, bool)`

GetCpusOk returns a tuple with the Cpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpus

`func (o *V0044Node) SetCpus(v int32)`

SetCpus sets Cpus field to given value.

### HasCpus

`func (o *V0044Node) HasCpus() bool`

HasCpus returns a boolean if a field has been set.

### GetEffectiveCpus

`func (o *V0044Node) GetEffectiveCpus() int32`

GetEffectiveCpus returns the EffectiveCpus field if non-nil, zero value otherwise.

### GetEffectiveCpusOk

`func (o *V0044Node) GetEffectiveCpusOk() (*int32, bool)`

GetEffectiveCpusOk returns a tuple with the EffectiveCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCpus

`func (o *V0044Node) SetEffectiveCpus(v int32)`

SetEffectiveCpus sets EffectiveCpus field to given value.

### HasEffectiveCpus

`func (o *V0044Node) HasEffectiveCpus() bool`

HasEffectiveCpus returns a boolean if a field has been set.

### GetSpecializedCpus

`func (o *V0044Node) GetSpecializedCpus() string`

GetSpecializedCpus returns the SpecializedCpus field if non-nil, zero value otherwise.

### GetSpecializedCpusOk

`func (o *V0044Node) GetSpecializedCpusOk() (*string, bool)`

GetSpecializedCpusOk returns a tuple with the SpecializedCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializedCpus

`func (o *V0044Node) SetSpecializedCpus(v string)`

SetSpecializedCpus sets SpecializedCpus field to given value.

### HasSpecializedCpus

`func (o *V0044Node) HasSpecializedCpus() bool`

HasSpecializedCpus returns a boolean if a field has been set.

### GetEnergy

`func (o *V0044Node) GetEnergy() V0044AcctGatherEnergy`

GetEnergy returns the Energy field if non-nil, zero value otherwise.

### GetEnergyOk

`func (o *V0044Node) GetEnergyOk() (*V0044AcctGatherEnergy, bool)`

GetEnergyOk returns a tuple with the Energy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnergy

`func (o *V0044Node) SetEnergy(v V0044AcctGatherEnergy)`

SetEnergy sets Energy field to given value.

### HasEnergy

`func (o *V0044Node) HasEnergy() bool`

HasEnergy returns a boolean if a field has been set.

### GetExternalSensors

`func (o *V0044Node) GetExternalSensors() map[string]interface{}`

GetExternalSensors returns the ExternalSensors field if non-nil, zero value otherwise.

### GetExternalSensorsOk

`func (o *V0044Node) GetExternalSensorsOk() (*map[string]interface{}, bool)`

GetExternalSensorsOk returns a tuple with the ExternalSensors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSensors

`func (o *V0044Node) SetExternalSensors(v map[string]interface{})`

SetExternalSensors sets ExternalSensors field to given value.

### HasExternalSensors

`func (o *V0044Node) HasExternalSensors() bool`

HasExternalSensors returns a boolean if a field has been set.

### GetExtra

`func (o *V0044Node) GetExtra() string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *V0044Node) GetExtraOk() (*string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *V0044Node) SetExtra(v string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *V0044Node) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetPower

`func (o *V0044Node) GetPower() map[string]interface{}`

GetPower returns the Power field if non-nil, zero value otherwise.

### GetPowerOk

`func (o *V0044Node) GetPowerOk() (*map[string]interface{}, bool)`

GetPowerOk returns a tuple with the Power field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPower

`func (o *V0044Node) SetPower(v map[string]interface{})`

SetPower sets Power field to given value.

### HasPower

`func (o *V0044Node) HasPower() bool`

HasPower returns a boolean if a field has been set.

### GetFeatures

`func (o *V0044Node) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *V0044Node) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *V0044Node) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *V0044Node) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetActiveFeatures

`func (o *V0044Node) GetActiveFeatures() []string`

GetActiveFeatures returns the ActiveFeatures field if non-nil, zero value otherwise.

### GetActiveFeaturesOk

`func (o *V0044Node) GetActiveFeaturesOk() (*[]string, bool)`

GetActiveFeaturesOk returns a tuple with the ActiveFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFeatures

`func (o *V0044Node) SetActiveFeatures(v []string)`

SetActiveFeatures sets ActiveFeatures field to given value.

### HasActiveFeatures

`func (o *V0044Node) HasActiveFeatures() bool`

HasActiveFeatures returns a boolean if a field has been set.

### GetGpuSpec

`func (o *V0044Node) GetGpuSpec() string`

GetGpuSpec returns the GpuSpec field if non-nil, zero value otherwise.

### GetGpuSpecOk

`func (o *V0044Node) GetGpuSpecOk() (*string, bool)`

GetGpuSpecOk returns a tuple with the GpuSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuSpec

`func (o *V0044Node) SetGpuSpec(v string)`

SetGpuSpec sets GpuSpec field to given value.

### HasGpuSpec

`func (o *V0044Node) HasGpuSpec() bool`

HasGpuSpec returns a boolean if a field has been set.

### GetGres

`func (o *V0044Node) GetGres() string`

GetGres returns the Gres field if non-nil, zero value otherwise.

### GetGresOk

`func (o *V0044Node) GetGresOk() (*string, bool)`

GetGresOk returns a tuple with the Gres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGres

`func (o *V0044Node) SetGres(v string)`

SetGres sets Gres field to given value.

### HasGres

`func (o *V0044Node) HasGres() bool`

HasGres returns a boolean if a field has been set.

### GetGresDrained

`func (o *V0044Node) GetGresDrained() string`

GetGresDrained returns the GresDrained field if non-nil, zero value otherwise.

### GetGresDrainedOk

`func (o *V0044Node) GetGresDrainedOk() (*string, bool)`

GetGresDrainedOk returns a tuple with the GresDrained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresDrained

`func (o *V0044Node) SetGresDrained(v string)`

SetGresDrained sets GresDrained field to given value.

### HasGresDrained

`func (o *V0044Node) HasGresDrained() bool`

HasGresDrained returns a boolean if a field has been set.

### GetGresUsed

`func (o *V0044Node) GetGresUsed() string`

GetGresUsed returns the GresUsed field if non-nil, zero value otherwise.

### GetGresUsedOk

`func (o *V0044Node) GetGresUsedOk() (*string, bool)`

GetGresUsedOk returns a tuple with the GresUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGresUsed

`func (o *V0044Node) SetGresUsed(v string)`

SetGresUsed sets GresUsed field to given value.

### HasGresUsed

`func (o *V0044Node) HasGresUsed() bool`

HasGresUsed returns a boolean if a field has been set.

### GetInstanceId

`func (o *V0044Node) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *V0044Node) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *V0044Node) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *V0044Node) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetInstanceType

`func (o *V0044Node) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *V0044Node) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *V0044Node) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *V0044Node) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetLastBusy

`func (o *V0044Node) GetLastBusy() V0044Uint64NoValStruct`

GetLastBusy returns the LastBusy field if non-nil, zero value otherwise.

### GetLastBusyOk

`func (o *V0044Node) GetLastBusyOk() (*V0044Uint64NoValStruct, bool)`

GetLastBusyOk returns a tuple with the LastBusy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBusy

`func (o *V0044Node) SetLastBusy(v V0044Uint64NoValStruct)`

SetLastBusy sets LastBusy field to given value.

### HasLastBusy

`func (o *V0044Node) HasLastBusy() bool`

HasLastBusy returns a boolean if a field has been set.

### GetMcsLabel

`func (o *V0044Node) GetMcsLabel() string`

GetMcsLabel returns the McsLabel field if non-nil, zero value otherwise.

### GetMcsLabelOk

`func (o *V0044Node) GetMcsLabelOk() (*string, bool)`

GetMcsLabelOk returns a tuple with the McsLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcsLabel

`func (o *V0044Node) SetMcsLabel(v string)`

SetMcsLabel sets McsLabel field to given value.

### HasMcsLabel

`func (o *V0044Node) HasMcsLabel() bool`

HasMcsLabel returns a boolean if a field has been set.

### GetSpecializedMemory

`func (o *V0044Node) GetSpecializedMemory() int64`

GetSpecializedMemory returns the SpecializedMemory field if non-nil, zero value otherwise.

### GetSpecializedMemoryOk

`func (o *V0044Node) GetSpecializedMemoryOk() (*int64, bool)`

GetSpecializedMemoryOk returns a tuple with the SpecializedMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecializedMemory

`func (o *V0044Node) SetSpecializedMemory(v int64)`

SetSpecializedMemory sets SpecializedMemory field to given value.

### HasSpecializedMemory

`func (o *V0044Node) HasSpecializedMemory() bool`

HasSpecializedMemory returns a boolean if a field has been set.

### GetName

`func (o *V0044Node) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V0044Node) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V0044Node) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V0044Node) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextStateAfterReboot

`func (o *V0044Node) GetNextStateAfterReboot() []string`

GetNextStateAfterReboot returns the NextStateAfterReboot field if non-nil, zero value otherwise.

### GetNextStateAfterRebootOk

`func (o *V0044Node) GetNextStateAfterRebootOk() (*[]string, bool)`

GetNextStateAfterRebootOk returns a tuple with the NextStateAfterReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStateAfterReboot

`func (o *V0044Node) SetNextStateAfterReboot(v []string)`

SetNextStateAfterReboot sets NextStateAfterReboot field to given value.

### HasNextStateAfterReboot

`func (o *V0044Node) HasNextStateAfterReboot() bool`

HasNextStateAfterReboot returns a boolean if a field has been set.

### GetAddress

`func (o *V0044Node) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *V0044Node) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *V0044Node) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *V0044Node) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetHostname

`func (o *V0044Node) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0044Node) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0044Node) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0044Node) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetState

`func (o *V0044Node) GetState() []string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *V0044Node) GetStateOk() (*[]string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *V0044Node) SetState(v []string)`

SetState sets State field to given value.

### HasState

`func (o *V0044Node) HasState() bool`

HasState returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *V0044Node) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *V0044Node) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *V0044Node) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *V0044Node) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOwner

`func (o *V0044Node) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *V0044Node) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *V0044Node) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *V0044Node) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPartitions

`func (o *V0044Node) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *V0044Node) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *V0044Node) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *V0044Node) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetPort

`func (o *V0044Node) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V0044Node) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V0044Node) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V0044Node) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRealMemory

`func (o *V0044Node) GetRealMemory() int64`

GetRealMemory returns the RealMemory field if non-nil, zero value otherwise.

### GetRealMemoryOk

`func (o *V0044Node) GetRealMemoryOk() (*int64, bool)`

GetRealMemoryOk returns a tuple with the RealMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealMemory

`func (o *V0044Node) SetRealMemory(v int64)`

SetRealMemory sets RealMemory field to given value.

### HasRealMemory

`func (o *V0044Node) HasRealMemory() bool`

HasRealMemory returns a boolean if a field has been set.

### GetResCoresPerGpu

`func (o *V0044Node) GetResCoresPerGpu() int32`

GetResCoresPerGpu returns the ResCoresPerGpu field if non-nil, zero value otherwise.

### GetResCoresPerGpuOk

`func (o *V0044Node) GetResCoresPerGpuOk() (*int32, bool)`

GetResCoresPerGpuOk returns a tuple with the ResCoresPerGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResCoresPerGpu

`func (o *V0044Node) SetResCoresPerGpu(v int32)`

SetResCoresPerGpu sets ResCoresPerGpu field to given value.

### HasResCoresPerGpu

`func (o *V0044Node) HasResCoresPerGpu() bool`

HasResCoresPerGpu returns a boolean if a field has been set.

### GetComment

`func (o *V0044Node) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *V0044Node) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *V0044Node) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *V0044Node) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *V0044Node) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V0044Node) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V0044Node) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V0044Node) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonChangedAt

`func (o *V0044Node) GetReasonChangedAt() V0044Uint64NoValStruct`

GetReasonChangedAt returns the ReasonChangedAt field if non-nil, zero value otherwise.

### GetReasonChangedAtOk

`func (o *V0044Node) GetReasonChangedAtOk() (*V0044Uint64NoValStruct, bool)`

GetReasonChangedAtOk returns a tuple with the ReasonChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonChangedAt

`func (o *V0044Node) SetReasonChangedAt(v V0044Uint64NoValStruct)`

SetReasonChangedAt sets ReasonChangedAt field to given value.

### HasReasonChangedAt

`func (o *V0044Node) HasReasonChangedAt() bool`

HasReasonChangedAt returns a boolean if a field has been set.

### GetReasonSetByUser

`func (o *V0044Node) GetReasonSetByUser() string`

GetReasonSetByUser returns the ReasonSetByUser field if non-nil, zero value otherwise.

### GetReasonSetByUserOk

`func (o *V0044Node) GetReasonSetByUserOk() (*string, bool)`

GetReasonSetByUserOk returns a tuple with the ReasonSetByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonSetByUser

`func (o *V0044Node) SetReasonSetByUser(v string)`

SetReasonSetByUser sets ReasonSetByUser field to given value.

### HasReasonSetByUser

`func (o *V0044Node) HasReasonSetByUser() bool`

HasReasonSetByUser returns a boolean if a field has been set.

### GetResumeAfter

`func (o *V0044Node) GetResumeAfter() V0044Uint64NoValStruct`

GetResumeAfter returns the ResumeAfter field if non-nil, zero value otherwise.

### GetResumeAfterOk

`func (o *V0044Node) GetResumeAfterOk() (*V0044Uint64NoValStruct, bool)`

GetResumeAfterOk returns a tuple with the ResumeAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAfter

`func (o *V0044Node) SetResumeAfter(v V0044Uint64NoValStruct)`

SetResumeAfter sets ResumeAfter field to given value.

### HasResumeAfter

`func (o *V0044Node) HasResumeAfter() bool`

HasResumeAfter returns a boolean if a field has been set.

### GetReservation

`func (o *V0044Node) GetReservation() string`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *V0044Node) GetReservationOk() (*string, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *V0044Node) SetReservation(v string)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *V0044Node) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetAllocMemory

`func (o *V0044Node) GetAllocMemory() int64`

GetAllocMemory returns the AllocMemory field if non-nil, zero value otherwise.

### GetAllocMemoryOk

`func (o *V0044Node) GetAllocMemoryOk() (*int64, bool)`

GetAllocMemoryOk returns a tuple with the AllocMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocMemory

`func (o *V0044Node) SetAllocMemory(v int64)`

SetAllocMemory sets AllocMemory field to given value.

### HasAllocMemory

`func (o *V0044Node) HasAllocMemory() bool`

HasAllocMemory returns a boolean if a field has been set.

### GetAllocCpus

`func (o *V0044Node) GetAllocCpus() int32`

GetAllocCpus returns the AllocCpus field if non-nil, zero value otherwise.

### GetAllocCpusOk

`func (o *V0044Node) GetAllocCpusOk() (*int32, bool)`

GetAllocCpusOk returns a tuple with the AllocCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocCpus

`func (o *V0044Node) SetAllocCpus(v int32)`

SetAllocCpus sets AllocCpus field to given value.

### HasAllocCpus

`func (o *V0044Node) HasAllocCpus() bool`

HasAllocCpus returns a boolean if a field has been set.

### GetAllocIdleCpus

`func (o *V0044Node) GetAllocIdleCpus() int32`

GetAllocIdleCpus returns the AllocIdleCpus field if non-nil, zero value otherwise.

### GetAllocIdleCpusOk

`func (o *V0044Node) GetAllocIdleCpusOk() (*int32, bool)`

GetAllocIdleCpusOk returns a tuple with the AllocIdleCpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocIdleCpus

`func (o *V0044Node) SetAllocIdleCpus(v int32)`

SetAllocIdleCpus sets AllocIdleCpus field to given value.

### HasAllocIdleCpus

`func (o *V0044Node) HasAllocIdleCpus() bool`

HasAllocIdleCpus returns a boolean if a field has been set.

### GetTresUsed

`func (o *V0044Node) GetTresUsed() string`

GetTresUsed returns the TresUsed field if non-nil, zero value otherwise.

### GetTresUsedOk

`func (o *V0044Node) GetTresUsedOk() (*string, bool)`

GetTresUsedOk returns a tuple with the TresUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresUsed

`func (o *V0044Node) SetTresUsed(v string)`

SetTresUsed sets TresUsed field to given value.

### HasTresUsed

`func (o *V0044Node) HasTresUsed() bool`

HasTresUsed returns a boolean if a field has been set.

### GetTresWeighted

`func (o *V0044Node) GetTresWeighted() float64`

GetTresWeighted returns the TresWeighted field if non-nil, zero value otherwise.

### GetTresWeightedOk

`func (o *V0044Node) GetTresWeightedOk() (*float64, bool)`

GetTresWeightedOk returns a tuple with the TresWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTresWeighted

`func (o *V0044Node) SetTresWeighted(v float64)`

SetTresWeighted sets TresWeighted field to given value.

### HasTresWeighted

`func (o *V0044Node) HasTresWeighted() bool`

HasTresWeighted returns a boolean if a field has been set.

### GetSlurmdStartTime

`func (o *V0044Node) GetSlurmdStartTime() V0044Uint64NoValStruct`

GetSlurmdStartTime returns the SlurmdStartTime field if non-nil, zero value otherwise.

### GetSlurmdStartTimeOk

`func (o *V0044Node) GetSlurmdStartTimeOk() (*V0044Uint64NoValStruct, bool)`

GetSlurmdStartTimeOk returns a tuple with the SlurmdStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlurmdStartTime

`func (o *V0044Node) SetSlurmdStartTime(v V0044Uint64NoValStruct)`

SetSlurmdStartTime sets SlurmdStartTime field to given value.

### HasSlurmdStartTime

`func (o *V0044Node) HasSlurmdStartTime() bool`

HasSlurmdStartTime returns a boolean if a field has been set.

### GetSockets

`func (o *V0044Node) GetSockets() int32`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *V0044Node) GetSocketsOk() (*int32, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *V0044Node) SetSockets(v int32)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *V0044Node) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetThreads

`func (o *V0044Node) GetThreads() int32`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *V0044Node) GetThreadsOk() (*int32, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *V0044Node) SetThreads(v int32)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *V0044Node) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetTemporaryDisk

`func (o *V0044Node) GetTemporaryDisk() int32`

GetTemporaryDisk returns the TemporaryDisk field if non-nil, zero value otherwise.

### GetTemporaryDiskOk

`func (o *V0044Node) GetTemporaryDiskOk() (*int32, bool)`

GetTemporaryDiskOk returns a tuple with the TemporaryDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemporaryDisk

`func (o *V0044Node) SetTemporaryDisk(v int32)`

SetTemporaryDisk sets TemporaryDisk field to given value.

### HasTemporaryDisk

`func (o *V0044Node) HasTemporaryDisk() bool`

HasTemporaryDisk returns a boolean if a field has been set.

### GetWeight

`func (o *V0044Node) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *V0044Node) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *V0044Node) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *V0044Node) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetTopology

`func (o *V0044Node) GetTopology() string`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *V0044Node) GetTopologyOk() (*string, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *V0044Node) SetTopology(v string)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *V0044Node) HasTopology() bool`

HasTopology returns a boolean if a field has been set.

### GetTres

`func (o *V0044Node) GetTres() string`

GetTres returns the Tres field if non-nil, zero value otherwise.

### GetTresOk

`func (o *V0044Node) GetTresOk() (*string, bool)`

GetTresOk returns a tuple with the Tres field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTres

`func (o *V0044Node) SetTres(v string)`

SetTres sets Tres field to given value.

### HasTres

`func (o *V0044Node) HasTres() bool`

HasTres returns a boolean if a field has been set.

### GetVersion

`func (o *V0044Node) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *V0044Node) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *V0044Node) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *V0044Node) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


