package types

import (
//"github.com/asyoume/paas_srv/pkg/utils"
)

func NewSystemLog() *SystemLog {
	log := SystemLog{
	//Id: utils.StrUUID(),
	}
	return &log
}

// 系统日志的类型
type SystemLog struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Msg   string `json:"type"`
	Level string `json:"level"`
	Time  int64  `json:"created"`
}

func (l *SystemLog) GetLevel() string {
	return l.Level
}

func (l *SystemLog) SetLevel(level string) {
	l.Level = level
}

func (l *SystemLog) SetTime(t int64) {
	l.Time = t
}
