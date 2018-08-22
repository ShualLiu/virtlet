// +build linux

/*
Copyright 2018 Mirantis

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

package nettools

import (
	"github.com/vishvananda/netlink"
	"syscall"
)

// these constants are only available on Linux
const (
	FAMILY_ALL     = netlink.FAMILY_ALL
	FAMILY_V4      = netlink.FAMILY_V4
	RTPROT_KERNEL  = syscall.RTPROT_KERNEL
	SCOPE_LINK     = netlink.SCOPE_LINK
	SCOPE_UNIVERSE = netlink.SCOPE_UNIVERSE
)

// there's bug in netlink_unspecified.go in netlink version we use
// that breaks non-Linux builds
func linkSetMaster(link netlink.Link, master *netlink.Bridge) error {
	return netlink.LinkSetMaster(link, master)
}