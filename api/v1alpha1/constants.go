package v1alpha1

const (
	Creating = "creating"
	Deleting = "deleting"
	Failed   = "failed"
	Created  = "created"
	Scaled  = "scaled"

	// kind of operator cr
	DBTypeClickHouse = "ClickHouseInstallation"
	DBTypePostgreSQL = "PostgreSQL"
	DBTypeMysql      = "MySQL"

	// type of the cluster application
	ClusterAppTypeClickHouse = "ClickHouse"
	ClusterAPPTypePostgreSQL = "PostgreSQL"
	ClusterAPPTypeMySQL      = "MySQL"

	// cluster status
	ClusterStatusUnknown = "unknown"
	// 状态更新异常
)
