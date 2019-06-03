/*

@Time : 2019/5/23
@Author : Jiangs

*/
package model

// DataSources.YAML MODEL
type Datasource struct {
	Name           string         `yaml:"name"`
	Type           string         `yaml:"type"`
	Url            string         `yaml:"url"`
	Database       string         `yaml:"database"`
	User           string         `yaml:"user"`
	SecureJsonData SecureJsonData `yaml:"secureJsonData"`
	JsonData       JsonData       `yaml:"jsonData"`
}

type SecureJsonData struct {
	Password string `yaml:"password"`
}

type JsonData struct {
	Sslmode         string `yaml:"sslmode"`
	MaxOpenConns    int    `yaml:"maxOpenConns"`
	MaxIdleConns    int    `yaml:"maxIdleConns"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
	PostgresVersion int    `yaml:"postgresVersion"`
	Timescaledb     bool   `yaml:"timescaledb"`
}

type PreDatasource struct {
	ApiVersion  int          `yaml:"apiVersion"`
	Datasources []Datasource `yaml:"datasources"`
}

//Notifiers.YAML MODEL
type PreNotifier struct {
	ApiVersion int        `yaml:"apiVersion"`
	Notifiers  []Notifier `yaml:"notifiers"`
}

type Notifier struct {
	Name       string   `yaml:"name"`
	Type       string   `yaml:"type"`
	Org_id     int      `yaml:"org_id"`
	Uid        string   `yaml:"uid"`
	Is_default bool     `yaml:"is_default"`
	Settings   Settings `yaml:"settings"`
}

type Settings struct {
	Addresses string `yaml:"addresses"`
}

// Dashboard.JSON MODEL
type Dashboard struct {
	Annotations   Annotations   `json:"annotations"`
	Editable      bool          `json:"editable"`
	GnetID        interface{}   `json:"gnetId"`
	GraphTooltip  int           `json:"graphTooltip"`
	Links         []interface{} `json:"links"`
	Panels        []Panels      `json:"panels"`
	Refresh       string        `json:"refresh"`
	SchemaVersion int           `json:"schemaVersion"`
	Style         string        `json:"style"`
	Tags          []interface{} `json:"tags"`
	Templating    Templating    `json:"templation"`
	Time          Time          `json:"time"`
	Timepicker    Timepicker    `json:"timepicker"`
	Timezone      string        `json:"timezone"`
	Title         string        `json:"title"`
	UID           string        `json:"uid"`
	Version       int           `json:"version"`
}

type Annotations struct {
	List []struct {
		BuiltIn    int    `json:"builtIn"`
		Datasource string `json:"datasource"`
		Enable     bool   `json:"enable"`
		Hide       bool   `json:"hide"`
		IconColor  string `json:"iconColor"`
		Name       string `json:"name"`
		Type       string `json:"type"`
	} `json:"list"`
}

type Templating struct {
	List []interface{} `json:"list"`
}

type Time struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type Timepicker struct {
	RefreshIntervals []string `json:"refresh_intervals"`
	TimeOptions      []string `json:"time_options"`
}

type Panels struct {
	CacheTimeout    interface{} `json:"cacheTimeout,omitempty"`
	ColorBackground bool        `json:"colorBackground,omitempty"`
	ColorValue      bool        `json:"colorValue,omitempty"`
	Colors          []string    `json:"colors,omitempty"`
	Datasource      string      `json:"datasource"`
	Format          string      `json:"format,omitempty"`
	Gauge           struct {
		MaxValue         int  `json:"maxValue"`
		MinValue         int  `json:"minValue"`
		Show             bool `json:"show"`
		ThresholdLabels  bool `json:"thresholdLabels"`
		ThresholdMarkers bool `json:"thresholdMarkers"`
	} `json:"gauge,omitempty"`
	GridPos struct {
		H int `json:"h"`
		W int `json:"w"`
		X int `json:"x"`
		Y int `json:"y"`
	} `json:"gridPos"`
	ID           int           `json:"id"`
	Interval     interface{}   `json:"interval,omitempty"`
	Links        []interface{} `json:"links"`
	MappingType  int           `json:"mappingType,omitempty"`
	MappingTypes []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"mappingTypes,omitempty"`
	MaxDataPoints   int         `json:"maxDataPoints,omitempty"`
	NullPointMode   string      `json:"nullPointMode,omitempty"`
	NullText        interface{} `json:"nullText,omitempty"`
	Postfix         string      `json:"postfix,omitempty"`
	PostfixFontSize string      `json:"postfixFontSize,omitempty"`
	Prefix          string      `json:"prefix,omitempty"`
	PrefixFontSize  string      `json:"prefixFontSize,omitempty"`
	RangeMaps       []struct {
		From string `json:"from"`
		Text string `json:"text"`
		To   string `json:"to"`
	} `json:"rangeMaps,omitempty"`
	Sparkline struct {
		FillColor string `json:"fillColor"`
		Full      bool   `json:"full"`
		LineColor string `json:"lineColor"`
		Show      bool   `json:"show"`
	} `json:"sparkline,omitempty"`
	TableColumn string `json:"tableColumn,omitempty"`
	Targets     []struct {
		Format       string        `json:"format"`
		Group        []interface{} `json:"group"`
		Hide         bool          `json:"hide"`
		MetricColumn string        `json:"metricColumn"`
		RawQuery     bool          `json:"rawQuery"`
		RawSQL       string        `json:"rawSql"`
		RefID        string        `json:"refId"`
		Select       [][]struct {
			Params []string `json:"params"`
			Type   string   `json:"type"`
		} `json:"select"`
		Table          string `json:"table"`
		TimeColumn     string `json:"timeColumn"`
		TimeColumnType string `json:"timeColumnType"`
		Where          []struct {
			Name   string        `json:"name"`
			Params []interface{} `json:"params"`
			Type   string        `json:"type"`
		} `json:"where"`
	} `json:"targets"`
	Thresholds    string `json:"thresholds,omitempty"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	ValueFontSize string `json:"valueFontSize,omitempty"`
	ValueMaps     []struct {
		Op    string `json:"op"`
		Text  string `json:"text"`
		Value string `json:"value"`
	} `json:"valueMaps,omitempty"`
	ValueName    string      `json:"valueName,omitempty"`
	ColorPostfix bool        `json:"colorPostfix,omitempty"`
	ColorPrefix  bool        `json:"colorPrefix,omitempty"`
	Decimals     interface{} `json:"decimals,omitempty"`
	TimeFrom     interface{} `json:"timeFrom,omitempty"`
	Transparent  bool        `json:"transparent,omitempty"`
	AliasColors  struct {
	} `json:"aliasColors,omitempty"`
	Bars       bool `json:"bars,omitempty"`
	DashLength int  `json:"dashLength,omitempty"`
	Dashes     bool `json:"dashes,omitempty"`
	Fill       int  `json:"fill,omitempty"`
	Legend     struct {
		Avg     bool `json:"avg"`
		Current bool `json:"current"`
		Max     bool `json:"max"`
		Min     bool `json:"min"`
		Show    bool `json:"show"`
		Total   bool `json:"total"`
		Values  bool `json:"values"`
	} `json:"legend,omitempty"`
	Lines           bool          `json:"lines,omitempty"`
	Linewidth       int           `json:"linewidth,omitempty"`
	Percentage      bool          `json:"percentage,omitempty"`
	Pointradius     int           `json:"pointradius,omitempty"`
	Points          bool          `json:"points,omitempty"`
	Renderer        string        `json:"renderer,omitempty"`
	SeriesOverrides []interface{} `json:"seriesOverrides,omitempty"`
	SpaceLength     int           `json:"spaceLength,omitempty"`
	Stack           bool          `json:"stack,omitempty"`
	SteppedLine     bool          `json:"steppedLine,omitempty"`
	TimeRegions     []interface{} `json:"timeRegions,omitempty"`
	TimeShift       interface{}   `json:"timeShift,omitempty"`
	Tooltip         struct {
		Shared    bool   `json:"shared"`
		Sort      int    `json:"sort"`
		ValueType string `json:"value_type"`
	} `json:"tooltip,omitempty"`
	Xaxis struct {
		Buckets interface{}   `json:"buckets"`
		Mode    string        `json:"mode"`
		Name    interface{}   `json:"name"`
		Show    bool          `json:"show"`
		Values  []interface{} `json:"values"`
	} `json:"xaxis,omitempty"`
	Yaxes []struct {
		Format  string      `json:"format"`
		Label   interface{} `json:"label"`
		LogBase int         `json:"logBase"`
		Max     interface{} `json:"max"`
		Min     interface{} `json:"min"`
		Show    bool        `json:"show"`
	} `json:"yaxes,omitempty"`
	Yaxis struct {
		Align      bool        `json:"align"`
		AlignLevel interface{} `json:"alignLevel"`
	} `json:"yaxis,omitempty"`
	Alert struct {
		Conditions []struct {
			Evaluator struct {
				Params []int  `json:"params"`
				Type   string `json:"type"`
			} `json:"evaluator"`
			Operator struct {
				Type string `json:"type"`
			} `json:"operator"`
			Query struct {
				Params []string `json:"params"`
			} `json:"query"`
			Reducer struct {
				Params []interface{} `json:"params"`
				Type   string        `json:"type"`
			} `json:"reducer"`
			Type string `json:"type"`
		} `json:"conditions"`
		ExecutionErrorState string        `json:"executionErrorState"`
		For                 string        `json:"for"`
		Frequency           string        `json:"frequency"`
		Handler             int           `json:"handler"`
		Name                string        `json:"name"`
		NoDataState         string        `json:"noDataState"`
		Notifications       []interface{} `json:"notifications"`
	} `json:"alert,omitempty"`
	Columns          []interface{} `json:"columns,omitempty"`
	FontSize         string        `json:"fontSize,omitempty"`
	HideTimeOverride bool          `json:"hideTimeOverride,omitempty"`
	PageSize         interface{}   `json:"pageSize,omitempty"`
	Scroll           bool          `json:"scroll,omitempty"`
	ShowHeader       bool          `json:"showHeader,omitempty"`
	Sort             struct {
		Col  int  `json:"col"`
		Desc bool `json:"desc"`
	} `json:"sort,omitempty"`
	Styles []struct {
		Alias      string        `json:"alias"`
		DateFormat string        `json:"dateFormat,omitempty"`
		Pattern    string        `json:"pattern"`
		Type       string        `json:"type"`
		ColorMode  interface{}   `json:"colorMode,omitempty"`
		Colors     []string      `json:"colors,omitempty"`
		Decimals   int           `json:"decimals,omitempty"`
		Thresholds []interface{} `json:"thresholds,omitempty"`
		Unit       string        `json:"unit,omitempty"`
	} `json:"styles,omitempty"`
	Transform string `json:"transform,omitempty"`
}
