// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: autotest-plan.proto

package pb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// TestPlanUpdateByHooksRequest update testPlan by hook (post /api/autotests/actions/plan-execute-callback )
type TestPlanUpdateByHookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event         string   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Action        string   `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	OrgID         string   `protobuf:"bytes,3,opt,name=orgID,proto3" json:"orgID,omitempty"`
	ProjectID     string   `protobuf:"bytes,4,opt,name=projectID,proto3" json:"projectID,omitempty"`
	ApplicationID string   `protobuf:"bytes,5,opt,name=applicationID,proto3" json:"applicationID,omitempty"`
	Env           string   `protobuf:"bytes,6,opt,name=env,proto3" json:"env,omitempty"`
	Timestamp     string   `protobuf:"bytes,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Content       *Content `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *TestPlanUpdateByHookRequest) Reset() {
	*x = TestPlanUpdateByHookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autotest_plan_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPlanUpdateByHookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPlanUpdateByHookRequest) ProtoMessage() {}

func (x *TestPlanUpdateByHookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autotest_plan_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPlanUpdateByHookRequest.ProtoReflect.Descriptor instead.
func (*TestPlanUpdateByHookRequest) Descriptor() ([]byte, []int) {
	return file_autotest_plan_proto_rawDescGZIP(), []int{0}
}

func (x *TestPlanUpdateByHookRequest) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *TestPlanUpdateByHookRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *TestPlanUpdateByHookRequest) GetOrgID() string {
	if x != nil {
		return x.OrgID
	}
	return ""
}

func (x *TestPlanUpdateByHookRequest) GetProjectID() string {
	if x != nil {
		return x.ProjectID
	}
	return ""
}

func (x *TestPlanUpdateByHookRequest) GetApplicationID() string {
	if x != nil {
		return x.ApplicationID
	}
	return ""
}

func (x *TestPlanUpdateByHookRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *TestPlanUpdateByHookRequest) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *TestPlanUpdateByHookRequest) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestPlanID  uint64  `protobuf:"varint,1,opt,name=testPlanID,proto3" json:"testPlanID,omitempty"`
	ExecuteTime string  `protobuf:"bytes,2,opt,name=executeTime,proto3" json:"executeTime,omitempty"`
	PassRate    float64 `protobuf:"fixed64,3,opt,name=passRate,proto3" json:"passRate,omitempty"`
	ApiTotalNum int64   `protobuf:"varint,4,opt,name=apiTotalNum,proto3" json:"apiTotalNum,omitempty"`
	// time duration of the test plan execution
	ExecuteDuration string `protobuf:"bytes,5,opt,name=executeDuration,proto3" json:"executeDuration,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autotest_plan_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_autotest_plan_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_autotest_plan_proto_rawDescGZIP(), []int{1}
}

func (x *Content) GetTestPlanID() uint64 {
	if x != nil {
		return x.TestPlanID
	}
	return 0
}

func (x *Content) GetExecuteTime() string {
	if x != nil {
		return x.ExecuteTime
	}
	return ""
}

func (x *Content) GetPassRate() float64 {
	if x != nil {
		return x.PassRate
	}
	return 0
}

func (x *Content) GetApiTotalNum() int64 {
	if x != nil {
		return x.ApiTotalNum
	}
	return 0
}

func (x *Content) GetExecuteDuration() string {
	if x != nil {
		return x.ExecuteDuration
	}
	return ""
}

type TestPlanUpdateByHookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data uint64 `protobuf:"varint,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TestPlanUpdateByHookResponse) Reset() {
	*x = TestPlanUpdateByHookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autotest_plan_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestPlanUpdateByHookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestPlanUpdateByHookResponse) ProtoMessage() {}

func (x *TestPlanUpdateByHookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autotest_plan_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestPlanUpdateByHookResponse.ProtoReflect.Descriptor instead.
func (*TestPlanUpdateByHookResponse) Descriptor() ([]byte, []int) {
	return file_autotest_plan_proto_rawDescGZIP(), []int{2}
}

func (x *TestPlanUpdateByHookResponse) GetData() uint64 {
	if x != nil {
		return x.Data
	}
	return 0
}

var File_autotest_plan_proto protoreflect.FileDescriptor

var file_autotest_plan_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x02, 0x0a, 0x1b, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x42, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0xb3, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x49, 0x44, 0x12, 0x20, 0x0a,
	0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x52, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x70, 0x69, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x61, 0x70, 0x69, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x12, 0x28, 0x0a,
	0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x1c, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x48, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xdd, 0x01, 0x0a, 0x0f,
	0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0xc9, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c,
	0x61, 0x6e, 0x42, 0x79, 0x48, 0x6f, 0x6f, 0x6b, 0x12, 0x3c, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50,
	0x6c, 0x61, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x48, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x65, 0x72, 0x64, 0x61, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x64, 0x6f, 0x70, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61,
	0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x48, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x2c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x2d, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x2d, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x65, 0x72, 0x64, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x64, 0x6f, 0x70, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x6c, 0x61, 0x6e, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autotest_plan_proto_rawDescOnce sync.Once
	file_autotest_plan_proto_rawDescData = file_autotest_plan_proto_rawDesc
)

func file_autotest_plan_proto_rawDescGZIP() []byte {
	file_autotest_plan_proto_rawDescOnce.Do(func() {
		file_autotest_plan_proto_rawDescData = protoimpl.X.CompressGZIP(file_autotest_plan_proto_rawDescData)
	})
	return file_autotest_plan_proto_rawDescData
}

var file_autotest_plan_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_autotest_plan_proto_goTypes = []interface{}{
	(*TestPlanUpdateByHookRequest)(nil),  // 0: erda.core.dop.autotest.testplan.TestPlanUpdateByHookRequest
	(*Content)(nil),                      // 1: erda.core.dop.autotest.testplan.Content
	(*TestPlanUpdateByHookResponse)(nil), // 2: erda.core.dop.autotest.testplan.TestPlanUpdateByHookResponse
}
var file_autotest_plan_proto_depIdxs = []int32{
	1, // 0: erda.core.dop.autotest.testplan.TestPlanUpdateByHookRequest.content:type_name -> erda.core.dop.autotest.testplan.Content
	0, // 1: erda.core.dop.autotest.testplan.TestPlanService.UpdateTestPlanByHook:input_type -> erda.core.dop.autotest.testplan.TestPlanUpdateByHookRequest
	2, // 2: erda.core.dop.autotest.testplan.TestPlanService.UpdateTestPlanByHook:output_type -> erda.core.dop.autotest.testplan.TestPlanUpdateByHookResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_autotest_plan_proto_init() }
func file_autotest_plan_proto_init() {
	if File_autotest_plan_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autotest_plan_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPlanUpdateByHookRequest); i {
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
		file_autotest_plan_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_autotest_plan_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestPlanUpdateByHookResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_autotest_plan_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autotest_plan_proto_goTypes,
		DependencyIndexes: file_autotest_plan_proto_depIdxs,
		MessageInfos:      file_autotest_plan_proto_msgTypes,
	}.Build()
	File_autotest_plan_proto = out.File
	file_autotest_plan_proto_rawDesc = nil
	file_autotest_plan_proto_goTypes = nil
	file_autotest_plan_proto_depIdxs = nil
}
