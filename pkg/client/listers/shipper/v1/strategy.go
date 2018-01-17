/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/bookingcom/shipper/pkg/apis/shipper/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StrategyLister helps list Strategies.
type StrategyLister interface {
	// List lists all Strategies in the indexer.
	List(selector labels.Selector) (ret []*v1.Strategy, err error)
	// Strategies returns an object that can list and get Strategies.
	Strategies(namespace string) StrategyNamespaceLister
	StrategyListerExpansion
}

// strategyLister implements the StrategyLister interface.
type strategyLister struct {
	indexer cache.Indexer
}

// NewStrategyLister returns a new StrategyLister.
func NewStrategyLister(indexer cache.Indexer) StrategyLister {
	return &strategyLister{indexer: indexer}
}

// List lists all Strategies in the indexer.
func (s *strategyLister) List(selector labels.Selector) (ret []*v1.Strategy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Strategy))
	})
	return ret, err
}

// Strategies returns an object that can list and get Strategies.
func (s *strategyLister) Strategies(namespace string) StrategyNamespaceLister {
	return strategyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StrategyNamespaceLister helps list and get Strategies.
type StrategyNamespaceLister interface {
	// List lists all Strategies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Strategy, err error)
	// Get retrieves the Strategy from the indexer for a given namespace and name.
	Get(name string) (*v1.Strategy, error)
	StrategyNamespaceListerExpansion
}

// strategyNamespaceLister implements the StrategyNamespaceLister
// interface.
type strategyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Strategies in the indexer for a given namespace.
func (s strategyNamespaceLister) List(selector labels.Selector) (ret []*v1.Strategy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Strategy))
	})
	return ret, err
}

// Get retrieves the Strategy from the indexer for a given namespace and name.
func (s strategyNamespaceLister) Get(name string) (*v1.Strategy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("strategy"), name)
	}
	return obj.(*v1.Strategy), nil
}
