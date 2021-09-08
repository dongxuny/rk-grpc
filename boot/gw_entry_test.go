// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package rkgrpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rookie-ninja/rk-entry/entry"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestNewGwEntry_WithoutOptions(t *testing.T) {
	entry := NewGwEntry()
	assert.NotNil(t, entry)

	assert.Equal(t, GwEntryNameDefault, entry.EntryName)
	assert.Equal(t, GwEntryType, entry.EntryType)
	assert.Equal(t, GwEntryDescription, entry.EntryDescription)
	assert.NotNil(t, entry.ZapLoggerEntry)
	assert.NotNil(t, entry.EventLoggerEntry)
	assert.Empty(t, entry.RegFuncsGw)
	assert.Empty(t, entry.GrpcDialOptions)
	assert.Empty(t, entry.ServerMuxOptions)
	assert.Empty(t, entry.GwMappingFilePaths)
	assert.Empty(t, entry.GwMapping)
}

func TestNewGwEntry_HappyCase(t *testing.T) {
	entryName := "ut-entry-name"
	gwMappingFilePath := "/ut-gw-mapping-fake.yaml"
	zapLoggerEntry := rkentry.NoopZapLoggerEntry()
	eventLoggerEntry := rkentry.NoopEventLoggerEntry()
	httpPort := uint64(80)
	grpcPort := uint64(1988)
	certEntry := &rkentry.CertEntry{}
	swEntry := &SwEntry{}
	tvEntry := &TvEntry{}
	promEntry := &PromEntry{}
	commonServiceEntry := &CommonServiceEntry{}
	grpcDialOption := grpc.WithBlock()

	entry := NewGwEntry(
		WithNameGw(entryName),
		WithGwMappingFilePathsGw(gwMappingFilePath),
		WithZapLoggerEntryGw(zapLoggerEntry),
		WithEventLoggerEntryGw(eventLoggerEntry),
		WithHttpPortGw(httpPort),
		WithGrpcPortGw(grpcPort),
		WithCertEntryGw(certEntry),
		WithSwEntryGw(swEntry),
		WithTvEntryGw(tvEntry),
		WithPromEntryGw(promEntry),
		WithCommonServiceEntryGw(commonServiceEntry),
		WithGrpcDialOptionsGw(grpcDialOption),
		WithRegFuncsGw(func(ctx context.Context, mux *runtime.ServeMux, s string, options []grpc.DialOption) error {
			return nil
		}),
	)
	assert.NotNil(t, entry)

	assert.Equal(t, entryName, entry.EntryName)
	assert.Equal(t, GwEntryType, entry.EntryType)
	assert.Equal(t, GwEntryDescription, entry.EntryDescription)
	assert.Equal(t, zapLoggerEntry, entry.ZapLoggerEntry)
	assert.Equal(t, eventLoggerEntry, entry.EventLoggerEntry)
	assert.Equal(t, httpPort, entry.HttpPort)
	assert.Equal(t, grpcPort, entry.GrpcPort)
	assert.Equal(t, certEntry, entry.CertEntry)
	assert.Equal(t, swEntry, entry.SwEntry)
	assert.Equal(t, tvEntry, entry.TvEntry)
	assert.Equal(t, promEntry, entry.PromEntry)
	assert.Equal(t, commonServiceEntry, entry.CommonServiceEntry)
	assert.NotEmpty(t, entry.GrpcDialOptions)
	assert.NotEmpty(t, entry.GrpcDialOptions)
	assert.NotEmpty(t, entry.GwMappingFilePaths)
}

func TestGwEntry_GetName_HappyCase(t *testing.T) {
	entryName := "ut-entry-name"

	entry := NewGwEntry(WithNameGw(entryName))
	assert.NotNil(t, entry)

	assert.Equal(t, entryName, entry.GetName())
}

func TestGwEntry_GetType_HappyCase(t *testing.T) {
	entry := NewGwEntry()
	assert.Equal(t, GwEntryType, entry.GetType())
}

func TestGwEntry_GetDescription_HappyCase(t *testing.T) {
	entry := NewGwEntry()
	assert.Equal(t, GwEntryDescription, entry.GetDescription())
}

func TestGwEntry_UnmarshalJSON(t *testing.T) {
	entry := NewGwEntry()
	assert.Nil(t, entry.UnmarshalJSON(nil))
}

