// Code generated by the feature package; DO NOT EDIT.

package feature

var appMetrics = MakeBoolFlag(
	"App Metrics",
	"appMetrics",
	"Bucky, Monitoring Team",
	false,
	Permanent,
	true,
)

// AppMetrics - Send UI Telementry to Tools cluster - should always be false in OSS
func AppMetrics() BoolFlag {
	return appMetrics
}

var backendExample = MakeBoolFlag(
	"Backend Example",
	"backendExample",
	"Gavin Cabbage",
	false,
	Permanent,
	false,
)

// BackendExample - A permanent backend example boolean flag
func BackendExample() BoolFlag {
	return backendExample
}

var frontendExample = MakeIntFlag(
	"Frontend Example",
	"frontendExample",
	"Gavin Cabbage",
	42,
	Temporary,
	true,
)

// FrontendExample - A temporary frontend example integer flag
func FrontendExample() IntFlag {
	return frontendExample
}

var pushDownWindowAggregateMean = MakeBoolFlag(
	"Push Down Window Aggregate Mean",
	"pushDownWindowAggregateMean",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownWindowAggregateMean - Enable Mean variant of PushDownWindowAggregateRule and PushDownBareAggregateRule
func PushDownWindowAggregateMean() BoolFlag {
	return pushDownWindowAggregateMean
}

var groupWindowAggregateTranspose = MakeBoolFlag(
	"Group Window Aggregate Transpose",
	"groupWindowAggregateTranspose",
	"Query Team",
	false,
	Temporary,
	false,
)

// GroupWindowAggregateTranspose - Enables the GroupWindowAggregateTransposeRule for all enabled window aggregates
func GroupWindowAggregateTranspose() BoolFlag {
	return groupWindowAggregateTranspose
}

var newAuth = MakeBoolFlag(
	"New Auth Package",
	"newAuth",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewAuthPackage - Enables the refactored authorization api
func NewAuthPackage() BoolFlag {
	return newAuth
}

var newLabels = MakeBoolFlag(
	"New Label Package",
	"newLabels",
	"Alirie Gray",
	false,
	Temporary,
	false,
)

// NewLabelPackage - Enables the refactored labels api
func NewLabelPackage() BoolFlag {
	return newLabels
}

var memoryOptimizedFill = MakeBoolFlag(
	"Memory Optimized Fill",
	"memoryOptimizedFill",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedFill - Enable the memory optimized fill()
func MemoryOptimizedFill() BoolFlag {
	return memoryOptimizedFill
}

var memoryOptimizedSchemaMutation = MakeBoolFlag(
	"Memory Optimized Schema Mutation",
	"memoryOptimizedSchemaMutation",
	"Query Team",
	false,
	Temporary,
	false,
)

// MemoryOptimizedSchemaMutation - Enable the memory optimized schema mutation functions
func MemoryOptimizedSchemaMutation() BoolFlag {
	return memoryOptimizedSchemaMutation
}

var queryTracing = MakeBoolFlag(
	"Query Tracing",
	"queryTracing",
	"Query Team",
	false,
	Permanent,
	false,
)

// QueryTracing - Turn on query tracing for queries that are sampled
func QueryTracing() BoolFlag {
	return queryTracing
}

var simpleTaskOptionsExtraction = MakeBoolFlag(
	"Simple Task Options Extraction",
	"simpleTaskOptionsExtraction",
	"Brett Buddin",
	false,
	Temporary,
	false,
)

// SimpleTaskOptionsExtraction - Simplified task options extraction to avoid undefined functions when saving tasks
func SimpleTaskOptionsExtraction() BoolFlag {
	return simpleTaskOptionsExtraction
}

var mergeFiltersRule = MakeBoolFlag(
	"Merged Filters Rule",
	"mergeFiltersRule",
	"Query Team",
	false,
	Temporary,
	false,
)

// MergedFiltersRule - Create one filter combining multiple single return statements
func MergedFiltersRule() BoolFlag {
	return mergeFiltersRule
}

var bandPlotType = MakeBoolFlag(
	"Band Plot Type",
	"bandPlotType",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// BandPlotType - Enables the creation of a band plot in Dashboards
func BandPlotType() BoolFlag {
	return bandPlotType
}

var mosaicGraphType = MakeBoolFlag(
	"Mosaic Graph Type",
	"mosaicGraphType",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// MosaicGraphType - Enables the creation of a mosaic graph in Dashboards
func MosaicGraphType() BoolFlag {
	return mosaicGraphType
}

var notebooks = MakeBoolFlag(
	"Notebooks",
	"notebooks",
	"Monitoring Team",
	false,
	Temporary,
	true,
)

// Notebooks - Determine if the notebook feature's route and navbar icon are visible to the user
func Notebooks() BoolFlag {
	return notebooks
}

var pushDownGroupAggregateMinMax = MakeBoolFlag(
	"Push Down Group Aggregate Min Max",
	"pushDownGroupAggregateMinMax",
	"Query Team",
	false,
	Temporary,
	false,
)

// PushDownGroupAggregateMinMax - Enable the min and max variants of the PushDownGroupAggregate planner rule
func PushDownGroupAggregateMinMax() BoolFlag {
	return pushDownGroupAggregateMinMax
}

var injectLatestSuccessTime = MakeBoolFlag(
	"Inject Latest Success Time",
	"injectLatestSuccessTime",
	"Compute Team",
	false,
	Temporary,
	false,
)

// InjectLatestSuccessTime - Inject the latest successful task run timestamp into a Task query extern when executing.
func InjectLatestSuccessTime() BoolFlag {
	return injectLatestSuccessTime
}

var enforceOrgDashboardLimits = MakeBoolFlag(
	"Enforce Organization Dashboard Limits",
	"enforceOrgDashboardLimits",
	"Compute Team",
	false,
	Temporary,
	false,
)

// EnforceOrganizationDashboardLimits - Enforces the default limit params for the dashboards api when orgs are set
func EnforceOrganizationDashboardLimits() BoolFlag {
	return enforceOrgDashboardLimits
}

var all = []Flag{
	appMetrics,
	backendExample,
	frontendExample,
	pushDownWindowAggregateMean,
	groupWindowAggregateTranspose,
	newAuth,
	newLabels,
	memoryOptimizedFill,
	memoryOptimizedSchemaMutation,
	queryTracing,
	simpleTaskOptionsExtraction,
	mergeFiltersRule,
	bandPlotType,
	mosaicGraphType,
	notebooks,
	pushDownGroupAggregateMinMax,
	injectLatestSuccessTime,
	enforceOrgDashboardLimits,
}

var byKey = map[string]Flag{
	"appMetrics":                    appMetrics,
	"backendExample":                backendExample,
	"frontendExample":               frontendExample,
	"pushDownWindowAggregateMean":   pushDownWindowAggregateMean,
	"groupWindowAggregateTranspose": groupWindowAggregateTranspose,
	"newAuth":                       newAuth,
	"newLabels":                     newLabels,
	"memoryOptimizedFill":           memoryOptimizedFill,
	"memoryOptimizedSchemaMutation": memoryOptimizedSchemaMutation,
	"queryTracing":                  queryTracing,
	"simpleTaskOptionsExtraction":   simpleTaskOptionsExtraction,
	"mergeFiltersRule":              mergeFiltersRule,
	"bandPlotType":                  bandPlotType,
	"mosaicGraphType":               mosaicGraphType,
	"notebooks":                     notebooks,
	"pushDownGroupAggregateMinMax":  pushDownGroupAggregateMinMax,
	"injectLatestSuccessTime":       injectLatestSuccessTime,
	"enforceOrgDashboardLimits":     enforceOrgDashboardLimits,
}
