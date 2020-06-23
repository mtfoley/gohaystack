package tags

import "github.com/owulveryck/gohaystack"

var (
	// AbsorptionLabel https://project-haystack.org/tag/absorption
	AbsorptionLabel = labelDB["absorption"]
	// AcLabel https://project-haystack.org/tag/ac
	AcLabel = labelDB["ac"]
	// ActiveLabel https://project-haystack.org/tag/active
	ActiveLabel = labelDB["active"]
	// AhuLabel https://project-haystack.org/tag/ahu
	AhuLabel = labelDB["ahu"]
	// AhuRefLabel https://project-haystack.org/tag/ahuRef
	AhuRefLabel = labelDB["ahuRef"]
	// AirLabel https://project-haystack.org/tag/air
	AirLabel = labelDB["air"]
	// AirCooledLabel https://project-haystack.org/tag/airCooled
	AirCooledLabel = labelDB["airCooled"]
	// AngleLabel https://project-haystack.org/tag/angle
	AngleLabel = labelDB["angle"]
	// ApparentLabel https://project-haystack.org/tag/apparent
	ApparentLabel = labelDB["apparent"]
	// AreaLabel https://project-haystack.org/tag/area
	AreaLabel = labelDB["area"]
	// AvgLabel https://project-haystack.org/tag/avg
	AvgLabel = labelDB["avg"]
	// BarometricLabel https://project-haystack.org/tag/barometric
	BarometricLabel = labelDB["barometric"]
	// BlowdownLabel https://project-haystack.org/tag/blowdown
	BlowdownLabel = labelDB["blowdown"]
	// BoilerLabel https://project-haystack.org/tag/boiler
	BoilerLabel = labelDB["boiler"]
	// BypassLabel https://project-haystack.org/tag/bypass
	BypassLabel = labelDB["bypass"]
	// CentrifugalLabel https://project-haystack.org/tag/centrifugal
	CentrifugalLabel = labelDB["centrifugal"]
	// ChilledLabel https://project-haystack.org/tag/chilled
	ChilledLabel = labelDB["chilled"]
	// ChilledBeamZoneLabel https://project-haystack.org/tag/chilledBeamZone
	ChilledBeamZoneLabel = labelDB["chilledBeamZone"]
	// ChilledWaterCoolLabel https://project-haystack.org/tag/chilledWaterCool
	ChilledWaterCoolLabel = labelDB["chilledWaterCool"]
	// ChilledWaterPlantLabel https://project-haystack.org/tag/chilledWaterPlant
	ChilledWaterPlantLabel = labelDB["chilledWaterPlant"]
	// ChilledWaterPlantRefLabel https://project-haystack.org/tag/chilledWaterPlantRef
	ChilledWaterPlantRefLabel = labelDB["chilledWaterPlantRef"]
	// ChillerLabel https://project-haystack.org/tag/chiller
	ChillerLabel = labelDB["chiller"]
	// CircLabel https://project-haystack.org/tag/circ
	CircLabel = labelDB["circ"]
	// CircuitLabel https://project-haystack.org/tag/circuit
	CircuitLabel = labelDB["circuit"]
	// ClosedLoopLabel https://project-haystack.org/tag/closedLoop
	ClosedLoopLabel = labelDB["closedLoop"]
	// CloudageLabel https://project-haystack.org/tag/cloudage
	CloudageLabel = labelDB["cloudage"]
	// CmdLabel https://project-haystack.org/tag/cmd
	CmdLabel = labelDB["cmd"]
	// CoLabel https://project-haystack.org/tag/co
	CoLabel = labelDB["co"]
	// Co2Label https://project-haystack.org/tag/co2
	Co2Label = labelDB["co2"]
	// ColdDeckLabel https://project-haystack.org/tag/coldDeck
	ColdDeckLabel = labelDB["coldDeck"]
	// CondensateLabel https://project-haystack.org/tag/condensate
	CondensateLabel = labelDB["condensate"]
	// CondenserLabel https://project-haystack.org/tag/condenser
	CondenserLabel = labelDB["condenser"]
	// ConnectionLabel https://project-haystack.org/tag/connection
	ConnectionLabel = labelDB["connection"]
	// ConstantVolumeLabel https://project-haystack.org/tag/constantVolume
	ConstantVolumeLabel = labelDB["constantVolume"]
	// CoolLabel https://project-haystack.org/tag/cool
	CoolLabel = labelDB["cool"]
	// CoolOnlyLabel https://project-haystack.org/tag/coolOnly
	CoolOnlyLabel = labelDB["coolOnly"]
	// CoolingLabel https://project-haystack.org/tag/cooling
	CoolingLabel = labelDB["cooling"]
	// CoolingCapacityLabel https://project-haystack.org/tag/coolingCapacity
	CoolingCapacityLabel = labelDB["coolingCapacity"]
	// CoolingTowerLabel https://project-haystack.org/tag/coolingTower
	CoolingTowerLabel = labelDB["coolingTower"]
	// CurLabel https://project-haystack.org/tag/cur
	CurLabel = labelDB["cur"]
	// CurErrLabel https://project-haystack.org/tag/curErr
	CurErrLabel = labelDB["curErr"]
	// CurStatusLabel https://project-haystack.org/tag/curStatus
	CurStatusLabel = labelDB["curStatus"]
	// CurValLabel https://project-haystack.org/tag/curVal
	CurValLabel = labelDB["curVal"]
	// CurrentLabel https://project-haystack.org/tag/current
	CurrentLabel = labelDB["current"]
	// DamperLabel https://project-haystack.org/tag/damper
	DamperLabel = labelDB["damper"]
	// DcLabel https://project-haystack.org/tag/dc
	DcLabel = labelDB["dc"]
	// DeltaLabel https://project-haystack.org/tag/delta
	DeltaLabel = labelDB["delta"]
	// DeviceLabel https://project-haystack.org/tag/device
	DeviceLabel = labelDB["device"]
	// Device1RefLabel https://project-haystack.org/tag/device1Ref
	Device1RefLabel = labelDB["device1Ref"]
	// Device2RefLabel https://project-haystack.org/tag/device2Ref
	Device2RefLabel = labelDB["device2Ref"]
	// DewLabel https://project-haystack.org/tag/dew
	DewLabel = labelDB["dew"]
	// DirectZoneLabel https://project-haystack.org/tag/directZone
	DirectZoneLabel = labelDB["directZone"]
	// DirectionLabel https://project-haystack.org/tag/direction
	DirectionLabel = labelDB["direction"]
	// DisLabel https://project-haystack.org/tag/dis
	DisLabel = labelDB["dis"]
	// DischargeLabel https://project-haystack.org/tag/discharge
	DischargeLabel = labelDB["discharge"]
	// DivertingLabel https://project-haystack.org/tag/diverting
	DivertingLabel = labelDB["diverting"]
	// DomesticLabel https://project-haystack.org/tag/domestic
	DomesticLabel = labelDB["domestic"]
	// DualDuctLabel https://project-haystack.org/tag/dualDuct
	DualDuctLabel = labelDB["dualDuct"]
	// DuctAreaLabel https://project-haystack.org/tag/ductArea
	DuctAreaLabel = labelDB["ductArea"]
	// DxCoolLabel https://project-haystack.org/tag/dxCool
	DxCoolLabel = labelDB["dxCool"]
	// EffectiveLabel https://project-haystack.org/tag/effective
	EffectiveLabel = labelDB["effective"]
	// EfficiencyLabel https://project-haystack.org/tag/efficiency
	EfficiencyLabel = labelDB["efficiency"]
	// ElecLabel https://project-haystack.org/tag/elec
	ElecLabel = labelDB["elec"]
	// ElecHeatLabel https://project-haystack.org/tag/elecHeat
	ElecHeatLabel = labelDB["elecHeat"]
	// ElecMeterLoadLabel https://project-haystack.org/tag/elecMeterLoad
	ElecMeterLoadLabel = labelDB["elecMeterLoad"]
	// ElecMeterRefLabel https://project-haystack.org/tag/elecMeterRef
	ElecMeterRefLabel = labelDB["elecMeterRef"]
	// ElecPanelLabel https://project-haystack.org/tag/elecPanel
	ElecPanelLabel = labelDB["elecPanel"]
	// ElecPanelOfLabel https://project-haystack.org/tag/elecPanelOf
	ElecPanelOfLabel = labelDB["elecPanelOf"]
	// ElecReheatLabel https://project-haystack.org/tag/elecReheat
	ElecReheatLabel = labelDB["elecReheat"]
	// EnableLabel https://project-haystack.org/tag/enable
	EnableLabel = labelDB["enable"]
	// EnergyLabel https://project-haystack.org/tag/energy
	EnergyLabel = labelDB["energy"]
	// EnteringLabel https://project-haystack.org/tag/entering
	EnteringLabel = labelDB["entering"]
	// EnumLabel https://project-haystack.org/tag/enum
	EnumLabel = labelDB["enum"]
	// EquipLabel https://project-haystack.org/tag/equip
	EquipLabel = labelDB["equip"]
	// EquipRefLabel https://project-haystack.org/tag/equipRef
	EquipRefLabel = labelDB["equipRef"]
	// EvaporatorLabel https://project-haystack.org/tag/evaporator
	EvaporatorLabel = labelDB["evaporator"]
	// ExhaustLabel https://project-haystack.org/tag/exhaust
	ExhaustLabel = labelDB["exhaust"]
	// ExportLabel https://project-haystack.org/tag/export
	ExportLabel = labelDB["export"]
	// FaceBypassLabel https://project-haystack.org/tag/faceBypass
	FaceBypassLabel = labelDB["faceBypass"]
	// FanLabel https://project-haystack.org/tag/fan
	FanLabel = labelDB["fan"]
	// FanPoweredLabel https://project-haystack.org/tag/fanPowered
	FanPoweredLabel = labelDB["fanPowered"]
	// FcuLabel https://project-haystack.org/tag/fcu
	FcuLabel = labelDB["fcu"]
	// FilterLabel https://project-haystack.org/tag/filter
	FilterLabel = labelDB["filter"]
	// FlowLabel https://project-haystack.org/tag/flow
	FlowLabel = labelDB["flow"]
	// FlueLabel https://project-haystack.org/tag/flue
	FlueLabel = labelDB["flue"]
	// FreezeStatLabel https://project-haystack.org/tag/freezeStat
	FreezeStatLabel = labelDB["freezeStat"]
	// FreqLabel https://project-haystack.org/tag/freq
	FreqLabel = labelDB["freq"]
	// GasLabel https://project-haystack.org/tag/gas
	GasLabel = labelDB["gas"]
	// GasHeatLabel https://project-haystack.org/tag/gasHeat
	GasHeatLabel = labelDB["gasHeat"]
	// GasMeterLoadLabel https://project-haystack.org/tag/gasMeterLoad
	GasMeterLoadLabel = labelDB["gasMeterLoad"]
	// GeoAddrLabel https://project-haystack.org/tag/geoAddr
	GeoAddrLabel = labelDB["geoAddr"]
	// GeoCityLabel https://project-haystack.org/tag/geoCity
	GeoCityLabel = labelDB["geoCity"]
	// GeoCoordLabel https://project-haystack.org/tag/geoCoord
	GeoCoordLabel = labelDB["geoCoord"]
	// GeoCountryLabel https://project-haystack.org/tag/geoCountry
	GeoCountryLabel = labelDB["geoCountry"]
	// GeoCountyLabel https://project-haystack.org/tag/geoCounty
	GeoCountyLabel = labelDB["geoCounty"]
	// GeoPostalCodeLabel https://project-haystack.org/tag/geoPostalCode
	GeoPostalCodeLabel = labelDB["geoPostalCode"]
	// GeoStateLabel https://project-haystack.org/tag/geoState
	GeoStateLabel = labelDB["geoState"]
	// GeoStreetLabel https://project-haystack.org/tag/geoStreet
	GeoStreetLabel = labelDB["geoStreet"]
	// HeaderLabel https://project-haystack.org/tag/header
	HeaderLabel = labelDB["header"]
	// HeatLabel https://project-haystack.org/tag/heat
	HeatLabel = labelDB["heat"]
	// HeatExchangerLabel https://project-haystack.org/tag/heatExchanger
	HeatExchangerLabel = labelDB["heatExchanger"]
	// HeatPumpLabel https://project-haystack.org/tag/heatPump
	HeatPumpLabel = labelDB["heatPump"]
	// HeatWheelLabel https://project-haystack.org/tag/heatWheel
	HeatWheelLabel = labelDB["heatWheel"]
	// HeatingLabel https://project-haystack.org/tag/heating
	HeatingLabel = labelDB["heating"]
	// HisLabel https://project-haystack.org/tag/his
	HisLabel = labelDB["his"]
	// HisErrLabel https://project-haystack.org/tag/hisErr
	HisErrLabel = labelDB["hisErr"]
	// HisInterpolateLabel https://project-haystack.org/tag/hisInterpolate
	HisInterpolateLabel = labelDB["hisInterpolate"]
	// HisStatusLabel https://project-haystack.org/tag/hisStatus
	HisStatusLabel = labelDB["hisStatus"]
	// HisTotalizedLabel https://project-haystack.org/tag/hisTotalized
	HisTotalizedLabel = labelDB["hisTotalized"]
	// HotLabel https://project-haystack.org/tag/hot
	HotLabel = labelDB["hot"]
	// HotDeckLabel https://project-haystack.org/tag/hotDeck
	HotDeckLabel = labelDB["hotDeck"]
	// HotWaterHeatLabel https://project-haystack.org/tag/hotWaterHeat
	HotWaterHeatLabel = labelDB["hotWaterHeat"]
	// HotWaterPlantLabel https://project-haystack.org/tag/hotWaterPlant
	HotWaterPlantLabel = labelDB["hotWaterPlant"]
	// HotWaterPlantRefLabel https://project-haystack.org/tag/hotWaterPlantRef
	HotWaterPlantRefLabel = labelDB["hotWaterPlantRef"]
	// HotWaterReheatLabel https://project-haystack.org/tag/hotWaterReheat
	HotWaterReheatLabel = labelDB["hotWaterReheat"]
	// HumidifierLabel https://project-haystack.org/tag/humidifier
	HumidifierLabel = labelDB["humidifier"]
	// HumidityLabel https://project-haystack.org/tag/humidity
	HumidityLabel = labelDB["humidity"]
	// HvacLabel https://project-haystack.org/tag/hvac
	HvacLabel = labelDB["hvac"]
	// IDLabel https://project-haystack.org/tag/iD
	IDLabel = labelDB["id"]
	// ImbalanceLabel https://project-haystack.org/tag/imbalance
	ImbalanceLabel = labelDB["imbalance"]
	// ImportLabel https://project-haystack.org/tag/import
	ImportLabel = labelDB["import"]
	// IrradianceLabel https://project-haystack.org/tag/irradiance
	IrradianceLabel = labelDB["irradiance"]
	// IsolationLabel https://project-haystack.org/tag/isolation
	IsolationLabel = labelDB["isolation"]
	// KindLabel https://project-haystack.org/tag/kind
	KindLabel = labelDB["kind"]
	// LeavingLabel https://project-haystack.org/tag/leaving
	LeavingLabel = labelDB["leaving"]
	// LevelLabel https://project-haystack.org/tag/level
	LevelLabel = labelDB["level"]
	// LightLevelLabel https://project-haystack.org/tag/lightLevel
	LightLevelLabel = labelDB["lightLevel"]
	// LightingLabel https://project-haystack.org/tag/lighting
	LightingLabel = labelDB["lighting"]
	// LightsLabel https://project-haystack.org/tag/lights
	LightsLabel = labelDB["lights"]
	// LightsGroupLabel https://project-haystack.org/tag/lightsGroup
	LightsGroupLabel = labelDB["lightsGroup"]
	// LoadLabel https://project-haystack.org/tag/load
	LoadLabel = labelDB["load"]
	// MagLabel https://project-haystack.org/tag/mag
	MagLabel = labelDB["mag"]
	// MakeupLabel https://project-haystack.org/tag/makeup
	MakeupLabel = labelDB["makeup"]
	// MauLabel https://project-haystack.org/tag/mau
	MauLabel = labelDB["mau"]
	// MaxLabel https://project-haystack.org/tag/max
	MaxLabel = labelDB["max"]
	// MaxValLabel https://project-haystack.org/tag/maxVal
	MaxValLabel = labelDB["maxVal"]
	// MeterLabel https://project-haystack.org/tag/meter
	MeterLabel = labelDB["meter"]
	// MinLabel https://project-haystack.org/tag/min
	MinLabel = labelDB["min"]
	// MinValLabel https://project-haystack.org/tag/minVal
	MinValLabel = labelDB["minVal"]
	// MixedLabel https://project-haystack.org/tag/mixed
	MixedLabel = labelDB["mixed"]
	// MixingLabel https://project-haystack.org/tag/mixing
	MixingLabel = labelDB["mixing"]
	// MultiZoneLabel https://project-haystack.org/tag/multiZone
	MultiZoneLabel = labelDB["multiZone"]
	// NetLabel https://project-haystack.org/tag/net
	NetLabel = labelDB["net"]
	// NetworkLabel https://project-haystack.org/tag/network
	NetworkLabel = labelDB["network"]
	// NetworkRefLabel https://project-haystack.org/tag/networkRef
	NetworkRefLabel = labelDB["networkRef"]
	// NeutralDeckLabel https://project-haystack.org/tag/neutralDeck
	NeutralDeckLabel = labelDB["neutralDeck"]
	// OccLabel https://project-haystack.org/tag/occ
	OccLabel = labelDB["occ"]
	// OccupancyIndicatorLabel https://project-haystack.org/tag/occupancyIndicator
	OccupancyIndicatorLabel = labelDB["occupancyIndicator"]
	// OccupiedLabel https://project-haystack.org/tag/occupied
	OccupiedLabel = labelDB["occupied"]
	// OilLabel https://project-haystack.org/tag/oil
	OilLabel = labelDB["oil"]
	// OpenLoopLabel https://project-haystack.org/tag/openLoop
	OpenLoopLabel = labelDB["openLoop"]
	// OutsideLabel https://project-haystack.org/tag/outside
	OutsideLabel = labelDB["outside"]
	// ParallelLabel https://project-haystack.org/tag/parallel
	ParallelLabel = labelDB["parallel"]
	// PerimeterHeatLabel https://project-haystack.org/tag/perimeterHeat
	PerimeterHeatLabel = labelDB["perimeterHeat"]
	// PfLabel https://project-haystack.org/tag/pf
	PfLabel = labelDB["pf"]
	// PhaseLabel https://project-haystack.org/tag/phase
	PhaseLabel = labelDB["phase"]
	// PointLabel https://project-haystack.org/tag/point
	PointLabel = labelDB["point"]
	// PowerLabel https://project-haystack.org/tag/power
	PowerLabel = labelDB["power"]
	// PrecipitationLabel https://project-haystack.org/tag/precipitation
	PrecipitationLabel = labelDB["precipitation"]
	// PressureLabel https://project-haystack.org/tag/pressure
	PressureLabel = labelDB["pressure"]
	// PressureDependentLabel https://project-haystack.org/tag/pressureDependent
	PressureDependentLabel = labelDB["pressureDependent"]
	// PressureIndependentLabel https://project-haystack.org/tag/pressureIndependent
	PressureIndependentLabel = labelDB["pressureIndependent"]
	// PrimaryFunctionLabel https://project-haystack.org/tag/primaryFunction
	PrimaryFunctionLabel = labelDB["primaryFunction"]
	// PrimaryLoopLabel https://project-haystack.org/tag/primaryLoop
	PrimaryLoopLabel = labelDB["primaryLoop"]
	// ProtocolLabel https://project-haystack.org/tag/protocol
	ProtocolLabel = labelDB["protocol"]
	// PumpLabel https://project-haystack.org/tag/pump
	PumpLabel = labelDB["pump"]
	// ReactiveLabel https://project-haystack.org/tag/reactive
	ReactiveLabel = labelDB["reactive"]
	// ReciprocalLabel https://project-haystack.org/tag/reciprocal
	ReciprocalLabel = labelDB["reciprocal"]
	// RefrigLabel https://project-haystack.org/tag/refrig
	RefrigLabel = labelDB["refrig"]
	// ReheatLabel https://project-haystack.org/tag/reheat
	ReheatLabel = labelDB["reheat"]
	// ReheatingLabel https://project-haystack.org/tag/reheating
	ReheatingLabel = labelDB["reheating"]
	// ReturnLabel https://project-haystack.org/tag/return
	ReturnLabel = labelDB["return"]
	// RooftopLabel https://project-haystack.org/tag/rooftop
	RooftopLabel = labelDB["rooftop"]
	// RunLabel https://project-haystack.org/tag/run
	RunLabel = labelDB["run"]
	// ScrewLabel https://project-haystack.org/tag/screw
	ScrewLabel = labelDB["screw"]
	// SecondaryLoopLabel https://project-haystack.org/tag/secondaryLoop
	SecondaryLoopLabel = labelDB["secondaryLoop"]
	// SensorLabel https://project-haystack.org/tag/sensor
	SensorLabel = labelDB["sensor"]
	// SeriesLabel https://project-haystack.org/tag/series
	SeriesLabel = labelDB["series"]
	// SingleDuctLabel https://project-haystack.org/tag/singleDuct
	SingleDuctLabel = labelDB["singleDuct"]
	// SiteLabel https://project-haystack.org/tag/site
	SiteLabel = labelDB["site"]
	// SiteMeterLabel https://project-haystack.org/tag/siteMeter
	SiteMeterLabel = labelDB["siteMeter"]
	// SitePanelLabel https://project-haystack.org/tag/sitePanel
	SitePanelLabel = labelDB["sitePanel"]
	// SiteRefLabel https://project-haystack.org/tag/siteRef
	SiteRefLabel = labelDB["siteRef"]
	// SolarLabel https://project-haystack.org/tag/solar
	SolarLabel = labelDB["solar"]
	// SpLabel https://project-haystack.org/tag/sp
	SpLabel = labelDB["sp"]
	// SpeedLabel https://project-haystack.org/tag/speed
	SpeedLabel = labelDB["speed"]
	// StageLabel https://project-haystack.org/tag/stage
	StageLabel = labelDB["stage"]
	// StandbyLabel https://project-haystack.org/tag/standby
	StandbyLabel = labelDB["standby"]
	// SteamLabel https://project-haystack.org/tag/steam
	SteamLabel = labelDB["steam"]
	// SteamHeatLabel https://project-haystack.org/tag/steamHeat
	SteamHeatLabel = labelDB["steamHeat"]
	// SteamMeterLoadLabel https://project-haystack.org/tag/steamMeterLoad
	SteamMeterLoadLabel = labelDB["steamMeterLoad"]
	// SteamPlantLabel https://project-haystack.org/tag/steamPlant
	SteamPlantLabel = labelDB["steamPlant"]
	// SteamPlantRefLabel https://project-haystack.org/tag/steamPlantRef
	SteamPlantRefLabel = labelDB["steamPlantRef"]
	// SubPanelOfLabel https://project-haystack.org/tag/subPanelOf
	SubPanelOfLabel = labelDB["subPanelOf"]
	// SubmeterOfLabel https://project-haystack.org/tag/submeterOf
	SubmeterOfLabel = labelDB["submeterOf"]
	// SunriseLabel https://project-haystack.org/tag/sunrise
	SunriseLabel = labelDB["sunrise"]
	// TankLabel https://project-haystack.org/tag/tank
	TankLabel = labelDB["tank"]
	// TempLabel https://project-haystack.org/tag/temp
	TempLabel = labelDB["temp"]
	// ThdLabel https://project-haystack.org/tag/thd
	ThdLabel = labelDB["thd"]
	// TotalLabel https://project-haystack.org/tag/total
	TotalLabel = labelDB["total"]
	// TripleDuctLabel https://project-haystack.org/tag/tripleDuct
	TripleDuctLabel = labelDB["tripleDuct"]
	// TzLabel https://project-haystack.org/tag/tz
	TzLabel = labelDB["tz"]
	// UnitLabel https://project-haystack.org/tag/unit
	UnitLabel = labelDB["unit"]
	// UnoccLabel https://project-haystack.org/tag/unocc
	UnoccLabel = labelDB["unocc"]
	// UvLabel https://project-haystack.org/tag/uv
	UvLabel = labelDB["uv"]
	// ValveLabel https://project-haystack.org/tag/valve
	ValveLabel = labelDB["valve"]
	// VariableVolumeLabel https://project-haystack.org/tag/variableVolume
	VariableVolumeLabel = labelDB["variableVolume"]
	// VavLabel https://project-haystack.org/tag/vav
	VavLabel = labelDB["vav"]
	// VavModeLabel https://project-haystack.org/tag/vavMode
	VavModeLabel = labelDB["vavMode"]
	// VavZoneLabel https://project-haystack.org/tag/vavZone
	VavZoneLabel = labelDB["vavZone"]
	// VfdLabel https://project-haystack.org/tag/vfd
	VfdLabel = labelDB["vfd"]
	// VisibilityLabel https://project-haystack.org/tag/visibility
	VisibilityLabel = labelDB["visibility"]
	// VoltLabel https://project-haystack.org/tag/volt
	VoltLabel = labelDB["volt"]
	// VolumeLabel https://project-haystack.org/tag/volume
	VolumeLabel = labelDB["volume"]
	// WaterLabel https://project-haystack.org/tag/water
	WaterLabel = labelDB["water"]
	// WaterCooledLabel https://project-haystack.org/tag/waterCooled
	WaterCooledLabel = labelDB["waterCooled"]
	// WaterMeterLoadLabel https://project-haystack.org/tag/waterMeterLoad
	WaterMeterLoadLabel = labelDB["waterMeterLoad"]
	// WeatherLabel https://project-haystack.org/tag/weather
	WeatherLabel = labelDB["weather"]
	// WeatherCondLabel https://project-haystack.org/tag/weatherCond
	WeatherCondLabel = labelDB["weatherCond"]
	// WeatherPointLabel https://project-haystack.org/tag/weatherPoint
	WeatherPointLabel = labelDB["weatherPoint"]
	// WeatherRefLabel https://project-haystack.org/tag/weatherRef
	WeatherRefLabel = labelDB["weatherRef"]
	// WetBulbLabel https://project-haystack.org/tag/wetBulb
	WetBulbLabel = labelDB["wetBulb"]
	// WindLabel https://project-haystack.org/tag/wind
	WindLabel = labelDB["wind"]
	// WritableLabel https://project-haystack.org/tag/writable
	WritableLabel = labelDB["writable"]
	// WriteErrLabel https://project-haystack.org/tag/writeErr
	WriteErrLabel = labelDB["writeErr"]
	// WriteLevelLabel https://project-haystack.org/tag/writeLevel
	WriteLevelLabel = labelDB["writeLevel"]
	// WriteStatusLabel https://project-haystack.org/tag/writeStatus
	WriteStatusLabel = labelDB["writeStatus"]
	// WriteValLabel https://project-haystack.org/tag/writeVal
	WriteValLabel = labelDB["writeVal"]
	// YearBuiltLabel https://project-haystack.org/tag/yearBuilt
	YearBuiltLabel = labelDB["yearBuilt"]
	// ZoneLabel https://project-haystack.org/tag/zone
	ZoneLabel = labelDB["zone"]
)

