// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package coredns

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "coredns", asset.ModuleFieldsPri, AssetCoredns); err != nil {
		panic(err)
	}
}

// AssetCoredns returns asset data.
// This is the base64 encoded zlib format compressed contents of module/coredns.
func AssetCoredns() string {
	return "eJzElk9P4zoUxff9FEddwRPk6T09saj0ZlMWw2IYBOwH175tPSR28HVA5dOPnD9tmoSGMh7hVXRtn/vzzfFNzvFImxmkdaQMTwCvfUozTOdVZDoBFLF0Ovfamhm+TAA06/HNqiKlCeAoJcE0w0pMgKWmVPGsXHoOIzJqpwjDb/Kw2NkiryMDWcJ4qPc9QFrjhTYM9sJr9loy/Fp4vJAjOBIKS2ezMtHl9V2t0EZp4wQR3kaHgA5AhTEfwHGUCk8K3sKvqSEBk3vWklrb98vVjC5rmzcXRstE2sL4vfmGPLVm1Zk4AF+Oe+tFClNkC3KwyyoFD2ZXhhNHTwWx/xMMTwW5DfrCQ+lV4USQSwwni0I+kk/+GqSxi58ku6BV8MeHeW8rCjQUWGv2duVEhgqGoQ2MMJZJWqPGy9k+DxfZewu7tC4TfrYF+dg5eLv/DFxkwQVN4LcOEtUjA6w709bZx/FYv1Ky2Hj6TNvc6VfqW+YY9iMsMgL1BlM3wQhPzDf9BtE7+4KN67rrrcVCd9JUf2zW4pnK5n75HQvtwTSOFhjiwmEe1PbxcnJ4tYYgjCqVB7l6Ew3FI21erFPHgXy1qeKyHFULD1KBKUTq0x+oDufWMCVOWhW3Pq3qNFk69SlzDqL1ZyJVqBTeFafCGkRYikynm3gM92uCUMoRc63dYHgnDOfWeZz8g/9xdYOTqxs8k+Pwbfvv9Az/luGLvfjF6en4W/3UfovbGuP4jtvHj9lyh6kO99w+UdymO8z0dtcNv7Pk4tqz0myuhV6kFP5wRvtIuM9xQYJi0EbBpLoMfx+8trmz3sal2d3PTt/AybRQ+RTWYeplPt1dyP4DOo6SQq4pWWsf10jzIIsgW3lnW7wyH/K0WGkzApVp5sj+rrAq4UNgvwIAAP//o0USBg=="
}
