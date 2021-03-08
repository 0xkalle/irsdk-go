package session

import "time"

type SessionData struct {
	WeekendInfo   WeekendInfo   `yaml:"WeekendInfo"`
	SessionInfo   SessionInfo   `yaml:"SessionInfo"`
	CameraInfo    CameraInfo    `yaml:"CameraInfo"`
	RadioInfo     RadioInfo     `yaml:"RadioInfo"`
	DriverInfo    DriverInfo    `yaml:"DriverInfo"`
	SplitTimeInfo SplitTimeInfo `yaml:"SplitTimeInfo"`
	CarSetup      CarSetup      `yaml:"CarSetup"`
}

type WeekendOptions struct {
	NumStarters                int       `yaml:"NumStarters"`
	StartingGrid               string    `yaml:"StartingGrid"`
	QualifyScoring             string    `yaml:"QualifyScoring"`
	CourseCautions             string    `yaml:"CourseCautions"`
	StandingStart              int       `yaml:"StandingStart"`
	ShortParadeLap             int       `yaml:"ShortParadeLap"`
	Restarts                   string    `yaml:"Restarts"`
	WeatherType                string    `yaml:"WeatherType"`
	Skies                      string    `yaml:"Skies"`
	WindDirection              string    `yaml:"WindDirection"`
	WindSpeed                  string    `yaml:"WindSpeed"`
	WeatherTemp                string    `yaml:"WeatherTemp"`
	RelativeHumidity           string    `yaml:"RelativeHumidity"`
	FogLevel                   string    `yaml:"FogLevel"`
	TimeOfDay                  string    `yaml:"TimeOfDay"`
	Date                       time.Time `yaml:"Date"`
	EarthRotationSpeedupFactor int       `yaml:"EarthRotationSpeedupFactor"`
	Unofficial                 int       `yaml:"Unofficial"`
	CommercialMode             string    `yaml:"CommercialMode"`
	NightMode                  string    `yaml:"NightMode"`
	IsFixedSetup               int       `yaml:"IsFixedSetup"`
	StrictLapsChecking         string    `yaml:"StrictLapsChecking"`
	HasOpenRegistration        int       `yaml:"HasOpenRegistration"`
	HardcoreLevel              int       `yaml:"HardcoreLevel"`
	NumJokerLaps               int       `yaml:"NumJokerLaps"`
	IncidentLimit              string    `yaml:"IncidentLimit"`
	FastRepairsLimit           string    `yaml:"FastRepairsLimit"`
	GreenWhiteCheckeredLimit   int       `yaml:"GreenWhiteCheckeredLimit"`
}
type TelemetryOptions struct {
	TelemetryDiskFile string `yaml:"TelemetryDiskFile"`
}
type WeekendInfo struct {
	TrackName              string           `yaml:"TrackName"`
	TrackID                int              `yaml:"TrackID"`
	TrackLength            string           `yaml:"TrackLength"`
	TrackDisplayName       string           `yaml:"TrackDisplayName"`
	TrackDisplayShortName  string           `yaml:"TrackDisplayShortName"`
	TrackConfigName        interface{}      `yaml:"TrackConfigName"`
	TrackCity              string           `yaml:"TrackCity"`
	TrackCountry           string           `yaml:"TrackCountry"`
	TrackAltitude          string           `yaml:"TrackAltitude"`
	TrackLatitude          string           `yaml:"TrackLatitude"`
	TrackLongitude         string           `yaml:"TrackLongitude"`
	TrackNorthOffset       string           `yaml:"TrackNorthOffset"`
	TrackNumTurns          int              `yaml:"TrackNumTurns"`
	TrackPitSpeedLimit     string           `yaml:"TrackPitSpeedLimit"`
	TrackType              string           `yaml:"TrackType"`
	TrackDirection         string           `yaml:"TrackDirection"`
	TrackWeatherType       string           `yaml:"TrackWeatherType"`
	TrackSkies             string           `yaml:"TrackSkies"`
	TrackSurfaceTemp       string           `yaml:"TrackSurfaceTemp"`
	TrackAirTemp           string           `yaml:"TrackAirTemp"`
	TrackAirPressure       string           `yaml:"TrackAirPressure"`
	TrackWindVel           string           `yaml:"TrackWindVel"`
	TrackWindDir           string           `yaml:"TrackWindDir"`
	TrackRelativeHumidity  string           `yaml:"TrackRelativeHumidity"`
	TrackFogLevel          string           `yaml:"TrackFogLevel"`
	TrackCleanup           int              `yaml:"TrackCleanup"`
	TrackDynamicTrack      int              `yaml:"TrackDynamicTrack"`
	TrackVersion           string           `yaml:"TrackVersion"`
	SeriesID               int              `yaml:"SeriesID"`
	SeasonID               int              `yaml:"SeasonID"`
	SessionID              int              `yaml:"SessionID"`
	SubSessionID           int              `yaml:"SubSessionID"`
	LeagueID               int              `yaml:"LeagueID"`
	Official               int              `yaml:"Official"`
	RaceWeek               int              `yaml:"RaceWeek"`
	EventType              string           `yaml:"EventType"`
	Category               string           `yaml:"Category"`
	SimMode                string           `yaml:"SimMode"`
	TeamRacing             int              `yaml:"TeamRacing"`
	MinDrivers             int              `yaml:"MinDrivers"`
	MaxDrivers             int              `yaml:"MaxDrivers"`
	DCRuleSet              string           `yaml:"DCRuleSet"`
	QualifierMustStartRace int              `yaml:"QualifierMustStartRace"`
	NumCarClasses          int              `yaml:"NumCarClasses"`
	NumCarTypes            int              `yaml:"NumCarTypes"`
	HeatRacing             int              `yaml:"HeatRacing"`
	BuildType              string           `yaml:"BuildType"`
	BuildTarget            string           `yaml:"BuildTarget"`
	BuildVersion           string           `yaml:"BuildVersion"`
	WeekendOptions         WeekendOptions   `yaml:"WeekendOptions"`
	TelemetryOptions       TelemetryOptions `yaml:"TelemetryOptions"`
}
type ResultsPositions struct {
	Position          int     `yaml:"Position"`
	ClassPosition     int     `yaml:"ClassPosition"`
	CarIdx            int     `yaml:"CarIdx"`
	Lap               int     `yaml:"Lap"`
	Time              float64 `yaml:"Time"`
	FastestLap        int     `yaml:"FastestLap"`
	FastestTime       float64 `yaml:"FastestTime"`
	LastTime          int     `yaml:"LastTime"`
	LapsLed           int     `yaml:"LapsLed"`
	LapsComplete      int     `yaml:"LapsComplete"`
	JokerLapsComplete int     `yaml:"JokerLapsComplete"`
	LapsDriven        int     `yaml:"LapsDriven"`
	Incidents         int     `yaml:"Incidents"`
	ReasonOutID       int     `yaml:"ReasonOutId"`
	ReasonOutStr      string  `yaml:"ReasonOutStr"`
}
type ResultsFastestLap struct {
	CarIdx      int     `yaml:"CarIdx"`
	FastestLap  int     `yaml:"FastestLap"`
	FastestTime float64 `yaml:"FastestTime"`
}
type Sessions struct {
	SessionNum              int                 `yaml:"SessionNum"`
	SessionLaps             string              `yaml:"SessionLaps"`
	SessionTime             string              `yaml:"SessionTime"`
	SessionNumLapsToAvg     int                 `yaml:"SessionNumLapsToAvg"`
	SessionType             string              `yaml:"SessionType"`
	SessionTrackRubberState string              `yaml:"SessionTrackRubberState"`
	SessionName             string              `yaml:"SessionName"`
	SessionSubType          interface{}         `yaml:"SessionSubType"`
	SessionSkipped          int                 `yaml:"SessionSkipped"`
	SessionRunGroupsUsed    int                 `yaml:"SessionRunGroupsUsed"`
	ResultsPositions        []ResultsPositions  `yaml:"ResultsPositions"`
	ResultsFastestLap       []ResultsFastestLap `yaml:"ResultsFastestLap"`
	ResultsAverageLapTime   int                 `yaml:"ResultsAverageLapTime"`
	ResultsNumCautionFlags  int                 `yaml:"ResultsNumCautionFlags"`
	ResultsNumCautionLaps   int                 `yaml:"ResultsNumCautionLaps"`
	ResultsNumLeadChanges   int                 `yaml:"ResultsNumLeadChanges"`
	ResultsLapsComplete     int                 `yaml:"ResultsLapsComplete"`
	ResultsOfficial         int                 `yaml:"ResultsOfficial"`
}
type SessionInfo struct {
	Sessions []Sessions `yaml:"Sessions"`
}
type Cameras struct {
	CameraNum  int    `yaml:"CameraNum"`
	CameraName string `yaml:"CameraName"`
}
type Groups struct {
	GroupNum  int       `yaml:"GroupNum"`
	GroupName string    `yaml:"GroupName"`
	Cameras   []Cameras `yaml:"Cameras"`
	IsScenic  bool      `yaml:"IsScenic,omitempty"`
}
type CameraInfo struct {
	Groups []Groups `yaml:"Groups"`
}
type Frequencies struct {
	FrequencyNum  int    `yaml:"FrequencyNum"`
	FrequencyName string `yaml:"FrequencyName"`
	Priority      int    `yaml:"Priority"`
	CarIdx        int    `yaml:"CarIdx"`
	EntryIdx      int    `yaml:"EntryIdx"`
	ClubID        int    `yaml:"ClubID"`
	CanScan       int    `yaml:"CanScan"`
	CanSquawk     int    `yaml:"CanSquawk"`
	Muted         int    `yaml:"Muted"`
	IsMutable     int    `yaml:"IsMutable"`
	IsDeletable   int    `yaml:"IsDeletable"`
}
type Radios struct {
	RadioNum            int           `yaml:"RadioNum"`
	HopCount            int           `yaml:"HopCount"`
	NumFrequencies      int           `yaml:"NumFrequencies"`
	TunedToFrequencyNum int           `yaml:"TunedToFrequencyNum"`
	ScanningIsOn        int           `yaml:"ScanningIsOn"`
	Frequencies         []Frequencies `yaml:"Frequencies"`
}
type RadioInfo struct {
	SelectedRadioNum int      `yaml:"SelectedRadioNum"`
	Radios           []Radios `yaml:"Radios"`
}
type Drivers struct {
	CarIdx                  int    `yaml:"CarIdx"`
	UserName                string `yaml:"UserName"`
	AbbrevName              string `yaml:"AbbrevName"`
	Initials                string `yaml:"Initials"`
	UserID                  int    `yaml:"UserID"`
	TeamID                  int    `yaml:"TeamID"`
	TeamName                string `yaml:"TeamName"`
	CarNumber               string `yaml:"CarNumber"`
	CarNumberRaw            int    `yaml:"CarNumberRaw"`
	CarPath                 string `yaml:"CarPath"`
	CarClassID              int    `yaml:"CarClassID"`
	CarID                   int    `yaml:"CarID"`
	CarIsPaceCar            int    `yaml:"CarIsPaceCar"`
	CarIsAI                 int    `yaml:"CarIsAI"`
	CarScreenName           string `yaml:"CarScreenName"`
	CarScreenNameShort      string `yaml:"CarScreenNameShort"`
	CarClassShortName       string `yaml:"CarClassShortName"`
	CarClassRelSpeed        int    `yaml:"CarClassRelSpeed"`
	CarClassLicenseLevel    int    `yaml:"CarClassLicenseLevel"`
	CarClassMaxFuelPct      string `yaml:"CarClassMaxFuelPct"`
	CarClassWeightPenalty   string `yaml:"CarClassWeightPenalty"`
	CarClassPowerAdjust     string `yaml:"CarClassPowerAdjust"`
	CarClassDryTireSetLimit string `yaml:"CarClassDryTireSetLimit"`
	CarClassColor           int    `yaml:"CarClassColor"`
	IRating                 int    `yaml:"IRating"`
	LicLevel                int    `yaml:"LicLevel"`
	LicSubLevel             int    `yaml:"LicSubLevel"`
	LicString               string `yaml:"LicString"`
	LicColor                string `yaml:"LicColor"` // Todo investigate returned: 0xundefined
	IsSpectator             int    `yaml:"IsSpectator"`
	CarDesignStr            string `yaml:"CarDesignStr"`
	HelmetDesignStr         string `yaml:"HelmetDesignStr"`
	SuitDesignStr           string `yaml:"SuitDesignStr"`
	CarNumberDesignStr      string `yaml:"CarNumberDesignStr"`
	CarSponsor1             int    `yaml:"CarSponsor_1"`
	CarSponsor2             int    `yaml:"CarSponsor_2"`
	ClubName                string `yaml:"ClubName"`
	DivisionName            string `yaml:"DivisionName"`
	CurDriverIncidentCount  int    `yaml:"CurDriverIncidentCount"`
	TeamIncidentCount       int    `yaml:"TeamIncidentCount"`
}
type DriverInfo struct {
	DriverCarIdx              int       `yaml:"DriverCarIdx"`
	DriverUserID              int       `yaml:"DriverUserID"`
	PaceCarIdx                int       `yaml:"PaceCarIdx"`
	DriverHeadPosX            float64   `yaml:"DriverHeadPosX"`
	DriverHeadPosY            float64   `yaml:"DriverHeadPosY"`
	DriverHeadPosZ            float64   `yaml:"DriverHeadPosZ"`
	DriverCarIdleRPM          int       `yaml:"DriverCarIdleRPM"`
	DriverCarRedLine          int       `yaml:"DriverCarRedLine"`
	DriverCarEngCylinderCount int       `yaml:"DriverCarEngCylinderCount"`
	DriverCarFuelKgPerLtr     float64   `yaml:"DriverCarFuelKgPerLtr"`
	DriverCarFuelMaxLtr       int       `yaml:"DriverCarFuelMaxLtr"`
	DriverCarMaxFuelPct       int       `yaml:"DriverCarMaxFuelPct"`
	DriverCarSLFirstRPM       int       `yaml:"DriverCarSLFirstRPM"`
	DriverCarSLShiftRPM       int       `yaml:"DriverCarSLShiftRPM"`
	DriverCarSLLastRPM        int       `yaml:"DriverCarSLLastRPM"`
	DriverCarSLBlinkRPM       int       `yaml:"DriverCarSLBlinkRPM"`
	DriverCarVersion          string    `yaml:"DriverCarVersion"`
	DriverPitTrkPct           float64   `yaml:"DriverPitTrkPct"`
	DriverCarEstLapTime       float64   `yaml:"DriverCarEstLapTime"`
	DriverSetupName           string    `yaml:"DriverSetupName"`
	DriverSetupIsModified     int       `yaml:"DriverSetupIsModified"`
	DriverSetupLoadTypeName   string    `yaml:"DriverSetupLoadTypeName"`
	DriverSetupPassedTech     int       `yaml:"DriverSetupPassedTech"`
	DriverIncidentCount       int       `yaml:"DriverIncidentCount"`
	Drivers                   []Drivers `yaml:"Drivers"`
}
type Sectors struct {
	SectorNum      int `yaml:"SectorNum"`
	SectorStartPct int `yaml:"SectorStartPct"`
}
type SplitTimeInfo struct {
	Sectors []Sectors `yaml:"Sectors"`
}
type TiresAeroLeftFront struct {
	StartingPressure string `yaml:"StartingPressure"`
	LastHotPressure  string `yaml:"LastHotPressure"`
	LastTempsOMI     string `yaml:"LastTempsOMI"`
	TreadRemaining   string `yaml:"TreadRemaining"`
}
type TiresAeroLeftRear struct {
	StartingPressure string `yaml:"StartingPressure"`
	LastHotPressure  string `yaml:"LastHotPressure"`
	LastTempsOMI     string `yaml:"LastTempsOMI"`
	TreadRemaining   string `yaml:"TreadRemaining"`
}
type TiresAeroRightFront struct {
	StartingPressure string `yaml:"StartingPressure"`
	LastHotPressure  string `yaml:"LastHotPressure"`
	LastTempsIMO     string `yaml:"LastTempsIMO"`
	TreadRemaining   string `yaml:"TreadRemaining"`
}
type TiresAeroRightRear struct {
	StartingPressure string `yaml:"StartingPressure"`
	LastHotPressure  string `yaml:"LastHotPressure"`
	LastTempsIMO     string `yaml:"LastTempsIMO"`
	TreadRemaining   string `yaml:"TreadRemaining"`
}
type AeroBalanceCalc struct {
	FrontRhAtSpeed string `yaml:"FrontRhAtSpeed"`
	RearRhAtSpeed  string `yaml:"RearRhAtSpeed"`
	WingSetting    string `yaml:"WingSetting"`
	FrontDownforce string `yaml:"FrontDownforce"`
}
type TiresAero struct {
	LeftFront       TiresAeroLeftFront  `yaml:"LeftFront"`
	LeftRear        TiresAeroLeftRear   `yaml:"LeftRear"`
	RightFront      TiresAeroRightFront `yaml:"RightFront"`
	RightRear       TiresAeroRightRear  `yaml:"RightRear"`
	AeroBalanceCalc AeroBalanceCalc     `yaml:"AeroBalanceCalc"`
}
type Front struct {
	ArbBlades      string `yaml:"ArbBlades"`
	ToeIn          string `yaml:"ToeIn"`
	FrontMasterCyl string `yaml:"FrontMasterCyl"`
	RearMasterCyl  string `yaml:"RearMasterCyl"`
	BrakePads      string `yaml:"BrakePads"`
	CrossWeight    string `yaml:"CrossWeight"`
}
type LeftFront struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	LsCompDamping     string `yaml:"LsCompDamping"`
	HsCompDamping     string `yaml:"HsCompDamping"`
	LsRbdDamping      string `yaml:"LsRbdDamping"`
	HsRbdDamping      string `yaml:"HsRbdDamping"`
	Camber            string `yaml:"Camber"`
	Caster            string `yaml:"Caster"`
}
type LeftRear struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	LsCompDamping     string `yaml:"LsCompDamping"`
	HsCompDamping     string `yaml:"HsCompDamping"`
	LsRbdDamping      string `yaml:"LsRbdDamping"`
	HsRbdDamping      string `yaml:"HsRbdDamping"`
	Camber            string `yaml:"Camber"`
	ToeIn             string `yaml:"ToeIn"`
}
type InCarDials struct {
	BrakePressureBias      string `yaml:"BrakePressureBias"`
	AbsSetting             string `yaml:"AbsSetting"`
	TractionControlSetting string `yaml:"TractionControlSetting"`
	EngineMapSetting       string `yaml:"EngineMapSetting"`
}
type RightFront struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	LsCompDamping     string `yaml:"LsCompDamping"`
	HsCompDamping     string `yaml:"HsCompDamping"`
	LsRbdDamping      string `yaml:"LsRbdDamping"`
	HsRbdDamping      string `yaml:"HsRbdDamping"`
	Camber            string `yaml:"Camber"`
	Caster            string `yaml:"Caster"`
}
type RightRear struct {
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	SpringRate        string `yaml:"SpringRate"`
	LsCompDamping     string `yaml:"LsCompDamping"`
	HsCompDamping     string `yaml:"HsCompDamping"`
	LsRbdDamping      string `yaml:"LsRbdDamping"`
	HsRbdDamping      string `yaml:"HsRbdDamping"`
	Camber            string `yaml:"Camber"`
	ToeIn             string `yaml:"ToeIn"`
}
type Rear struct {
	FuelLevel   string `yaml:"FuelLevel"`
	ArbBlades   string `yaml:"ArbBlades"`
	DiffPreload string `yaml:"DiffPreload"`
	WingAngle   string `yaml:"WingAngle"`
}
type Chassis struct {
	Front      Front      `yaml:"Front"`
	LeftFront  LeftFront  `yaml:"LeftFront"`
	LeftRear   LeftRear   `yaml:"LeftRear"`
	InCarDials InCarDials `yaml:"InCarDials"`
	RightFront RightFront `yaml:"RightFront"`
	RightRear  RightRear  `yaml:"RightRear"`
	Rear       Rear       `yaml:"Rear"`
}

