//
// SPDX-License-Identifier: BSD-3-Clause

package redfish

import (
	"encoding/json"
	"reflect"
	"strconv"
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ciferlu1024/gofish/common"
)

// InputType is the type of power input.
type InputType string

const (

	// ACInputType Alternating Current (AC) input range.
	ACInputType InputType = "AC"
	// DCInputType Direct Current (DC) input range.
	DCInputType InputType = "DC"
)

// LineInputVoltageType is the type of line voltage.
type LineInputVoltageType string

const (

	// UnknownLineInputVoltageType The power supply line input voltage type
	// cannot be determined.
	UnknownLineInputVoltageType LineInputVoltageType = "Unknown"
	// ACLowLineLineInputVoltageType 100-127V AC input.
	ACLowLineLineInputVoltageType LineInputVoltageType = "ACLowLine"
	// ACMidLineLineInputVoltageType 200-240V AC input.
	ACMidLineLineInputVoltageType LineInputVoltageType = "ACMidLine"
	// ACHighLineLineInputVoltageType 277V AC input.
	ACHighLineLineInputVoltageType LineInputVoltageType = "ACHighLine"
	// DCNeg48VLineInputVoltageType -48V DC input.
	DCNeg48VLineInputVoltageType LineInputVoltageType = "DCNeg48V"
	// DC380VLineInputVoltageType High Voltage DC input (380V).
	DC380VLineInputVoltageType LineInputVoltageType = "DC380V"
	// AC120VLineInputVoltageType AC 120V nominal input.
	AC120VLineInputVoltageType LineInputVoltageType = "AC120V"
	// AC240VLineInputVoltageType AC 240V nominal input.
	AC240VLineInputVoltageType LineInputVoltageType = "AC240V"
	// AC277VLineInputVoltageType AC 277V nominal input.
	AC277VLineInputVoltageType LineInputVoltageType = "AC277V"
	// ACandDCWideRangeLineInputVoltageType Wide range AC or DC input.
	ACandDCWideRangeLineInputVoltageType LineInputVoltageType = "ACandDCWideRange"
	// ACWideRangeLineInputVoltageType Wide range AC input.
	ACWideRangeLineInputVoltageType LineInputVoltageType = "ACWideRange"
	// DC240VLineInputVoltageType DC 240V nominal input.
	DC240VLineInputVoltageType LineInputVoltageType = "DC240V"
)

// PowerLimitException is the type of power limit exception.
type PowerLimitException string

const (

	// NoActionPowerLimitException Take no action when the limit is exceeded.
	NoActionPowerLimitException PowerLimitException = "NoAction"
	// HardPowerOffPowerLimitException Turn the power off immediately when
	// the limit is exceeded.
	HardPowerOffPowerLimitException PowerLimitException = "HardPowerOff"
	// LogEventOnlyPowerLimitException Log an event when the limit is
	// exceeded, but take no further action.
	LogEventOnlyPowerLimitException PowerLimitException = "LogEventOnly"
	// OemPowerLimitException Take an OEM-defined action.
	OemPowerLimitException PowerLimitException = "Oem"
)

// PowerSupplyType is the type of power supply.
type PowerSupplyType string

const (

	// UnknownPowerSupplyType The power supply type cannot be determined.
	UnknownPowerSupplyType PowerSupplyType = "Unknown"
	// ACPowerSupplyType Alternating Current (AC) power supply.
	ACPowerSupplyType PowerSupplyType = "AC"
	// DCPowerSupplyType Direct Current (DC) power supply.
	DCPowerSupplyType PowerSupplyType = "DC"
	// ACorDCPowerSupplyType Power Supply supports both DC or AC.
	ACorDCPowerSupplyType PowerSupplyType = "ACorDC"
)

