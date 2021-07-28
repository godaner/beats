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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package pensando

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "pensando", asset.ModuleFieldsPri, AssetPensando); err != nil {
		panic(err)
	}
}

// AssetPensando returns asset data.
// This is the base64 encoded zlib format compressed contents of module/pensando.
func AssetPensando() string {
	return "eJy0k92qnDAQx+99inmBsw/gRaGwHOhFYekp9FJSM3HDiU5IRmTfviSrbrSR/aidy5n4//+cjzf4xEsJFjsvOkkFAGs2WMJpzEABINHXTlvW1JXwpQCA+QP4TrI3WAAojUb6MlbfoBMtLmRD8MViCY2j3o6ZjHKI96gFylF7AzHU+MP4JDVLDaUaYE7mDEM4NCg8lvAbWSR5iUr0hqsoXoISxuOinGVNcMndaI/vv5Ina96UWdRBclGayD/xMpCTq9o2SIyvUQ+oAz4jKEPDAfLG1lZ6LX411h1jgw6edLbW6FpE+2/HvKlEz7qLbyohpUPv1y6Lf38WYZQklTptNCBlseQ4D/JaL05B7zEK7fA6sv124DhrkprXIG9vHTHVZHZ0P02Sd81db7DSfw35XzbgR28wbB+fBcMgPLSC6zPKje579D7Mfwvitel/XFUDx90eTASeBeOOjZgYom6KsdUJ6l2N/+MoP6IyiNttPoay902OHHY8zTsQrFv0LFqbR5Chq8/5/5wVR3dDzQGKPwEAAP//dBvr3Q=="
}
