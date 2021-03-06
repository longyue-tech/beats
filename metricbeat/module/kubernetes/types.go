// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package kubernetes

type Summary struct {
	Node struct {
		CPU struct {
			Time                 string `json:"time"`
			UsageCoreNanoSeconds int64  `json:"usageCoreNanoSeconds"`
			UsageNanoCores       int64  `json:"usageNanoCores"`
		} `json:"cpu"`
		Fs struct {
			AvailableBytes int64 `json:"availableBytes"`
			CapacityBytes  int64 `json:"capacityBytes"`
			Inodes         int64 `json:"inodes"`
			InodesFree     int64 `json:"inodesFree"`
			InodesUsed     int64 `json:"inodesUsed"`
			UsedBytes      int64 `json:"usedBytes"`
		} `json:"fs"`
		Memory struct {
			AvailableBytes  int64  `json:"availableBytes"`
			MajorPageFaults int64  `json:"majorPageFaults"`
			PageFaults      int64  `json:"pageFaults"`
			RssBytes        int64  `json:"rssBytes"`
			Time            string `json:"time"`
			UsageBytes      int64  `json:"usageBytes"`
			WorkingSetBytes int64  `json:"workingSetBytes"`
		} `json:"memory"`
		Network struct {
			RxBytes  int64  `json:"rxBytes"`
			RxErrors int64  `json:"rxErrors"`
			Time     string `json:"time"`
			TxBytes  int64  `json:"txBytes"`
			TxErrors int64  `json:"txErrors"`
		} `json:"network"`
		NodeName string `json:"nodeName"`
		Runtime  struct {
			ImageFs struct {
				AvailableBytes int64 `json:"availableBytes"`
				CapacityBytes  int64 `json:"capacityBytes"`
				UsedBytes      int64 `json:"usedBytes"`
			} `json:"imageFs"`
		} `json:"runtime"`
		StartTime        string `json:"startTime"`
		SystemContainers []struct {
			CPU struct {
				Time                 string `json:"time"`
				UsageCoreNanoSeconds int64  `json:"usageCoreNanoSeconds"`
				UsageNanoCores       int64  `json:"usageNanoCores"`
			} `json:"cpu"`
			Memory struct {
				MajorPageFaults int64  `json:"majorPageFaults"`
				PageFaults      int64  `json:"pageFaults"`
				RssBytes        int64  `json:"rssBytes"`
				Time            string `json:"time"`
				UsageBytes      int64  `json:"usageBytes"`
				WorkingSetBytes int64  `json:"workingSetBytes"`
			} `json:"memory"`
			Name               string      `json:"name"`
			StartTime          string      `json:"startTime"`
			UserDefinedMetrics interface{} `json:"userDefinedMetrics"`
		} `json:"systemContainers"`
	} `json:"node"`
	Pods []struct {
		Containers []struct {
			CPU struct {
				Time                 string `json:"time"`
				UsageCoreNanoSeconds int64  `json:"usageCoreNanoSeconds"`
				UsageNanoCores       int64  `json:"usageNanoCores"`
			} `json:"cpu"`
			Logs struct {
				AvailableBytes int64 `json:"availableBytes"`
				CapacityBytes  int64 `json:"capacityBytes"`
				Inodes         int64 `json:"inodes"`
				InodesFree     int64 `json:"inodesFree"`
				InodesUsed     int64 `json:"inodesUsed"`
				UsedBytes      int64 `json:"usedBytes"`
			} `json:"logs"`
			Memory struct {
				AvailableBytes  int64  `json:"availableBytes"`
				MajorPageFaults int64  `json:"majorPageFaults"`
				PageFaults      int64  `json:"pageFaults"`
				RssBytes        int64  `json:"rssBytes"`
				Time            string `json:"time"`
				UsageBytes      int64  `json:"usageBytes"`
				WorkingSetBytes int64  `json:"workingSetBytes"`
			} `json:"memory"`
			Name   string `json:"name"`
			Rootfs struct {
				AvailableBytes int64 `json:"availableBytes"`
				CapacityBytes  int64 `json:"capacityBytes"`
				InodesUsed     int64 `json:"inodesUsed"`
				UsedBytes      int64 `json:"usedBytes"`
			} `json:"rootfs"`
			StartTime          string      `json:"startTime"`
			UserDefinedMetrics interface{} `json:"userDefinedMetrics"`
		} `json:"containers"`
		Network struct {
			RxBytes  int64  `json:"rxBytes"`
			RxErrors int64  `json:"rxErrors"`
			Time     string `json:"time"`
			TxBytes  int64  `json:"txBytes"`
			TxErrors int64  `json:"txErrors"`
		} `json:"network"`
		PodRef struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
			UID       string `json:"uid"`
		} `json:"podRef"`
		StartTime string `json:"startTime"`
		Volume    []struct {
			AvailableBytes int64  `json:"availableBytes"`
			CapacityBytes  int64  `json:"capacityBytes"`
			Inodes         int64  `json:"inodes"`
			InodesFree     int64  `json:"inodesFree"`
			InodesUsed     int64  `json:"inodesUsed"`
			Name           string `json:"name"`
			UsedBytes      int64  `json:"usedBytes"`
		} `json:"volume"`
	} `json:"pods"`
}
