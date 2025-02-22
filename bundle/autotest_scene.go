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

package bundle

import (
	"fmt"
	"strconv"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle/apierrors"
	"github.com/erda-project/erda/pkg/httputil"
)

func (b *Bundle) ExecuteDiceAutotestSceneStep(req apistructs.AutotestExecuteSceneStepRequest) (*apistructs.AutotestExecuteSceneStepResp, error) {
	host, err := b.urls.QA()
	if err != nil {
		return nil, err
	}
	hc := b.hc
	var rsp apistructs.AutotestExecuteSceneStepResp
	httpResp, err := hc.Post(host).Path(fmt.Sprintf("/api/autotests/scenes-step/%v/actions/execute", req.SceneStepID)).
		Header(httputil.UserHeader, req.UserID).
		JSONBody(&req).
		Do().JSON(&rsp)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return nil, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return &rsp, nil
}

func (b *Bundle) ExecuteDiceAutotestScene(req apistructs.AutotestExecuteSceneRequest) (*apistructs.AutotestExecuteSceneResponse, error) {
	host, err := b.urls.QA()
	if err != nil {
		return nil, err
	}
	hc := b.hc
	var rsp apistructs.AutotestExecuteSceneResponse
	httpResp, err := hc.Post(host).Path(fmt.Sprintf("/api/autotests/scenes/%v/actions/execute", req.AutoTestScene.ID)).
		Header(httputil.UserHeader, req.UserID).
		JSONBody(&req).
		Do().JSON(&rsp)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return nil, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return &rsp, nil
}

func (b *Bundle) CancelDiceAutotestScene(req apistructs.AutotestCancelSceneRequest) (string, error) {
	host, err := b.urls.QA()
	if err != nil {
		return "", err
	}
	hc := b.hc
	var rsp apistructs.AutotestCancelSceneResponse
	httpResp, err := hc.Post(host).Path(fmt.Sprintf("/api/autotests/scenes/%v/actions/cancel", req.AutoTestScene.ID)).
		Header(httputil.UserHeader, req.UserID).
		JSONBody(&req).
		Do().JSON(&rsp)
	if err != nil {
		return "", apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return "", toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return rsp.Data, nil
}

func (b *Bundle) CreateAutoTestScene(req apistructs.AutotestSceneRequest) (uint64, error) {
	host, err := b.urls.QA()
	if err != nil {
		return 0, err
	}
	hc := b.hc
	var rsp apistructs.AutotestCreateSceneResponse
	httpResp, err := hc.Post(host).Path(fmt.Sprintf("/api/autotests/scenes")).
		Header(httputil.UserHeader, req.UserID).
		JSONBody(&req).
		Do().JSON(&rsp)
	if err != nil {
		return 0, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return 0, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return rsp.Data, nil
}

func (b *Bundle) UpdateAutoTestScene(req apistructs.AutotestSceneSceneUpdateRequest) (uint64, error) {
	host, err := b.urls.QA()
	if err != nil {
		return 0, err
	}
	hc := b.hc
	var rsp apistructs.AutotestCreateSceneResponse
	httpResp, err := hc.Put(host).Path(fmt.Sprintf("/api/autotests/scenes/"+strconv.FormatInt(int64(req.SceneID), 10))).
		Header(httputil.UserHeader, req.UserID).
		JSONBody(&req).
		Do().JSON(&rsp)
	if err != nil {
		return 0, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return 0, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return rsp.Data, nil
}

func (b *Bundle) MoveAutoTestScene(req apistructs.AutotestSceneRequest) (uint64, error) {
	host, err := b.urls.QA()
	if err != nil {
		return 0, err
	}
	hc := b.hc
	var rsp apistructs.AutotestCreateSceneResponse
	httpResp, err := hc.Put(host).Path(fmt.Sprintf("/api/autotests/scenes/actions/move-scene")).
		Header(httputil.UserHeader, req.UserID).
		JSONBody(&req).
		Do().JSON(&rsp)
	if err != nil {
		return 0, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return 0, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return rsp.Data, nil
}

func (b *Bundle) ListAutoTestScene(req apistructs.AutotestSceneRequest) (uint64, []apistructs.AutoTestScene, error) {
	host, err := b.urls.QA()
	if err != nil {
		return 0, nil, err
	}
	hc := b.hc
	var rsp apistructs.AutotestListSceneResponse
	httpResp, err := hc.Get(host).Path(fmt.Sprintf("/api/autotests/scenes")).
		Header(httputil.UserHeader, req.UserID).
		Params(req.URLQueryString()).
		Do().JSON(&rsp)
	if err != nil {
		return 0, nil, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return 0, nil, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return rsp.Data.Total, rsp.Data.List, nil
}

func (b *Bundle) GetAutoTestScene(req apistructs.AutotestSceneRequest) (*apistructs.AutoTestScene, error) {
	host, err := b.urls.QA()
	if err != nil {
		return nil, err
	}
	hc := b.hc
	var rsp apistructs.AutotestGetSceneResponse
	httpResp, err := hc.Get(host).Path(fmt.Sprintf("/api/autotests/scenes/"+strconv.FormatInt(int64(req.SceneID), 10))).
		Header(httputil.UserHeader, req.UserID).
		Params(req.URLQueryString()).
		Do().JSON(&rsp)
	if err != nil {
		return nil, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return nil, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return &rsp.Data, nil
}

func (b *Bundle) DeleteAutoTestScene(req apistructs.AutotestSceneRequest) error {
	host, err := b.urls.QA()
	if err != nil {
		return err
	}
	hc := b.hc
	var rsp apistructs.AutotestCancelSceneResponse
	httpResp, err := hc.Delete(host).Path(fmt.Sprintf("/api/autotests/scenes/"+strconv.FormatInt(int64(req.SceneID), 10))).
		Header(httputil.UserHeader, req.UserID).
		Do().JSON(&rsp)
	if err != nil {
		return apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return nil
}

func (b *Bundle) CopyAutoTestScene(req apistructs.AutotestSceneCopyRequest) (uint64, error) {
	host, err := b.urls.QA()
	if err != nil {
		return 0, err
	}
	hc := b.hc
	var rsp apistructs.AutotestCreateSceneResponse
	httpResp, err := hc.Post(host).Path(fmt.Sprintf("/api/autotests/scenes/actions/copy")).
		Header(httputil.UserHeader, req.UserID).
		JSONBody(&req).
		Do().JSON(&rsp)
	if err != nil {
		return 0, apierrors.ErrInvoke.InternalError(err)
	}
	if !httpResp.IsOK() || !rsp.Success {
		return 0, toAPIError(httpResp.StatusCode(), rsp.Error)
	}

	return rsp.Data, nil
}
