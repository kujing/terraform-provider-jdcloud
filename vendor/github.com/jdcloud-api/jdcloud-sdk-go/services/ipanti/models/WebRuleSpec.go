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

package models


type WebRuleSpec struct {

    /* 子域名 (Optional) */
    Domain *string `json:"domain"`

    /* 协议：HTTP、HTTPS、HTTP_HTTPS (Optional) */
    Protocol *string `json:"protocol"`

    /* HTTP协议的端口号，如80,81，多个端口号使用逗号分隔 (Optional) */
    Port *string `json:"port"`

    /* HTTPS协议的端口号，如443,8443，多个端口号使用逗号分隔 (Optional) */
    HttpsPort *string `json:"httpsPort"`

    /* 回源类型：A或者CNAME (Optional) */
    OriginType *string `json:"originType"`

    /*  (Optional) */
    OriginAddr []OriginAddrItem `json:"originAddr"`

    /*  (Optional) */
    OnlineAddr []string `json:"onlineAddr"`

    /* 回源域名,originType为CNAME时需要指定该字段 (Optional) */
    OriginDomain *string `json:"originDomain"`

    /* 证书内容 (Optional) */
    HttpsCertContent *string `json:"httpsCertContent"`

    /* 证书私钥 (Optional) */
    HttpsRsaKey *string `json:"httpsRsaKey"`

    /* 转发规则：wrr->带权重的轮询，rr->不带权重的轮询 (Optional) */
    Algorithm *string `json:"algorithm"`

    /* 是否开启https强制跳转，当protocol为HTTP_HTTPS时可以配置该属性 0为不强跳 1为开启强跳 (Optional) */
    ForceJump *int `json:"forceJump"`

    /* 是否为自定义端口号，0为默认 1为自定义 (Optional) */
    CustomPortStatus *int `json:"customPortStatus"`

    /* 是否开启http回源，0为不开启 1为开启，当勾选HTTPS时可以配置该属性 (Optional) */
    HttpOrigin *int `json:"httpOrigin"`
}
