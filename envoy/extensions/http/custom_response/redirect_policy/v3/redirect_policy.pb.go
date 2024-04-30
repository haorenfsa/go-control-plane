// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.25.3
// source: envoy/extensions/http/custom_response/redirect_policy/v3/redirect_policy.proto

package redirect_policyv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/cncf/xds/go/xds/annotations/v3"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	v31 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Custom response policy to internally redirect the original response to a different
// upstream.
// [#next-free-field: 7]
type RedirectPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to RedirectActionSpecifier:
	//
	//	*RedirectPolicy_Uri
	//	*RedirectPolicy_RedirectAction
	RedirectActionSpecifier isRedirectPolicy_RedirectActionSpecifier `protobuf_oneof:"redirect_action_specifier"`
	// The new response status code if specified. This is used to override the
	// status code of the response from the new upstream if it is not an error status.
	StatusCode *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// HTTP headers to add to the response. This allows the
	// response policy to append, to add or to override headers of
	// the original response for local body, or the custom response from the
	// remote body, before it is sent to a downstream client.
	// Note that these are not applied if the redirected response is an error
	// response.
	ResponseHeadersToAdd []*v3.HeaderValueOption `protobuf:"bytes,4,rep,name=response_headers_to_add,json=responseHeadersToAdd,proto3" json:"response_headers_to_add,omitempty"`
	// HTTP headers to add to the request before it is internally redirected.
	RequestHeadersToAdd []*v3.HeaderValueOption `protobuf:"bytes,5,rep,name=request_headers_to_add,json=requestHeadersToAdd,proto3" json:"request_headers_to_add,omitempty"`
	// Custom action to modify request headers before selection of the
	// redirected route.
	// [#comment: TODO(pradeepcrao) add an extension category.]
	ModifyRequestHeadersAction *v3.TypedExtensionConfig `protobuf:"bytes,6,opt,name=modify_request_headers_action,json=modifyRequestHeadersAction,proto3" json:"modify_request_headers_action,omitempty"`
}

func (x *RedirectPolicy) Reset() {
	*x = RedirectPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedirectPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedirectPolicy) ProtoMessage() {}

func (x *RedirectPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedirectPolicy.ProtoReflect.Descriptor instead.
func (*RedirectPolicy) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescGZIP(), []int{0}
}

func (m *RedirectPolicy) GetRedirectActionSpecifier() isRedirectPolicy_RedirectActionSpecifier {
	if m != nil {
		return m.RedirectActionSpecifier
	}
	return nil
}

func (x *RedirectPolicy) GetUri() string {
	if x, ok := x.GetRedirectActionSpecifier().(*RedirectPolicy_Uri); ok {
		return x.Uri
	}
	return ""
}

func (x *RedirectPolicy) GetRedirectAction() *v31.RedirectAction {
	if x, ok := x.GetRedirectActionSpecifier().(*RedirectPolicy_RedirectAction); ok {
		return x.RedirectAction
	}
	return nil
}

func (x *RedirectPolicy) GetStatusCode() *wrapperspb.UInt32Value {
	if x != nil {
		return x.StatusCode
	}
	return nil
}

func (x *RedirectPolicy) GetResponseHeadersToAdd() []*v3.HeaderValueOption {
	if x != nil {
		return x.ResponseHeadersToAdd
	}
	return nil
}

func (x *RedirectPolicy) GetRequestHeadersToAdd() []*v3.HeaderValueOption {
	if x != nil {
		return x.RequestHeadersToAdd
	}
	return nil
}

func (x *RedirectPolicy) GetModifyRequestHeadersAction() *v3.TypedExtensionConfig {
	if x != nil {
		return x.ModifyRequestHeadersAction
	}
	return nil
}

type isRedirectPolicy_RedirectActionSpecifier interface {
	isRedirectPolicy_RedirectActionSpecifier()
}

type RedirectPolicy_Uri struct {
	// The Http URI to redirect the original request to, to get the custom
	// response.
	// It should be a full FQDN with protocol, host and path.
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//	uri: https://www.mydomain.com/path/to/404.txt
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3,oneof"`
}

