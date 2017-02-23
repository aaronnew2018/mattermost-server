// Copyright (c) 2015 Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"io"
	"strings"
)

const (
	PUSH_NOTIFY_APPLE                = "apple"
	PUSH_NOTIFY_ANDROID              = "android"
	PUSH_NOTIFY_APPLE_REACT_NATIVE   = "apple_rn"
	PUSH_NOTIFY_ANDROID_REACT_NATIVE = "android_rn"

	PUSH_TYPE_MESSAGE = "message"
	PUSH_TYPE_CLEAR   = "clear"

	CATEGORY_DM = "DIRECT_MESSAGE"

	MHPNS = "https://push.mattermost.com"
)

type PushNotification struct {
	Platform         string `json:"platform"`
	ServerId         string `json:"server_id"`
	DeviceId         string `json:"device_id"`
	Category         string `json:"category"`
	Sound            string `json:"sound"`
	Message          string `json:"message"`
	Badge            int    `json:"badge"`
	ContentAvailable int    `json:"cont_ava"`
	TeamId           string `json:"team_id"`
	ChannelId        string `json:"channel_id"`
	ChannelName      string `json:"channel_name"`
	Type             string `json:"type"`
}

func (me *PushNotification) ToJson() string {
	b, err := json.Marshal(me)
	if err != nil {
		return ""
	} else {
		return string(b)
	}
}

func (me *PushNotification) SetDeviceIdAndPlatform(deviceId string) {

	parts := strings.Split(deviceId, ":")
	if len(parts) == 2 {
		me.Platform = parts[0]
		me.DeviceId = parts[1]
	}
}

func PushNotificationFromJson(data io.Reader) *PushNotification {
	decoder := json.NewDecoder(data)
	var me PushNotification
	err := decoder.Decode(&me)
	if err == nil {
		return &me
	} else {
		return nil
	}
}
