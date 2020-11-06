package structures

import (
	"k8s.io/kops/pkg/apis/kops"
)

func ExpandAuthorizationSpec(in map[string]interface{}) kops.AuthorizationSpec {
	if in == nil {
		panic("expand AuthorizationSpec failure, in is nil")
	}
	return kops.AuthorizationSpec{
		AlwaysAllow: func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
			return func(in interface{}) *kops.AlwaysAllowAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.AlwaysAllowAuthorizationSpec) *kops.AlwaysAllowAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.AlwaysAllowAuthorizationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.AlwaysAllowAuthorizationSpec{}
					}
					return (ExpandAlwaysAllowAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["always_allow"]),
		RBAC: func(in interface{}) *kops.RBACAuthorizationSpec {
			return func(in interface{}) *kops.RBACAuthorizationSpec {
				if in == nil {
					return nil
				}
				if _, ok := in.([]interface{}); ok && len(in.([]interface{})) == 0 {
					return nil
				}
				return func(in kops.RBACAuthorizationSpec) *kops.RBACAuthorizationSpec {
					return &in
				}(func(in interface{}) kops.RBACAuthorizationSpec {
					if len(in.([]interface{})) == 0 || in.([]interface{})[0] == nil {
						return kops.RBACAuthorizationSpec{}
					}
					return (ExpandRBACAuthorizationSpec(in.([]interface{})[0].(map[string]interface{})))
				}(in))
			}(in)
		}(in["rbac"]),
	}
}

func FlattenAuthorizationSpec(in kops.AuthorizationSpec) map[string]interface{} {
	return map[string]interface{}{
		"always_allow": func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
			return func(in *kops.AlwaysAllowAuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.AlwaysAllowAuthorizationSpec) interface{} {
					return func(in kops.AlwaysAllowAuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenAlwaysAllowAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.AlwaysAllow),
		"rbac": func(in *kops.RBACAuthorizationSpec) interface{} {
			return func(in *kops.RBACAuthorizationSpec) interface{} {
				if in == nil {
					return nil
				}
				return func(in kops.RBACAuthorizationSpec) interface{} {
					return func(in kops.RBACAuthorizationSpec) []map[string]interface{} {
						return []map[string]interface{}{FlattenRBACAuthorizationSpec(in)}
					}(in)
				}(*in)
			}(in)
		}(in.RBAC),
	}
}
