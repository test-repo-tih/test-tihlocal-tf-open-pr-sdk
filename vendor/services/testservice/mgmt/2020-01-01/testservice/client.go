// Package testservice implements the Azure ARM Testservice service API version 2020-01-01.
//
// Test Service Client for SDK Automation integration test.
package testservice

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
    "github.com/Azure/go-autorest/autorest"
)

const (
// DefaultBaseURI is the default URI used for the service Testservice
DefaultBaseURI = "https://management.azure.com")

// BaseClient is the base client for Testservice.
type BaseClient struct {
    autorest.Client
    BaseURI string
}

// New creates an instance of the BaseClient client.
func New()BaseClient {
    return NewWithBaseURI(DefaultBaseURI, )
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string, ) BaseClient {
    return BaseClient{
        Client: autorest.NewClientWithUserAgent(UserAgent()),
        BaseURI: baseURI,
    }
}

