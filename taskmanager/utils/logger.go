package utils

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

// 日志输出目标常量
const (
	WriteConsole = "console" // 输出到控制台
	WriteFile    = "file"    // 输出到文件
	WriteBoth    = "both"    // 同时输出到控制台和文件
)

type ZapConfig struct {
	Prefix     string         `yaml:"prefix" mapstructure:"prefix"`
	TimeFormat string         `yaml:"timeFormat" mapstructure:"timeFormat"`
	Level      string         `yaml:"level" mapstructure:"level"`
	Caller     bool           `yaml:"caller" mapstructure:"caller"`
	StackTrace bool           `yaml:"stackTrace" mapstructure:"stackTrace"`
	Writer     string         `yaml:"writer" mapstructure:"writer"`
	Encode     string         `yaml:"encode" mapstructure:"encode"`
	LogFile    *LogFileConfig `yaml:"logFile" mapstructure:"logFile"`
}

type LogFileConfig struct {
	MaxSize  int      `yaml:"maxSize" mapstructure:"maxSize"`
	BackUps  int      `yaml:"backups" mapstructure:"backups"`
	Compress bool     `yaml:"compress" mapstructure:"compress"`
	Output   []string `yaml:"output" mapstructure:"output"`
	Errput   []string `yaml:"errput" mapstructure:"errput"`
}

func zapEncoder(config *ZapConfig) zapcore.Encoder {
	//用viper读取配置 ,  考虑做多个配置文件，还是一个配置文件塞多个配置。
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/") // 使用变量  //环境变量
	viper.AddConfigPath(".")      // 在工作目录下查找

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	//也就是说自定义的 zap编码
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "Time",
		LevelKey:      "Level",
		NameKey:       "Logger",
		CallerKey:     "Caller",
		MessageKey:    "Message",
		StacktraceKey: "StackTrace",
		LineEnding:    zapcore.DefaultLineEnding,
		FunctionKey:   zapcore.OmitKey,
	}
	// 自定义时间格式
	//encoderConfig.EncodeTime = CustomTimeFormatEncoder
	// 日志级别大写
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	// 秒级时间间隔
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	// 简短的调用者输出
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	// 完整的序列化logger名称
	encoderConfig.EncodeName = zapcore.FullNameEncoder

	// 最终的日志编码 json或者console

	switch config.Encode {
	case "json":
		{
			return zapcore.NewJSONEncoder(encoderConfig)
		}
	case "console":
		{
			return zapcore.NewConsoleEncoder(encoderConfig)
		}
	}
	// 默认console
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func zapWriteSyncer(cfg *ZapConfig) zapcore.WriteSyncer {
	syncers := make([]zapcore.WriteSyncer, 0, 2)
	// 如果开启了日志控制台输出，就加入控制台书写器
	if cfg.Writer == WriteBoth || cfg.Writer == WriteConsole {
		syncers = append(syncers, zapcore.AddSync(os.Stdout))
	}

	// 如果开启了日志文件存储，就根据文件路径切片加入书写器
	if cfg.Writer == WriteBoth || cfg.Writer == WriteFile {
		// 添加日志输出器
		for _, path := range cfg.LogFile.Output {
			logger := &lumberjack.Logger{
				Filename:   path,                 //文件路径
				MaxSize:    cfg.LogFile.MaxSize,  //分割文件的大小
				MaxBackups: cfg.LogFile.BackUps,  //备份次数
				Compress:   cfg.LogFile.Compress, // 是否压缩
				LocalTime:  true,                 //使用本地时间
			}
			syncers = append(syncers, zapcore.Lock(zapcore.AddSync(logger)))
		}
	}
	return zap.CombineWriteSyncers(syncers...)
}

func zapLevelEnabler(cfg *ZapConfig) zapcore.LevelEnabler {
	switch cfg.Level {
	case string(zapcore.DebugLevel):
		return zap.DebugLevel
	case string(zapcore.InfoLevel):
		return zap.InfoLevel
	case string(zapcore.ErrorLevel):
		return zap.ErrorLevel
	case string(zapcore.PanicLevel):
		return zap.PanicLevel
	case string(zapcore.FatalLevel):
		return zap.FatalLevel
	}
	// 默认Debug级别
	return zap.DebugLevel
}

func InitZap(config *ZapConfig) *zap.Logger {
	// 构建编码器
	encoder := zapEncoder(config)
	// 构建日志级别
	levelEnabler := zapLevelEnabler(config)
	// 最后获得Core和Options
	subCore, options := tee(config, encoder, levelEnabler)
	// 创建Logger
	return zap.New(subCore, options...)
}

// 将所有合并
func tee(cfg *ZapConfig, encoder zapcore.Encoder, levelEnabler zapcore.LevelEnabler) (core zapcore.Core, options []zap.Option) {
	sink := zapWriteSyncer(cfg)
	return zapcore.NewCore(encoder, sink, levelEnabler), buildOptions(cfg, levelEnabler)
}

// 构建Option
func buildOptions(cfg *ZapConfig, levelEnabler zapcore.LevelEnabler) (options []zap.Option) {
	if cfg.Caller {
		options = append(options, zap.AddCaller())
	}

	if cfg.StackTrace {
		options = append(options, zap.AddStacktrace(levelEnabler))
	}
	return
}
