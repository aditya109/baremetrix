package models

// SummarySpecificDirectives contains flags and other relative informations.
type SummarySpecificDirectives struct {
	Tenant                  string
	PlayName                string
	FileSpecs               FileSpecs
	ShouldUseDirectives     bool
	TimeStamp               string
	Iteration               string
	ShouldUseGraphIndicator bool
	GraphType               string
}
