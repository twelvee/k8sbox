// Package create is an entry point and help tools for create commands
package create

func getExample() string {
	return `
	boxie create box -toml|--toml-file=./box.toml

	boxie create box -json|--json-spec="{json_spec}"

	boxie create box -t|--type={helm} -n|--name="my-box" -c|--chart="helm_chart_yaml_as_string" -v|--values="helm_values_yaml_as_string" -ns|--namespace="my-namespace" 
	`
}
