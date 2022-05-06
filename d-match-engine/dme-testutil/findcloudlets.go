// Copyright 2022 MobiledgeX, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dmetest

import dme "github.com/edgexr/edge-cloud/d-match-engine/dme-proto"

type FindCloudletRR struct {
	Reg            dme.RegisterClientRequest
	Req            dme.FindCloudletRequest
	Reply          dme.FindCloudletReply
	ReplyCarrier   string
	ReplyCloudlet  string
	ReplyAlternate dme.FindCloudletReply
}

type GetAppInstListRR struct {
	Reg   dme.RegisterClientRequest
	Req   dme.AppInstListRequest
	Reply dme.AppInstListReply
}

// FindCloudlet API test data.
// Replies are based on AppInst data generated by GenerateAppInsts()
// in this package.
var FindCloudletData = []FindCloudletRR{
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Untomt",
			AppName: "Untomt",
			AppVers: "1.1",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "GDDT",
			GpsLocation: &dme.Loc{Latitude: 50.65, Longitude: 6.341},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[2].Uri,
			CloudletLocation: &Cloudlets[2].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[2].CarrierName,
		ReplyCloudlet: Cloudlets[2].Name,
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName:      "Untomt",
			AppName:      "Untomt",
			AppVers:      "1.1",
			UniqueId:     "123",
			UniqueIdType: "1000Realities",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "GDDT",
			GpsLocation: &dme.Loc{Latitude: 51.65, Longitude: 9.341},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[1].Uri,
			CloudletLocation: &Cloudlets[1].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[1].CarrierName,
		ReplyCloudlet: Cloudlets[1].Name,
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName:      "Untomt",
			AppName:      "Untomt",
			AppVers:      "1.1",
			UniqueId:     "1234",
			UniqueIdType: "1000Realities",
		}, // ATT does not exist and so should return default cloudlet
		Req: dme.FindCloudletRequest{
			CarrierName: "ATT",
			GpsLocation: &dme.Loc{Latitude: 52.65, Longitude: 10.341},
		},
		Reply: dme.FindCloudletReply{
			Status: 2,
		},
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Untomt",
			AppName: "Untomt",
			AppVers: "1.1",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "GDDT",
			GpsLocation: &dme.Loc{Latitude: 50.75, Longitude: 7.9050},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[2].Uri,
			CloudletLocation: &Cloudlets[2].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[2].CarrierName,
		ReplyCloudlet: Cloudlets[2].Name,
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Atlantic Labs",
			AppName: "Pillimo-go",
			AppVers: "2.1",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "GDDT",
			GpsLocation: &dme.Loc{Latitude: 52.75, Longitude: 12.9050},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[1].Uri,
			CloudletLocation: &Cloudlets[1].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[1].CarrierName,
		ReplyCloudlet: Cloudlets[1].Name,
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Atlantic Labs",
			AppName: "HarryPotter-go",
			AppVers: "1.0",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "GDDT",
			GpsLocation: &dme.Loc{Latitude: 50.75, Longitude: 11.9050},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[1].Uri,
			CloudletLocation: &Cloudlets[1].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[1].CarrierName,
		ReplyCloudlet: Cloudlets[1].Name,
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Ever.AI",
			AppName: "Ever",
			AppVers: "1.7",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "DMUUS",
			GpsLocation: &dme.Loc{Latitude: 47.75, Longitude: 122.9050},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[3].Uri,
			CloudletLocation: &Cloudlets[3].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[3].CarrierName,
		ReplyCloudlet: Cloudlets[3].Name,
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Ever.AI",
			AppName: "Ever",
			AppVers: "1.7",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "",
			GpsLocation: &dme.Loc{Latitude: 47.75, Longitude: 122.9050},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[3].Uri,
			CloudletLocation: &Cloudlets[3].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[3].CarrierName,
		ReplyCloudlet: Cloudlets[3].Name,
	},
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Ever.AI",
			AppName: "Ever",
			AppVers: "1.7",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "",
			GpsLocation: &dme.Loc{Latitude: 48.31, Longitude: 11.66},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[2].Uri,
			CloudletLocation: &Cloudlets[2].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[2].CarrierName,
		ReplyCloudlet: Cloudlets[2].Name,
	},
}

var FindCloudletAllianceOrg = []FindCloudletRR{
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Atlantic Labs",
			AppName: "HarryPotter-go",
			AppVers: "1.0",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "DMUUS",
			GpsLocation: &dme.Loc{Latitude: 50.75, Longitude: 11.9050},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[1].Uri,
			CloudletLocation: &Cloudlets[1].Location,
			Status:           1,
		},
		ReplyCarrier:  "DMUUS", // real cloudlet is GDDT, but treat as DMUUS
		ReplyCloudlet: Cloudlets[1].Name,
	},
}

var FindCloudletNoAllianceOrg = []FindCloudletRR{
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Atlantic Labs",
			AppName: "HarryPotter-go",
			AppVers: "1.0",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "DMUUS",
			GpsLocation: &dme.Loc{Latitude: 50.75, Longitude: 11.9050},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[3].Uri,
			CloudletLocation: &Cloudlets[3].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[3].CarrierName,
		ReplyCloudlet: Cloudlets[3].Name,
	},
}

