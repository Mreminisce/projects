package config

import (
	"github.com/lexkong/log"
)

func initLog() {
	logConfig := log.PassLagerCfg{
		// 输出位置：file 指定的日志文件、stdout 标准输出，也可以同时选择
		Writers: "file,stdout",
		// 日志级别：DEBUG、INFO、WARN、ERROR、FATAL
		LoggerLevel: "DEBUG",
		// 日志文件位置
		LoggerFile: logFilePath,
		// 日志的输出格式：true 输出成 JSON 格式，false 输出成 plaintext 格式
		LogFormatText: false,
		// rotate 依据：daily 根据天进行转存、size 根据大小进行转存
		RollingPolicy: "size",
		// rotate 转存时间，配合 rollingPolicy: daily 使用
		LogRotateDate: 1,
		// rotate 转存大小，配合 rollingPolicy: size 使用(大于 1mb 会压缩为 zip)
		LogRotateSize: 1,
		// 当日志文件达到转存标准时 log 系统会进行压缩备份，这里指定备份文件的最多个数
		LogBackupCount: 7,
	}
	log.InitWithConfig(&logConfig)
}
