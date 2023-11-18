/*
Copyright 2023.

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

package venus

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/retry"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	venusv1 "github.com/huweihuang/venus/api/venus/v1"
)

// MemcacheReconciler reconciles a Memcache object
type MemcacheReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=venus.huweihuang.com,resources=memcaches,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=venus.huweihuang.com,resources=memcaches/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=venus.huweihuang.com,resources=memcaches/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Memcache object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *MemcacheReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MemcacheReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&venusv1.Memcache{}).
		Complete(r)
}

// UpdateStatus updates Status with retry.RetryOnConflict
// fix error: the object has been modified; please apply your changes to the latest version and try again
func (r *MemcacheReconciler) UpdateStatus(ctx context.Context, req ctrl.Request, obj *venusv1.Memcache) error {
	status := obj.DeepCopy().Status
	return retry.RetryOnConflict(retry.DefaultBackoff, func() (err error) {
		if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
			return err
		}
		obj.Status = status
		return r.Status().Update(ctx, obj)
	})
}
