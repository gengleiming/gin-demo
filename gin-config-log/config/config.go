package config

// Configuration mapstructure 标签需对应 config.yaml 中的配置名称， viper 会识别标签
type Configuration struct {
	App App `mapstructure:"app" json:"app" yaml:"app"`
	Log Log `mapstructure:"log" json:"log" yaml:"log"`
}
