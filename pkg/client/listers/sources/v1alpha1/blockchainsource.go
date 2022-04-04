/*
Copyright 2022 The Knative Authors

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "knative.dev/eventing-blockchain/pkg/apis/sources/v1alpha1"
)

// BlockchainSourceLister helps list BlockchainSources.
// All objects returned here must be treated as read-only.
type BlockchainSourceLister interface {
	// List lists all BlockchainSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BlockchainSource, err error)
	// BlockchainSources returns an object that can list and get BlockchainSources.
	BlockchainSources(namespace string) BlockchainSourceNamespaceLister
	BlockchainSourceListerExpansion
}

// blockchainSourceLister implements the BlockchainSourceLister interface.
type blockchainSourceLister struct {
	indexer cache.Indexer
}

// NewBlockchainSourceLister returns a new BlockchainSourceLister.
func NewBlockchainSourceLister(indexer cache.Indexer) BlockchainSourceLister {
	return &blockchainSourceLister{indexer: indexer}
}

// List lists all BlockchainSources in the indexer.
func (s *blockchainSourceLister) List(selector labels.Selector) (ret []*v1alpha1.BlockchainSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BlockchainSource))
	})
	return ret, err
}

// BlockchainSources returns an object that can list and get BlockchainSources.
func (s *blockchainSourceLister) BlockchainSources(namespace string) BlockchainSourceNamespaceLister {
	return blockchainSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BlockchainSourceNamespaceLister helps list and get BlockchainSources.
// All objects returned here must be treated as read-only.
type BlockchainSourceNamespaceLister interface {
	// List lists all BlockchainSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BlockchainSource, err error)
	// Get retrieves the BlockchainSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BlockchainSource, error)
	BlockchainSourceNamespaceListerExpansion
}

// blockchainSourceNamespaceLister implements the BlockchainSourceNamespaceLister
// interface.
type blockchainSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BlockchainSources in the indexer for a given namespace.
func (s blockchainSourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BlockchainSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BlockchainSource))
	})
	return ret, err
}

// Get retrieves the BlockchainSource from the indexer for a given namespace and name.
func (s blockchainSourceNamespaceLister) Get(name string) (*v1alpha1.BlockchainSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("blockchainsource"), name)
	}
	return obj.(*v1alpha1.BlockchainSource), nil
}