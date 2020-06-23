package handlers

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	mappers "github.com/epc-blockchain/go-xpx-chain-sdk/mocks"

	mocksSubscribers "github.com/epc-blockchain/go-xpx-chain-sdk/mocks/websocket/subscribers"
	"github.com/epc-blockchain/go-xpx-chain-sdk/sdk"
	"github.com/epc-blockchain/go-xpx-chain-sdk/sdk/websocket/subscribers"
)

func Test_cosignatureHandler_Handle(t *testing.T) {
	type fields struct {
		messageMapper sdk.CosignatureMapper
		handlers      subscribers.Cosignature
	}
	type args struct {
		address *sdk.Address
		resp    []byte
	}

	address := new(sdk.Address)

	obj := new(sdk.SignerInfo)
	messageMapperMock := new(mappers.CosignatureMapper)
	messageMapperMock.On("MapCosignature", mock.Anything).Return(obj, nil)

	handlerFunc1 := func(*sdk.SignerInfo) bool {
		return false
	}

	handlerFunc2 := func(*sdk.SignerInfo) bool {
		return true
	}

	blockHandler1 := subscribers.CosignatureHandler(handlerFunc1)
	blockHandler2 := subscribers.CosignatureHandler(handlerFunc2)

	blockHandlers := []*subscribers.CosignatureHandler{
		&blockHandler1,
		&blockHandler2,
	}

	blockHandlersMock := new(mocksSubscribers.Cosignature)
	blockHandlersMock.On("GetHandlers", mock.Anything).Return(nil).Once().
		On("GetHandlers", mock.Anything).Return(blockHandlers).
		On("RemoveHandlers", mock.Anything, mock.Anything).Return(true, nil).
		On("HasHandlers", mock.Anything).Return(true, nil)

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "empty handlers",
			fields: fields{
				handlers:      blockHandlersMock,
				messageMapper: messageMapperMock,
			},
			args: args{
				address: address,
			},
			want: true,
		},
		{
			name: "remove handlers without error",
			fields: fields{
				handlers:      blockHandlersMock,
				messageMapper: messageMapperMock,
			},
			args: args{
				address: address,
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &cosignatureHandler{
				messageMapper: tt.fields.messageMapper,
				handlers:      tt.fields.handlers,
			}
			got := h.Handle(tt.args.address, tt.args.resp)
			assert.Equal(t, got, tt.want)
		})
	}
}