type RedirectPolicy_RedirectAction struct {
	// Specify elements of the redirect url individually.
	// Note: Do not specify the “response_code“ field in “redirect_action“, use
	// “status_code“ instead.
	// The following fields in “redirect_action“ are currently not supported,
	// and specifying them will cause the config to be rejected:
	// - “prefix_rewrite“
	// - “regex_rewrite“
	RedirectAction *v31.RedirectAction `protobuf:"bytes,2,opt,name=redirect_action,json=redirectAction,proto3,oneof"`
}

func (*RedirectPolicy_Uri) isRedirectPolicy_RedirectActionSpecifier() {}

func (*RedirectPolicy_RedirectAction) isRedirectPolicy_RedirectActionSpecifier() {}

var File_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto protoreflect.FileDescriptor

var file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x78, 0x64, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x04, 0x0a, 0x0e, 0x52, 0x65, 0x64,
	0x69, 0x72, 0x65, 0x63, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x1b, 0x0a, 0x03, 0x75,
	0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x50, 0x0a, 0x0f, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0xfa,
	0x42, 0x07, 0x2a, 0x05, 0x18, 0xe7, 0x07, 0x28, 0x64, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x69, 0x0a, 0x17, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x09, 0xfa, 0x42, 0x06, 0x92, 0x01, 0x03, 0x10, 0xe8, 0x07, 0x52, 0x14, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64,
	0x12, 0x67, 0x0a, 0x16, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x92, 0x01,
	0x03, 0x10, 0xe8, 0x07, 0x52, 0x13, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x12, 0x6d, 0x0a, 0x1d, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x1a, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x20, 0x0a, 0x19, 0x72, 0x65, 0x64, 0x69,
	0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0xe2, 0x01, 0xba, 0x80, 0xc8,
	0xd1, 0x06, 0x02, 0x10, 0x02, 0xd2, 0xc6, 0xa4, 0xe1, 0x06, 0x02, 0x08, 0x01, 0x0a, 0x46, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x13, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x71, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x72, 0x65,
	0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x76, 0x33, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescOnce sync.Once
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescData = file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDesc
)

func file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescData)
	})
	return file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDescData
}

var file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_goTypes = []interface{}{
	(*RedirectPolicy)(nil),          // 0: envoy.extensions.http.custom_response.redirect_policy.v3.RedirectPolicy
	(*v31.RedirectAction)(nil),      // 1: envoy.config.route.v3.RedirectAction
	(*wrapperspb.UInt32Value)(nil),  // 2: google.protobuf.UInt32Value
	(*v3.HeaderValueOption)(nil),    // 3: envoy.config.core.v3.HeaderValueOption
	(*v3.TypedExtensionConfig)(nil), // 4: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_depIdxs = []int32{
	1, // 0: envoy.extensions.http.custom_response.redirect_policy.v3.RedirectPolicy.redirect_action:type_name -> envoy.config.route.v3.RedirectAction
	2, // 1: envoy.extensions.http.custom_response.redirect_policy.v3.RedirectPolicy.status_code:type_name -> google.protobuf.UInt32Value
	3, // 2: envoy.extensions.http.custom_response.redirect_policy.v3.RedirectPolicy.response_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	3, // 3: envoy.extensions.http.custom_response.redirect_policy.v3.RedirectPolicy.request_headers_to_add:type_name -> envoy.config.core.v3.HeaderValueOption
	4, // 4: envoy.extensions.http.custom_response.redirect_policy.v3.RedirectPolicy.modify_request_headers_action:type_name -> envoy.config.core.v3.TypedExtensionConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_init()
}
func file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_init() {
	if File_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedirectPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*RedirectPolicy_Uri)(nil),
		(*RedirectPolicy_RedirectAction)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto = out.File
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_rawDesc = nil
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_goTypes = nil
	file_envoy_extensions_http_custom_response_redirect_policy_v3_redirect_policy_proto_depIdxs = nil
}
