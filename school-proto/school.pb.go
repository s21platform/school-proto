// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: school.proto

package school_proto

import (
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

// Запрос на авторизацию
type SchoolLoginRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// email от edu.21-school.ru формата nickname@student.21-school.ru
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// Пароль пользователя от edu.21-school.ru
	Password      string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SchoolLoginRequest) Reset() {
	*x = SchoolLoginRequest{}
	mi := &file_school_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchoolLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolLoginRequest) ProtoMessage() {}

func (x *SchoolLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchoolLoginRequest.ProtoReflect.Descriptor instead.
func (*SchoolLoginRequest) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{0}
}

func (x *SchoolLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SchoolLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Ответ на запрос авторизации
type SchoolLoginResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Авторизационный токен JWT
	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SchoolLoginResponse) Reset() {
	*x = SchoolLoginResponse{}
	mi := &file_school_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SchoolLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchoolLoginResponse) ProtoMessage() {}

func (x *SchoolLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchoolLoginResponse.ProtoReflect.Descriptor instead.
func (*SchoolLoginResponse) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{1}
}

func (x *SchoolLoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// Запрос на получение всех кампусов
type Empty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_school_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{2}
}

// Сообщение для выходных данных
type Campus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Uuid кампуса
	CampusUuid string `protobuf:"bytes,1,opt,name=campusUuid,proto3" json:"campusUuid,omitempty"`
	// Сокращенное название кампуса
	ShortName string `protobuf:"bytes,2,opt,name=shortName,proto3" json:"shortName,omitempty"`
	// Полное название кампуса
	FullName      string `protobuf:"bytes,3,opt,name=fullName,proto3" json:"fullName,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Campus) Reset() {
	*x = Campus{}
	mi := &file_school_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Campus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Campus) ProtoMessage() {}

func (x *Campus) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Campus.ProtoReflect.Descriptor instead.
func (*Campus) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{3}
}

func (x *Campus) GetCampusUuid() string {
	if x != nil {
		return x.CampusUuid
	}
	return ""
}

func (x *Campus) GetShortName() string {
	if x != nil {
		return x.ShortName
	}
	return ""
}

func (x *Campus) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

// Ответ на запрос на получение всех кампусов
type CampusesOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Campuses      []*Campus              `protobuf:"bytes,1,rep,name=campuses,proto3" json:"campuses,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampusesOut) Reset() {
	*x = CampusesOut{}
	mi := &file_school_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampusesOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampusesOut) ProtoMessage() {}

func (x *CampusesOut) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampusesOut.ProtoReflect.Descriptor instead.
func (*CampusesOut) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{4}
}

func (x *CampusesOut) GetCampuses() []*Campus {
	if x != nil {
		return x.Campuses
	}
	return nil
}

// Запрос на получение всех трайбов кампуса
type CampusUuidIn struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Uuid кампуса, который призходит от community сервиса
	CampusUuid    string `protobuf:"bytes,1,opt,name=campus_uuid,json=campusUuid,proto3" json:"campus_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampusUuidIn) Reset() {
	*x = CampusUuidIn{}
	mi := &file_school_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampusUuidIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampusUuidIn) ProtoMessage() {}

func (x *CampusUuidIn) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampusUuidIn.ProtoReflect.Descriptor instead.
func (*CampusUuidIn) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{5}
}

func (x *CampusUuidIn) GetCampusUuid() string {
	if x != nil {
		return x.CampusUuid
	}
	return ""
}

// Сообщение для выходных данных
type Tribe struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id  тайба
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Название трайба
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tribe) Reset() {
	*x = Tribe{}
	mi := &file_school_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tribe) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tribe) ProtoMessage() {}

func (x *Tribe) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tribe.ProtoReflect.Descriptor instead.
func (*Tribe) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{6}
}

func (x *Tribe) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tribe) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Ответ на запрос на получение всех трайбов кампуса
type TribesOut struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Список трайбов
	Tribes        []*Tribe `protobuf:"bytes,1,rep,name=tribes,proto3" json:"tribes,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TribesOut) Reset() {
	*x = TribesOut{}
	mi := &file_school_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TribesOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TribesOut) ProtoMessage() {}

