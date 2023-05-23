package utils

import (
	"encoding/json"
	"os"

	"github.com/google/go-cmp/cmp"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
)

const saveDir = "/tmp/k8sbox_saves"
const savesFile = "/tmp/k8sbox_saves/save"

func ensureSaveFileAvailable() error {
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

func IsBoxSaved(environmentId string, sbox structs.Box) (bool, error) {
	err := ensureSaveFileAvailable()
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
		if env.Id == environmentId {
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

func IsEnvironmentSaved(id string) (bool, error) {
	err := ensureSaveFileAvailable()
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
		if env.Id == id {
			return true, nil
		}
	}
	return false, nil
}

func SaveEnvironment(environment structs.Environment) error {
	err := ensureSaveFileAvailable()
	if err != nil {
		return err
	}

	isSavedAlready, err := IsEnvironmentSaved(environment.Id)
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

func GetEnvironment(id string) (*structs.Environment, error) {
	err := ensureSaveFileAvailable()
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
		if env.Id == id {
			return &env, nil
		}
	}
	return nil, nil
}

func GetBox(environmentId string, boxName string, boxNamespace string, boxTempDirectory string) (*structs.Box, error) {
	err := ensureSaveFileAvailable()
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
		if env.Id == environmentId {
			for _, box := range env.Boxes {
				if box.Name == boxName && box.Namespace == boxNamespace && box.TempDirectory == boxTempDirectory {
					return &box, nil
				}
			}
			return nil, nil
		}
	}
	return nil, nil
}

func SaveBox(box structs.Box, environmentId string) error {
	err := ensureSaveFileAvailable()
	if err != nil {
		return err
	}
	saved, err := IsEnvironmentSaved(environmentId)
	if err != nil || !saved {
		return err
	}
	saved, err = IsBoxSaved(environmentId, box)
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
		if env.Id == environmentId {
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

func RemoveBox(box structs.Box, environmentId string) error {
	err := ensureSaveFileAvailable()
	if err != nil {
		return err
	}
	saved, err := IsEnvironmentSaved(environmentId)
	if err != nil || !saved {
		return err
	}
	saved, err = IsBoxSaved(environmentId, box)
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
		if env.Id == environmentId {
			for i, savedBox := range env.Boxes {
				if cmp.Equal(box, savedBox) {
					targets[t].Boxes[i] = targets[t].Boxes[len(env.Boxes)-1]
					targets[t].Boxes = targets[t].Boxes[:len(env.Boxes)-1]
					err := os.RemoveAll(box.TempDirectory)
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

func RemoveEnvironment(id string) error {
	err := ensureSaveFileAvailable()
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
		if env.Id == id {
			targets[i] = targets[len(targets)-1]
			targets = targets[:len(targets)-1]
		}
	}
	content, err = json.Marshal(targets)
	err = os.WriteFile(savesFile, content, 0644)
	if err != nil {
		return err
	}

	return nil
}
