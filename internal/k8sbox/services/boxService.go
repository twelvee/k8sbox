// Package services contains buisness-logic methods of the models
package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/twelvee/k8sbox/pkg/k8sbox/structs"
	"github.com/twelvee/k8sbox/pkg/k8sbox/utils"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/engine"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes"
)

// NewBoxService creates a new BoxService
func NewBoxService() structs.BoxService {
	return structs.BoxService{
		ProcessEnvValues:        processEnvValues,
		ValidateBoxes:           validateBoxes,
		FillEmptyFields:         fillEmptyFields,
		UninstallBox:            uninstallBox,
		DescribeBoxApplications: describeBoxApplications,
		ExpandBoxVariables:      expandBoxVariables,
	}
}

func fillEmptyFields(environment structs.Environment, box *structs.Box) error {
	if len(strings.TrimSpace(box.Namespace)) == 0 {
		if len(strings.TrimSpace(environment.Namespace)) == 0 {
			box.Namespace = strings.ToLower(strings.Join([]string{"k8srun", utils.GetShortNamespace(8)}, "-"))
		} else {
			box.Namespace = environment.Namespace
		}
	}
	if len(strings.TrimSpace(box.Name)) == 0 {
		box.Name = strings.ToLower(strings.Join([]string{"k8srun", utils.GetShortNamespace(8)}, "-"))
	}

	if box.Type == structs.Helm() {
		err := createHelmRenders(environment, box)
		if err != nil {
			return err
		}
	}

	return nil
}

func createHelmRenders(environment structs.Environment, box *structs.Box) error {
	chart, err := loader.Load(filepath.Dir(box.Chart))
	if err != nil {
		return err
	}
	releaseOptions := chartutil.ReleaseOptions{
		Name:      box.Name,
		Namespace: box.Namespace,
		Revision:  1,
		IsInstall: true,
	}

	e := engine.New(restConfig)
	var replacedValues map[string]interface{}
	if len(strings.TrimSpace(environment.Variables)) != 0 {
		replacedValues = processEnvValues(chart.Values, environment.Variables)
	}
	vals, err := chartutil.ToRenderValues(chart, replacedValues, releaseOptions, nil)
	if err != nil {
		return err
	}
	render, err := e.Render(chart, vals)
	if err != nil {
		return err
	}
	box.HelmRender = render
	return nil
}

func processEnvValues(values map[string]interface{}, dotenvPath string) map[string]interface{} {
	if len(dotenvPath) > 0 {
		err := godotenv.Load(dotenvPath)
		if err != nil {
			panic(err)
		}
	}
	for k, v := range values {
		if len(os.ExpandEnv(fmt.Sprintf("%v", v))) == 0 {
			continue
		}
		values[k] = os.ExpandEnv(fmt.Sprintf("%v", v))
	}
	return values
}

func validateBoxes(boxes []structs.Box) error {
	var messages []string
	for index, box := range boxes {
		if len(strings.TrimSpace(box.Type)) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Type is missing", index))
		}

		if len(box.Applications) == 0 && box.Type != structs.Helm() {
			messages = append(messages, fmt.Sprintf("-> Box %d: Applications are missing", index))
		}

		if len(strings.TrimSpace(box.Chart)) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Chart is missing", index))
		}
		_, err := os.Stat(box.Chart)
		if err != nil {
			messages = append(messages, fmt.Sprintf("-> Box %d: Chart file can't be opened (%s)", index, box.Chart))
		}

		if len(strings.TrimSpace(box.Values)) == 0 {
			messages = append(messages, fmt.Sprintf("-> Box %d: Values are missing", index))
		}
		_, err = os.Stat(box.Values)
		if err != nil {
			messages = append(messages, fmt.Sprintf("-> Box %d: Values file can't be opened (%s)", index, box.Values))
		}

		if box.Type != structs.Helm() {
			applicationsErrors := validateApplications(box.Applications)
			if len(applicationsErrors) > 0 {
				for _, err := range applicationsErrors {
					messages = append(messages, fmt.Sprintf("-> Box %d: \n\r%s", index, err))
				}
			}
		}
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n\r"))
	}
	return nil
}

