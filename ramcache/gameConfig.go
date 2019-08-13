package ramcache

import (
	"encoding/json"
)

func getPlatformCofig(platform, gamePlatformCode string, cfgObj interface{}) error {
	gamePlatformAPiConfigs, _ := GamePlatformAPiConfigs.Load(platform)
	gamePlatformAPiConfigDict := gamePlatformAPiConfigs.(map[string]string)
	cfgVal, cfgExist := gamePlatformAPiConfigDict[gamePlatformCode]
	if cfgExist {
		return json.Unmarshal([]byte(cfgVal), &cfgObj)
	}
	return nil
}

type FGConfig struct {
	FGURL string
	//FG平台开发者ID
	FgMerchantName string
	FgMerchantCode string
}

func GetFGConfig(platform string) FGConfig {
	cfg := FGConfig{}
	getPlatformCofig(platform, "FG", &cfg)
	return cfg
}

type AEConfig struct {
	EAURL  string
	EAKEY  string
	SITEID string
}

func GetAEConfig(platform string) AEConfig {
	cfg := AEConfig{}
	getPlatformCofig(platform, "AE", &cfg)
	return cfg
}

type MGConfig struct {
	MGURL    string
	TOKENURL string
	MGNAME   string
	MGKEY    string
}

func GetMGConfig(platform string) MGConfig {
	cfg := MGConfig{}
	getPlatformCofig(platform, "MG", &cfg)
	return cfg
}

type KYConfig struct {
	KYURL    string
	KYBETURL string
	TRIALURL string
	AGENT    string
	LINE     string
	DESKEY   string
	MD5KEY   string
}

func GetKYConfig(platform string) KYConfig {
	cfg := KYConfig{}
	getPlatformCofig(platform, "KY", &cfg)
	return cfg
}

type AGConfig struct {
	GIURL   string
	GCIURL  string
	CAGENT  string
	MD5KEY  string
	DESKEY  string
	CAGENTQ string
	FTPNAME string
	FTPPWD  string
	FTPURL  string
}

func GetAGConfig(platform string) AGConfig {
	cfg := AGConfig{}
	getPlatformCofig(platform, "AG", &cfg)
	return cfg
}

type LYConfig struct {
	LYURL    string
	LYBETURL string
	TRIALURL string
	AGENT    string
	LINE     string
	DESKEY   string
	MD5KEY   string
}

func GetLYConfig(platform string) LYConfig {
	cfg := LYConfig{}
	getPlatformCofig(platform, "LY", &cfg)
	return cfg
}

type NWConfig struct {
	URL      string
	BETURL   string
	TRIALURL string
	AGENT    string
	LINE     string
	DESKEY   string
	MD5KEY   string
}

func GetNWConfig(platform string) NWConfig {
	cfg := NWConfig{}
	getPlatformCofig(platform, "NW", &cfg)
	return cfg
}

type VGConfig struct {
	CHANNEL  string
	URL      string
	BETURL   string
	PASSWORD string
}

func GetVGConfig(platform string) VGConfig {
	cfg := VGConfig{}
	getPlatformCofig(platform, "VG", &cfg)
	return cfg
}

type JDBConfig struct {
	DC     string
	URL    string
	IV     string
	KEY    string
	PARENT string
}

func GetJDBConfig(platform string) JDBConfig {
	cfg := JDBConfig{}
	getPlatformCofig(platform, "JDB", &cfg)
	return cfg
}

type OGConfig struct {
	XOPERATOR string
	URL       string
	BETURL    string
	XKEY      string
	PREFIX    string
}

func GetOGConfig(platform string) OGConfig {
	cfg := OGConfig{}
	getPlatformCofig(platform, "OG", &cfg)
	return cfg
}

type UGConfig struct {
	BaseUrl string
	Key     string
}

func GetUGConfig(platform string) UGConfig {
	cfg := UGConfig{}
	getPlatformCofig(platform, "UG", &cfg)
	return cfg
}