func (x *TribesOut) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TribesOut.ProtoReflect.Descriptor instead.
func (*TribesOut) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{7}
}

func (x *TribesOut) GetTribes() []*Tribe {
	if x != nil {
		return x.Tribes
	}
	return nil
}

// Запрос на получение списка пиров
type GetPeersIn struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Uuid кампуса
	CampusUuid string `protobuf:"bytes,1,opt,name=campusUuid,proto3" json:"campusUuid,omitempty"`
	// Kоличество записей для возвращения
	Limit int64 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	// Смещение
	Offset        int64 `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPeersIn) Reset() {
	*x = GetPeersIn{}
	mi := &file_school_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPeersIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeersIn) ProtoMessage() {}

func (x *GetPeersIn) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeersIn.ProtoReflect.Descriptor instead.
func (*GetPeersIn) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{8}
}

func (x *GetPeersIn) GetCampusUuid() string {
	if x != nil {
		return x.CampusUuid
	}
	return ""
}

func (x *GetPeersIn) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetPeersIn) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

// Ответ на запрос на получение списка пиров
type GetPeersOut struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Peer          []string               `protobuf:"bytes,1,rep,name=peer,proto3" json:"peer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPeersOut) Reset() {
	*x = GetPeersOut{}
	mi := &file_school_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPeersOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeersOut) ProtoMessage() {}

func (x *GetPeersOut) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeersOut.ProtoReflect.Descriptor instead.
func (*GetPeersOut) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{9}
}

func (x *GetPeersOut) GetPeer() []string {
	if x != nil {
		return x.Peer
	}
	return nil
}

type GetParticipantDataIn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Login         string                 `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetParticipantDataIn) Reset() {
	*x = GetParticipantDataIn{}
	mi := &file_school_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetParticipantDataIn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParticipantDataIn) ProtoMessage() {}

func (x *GetParticipantDataIn) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParticipantDataIn.ProtoReflect.Descriptor instead.
func (*GetParticipantDataIn) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{10}
}

func (x *GetParticipantDataIn) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

type GetParticipantDataOut struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	ClassName            string                 `protobuf:"bytes,1,opt,name=className,proto3" json:"className,omitempty"`
	ParallelName         string                 `protobuf:"bytes,2,opt,name=parallelName,proto3" json:"parallelName,omitempty"`
	ExpValue             int64                  `protobuf:"varint,3,opt,name=expValue,proto3" json:"expValue,omitempty"`
	Level                int32                  `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	ExpToNextLevel       int64                  `protobuf:"varint,5,opt,name=expToNextLevel,proto3" json:"expToNextLevel,omitempty"`
	CampusUuid           string                 `protobuf:"bytes,6,opt,name=campusUuid,proto3" json:"campusUuid,omitempty"`
	Status               string                 `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Skills               []*Skills              `protobuf:"bytes,8,rep,name=skills,proto3" json:"skills,omitempty"`
	PeerReviewPoints     int64                  `protobuf:"varint,9,opt,name=peerReviewPoints,proto3" json:"peerReviewPoints,omitempty"`
	PeerCodeReviewPoints int64                  `protobuf:"varint,10,opt,name=peerCodeReviewPoints,proto3" json:"peerCodeReviewPoints,omitempty"`
	Coins                int64                  `protobuf:"varint,11,opt,name=coins,proto3" json:"coins,omitempty"`
	Badges               []*Badges              `protobuf:"bytes,12,rep,name=badges,proto3" json:"badges,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *GetParticipantDataOut) Reset() {
	*x = GetParticipantDataOut{}
	mi := &file_school_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetParticipantDataOut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParticipantDataOut) ProtoMessage() {}

func (x *GetParticipantDataOut) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParticipantDataOut.ProtoReflect.Descriptor instead.
func (*GetParticipantDataOut) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{11}
}

func (x *GetParticipantDataOut) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *GetParticipantDataOut) GetParallelName() string {
	if x != nil {
		return x.ParallelName
	}
	return ""
}

func (x *GetParticipantDataOut) GetExpValue() int64 {
	if x != nil {
		return x.ExpValue
	}
	return 0
}

