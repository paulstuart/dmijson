package dmijson

import (
	"encoding/json"
)

//go:generate sudo bin/dmij -f test.json
//go:generate go-bindata -pkg=dmijson bin

type BaseBoardInformation struct {
	Manufacturer           string
	ProductName            string
	Version                string
	SerialNumber           string
	AssetTag               string
	Features               []string
	LocationInChassis      string
	ChassisHandle          string
	Type                   string
	ContainedObjectHandles int
}

type BIOSInformation struct {
	Vendor          string
	Version         string
	ReleaseDate     string
	Address         string
	RuntimeSize     string
	ROMSize         string
	Characteristics []string
	BIOSRevision    string
}

type BIOSLanguageInformation struct {
	LanguageDescriptionFormat  string
	InstallableLanguages       int
	CurrentlyInstalledLanguage string
}

type CacheInformation struct {
	SocketDesignation   string
	Configuration       string
	OperationalMode     string
	Location            string
	InstalledSize       string
	MaximumSize         string
	SupportedSRAMTypes  []string
	InstalledSRAMType   string
	Speed               string
	ErrorCorrectionType string
	SystemType          string
	Associativity       string
}

type ChassisInformation struct {
	Manufacturer       string
	Type               string
	Lock               string
	Version            string
	SerialNumber       string
	AssetTag           string
	BootupState        string
	PowerSupplyState   string
	ThermalState       string
	SecurityStatus     string
	OEMInformation     string
	Height             string
	NumberOfPowerCords int
	ContainedElements  int
	SKUNumber          string
}

type CoolingDevice struct {
	TemperatureProbeHandle string
	Type                   string
	Status                 string
	CoolingUnitGroup       int
	OEMspecificInformation string
	NominalSpeed           string
	Description            string
}

type ElectricalCurrentProbe struct {
	Description            string
	Location               string
	Status                 string
	MaximumValue           string
	MinimumValue           string
	Resolution             string
	Tolerance              string
	Accuracy               string
	OEMspecificInformation string
	NominalValue           string
}

type IPMIDeviceInformation struct {
	InterfaceType        string
	SpecificationVersion string
	I2CSlaveAddress      string
	NVStorageDevice      string
	BaseAddress          string
	RegisterSpacing      string
}

type ManagementDevice struct {
	Description string
	Type        string
	Address     string
	AddressType string
}

type ManagementDeviceComponent struct {
	Description            string
	ManagementDeviceHandle string
	ComponentHandle        string
	ThresholdHandle        string
}

type ManagementDeviceThresholdData struct {
	LowerNoncriticalThreshold    int
	UpperNoncriticalThreshold    int
	LowerCriticalThreshold       int
	UpperCriticalThreshold       int
	LowerNonrecoverableThreshold int
	UpperNonrecoverableThreshold int
}

type MemoryArrayMappedAddress struct {
	StartingAddress     string
	EndingAddress       string
	RangeSize           string
	PhysicalArrayHandle string
	PartitionWidth      int
}

type MemoryDevice struct {
	ArrayHandle            string
	ErrorInformationHandle string
	TotalWidth             string
	DataWidth              string
	Size                   string
	FormFactor             string
	Set                    string
	Locator                string
	BankLocator            string
	Type                   string
	TypeDetail             string
	Speed                  string
	Manufacturer           string
	SerialNumber           string
	AssetTag               string
	PartNumber             string
	Rank                   string
	ConfiguredClockSpeed   string
}

type MemoryDeviceMappedAddress struct {
	StartingAddress                string
	EndingAddress                  string
	RangeSize                      string
	PhysicalDeviceHandle           string
	MemoryArrayMappedAddressHandle string
	PartitionRowPosition           int
}

type OEMStrings map[string]string

type OnboardDevice struct {
	ReferenceDesignation string
	Type                 string
	Status               string
	TypeInstance         int
	BusAddress           string
}

type OnBoardDeviceInformation struct {
	Type        string
	Status      string
	Description string
}

type PhysicalMemoryArray struct {
	Location               string
	Use                    string
	ErrorCorrectionType    string
	MaximumCapacity        string
	ErrorInformationHandle string
	NumberOfDevices        int
}

type PortConnectorInformation struct {
	InternalReferenceDesignator string
	InternalConnectorType       string
	ExternalReferenceDesignator string
	ExternalConnectorType       string
	PortType                    string
}

type ProcessorInformation struct {
	SocketDesignation string
	Type              string
	Family            string
	Manufacturer      string
	ID                string
	Signature         string
	Flags             []string
	Version           string
	Voltage           string
	ExternalClock     string
	MaxSpeed          string
	CurrentSpeed      string
	Status            string
	Upgrade           string
	L1CacheHandle     string
	L2CacheHandle     string
	L3CacheHandle     string
	SerialNumber      string
	AssetTag          string
	PartNumber        string
	CoreCount         int
	CoreEnabled       int
	ThreadCount       int
	Characteristics   []string
}

