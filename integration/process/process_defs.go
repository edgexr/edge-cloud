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

package process

import (
	"os/exec"
)

type Vault struct {
	Common     `yaml:",inline"`
	DmeSecret  string
	Regions    string
	VaultDatas []VaultData
	ListenAddr string
	RootToken  string
	CADir      string
	cmd        *exec.Cmd
}

type VaultData struct {
	Path string
	Data map[string]string
}

type Etcd struct {
	Common         `yaml:",inline"`
	DataDir        string
	PeerAddrs      string
	ClientAddrs    string
	InitialCluster string
	cmd            *exec.Cmd
}
type Controller struct {
	Common               `yaml:",inline"`
	NodeCommon           `yaml:",inline"`
	RedisClientCommon    `yaml:",inline"`
	EtcdAddrs            string
	ApiAddr              string
	HttpAddr             string
	NotifyAddr           string
	NotifyRootAddrs      string
	NotifyParentAddrs    string
	EdgeTurnAddr         string
	InfluxAddr           string
	Region               string
	cmd                  *exec.Cmd
	TestMode             bool
	RegistryFQDN         string
	ArtifactoryFQDN      string
	CloudletRegistryPath string
	VersionTag           string
	CloudletVMImagePath  string
	CheckpointInterval   string
	AppDNSRoot           string
	ChefServerPath       string
	ThanosRecvAddr       string
}
type Dme struct {
	Common      `yaml:",inline"`
	NodeCommon  `yaml:",inline"`
	ApiAddr     string
	HttpAddr    string
	NotifyAddrs string
	LocVerUrl   string
	TokSrvUrl   string
	QosPosUrl   string
	QosSesAddr  string
	Carrier     string
	CloudletKey string
	CookieExpr  string
	Region      string
	cmd         *exec.Cmd
}
type Crm struct {
	Common              `yaml:",inline"`
	NodeCommon          `yaml:",inline"`
	RedisClientCommon   `yaml:",inline"`
	NotifyAddrs         string
	NotifySrvAddr       string
	CloudletKey         string
	Platform            string
	Plugin              string
	cmd                 *exec.Cmd
	PhysicalName        string
	TestMode            bool
	Span                string
	ContainerVersion    string
	VMImageVersion      string
	CloudletVMImagePath string
	Region              string
	CommercialCerts     bool
	AppDNSRoot          string
	ChefServerPath      string
	CacheDir            string
	HARole              HARole
}
type LocApiSim struct {
	Common  `yaml:",inline"`
	Port    int
	Locfile string
	Geofile string
	Country string
	cmd     *exec.Cmd
}
type TokSrvSim struct {
	Common `yaml:",inline"`
	Port   int
	Token  string
	cmd    *exec.Cmd
}
type SampleApp struct {
	Common       `yaml:",inline"`
	Exename      string
	Args         []string
	Command      string
	VolumeMounts []string
	cmd          *exec.Cmd
}
type Influx struct {
	Common   `yaml:",inline"`
	DataDir  string
	HttpAddr string
	BindAddr string
	Config   string // set during Start
	TLS      TLSCerts
	Auth     LocalAuth
	cmd      *exec.Cmd
}
type ClusterSvc struct {
	Common         `yaml:",inline"`
	NodeCommon     `yaml:",inline"`
	NotifyAddrs    string
	CtrlAddrs      string
	PromPorts      string
	InfluxDB       string
	Interval       string
	Region         string
	PluginRequired bool
	cmd            *exec.Cmd
}
type DockerGeneric struct {
	Common        `yaml:",inline"`
	Links         []string
	DockerNetwork string
	DockerEnvVars map[string]string
	TLS           TLSCerts
	cmd           *exec.Cmd
}
type DockerNetwork struct {
	Common `yaml:",inline"`
}
type Jaeger struct {
	DockerGeneric `yaml:",inline"`
}
type ElasticSearch struct {
	DockerGeneric `yaml:",inline"`
	Type          string
}
type NginxProxy struct {
	DockerGeneric `yaml:",inline"`
	Servers       []NginxServerConfig
}
type NginxServerConfig struct {
	ServerName string
	Port       string
	TlsPort    string
	Target     string
}
type Traefik struct {
	Common        `yaml:",inline"`
	DockerNetwork string
	TLS           TLSCerts
	cmd           *exec.Cmd
}

type RedisCache struct {
	Common     `yaml:",inline"`
	cmd        *exec.Cmd
	Type       string
	Port       string
	MasterPort string
}

type NotifyRoot struct {
	Common     `yaml:",inline"`
	NodeCommon `yaml:",inline"`
	cmd        *exec.Cmd
}
type EdgeTurn struct {
	Common     `yaml:",inline"`
	NodeCommon `yaml:",inline"`
	cmd        *exec.Cmd
	ListenAddr string
	ProxyAddr  string
	Region     string
	TestMode   bool
}
