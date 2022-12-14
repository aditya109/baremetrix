package models

// Config is central container for configuration object.
type Config struct {
	Instance Instance `json:"instance"`
}

// Instance contains instance-specific configuration.
type Instance struct {
	Env   Env   `json:"environment"` // Contains the environment the project is running in.
	Specs Specs `json:"specs"`       // Contains the project-specific configurations.
}

type Env struct {
	Type         string `json:"type"`          // Type specifies the type of environment.
	InfraDetails string `json:"infra_details"` // InfraDetails is a short description of infrastructure of the environment on which the load testing is being.
}

// Specs contains project-specific configuration.
type Specs struct {
	LevelledLog        LogSpecs           `json:"levelled_log"`  // Specifies the logging-related configuration.
	SummarySpecs       SummarySpecs       `json:"summary"`       // Specifies the summary-related configuration.
	PlaySpecs          PlaySpecs          `json:"play"`          // Specifies the play files-related configuration.
	PrerequisiteSpecs  PrerequisiteSpecs  `json:"prerequisite"`  // Specifies the prerequisite call-related configuration.
	VisualizationSpecs VisualizationSpecs `json:"visualization"` // Specifies the visualization graph generation-related configuration.
	Flow               FlowSpecs          `json:"flow"`          //Specifies the flow-related configuration.
}

// LogSpecs contains logging-related configuration.
type LogSpecs struct {
	FileSpecs             []FileSpecs `json:"file_specs"`               // Specifies the file-related configuration.
	EnableLoggingToFile   bool        `json:"enable_logging_to_file"`   // Flag to enable logging to file.
	EnableLoggingToStdout bool        `json:"enable_logging_to_stdout"` // Flag to enable logging to stdout.
	EnableColors          bool        `json:"enable_colors"`            // Flag to enable color for stdout, for text formatter.
	EnableFullTimeStamp   bool        `json:"enable_full_timeStamp"`    // Flag to enable full timestamp.
	OutputFormatter       string      `json:"output_formatter"`         // Can support json and text
}

// FileSpecs contains file-structure related configuration.
type FileSpecs struct {
	Name               []string `json:"filename"`             // Specifies the names required for file creation.
	Extension          string   `json:"fileextension"`        // Specifies the extension required for file creation.
	ContainerDirectory string   `json:"container_directory"`  // Specifies the container directory
	ShouldRun          bool     `json:"should_run,omitempty"` // Specifies if the directory is to be considered for execution, optional
}

// SummarySpecs contains summary -related configuration.
type SummarySpecs struct {
	FileSpecs                     []FileSpecs `json:"file_specs"`                       // Specifies the file-related configuration.
	ExpectedLatencyInMilliseconds int         `json:"expected_latency_in_milliseconds"` // Specifies the expected latency for API latency in ideal scenario.
}

// PlaySpecs contains playfiles-related configuration.
type PlaySpecs struct {
	FileSpecs []FileSpecs `json:"file_specs"` // Specifies the file-related configuration.
}

// PrerequisiteSpecs contains prerequisite call-related configuration.
type PrerequisiteSpecs struct {
	HttpTimeoutInSeconds           int `json:"http_timeout_in_seconds"`            // Specifies the timeout required in the pre-requisite HTTP API calls.
	SupportedBatchSize             int `json:"batch_size"`                         // Specifies the batch size to be considered.
	ConditionalBatchBuffer         int `json:"conditional_batch_buffer"`           // Specifies the batch incremental buffer in case of batch# = 0.
	ConditionalBatchBufferForFlows int `json:"conditional_batch_buffer_for_flows"` // Specifies the batch incremental buffer for flows in case of batch# = 0.
}

// VisualizationSpecs contains visualization-related configuration.
type VisualizationSpecs struct {
	FileSpecs  []FileSpecs `json:"file_specs"`  // Specifies the file-related configuration.
	GraphTypes []GraphType `json:"graph_types"` // Specifies the graph-type configuration.
}

// GraphType contains details for graph which are to be generated at the end of each of playfile execution.
type GraphType struct {
	Name       string      `json:"name"` // Specifies the graph-type name.
	Title      string      `json:"title"`
	IsEnabled  bool        `json:"is_enabled"`   // Flag to enable graph generation.
	XAxisLabel string      `json:"x_axis_label"` // Specifies the X-Axis label.
	YAxisLabel string      `json:"y_axis_label"` // Specifies the Y-Axis label.
	InnerPlots []InnerPlot `json:"inner_plots"`  // Specifies a list of inner plots
}

// InnerPlot contains details about the legends of the graph.
type InnerPlot struct {
	Name  string `json:"name"`  // Specifies the name of the plot.
	Label string `json:"label"` // Specifies the label of the plot (same is used for legends as well.).
}

// FlowSpecs contains list of flows-supported by the system.
type FlowSpecs struct {
	FlowTypes []FlowType `json:"flow_types"` // Specifies the flow types configured within the application.
}

// FlowType contains the details of a particular flow supported by the system.
type FlowType struct {
	Name string   `json:"name"` // Specifies the name of the flow.
	Id   string   `json:"id"`   // Specifies the id of the flow.
	Flow []string `json:"flow"` // Specifies the list of api of the flow in the sequence.
}