func TestGwEntry_MarshalJSON_HappyCase(t *testing.T) {
	entry := NewGwEntry()
	bytes, err := entry.MarshalJSON()

	assert.Nil(t, err)
	assert.NotEmpty(t, bytes)

	str := string(bytes)

	assert.Contains(t, str, "entryName")
	assert.Contains(t, str, "entryType")
	assert.Contains(t, str, "grpcPort")
	assert.Contains(t, str, "httpPort")
	assert.Contains(t, str, "zapLoggerEntry")
	assert.Contains(t, str, "eventLoggerEntry")
	assert.Contains(t, str, "swEnabled")
	assert.Contains(t, str, "tvEnabled")
	assert.Contains(t, str, "promEnabled")
	assert.Contains(t, str, "commonServiceEnabled")
	assert.Contains(t, str, "grpcTlsEnabled")
	assert.Contains(t, str, "serverTlsEnabled")
}

func TestGwEntry_String_HappyCase(t *testing.T) {
	entry := NewGwEntry()
	str := entry.String()
	assert.Contains(t, str, "entryName")
	assert.Contains(t, str, "entryType")
	assert.Contains(t, str, "grpcPort")
	assert.Contains(t, str, "httpPort")
	assert.Contains(t, str, "zapLoggerEntry")
	assert.Contains(t, str, "eventLoggerEntry")
	assert.Contains(t, str, "swEnabled")
	assert.Contains(t, str, "tvEnabled")
	assert.Contains(t, str, "promEnabled")
	assert.Contains(t, str, "commonServiceEnabled")
	assert.Contains(t, str, "grpcTlsEnabled")
	assert.Contains(t, str, "serverTlsEnabled")
}

func TestGwEntry_IsSwEnabled_ExpectTrue(t *testing.T) {
	swEntry := NewSwEntry()
	entry := NewGwEntry(WithSwEntryGw(swEntry))

	assert.True(t, entry.IsSwEnabled())
}

func TestGwEntry_IsSwEnabled_ExpectFalse(t *testing.T) {
	entry := NewGwEntry()

	assert.False(t, entry.IsSwEnabled())
}

func TestGwEntry_IsTvEnabled_ExpectTrue(t *testing.T) {
	tvEntry := NewTvEntry()
	entry := NewGwEntry(WithTvEntryGw(tvEntry))

	assert.True(t, entry.IsTvEnabled())
}

func TestGwEntry_IsTvEnabled_ExpectFalse(t *testing.T) {
	entry := NewGwEntry()

	assert.False(t, entry.IsTvEnabled())
}

func TestGwEntry_IsPromEnabled_ExpectTrue(t *testing.T) {
	promEntry := NewPromEntry()
	entry := NewGwEntry(WithPromEntryGw(promEntry))

	assert.True(t, entry.IsPromEnabled())
}

func TestGwEntry_IsPromEnabled_ExpectFalse(t *testing.T) {
	entry := NewGwEntry()

	assert.False(t, entry.IsPromEnabled())
}

func TestGwEntry_IsCommonServiceEnabled_ExpectTrue(t *testing.T) {
	commonServiceEntry := NewCommonServiceEntry()
	entry := NewGwEntry(WithCommonServiceEntryGw(commonServiceEntry))

	assert.True(t, entry.IsCommonServiceEnabled())
}

func TestGwEntry_IsCommonServiceEnabled_ExpectFalse(t *testing.T) {
	entry := NewGwEntry()

	assert.False(t, entry.IsCommonServiceEnabled())
}

func TestGwEntry_IsClientTlsEnabled_ExpectFalse(t *testing.T) {
	// Without client cert
	certEntry := &rkentry.CertEntry{
		Store: &rkentry.CertStore{},
	}

	entry := NewGwEntry(WithCertEntryGw(certEntry))
	assert.False(t, entry.IsGrpcTlsEnabled())

	// Without Store
	certEntry = &rkentry.CertEntry{}

	entry = NewGwEntry(WithCertEntryGw(certEntry))
	assert.False(t, entry.IsGrpcTlsEnabled())

	// Without cert entry
	entry = NewGwEntry()
	assert.False(t, entry.IsGrpcTlsEnabled())
}

func TestGwEntry_IsServerTlsEnabled_ExpectTrue(t *testing.T) {
	certEntry := &rkentry.CertEntry{
		Store: &rkentry.CertStore{
			ServerCert: []byte("ut-server-cert"),
		},
	}

	entry := NewGwEntry(WithCertEntryGw(certEntry))
	assert.True(t, entry.IsServerTlsEnabled())
}

func TestGwEntry_IsServerTlsEnabled_ExpectFalse(t *testing.T) {
	// Without client cert
	certEntry := &rkentry.CertEntry{
		Store: &rkentry.CertStore{},
	}

	entry := NewGwEntry(WithCertEntryGw(certEntry))
	assert.False(t, entry.IsServerTlsEnabled())

	// Without Store
	certEntry = &rkentry.CertEntry{}

	entry = NewGwEntry(WithCertEntryGw(certEntry))
	assert.False(t, entry.IsServerTlsEnabled())

	// Without cert entry
	entry = NewGwEntry()
	assert.False(t, entry.IsServerTlsEnabled())
}
