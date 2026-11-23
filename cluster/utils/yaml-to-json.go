package utils

import "sigs.k8s.io/yaml"

func YamlToJson(file string) []byte {
	jsofile, err := yaml.YAMLToJSON([]byte(file))
	if err != nil {
		panic(err)
	}
	return jsofile
}

func JsonToYaml(file string) []byte {
	jsofile, err := yaml.JSONToYAML([]byte(file))
	if err != nil {
		panic(err)
	}
	return jsofile
}
