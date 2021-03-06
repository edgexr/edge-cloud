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

package log

import (
	"fmt"
	"strings"
	"testing"

	yaml "github.com/mobiledgex/yaml/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type yamltest struct {
	Lvl  DebugLevel
	Lvls []DebugLevel
}

func TestYaml(t *testing.T) {
	y := yamltest{
		Lvl:  DebugLevel_api,
		Lvls: []DebugLevel{DebugLevel_notify, DebugLevel_etcd},
	}
	out, err := yaml.Marshal(&y)
	assert.Nil(t, err, "marshal yaml")

	strout := string(out)
	fmt.Println(strout)
	assert.True(t, strings.Index(strout, "lvl: Api") != -1, "has api")
	assert.True(t, strings.Index(strout, "lvls: [Notify, Etcd]") != -1, "has array")

	yin := yamltest{}
	err = yaml.Unmarshal(out, &yin)
	assert.Nil(t, err, "unmarshal yaml")
	assert.Equal(t, y, yin, "equal")
}

func TestBits(t *testing.T) {
	lvl := uint64(1) | uint64(1)<<10
	SetDebugLevel(lvl)
	assert.Equal(t, lvl, debugLevel)
	ClearDebugLevel(uint64(1))
	assert.Equal(t, uint64(1)<<10, debugLevel)
	ClearDebugLevel(uint64(1) << 10)
	assert.Equal(t, uint64(0), debugLevel)

	SetDebugLevelEnum(DebugLevel_api)
	assert.True(t, debugLevel&DebugLevelApi != 0)
	SetDebugLevelEnum(DebugLevel_notify)
	assert.True(t, debugLevel&DebugLevelNotify != 0)
	ClearDebugLevelEnum(DebugLevel_api)
	assert.True(t, debugLevel&DebugLevelApi == 0)
	ClearDebugLevelEnum(DebugLevel_notify)
	assert.True(t, debugLevel&DebugLevelNotify == 0)

	SetDebugLevels([]DebugLevel{DebugLevel_api, DebugLevel_etcd})
	assert.True(t, debugLevel&DebugLevelApi != 0)
	assert.True(t, debugLevel&DebugLevelEtcd != 0)
	ClearDebugLevels([]DebugLevel{DebugLevel_notify, DebugLevel_etcd})
	assert.True(t, debugLevel&DebugLevelApi != 0)
	assert.True(t, debugLevel&DebugLevelNotify == 0)
	assert.True(t, debugLevel&DebugLevelEtcd == 0)
	ClearDebugLevels([]DebugLevel{DebugLevel_api, DebugLevel_etcd})
	assert.Equal(t, uint64(0), debugLevel)
}

func TestDebugStrs(t *testing.T) {
	testDebugStrsAdd(t, "etcd,dmereq", "etcd,dmereq")
	testDebugStrsAdd(t, "notify,locapi", "etcd,notify,dmereq,locapi")
	testDebugStrsClear(t, "etcd,locapi", "notify,dmereq")
	testDebugStrsAdd(t, "api,infra,metrics", "api,notify,dmereq,infra,metrics")
	testDebugStrsClear(t, "infra", "api,notify,dmereq,metrics")
	testDebugStrsClear(t, "api,notify,dmereq,metrics", "")
}

func testDebugStrsAdd(t *testing.T, add, expected string) {
	SetDebugLevelStrs(add)
	actual := GetDebugLevelStrs()
	require.Equal(t, expected, actual)
}

func testDebugStrsClear(t *testing.T, remove, expected string) {
	ClearDebugLevelStrs(remove)
	actual := GetDebugLevelStrs()
	require.Equal(t, expected, actual)
}
