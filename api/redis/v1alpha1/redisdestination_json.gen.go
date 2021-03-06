// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/redis/v1alpha1/redisdestination.proto

// RedisDestination defines policies that apply to redis traffic intended for a redis service

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "istio.io/api/networking/v1alpha3"
	_ "istio.io/gogo-genproto/googleapis/google/api"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for RedisDestination
func (this *RedisDestination) MarshalJSON() ([]byte, error) {
	str, err := RedisdestinationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RedisDestination
func (this *RedisDestination) UnmarshalJSON(b []byte) error {
	return RedisdestinationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for ConnectionPoolSettings
func (this *ConnectionPoolSettings) MarshalJSON() ([]byte, error) {
	str, err := RedisdestinationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ConnectionPoolSettings
func (this *ConnectionPoolSettings) UnmarshalJSON(b []byte) error {
	return RedisdestinationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for Auth
func (this *Auth) MarshalJSON() ([]byte, error) {
	str, err := RedisdestinationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for Auth
func (this *Auth) UnmarshalJSON(b []byte) error {
	return RedisdestinationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for SecretReference
func (this *SecretReference) MarshalJSON() ([]byte, error) {
	str, err := RedisdestinationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for SecretReference
func (this *SecretReference) UnmarshalJSON(b []byte) error {
	return RedisdestinationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PlainAuth
func (this *PlainAuth) MarshalJSON() ([]byte, error) {
	str, err := RedisdestinationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PlainAuth
func (this *PlainAuth) UnmarshalJSON(b []byte) error {
	return RedisdestinationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for RedisSettings
func (this *RedisSettings) MarshalJSON() ([]byte, error) {
	str, err := RedisdestinationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for RedisSettings
func (this *RedisSettings) UnmarshalJSON(b []byte) error {
	return RedisdestinationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TrafficPolicy
func (this *TrafficPolicy) MarshalJSON() ([]byte, error) {
	str, err := RedisdestinationMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TrafficPolicy
func (this *TrafficPolicy) UnmarshalJSON(b []byte) error {
	return RedisdestinationUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	RedisdestinationMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	RedisdestinationUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