func installBox(box *structs.Box, environment structs.Environment) ([]*runtime.Object, error) {
	var objects []*runtime.Object

	r := utils.ConvertHelmRenderToYaml(box.HelmRender)
	for _, rend := range r {
		obj, err := utils.CreateRuntimeObject(rend)
		if err != nil {
			return nil, err
		}

		mapping, err := utils.CreateRestMapper(k8sclient, obj)
		if err != nil {
			return nil, err
		}

		restClient, err := utils.NewRestClient(*restConfig, mapping.GroupVersionKind.GroupVersion())
		if err != nil {
			return nil, err
		}

		// Use the REST helper to create the object in the box namespace.
		restHelper := resource.NewHelper(restClient, mapping)
		rtobj, err := restHelper.Create(box.Namespace, false, obj)
		if err != nil {
			return nil, err
		}
		objects = append(objects, &rtobj)
	}

	return objects, nil
}

func uninstallBox(environment structs.Environment, box structs.Box) ([]*runtime.Object, error) {
	var objects []*runtime.Object

	r := utils.ConvertHelmRenderToYaml(box.HelmRender)
	for _, rend := range r {
		obj, err := utils.CreateRuntimeObject(rend)
		if err != nil {
			return nil, err
		}

		mapping, err := utils.CreateRestMapper(k8sclient, obj)
		if err != nil {
			return nil, err
		}

		restClient, err := utils.NewRestClient(*restConfig, mapping.GroupVersionKind.GroupVersion())
		if err != nil {
			return nil, err
		}

		// Use the REST helper to create the object in the box namespace.
		restHelper := resource.NewHelper(restClient, mapping)

		name, err := meta.NewAccessor().Name(obj)
		if err != nil {
			return nil, err
		}

		rtobj, err := restHelper.Delete(box.Namespace, name)
		if err != nil {
			return nil, err
		}
		objects = append(objects, &rtobj)
	}

	err := deleteSavedBox(environment, box)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func describeBoxApplications(environment structs.Environment, box structs.Box) error {
	r := utils.ConvertHelmRenderToYaml(box.HelmRender)
	for _, rend := range r {
		obj, err := utils.CreateRuntimeObject(rend)
		if err != nil {
			return err
		}
		var describeFunc func(*kubernetes.Clientset, string, string) error
		switch kind := obj.GetObjectKind().GroupVersionKind().Kind; kind {
		case structs.KIND_POD:
			describeFunc = describePod
		case structs.KIND_POD_TEMPLATE:
			describeFunc = describePodTemplate
		case structs.KIND_REPLICATION_CONTROLLER:
			describeFunc = describeReplicationController
		case structs.KIND_REPLICA_SET:
			describeFunc = describeReplicaSet
		case structs.KIND_DEPLOYMENT:
			describeFunc = describeDeployment
		case structs.KIND_STATEFUL_SET:
			describeFunc = describeStatefulSet
		case structs.KIND_CONTROLLER_REVISION:
			describeFunc = describeControllerRevision
		case structs.KIND_DAEMON_SET:
			describeFunc = describeDaemonSet
		case structs.KIND_JOB:
			describeFunc = describeJob
		case structs.KIND_CRONJOB:
			describeFunc = describeCronjob
		case structs.KIND_HPA:
			describeFunc = describeHPA
		case structs.KIND_SERVICE:
			describeFunc = describeService
		case structs.KIND_INGRESS:
			describeFunc = describeIngress
		}

		name, err := meta.NewAccessor().Name(obj)
		if err != nil {
			return err
		}
		if describeFunc != nil {
			err = describeFunc(k8sclient, box.Namespace, name)
			if err != nil {
				return err
			}
			fmt.Println()
		}
	}
	return nil
}

func expandBoxVariables(boxes []structs.Box) []structs.Box {
	var newBoxes []structs.Box

	for _, b := range boxes {
		b.Name = os.ExpandEnv(b.Name)
		b.Namespace = os.ExpandEnv(b.Namespace)
		b.Type = os.ExpandEnv(b.Type)
		b.Chart = os.ExpandEnv(b.Chart)
		b.Values = os.ExpandEnv(b.Values)
		b.Applications = ExpandApplications(b.Applications)
		newBoxes = append(newBoxes, b)
	}
	return newBoxes
}
