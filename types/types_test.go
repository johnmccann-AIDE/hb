package types

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

// HelperLoadBytes allows you to use relative path testdata directory as a place
// to load and store your data
func HelperLoadBytes(tb testing.TB, name string) []byte {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)     // nolint : gosec
	if err != nil {
		tb.Fatal(err)
	}
	return bytes
}

func TestDeviceYamlParsing(t *testing.T) {
	var devices Devices
	err := devices.Parse(HelperLoadBytes(t, "./devices/devices.yml"))
	assert.Nil(t, err, "Failed to parse yaml representation of Devices")
	assert.Len(t, devices.Device, 3, "Expected to parse 3 Devices")
	assert.EqualValues(t, "4200_1", devices.Device[0].DeviceID, "Yaml type with a hyphen, didn't decode correctly")
}

func TestDeviceOmit(t *testing.T) {
	var devices Devices
	_ = devices.Parse(HelperLoadBytes(t, "./devices/devices.yml"))
	maxDevice, err := yaml.Marshal(devices.Device[0])
	if err != nil {
		assert.Nil(t, err, "Failed to marshal devices to json")
	}
	assert.NotContains(t, string(maxDevice), "cisco", "cisco should be ignored")

	minDevice, err := json.Marshal(devices.Device[2])
	if err != nil {
		assert.Nil(t, err, "Failed to marshal devices to json")
	}
	assert.NotContains(t, string(minDevice), "authentication", "Optional Authentication type was not ignored")
	assert.NotContains(t, string(minDevice), "iAgent", "Optional NETCONF type was not ignored")
	assert.NotContains(t, string(minDevice), "open-config", "Optional Open Config type was not ignored")
	assert.NotContains(t, string(minDevice), "snmp", "Optional SNMP type was not ignored")
	assert.NotContains(t, string(minDevice), "vendor", "Optional vendor type was not ignored")
	partialDevice, err := json.Marshal(devices.Device[1])
	if err != nil {
		assert.Nil(t, err, "Failed to marshal devices to json")
	}
	assert.NotContains(t, string(partialDevice), "v2", "Optional Community type was not ignored")
	assert.NotContains(t, string(partialDevice), "juniper", "Optional Vendor juniper was not ignored")
}

func TestDeviceGroupYamlParsing(t *testing.T) {
	var deviceGroups DeviceGroups
	err := deviceGroups.Parse(HelperLoadBytes(t, "./device-groups/deviceGroups.yml"))
	assert.Nil(t, err, "Failed to parse yaml representation of DeviceGroups")
	assert.Len(t, deviceGroups.DeviceGroup, 2, "Expected to parse 2 DeviceGroups")
	assert.EqualValues(t, "l2-test-group", deviceGroups.DeviceGroup[0].DeviceGroupName, "Yaml type with a hyphen, didn't decode correctly")
}

func TestDeviceGroupOmit(t *testing.T) {
	var deviceGroups DeviceGroups
	_ = deviceGroups.Parse(HelperLoadBytes(t, "./device-groups/deviceGroups.yml"))
	minDevice, err := json.Marshal(deviceGroups.DeviceGroup[1])
	if err != nil {
		assert.Nil(t, err, "Failed to marshal devicegroups to json")
	}
	assert.NotContains(t, string(minDevice), "description", "Optional Description type was not ignored")
	assert.NotContains(t, string(minDevice), "devices", "Optional Devices type was not ignored")
	assert.NotContains(t, string(minDevice), "playbooks", "Optional Playbooks type was not ignored")
	assert.NotContains(t, string(minDevice), "authentication", "Optional Authentication type was not ignored")
	assert.NotContains(t, string(minDevice), "native-gpb", "Optional Native GPB type was not ignored")
}

func TestPlaybookInstanceYamlParsing(t *testing.T) {
	var playbookInstances PlaybookInstances
	err := playbookInstances.Parse(HelperLoadBytes(t, "./playbook-instances/playbook-instances.yml"))
	assert.Nil(t, err, "Failed to parse yaml representation of Playbook Instances (DeviceGroups)")
	assert.Len(t, playbookInstances.DeviceGroup, 1, "Expected to parse 1 Instances")
	assert.EqualValues(t, "ptp-test-group", playbookInstances.DeviceGroup[0].DeviceGroupName, "Yaml type with a hyphen, didn't decode correctly")
}

func TestPlaybooksYamlParsing(t *testing.T) {
	var playbooks Playbooks
	err := playbooks.Parse(HelperLoadBytes(t, "./playbooks/playbooks.yml"))
	assert.Nil(t, err, "Failed to parse yaml representation of Playbooks")
	assert.Len(t, playbooks.Playbooks, 1, "Expected to parse 1 playbook")
	assert.EqualValues(t, "interface-status-test", playbooks.Playbooks[0].PlayBookName, "Yaml type with a hyphen, didn't decode correctly")
}