func (x *GetParticipantDataOut) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *GetParticipantDataOut) GetExpToNextLevel() int64 {
	if x != nil {
		return x.ExpToNextLevel
	}
	return 0
}

func (x *GetParticipantDataOut) GetCampusUuid() string {
	if x != nil {
		return x.CampusUuid
	}
	return ""
}

func (x *GetParticipantDataOut) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetParticipantDataOut) GetSkills() []*Skills {
	if x != nil {
		return x.Skills
	}
	return nil
}

func (x *GetParticipantDataOut) GetPeerReviewPoints() int64 {
	if x != nil {
		return x.PeerReviewPoints
	}
	return 0
}

func (x *GetParticipantDataOut) GetPeerCodeReviewPoints() int64 {
	if x != nil {
		return x.PeerCodeReviewPoints
	}
	return 0
}

func (x *GetParticipantDataOut) GetCoins() int64 {
	if x != nil {
		return x.Coins
	}
	return 0
}

func (x *GetParticipantDataOut) GetBadges() []*Badges {
	if x != nil {
		return x.Badges
	}
	return nil
}

type Skills struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Points        int32                  `protobuf:"varint,2,opt,name=points,proto3" json:"points,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Skills) Reset() {
	*x = Skills{}
	mi := &file_school_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Skills) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skills) ProtoMessage() {}

func (x *Skills) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skills.ProtoReflect.Descriptor instead.
func (*Skills) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{12}
}

func (x *Skills) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Skills) GetPoints() int32 {
	if x != nil {
		return x.Points
	}
	return 0
}

type Badges struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ReceiptDateTime string                 `protobuf:"bytes,2,opt,name=receiptDateTime,proto3" json:"receiptDateTime,omitempty"`
	IconURL         string                 `protobuf:"bytes,3,opt,name=iconURL,proto3" json:"iconURL,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Badges) Reset() {
	*x = Badges{}
	mi := &file_school_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Badges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Badges) ProtoMessage() {}

func (x *Badges) ProtoReflect() protoreflect.Message {
	mi := &file_school_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Badges.ProtoReflect.Descriptor instead.
func (*Badges) Descriptor() ([]byte, []int) {
	return file_school_proto_rawDescGZIP(), []int{13}
}

func (x *Badges) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Badges) GetReceiptDateTime() string {
	if x != nil {
		return x.ReceiptDateTime
	}
	return ""
}

func (x *Badges) GetIconURL() string {
	if x != nil {
		return x.IconURL
	}
	return ""
}

var File_school_proto protoreflect.FileDescriptor

var file_school_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46,
	0x0a, 0x12, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x62, 0x0a, 0x06,
	0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73,
	0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x32, 0x0a, 0x0b, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x12,
	0x23, 0x0a, 0x08, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x52, 0x08, 0x63, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x65, 0x73, 0x22, 0x2f, 0x0a, 0x0c, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75,
	0x69, 0x64, 0x49, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x75,
	0x73, 0x55, 0x75, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x05, 0x54, 0x72, 0x69, 0x62, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x09, 0x54, 0x72, 0x69, 0x62, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x12,
	0x1e, 0x0a, 0x06, 0x74, 0x72, 0x69, 0x62, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x54, 0x72, 0x69, 0x62, 0x65, 0x52, 0x06, 0x74, 0x72, 0x69, 0x62, 0x65, 0x73, 0x22,
	0x5a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x21, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x65,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x70, 0x65, 0x65, 0x72, 0x22, 0x2c,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x22, 0xa3, 0x03, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x4f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x61,
	0x6c, 0x6c, 0x65, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x65, 0x78, 0x70, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78,
	0x70, 0x54, 0x6f, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x54, 0x6f, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x6b,
	0x69, 0x6c, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x6b, 0x69,
	0x6c, 0x6c, 0x73, 0x52, 0x06, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x70,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x70, 0x65, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x65, 0x65, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x14, 0x70, 0x65, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x69, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x69, 0x6e,
	0x73, 0x12, 0x1f, 0x0a, 0x06, 0x62, 0x61, 0x64, 0x67, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x42, 0x61, 0x64, 0x67, 0x65, 0x73, 0x52, 0x06, 0x62, 0x61, 0x64, 0x67,
	0x65, 0x73, 0x22, 0x34, 0x0a, 0x06, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x60, 0x0a, 0x06, 0x42, 0x61, 0x64, 0x67,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x44, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x69, 0x63, 0x6f, 0x6e, 0x55, 0x52, 0x4c, 0x32, 0x92, 0x02, 0x0a, 0x0d, 0x53,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x53, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x63, 0x68,
	0x6f, 0x6f, 0x6c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x25, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x43, 0x61, 0x6d, 0x70,
	0x75, 0x73, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x54, 0x72, 0x69, 0x62, 0x65, 0x73, 0x42, 0x79, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75,
	0x69, 0x64, 0x12, 0x0d, 0x2e, 0x43, 0x61, 0x6d, 0x70, 0x75, 0x73, 0x55, 0x75, 0x69, 0x64, 0x49,
	0x6e, 0x1a, 0x0a, 0x2e, 0x54, 0x72, 0x69, 0x62, 0x65, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12,
	0x27, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x12, 0x0b, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x65, 0x72, 0x73, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x15,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x49, 0x6e, 0x1a, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x4f, 0x75, 0x74, 0x22, 0x00, 0x42,
	0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_school_proto_rawDescOnce sync.Once
	file_school_proto_rawDescData []byte
)

