// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package apis

import (
    "github.com/jdcloud-api/jdcloud-sdk-go/core"
    streambus "github.com/jdcloud-api/jdcloud-sdk-go/services/streambus/models"
)

type UpdateTopicRequest struct {

    core.JDCloudRequest

    /* 地域ID  */
    RegionId string `json:"regionId"`

    /* 当更新主题时只需要修改topicModel中的topic中的属性即可；创建归档需要指定target以及归档的目的地(mysql,京东云 Elasticsearch,对象存储,数据计算服务)参数  */
    TopicModel *streambus.TopicModel `json:"topicModel"`
}

/*
 * param regionId: 地域ID (Required)
 * param topicModel: 当更新主题时只需要修改topicModel中的topic中的属性即可；创建归档需要指定target以及归档的目的地(mysql,京东云 Elasticsearch,对象存储,数据计算服务)参数 (Required)
 *
 * @Deprecated, not compatible when mandatory parameters changed
 */
func NewUpdateTopicRequest(
    regionId string,
    topicModel *streambus.TopicModel,
) *UpdateTopicRequest {

	return &UpdateTopicRequest{
        JDCloudRequest: core.JDCloudRequest{
			URL:     "/regions/{regionId}/topic",
			Method:  "PUT",
			Header:  nil,
			Version: "v1",
		},
        RegionId: regionId,
        TopicModel: topicModel,
	}
}

/*
 * param regionId: 地域ID (Required)
 * param topicModel: 当更新主题时只需要修改topicModel中的topic中的属性即可；创建归档需要指定target以及归档的目的地(mysql,京东云 Elasticsearch,对象存储,数据计算服务)参数 (Required)
 */
func NewUpdateTopicRequestWithAllParams(
    regionId string,
    topicModel *streambus.TopicModel,
) *UpdateTopicRequest {

    return &UpdateTopicRequest{
        JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topic",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
        RegionId: regionId,
        TopicModel: topicModel,
    }
}

/* This constructor has better compatible ability when API parameters changed */
func NewUpdateTopicRequestWithoutParam() *UpdateTopicRequest {

    return &UpdateTopicRequest{
            JDCloudRequest: core.JDCloudRequest{
            URL:     "/regions/{regionId}/topic",
            Method:  "PUT",
            Header:  nil,
            Version: "v1",
        },
    }
}

/* param regionId: 地域ID(Required) */
func (r *UpdateTopicRequest) SetRegionId(regionId string) {
    r.RegionId = regionId
}

/* param topicModel: 当更新主题时只需要修改topicModel中的topic中的属性即可；创建归档需要指定target以及归档的目的地(mysql,京东云 Elasticsearch,对象存储,数据计算服务)参数(Required) */
func (r *UpdateTopicRequest) SetTopicModel(topicModel *streambus.TopicModel) {
    r.TopicModel = topicModel
}

// GetRegionId returns path parameter 'regionId' if exist,
// otherwise return empty string
func (r UpdateTopicRequest) GetRegionId() string {
    return r.RegionId
}

type UpdateTopicResponse struct {
    RequestID string `json:"requestId"`
    Error core.ErrorResponse `json:"error"`
    Result UpdateTopicResult `json:"result"`
}

type UpdateTopicResult struct {
    Status bool `json:"status"`
    Message string `json:"message"`
}