// Package put is an entry point and help tools for put commands
package put

func getExample() string {
	return `
	boxie put box -toml|--toml-file=./box.toml --force // --force to update values with same box name

	boxie put box -json|--json-spec="{json_spec}"

	boxie put box -t|--type={helm} -n|--name="my-box" -c|--chart="helm_chart_yaml_as_string" -v|--values="helm_values_yaml_as_string" --namespace="my-namespace" 
	`
}
