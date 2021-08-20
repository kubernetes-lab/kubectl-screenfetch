package k8s

import (
	"context"
	"flag"
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type KubernetesInfo struct {
	CurrentContext   string
	Version          string
	NodesCount       uint64
	NamespacesCount  uint64
	DeploymentsCount uint64
	PodsCount        uint64
	ServicesCount    uint64
	IngressesCount   uint64
	PVsCount         uint64
	CRI              string
	CNI              string
	CSI              string
}

func GetKubernetesInfo(ctx context.Context) (k KubernetesInfo) {
	var kubeconfigPath *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfigPath = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfigPath = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	kubeconfig, err := clientcmd.LoadFromFile(*kubeconfigPath)
	if err != nil {
		panic(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfigPath)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	k.CurrentContext = kubeconfig.CurrentContext

	if nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{}); err == nil {
		k.NodesCount = uint64(len(nodes.Items))
		for _, node := range nodes.Items {
			k.Version = node.Status.NodeInfo.KubeletVersion
			k.CRI = node.Status.NodeInfo.ContainerRuntimeVersion
			break
		}
	}

	if namespaces, err := clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{}); err == nil {
		k.NamespacesCount = uint64(len(namespaces.Items))
	}

	if deployments, err := clientset.AppsV1().Deployments(metav1.NamespaceAll).List(ctx, metav1.ListOptions{}); err == nil {
		k.DeploymentsCount = uint64(len(deployments.Items))
	}

	if pods, err := clientset.CoreV1().Pods(metav1.NamespaceAll).List(ctx, metav1.ListOptions{}); err == nil {
		k.PodsCount = uint64(len(pods.Items))
	}

	if services, err := clientset.CoreV1().Services(metav1.NamespaceAll).List(ctx, metav1.ListOptions{}); err == nil {
		k.ServicesCount = uint64(len(services.Items))
	}

	if ingresses, err := clientset.NetworkingV1().Ingresses(metav1.NamespaceAll).List(ctx, metav1.ListOptions{}); err == nil {
		k.IngressesCount = uint64(len(ingresses.Items))
	}

	if pvs, err := clientset.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{}); err == nil {
		k.PVsCount = uint64(len(pvs.Items))
	}

	if csis, err := clientset.StorageV1().CSIDrivers().List(ctx, metav1.ListOptions{}); err == nil {
		for _, csi := range csis.Items {
			k.CSI += fmt.Sprintf("%s,", csi.Name)
		}
	}

	filepath.WalkDir("/etc/cni/net.d/", func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() {
			hyphenIndex := strings.Index(d.Name(), "-")
			dotIndex := strings.LastIndex(d.Name(), ".")
			k.CNI = d.Name()[hyphenIndex+1:dotIndex]
		}
		return err
	})

	return
}
