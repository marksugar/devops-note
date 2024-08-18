package logger

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(mode string) (err error) {
	// writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	writeSyncer := getLogWriter(
		viper.GetString("log.filename"),
		viper.GetInt("log.max_size"),
		viper.GetInt("log.max_backups"),
		viper.GetInt("log.max_age"),
	)
	encoder := getEncoder()
	var levellog = new(zapcore.Level)
	err = levellog.UnmarshalText([]byte(viper.GetString("log.level")))
	if err != nil {
		return
	}

	var core zapcore.Core
	if mode == "dev" {
		// 日志打印到终端并且写入文件
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, levellog),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, levellog)
	}

	lg := zap.New(core, zap.AddCaller()) // zap.addcaller是添加调用信息
	// lg替换全局
	zap.ReplaceGlobals(lg)
	return
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		// 字段写入
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),                                 // 请求状态
			zap.String("method", c.Request.Method),                               // 请求类型
			zap.String("path", path),                                             // 请求路径
			zap.String("query", query),                                           // 请求参数
			zap.String("ip", c.ClientIP()),                                       // 请求ip
			zap.String("user-agent", c.Request.UserAgent()),                      // 请求浏览器
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()), // 请求错误
			zap.Duration("cost", cost),                                           // 请求时间
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func getEncoder() zapcore.Encoder {
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) // json格式
	// return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()) // 正常格式

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 可读的时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 日志写入文件并且切割压缩
func getLogWriter(name string, maxsize, maxbackups, maxage int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   name,
		MaxSize:    maxsize,    // 备份文件大小MB
		MaxBackups: maxbackups, // 备份数量
		MaxAge:     maxage,     // 备份天数
		Compress:   false,      // 备份是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}
