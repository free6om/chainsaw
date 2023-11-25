package validation

import (
	"github.com/kyverno/chainsaw/pkg/apis/v1alpha1"
	"k8s.io/apimachinery/pkg/util/validation/field"
)

func ValidateDelete(path *field.Path, obj *v1alpha1.Delete) field.ErrorList {
	var errs field.ErrorList
	if obj != nil {
		errs = append(errs, ValidateObjectReference(path, obj.ObjectReference)...)
	}
	return errs
}