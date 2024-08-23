package gc

type DBGCConfig struct {
	TableName     string // 表名
	ComparedField string // 比较字段
	ComparedValue string // 比较值
}

const (
	GCExecFmt = "DELETE FROM %s WHERE %s < ?"
)
