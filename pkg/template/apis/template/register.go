package template

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kapi "k8s.io/kubernetes/pkg/api"
)

const (
	GroupName       = "template.openshift.io"
	LegacyGroupName = ""
)

// SchemeGroupVersion is group version used to register these objects
var (
	SchemeGroupVersion       = schema.GroupVersion{Group: GroupName, Version: runtime.APIVersionInternal}
	LegacySchemeGroupVersion = schema.GroupVersion{Group: LegacyGroupName, Version: runtime.APIVersionInternal}

	LegacySchemeBuilder    = runtime.NewSchemeBuilder(addLegacyKnownTypes, RegisterDeepCopies)
	AddToSchemeInCoreGroup = LegacySchemeBuilder.AddToScheme

	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// LegacyKind takes an unqualified kind and returns back a Group qualified GroupKind
func LegacyKind(kind string) schema.GroupKind {
	return LegacySchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns back a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func LegacyResource(resource string) schema.GroupResource {
	return LegacySchemeGroupVersion.WithResource(resource).GroupResource()
}

// IsKindOrLegacy checks if the provided GroupKind matches with the given kind by looking
// up the API group and also the legacy API.
func IsKindOrLegacy(kind string, gk schema.GroupKind) bool {
	return gk == Kind(kind) || gk == LegacyKind(kind)
}

// IsResourceOrLegacy checks if the provided GroupResources matches with the given
// resource by looking up the API group and also the legacy API.
func IsResourceOrLegacy(resource string, gr schema.GroupResource) bool {
	return gr == Resource(resource) || gr == LegacyResource(resource)
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Template{},
		&TemplateList{},
		&TemplateInstance{},
		&TemplateInstanceList{},
		&BrokerTemplateInstance{},
		&BrokerTemplateInstanceList{},
		&kapi.List{},
	)
	return nil
}

func addLegacyKnownTypes(scheme *runtime.Scheme) error {
	types := []runtime.Object{
		&Template{},
		&TemplateList{},
	}
	scheme.AddKnownTypes(LegacySchemeGroupVersion, types...)
	scheme.AddKnownTypeWithName(LegacySchemeGroupVersion.WithKind("TemplateConfig"), &Template{})
	scheme.AddKnownTypeWithName(LegacySchemeGroupVersion.WithKind("ProcessedTemplate"), &Template{})
	return nil
}
