package main

import (
	"flag"
	"log"

	"k8s.io/client-go/tools/clientcmd"

	kubeflowbackend "github.com/caicloud/ciao/pkg/backend/kubeflow"
	interpreter "github.com/caicloud/ciao/pkg/interpreter/simple"
	"github.com/caicloud/ciao/pkg/kernel"
	"github.com/caicloud/ciao/pkg/manager"
	s2i "github.com/caicloud/ciao/pkg/s2i/simple"
	"github.com/caicloud/ciao/version"
)

const RecommendedKubeConfigPathEnv = "KUBECONFIG"

func main() {

	// Parse the connection file.
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln("Need a command line argument specifying the connection file.")
	}

	// This is a trick!
	kubeconfig := "/var/run/kubernetes/admin.kubeconfig"

	// // Note: ENV KUBECONFIG will overwrite user defined Kubeconfig option.
	// if len(os.Getenv(RecommendedKubeConfigPathEnv)) > 0 {
	// 	// use the current context in kubeconfig
	// 	// This is very useful for running locally.
	// 	kubeconfig = os.Getenv(RecommendedKubeConfigPathEnv)
	// }

	// Get kubernetes config.
	kcfg, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeflowBackend, err := kubeflowbackend.New(kcfg)
	if err != nil {
		log.Fatalf("Error building kubeflow backend: %s", err.Error())
	}

	// TODO: Using a real s2i client.
	s2iClient := s2i.New()

	interpreter := interpreter.New()

	manager := manager.New(kubeflowBackend, s2iClient, interpreter)

	kernel := kernel.New(version.ProtocolVersion, version.Version, flag.Arg(0), manager)

	log.Println("Running Kubeflow kernel for Jupyter...")
	kernel.RunKernel()
}
