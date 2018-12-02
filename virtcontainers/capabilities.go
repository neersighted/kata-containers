// Copyright (c) 2017 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package virtcontainers

const (
	blockDeviceSupport = 1 << iota
	blockDeviceHotplugSupport
	multiQueueSupport
	fsSharingUnsupported
)

type capabilities struct {
	flags uint
}

func (caps *capabilities) isBlockDeviceSupported() bool {
	if caps.flags&blockDeviceSupport != 0 {
		return true
	}
	return false
}

func (caps *capabilities) setBlockDeviceSupport() {
	caps.flags = caps.flags | blockDeviceSupport
}

func (caps *capabilities) isBlockDeviceHotplugSupported() bool {
	if caps.flags&blockDeviceHotplugSupport != 0 {
		return true
	}
	return false
}

func (caps *capabilities) setBlockDeviceHotplugSupport() {
	caps.flags |= blockDeviceHotplugSupport
}

func (caps *capabilities) isMultiQueueSupported() bool {
	if caps.flags&multiQueueSupport != 0 {
		return true
	}
	return false
}

func (caps *capabilities) setMultiQueueSupport() {
	caps.flags |= multiQueueSupport
}

func (caps *capabilities) isFsSharingSupported() bool {
	return caps.flags&fsSharingUnsupported == 0
}

func (caps *capabilities) setFsSharingUnsupported() {
	caps.flags |= fsSharingUnsupported
}
