// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.3
// source: location_weather.proto

package location_weather

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WeatherData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LocationId    int32                  `protobuf:"varint,1,opt,name=LocationId,proto3" json:"LocationId,omitempty"`
	Temperature   float64                `protobuf:"fixed64,2,opt,name=Temperature,proto3" json:"Temperature,omitempty"`
	Humidity      int32                  `protobuf:"varint,3,opt,name=Humidity,proto3" json:"Humidity,omitempty"`
	WindSpeed     float64                `protobuf:"fixed64,4,opt,name=WindSpeed,proto3" json:"WindSpeed,omitempty"`
	Pressure      float64                `protobuf:"fixed64,5,opt,name=Pressure,proto3" json:"Pressure,omitempty"`
	Precip        float64                `protobuf:"fixed64,6,opt,name=Precip,proto3" json:"Precip,omitempty"`
	Cloud         int32                  `protobuf:"varint,7,opt,name=Cloud,proto3" json:"Cloud,omitempty"`
	UpdateAt      string                 `protobuf:"bytes,8,opt,name=UpdateAt,proto3" json:"UpdateAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WeatherData) Reset() {
	*x = WeatherData{}
	mi := &file_location_weather_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeatherData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeatherData) ProtoMessage() {}

func (x *WeatherData) ProtoReflect() protoreflect.Message {
	mi := &file_location_weather_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeatherData.ProtoReflect.Descriptor instead.
func (*WeatherData) Descriptor() ([]byte, []int) {
	return file_location_weather_proto_rawDescGZIP(), []int{0}
}

func (x *WeatherData) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

func (x *WeatherData) GetTemperature() float64 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *WeatherData) GetHumidity() int32 {
	if x != nil {
		return x.Humidity
	}
	return 0
}

func (x *WeatherData) GetWindSpeed() float64 {
	if x != nil {
		return x.WindSpeed
	}
	return 0
}

func (x *WeatherData) GetPressure() float64 {
	if x != nil {
		return x.Pressure
	}
	return 0
}

func (x *WeatherData) GetPrecip() float64 {
	if x != nil {
		return x.Precip
	}
	return 0
}

func (x *WeatherData) GetCloud() int32 {
	if x != nil {
		return x.Cloud
	}
	return 0
}

func (x *WeatherData) GetUpdateAt() string {
	if x != nil {
		return x.UpdateAt
	}
	return ""
}

type GetLocationByIdRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LocationId    int32                  `protobuf:"varint,1,opt,name=locationId,proto3" json:"locationId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLocationByIdRequest) Reset() {
	*x = GetLocationByIdRequest{}
	mi := &file_location_weather_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLocationByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationByIdRequest) ProtoMessage() {}

func (x *GetLocationByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_location_weather_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationByIdRequest.ProtoReflect.Descriptor instead.
func (*GetLocationByIdRequest) Descriptor() ([]byte, []int) {
	return file_location_weather_proto_rawDescGZIP(), []int{1}
}

func (x *GetLocationByIdRequest) GetLocationId() int32 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

type GetLocationByIdResponce struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Weather       *WeatherData           `protobuf:"bytes,1,opt,name=Weather,proto3" json:"Weather,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLocationByIdResponce) Reset() {
	*x = GetLocationByIdResponce{}
	mi := &file_location_weather_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLocationByIdResponce) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationByIdResponce) ProtoMessage() {}

func (x *GetLocationByIdResponce) ProtoReflect() protoreflect.Message {
	mi := &file_location_weather_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationByIdResponce.ProtoReflect.Descriptor instead.
func (*GetLocationByIdResponce) Descriptor() ([]byte, []int) {
	return file_location_weather_proto_rawDescGZIP(), []int{2}
}

func (x *GetLocationByIdResponce) GetWeather() *WeatherData {
	if x != nil {
		return x.Weather
	}
	return nil
}

var File_location_weather_proto protoreflect.FileDescriptor

var file_location_weather_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0b, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x54,
	0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x48, 0x75,
	0x6d, 0x69, 0x64, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x57, 0x69, 0x6e, 0x64, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x50, 0x72, 0x65, 0x73, 0x73, 0x75, 0x72, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x65, 0x63, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x06, 0x50, 0x72, 0x65, 0x63, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x75,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x22, 0x38, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x07, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x61, 0x74,
	0x68, 0x65, 0x72, 0x2e, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x07, 0x57, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x32, 0x9c, 0x01, 0x0a, 0x0e, 0x57, 0x65, 0x61,
	0x74, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x89, 0x01, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x28, 0x2e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x61, 0x74, 0x68,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x7d, 0x42, 0x66, 0x5a, 0x64, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x56, 0x6f, 0x76, 0x61, 0x2d, 0x6c, 0x75, 0x6b, 0x2f, 0x77,
	0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x65, 0x61, 0x74, 0x68, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_location_weather_proto_rawDescOnce sync.Once
	file_location_weather_proto_rawDescData []byte
)

func file_location_weather_proto_rawDescGZIP() []byte {
	file_location_weather_proto_rawDescOnce.Do(func() {
		file_location_weather_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_location_weather_proto_rawDesc), len(file_location_weather_proto_rawDesc)))
	})
	return file_location_weather_proto_rawDescData
}

var file_location_weather_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_location_weather_proto_goTypes = []any{
	(*WeatherData)(nil),             // 0: location_weather.WeatherData
	(*GetLocationByIdRequest)(nil),  // 1: location_weather.GetLocationByIdRequest
	(*GetLocationByIdResponce)(nil), // 2: location_weather.GetLocationByIdResponce
}
var file_location_weather_proto_depIdxs = []int32{
	0, // 0: location_weather.GetLocationByIdResponce.Weather:type_name -> location_weather.WeatherData
	1, // 1: location_weather.WeatherService.GetLocationById:input_type -> location_weather.GetLocationByIdRequest
	2, // 2: location_weather.WeatherService.GetLocationById:output_type -> location_weather.GetLocationByIdResponce
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_location_weather_proto_init() }
func file_location_weather_proto_init() {
	if File_location_weather_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_location_weather_proto_rawDesc), len(file_location_weather_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_location_weather_proto_goTypes,
		DependencyIndexes: file_location_weather_proto_depIdxs,
		MessageInfos:      file_location_weather_proto_msgTypes,
	}.Build()
	File_location_weather_proto = out.File
	file_location_weather_proto_goTypes = nil
	file_location_weather_proto_depIdxs = nil
}
