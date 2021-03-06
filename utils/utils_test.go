package utils

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/stretchr/testify/assert"
)

func TestPanicOnError(t *testing.T) {
	assert.Panics(t, func() { PanicOnError(fmt.Errorf("some error")) })
	assert.NotPanics(t, func() { PanicOnError(nil) })
}

func TestMax(t *testing.T) {
	v1 := uint32(0)
	v2 := uint32(3418911847)
	v3 := uint32(356309450)

	assert.Equal(t, Max(v1), v1)
	assert.Equal(t, Max(v2), v2)
	assert.Equal(t, Max(v3), v3)

	assert.Equal(t, Max(v1, v1), v1)
	assert.Equal(t, Max(v1, v2), v2)
	assert.Equal(t, Max(v1, v3), v3)
	assert.Equal(t, Max(v2, v1), v2)
	assert.Equal(t, Max(v2, v2), v2)
	assert.Equal(t, Max(v2, v3), v2)
	assert.Equal(t, Max(v3, v1), v3)
	assert.Equal(t, Max(v3, v2), v2)
	assert.Equal(t, Max(v3, v3), v3)

	assert.Equal(t, Max(v1, v2, v3), v2)
	assert.Equal(t, Max(v1, v2, v3, v1), v2)
}

func TestXpubToNetwork(t *testing.T) {
	assert.Equal(t, XpubToNetwork("xpub6C774QqLVXvX3WBMACHRVdWTyPphFh45cXFvawg9eFuNAK2DNPsWDf1zJcSyZWY59FNspYUCAUJJXhmVzCPcWzLWDm6yEQSN9982pBAsj1k"), Mainnet)

	assert.Equal(t, XpubToNetwork("tpubDC5s7LsM3QFZz8CKNz8ePa2wpvQiq5LsGXrkoaaGsLhNx44wTr13XqoKEMCFPWMK4yen2DsLN7ArrZuqRqQE24Y9kNN51bpcjNdbWpJngdG"), Testnet)

	assert.Panics(t, func() { XpubToNetwork("foobar") })
}

func TestAddressToNetwork(t *testing.T) {
	assert.Equal(t, AddressToNetwork("19YomTTzGd55JM18pmj6Vv2F7ZqkaQDnRF"), Mainnet)
	assert.Equal(t, AddressToNetwork("3DmcpZprPpPLFsBsuMeGTik11DyQVsadQK"), Mainnet)

	assert.Equal(t, AddressToNetwork("mm8xEm6YS8B7ErLYYqcdF6URWkS1BWnqtY"), Testnet)
	assert.Equal(t, AddressToNetwork("2MvmkK3F4vT2h3gLjxz66SwQ5zW5XbsdZLu"), Testnet)
	assert.Equal(t, AddressToNetwork("n3s7pVRvCEuXfF5fyh74JXmYg45q4Wev86"), Testnet)

	assert.Panics(t, func() { AddressToNetwork("foobar") })
}

func TestChainConfig(t *testing.T) {
	assert.Equal(t, &chaincfg.MainNetParams, Mainnet.ChainConfig())
	assert.Equal(t, &chaincfg.TestNet3Params, Testnet.ChainConfig())
}

func TestGenesisBlock(t *testing.T) {
	assert.Equal(t, "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f", GenesisBlock(Mainnet))
	assert.Equal(t, "000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943", GenesisBlock(Testnet))
}

func TestVerifyMandN(t *testing.T) {
	assert.Error(t, VerifyMandN(0, 10))
	assert.Error(t, VerifyMandN(12, 10))
	assert.Error(t, VerifyMandN(5, 30))

	assert.Nil(t, VerifyMandN(1, 1))
	assert.Nil(t, VerifyMandN(1, 2))
	assert.Nil(t, VerifyMandN(2, 4))
	assert.Nil(t, VerifyMandN(5, 20))
	assert.Nil(t, VerifyMandN(20, 20))
}

func TestGetDefaultServer(t *testing.T) {
	host, port := GetDefaultServer(Testnet, Electrum, "foobar:s1234")
	assert.Equal(t, "foobar", host)
	assert.Equal(t, "s1234", port)

	host, port = GetDefaultServer(Testnet, Electrum, "192.0.2.5:s1234")
	assert.Equal(t, "192.0.2.5", host)
	assert.Equal(t, "s1234", port)

	host, port = GetDefaultServer(Testnet, Electrum, "[2001:db8::1]:s1234")
	assert.Equal(t, "2001:db8::1", host)
	assert.Equal(t, "s1234", port)

	host, port = GetDefaultServer(Testnet, Btcd, "foobar:1234")
	assert.Equal(t, "foobar", host)
	assert.Equal(t, "1234", port)

	host, port = GetDefaultServer(Testnet, Electrum, "")
	assert.NotEqual(t, "localhost", host)
	assert.Equal(t, "s53012", port)

	host, port = GetDefaultServer(Testnet, Btcd, "")
	assert.Equal(t, "localhost", host)
	assert.Equal(t, "18334", port)
}
