// Package services contains buisness-logic methods of the models
package services

import (
	"errors"
	"fmt"
	"os"
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
		ProcessEnvValues:   processEnvValues,
		ValidateBoxes:      validateBoxes,
		FillEmptyFields:    fillEmptyFields,
		UninstallBox:       UninstallBox,
		GetBox:             GetBox,
		ExpandBoxVariables: expandBoxVariables,
	}
}

func fillEmptyFields(box *structs.Box, environmentNamespace string) error {
	if len(strings.TrimSpace(box.Namespace)) == 0 {
		if len(strings.TrimSpace(environmentNamespace)) == 0 {
			box.Namespace = strings.ToLower(strings.Join([]string{"k8srun", utils.GetShortNamespace(8)}, "-"))
		} else {
			box.Namespace = environmentNamespace
		}
	}
	if len(strings.TrimSpace(box.Name)) == 0 {
		box.Name = strings.ToLower(strings.Join([]string{"k8srun", utils.GetShortNamespace(8)}, "-"))
	}

	return nil
}

func processEnvValues(values map[string]interface{}, dotenvPath string) map[string]interface{} {
	if len(dotenvPath) > 0 {
		godotenv.Load(dotenvPath)
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

		if len(box.Applications) == 0 {
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

		applicationsErrors := validateApplications(box.Applications)
		if len(applicationsErrors) > 0 {
			for _, err := range applicationsErrors {
				messages = append(messages, fmt.Sprintf("-> Box %d: \n\r%s", index, err))
			}
		}
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, "\n\r"))
	}
	return nil
}

// InstallBox will deploy your box applications into your k8s cluster
func InstallBox(box *structs.Box, environment structs.Environment) ([]*runtime.Object, error) {
	var objects []*runtime.Object
	_, restConfig := GetConfigFromKubeconfig(box.Namespace)
	chart, err := loader.Load(box.TempDirectory)
	if err != nil {
		return nil, err
	}

	releaseOptions := chartutil.ReleaseOptions{
		Name:      box.Name,
		Namespace: box.Namespace,
		Revision:  1,
		IsInstall: true,
	}

	e := engine.New(restConfig)
	vals, err := chartutil.ToRenderValues(chart, processEnvValues(chart.Values, environment.Variables), releaseOptions, nil)
	if err != nil {
		return nil, err
	}
	render, err := e.Render(chart, vals)
	if err != nil {
		return nil, err
	}

	k8sclient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	r := utils.ConvertHelmRenderToYaml(render)
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

	utils.SaveBox(*box, environment.ID)
	return objects, nil
}

// UninstallBox will uninstall your box applications from your k8s cluster
func UninstallBox(box *structs.Box, environment structs.Environment) ([]*runtime.Object, error) {
	var objects []*runtime.Object
	_, restConfig := GetConfigFromKubeconfig(box.Namespace)

	chart, err := loader.Load(box.TempDirectory)
	if err != nil {
		return nil, err
	}

	releaseOptions := chartutil.ReleaseOptions{
		Name:      box.Name,
		Namespace: box.Namespace,
		Revision:  1,
		IsInstall: true,
	}

	e := engine.New(restConfig)
	vals, err := chartutil.ToRenderValues(chart, processEnvValues(chart.Values, environment.Variables), releaseOptions, nil)
	if err != nil {
		return nil, err
	}
	render, err := e.Render(chart, vals)
	if err != nil {
		return nil, err
	}

	k8sclient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	r := utils.ConvertHelmRenderToYaml(render)
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

	err = utils.RemoveBox(*box, environment.ID)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

// GetBox will return a helm release from your k8s cluster
func GetBox(box *structs.Box) ([]*runtime.Object, error) {
	_, restConfig := GetConfigFromKubeconfig(box.Namespace)
	var objects []*runtime.Object
	chart, err := loader.Load(box.TempDirectory)
	if err != nil {
		return nil, err
	}

	releaseOptions := chartutil.ReleaseOptions{
		Name:      box.Name,
		Namespace: box.Namespace,
		Revision:  1,
		IsInstall: true,
	}

	e := engine.New(restConfig)
	vals, err := chartutil.ToRenderValues(chart, chart.Values, releaseOptions, nil)
	if err != nil {
		return nil, err
	}
	render, err := e.Render(chart, vals)
	if err != nil {
		return nil, err
	}

	k8sclient, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}
	r := utils.ConvertHelmRenderToYaml(render)
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

		rtobj, err := restHelper.Get(box.Namespace, name)
		if err != nil {
			return nil, err
		}
		objects = append(objects, &rtobj)
	}
	return objects, err
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
