// +build cgo

/**
 * Copyright (c) 2020-2021, Intel Corporation.
 * SPDX-License-Identifier: BSD-3-Clause
 **
 * This package introduce wrapper for ipmctl library written in C.
 * api_parser.go file contains all functions bounded to ipmctl library
 * C types wrapped in GO by api_types.go file. All these functions are used
 * for parsing between exposed types
 */

package nvm

// #cgo pkg-config: libipmctl
// #include <include/nvm_management.h>
import "C"

import (
	"strconv"
)

func strGo2C(str string, length uint) []C.char {
	result := make([]C.char, length)
	for i, c := range str {
		result[i] = C.char(c)
	}
	// terminate string
	result[len(str)] = 0
	return result
}

func bytesToString(value []nvmUint8) string {
	result := "0x"
	for _, v := range value {
		result += strconv.FormatUint(uint64(v), 16)
	}
	return result
}

func uint16ToString(value []nvmUint16) string {
	result := "0x"
	for _, v := range value {
		result += strconv.FormatUint(uint64(v), 16)
	}
	return result
}

func (uid *nvmUID) toCharArray() []C.char {
	uidStr := string(*uid)
	return strGo2C(uidStr, nvmMaxUIDLen)
}

func toString(value uint64, base int) string {
	prefix := ""
	if 16 == base {
		prefix = "0x"
	} else if 2 == base {
		prefix = "0b"
	}
	return prefix + strconv.FormatUint(value, base)
}

func (value *nvmBool) toString(base int) string {
	if *value {
		return toString(uint64(1), base)
	}
	return toString(uint64(0), base)
}

func (value *nvmUint8) toString(base int) string {
	return toString(uint64(*value), base)
}

func (value *nvmUint16) toString(base int) string {
	return toString(uint64(*value), base)
}

func (value *nvmUint32) toString(base int) string {
	return toString(uint64(*value), base)
}

func (value *nvmUint64) toString(base int) string {
	return toString(uint64(*value), base)
}

func (value *enumAttr) toUint32() uint32 {
	valueUint32 := uint32(*value)
	return valueUint32
}

func (value *nvmBool) toNvmUint64() nvmUint64 {
	if bool(*value) {
		return nvmUint64(1)
	}
	return nvmUint64(0)
}
