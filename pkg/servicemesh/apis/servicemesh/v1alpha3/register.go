package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Define your schema name and the version
var SchemeGroupVersion = schema.GroupVersion{
	Group:   "istio.openshift.com",
	Version: "v1alpha3",
}

var InternalSchemeGroupVersion = schema.GroupVersion{
	Group:   "istio.openshift.com",
	Version: runtime.APIVersionInternal,
}

var (
	SchemeBuilder      runtime.SchemeBuilder
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func init() {
	// We only register manually written functions here. The registration of the
	// generated functions takes place in the generated files. The separation
	// makes the code compile even when the generated files are missing.
	localSchemeBuilder.Register(addKnownTypes)
	localSchemeBuilder.Register(addKnownInternalTypes)
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&ServiceMeshMemberRoll{},
		&ServiceMeshMemberRollList{},
	)

	metav1.AddToGroupVersion(
		scheme,
		SchemeGroupVersion,
	)

	return nil
}

func addKnownInternalTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		InternalSchemeGroupVersion,
		&ServiceMeshMemberRoll{},
		&ServiceMeshMemberRollList{},
	)

	return nil
}
