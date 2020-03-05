package config

type Config struct {
	Ip     string  `yaml:"ip"`
	Period float64 `yaml:"period"`

	JoinSpace      Trajectory `yaml:"joinSpace"`
	CartesianSpace Trajectory `yaml:"cartesianSpace"`

	Workspace [][]float64 `json:"workspace"`

	Motors []Motor `yaml:"motors"`
}

type Trajectory struct {
	Acceleration float64 `yaml:"acceleration"`
	Deceleration float64 `yaml:"deceleration"`
	Speed        float64 `yaml:"speed"`
	MinSpeed     float64 `yaml:"minSpeed"`
}

type Motor struct {
	Drum         Drum      `yaml:"drum"`
	ExitPoint    []float64 `yaml:"exitPoint"`
	RollerRadius float64   `yaml:"rollerRadius"`
}

type Drum struct {
	H float64 `yaml:"h"`
	R float64 `yaml:"r"`
}
