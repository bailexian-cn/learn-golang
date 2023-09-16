package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"testing"
)

var kubeConfigPath string

func TestName(t *testing.T) {
	kubeConfigPath = "C:\\Users\\34552\\Desktop\\debug\\kubeconfig"
	byteArr, err := os.ReadFile(kubeConfigPath)
	if err != nil {
		t.Error(err)
	}
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(byteArr)
	if err != nil {
		t.Error(err)
	}
	c, err := client.New(restConfig, client.Options{})
	itemList := &corev1.ServiceList{}

	if err := c.List(context.Background(), itemList, &client.ListOptions{}); err != nil {
		t.Error(err, "list load balancer service")
	}
	// 删除所有service
	if itemList == nil || len(itemList.Items) < 1 {
		return
	}
	var errList []error
	for _, item := range itemList.Items {
		if item.Spec.Type == corev1.ServiceTypeLoadBalancer {
			e1 := c.Delete(context.Background(), &item)
			if e1 != nil {
				errList = append(errList, e1)
			}
		}
	}
	if len(errList) > 0 {
		t.Error(errList[0], "failed to delete service")
	}
}
