package imaginary

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/go-yaml/yaml"
)

var (
	profilesEnv = "profiles.cfg"
	config      profileConfig
)

type profileConfig struct {
	isLoaded  bool
	Thumbnail sizeConfig `yaml:"thumbnail"`
	Small     sizeConfig `yaml:"small"`
	Medium    sizeConfig `yaml:"medium"`
}

type sizeConfig struct {
	Width  string `yaml:"width"`
	Height string `yaml:"height"`
}

func Resize(profileName string, originPath string, outputPath string) error {
	p, err := loadProfile(profileName)
	if err != nil {
		return err
	}

	var args = []string{
		originPath,
		"--output", outputPath,
		"--size", strings.Join([]string{p.Width, "x", p.Height}, ""),
		//"--crop",
	}

	lib, _ := getLib()
	exec.Command(lib, args...).Run()

	return nil
}

func AvailableProfiles() string {
	return strings.Join([]string{
		"thumbnail",
		"small",
		"medium",
	}, ", ")
}

func loadProfile(profileName string) (settings *sizeConfig, err error) {
	if config.isLoaded {
		return config.getSettings(profileName)
	}

	cwd, _ := os.Getwd()

	data, err := ioutil.ReadFile(filepath.Join(cwd, profilesEnv))
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	config.isLoaded = true

	return config.getSettings(profileName)
}

func getLib() (string, error) {
	return exec.LookPath("vipsthumbnail")
}

func (p *profileConfig) getSettings(field string) (settings *sizeConfig, err error) {
	f := reflect.Indirect(reflect.ValueOf(p)).FieldByName(strings.Title(field))
	if f.IsValid() {
		profile := &sizeConfig{
			f.FieldByName("Width").String(),
			f.FieldByName("Height").String(),
		}

		return profile, nil
	}

	return nil, errors.New("Profile not found")
}