type SuspensionFront struct {
	ToeIn       string `yaml:"ToeIn"`
	CrossWeight string `yaml:"CrossWeight"`
	AntiRollBar string `yaml:"AntiRollBar"`
}

type SuspensionLeftFront struct {
	ColdPressure      string `yaml:"ColdPressure"`
	LastHotPressure   string `yaml:"LastHotPressure"`
	LastTempsOMI      string `yaml:"LastTempsOMI"`
	TreadRemaining    string `yaml:"TreadRemaining"`
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	Camber            string `yaml:"Camber"`
}

type SuspensionLeftRear struct {
	ColdPressure      string `yaml:"ColdPressure"`
	LastHotPressure   string `yaml:"LastHotPressure"`
	LastTempsOMI      string `yaml:"LastTempsOMI"`
	TreadRemaining    string `yaml:"TreadRemaining"`
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	Camber            string `yaml:"Camber"`
}

type SuspensionRightFront struct {
	ColdPressure      string `yaml:"ColdPressure"`
	LastHotPressure   string `yaml:"LastHotPressure"`
	LastTempsIMO      string `yaml:"LastTempsIMO"`
	TreadRemaining    string `yaml:"TreadRemaining"`
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	Camber            string `yaml:"Camber"`
}

type SuspensionRightRear struct {
	ColdPressure      string `yaml:"ColdPressure"`
	LastHotPressure   string `yaml:"LastHotPressure"`
	LastTempsIMO      string `yaml:"LastTempsIMO"`
	TreadRemaining    string `yaml:"TreadRemaining"`
	CornerWeight      string `yaml:"CornerWeight"`
	RideHeight        string `yaml:"RideHeight"`
	SpringPerchOffset string `yaml:"SpringPerchOffset"`
	BumpStiffness     string `yaml:"BumpStiffness"`
	ReboundStiffness  string `yaml:"ReboundStiffness"`
	Camber            string `yaml:"Camber"`
}

type SuspensionRear struct {
	FuelLevel   string `yaml:"FuelLevel"`
	ToeIn       string `yaml:"ToeIn"`
	AntiRollBar string `yaml:"AntiRollBar"`
}

type Suspension struct {
	Front      SuspensionFront      `yaml:"Front"`
	LeftFront  SuspensionLeftFront  `yaml:"LeftFront"`
	LeftRear   SuspensionLeftRear   `yaml:"LeftRear"`
	RightFront SuspensionRightFront `yaml:"RightFront"`
	RightRear  SuspensionRightRear  `yaml:"RightRear"`
	Rear       SuspensionRear       `yaml:"Rear"`
}

type CarSetup struct {
	UpdateCount int        `yaml:"UpdateCount"`
	TiresAero   TiresAero  `yaml:"TiresAero"`
	Chassis     Chassis    `yaml:"Chassis"`
	Suspension  Suspension `yaml:"Suspension"`
}
