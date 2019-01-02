package e2e

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"time"

	. "github.com/onsi/ginkgo"
	//. "github.com/onsi/gomega"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/test/e2e/framework"
)

var _ = Describe("gcs", func() {
	f := framework.NewDefaultFramework("gcs")

	// filled in BeforeEach
	var c kubernetes.Interface
	var ns string

	var catalogNs string
	var outputDir string
	var templateDir string

	BeforeEach(func() {
		c = f.ClientSet
		ns = f.Namespace.Name

		var err error
		outputDir, err = ioutil.TempDir("", "gcs")
		framework.ExpectNoError(err)
	})

	AfterEach(func() {
		_, _ = framework.RunKubectl("delete", "-f", outputDir)
	})

	Describe("Gluster Container Service", func() {
		It("should allow persistent storage backed by glusterfs", func() {
			nsFlag := fmt.Sprintf("--namespace=%v", ns)

			By("waiting for gluterfs-csi storageclass")
			scName := "glusterfs-csi"
			var sc *storagev1.StorageClass
			framework.ExpectNoError(wait.PollImmediate(2*time.Second, 30*time.Second, func() (bool, error) {
				var err error
				sc, err = c.StorageV1().StorageClasses().Get(scName, metav1.GetOptions{})
				if errors.IsNotFound(err) {
					return false, nil
				}
				if err != nil {
					return false, err
				}
				return true, nil
			}), "waiting for glusterfs-csi storageclass")
			framework.Logf("%v", sc)

			//By("creating a pvc")
			//pvc := &v1.PersistentVolumeClaim{
			//ObjectMeta: metav1.ObjectMeta{
			//GenerateName: "pvc-",
			//},
			//Spec: v1.PersistentVolumeClaimSpec{
			//AccessModes: []v1.PersistentVolumeAccessMode{
			//v1.ReadWriteMany,
			//},
			//Resources: v1.ResourceRequirements{
			//Requests: v1.ResourceList{
			//v1.ResourceName(v1.ResourceStorage): resource.MustParse("1Gi"),
			//},
			//},
			//StorageClassName: &scName,
			//},
			//}
			//var err error
			//pvc, err = c.CoreV1().PersistentVolumeClaims(ns).Create(pvc)
			//framework.ExpectNoError(err)

			//By("waiting for pvc to have an efs pv provisioned for it")
			//framework.ExpectNoError(framework.WaitForPersistentVolumeClaimPhase(v1.ClaimBound, c, ns, pvc.Name, framework.Poll, 1*time.Minute))
			//defer framework.ExpectNoError(framework.DeletePersistentVolumeClaim(c, pvc.Name, ns))
		})
	})
})
