package main

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	chartPath := "/tmp/my-chart-0.1.0.tgz"
	chart, err := loader.Load(chartPath)
	if err != nil {
		panic(err)
	}

	kubeconfigPath := "/tmp/my-kubeconfig"
	releaseName := "my-release"
	releaseNamespace := "default"
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(kubeconfigPath, "", releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		panic(err)
	}

	iCli := action.NewInstall(actionConfig)
	iCli.Namespace = releaseNamespace
	iCli.ReleaseName = releaseName
	rel, err := iCli.Run(chart, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully installed release: ", rel.Name)
}