// InputRange shall describe an input range that the associated power supply is
// able to utilize.
type InputRange struct {
	// InputType shall contain the input type (AC or DC) of the associated range.
	InputType InputType
	// MaximumFrequencyHz shall contain the value in Hertz of the maximum line
	// input frequency which the power supply is capable of consuming for this range.
	MaximumFrequencyHz float64
	// MaximumVoltage shall contain the value in Volts of the maximum line input
	// voltage which the power supply is capable of consuming for this range.
	MaximumVoltage float64
	// MinimumFrequencyHz shall contain the value in Hertz of the minimum line
	// input frequency which the power supply is capable of consuming for this range.
	MinimumFrequencyHz float64
	// MinimumVoltage shall contain the value in Volts of the minimum line input
	// voltage which the power supply is capable of consuming for this range.
	MinimumVoltage float64
	// OutputWattage shall contain the maximum amount of power, in Watts, that
	// the associated power supply is rated to deliver while operating in this input range.
	OutputWattage float64
}

// Power is used to represent a power metrics resource for a Redfish
// implementation.
type Power struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IndicatorLED shall contain the indicator light state for the indicator
	// light associated with this power supply.
	IndicatorLED common.IndicatorLED
	// PowerControl shall be the definition for power control (power reading and
	// limiting) for a Redfish implementation.
	PowerControl []PowerControl
	// PowerControlCount is the number of objects.
	PowerControlCount int `json:"PowerControl@odata.count"`
	// PowerSupplies shall contain details of the power supplies associated with
	// this system or device.
	PowerSupplies []PowerSupply
	// PowerSuppliesCount is the number of objects.
	PowerSuppliesCount int `json:"PowerSupplies@odata.count"`
	// Redundancy shall contain redundancy information for the power subsystem
	// of this system or device.
	Redundancy []Redundancy
	// RedundancyCount is the number of objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Voltages shall be the definition for voltage
	// sensors for a Redfish implementation.
	Voltages []Voltage
	// VoltagesCount is the number of objects.
	VoltagesCount int `json:"Voltages@odata.count"`
}

// GetPower will get a Power instance from the service.
func GetPower(c common.Client, uri string) (*Power, error) {
	fmt.Println("******************power.go getpower", uri)
	resp, err := c.Get(uri)
	if err != nil {
		fmt.Println("*********************power.go getpower get 报错！", err)
		return nil, err
	}else{
		fmt.Println("*********************power.go getpower get 没有报错！")
	}
	defer resp.Body.Close()

	// os.Stdout 输出原始json内容!
        mybodys, _ := ioutil.ReadAll(resp.Body)
        var out bytes.Buffer
        err = json.Indent(&out, mybodys, "", "\t")
        if err != nil {
                fmt.Println("**************************power.go json body 报错!", err)
        }else{
                fmt.Println("**************************power.go json body: 已获取\n")
        }
        //out.WriteTo(os.Stdout)

        file, _ := os.Create("/tmp/powerjson.txt")
        defer file.Close()
        out.WriteTo(file)

        // 读取json文件获取json数据
        jsonFile, err := os.Open("/tmp/powerjson.txt")
        if err != nil {
                fmt.Println("error opening power json file")
        }else{
                fmt.Println("已打开power json文件")
        }

        defer jsonFile.Close()
        jsonData, err := ioutil.ReadAll(jsonFile)
        if err!= nil {
                fmt.Println("error reading power json file")
        }else{
                fmt.Println("已读取power json数据")
        }

        // 重新解析json数据

        var r interface{}
        err = json.Unmarshal(jsonData, &r)
        // fmt.Println("r的值：", r)

        // 修改json数据部分字段的格式
        newbodymap, _ := r.(map[string]interface{})

//	var a float64 = 0
//	var b string = "0"
//	newbodymap["PowerControl"].(map[string]interface{})["PowerCapacityWatts"] = a
//	newbodymap["PowerControl"].(map[string]interface{})["PowerLimit"].(map[string]interface{})["LimitInWatts"] = a
//	newbodymap["PowerControl"].(map[string]interface{})["PowerLimit"].(map[string]interface{})["LimitException"] = b
//	newbodymap["PowerControl"].(map[string]interface{})["PowerConsumedWatts"] = strconv.FormatFloat(newbodymap["PowerControl"].(map[string]interface{})["PowerConsumedWatts"].(float64), 'f', -1, 64)

//	delete(newbodymap["PowerControl"].(map[string]interface{}), "PowerConsumedWatts")
//	delete(newbodymap["PowerControl"].(map[string]interface{}), "PowerLimit")
//	delete(newbodymap["PowerControl"].(map[string]interface{}), "PowerCapacityWatts")
	delete(newbodymap, "PowerControl")
	//fmt.Printf("powercontrol的值:%v , 类型:%T \n", newbodymap["PowerControl"], newbodymap["PowerControl"])


        newbodyjson, err := json.Marshal(newbodymap)
        if err != nil {
                fmt.Println("*************newbodyjson err:", err)
        }

	var power Power
	var newjsonreader io.Reader
	newjsonreader = strings.NewReader(string(newbodyjson))
	err = json.NewDecoder(newjsonreader).Decode(&power)
	//err = json.NewDecoder(resp.Body).Decode(&power)
	if err != nil {
		return nil, err
	}

	power.SetClient(c)
	return &power, nil
}

