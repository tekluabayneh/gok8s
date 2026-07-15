package mapper

import "github.com/tekluabayneh/gok8s/config"

// TODO
// pod configation  need to be change as is it only contain few of configuration for testing
func ToPod(conf config.Pod) config.Pod {
	pod := config.Pod{
		APIVersion: conf.APIVersion,
		Kind:       conf.Kind,
		Metadata: config.ObjectMeta{
			Name:      conf.Metadata.Name,
			Namespace: conf.Metadata.Namespace,
		},
	}

	return pod
}
