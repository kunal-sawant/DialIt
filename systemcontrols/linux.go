//go:build linux

package systemcontrols

import (
	"fmt"
	"os/exec"
	"strings"
)

type linuxController struct {
	displayName string
}

func newPlatformController() (Controller, error) {
	cmd := exec.Command("sh", "-c", "xrandr | grep primary | cut -d' ' -f1")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get display info: %v", err)
	}
	displayName := strings.TrimSpace(string(output))

	return &linuxController{
		displayName: displayName,
	}, nil
}

func (l *linuxController) SetBrightness(level int) error {
	brightness := float64(level) / 100.0
	cmd := exec.Command("xrandr", "--output", l.displayName, "--brightness", fmt.Sprintf("%.2f", brightness))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set brightness: %v, output: %s", err, string(output))
	}
	return nil
}

func (l *linuxController) GetBrightness() (int, error) {
	// TODO: Impl

	// cmd := exec.Command("xrandr", "--verbose")
	// // output, err := cmd.Output()
	// // if err != nil {
	// // 	return 0, err
	// // }
	return 0, nil
}

func (l *linuxController) SetVolume(level int) error {
	cmd := exec.Command("pactl", "set-sink-volume", "@DEFAULT_SINK@", fmt.Sprintf("%d%%", level))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to set volume: %v, output: %s", err, string(output))
	}
	return nil
}

func (l *linuxController) GetVolume() (int, error) {
	//  TODO: Impl
	//  cmd := exec.Command("pactl", "get-sink-volume", "@DEFAULT_SINK@")
	//  output, err := cmd.Output()
	//  if err != nil {
	// 	 return 0, err
	//  }
	// Parse output to extract volume level
	// Implementation details here...
	return 0, nil
}