func file_school_proto_rawDescGZIP() []byte {
	file_school_proto_rawDescOnce.Do(func() {
		file_school_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_school_proto_rawDesc), len(file_school_proto_rawDesc)))
	})
	return file_school_proto_rawDescData
}

var file_school_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_school_proto_goTypes = []any{
	(*SchoolLoginRequest)(nil),    // 0: SchoolLoginRequest
	(*SchoolLoginResponse)(nil),   // 1: SchoolLoginResponse
	(*Empty)(nil),                 // 2: Empty
	(*Campus)(nil),                // 3: Campus
	(*CampusesOut)(nil),           // 4: CampusesOut
	(*CampusUuidIn)(nil),          // 5: CampusUuidIn
	(*Tribe)(nil),                 // 6: Tribe
	(*TribesOut)(nil),             // 7: TribesOut
	(*GetPeersIn)(nil),            // 8: GetPeersIn
	(*GetPeersOut)(nil),           // 9: GetPeersOut
	(*GetParticipantDataIn)(nil),  // 10: GetParticipantDataIn
	(*GetParticipantDataOut)(nil), // 11: GetParticipantDataOut
	(*Skills)(nil),                // 12: Skills
	(*Badges)(nil),                // 13: Badges
}
var file_school_proto_depIdxs = []int32{
	3,  // 0: CampusesOut.campuses:type_name -> Campus
	6,  // 1: TribesOut.tribes:type_name -> Tribe
	12, // 2: GetParticipantDataOut.skills:type_name -> Skills
	13, // 3: GetParticipantDataOut.badges:type_name -> Badges
	0,  // 4: SchoolService.Login:input_type -> SchoolLoginRequest
	2,  // 5: SchoolService.GetCampuses:input_type -> Empty
	5,  // 6: SchoolService.GetTribesByCampusUuid:input_type -> CampusUuidIn
	8,  // 7: SchoolService.GetPeers:input_type -> GetPeersIn
	10, // 8: SchoolService.GetParticipantData:input_type -> GetParticipantDataIn
	1,  // 9: SchoolService.Login:output_type -> SchoolLoginResponse
	4,  // 10: SchoolService.GetCampuses:output_type -> CampusesOut
	7,  // 11: SchoolService.GetTribesByCampusUuid:output_type -> TribesOut
	9,  // 12: SchoolService.GetPeers:output_type -> GetPeersOut
	11, // 13: SchoolService.GetParticipantData:output_type -> GetParticipantDataOut
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_school_proto_init() }
func file_school_proto_init() {
	if File_school_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_school_proto_rawDesc), len(file_school_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_school_proto_goTypes,
		DependencyIndexes: file_school_proto_depIdxs,
		MessageInfos:      file_school_proto_msgTypes,
	}.Build()
	File_school_proto = out.File
	file_school_proto_goTypes = nil
	file_school_proto_depIdxs = nil
}
