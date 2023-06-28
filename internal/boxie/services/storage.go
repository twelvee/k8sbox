// Package services contains buisness-logic methods of the models
package services

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-cmp/cmp"
	"github.com/twelvee/boxie/pkg/boxie/structs"
	"github.com/twelvee/boxie/pkg/boxie/utils"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applyv1 "k8s.io/client-go/applyconfigurations/core/v1"
	applymetav1 "k8s.io/client-go/applyconfigurations/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// NewStorageService creates a new StorageService
func NewStorageService() structs.StorageService {
	return structs.StorageService{
		EnsureStorageAvailable: ensureStorageAvailable,
		SaveEnvironment:        saveEnvironment,
		DeleteEnvironment:      deleteSavedEnvironment,
		DeleteBox:              deleteSavedBox,
		GetEnvironments:        getSavedEnvironments,
		GetEnvironment:         getSavedEnvironment,
		IsEnvironmentSaved:     isEnvironmentSaved,
	}
}

const CONFIG_MAP_NAME string = "boxie-configmap"

var storageType structs.StorageType

func ensureStorageAvailable(namespace string) error {
	// Make volume is default storage type
	storageType = structs.TYPE_VOLUME
	if os.Getenv("BOXIE_STORAGE_TYPE") == string(structs.TYPE_FILESYSTEM) {
		storageType = structs.TYPE_FILESYSTEM
		err := ensureSaveFileExists()
		return err
	}

	err := ensureVolumeExists(namespace)
	return err
}

func ensureSaveFileExists() error {
	// TODO: Move from utils to internal package
	return utils.EnsureSaveFileAvailable()
}

func ensureVolumeExists(namespace string) error {
	_, err := k8sclient.CoreV1().ConfigMaps(namespace).Get(context.Background(), CONFIG_MAP_NAME, v1.GetOptions{})
	if err != nil {
		err = createEmptyConfigMap(k8sclient, namespace)
		if err != nil {
			return err
		}
	}
	return nil
}

