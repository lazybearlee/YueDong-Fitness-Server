package core

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

// ZapLoggerInit get zap logger
func ZapLoggerInit() (logger *zap.Logger) {
	// 检查是否存在Director文件夹
	if ok := utils.DirectoryExists(global.FITNESS_CONFIG.Zap.Director); !ok {
		// 创建Director文件夹
		_ = utils.CreateDir(global.FITNESS_CONFIG.Zap.Director)
	}
	// 获取日志级别
	levels := global.FITNESS_CONFIG.Zap.Levels()
	// 根据日志级别个数创建zapcore.Core，然后统一新建
	cores := make([]zapcore.Core, 0, len(levels))
	for i := 0; i < len(levels); i++ {
		core := NewZapCore(levels[i])
		cores = append(cores, core)
	}
	logger = zap.New(zapcore.NewTee(cores...))
	// 是否显示行号
	if global.FITNESS_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	zap.ReplaceGlobals(logger)
	return logger
}

type ZapCore struct {
	level zapcore.Level
	zapcore.Core
}

func NewZapCore(level zapcore.Level) *ZapCore {
	entity := &ZapCore{level: level}
	syncer := entity.WriteSyncer()
	levelEnabler := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == level
	})
	entity.Core = zapcore.NewCore(global.FITNESS_CONFIG.Zap.Encoder(), syncer, levelEnabler)
	return entity
}

func (z *ZapCore) WriteSyncer(formats ...string) zapcore.WriteSyncer {
	// define a new cutter, which is a middleware for log
	cutter := utils.NewCutter(
		global.FITNESS_CONFIG.Zap.Director,
		z.level.String(),
		global.FITNESS_CONFIG.Zap.RetentionDay,
		utils.CutterWithLayout(time.DateOnly),
		utils.CutterWithFormats(formats...),
	)
	if global.FITNESS_CONFIG.Zap.LogInConsole {
		multiSyncer := zapcore.NewMultiWriteSyncer(os.Stdout, cutter)
		return zapcore.AddSync(multiSyncer)
	}
	return zapcore.AddSync(cutter)
}

func (z *ZapCore) Enabled(level zapcore.Level) bool {
	return z.level == level
}

func (z *ZapCore) With(fields []zapcore.Field) zapcore.Core {
	return z.Core.With(fields)
}

func (z *ZapCore) Check(entry zapcore.Entry, check *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	if z.Enabled(entry.Level) {
		return check.AddCore(entry, z)
	}
	return check
}

func (z *ZapCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	for i := 0; i < len(fields); i++ {
		if fields[i].Key == "business" || fields[i].Key == "folder" || fields[i].Key == "directory" {
			syncer := z.WriteSyncer(fields[i].String)
			z.Core = zapcore.NewCore(global.FITNESS_CONFIG.Zap.Encoder(), syncer, z.level)
		}
	}
	return z.Core.Write(entry, fields)
}

func (z *ZapCore) Sync() error {
	return z.Core.Sync()
}
