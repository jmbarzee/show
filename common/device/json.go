package device

import (
	"encoding/json"
	"fmt"

	"github.com/jmbarzee/show/common"
)

type ErrorDeviceTypeNotRegistered struct {
	nodeType string
}

func (e ErrorDeviceTypeNotRegistered) Error() string {
	return fmt.Sprintf("Failed to find registered device for UnmarshallingJSON: %s", e.nodeType)
}

type typed struct {
	Type string
}

type builder func() common.Device

var register = map[string]builder{}

func Register(db builder) {
	d := db()
	register[d.GetType()] = db
}

func UnmarshalJSON(data json.RawMessage) (common.Device, error) {
	temp := &typed{}

	if err := json.Unmarshal(data, temp); err != nil {
		return nil, err
	}

	db, ok := register[temp.Type]
	if !ok {
		fmt.Printf("%s", data)
		return nil, ErrorDeviceTypeNotRegistered{temp.Type}
	}

	d := db()

	err := d.UnmarshalJSON(data)
	return d, err
}

func UnmarshalJSONs(datas []json.RawMessage) ([]common.Device, error) {
	devices := make([]common.Device, len(datas))
	for i, data := range datas {
		device, err := UnmarshalJSON(data)
		if err != nil {
			return nil, err
		}

		devices[i] = device
	}
	return devices, nil
}
