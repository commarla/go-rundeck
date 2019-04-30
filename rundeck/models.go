package rundeck

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
    "io"
)

// The package's fully qualified name.
const fqdn = "rundeck"

        // FileState enumerates the values for file state.
    type FileState string

    const (
                // Deleted ...
        Deleted FileState = "deleted"
                // Expired ...
        Expired FileState = "expired"
                // Retained ...
        Retained FileState = "retained"
                // Temp ...
        Temp FileState = "temp"
            )
    // PossibleFileStateValues returns an array of possible values for the FileState const type.
    func PossibleFileStateValues() []FileState {
        return []FileState{Deleted,Expired,Retained,Temp}
    }

        // KeyType enumerates the values for key type.
    type KeyType string

    const (
                // Private ...
        Private KeyType = "private"
                // Public ...
        Public KeyType = "public"
            )
    // PossibleKeyTypeValues returns an array of possible values for the KeyType const type.
    func PossibleKeyTypeValues() []KeyType {
        return []KeyType{Private,Public}
    }

        // Loglevel enumerates the values for loglevel.
    type Loglevel string

    const (
                // DEBUG ...
        DEBUG Loglevel = "DEBUG"
                // ERROR ...
        ERROR Loglevel = "ERROR"
                // INFO ...
        INFO Loglevel = "INFO"
                // VERBOSE ...
        VERBOSE Loglevel = "VERBOSE"
                // WARN ...
        WARN Loglevel = "WARN"
            )
    // PossibleLoglevelValues returns an array of possible values for the Loglevel const type.
    func PossibleLoglevelValues() []Loglevel {
        return []Loglevel{DEBUG,ERROR,INFO,VERBOSE,WARN}
    }

        // Loglevel1 enumerates the values for loglevel 1.
    type Loglevel1 string

    const (
                // Loglevel1DEBUG ...
        Loglevel1DEBUG Loglevel1 = "DEBUG"
                // Loglevel1ERROR ...
        Loglevel1ERROR Loglevel1 = "ERROR"
                // Loglevel1INFO ...
        Loglevel1INFO Loglevel1 = "INFO"
                // Loglevel1VERBOSE ...
        Loglevel1VERBOSE Loglevel1 = "VERBOSE"
                // Loglevel1WARN ...
        Loglevel1WARN Loglevel1 = "WARN"
            )
    // PossibleLoglevel1Values returns an array of possible values for the Loglevel1 const type.
    func PossibleLoglevel1Values() []Loglevel1 {
        return []Loglevel1{Loglevel1DEBUG,Loglevel1ERROR,Loglevel1INFO,Loglevel1VERBOSE,Loglevel1WARN}
    }

        // Status enumerates the values for status.
    type Status string

    const (
                // Aborted ...
        Aborted Status = "aborted"
                // Failed ...
        Failed Status = "failed"
                // FailedWithRetry ...
        FailedWithRetry Status = "failed-with-retry"
                // Other ...
        Other Status = "other"
                // Running ...
        Running Status = "running"
                // Scheduled ...
        Scheduled Status = "scheduled"
                // Succeeded ...
        Succeeded Status = "succeeded"
                // Timedout ...
        Timedout Status = "timedout"
            )
    // PossibleStatusValues returns an array of possible values for the Status const type.
    func PossibleStatusValues() []Status {
        return []Status{Aborted,Failed,FailedWithRetry,Other,Running,Scheduled,Succeeded,Timedout}
    }

            // ACLList ...
            type ACLList struct {
            autorest.Response `json:"-"`
            Path *string `json:"path,omitempty"`
            Type *string `json:"type,omitempty"`
            Href *string `json:"href,omitempty"`
            Resources *[]ACLReference `json:"resources,omitempty"`
            }

            // ACLPolicyResponse ...
            type ACLPolicyResponse struct {
            autorest.Response `json:"-"`
            // Contents - Policy as JSON encoded YAML string
            Contents *string `json:"contents,omitempty"`
            }

            // ACLReference ...
            type ACLReference struct {
            Path *string `json:"path,omitempty"`
            Type *string `json:"type,omitempty"`
            Name *string `json:"name,omitempty"`
            Href *string `json:"href,omitempty"`
            }

            // BulkJobFailedInfo ...
            type BulkJobFailedInfo struct {
            ID *string `json:"id,omitempty"`
            ErrorCode *string `json:"errorCode,omitempty"`
            Message *string `json:"message,omitempty"`
            }

            // BulkJobSucceededInfo ...
            type BulkJobSucceededInfo struct {
            ID *string `json:"id,omitempty"`
            Message *string `json:"message,omitempty"`
            }

            // ExecuteJobRequest ...
            type ExecuteJobRequest struct {
            // ArgString - Argument string to pass to the job, of the form `-opt value -opt2 value ...`
            ArgString *string `json:"argString,omitempty"`
            // Loglevel - Loglevel to use. Possible values include: 'Loglevel1DEBUG', 'Loglevel1VERBOSE', 'Loglevel1INFO', 'Loglevel1WARN', 'Loglevel1ERROR'
            Loglevel Loglevel1 `json:"loglevel,omitempty"`
            // AsUser - A username identifying the user who ran the job. Requires `runAs` permission.
            AsUser *string `json:"asUser,omitempty"`
            // Filter - A node filter string.
            Filter *string `json:"filter,omitempty"`
            // RunAtTime - Specify a time to run the job (*API v18* or later). `ISO-8601` format with optional milliseconds.
            RunAtTime *string `json:"runAtTime,omitempty"`
            // Options - Option value for option named `OPTNAME`. If specified the `argString` value is ignored (*API v18* or later).
            Options interface{} `json:"options,omitempty"`
            }

            // Execution ...
            type Execution struct {
            autorest.Response `json:"-"`
            ID *float64 `json:"id,omitempty"`
            Href *string `json:"href,omitempty"`
            Permalink *string `json:"permalink,omitempty"`
            // Status - Possible values include: 'Running', 'Succeeded', 'Failed', 'Aborted', 'Timedout', 'FailedWithRetry', 'Scheduled', 'Other'
            Status Status `json:"status,omitempty"`
            CustomStatus *string `json:"customStatus,omitempty"`
            Project *string `json:"project,omitempty"`
            User *string `json:"user,omitempty"`
            ServerUUID *string `json:"serverUUID,omitempty"`
            DateStarted *ExecutionDateStarted `json:"date-started,omitempty"`
            Job *JobMetadata `json:"job,omitempty"`
            Description *string `json:"description,omitempty"`
            Argstring *string `json:"argstring,omitempty"`
            SuccessfulNodes *[]string `json:"successfulNodes,omitempty"`
            }

            // ExecutionBulkDeleteRequest ...
            type ExecutionBulkDeleteRequest struct {
            Ids *[]string `json:"ids,omitempty"`
            }

            // ExecutionDateStarted ...
            type ExecutionDateStarted struct {
            Unixtime *float64 `json:"unixtime,omitempty"`
            Date *string `json:"date,omitempty"`
            }

            // ExecutionInputFilesListOKResponse ...
            type ExecutionInputFilesListOKResponse struct {
            autorest.Response `json:"-"`
            Files *[]JobInputFileInfo `json:"files,omitempty"`
            }

            // ExecutionList ...
            type ExecutionList struct {
            autorest.Response `json:"-"`
            Paging *Paging `json:"paging,omitempty"`
            Executions *[]Execution `json:"executions,omitempty"`
            }

            // IncompleteLogExecution ...
            type IncompleteLogExecution struct {
            ID *string `json:"id,omitempty"`
            Project *string `json:"project,omitempty"`
            Href *string `json:"href,omitempty"`
            Permalink *string `json:"permalink,omitempty"`
            Storage *IncompleteLogExecutionStorage `json:"storage,omitempty"`
            Errors *[]string `json:"errors,omitempty"`
            }

            // IncompleteLogExecutions ...
            type IncompleteLogExecutions struct {
            autorest.Response `json:"-"`
            Total *float64 `json:"total,omitempty"`
            Max *float64 `json:"max,omitempty"`
            Offset *float64 `json:"offset,omitempty"`
            Executions *[]IncompleteLogExecution `json:"executions,omitempty"`
            }

            // IncompleteLogExecutionStorage ...
            type IncompleteLogExecutionStorage struct {
            LocalFilesPresent *bool `json:"localFilesPresent,omitempty"`
            IncompleteFiletypes *[]string `json:"incompleteFiletypes,omitempty"`
            Queued *bool `json:"queued,omitempty"`
            Failed *bool `json:"failed,omitempty"`
            Date *string `json:"date,omitempty"`
            }

            // InvalidACLPolicyResponse ...
            type InvalidACLPolicyResponse struct {
            Valid *bool `json:"valid,omitempty"`
            Policies *[]InvalidACLPolicyResponsePoliciesItem `json:"policies,omitempty"`
            }

            // InvalidACLPolicyResponsePoliciesItem ...
            type InvalidACLPolicyResponsePoliciesItem struct {
            Policy *string `json:"policy,omitempty"`
            Errors *[]string `json:"errors,omitempty"`
            }

            // Job ...
            type Job struct {
            ID *string `json:"id,omitempty"`
            Name *string `json:"name,omitempty"`
            Group *string `json:"group,omitempty"`
            Project *string `json:"project,omitempty"`
            Description *string `json:"description,omitempty"`
            Href *string `json:"href,omitempty"`
            Permalink *string `json:"permalink,omitempty"`
            Scheduled *bool `json:"scheduled,omitempty"`
            ScheduleEnabled *bool `json:"scheduleEnabled,omitempty"`
            ServerNodeUUID *string `json:"serverNodeUUID,omitempty"`
            ServerOwner *string `json:"serverOwner,omitempty"`
            Enabled *bool `json:"enabled,omitempty"`
            }

            // JobBulkDeleteRequest ...
            type JobBulkDeleteRequest struct {
            Ids *[]string `json:"ids,omitempty"`
            }

            // JobBulkOperationResponse ...
            type JobBulkOperationResponse struct {
            autorest.Response `json:"-"`
            // RequestCount - The number of job IDs that were in the delete request.
            RequestCount *float64 `json:"requestCount,omitempty"`
            Allsuccessful *bool `json:"allsuccessful,omitempty"`
            Succeeded *[]BulkJobSucceededInfo `json:"succeeded,omitempty"`
            Failed *[]BulkJobFailedInfo `json:"failed,omitempty"`
            }

            // JobExecutionBulkDisableRequest ...
            type JobExecutionBulkDisableRequest struct {
            Ids *[]string `json:"ids,omitempty"`
            }

            // JobExecutionBulkEnableRequest ...
            type JobExecutionBulkEnableRequest struct {
            Ids *[]string `json:"ids,omitempty"`
            }

            // JobExecutionDelete ...
            type JobExecutionDelete struct {
            autorest.Response `json:"-"`
            FailedCount *float64 `json:"failedCount,omitempty"`
            SuccessCount *float64 `json:"successCount,omitempty"`
            Allsuccessful *bool `json:"allsuccessful,omitempty"`
            RequestCount *float64 `json:"requestCount,omitempty"`
            Failures *[]JobExecutionDeleteFailuresItem `json:"failures,omitempty"`
            }

            // JobExecutionDeleteFailuresItem ...
            type JobExecutionDeleteFailuresItem struct {
            ID *string `json:"id,omitempty"`
            Message *string `json:"message,omitempty"`
            }

            // JobExecutionDisableOKResponse ...
            type JobExecutionDisableOKResponse struct {
            autorest.Response `json:"-"`
            Success *bool `json:"success,omitempty"`
            }

            // JobExecutionEnableOKResponse ...
            type JobExecutionEnableOKResponse struct {
            autorest.Response `json:"-"`
            Success *bool `json:"success,omitempty"`
            }

            // JobGetOKResponse ...
            type JobGetOKResponse struct {
            autorest.Response `json:"-"`
            Content *string `json:"content,omitempty"`
            }

            // JobInputFileInfo ...
            type JobInputFileInfo struct {
            autorest.Response `json:"-"`
            ID *string `json:"id,omitempty"`
            User *string `json:"user,omitempty"`
            // FileState - Possible values include: 'Temp', 'Deleted', 'Expired', 'Retained'
            FileState FileState `json:"fileState,omitempty"`
            Sha *string `json:"sha,omitempty"`
            JobID *string `json:"jobId,omitempty"`
            DateCreated *string `json:"dateCreated,omitempty"`
            ServerNodeUUID *string `json:"serverNodeUUID,omitempty"`
            FileName *string `json:"fileName,omitempty"`
            Size *int32 `json:"size,omitempty"`
            ExpirationDate *string `json:"expirationDate,omitempty"`
            ExecID *string `json:"execId,omitempty"`
            }

            // JobInputFileListResponse ...
            type JobInputFileListResponse struct {
            autorest.Response `json:"-"`
            Paging *Paging `json:"paging,omitempty"`
            Files *[]JobInputFileInfo `json:"files,omitempty"`
            }

            // JobMetadata ...
            type JobMetadata struct {
            autorest.Response `json:"-"`
            ID *string `json:"id,omitempty"`
            Name *string `json:"name,omitempty"`
            Group *string `json:"group,omitempty"`
            Project *string `json:"project,omitempty"`
            Description *string `json:"description,omitempty"`
            Href *string `json:"href,omitempty"`
            Permalink *string `json:"permalink,omitempty"`
            Scheduled *bool `json:"scheduled,omitempty"`
            ScheduleEnabled *bool `json:"scheduleEnabled,omitempty"`
            AverageDuration *float64 `json:"averageDuration,omitempty"`
            Options interface{} `json:"options,omitempty"`
            }

            // JobReference ...
            type JobReference struct {
            Href *string `json:"href,omitempty"`
            Permalink *string `json:"permalink,omitempty"`
            ID *string `json:"id,omitempty"`
            PerviousOwner *string `json:"pervious-owner,omitempty"`
            }

            // JobScheduleBulkDisableRequest ...
            type JobScheduleBulkDisableRequest struct {
            Ids *[]string `json:"ids,omitempty"`
            }

            // JobScheduleBulkEnableRequest ...
            type JobScheduleBulkEnableRequest struct {
            Ids *[]string `json:"ids,omitempty"`
            }

            // JobScheduleDisableOKResponse ...
            type JobScheduleDisableOKResponse struct {
            autorest.Response `json:"-"`
            Success *bool `json:"success,omitempty"`
            }

            // JobScheduleEnableOKResponse ...
            type JobScheduleEnableOKResponse struct {
            autorest.Response `json:"-"`
            Success *bool `json:"success,omitempty"`
            }

            // ListJob ...
            type ListJob struct {
            autorest.Response `json:"-"`
            Value *[]Job `json:"value,omitempty"`
            }

            // ListProjectListOKResponseItem ...
            type ListProjectListOKResponseItem struct {
            autorest.Response `json:"-"`
            Value *[]ProjectListOKResponseItem `json:"value,omitempty"`
            }

            // ListUser ...
            type ListUser struct {
            autorest.Response `json:"-"`
            Value *[]User `json:"value,omitempty"`
            }

            // LogStorage ...
            type LogStorage struct {
            autorest.Response `json:"-"`
            // Enabled - True if plugin is configured
            Enabled *bool `json:"enabled,omitempty"`
            // PluginName - Name of the configured plugin
            PluginName *string `json:"pluginName,omitempty"`
            // SucceededCount - Number of successful storage requests
            SucceededCount *float64 `json:"succeededCount,omitempty"`
            // FailedCount - Number of failed storage requests
            FailedCount *float64 `json:"failedCount,omitempty"`
            // QueuedCount - Number of queued storage requests
            QueuedCount *float64 `json:"queuedCount,omitempty"`
            // TotalCount - Total number of storage requests (currently queued plus previously processed)
            TotalCount *float64 `json:"totalCount,omitempty"`
            // IncompleteCount - Number of storage requests which have not completed successfully
            IncompleteCount *float64 `json:"incompleteCount,omitempty"`
            // MissingCount - Number of executions for this cluster node which have no associated storage requests
            MissingCount *float64 `json:"missingCount,omitempty"`
            }

            // MetricListOKResponse ...
            type MetricListOKResponse struct {
            autorest.Response `json:"-"`
            Links *MetricListOKResponseLinks `json:"_links,omitempty"`
            }

            // MetricListOKResponseLinks ...
            type MetricListOKResponseLinks struct {
            Metrics *MetricListOKResponseLinksMetrics `json:"metrics,omitempty"`
            Ping *MetricListOKResponseLinksPing `json:"ping,omitempty"`
            Threads *MetricListOKResponseLinksThreads `json:"threads,omitempty"`
            Healthcheck *MetricListOKResponseLinksHealthcheck `json:"healthcheck,omitempty"`
            }

            // MetricListOKResponseLinksHealthcheck ...
            type MetricListOKResponseLinksHealthcheck struct {
            Href *string `json:"href,omitempty"`
            }

            // MetricListOKResponseLinksMetrics ...
            type MetricListOKResponseLinksMetrics struct {
            Href *string `json:"href,omitempty"`
            }

            // MetricListOKResponseLinksPing ...
            type MetricListOKResponseLinksPing struct {
            Href *string `json:"href,omitempty"`
            }

            // MetricListOKResponseLinksThreads ...
            type MetricListOKResponseLinksThreads struct {
            Href *string `json:"href,omitempty"`
            }

            // ModifyUserRequest ...
            type ModifyUserRequest struct {
            FirstName *string `json:"firstName,omitempty"`
            LastName *string `json:"lastName,omitempty"`
            Email *string `json:"email,omitempty"`
            }

            // MotdUpdateRequest ...
            type MotdUpdateRequest struct {
            Contents *string `json:"contents,omitempty"`
            }

            // Paging ...
            type Paging struct {
            Count *int32 `json:"count,omitempty"`
            Total *int32 `json:"total,omitempty"`
            Offset *int32 `json:"offset,omitempty"`
            Max *int32 `json:"max,omitempty"`
            }

            // Project ...
            type Project struct {
            autorest.Response `json:"-"`
            Description *string `json:"description,omitempty"`
            Name *string `json:"name,omitempty"`
            URL *string `json:"url,omitempty"`
            Config interface{} `json:"config,omitempty"`
            }

            // ProjectConfigKeyGetOKResponse ...
            type ProjectConfigKeyGetOKResponse struct {
            autorest.Response `json:"-"`
            Key *string `json:"key,omitempty"`
            Value *string `json:"value,omitempty"`
            }

            // ProjectConfigKeySetOKResponse ...
            type ProjectConfigKeySetOKResponse struct {
            autorest.Response `json:"-"`
            Key *string `json:"key,omitempty"`
            Value *string `json:"value,omitempty"`
            }

            // ProjectConfigKeySetRequest ...
            type ProjectConfigKeySetRequest struct {
            Value *string `json:"value,omitempty"`
            }

            // ProjectCreateRequest ...
            type ProjectCreateRequest struct {
            Name *string `json:"name,omitempty"`
            Config interface{} `json:"config,omitempty"`
            }

            // ProjectListOKResponseItem ...
            type ProjectListOKResponseItem struct {
            Name *string `json:"name,omitempty"`
            Description *string `json:"description,omitempty"`
            URL *string `json:"url,omitempty"`
            }

            // ProjectMotdGetOKResponse ...
            type ProjectMotdGetOKResponse struct {
            autorest.Response `json:"-"`
            Contents *string `json:"contents,omitempty"`
            }

            // ProjectReadmeGetOKResponse ...
            type ProjectReadmeGetOKResponse struct {
            autorest.Response `json:"-"`
            Contents *string `json:"contents,omitempty"`
            }

            // ReadCloser ...
            type ReadCloser struct {
            autorest.Response `json:"-"`
            Value *io.ReadCloser `json:"value,omitempty"`
            }

            // ReadmeUpdateRequest ...
            type ReadmeUpdateRequest struct {
            Contents *string `json:"contents,omitempty"`
            }

            // RetryExecutionRequest ...
            type RetryExecutionRequest struct {
            // ArgString - Argument string to pass to the job, of the form `-opt value -opt2 value ...`
            ArgString *string `json:"argString,omitempty"`
            // Loglevel - Loglevel to use. Possible values include: 'DEBUG', 'VERBOSE', 'INFO', 'WARN', 'ERROR'
            Loglevel Loglevel `json:"loglevel,omitempty"`
            // AsUser - A username identifying the user who ran the job. Requires `runAs` permission.
            AsUser *string `json:"asUser,omitempty"`
            // Options - Option value for option named `OPTNAME`. If specified the `argString` value is ignored (*API v18* or later).
            Options interface{} `json:"options,omitempty"`
            }

            // SchedulerTakeoverRequest ...
            type SchedulerTakeoverRequest struct {
            Server *SchedulerTakeoverRequestServer `json:"server,omitempty"`
            Project *string `json:"project,omitempty"`
            Job *SchedulerTakeoverRequestJob `json:"job,omitempty"`
            }

            // SchedulerTakeoverRequestJob ...
            type SchedulerTakeoverRequestJob struct {
            ID *string `json:"id,omitempty"`
            }

            // SchedulerTakeoverRequestServer ...
            type SchedulerTakeoverRequestServer struct {
            UUID *string `json:"uuid,omitempty"`
            All *bool `json:"all,omitempty"`
            }

            // SetObject ...
            type SetObject struct {
            autorest.Response `json:"-"`
            Value interface{} `json:"value,omitempty"`
            }

            // StorageKeyListResponse ...
            type StorageKeyListResponse struct {
            autorest.Response `json:"-"`
            Resources *[]StorageKeyMetadata `json:"resources,omitempty"`
            URL *string `json:"url,omitempty"`
            Type *string `json:"type,omitempty"`
            Path *string `json:"path,omitempty"`
            }

            // StorageKeyMetadata ...
            type StorageKeyMetadata struct {
            Meta *StorageKeyMetadataMeta `json:"meta,omitempty"`
            URL *string `json:"url,omitempty"`
            Name *string `json:"name,omitempty"`
            Type *string `json:"type,omitempty"`
            Path *string `json:"path,omitempty"`
            }

            // StorageKeyMetadataMeta ...
            type StorageKeyMetadataMeta struct {
            // RundeckKeyType - Possible values include: 'Private', 'Public'
            RundeckKeyType KeyType `json:"Rundeck-key-type,omitempty"`
            RundeckContentMask *string `json:"Rundeck-content-mask,omitempty"`
            RundeckContentSize *string `json:"Rundeck-content-size,omitempty"`
            RundeckContentType *string `json:"Rundeck-content-type,omitempty"`
            }

            // String ...
            type String struct {
            autorest.Response `json:"-"`
            Value *string `json:"value,omitempty"`
            }

            // SystemACLPolicyCreateRequest ...
            type SystemACLPolicyCreateRequest struct {
            Contents *string `json:"contents,omitempty"`
            }

            // SystemACLPolicyUpdateRequest ...
            type SystemACLPolicyUpdateRequest struct {
            Contents *string `json:"contents,omitempty"`
            }

            // SystemExecutionsDisableOKResponse ...
            type SystemExecutionsDisableOKResponse struct {
            autorest.Response `json:"-"`
            ExecutionMode *string `json:"executionMode,omitempty"`
            }

            // SystemExecutionsEnableOKResponse ...
            type SystemExecutionsEnableOKResponse struct {
            autorest.Response `json:"-"`
            ExecutionMode *string `json:"executionMode,omitempty"`
            }

            // SystemIncompleteLogStorageExecutionsResumeOKResponse ...
            type SystemIncompleteLogStorageExecutionsResumeOKResponse struct {
            autorest.Response `json:"-"`
            Resumed *bool `json:"resumed,omitempty"`
            }

            // SystemInfo ...
            type SystemInfo struct {
            autorest.Response `json:"-"`
            System *SystemInfoSystem `json:"system,omitempty"`
            }

            // SystemInfoSystem ...
            type SystemInfoSystem struct {
            Timestamp *SystemInfoSystemTimestamp `json:"timestamp,omitempty"`
            RundeckProperty *SystemInfoSystemRundeck `json:"rundeck,omitempty"`
            Executions *SystemInfoSystemExecutions `json:"executions,omitempty"`
            Os *SystemInfoSystemOs `json:"os,omitempty"`
            Jvm *SystemInfoSystemJvm `json:"jvm,omitempty"`
            Stats *SystemInfoSystemStats `json:"stats,omitempty"`
            Metrics *SystemInfoSystemMetrics `json:"metrics,omitempty"`
            ThreadDump *SystemInfoSystemThreadDump `json:"threadDump,omitempty"`
            }

            // SystemInfoSystemExecutions ...
            type SystemInfoSystemExecutions struct {
            Active *bool `json:"active,omitempty"`
            ExecutionMode *string `json:"executionMode,omitempty"`
            }

            // SystemInfoSystemJvm ...
            type SystemInfoSystemJvm struct {
            Name *string `json:"name,omitempty"`
            Vendor *string `json:"vendor,omitempty"`
            Version *string `json:"version,omitempty"`
            ImplementationVersion *string `json:"implementationVersion,omitempty"`
            }

            // SystemInfoSystemMetrics ...
            type SystemInfoSystemMetrics struct {
            Href *string `json:"href,omitempty"`
            ContentType *string `json:"contentType,omitempty"`
            }

            // SystemInfoSystemOs ...
            type SystemInfoSystemOs struct {
            Arch *string `json:"arch,omitempty"`
            Name *string `json:"name,omitempty"`
            Version *string `json:"version,omitempty"`
            }

            // SystemInfoSystemRundeck ...
            type SystemInfoSystemRundeck struct {
            Version *string `json:"version,omitempty"`
            Build *string `json:"build,omitempty"`
            Node *string `json:"node,omitempty"`
            Base *string `json:"base,omitempty"`
            Apiversion *float64 `json:"apiversion,omitempty"`
            ServerUUID *string `json:"serverUUID,omitempty"`
            }

            // SystemInfoSystemStats ...
            type SystemInfoSystemStats struct {
            Uptime *SystemInfoSystemStatsUptime `json:"uptime,omitempty"`
            CPU *SystemInfoSystemStatsCPU `json:"cpu,omitempty"`
            Memory *SystemInfoSystemStatsMemory `json:"memory,omitempty"`
            Scheduler *SystemInfoSystemStatsScheduler `json:"scheduler,omitempty"`
            Threads *SystemInfoSystemStatsThreads `json:"threads,omitempty"`
            }

            // SystemInfoSystemStatsCPU ...
            type SystemInfoSystemStatsCPU struct {
            LoadAverage *SystemInfoSystemStatsCPULoadAverage `json:"loadAverage,omitempty"`
            Processors *float64 `json:"processors,omitempty"`
            }

            // SystemInfoSystemStatsCPULoadAverage ...
            type SystemInfoSystemStatsCPULoadAverage struct {
            Unit *string `json:"unit,omitempty"`
            Average *float64 `json:"average,omitempty"`
            }

            // SystemInfoSystemStatsMemory ...
            type SystemInfoSystemStatsMemory struct {
            Unit *string `json:"unit,omitempty"`
            Max *float64 `json:"max,omitempty"`
            Free *float64 `json:"free,omitempty"`
            Total *float64 `json:"total,omitempty"`
            }

            // SystemInfoSystemStatsScheduler ...
            type SystemInfoSystemStatsScheduler struct {
            Running *float64 `json:"running,omitempty"`
            ThreadPoolSize *float64 `json:"threadPoolSize,omitempty"`
            }

            // SystemInfoSystemStatsThreads ...
            type SystemInfoSystemStatsThreads struct {
            Active *float64 `json:"active,omitempty"`
            }

            // SystemInfoSystemStatsUptime ...
            type SystemInfoSystemStatsUptime struct {
            Duration *float64 `json:"duration,omitempty"`
            Unit *string `json:"unit,omitempty"`
            Since *SystemInfoSystemStatsUptimeSince `json:"since,omitempty"`
            }

            // SystemInfoSystemStatsUptimeSince ...
            type SystemInfoSystemStatsUptimeSince struct {
            Epoch *float64 `json:"epoch,omitempty"`
            Unit *string `json:"unit,omitempty"`
            Datetime *string `json:"datetime,omitempty"`
            }

            // SystemInfoSystemThreadDump ...
            type SystemInfoSystemThreadDump struct {
            Href *string `json:"href,omitempty"`
            ContentType *string `json:"contentType,omitempty"`
            }

            // SystemInfoSystemTimestamp ...
            type SystemInfoSystemTimestamp struct {
            Epoch *float64 `json:"epoch,omitempty"`
            Unit *string `json:"unit,omitempty"`
            Datetime *string `json:"datetime,omitempty"`
            }

            // TakeoverScheduleRequest ...
            type TakeoverScheduleRequest struct {
            Server *TakeoverScheduleRequestServer `json:"server,omitempty"`
            Project *string `json:"project,omitempty"`
            Job *TakeoverScheduleRequestJob `json:"job,omitempty"`
            }

            // TakeoverScheduleRequestJob ...
            type TakeoverScheduleRequestJob struct {
            ID *string `json:"id,omitempty"`
            }

            // TakeoverScheduleRequestServer ...
            type TakeoverScheduleRequestServer struct {
            UUID *string `json:"uuid,omitempty"`
            All *bool `json:"all,omitempty"`
            }

            // TakeoverScheduleResponse ...
            type TakeoverScheduleResponse struct {
            autorest.Response `json:"-"`
            TakeoverSchedule *TakeoverScheduleResponseTakeoverSchedule `json:"takeoverSchedule,omitempty"`
            Self *TakeoverScheduleResponseSelf `json:"self,omitempty"`
            Message *string `json:"message,omitempty"`
            Apiversion *float64 `json:"apiversion,omitempty"`
            Success *bool `json:"success,omitempty"`
            }

            // TakeoverScheduleResponseSelf ...
            type TakeoverScheduleResponseSelf struct {
            Server *TakeoverScheduleResponseSelfServer `json:"server,omitempty"`
            }

            // TakeoverScheduleResponseSelfServer ...
            type TakeoverScheduleResponseSelfServer struct {
            UUID *string `json:"uuid,omitempty"`
            }

            // TakeoverScheduleResponseTakeoverSchedule ...
            type TakeoverScheduleResponseTakeoverSchedule struct {
            Project *string `json:"project,omitempty"`
            Jobs *TakeoverScheduleResponseTakeoverScheduleJobs `json:"jobs,omitempty"`
            Server *TakeoverScheduleResponseTakeoverScheduleServer `json:"server,omitempty"`
            }

            // TakeoverScheduleResponseTakeoverScheduleJobs ...
            type TakeoverScheduleResponseTakeoverScheduleJobs struct {
            Failed *[]JobReference `json:"failed,omitempty"`
            Successfull *[]JobReference `json:"successfull,omitempty"`
            Total *float64 `json:"total,omitempty"`
            }

            // TakeoverScheduleResponseTakeoverScheduleServer ...
            type TakeoverScheduleResponseTakeoverScheduleServer struct {
            UUID *string `json:"uuid,omitempty"`
            All *bool `json:"all,omitempty"`
            }

            // User ...
            type User struct {
            autorest.Response `json:"-"`
            Login *string `json:"login,omitempty"`
            FirstName *string `json:"firstName,omitempty"`
            LastName *string `json:"lastName,omitempty"`
            Email *string `json:"email,omitempty"`
            }

            // UserRoleListOKResponse ...
            type UserRoleListOKResponse struct {
            autorest.Response `json:"-"`
            Roles *[]string `json:"roles,omitempty"`
            }

