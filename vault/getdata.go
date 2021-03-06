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

package vault

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type EnvData struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type VaultEnvData struct {
	Env []EnvData `json:"env"`
}

type PublicCert struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
	TTL  int64  `json:"ttl"` // in seconds
}

func GetData(config *Config, path string, version int, data interface{}) error {
	if config == nil {
		return fmt.Errorf("no vault Config specified")
	}
	client, err := config.Login()
	if err != nil {
		return err
	}
	vdat, err := GetKV(client, path, version)
	if err != nil {
		return err
	}
	return mapstructure.WeakDecode(vdat["data"], data)
}

func PutData(config *Config, path string, data interface{}) error {
	client, err := config.Login()
	if err != nil {
		return err
	}
	vdata := map[string]interface{}{
		"data": data,
	}
	out, err := json.Marshal(vdata)
	if err != nil {
		return fmt.Errorf("Failed to marshal data to json: %v", err)
	}

	var vaultData map[string]interface{}
	err = json.Unmarshal(out, &vaultData)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal json to vault data: %v", err)
	}
	return PutKV(client, path, vaultData)
}

func GetEnvVars(config *Config, path string) (map[string]string, error) {
	envData := &VaultEnvData{}
	err := GetData(config, path, 0, envData)
	if err != nil {
		return nil, err
	}
	vars := make(map[string]string, 1)
	for _, envData := range envData.Env {
		vars[envData.Name] = envData.Value
	}
	return vars, nil
}

func GetPublicCert(config *Config, commonName string) (*PublicCert, error) {
	if config == nil {
		return nil, fmt.Errorf("no vault Config specified")
	}
	client, err := config.Login()
	if err != nil {
		return nil, err
	}
	path := "/certs/cert/" + commonName
	vdat, err := GetKV(client, path, 0)
	if err != nil {
		return nil, err
	}
	pubCert := &PublicCert{}
	err = mapstructure.WeakDecode(vdat, pubCert)
	if err != nil {
		return nil, err
	}
	return pubCert, nil
}