var (
	// Absorption Marker Tag
	Absorption = markerDB["absorption"]
	// Ac Marker Tag
	Ac = markerDB["ac"]
	// Active Marker Tag
	Active = markerDB["active"]
	// Ahu Marker Tag
	Ahu = markerDB["ahu"]
	// Air Marker Tag
	Air = markerDB["air"]
	// AirCooled Marker Tag
	AirCooled = markerDB["airCooled"]
	// Angle Marker Tag
	Angle = markerDB["angle"]
	// Apparent Marker Tag
	Apparent = markerDB["apparent"]
	// Avg Marker Tag
	Avg = markerDB["avg"]
	// Barometric Marker Tag
	Barometric = markerDB["barometric"]
	// Blowdown Marker Tag
	Blowdown = markerDB["blowdown"]
	// Boiler Marker Tag
	Boiler = markerDB["boiler"]
	// Bypass Marker Tag
	Bypass = markerDB["bypass"]
	// Centrifugal Marker Tag
	Centrifugal = markerDB["centrifugal"]
	// Chilled Marker Tag
	Chilled = markerDB["chilled"]
	// ChilledBeamZone Marker Tag
	ChilledBeamZone = markerDB["chilledBeamZone"]
	// ChilledWaterCool Marker Tag
	ChilledWaterCool = markerDB["chilledWaterCool"]
	// ChilledWaterPlant Marker Tag
	ChilledWaterPlant = markerDB["chilledWaterPlant"]
	// Chiller Marker Tag
	Chiller = markerDB["chiller"]
	// Circ Marker Tag
	Circ = markerDB["circ"]
	// Circuit Marker Tag
	Circuit = markerDB["circuit"]
	// ClosedLoop Marker Tag
	ClosedLoop = markerDB["closedLoop"]
	// Cloudage Marker Tag
	Cloudage = markerDB["cloudage"]
	// Cmd Marker Tag
	Cmd = markerDB["cmd"]
	// Co Marker Tag
	Co = markerDB["co"]
	// Co2 Marker Tag
	Co2 = markerDB["co2"]
	// ColdDeck Marker Tag
	ColdDeck = markerDB["coldDeck"]
	// Condensate Marker Tag
	Condensate = markerDB["condensate"]
	// Condenser Marker Tag
	Condenser = markerDB["condenser"]
	// Connection Marker Tag
	Connection = markerDB["connection"]
	// ConstantVolume Marker Tag
	ConstantVolume = markerDB["constantVolume"]
	// Cool Marker Tag
	Cool = markerDB["cool"]
	// CoolOnly Marker Tag
	CoolOnly = markerDB["coolOnly"]
	// Cooling Marker Tag
	Cooling = markerDB["cooling"]
	// CoolingTower Marker Tag
	CoolingTower = markerDB["coolingTower"]
	// Cur Marker Tag
	Cur = markerDB["cur"]
	// Current Marker Tag
	Current = markerDB["current"]
	// Damper Marker Tag
	Damper = markerDB["damper"]
	// Dc Marker Tag
	Dc = markerDB["dc"]
	// Delta Marker Tag
	Delta = markerDB["delta"]
	// Device Marker Tag
	Device = markerDB["device"]
	// Dew Marker Tag
	Dew = markerDB["dew"]
	// DirectZone Marker Tag
	DirectZone = markerDB["directZone"]
	// Direction Marker Tag
	Direction = markerDB["direction"]
	// Discharge Marker Tag
	Discharge = markerDB["discharge"]
	// Diverting Marker Tag
	Diverting = markerDB["diverting"]
	// Domestic Marker Tag
	Domestic = markerDB["domestic"]
	// DualDuct Marker Tag
	DualDuct = markerDB["dualDuct"]
	// DuctArea Marker Tag
	DuctArea = markerDB["ductArea"]
	// DxCool Marker Tag
	DxCool = markerDB["dxCool"]
	// Effective Marker Tag
	Effective = markerDB["effective"]
	// Efficiency Marker Tag
	Efficiency = markerDB["efficiency"]
	// Elec Marker Tag
	Elec = markerDB["elec"]
	// ElecHeat Marker Tag
	ElecHeat = markerDB["elecHeat"]
	// ElecPanel Marker Tag
	ElecPanel = markerDB["elecPanel"]
	// ElecReheat Marker Tag
	ElecReheat = markerDB["elecReheat"]
	// Enable Marker Tag
	Enable = markerDB["enable"]
	// Energy Marker Tag
	Energy = markerDB["energy"]
	// Entering Marker Tag
	Entering = markerDB["entering"]
	// Equip Marker Tag
	Equip = markerDB["equip"]
	// Evaporator Marker Tag
	Evaporator = markerDB["evaporator"]
	// Exhaust Marker Tag
	Exhaust = markerDB["exhaust"]
	// Export Marker Tag
	Export = markerDB["export"]
	// FaceBypass Marker Tag
	FaceBypass = markerDB["faceBypass"]
	// Fan Marker Tag
	Fan = markerDB["fan"]
	// FanPowered Marker Tag
	FanPowered = markerDB["fanPowered"]
	// Fcu Marker Tag
	Fcu = markerDB["fcu"]
	// Filter Marker Tag
	Filter = markerDB["filter"]
	// Flow Marker Tag
	Flow = markerDB["flow"]
	// Flue Marker Tag
	Flue = markerDB["flue"]
	// FreezeStat Marker Tag
	FreezeStat = markerDB["freezeStat"]
	// Freq Marker Tag
	Freq = markerDB["freq"]
	// Gas Marker Tag
	Gas = markerDB["gas"]
	// GasHeat Marker Tag
	GasHeat = markerDB["gasHeat"]
	// Header Marker Tag
	Header = markerDB["header"]
	// Heat Marker Tag
	Heat = markerDB["heat"]
	// HeatExchanger Marker Tag
	HeatExchanger = markerDB["heatExchanger"]
	// HeatPump Marker Tag
	HeatPump = markerDB["heatPump"]
	// HeatWheel Marker Tag
	HeatWheel = markerDB["heatWheel"]
	// Heating Marker Tag
	Heating = markerDB["heating"]
	// His Marker Tag
	His = markerDB["his"]
	// HisTotalized Marker Tag
	HisTotalized = markerDB["hisTotalized"]
	// Hot Marker Tag
	Hot = markerDB["hot"]
	// HotDeck Marker Tag
	HotDeck = markerDB["hotDeck"]
	// HotWaterHeat Marker Tag
	HotWaterHeat = markerDB["hotWaterHeat"]
	// HotWaterPlant Marker Tag
	HotWaterPlant = markerDB["hotWaterPlant"]
	// HotWaterReheat Marker Tag
	HotWaterReheat = markerDB["hotWaterReheat"]
	// Humidifier Marker Tag
	Humidifier = markerDB["humidifier"]
	// Humidity Marker Tag
	Humidity = markerDB["humidity"]
	// Hvac Marker Tag
	Hvac = markerDB["hvac"]
	// Imbalance Marker Tag
	Imbalance = markerDB["imbalance"]
	// Import Marker Tag
	Import = markerDB["import"]
	// Irradiance Marker Tag
	Irradiance = markerDB["irradiance"]
	// Isolation Marker Tag
	Isolation = markerDB["isolation"]
	// Leaving Marker Tag
	Leaving = markerDB["leaving"]
	// Level Marker Tag
	Level = markerDB["level"]
	// LightLevel Marker Tag
	LightLevel = markerDB["lightLevel"]
	// Lighting Marker Tag
	Lighting = markerDB["lighting"]
	// Lights Marker Tag
	Lights = markerDB["lights"]
	// LightsGroup Marker Tag
	LightsGroup = markerDB["lightsGroup"]
	// Load Marker Tag
	Load = markerDB["load"]
	// Mag Marker Tag
	Mag = markerDB["mag"]
	// Makeup Marker Tag
	Makeup = markerDB["makeup"]
	// Mau Marker Tag
	Mau = markerDB["mau"]
	// Max Marker Tag
	Max = markerDB["max"]
	// Meter Marker Tag
	Meter = markerDB["meter"]
	// Min Marker Tag
	Min = markerDB["min"]
	// Mixed Marker Tag
	Mixed = markerDB["mixed"]
	// Mixing Marker Tag
	Mixing = markerDB["mixing"]
	// MultiZone Marker Tag
	MultiZone = markerDB["multiZone"]
	// Net Marker Tag
	Net = markerDB["net"]
	// Network Marker Tag
	Network = markerDB["network"]
	// NeutralDeck Marker Tag
	NeutralDeck = markerDB["neutralDeck"]
	// Occ Marker Tag
	Occ = markerDB["occ"]
	// OccupancyIndicator Marker Tag
	OccupancyIndicator = markerDB["occupancyIndicator"]
	// Occupied Marker Tag
	Occupied = markerDB["occupied"]
	// Oil Marker Tag
	Oil = markerDB["oil"]
	// OpenLoop Marker Tag
	OpenLoop = markerDB["openLoop"]
	// Outside Marker Tag
	Outside = markerDB["outside"]
	// Parallel Marker Tag
	Parallel = markerDB["parallel"]
	// PerimeterHeat Marker Tag
	PerimeterHeat = markerDB["perimeterHeat"]
	// Pf Marker Tag
	Pf = markerDB["pf"]
	// Point Marker Tag
	Point = markerDB["point"]
	// Power Marker Tag
	Power = markerDB["power"]
	// Precipitation Marker Tag
	Precipitation = markerDB["precipitation"]
	// Pressure Marker Tag
	Pressure = markerDB["pressure"]
	// PressureDependent Marker Tag
	PressureDependent = markerDB["pressureDependent"]
	// PressureIndependent Marker Tag
	PressureIndependent = markerDB["pressureIndependent"]
	// PrimaryLoop Marker Tag
	PrimaryLoop = markerDB["primaryLoop"]
	// Pump Marker Tag
	Pump = markerDB["pump"]
	// Reactive Marker Tag
	Reactive = markerDB["reactive"]
	// Reciprocal Marker Tag
	Reciprocal = markerDB["reciprocal"]
	// Refrig Marker Tag
	Refrig = markerDB["refrig"]
	// Reheat Marker Tag
	Reheat = markerDB["reheat"]
	// Reheating Marker Tag
	Reheating = markerDB["reheating"]
	// Return Marker Tag
	Return = markerDB["return"]
	// Rooftop Marker Tag
	Rooftop = markerDB["rooftop"]
	// Run Marker Tag
	Run = markerDB["run"]
	// Screw Marker Tag
	Screw = markerDB["screw"]
	// SecondaryLoop Marker Tag
	SecondaryLoop = markerDB["secondaryLoop"]
	// Sensor Marker Tag
	Sensor = markerDB["sensor"]
	// Series Marker Tag
	Series = markerDB["series"]
	// SingleDuct Marker Tag
	SingleDuct = markerDB["singleDuct"]
	// Site Marker Tag
	Site = markerDB["site"]
	// SiteMeter Marker Tag
	SiteMeter = markerDB["siteMeter"]
	// SitePanel Marker Tag
	SitePanel = markerDB["sitePanel"]
	// Solar Marker Tag
	Solar = markerDB["solar"]
	// Sp Marker Tag
	Sp = markerDB["sp"]
	// Speed Marker Tag
	Speed = markerDB["speed"]
	// Standby Marker Tag
	Standby = markerDB["standby"]
	// Steam Marker Tag
	Steam = markerDB["steam"]
	// SteamHeat Marker Tag
	SteamHeat = markerDB["steamHeat"]
	// SteamPlant Marker Tag
	SteamPlant = markerDB["steamPlant"]
	// Sunrise Marker Tag
	Sunrise = markerDB["sunrise"]
	// Tank Marker Tag
	Tank = markerDB["tank"]
	// Temp Marker Tag
	Temp = markerDB["temp"]
	// Thd Marker Tag
	Thd = markerDB["thd"]
	// Total Marker Tag
	Total = markerDB["total"]
	// TripleDuct Marker Tag
	TripleDuct = markerDB["tripleDuct"]
	// Unocc Marker Tag
	Unocc = markerDB["unocc"]
	// Uv Marker Tag
	Uv = markerDB["uv"]
	// Valve Marker Tag
	Valve = markerDB["valve"]
	// VariableVolume Marker Tag
	VariableVolume = markerDB["variableVolume"]
	// Vav Marker Tag
	Vav = markerDB["vav"]
	// VavMode Marker Tag
	VavMode = markerDB["vavMode"]
	// VavZone Marker Tag
	VavZone = markerDB["vavZone"]
	// Vfd Marker Tag
	Vfd = markerDB["vfd"]
	// Visibility Marker Tag
	Visibility = markerDB["visibility"]
	// Volt Marker Tag
	Volt = markerDB["volt"]
	// Volume Marker Tag
	Volume = markerDB["volume"]
	// Water Marker Tag
	Water = markerDB["water"]
	// WaterCooled Marker Tag
	WaterCooled = markerDB["waterCooled"]
	// Weather Marker Tag
	Weather = markerDB["weather"]
	// WeatherCond Marker Tag
	WeatherCond = markerDB["weatherCond"]
	// WeatherPoint Marker Tag
	WeatherPoint = markerDB["weatherPoint"]
	// WetBulb Marker Tag
	WetBulb = markerDB["wetBulb"]
	// Wind Marker Tag
	Wind = markerDB["wind"]
	// Writable Marker Tag
	Writable = markerDB["writable"]
	// Zone Marker Tag
	Zone = markerDB["zone"]
)

// Marker is a helper function to generate a https://project-haystack.org/tag/ Marker is a helper function to generate a Tag
func Marker(l *gohaystack.Label) func() (*gohaystack.Label, *gohaystack.Value) {
	return func() (*gohaystack.Label, *gohaystack.Value) {
		return l, gohaystack.MarkerValue
	}
}
