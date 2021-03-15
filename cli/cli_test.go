package cli

import (
	"fmt"
	"os"
	"testing"
)

func TestVersionParseNoArgs(t *testing.T) {
	got, _, _ := ParseVersion([]string{})
	want := NewVersionArguments(false, Version)

	if got != want {
		t.Errorf("ParseVersion([]string{}) got %+v but wanted %+v", got, want)
	}
}

func TestVersionParseArgsHelp(t *testing.T) {
	got, _, _ := ParseVersion([]string{"--help"})
	want := NewVersionArguments(true, Version)

	if got != want {
		t.Errorf("ParseVersion([]string{\"--help\"}) got %+v but wanted %+v", got, want)
	}
}

func TestStartParseNoArgs(t *testing.T) {
	got, _, _ := ParseStart([]string{})
	configFilePath, err := defaultConfigFilePath()
	want := NewStartArguments(false, configFilePath)

	if err != nil {
		t.Fatalf("Could not get default configuration file path:\n%s", err.Error())
	}

	if got != want {
		t.Errorf("ParseStart([]string{}) got %+v but wanted %+v", got, want)
	}
}

func TestStartParseArgsHelp(t *testing.T) {
	got, _, _ := ParseStart([]string{"--help"})
	configFilePath, err := defaultConfigFilePath()
	want := NewStartArguments(true, configFilePath)

	if err != nil {
		t.Fatalf("Could not get default configuration file path:\n%s", err.Error())
	}

	if got != want {
		t.Errorf("ParseStart([]string{\"--help\"}) got %+v but wanted %+v", got, want)
	}
}

func TestLoadServerConfig(t *testing.T) {
	cwd, err := os.Getwd()

	if err != nil {
		t.Fatalf("Could not get current working directory path:\n%s", err.Error())
	}

	got, err := LoadServerConfig(fmt.Sprintf("%s\\cli_test\\test_config.yaml", cwd))
	want := NewServerConfig("0.0.0.0", 8081, 45041, "./worlds/alt-world.json", 32, true, 1024)

	if err != nil {
		t.Fatalf("Could not open file:\n%s", err.Error())
	}

	if got != want {
		t.Errorf("LoadServerConfig(\"./cli_test/test_config.yaml\") got %+v but wanted %+v", got, want)
	}
}

func TestLoadServerConfigWrongPath(t *testing.T) {
	cwd, err := os.Getwd()

	if err != nil {
		t.Fatalf("Could not get current working directory path:\n%s", err.Error())
	}

	_, err = LoadServerConfig(fmt.Sprintf("%s\\cli_test\\inexistent.yaml", cwd))

	if err == nil {
		t.Errorf("LoadServerConfig() called with invalid path didn't spawn error")
	}
}
