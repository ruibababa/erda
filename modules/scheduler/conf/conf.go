// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package conf

import (
	"github.com/erda-project/erda/pkg/envconf"
)

// Conf scheduler conf, use envconf to load configuration
type Conf struct {
	// Debug Control log level
	Debug bool `env:"DEBUG" default:"false"`
	// PoolSize goroutine pool size
	PoolSize int `env:"POOL_SIZE" default:"50"`
	// ListenAddr scheduler listening address , eg: ":9091"
	ListenAddr             string `env:"LISTEN_ADDR" default:":9091"`
	DefaultRuntimeExecutor string `env:"DEFAULT_RUNTIME_EXECUTOR" default:"MARATHON"`
	// TraceLogEnv shows the key of environment variable defined for tracing log
	TraceLogEnv string `env:"TRACELOGENV" default:"TERMINUS_DEFINE_TAG"`
	// PlaceHolderImage Image used to occupy the seat when disassembling the service deployment
	PlaceHolderImage string `env:"PLACEHOLDER_IMAGE" default:"registry.cn-hangzhou.aliyuncs.com/terminus/busybox"`

	KafkaBrokers        string `env:"BOOTSTRAP_SERVERS"`
	KafkaContainerTopic string `env:"CMDB_CONTAINER_TOPIC"`
	KafkaGroup          string `env:"CMDB_GROUP"`

	TerminalSecurity bool `env:"TERMINAL_SECURITY" default:"false"`
}

var cfg Conf

// Load environment variable
func Load() {
	envconf.MustLoad(&cfg)
}

// Debug return cfg.Debug
func Debug() bool {
	return cfg.Debug
}

// PoolSize return cfg.PoolSize
func PoolSize() int {
	return cfg.PoolSize
}

// ListenAddr return cfg.ListenAddr
func ListenAddr() string {
	return cfg.ListenAddr
}

// DefaultRuntimeExecutor return cfg.DefaultRuntimeExecutor
func DefaultRuntimeExecutor() string {
	return cfg.DefaultRuntimeExecutor
}

// TraceLogEnv return cfg.TraceLogEnv
func TraceLogEnv() string {
	return cfg.TraceLogEnv
}

// PlaceHolderImage return cfg.PlaceHolderImage
func PlaceHolderImage() string {
	return cfg.PlaceHolderImage
}

func KafkaBrokers() string {
	return cfg.KafkaBrokers
}
func KafkaContainerTopic() string {
	return cfg.KafkaContainerTopic
}
func KafkaGroup() string {
	return cfg.KafkaGroup
}

// TerminalSecurity return cfg.TerminalSecurity
func TerminalSecurity() bool {
	return cfg.TerminalSecurity
}
