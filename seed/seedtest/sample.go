// -*- Mode: Go; indent-tabs-mode: t -*-

/*
 * Copyright (C) 2015-2019 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package seedtest

var SampleSnapYaml = map[string]string{
	"core": `name: core
type: os
version: 1.0
`,
	"pc-kernel": `name: pc-kernel
type: kernel
version: 1.0
`,
	"pc": `name: pc
type: gadget
version: 1.0
`,
	"classic-gadget": `name: classic-gadget
version: 1.0
type: gadget
`,
	"required": `name: required
type: app
version: 1.0
`,
	"classic-snap": `name: classic-snap
type: app
confinement: classic
version: 1.0
`,
	"snapd": `name: snapd
type: snapd
version: 1.0
`,
	"core18": `name: core18
type: base
version: 1.0
`,
	"pc-kernel=18": `name: pc-kernel
type: kernel
version: 1.0
`,
	"pc=18": `name: pc
type: gadget
base: core18
version: 1.0
`,
	"classic-gadget18": `name: classic-gadget18
version: 1.0
base: core18
type: gadget
`,
	"required18": `name: required18
type: app
base: core18
version: 1.0
`,
	"core20": `name: core20
type: base
version: 1.0
`,
	"pc-kernel=20": `name: pc-kernel
type: kernel
version: 1.0
`,
	"pc=20": `name: pc
type: gadget
base: core20
version: 1.0
`,
	"required20": `name: required20
type: app
base: core20
version: 1.0
components:
  comp1:
    type: standard
  comp2:
    type: standard
`,
	"required20+comp1": `component: required20+comp1
type: standard
version: 1.0
`,
	"required20+comp1_kernel": `component: required20+comp1
type: kernel-modules
version: 1.0
`,
	"required20+comp2": `component: required20+comp2
type: standard
version: 2.0
`,
	"required20+unknown": `component: required20+unknown
type: standard
version: 2.0
`,
	"aux-info-test": `name: aux-info-test
type: app
base: core20
version: 1.0
links:
  contact:
    - mailto:author@example.com
`,
	"component-test": `name: component-test
type: app
base: core20
version: 1.0
components:
  comp1:
    type: standard
  comp2:
    type: standard
  comp3:
    type: standard
`,
	"component-test+comp1": `component: component-test+comp1
type: standard
version: 1.0
`,
	"component-test+comp2": `component: component-test+comp2
type: standard
version: 2.0
`,
	"component-test+comp3": `component: component-test+comp3
type: standard
version: 2.0
`,
	"local-component-test": `name: local-component-test
type: app
base: core20
version: 1.0
components:
  comp4:
    type: standard
`,
	"local-component-test+comp4": `component: local-component-test+comp4
type: standard
version: 1.0
`,
	"optional20-a": `name: optional20-a
type: app
base: core20
version: 1.0
`,
	"optional20-b": `name: optional20-b
type: app
base: core20
version: 1.0`,
	"uboot-gadget=20": `name: uboot-gadget
type: gadget
base: core20
version: 1.0
`,
	"arm-kernel=20": `name: arm-kernel
type: kernel
version: 1.0
`,
	"test-devmode=20": `name: test-devmode
type: app
base: core20
version: 1.0
confinement: devmode
`,
	"core22": `name: core22
type: base
version: 1.0
`,
	"core24": `name: core24
type: base
version: 1.0
`,
	"pc-kernel=22": `name: pc-kernel
type: kernel
version: 1.0
`,
	"pc-kernel=24": `name: pc-kernel
type: kernel
version: 1.0
`,
	"pc-kernel=22+kmods": `name: pc-kernel
type: kernel
version: 1.0
components:
  kcomp1:
    type: kernel-modules
  kcomp2:
    type: kernel-modules
`,
	"pc-kernel=24+kmods": `name: pc-kernel
type: kernel
version: 1.0
components:
  kcomp1:
    type: kernel-modules
  kcomp2:
    type: kernel-modules
  kcomp3:
    type: kernel-modules
`,
	"pc-kernel+kcomp1": `component: pc-kernel+kcomp1
type: kernel-modules
version: 1.0
`,
	"pc-kernel+kcomp2": `component: pc-kernel+kcomp2
type: kernel-modules
version: 1.0
`,
	"pc-kernel+kcomp3": `component: pc-kernel+kcomp3
type: kernel-modules
version: 1.0
`,
	"pc=22": `name: pc
type: gadget
base: core22
version: 1.0
`,
	"pc=24": `name: pc
type: gadget
base: core24
version: 1.0
`,
	"optional22": `name: optional22
type: app
base: core22
version: 1.0
components:
  comp1:
    type: standard
`,
	"optional22+comp1": `component: optional22+comp1
type: standard
version: 1.0
`,
	"optional24": `name: optional24
type: app
base: core24
version: 1.0
components:
  comp1:
    type: standard
`,
	"optional24+comp1": `component: optional24+comp1
type: standard
version: 1.0
`,
}

func MergeSampleSnapYaml(snapYaml ...map[string]string) map[string]string {
	if len(snapYaml) == 0 {
		return nil
	}
	merged := make(map[string]string, len(snapYaml[0]))
	for _, m := range snapYaml {
		for yamlKey, yaml := range m {
			merged[yamlKey] = yaml
		}
	}
	return merged
}
