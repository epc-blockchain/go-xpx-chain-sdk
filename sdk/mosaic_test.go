// Copyright 2018 ProximaX Limited. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package sdk

import (
	"fmt"
	"github.com/proximax-storage/go-xpx-utils/mock"
	"github.com/proximax-storage/go-xpx-utils/tests"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

var (
	mosaicClient = mockServer.getPublicTestClientUnsafe().Mosaic
)

const (
	testMosaicPathID = "D525AD41D95FCF29"

	testMosaicInfoJson = `
{
  "mosaic": {
    "mosaicId": [
      3646934825,
      3576016193
    ],
    "supply": [
      3403414400,
      2095475
    ],
    "height": [
      1,
      0
    ],
    "owner": "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
	"revision": 1,
    "properties": [
      [
        2,
        0
      ],
      [
        6,
        0
      ],
      [
        0,
        0
      ]
    ],
    "levy": {}
  }
}`

	testMosaicNamesJson = `[
   {
      "mosaicId":[
         519256100,
         642862634
      ],
      "names":[
         "cat.storage"
      ]
   },
   {
      "mosaicId":[
         481110499,
         231112638
      ],
      "names":[
         "cat.currency"
      ]
   }
]`
)

var (
	mosaicCorr = &MosaicInfo{
		MosaicId: bigIntToMosaicId(uint64DTO{3646934825, 3576016193}.toBigInt()),
		Supply:   uint64DTO{3403414400, 2095475}.toBigInt(),
		Height:   big.NewInt(1),
		Owner: &PublicAccount{
			Address: &Address{
				Type:    mosaicClient.client.config.NetworkType,
				Address: "VBFBW6TUGLEWQIBCMTBMXXQORZKUP3WTVX36ZFE7",
			},

			PublicKey: "321DE652C4D3362FC2DDF7800F6582F4A10CFEA134B81F8AB6E4BE78BBA4D18E",
		},
		Revision: 1,
		Properties: &MosaicProperties{
			Transferable: true,
			Divisibility: 6,
			Duration:     big.NewInt(0),
		},
	}

	mosaicNames = []*MosaicName{
		{
			bigIntToMosaicId(big.NewInt(0x26514E2A1EF33824)),
			[]string{"cat.storage"},
		},
		{
			bigIntToMosaicId(big.NewInt(0x0DC67FBE1CAD29E3)),
			[]string{"cat.currency"},
		},
	}
)

func TestMosaicService_GetMosaic(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     fmt.Sprintf(mosaicRoute, testMosaicPathID),
		RespBody: testMosaicInfoJson,
	})

	mscInfo, err := mosaicClient.GetMosaic(ctx, mosaicCorr.MosaicId)

	assert.Nilf(t, err, "MosaicService.GetMosaic returned error: %s", err)
	tests.ValidateStringers(t, mosaicCorr, mscInfo)
}

func TestMosaicService_GetMosaics(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServer.AddRouter(&mock.Router{
			Path:     mosaicsRoute,
			RespBody: "[" + testMosaicInfoJson + "]",
			ReqJsonBodyStruct: struct {
				MosaicIds []string `json:"mosaicIds"`
			}{},
		})

		mscInfoArr, err := mosaicClient.GetMosaics(ctx, []*MosaicId{mosaicCorr.MosaicId})

		assert.Nilf(t, err, "MosaicService.GetMosaics returned error: %s", err)

		for _, mscInfo := range mscInfoArr {
			tests.ValidateStringers(t, mosaicCorr, mscInfo)
		}
	})

	t.Run("empty url params", func(t *testing.T) {
		_, err := mosaicClient.GetMosaics(ctx, []*MosaicId{})

		assert.NotNil(t, err, "MosaicService.GetMosaics returned error: %s", err)
	})
}

func TestMosaicService_GetMosaicsNames(t *testing.T) {
	mockServer.AddRouter(&mock.Router{
		Path:     mosaicNamesRoute,
		RespBody: testMosaicNamesJson,
		ReqJsonBodyStruct: struct {
			MosaicIds []string `json:"mosaicIds"`
		}{},
	})

	mscNameArr, err := mosaicClient.GetMosaicsNames(ctx, mosaicNames[0].MosaicId, mosaicNames[1].MosaicId)

	assert.Nilf(t, err, "MosaicService.GetMosaics returned error: %s", err)

	for i, mscName := range mscNameArr {
		tests.ValidateStringers(t, mosaicNames[i], mscName)
	}
}