// ListReferencedPowers gets the collection of Power from
// a provided reference.
func ListReferencedPowers(c common.Client, link string) ([]*Power, error) { //nolint:dupl
	var result []*Power
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		fmt.Println("power.go ListReferencedPowers getcollection 有报错！")
		return result, err
	}else{
		fmt.Println("power.go ListReferencedPowers getcollection 没有错！")
	}

	collectionError := common.NewCollectionError()
	for _, powerLink := range links.ItemLinks {
		power, err := GetPower(c, powerLink)
		if err != nil {
			collectionError.Failures[powerLink] = err
		} else {
			result = append(result, power)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// PowerControl is
type PowerControl struct {
	common.Entity

	// MemberID shall uniquely identify the member within the collection. For
	// services supporting Redfish v1.6 or higher, this value shall be the
	// zero-based array index.
	MemberID string `json:"MemberId"`
	// PhysicalContext shall be a description of the affected device(s) or region
	// within the chassis to which this power control applies.
	PhysicalContext common.PhysicalContext
	// PowerAllocatedWatts shall represent the total power currently allocated
	// to chassis resources.
	PowerAllocatedWatts float64
	// PowerAvailableWatts shall represent the amount of power capacity (in
	// Watts) not already allocated and shall equal PowerCapacityWatts -
	// PowerAllocatedWatts.
	PowerAvailableWatts float64
	// PowerCapacityWatts shall represent the total power capacity that is
	// available for allocation to the chassis resources.
	PowerCapacityWatts float64
	// PowerConsumedWatts shall represent the actual power being consumed (in
	// Watts) by the chassis.
	PowerConsumedWatts float64
	// PowerLimit shall contain power limit status and configuration information
	// for this chassis.
	PowerLimit PowerLimit
	// PowerMetrics shall contain power metrics for power readings (interval,
	// minimum/maximum/average power consumption) for the chassis.
	PowerMetrics PowerMetric
	// PowerRequestedWatts shall represent the
	// amount of power (in Watts) that the chassis resource is currently
	// requesting be budgeted to it for future use.
	PowerRequestedWatts float64
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a PowerControl object from the raw JSON.
func (powercontrol *PowerControl) UnmarshalJSON(b []byte) error { // nolint:dupl
	type temp PowerControl
	type t1 struct {
		temp
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		fmt.Println("*******power.go UnmarshalJSON powercontrol 解析有报错！")
		// See if we need to handle converting MemberID
		var t2 struct {
			t1
			MemberID int `json:"MemberId"`
		}
		err2 := json.Unmarshal(b, &t2)

		if err2 != nil {
			// Return the original error
			return err
		}

		// Convert the numeric member ID to a string
		t = t2.t1
		t.temp.MemberID = strconv.Itoa(t2.MemberID)
		fmt.Println("*****power.go powercontrol 解析结果: ", t.temp)
	}

	// Extract the links to other entities for later
	*powercontrol = PowerControl(t.temp)

	return nil
}

// PowerLimit shall contain power limit status and
// configuration information for this chassis.

type PowerLimit struct {
	// CorrectionInMs shall represent the time
	// interval in ms required for the limiting process to react and reduce
	// the power consumption below the limit.
	CorrectionInMs int64
	// LimitException shall represent the
	// action to be taken if the resource power consumption can not be
	// limited below the specified limit after several correction time
	// periods.
	LimitException PowerLimitException
	// LimitInWatts shall represent the power
	// cap limit in watts for the resource. If set to null, power capping
	// shall be disabled.
	LimitInWatts float64
}

// PowerMetric shall contain power metrics for power
// readings (interval, minimum/maximum/average power consumption) for a
// resource.
type PowerMetric struct {
	// AverageConsumedWatts shall represent the
	// average power level that occurred averaged over the last IntervalInMin
	// minutes.
	AverageConsumedWatts float64
	// IntervalInMin shall represent the time
	// interval (or window), in minutes, in which the PowerMetrics properties
	// are measured over.
	// Should be an integer, but some Dell implementations return as a float.
	IntervalInMin float64
	// MaxConsumedWatts shall represent the
	// maximum power level in watts that occurred within the last
	// IntervalInMin minutes.
	MaxConsumedWatts float64
	// MinConsumedWatts shall represent the
	// minimum power level in watts that occurred within the last
	// IntervalInMin minutes.
	MinConsumedWatts float64
}

// PowerSupply is Details of a power supplies associated with this system
// or device.
type PowerSupply struct {
	common.Entity

	// assembly shall be a link to a resource of type Assembly.
	assembly string
	// EfficiencyPercent shall contain the value of the measured power
	// efficiency, as a percentage, of the associated power supply.
	EfficiencyPercent float64
	// FirmwareVersion shall contain the firmware version as
	// defined by the manufacturer for the associated power supply.
	FirmwareVersion string
	// HotPluggable shall indicate whether the
	// device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Devices indicated
	// as hot-pluggable shall allow the device to become operable without
	// altering the operational state of the underlying equipment. Devices
	// that cannot be inserted or removed from equipment in operation, or
	// devices that cannot become operable without affecting the operational
	// state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// IndicatorLED shall contain the indicator
	// light state for the indicator light associated with this power supply.
	IndicatorLED common.IndicatorLED
	// InputRanges shall be a collection of ranges usable by the power supply unit.
	InputRanges []InputRange
	// LastPowerOutputWatts shall contain the average power
	// output, measured in Watts, of the associated power supply.
	LastPowerOutputWatts float64
	// LineInputVoltage shall contain the value in Volts of
	// the line input voltage (measured or configured for) that the power
	// supply has been configured to operate with or is currently receiving.
	LineInputVoltage float64
	// LineInputVoltageType shall contain the type of input
	// line voltage supported by the associated power supply.
	LineInputVoltageType LineInputVoltageType
	// Location shall contain location information of the
	// associated power supply.
	Location common.Location
	// Manufacturer shall be the name of the
	// organization responsible for producing the power supply. This
	// organization might be the entity from whom the power supply is
	// purchased, but this is not necessarily true.
	Manufacturer string
	// MemberID shall uniquely identify the
	// member within the collection. For services supporting Redfish v1.6 or
	// higher, this value shall be the zero-based array index.
	MemberID string `json:"MemberId"`
	// Model shall contain the model information as defined
	// by the manufacturer for the associated power supply.
	Model string
	// PartNumber shall contain the part number as defined
	// by the manufacturer for the associated power supply.
	PartNumber string
	// PowerCapacityWatts shall contain the maximum amount
	// of power, in Watts, that the associated power supply is rated to
	// deliver.
	PowerCapacityWatts float64
	// PowerInputWatts shall contain the value of the
	// measured input power, in Watts, of the associated power supply.
	PowerInputWatts float64
	// PowerOutputWatts shall contain the value of the
	// measured output power, in Watts, of the associated power supply.
	PowerOutputWatts float64
	// PowerSupplyType shall contain the input power type
	// (AC or DC) of the associated power supply.
	PowerSupplyType PowerSupplyType
	// Redundancy is used to show redundancy for power supplies and other
	// elements in this resource. The use of IDs within these arrays shall
	// reference the members of the redundancy groups.
	Redundancy []Redundancy
	// RedundancyCount is the number of objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SerialNumber shall contain the serial number as
	// defined by the manufacturer for the associated power supply.
	SerialNumber string
	// SparePartNumber shall contain the spare or
	// replacement part number as defined by the manufacturer for the
	// associated power supply.
	SparePartNumber string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerSupply object from the raw JSON.
func (powersupply *PowerSupply) UnmarshalJSON(b []byte) error {
	type temp PowerSupply
	var t struct {
		temp
		Assembly common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*powersupply = PowerSupply(t.temp)
	powersupply.assembly = string(t.Assembly)

	// This is a read/write object, so we need to save the raw object data for later
	powersupply.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (powersupply *PowerSupply) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PowerSupply)
	err := original.UnmarshalJSON(powersupply.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(powersupply).Elem()

	return powersupply.Entity.Update(originalElement, currentElement, readWriteFields)
}

// Voltage is a voltage representation.
type Voltage struct {
	common.Entity

	// LowerThresholdCritical shall indicate
	// the present reading is below the normal range but is not yet fatal.
	// Units shall use the same units as the related ReadingVolts property.
	LowerThresholdCritical float64
	// LowerThresholdFatal shall indicate the
	// present reading is below the normal range and is fatal. Units shall
	// use the same units as the related ReadingVolts property.
	LowerThresholdFatal float64
	// LowerThresholdNonCritical shall indicate
	// the present reading is below the normal range but is not critical.
	// Units shall use the same units as the related ReadingVolts property.
	LowerThresholdNonCritical float64
	// MaxReadingRange shall indicate the
	// highest possible value for ReadingVolts. Units shall use the same
	// units as the related ReadingVolts property.
	MaxReadingRange float64
	// MemberID shall uniquely identify the member within the collection. For
	// services supporting Redfish v1.6 or higher, this value shall be the
	// zero-based array index.
	MemberID string `json:"MemberId"`
	// MinReadingRange shall indicate the lowest possible value for ReadingVolts.
	// Units shall use the same units as the related ReadingVolts property.
	MinReadingRange float64
	// PhysicalContext shall be a description
	// of the affected device or region within the chassis to which this
	// voltage measurement applies.
	PhysicalContext string
	// ReadingVolts shall be the present
	// reading of the voltage sensor's reading.
	ReadingVolts float64
	// SensorNumber shall be a numerical
	// identifier for this voltage sensor that is unique within this
	// resource.
	SensorNumber int
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// UpperThresholdCritical shall indicate
	// the present reading is above the normal range but is not yet fatal.
	// Units shall use the same units as the related ReadingVolts property.
	UpperThresholdCritical float64
	// UpperThresholdFatal shall indicate the
	// present reading is above the normal range and is fatal. Units shall
	// use the same units as the related ReadingVolts property.
	UpperThresholdFatal float64
	// UpperThresholdNonCritical shall indicate
	// the present reading is above the normal range but is not critical.
	// Units shall use the same units as the related ReadingVolts property.
	UpperThresholdNonCritical float64
}

// UnmarshalJSON unmarshals a Voltage object from the raw JSON.
func (voltage *Voltage) UnmarshalJSON(b []byte) error { // nolint:dupl
	type temp Voltage
	type t1 struct {
		temp
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		// See if we need to handle converting MemberID
		var t2 struct {
			t1
			MemberID int `json:"MemberId"`
		}
		err2 := json.Unmarshal(b, &t2)

		if err2 != nil {
			// Return the original error
			return err
		}

		// Convert the numeric member ID to a string
		t = t2.t1
		t.temp.MemberID = strconv.Itoa(t2.MemberID)
	}

	// Extract the links to other entities for later
	*voltage = Voltage(t.temp)

	return nil
}
