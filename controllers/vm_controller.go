/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"github.com/go-logr/logr"
	//"time"

	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mscloudv1 "crd/api/v1"
)

// VMReconciler reconciles a VM object
type VMReconciler struct {
	client.Client
	Log logr.Logger
}

// +kubebuilder:rbac:groups=mscloud.xiliangam.com,resources=vms,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=mscloud.xiliangam.com,resources=vms/status,verbs=get;update;patch

func (r *VMReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("vm", req.NamespacedName)

	// 1. xiliangma 获取vm 信息
	var vm mscloudv1.VM
	if err := r.Get(ctx, req.NamespacedName, &vm); err != nil {
		log.Error(err, "unable to get vm")
	} else {
		fmt.Println("Get vm spec info success, ", vm.Spec.Name, vm.Spec.Type, vm.Spec.CPU, vm.Spec.Memory, vm.Spec.HA)

	}

	// 2. 更新虚拟机状态
	//vm.Status.UpdateLastTime = metav1.Now()
	//vm.Status.Status = "Running"
	//if err := r.Status().Update(ctx, &vm); err != nil {
	//	log.Error(err, "not update vm  status.")
	//}

	// 3. 删除虚拟机
	//time.Sleep(time.Second * 5)
	//if err := r.Delete(ctx, &vm); err != nil {
	//	log.Error(err, "unable to delete vm ", "vm", vm)
	//}

	return ctrl.Result{}, nil
}

func (r *VMReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mscloudv1.VM{}).
		Complete(r)
}
