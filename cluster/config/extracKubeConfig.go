package config

import (
	"encoding/json"

	Types "github.com/tekluabayneh/gok8s/types"
	"github.com/tekluabayneh/gok8s/utils"
)

type Kubeconfig struct {
	APIVersion     string         `json:"apiVersion"`
	Kind           string         `json:"kind"`
	CurrentContext string         `json:"current-context"`
	Users          []NamedUser    `json:"users"`
	Clusters       []NamedCluster `json:"clusters"`
	Contexts       []NamedContext `json:"contexts"`
}

type NamedUser struct {
	Name string `json:"name"`
	User struct {
		Token                 string `json:"token"`
		ClientCertificateData string `json:"client-certificate-data"`
		ClientKeyData         string `json:"client-key-data"`
	} `json:"user"`
}

type NamedCluster struct {
	Name    string `json:"name"`
	Cluster struct {
		Server                   string `json:"server"`
		CertificateAuthorityData string `json:"certificate-authority-data"`
		InsecureSkipTLSVerify    bool   `json:"insecure-skip-tls-verify"`
	} `json:"cluster"`
}

type NamedContext struct {
	Name    string `json:"name"`
	Context struct {
		Cluster string `json:"cluster"`
		User    string `json:"user"`
	} `json:"context"`
}

func ExtractKubeConfig() Types.ApiKubeconfig {
	file := ReadFile(".kube/config")
	jsofile := utils.YamlToJson(file)

	da := Types.ApiKubeconfig{}
	var data Kubeconfig
	if err := json.Unmarshal(jsofile, &data); err != nil {
		panic(err)
	}
	da.CurrentContext = data.CurrentContext

	for _, tok := range data.Users {
		da.ClientKeyData = tok.User.ClientKeyData
		da.ClientCertificateData = tok.User.ClientCertificateData
		if tok.User.Token != "" {
			da.Token = tok.User.Token
		}

	}

	for _, clus := range data.Clusters {
		if clus.Name == data.CurrentContext {
			da.BaseURL = clus.Cluster.Server
			da.CertificateAuthorityData = clus.Cluster.CertificateAuthorityData
			da.InsecureSkipTLSVerify = clus.Cluster.InsecureSkipTLSVerify

		}
	}
	for _, clus := range data.Contexts {
		if clus.Name == data.CurrentContext {
			da.ContextUser = clus.Context.User
			da.ContextCluster = clus.Context.Cluster
		}
	}

	return da
}
