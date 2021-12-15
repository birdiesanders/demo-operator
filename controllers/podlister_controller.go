/*
Copyright 2021.

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

	//"k8s.io/apimachinery/pkg/labels"
	//types "k8s.io/apimachinery/pkg/types"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"

	//"k8s.io/apimachinery/pkg/api/errors"
	"github.com/go-logr/logr"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	podlisterv1alpha1 "github.com/birdiesanders/demo-operator/api/v1alpha1"
	//"github.com/onsi/ginkgo/types"
)

// PodListerReconciler reconciles a PodLister object
type PodListerReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=pod-lister.example.com,resources=podlisters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=pod-lister.example.com,resources=podlisters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=pod-lister.example.com,resources=podlisters/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch;
//+kubebuilder:rbac:groups=core,resources=pod,verbs=get;list;watch;
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodLister object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *PodListerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	r.Log.Info("Starting Reconcile:", "namespace", req.NamespacedName.Namespace, "name", req.Name)

	//r.Log.Info("Looking for pods with label...", "context", ctx)
	podList := v1.PodList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
	}

	listOptions := []client.ListOption{
		client.InNamespace(req.NamespacedName.Namespace),
		client.MatchingLabels{"hello": "civo"},
	}
	err := r.List(ctx, &podList, listOptions...)

	if err != nil {
		return ctrl.Result{}, err
	}
	for _, pod := range podList.Items {

		r.Log.Info("Pod Found!", "pod", pod)
		deployment := &appsv1.Deployment{}
		r.Get(ctx, types.NamespacedName{Namespace: pod.Namespace, Name: "ngin"}, deployment)
		r.Log.Info("Found deployment:", "deployment", deployment.Name)
		r.Log.Info("--------------------------------------------------------------------------------------")
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodListerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&podlisterv1alpha1.PodLister{}).
		Complete(r)
}