var FindCloudletHAData = []FindCloudletRR{
	FindCloudletRR{
		Reg: dme.RegisterClientRequest{
			OrgName: "Ever.AI",
			AppName: "Ever",
			AppVers: "1.7",
		},
		Req: dme.FindCloudletRequest{
			CarrierName: "",
			GpsLocation: &dme.Loc{Latitude: -50, Longitude: -100},
		},
		Reply: dme.FindCloudletReply{
			Fqdn:             Cloudlets[4].Uri,
			CloudletLocation: &Cloudlets[4].Location,
			Status:           1,
		},
		ReplyCarrier:  Cloudlets[4].CarrierName,
		ReplyCloudlet: Cloudlets[4].Name,
		ReplyAlternate: dme.FindCloudletReply{
			Fqdn:             Cloudlets[5].Uri,
			CloudletLocation: &Cloudlets[5].Location,
			Status:           1,
		},
	},
}

// copy of FindCloudletData[3] with a changed reply to Sunnydale cloudlet
var DisabledCloudletRR = FindCloudletRR{
	Reg: dme.RegisterClientRequest{
		OrgName: "Untomt",
		AppName: "Untomt",
		AppVers: "1.1",
	},
	Req: dme.FindCloudletRequest{
		CarrierName: "GDDT",
		GpsLocation: &dme.Loc{Latitude: 50.75, Longitude: 7.9050},
	},
	Reply: dme.FindCloudletReply{
		Fqdn:             Cloudlets[1].Uri,
		CloudletLocation: &Cloudlets[1].Location,
		Status:           1,
	},
}

var GetAppInstListData = []GetAppInstListRR{
	GetAppInstListRR{
		Reg: dme.RegisterClientRequest{ // Pillimo-go
			OrgName: Apps[1].Organization,
			AppName: Apps[1].Name,
			AppVers: Apps[1].Vers,
		},
		Req: dme.AppInstListRequest{
			CarrierName: "",
			GpsLocation: &dme.Loc{
				Latitude:  51,
				Longitude: 11,
			},
			Limit: 4,
		},
		Reply: dme.AppInstListReply{
			Status: dme.AppInstListReply_AI_SUCCESS,
			Cloudlets: []*dme.CloudletLocation{
				&dme.CloudletLocation{
					CarrierName:  "GDDT",
					CloudletName: "Sunnydale",
					GpsLocation:  &Cloudlets[1].Location,
				},
				&dme.CloudletLocation{
					CarrierName:  "GDDT",
					CloudletName: "Buckhorn",
					GpsLocation:  &Cloudlets[0].Location,
				},
				&dme.CloudletLocation{
					CarrierName:  "GDDT",
					CloudletName: "Beacon",
					GpsLocation:  &Cloudlets[2].Location,
				},
				&dme.CloudletLocation{
					CarrierName:  "DMUUS",
					CloudletName: "San Francisco",
					GpsLocation:  &Cloudlets[3].Location,
				},
			},
		},
	},
}

var GetAppInstListAllianceOrg = []GetAppInstListRR{
	GetAppInstListRR{
		Reg: dme.RegisterClientRequest{ //HarryPotter-go
			OrgName: Apps[2].Organization,
			AppName: Apps[2].Name,
			AppVers: Apps[2].Vers,
		},
		Req: dme.AppInstListRequest{
			CarrierName: "DMUUS",
			GpsLocation: &dme.Loc{
				Latitude:  50.75,
				Longitude: 11.9050,
			},
			Limit: 4,
		},
		Reply: dme.AppInstListReply{
			Status: dme.AppInstListReply_AI_SUCCESS,
			Cloudlets: []*dme.CloudletLocation{
				&dme.CloudletLocation{
					CarrierName:  "GDDT",
					CloudletName: "Sunnydale",
					GpsLocation:  &Cloudlets[1].Location,
				},
				&dme.CloudletLocation{
					CarrierName:  "DMUUS",
					CloudletName: "San Francisco",
					GpsLocation:  &Cloudlets[3].Location,
				},
			},
		},
	},
}

var GetAppInstListNoAllianceOrg = []GetAppInstListRR{
	GetAppInstListRR{
		Reg: dme.RegisterClientRequest{ //HarryPotter-go
			OrgName: Apps[2].Organization,
			AppName: Apps[2].Name,
			AppVers: Apps[2].Vers,
		},
		Req: dme.AppInstListRequest{
			CarrierName: "DMUUS",
			GpsLocation: &dme.Loc{
				Latitude:  50.75,
				Longitude: 11.9050,
			},
			Limit: 4,
		},
		Reply: dme.AppInstListReply{
			Status: dme.AppInstListReply_AI_SUCCESS,
			Cloudlets: []*dme.CloudletLocation{
				&dme.CloudletLocation{
					CarrierName:  "DMUUS",
					CloudletName: "San Francisco",
					GpsLocation:  &Cloudlets[3].Location,
				},
			},
		},
	},
}
