package initialize

import (
	"github.com/lazybearlee/yuedong-fitness/core/initialize/app"
	"github.com/lazybearlee/yuedong-fitness/core/initialize/system"
	"sort"
)

// Order 用于设置初始化顺序，避免出现依赖问题
const (
	SystemOrder   = 10
	InternalOrder = 30
)

// 初始化顺序
const (
	JwtBlacklistOrder = SystemOrder + 5
	CasbinOrder       = SystemOrder + 5 // 预留前4个位置给可能的系统其他表
	AuthorityOrder    = CasbinOrder + 1
	UserOrder         = AuthorityOrder + 1
	FileOrder         = UserOrder + 1

	PlanOrder         = InternalOrder + 5
	RecordOrder       = InternalOrder + 5
	HealthStatusOrder = InternalOrder + 5
	DeviceOrder       = InternalOrder + 5
	PlanStageOrder    = PlanOrder + 1
)

// TablesInitializer 定义初始化接口
type TablesInitializer interface {
	Name() string // 初始化器名称
	MigrateTable() (err error)
	InitializeData() (err error)
	TableCreated() bool
	DataInitialized() bool
}

type TablesInitializerWithOrder struct {
	order int
	TablesInitializer
}

type TablesInitializerSlice []*TablesInitializerWithOrder

/* -- sort.Interface -- */

func (tis TablesInitializerSlice) Len() int {
	return len(tis)
}

func (tis TablesInitializerSlice) Less(i, j int) bool {
	return tis[i].order < tis[j].order
}

func (tis TablesInitializerSlice) Swap(i, j int) {
	tis[i], tis[j] = tis[j], tis[i]
}

var (
	initializers TablesInitializerSlice                 // initializers 用于存储所有的初始化器
	cache        map[string]*TablesInitializerWithOrder // cache 用于存储初始化器
)

// RegisterInitializer 注册初始化器
func RegisterInitializer(order int, initializer TablesInitializer) {
	if initializers == nil {
		initializers = TablesInitializerSlice{}
	}
	if cache == nil {
		cache = map[string]*TablesInitializerWithOrder{}
	}
	name := initializer.Name()
	if _, existed := cache[name]; existed {
		panic("Name conflict on " + name)
	}
	ti := TablesInitializerWithOrder{order, initializer}
	initializers = append(initializers, &ti)
	cache[name] = &ti
}

// LoadInitializers 加载所有的初始化器
func LoadInitializers() {
	RegisterInitializer(CasbinOrder, &system.CasbinInitializer{})
	RegisterInitializer(AuthorityOrder, &system.AuthorityInitializer{})
	RegisterInitializer(UserOrder, &system.UserInitializer{})
	RegisterInitializer(JwtBlacklistOrder, &system.JwtBlacklistInitializer{})
	RegisterInitializer(FileOrder, &system.FileInitializer{})

	RegisterInitializer(PlanOrder, &app.ExercisePlanInitializer{})
	RegisterInitializer(RecordOrder, &app.ExerciseRecordInitializer{})
	RegisterInitializer(HealthStatusOrder, &app.HealthStatusInitializer{})
	RegisterInitializer(PlanStageOrder, &app.PlanStageInitializer{})

	// 给所有的初始化器排序
	sort.Sort(initializers)
}
