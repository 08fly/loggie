/*
Copyright 2021 Loggie Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package index

import (
	"github.com/loggie-io/loggie/pkg/control"
	"github.com/loggie-io/loggie/pkg/discovery/kubernetes/apis/loggie/v1beta1"
	"github.com/loggie-io/loggie/pkg/pipeline"
	"sync"
)

type LogConfigTypeClusterIndex struct {
	pipeConfigs map[string]*TypeClusterPipeConfig // key: logConfigNamespace/Name, value: pipeline configs

	mutex sync.RWMutex
}

type TypeClusterPipeConfig struct {
	Raw []pipeline.Config
	Lgc *v1beta1.LogConfig
}

func NewLogConfigTypeLoggieIndex() *LogConfigTypeClusterIndex {
	return &LogConfigTypeClusterIndex{
		pipeConfigs: make(map[string]*TypeClusterPipeConfig),
	}
}

func (index *LogConfigTypeClusterIndex) GetConfig(logConfigKey string) ([]pipeline.Config, bool) {
	index.mutex.RLock()
	defer index.mutex.RUnlock()

	cfg, ok := index.pipeConfigs[logConfigKey]
	if !ok {
		return nil, false
	}
	return cfg.Raw, true
}

func (index *LogConfigTypeClusterIndex) DeleteConfig(logConfigKey string) bool {
	index.mutex.Lock()
	defer index.mutex.Unlock()

	_, ok := index.GetConfig(logConfigKey)
	if !ok {
		return false
	}
	delete(index.pipeConfigs, logConfigKey)
	return true
}

func (index *LogConfigTypeClusterIndex) setConfig(logConfigKey string, p []pipeline.Config, lgc *v1beta1.LogConfig) {
	index.mutex.Lock()
	defer index.mutex.Unlock()

	index.pipeConfigs[logConfigKey] = &TypeClusterPipeConfig{
		Raw: p,
		Lgc: lgc,
	}
}

func (index *LogConfigTypeClusterIndex) ValidateAndSetConfig(logConfigKey string, p []pipeline.Config, lgc *v1beta1.LogConfig) error {
	index.setConfig(logConfigKey, p, lgc)
	if err := index.GetAll().ValidateUniquePipeName(); err != nil {
		index.DeleteConfig(logConfigKey)
		return err
	}
	return nil
}

func (index *LogConfigTypeClusterIndex) GetAll() *control.PipelineConfig {
	index.mutex.RLock()
	defer index.mutex.RUnlock()

	var cfgRaws []pipeline.Config
	for _, v := range index.pipeConfigs {
		cfgRaws = append(cfgRaws, v.Raw...)
	}
	all := &control.PipelineConfig{
		Pipelines: cfgRaws,
	}
	return all
}

func (index *LogConfigTypeClusterIndex) GetAllConfigMap() map[string]*TypeClusterPipeConfig {
	index.mutex.RLock()
	defer index.mutex.RUnlock()

	return index.pipeConfigs
}