type SystemBootInformation struct {
	Status string
}

type SystemConfigurationOptions struct {
	Option1 string
}

type SystemEventLog struct {
	AreaLength                  string
	HeaderStartOffset           string
	HeaderLength                string
	DataStartOffset             string
	AccessMethod                string
	AccessAddress               string
	Status                      string
	ChangeToken                 string
	HeaderFormat                string
	SupportedLogTypeDescriptors int
	Descriptor1                 string
	DataFormat1                 string
	Descriptor2                 string
	DataFormat2                 string
	Descriptor3                 string
	DataFormat3                 string
	Descriptor4                 string
	DataFormat4                 string
	Descriptor5                 string
	DataFormat5                 string
	Descriptor6                 string
	DataFormat6                 string
	Descriptor7                 string
	DataFormat7                 string
	Descriptor8                 string
	DataFormat8                 string
	Descriptor9                 string
	DataFormat9                 string
	Descriptor10                string
	DataFormat10                string
	Descriptor11                string
	DataFormat11                string
	Descriptor12                string
	DataFormat12                string
	Descriptor13                string
	DataFormat13                string
	Descriptor14                string
	DataFormat14                string
	Descriptor15                string
	DataFormat15                string
	Descriptor16                string
	DataFormat16                string
	Descriptor17                string
	DataFormat17                string
	Descriptor18                string
	DataFormat18                string
	Descriptor19                string
	DataFormat19                string
	Descriptor20                string
	DataFormat20                string
	Descriptor21                string
	DataFormat21                string
	Descriptor22                string
	DataFormat22                string
	Descriptor23                string
	DataFormat23                string
	Descriptor24                string
	DataFormat24                string
	Descriptor25                string
	DataFormat25                string
}

type SystemInformation struct {
	Manufacturer string
	ProductName  string
	Version      string
	SerialNumber string
	UUID         string
	WakeupType   string
	SKUNumber    string
	Family       string
}

type SystemPowerSupply struct {
	PowerUnitGroup             int
	Location                   string
	Name                       string
	Manufacturer               string
	SerialNumber               string
	AssetTag                   string
	ModelPartNumber            string
	Revision                   string
	MaxPowerCapacity           string
	Status                     string
	Type                       string
	InputVoltageRangeSwitching string
	Plugged                    string
	HotReplaceable             string
	InputVoltageProbeHandle    string
	CoolingDeviceHandle        string
	InputCurrentProbeHandle    string
}

type SystemSlotInformation struct {
	Designation     string
	Type            string
	CurrentUsage    string
	Length          string
	ID              int
	Characteristics []string
	BusAddress      string
}

type Probe struct {
	Description            string
	Location               string
	Status                 string
	MaximumValue           string
	MinimumValue           string
	Resolution             string
	Tolerance              string
	Accuracy               string
	OEMspecificInformation string
	NominalValue           string
}

type DMI struct {
	BaseBoardInformation          BaseBoardInformation
	BIOSInformation               BIOSInformation
	BIOSLanguageInformation       BIOSLanguageInformation
	CacheInformation              []CacheInformation
	ChassisInformation            ChassisInformation
	CoolingDevice                 []CoolingDevice
	ElectricalCurrentProbe        []ElectricalCurrentProbe
	IPMIDeviceInformation         IPMIDeviceInformation
	ManagementDevice              []ManagementDevice
	ManagementDeviceComponent     []ManagementDeviceComponent
	ManagementDeviceThresholdData []ManagementDeviceThresholdData
	MemoryArrayMappedAddress      []MemoryArrayMappedAddress
	MemoryDevice                  []MemoryDevice
	MemoryDeviceMappedAddress     []MemoryDeviceMappedAddress
	OEMStrings                    OEMStrings
	OnboardDevice                 []OnboardDevice
	OnBoardDevice1Information     OnBoardDeviceInformation
	OnBoardDevice2Information     OnBoardDeviceInformation
	OnBoardDevice3Information     OnBoardDeviceInformation
	OnBoardDevice4Information     OnBoardDeviceInformation
	PhysicalMemoryArray           []PhysicalMemoryArray
	PortConnectorInformation      []PortConnectorInformation
	ProcessorInformation          []ProcessorInformation
	SystemBootInformation         SystemBootInformation
	SystemConfigurationOptions    SystemConfigurationOptions
	SystemEventLog                SystemEventLog
	SystemInformation             SystemInformation
	SystemPowerSupply             SystemPowerSupply
	SystemSlotInformation         []SystemSlotInformation
	TemperatureProbe              []Probe
	VoltageProbe                  []Probe
}

func FromJSON(j string) (DMI, error) {
	d := DMI{}
	return d, json.Unmarshal([]byte(j), &d)
}

func (d DMI) ToJSON() string {
	j, _ := json.MarshalIndent(d, "  ", "  ")
	return string(j)
}

func Script() (string, error) {
	s, err := Asset("bin/dmij")
	return string(s), err
}
