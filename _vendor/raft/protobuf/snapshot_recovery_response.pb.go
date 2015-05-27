// Code generated by protoc-gen-gogo.
// source: snapshot_recovery_response.proto
// DO NOT EDIT!

package protobuf

import proto "golang.org/x/gogoprotobuf/proto"
import math "math"

// discarding unused import gogoproto "golang.org/x/gogoprotobuf/gogoproto/gogo.pb"

import io6 "io"
import fmt24 "fmt"
import code_google_com_p_gogoprotobuf_proto12 "golang.org/x/gogoprotobuf/proto"

import fmt25 "fmt"
import strings12 "strings"
import reflect12 "reflect"

import fmt26 "fmt"
import strings13 "strings"
import code_google_com_p_gogoprotobuf_proto13 "golang.org/x/gogoprotobuf/proto"
import sort6 "sort"
import strconv6 "strconv"
import reflect13 "reflect"

import fmt27 "fmt"
import bytes6 "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type SnapshotRecoveryResponse struct {
	Term             *uint64 `protobuf:"varint,1,req" json:"Term,omitempty"`
	Success          *bool   `protobuf:"varint,2,req" json:"Success,omitempty"`
	CommitIndex      *uint64 `protobuf:"varint,3,req" json:"CommitIndex,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *SnapshotRecoveryResponse) Reset()      { *m = SnapshotRecoveryResponse{} }
func (*SnapshotRecoveryResponse) ProtoMessage() {}

func (m *SnapshotRecoveryResponse) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

func (m *SnapshotRecoveryResponse) GetSuccess() bool {
	if m != nil && m.Success != nil {
		return *m.Success
	}
	return false
}

func (m *SnapshotRecoveryResponse) GetCommitIndex() uint64 {
	if m != nil && m.CommitIndex != nil {
		return *m.CommitIndex
	}
	return 0
}

func init() {
}
func (m *SnapshotRecoveryResponse) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io6.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt24.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io6.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Term = &v
		case 2:
			if wireType != 0 {
				return fmt24.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io6.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.Success = &b
		case 3:
			if wireType != 0 {
				return fmt24.Errorf("proto: wrong wireType = %d for field CommitIndex", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io6.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CommitIndex = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto12.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io6.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (this *SnapshotRecoveryResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings12.Join([]string{`&SnapshotRecoveryResponse{`,
		`Term:` + valueToStringSnapshotRecoveryResponse(this.Term) + `,`,
		`Success:` + valueToStringSnapshotRecoveryResponse(this.Success) + `,`,
		`CommitIndex:` + valueToStringSnapshotRecoveryResponse(this.CommitIndex) + `,`,
		`XXX_unrecognized:` + fmt25.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSnapshotRecoveryResponse(v interface{}) string {
	rv := reflect12.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect12.Indirect(rv).Interface()
	return fmt25.Sprintf("*%v", pv)
}
func (m *SnapshotRecoveryResponse) Size() (n int) {
	var l int
	_ = l
	if m.Term != nil {
		n += 1 + sovSnapshotRecoveryResponse(uint64(*m.Term))
	}
	if m.Success != nil {
		n += 2
	}
	if m.CommitIndex != nil {
		n += 1 + sovSnapshotRecoveryResponse(uint64(*m.CommitIndex))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSnapshotRecoveryResponse(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSnapshotRecoveryResponse(x uint64) (n int) {
	return sovSnapshotRecoveryResponse(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func NewPopulatedSnapshotRecoveryResponse(r randySnapshotRecoveryResponse, easy bool) *SnapshotRecoveryResponse {
	this := &SnapshotRecoveryResponse{}
	v1 := uint64(r.Uint32())
	this.Term = &v1
	v2 := bool(r.Intn(2) == 0)
	this.Success = &v2
	v3 := uint64(r.Uint32())
	this.CommitIndex = &v3
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedSnapshotRecoveryResponse(r, 4)
	}
	return this
}

type randySnapshotRecoveryResponse interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneSnapshotRecoveryResponse(r randySnapshotRecoveryResponse) rune {
	res := rune(r.Uint32() % 1112064)
	if 55296 <= res {
		res += 2047
	}
	return res
}
func randStringSnapshotRecoveryResponse(r randySnapshotRecoveryResponse) string {
	v4 := r.Intn(100)
	tmps := make([]rune, v4)
	for i := 0; i < v4; i++ {
		tmps[i] = randUTF8RuneSnapshotRecoveryResponse(r)
	}
	return string(tmps)
}
func randUnrecognizedSnapshotRecoveryResponse(r randySnapshotRecoveryResponse, maxFieldNumber int) (data []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		data = randFieldSnapshotRecoveryResponse(data, r, fieldNumber, wire)
	}
	return data
}
func randFieldSnapshotRecoveryResponse(data []byte, r randySnapshotRecoveryResponse, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		data = encodeVarintPopulateSnapshotRecoveryResponse(data, uint64(key))
		v5 := r.Int63()
		if r.Intn(2) == 0 {
			v5 *= -1
		}
		data = encodeVarintPopulateSnapshotRecoveryResponse(data, uint64(v5))
	case 1:
		data = encodeVarintPopulateSnapshotRecoveryResponse(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		data = encodeVarintPopulateSnapshotRecoveryResponse(data, uint64(key))
		ll := r.Intn(100)
		data = encodeVarintPopulateSnapshotRecoveryResponse(data, uint64(ll))
		for j := 0; j < ll; j++ {
			data = append(data, byte(r.Intn(256)))
		}
	default:
		data = encodeVarintPopulateSnapshotRecoveryResponse(data, uint64(key))
		data = append(data, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return data
}
func encodeVarintPopulateSnapshotRecoveryResponse(data []byte, v uint64) []byte {
	for v >= 1<<7 {
		data = append(data, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	data = append(data, uint8(v))
	return data
}
func (m *SnapshotRecoveryResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *SnapshotRecoveryResponse) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Term != nil {
		data[i] = 0x8
		i++
		i = encodeVarintSnapshotRecoveryResponse(data, i, uint64(*m.Term))
	}
	if m.Success != nil {
		data[i] = 0x10
		i++
		if *m.Success {
			data[i] = 1
		} else {
			data[i] = 0
		}
		i++
	}
	if m.CommitIndex != nil {
		data[i] = 0x18
		i++
		i = encodeVarintSnapshotRecoveryResponse(data, i, uint64(*m.CommitIndex))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64SnapshotRecoveryResponse(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32SnapshotRecoveryResponse(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSnapshotRecoveryResponse(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (this *SnapshotRecoveryResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings13.Join([]string{`&protobuf.SnapshotRecoveryResponse{` + `Term:` + valueToGoStringSnapshotRecoveryResponse(this.Term, "uint64"), `Success:` + valueToGoStringSnapshotRecoveryResponse(this.Success, "bool"), `CommitIndex:` + valueToGoStringSnapshotRecoveryResponse(this.CommitIndex, "uint64"), `XXX_unrecognized:` + fmt26.Sprintf("%#v", this.XXX_unrecognized) + `}`}, ", ")
	return s
}
func valueToGoStringSnapshotRecoveryResponse(v interface{}, typ string) string {
	rv := reflect13.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect13.Indirect(rv).Interface()
	return fmt26.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringSnapshotRecoveryResponse(e map[int32]code_google_com_p_gogoprotobuf_proto13.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort6.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv6.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings13.Join(ss, ",") + "}"
	return s
}
func (this *SnapshotRecoveryResponse) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt27.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*SnapshotRecoveryResponse)
	if !ok {
		return fmt27.Errorf("that is not of type *SnapshotRecoveryResponse")
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt27.Errorf("that is type *SnapshotRecoveryResponse but is nil && this != nil")
	} else if this == nil {
		return fmt27.Errorf("that is type *SnapshotRecoveryResponsebut is not nil && this == nil")
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return fmt27.Errorf("Term this(%v) Not Equal that(%v)", *this.Term, *that1.Term)
		}
	} else if this.Term != nil {
		return fmt27.Errorf("this.Term == nil && that.Term != nil")
	} else if that1.Term != nil {
		return fmt27.Errorf("Term this(%v) Not Equal that(%v)", this.Term, that1.Term)
	}
	if this.Success != nil && that1.Success != nil {
		if *this.Success != *that1.Success {
			return fmt27.Errorf("Success this(%v) Not Equal that(%v)", *this.Success, *that1.Success)
		}
	} else if this.Success != nil {
		return fmt27.Errorf("this.Success == nil && that.Success != nil")
	} else if that1.Success != nil {
		return fmt27.Errorf("Success this(%v) Not Equal that(%v)", this.Success, that1.Success)
	}
	if this.CommitIndex != nil && that1.CommitIndex != nil {
		if *this.CommitIndex != *that1.CommitIndex {
			return fmt27.Errorf("CommitIndex this(%v) Not Equal that(%v)", *this.CommitIndex, *that1.CommitIndex)
		}
	} else if this.CommitIndex != nil {
		return fmt27.Errorf("this.CommitIndex == nil && that.CommitIndex != nil")
	} else if that1.CommitIndex != nil {
		return fmt27.Errorf("CommitIndex this(%v) Not Equal that(%v)", this.CommitIndex, that1.CommitIndex)
	}
	if !bytes6.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return fmt27.Errorf("XXX_unrecognized this(%v) Not Equal that(%v)", this.XXX_unrecognized, that1.XXX_unrecognized)
	}
	return nil
}
func (this *SnapshotRecoveryResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*SnapshotRecoveryResponse)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.Term != nil && that1.Term != nil {
		if *this.Term != *that1.Term {
			return false
		}
	} else if this.Term != nil {
		return false
	} else if that1.Term != nil {
		return false
	}
	if this.Success != nil && that1.Success != nil {
		if *this.Success != *that1.Success {
			return false
		}
	} else if this.Success != nil {
		return false
	} else if that1.Success != nil {
		return false
	}
	if this.CommitIndex != nil && that1.CommitIndex != nil {
		if *this.CommitIndex != *that1.CommitIndex {
			return false
		}
	} else if this.CommitIndex != nil {
		return false
	} else if that1.CommitIndex != nil {
		return false
	}
	if !bytes6.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