func createEmptyConfigMap(k8sclient *kubernetes.Clientset, namespace string) error {
	configMap := corev1.ConfigMap{
		TypeMeta: v1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      CONFIG_MAP_NAME,
			Namespace: namespace,
		},
	}
	_, err := k8sclient.CoreV1().ConfigMaps(namespace).Create(context.Background(), &configMap, v1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func saveEnvironment(environment structs.Environment) error {
	err := ensureStorageAvailable(environment.Namespace)
	if err != nil {
		return err
	}
	if storageType == structs.TYPE_FILESYSTEM {
		return saveEnvironmentToFilesystem(environment)
	}
	return saveEnvironmentToVolume(environment)
}

func saveEnvironmentToFilesystem(environment structs.Environment) error {
	return utils.SaveEnvironment(environment)
}

func saveEnvironmentToVolume(environment structs.Environment) error {
	configMap, err := k8sclient.CoreV1().ConfigMaps(environment.Namespace).Get(context.Background(), CONFIG_MAP_NAME, v1.GetOptions{})
	if err != nil {
		return err
	}
	savedEnvironments, err := getEnvironmentsFromConfigMap(*configMap)
	if err != nil {
		return err
	}
	savedEnvironments = append(savedEnvironments, environment)
	return updateConfigMap(configMap, savedEnvironments, environment.Namespace)
}

func updateConfigMap(configMap *corev1.ConfigMap, savedEnvironments []structs.Environment, namespace string) error {
	newBinaryData := make(map[string][]byte)
	for _, e := range savedEnvironments {
		nbd, err := json.Marshal(e)
		if err != nil {
			return err
		}
		newBinaryData[e.ID] = nbd
	}

	apiVersion, kind := "v1", "ConfigMap"
	applyConfig := applyv1.ConfigMapApplyConfiguration{
		TypeMetaApplyConfiguration: applymetav1.TypeMetaApplyConfiguration{
			Kind:       &kind,
			APIVersion: &apiVersion,
		},
		ObjectMetaApplyConfiguration: &applymetav1.ObjectMetaApplyConfiguration{
			Name:      &configMap.Name,
			Namespace: &configMap.Namespace,
		},
		BinaryData: newBinaryData,
	}
	_, err := k8sclient.CoreV1().ConfigMaps(namespace).Apply(context.Background(), &applyConfig, v1.ApplyOptions{FieldManager: "boxie"})
	if err != nil {
		return err
	}
	return nil
}

func getEnvironmentsFromConfigMap(configMap corev1.ConfigMap) ([]structs.Environment, error) {
	var savedEnvironments []structs.Environment
	for _, e := range configMap.BinaryData {
		var env structs.Environment
		err := json.Unmarshal(e, &env)
		if err != nil {
			return nil, err
		}
		savedEnvironments = append(savedEnvironments, env)
	}
	return savedEnvironments, nil
}

func isEnvironmentSaved(environment structs.Environment) (bool, error) {
	err := ensureStorageAvailable(environment.Namespace)
	if err != nil {
		return false, err
	}
	if storageType == structs.TYPE_FILESYSTEM {
		return isEnvironmentSavedToFilesystem(environment)
	}
	return isEnvironmentSavedToVolume(environment)
}

func isEnvironmentSavedToFilesystem(environment structs.Environment) (bool, error) {
	return utils.IsEnvironmentSaved(environment.ID)
}

func isEnvironmentSavedToVolume(environment structs.Environment) (bool, error) {
	configMap, err := k8sclient.CoreV1().ConfigMaps(environment.Namespace).Get(context.Background(), CONFIG_MAP_NAME, v1.GetOptions{})
	if err != nil {
		return false, err
	}
	savedEnvironments, err := getEnvironmentsFromConfigMap(*configMap)
	if err != nil {
		return false, err
	}
	currentEnvironment := -1
	for i, env := range savedEnvironments {
		if env.ID == environment.ID {
			currentEnvironment = i
			break
		}
	}

	if currentEnvironment == -1 {
		return false, nil
	}
	return true, nil
}

func deleteSavedEnvironment(environment structs.Environment) error {
	err := ensureStorageAvailable(environment.Namespace)
	if err != nil {
		return err
	}
	if storageType == structs.TYPE_FILESYSTEM {
		return deleteEnvironmentFromFilesystem(environment)
	}
	return deleteEnvironmentFromVolume(environment)
}

func deleteEnvironmentFromFilesystem(environment structs.Environment) error {
	return utils.RemoveEnvironment(environment.ID)
}

func deleteEnvironmentFromVolume(environment structs.Environment) error {
	configMap, err := k8sclient.CoreV1().ConfigMaps(environment.Namespace).Get(context.Background(), CONFIG_MAP_NAME, v1.GetOptions{})
	if err != nil {
		return err
	}
	savedEnvironments, err := getEnvironmentsFromConfigMap(*configMap)
	if err != nil {
		return err
	}
	var newSavedEnvironments []structs.Environment
	for _, env := range savedEnvironments {
		if env.ID != environment.ID {
			newSavedEnvironments = append(newSavedEnvironments, env)
		}
	}

	return updateConfigMap(configMap, newSavedEnvironments, environment.Namespace)
}

func deleteSavedBox(environment structs.Environment, box structs.Box) error {
	err := ensureStorageAvailable(environment.Namespace)
	if err != nil {
		return err
	}
	if storageType == structs.TYPE_FILESYSTEM {
		return deleteBoxFromFilesystem(environment, box)
	}
	return deleteBoxFromVolume(environment, box)
}

func deleteBoxFromFilesystem(environment structs.Environment, box structs.Box) error {
	return utils.RemoveBox(box, environment.ID)
}

func deleteBoxFromVolume(environment structs.Environment, box structs.Box) error {
	configMap, err := k8sclient.CoreV1().ConfigMaps(environment.Namespace).Get(context.Background(), CONFIG_MAP_NAME, v1.GetOptions{})
	if err != nil {
		return err
	}
	savedEnvironments, err := getEnvironmentsFromConfigMap(*configMap)
	if err != nil {
		return err
	}
	currentEnvironment := -1
	for i, env := range savedEnvironments {
		if env.ID == environment.ID {
			currentEnvironment = i
			break
		}
	}

	if currentEnvironment != -1 {
		for j, b := range savedEnvironments[currentEnvironment].Boxes {
			if cmp.Equal(b, box) {
				savedEnvironments[currentEnvironment].Boxes[j] = savedEnvironments[currentEnvironment].Boxes[len(savedEnvironments[currentEnvironment].Boxes)-1]
				savedEnvironments[currentEnvironment].Boxes = savedEnvironments[currentEnvironment].Boxes[:len(savedEnvironments[currentEnvironment].Boxes)-1]
				break
			}
		}
	}

	return updateConfigMap(configMap, savedEnvironments, environment.Namespace)
}

func getSavedEnvironments(namespace string) ([]structs.Environment, error) {
	err := ensureStorageAvailable(namespace)
	if err != nil {
		return nil, err
	}
	if storageType == structs.TYPE_FILESYSTEM {
		return getEnvironmentsFromFilesystem(namespace)
	}
	return getEnvironmentsFromVolume(namespace)
}

func getEnvironmentsFromFilesystem(namespace string) ([]structs.Environment, error) {
	return utils.GetEnvironments()
}

func getEnvironmentsFromVolume(namespace string) ([]structs.Environment, error) {
	configMap, err := k8sclient.CoreV1().ConfigMaps(namespace).Get(context.Background(), CONFIG_MAP_NAME, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return getEnvironmentsFromConfigMap(*configMap)
}

func getSavedEnvironment(namespace string, id string) (*structs.Environment, error) {
	err := ensureStorageAvailable(namespace)
	if err != nil {
		return nil, err
	}
	if storageType == structs.TYPE_FILESYSTEM {
		return getEnvironmentFromFilesystem(namespace, id)
	}
	return getEnvironmentFromVolume(namespace, id)
}

func getEnvironmentFromFilesystem(namespace string, id string) (*structs.Environment, error) {
	return utils.GetEnvironment(id)
}

func getEnvironmentFromVolume(namespace string, id string) (*structs.Environment, error) {
	configMap, err := k8sclient.CoreV1().ConfigMaps(namespace).Get(context.Background(), CONFIG_MAP_NAME, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	for k, e := range configMap.BinaryData {
		if k == id {
			var env structs.Environment
			err := json.Unmarshal(e, &env)
			if err != nil {
				return nil, err
			}
			return &env, nil
		}
	}
	return nil, fmt.Errorf("No environment found.")
}
