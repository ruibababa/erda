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

package reconciler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/modules/pipeline/pipengine/queue/throttler"
	"github.com/erda-project/erda/pkg/jsonstore"
)

func makeThrottlerBackupKey(name string) string {
	return fmt.Sprintf("/devops/pipeline/throttler/reconciler/%s", name)
}

// loadThrottler 从存储中加载 throttler
func (r *Reconciler) loadThrottler() error {
	// init throttler
	r.Throttler = throttler.NewNamedThrottler("default", nil)

	ctx := context.Background()
	var backup json.RawMessage
	if err := r.js.Get(ctx, makeThrottlerBackupKey(r.Throttler.Name()), &backup); err != nil {
		if err == jsonstore.NotFoundErr {
			return nil
		}
		return fmt.Errorf("reconciler: failed to load throttler from etcd, err: %v", err)
	}
	err := r.Throttler.Import(backup)
	if err == nil {
		return nil
	}
	// 加载失败可忽略，任务目前没有存队列信息，无法恢复，原来在队列中的任务 popPending 都会返回可弹出，不影响新的任务
	logrus.Warnf("reconciler: failed to load throttler, ignore, import err: %v", err)
	// load from database
	return nil
}

const logPrefixContinueBackupThrottler = "[throttler backup]"

// ContinueBackupThrottler 持续备份 throttler
func (r *Reconciler) ContinueBackupThrottler() {
	done := make(chan struct{})
	errDone := make(chan error)

	var costTime time.Duration

	for {
		go func() {
			// 执行 loading
			begin := time.Now()
			backup := r.Throttler.Export()
			end := time.Now()
			costTime = end.Sub(begin)
			if err := r.js.Put(context.Background(), makeThrottlerBackupKey(r.Throttler.Name()), backup); err != nil {
				errDone <- err
				return
			}
			done <- struct{}{}
		}()

		select {
		// 正常结束，等待 1min 后开始下一次处理
		case <-done:
			logrus.Debugf("%s: sleep 30s for next backup (cost %s this time)", logPrefixContinueBackupThrottler, costTime)
			time.Sleep(time.Second * 30)

		// 异常结束，等待 10s 后尽快开始下一次处理
		case err := <-errDone:
			logrus.Errorf("%s: failed to load, wait 10s for next loading, err: %v", logPrefixContinueBackupThrottler, err)
			time.Sleep(time.Second * 10)
		}
	}
}
