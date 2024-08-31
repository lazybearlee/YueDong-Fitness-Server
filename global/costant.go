package global

const (
	DefaultRouter = "health"

	ConfigEnv         = "FITNESS_CONFIG" // used for find the config in the system environment variables
	ConfigDefaultFile = temp_root + "config.yaml"
	ConfigReleaseFile = temp_root + "config.release.yaml"
	ConfigTestFile    = temp_root + "config.test.yaml"

	CommonUser = 9870 // 普通用户
	AdminUser  = 7890 // 管理员用户
	AdminSuper = 789  // 超级管理员用户

	CommonUserStr = "9870" // 普通用户
	AdminUserStr  = "7890" // 管理员用户
	AdminSuperStr = "789"  // 超级管理员用户

	StepRankYesterday = "rank_yesterday"
	StepRankToday     = "rank_today"
	DistanceRankToday = "distance_today"
)

const temp_root = "D:\\RoadToCs\\YueDong-Fitness-Server\\"
