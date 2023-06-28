// Package utils is a useful utils that boxie use. Methods are usually exported
package utils

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/twelvee/boxie/pkg/boxie/structs"
)

const saveDir = "/tmp/boxie_saves"
const savesFile = "/tmp/boxie_saves/save"

func EnsureSaveFileAvailable() error {
	// TODO: check useless calls of this method
	_, err := os.Stat(savesFile)
	if os.IsNotExist(err) {
		os.Mkdir(saveDir, 0750)
		dum, err := json.Marshal([]structs.Environment{})
		if err != nil {
			return err
		}
		err = os.WriteFile(savesFile, dum, 0644)
		if err != nil {
			return err
		}
	} else {
		return err
	}
	return nil
}

// IsBoxSaved check if the box is already saved or not
func IsBoxSaved(environmentID string, sbox structs.Box) (bool, error) {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return false, err
	}
	content, err := os.ReadFile(savesFile)
	if err != nil {
		return false, err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)
	if err != nil {
		return false, err
	}

	for _, env := range targets {
		if env.ID == environmentID {
			for _, box := range env.Boxes {
				if cmp.Equal(box, sbox) {
					return true, nil
				}
			}
			return false, nil
		}
	}
	return false, nil
}

// IsEnvironmentSaved check if the environment is already saved or not
func IsEnvironmentSaved(id string) (bool, error) {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return false, err
	}
	content, err := os.ReadFile(savesFile)
	if err != nil {
		return false, err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)
	if err != nil {
		return false, err
	}

	for _, env := range targets {
		if env.ID == id {
			return true, nil
		}
	}
	return false, nil
}

// SaveEnvironment will save your environment to tmp folder
func SaveEnvironment(environment structs.Environment) error {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return err
	}

	isSavedAlready, err := IsEnvironmentSaved(environment.ID)
	if err != nil {
		return err
	}

	if !isSavedAlready {
		content, err := os.ReadFile(savesFile)
		if err != nil {
			return err
		}
		targets := []structs.Environment{}
		err = json.Unmarshal(content, &targets)

		targets = append(targets, environment)
		content, err = json.Marshal(targets)
		err = os.WriteFile(savesFile, content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetEnvironment will return your environment from tmp folder
func GetEnvironment(id string) (*structs.Environment, error) {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(savesFile)
	if err != nil {
		return nil, err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)
	if err != nil {
		return nil, err
	}

	for _, env := range targets {
		if env.ID == strings.TrimSpace(id) {
			return &env, nil
		}
	}
	return nil, errors.New("Environment not found.")
}

// GetEnvironments will return all your environments from tmp folder
func GetEnvironments() ([]structs.Environment, error) {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(savesFile)
	if err != nil {
		return nil, err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)
	if err != nil {
		return nil, err
	}

	return targets, nil
}

// GetBox will return your box from tmp folder
func GetBox(environmentID string, boxName string, boxNamespace string) (*structs.Box, error) {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return nil, err
	}
	content, err := os.ReadFile(savesFile)
	if err != nil {
		return nil, err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)
	if err != nil {
		return nil, err
	}

	for _, env := range targets {
		if env.ID == environmentID {
			for _, box := range env.Boxes {
				if box.Name == boxName && box.Namespace == boxNamespace {
					return &box, nil
				}
			}
			return nil, nil
		}
	}
	return nil, nil
}

// SaveBox will save your box to tmp folder
func SaveBox(box structs.Box, environmentID string) error {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return err
	}
	saved, err := IsEnvironmentSaved(environmentID)
	if err != nil || !saved {
		return err
	}
	saved, err = IsBoxSaved(environmentID, box)
	if err != nil || saved {
		// already saved on this environment id or something went wrong
		return err
	}

	content, err := os.ReadFile(savesFile)
	if err != nil {
		return err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)
	for i, env := range targets {
		if env.ID == environmentID {
			targets[i].Boxes = append(env.Boxes, box)
		}
	}

	content, err = json.Marshal(targets)
	err = os.WriteFile(savesFile, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

// RemoveBox will remove your box from tmp folder
func RemoveBox(box structs.Box, environmentID string) error {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return err
	}
	saved, err := IsEnvironmentSaved(environmentID)
	if err != nil || !saved {
		return err
	}
	saved, err = IsBoxSaved(environmentID, box)
	if err != nil || !saved {
		// not saved on this environment id or something went wrong
		return err
	}

	content, err := os.ReadFile(savesFile)
	if err != nil {
		return err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)
	for t, env := range targets {
		if env.ID == environmentID {
			for i, savedBox := range env.Boxes {
				if cmp.Equal(box, savedBox) {
					targets[t].Boxes[i] = targets[t].Boxes[len(env.Boxes)-1]
					targets[t].Boxes = targets[t].Boxes[:len(env.Boxes)-1]
					if err != nil {
						return err
					}
					break
				}
			}
		}
	}

	content, err = json.Marshal(targets)
	err = os.WriteFile(savesFile, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

// RemoveEnvironment will remove your environment from tmp folder
func RemoveEnvironment(id string) error {
	err := EnsureSaveFileAvailable()
	if err != nil {
		return err
	}

	isSaved, err := IsEnvironmentSaved(id)
	if err != nil || !isSaved {
		return err
	}

	content, err := os.ReadFile(savesFile)
	if err != nil {
		return err
	}
	targets := []structs.Environment{}
	err = json.Unmarshal(content, &targets)

	for i, env := range targets {
		if env.ID == id {
			targets[i] = targets[len(targets)-1]
			targets = targets[:len(targets)-1]
			if err != nil {
				return err
			}
		}
	}
	content, err = json.Marshal(targets)
	err = os.WriteFile(savesFile, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
