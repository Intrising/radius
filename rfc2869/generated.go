// Code generated by radius-dict-gen. DO NOT EDIT.

package rfc2869

import (
	"strconv"
	"time"

	"github.com/Intrising/radius"
)

const (
	AcctInputGigawords_Type   radius.Type = 52
	AcctOutputGigawords_Type  radius.Type = 53
	EventTimestamp_Type       radius.Type = 55
	ARAPZoneAccess_Type       radius.Type = 72
	ARAPSecurity_Type         radius.Type = 73
	ARAPSecurityData_Type     radius.Type = 74
	PasswordRetry_Type        radius.Type = 75
	Prompt_Type               radius.Type = 76
	ConnectInfo_Type          radius.Type = 77
	ConfigurationToken_Type   radius.Type = 78
	EAPMessage_Type           radius.Type = 79
	MessageAuthenticator_Type radius.Type = 80
	AcctInterimInterval_Type  radius.Type = 85
	NASPortID_Type            radius.Type = 87
	FramedPool_Type           radius.Type = 88
)

type AcctInputGigawords uint32

var AcctInputGigawords_Strings = map[AcctInputGigawords]string{}

func (a AcctInputGigawords) String() string {
	if str, ok := AcctInputGigawords_Strings[a]; ok {
		return str
	}
	return "AcctInputGigawords(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctInputGigawords_Add(p *radius.Packet, value AcctInputGigawords) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctInputGigawords_Type, a)
	return nil
}

func AcctInputGigawords_Get(p *radius.Packet) (value AcctInputGigawords) {
	value, _ = AcctInputGigawords_Lookup(p)
	return
}

func AcctInputGigawords_Gets(p *radius.Packet) (values []AcctInputGigawords, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctInputGigawords_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctInputGigawords(i))
	}
	return
}

func AcctInputGigawords_Lookup(p *radius.Packet) (value AcctInputGigawords, err error) {
	a, ok := p.Lookup(AcctInputGigawords_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctInputGigawords(i)
	return
}

func AcctInputGigawords_Set(p *radius.Packet, value AcctInputGigawords) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctInputGigawords_Type, a)
	return nil
}

type AcctOutputGigawords uint32

var AcctOutputGigawords_Strings = map[AcctOutputGigawords]string{}

func (a AcctOutputGigawords) String() string {
	if str, ok := AcctOutputGigawords_Strings[a]; ok {
		return str
	}
	return "AcctOutputGigawords(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctOutputGigawords_Add(p *radius.Packet, value AcctOutputGigawords) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctOutputGigawords_Type, a)
	return nil
}

func AcctOutputGigawords_Get(p *radius.Packet) (value AcctOutputGigawords) {
	value, _ = AcctOutputGigawords_Lookup(p)
	return
}

func AcctOutputGigawords_Gets(p *radius.Packet) (values []AcctOutputGigawords, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctOutputGigawords_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctOutputGigawords(i))
	}
	return
}

func AcctOutputGigawords_Lookup(p *radius.Packet) (value AcctOutputGigawords, err error) {
	a, ok := p.Lookup(AcctOutputGigawords_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctOutputGigawords(i)
	return
}

func AcctOutputGigawords_Set(p *radius.Packet, value AcctOutputGigawords) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctOutputGigawords_Type, a)
	return nil
}

func EventTimestamp_Add(p *radius.Packet, value time.Time) (err error) {
	var a radius.Attribute
	a, err = radius.NewDate(value)
	if err != nil {
		return
	}
	p.Add(EventTimestamp_Type, a)
	return nil
}

func EventTimestamp_Get(p *radius.Packet) (value time.Time) {
	value, _ = EventTimestamp_Lookup(p)
	return
}

func EventTimestamp_Gets(p *radius.Packet) (values []time.Time, err error) {
	var i time.Time
	for _, attr := range p.Attributes[EventTimestamp_Type] {
		i, err = radius.Date(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EventTimestamp_Lookup(p *radius.Packet) (value time.Time, err error) {
	a, ok := p.Lookup(EventTimestamp_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value, err = radius.Date(a)
	return
}

func EventTimestamp_Set(p *radius.Packet, value time.Time) (err error) {
	var a radius.Attribute
	a, err = radius.NewDate(value)
	if err != nil {
		return
	}
	p.Set(EventTimestamp_Type, a)
	return nil
}

type ARAPZoneAccess uint32

const (
	ARAPZoneAccess_Value_DefaultZone         ARAPZoneAccess = 1
	ARAPZoneAccess_Value_ZoneFilterInclusive ARAPZoneAccess = 2
	ARAPZoneAccess_Value_ZoneFilterExclusive ARAPZoneAccess = 4
)

var ARAPZoneAccess_Strings = map[ARAPZoneAccess]string{
	ARAPZoneAccess_Value_DefaultZone:         "Default-Zone",
	ARAPZoneAccess_Value_ZoneFilterInclusive: "Zone-Filter-Inclusive",
	ARAPZoneAccess_Value_ZoneFilterExclusive: "Zone-Filter-Exclusive",
}

func (a ARAPZoneAccess) String() string {
	if str, ok := ARAPZoneAccess_Strings[a]; ok {
		return str
	}
	return "ARAPZoneAccess(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ARAPZoneAccess_Add(p *radius.Packet, value ARAPZoneAccess) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(ARAPZoneAccess_Type, a)
	return nil
}

func ARAPZoneAccess_Get(p *radius.Packet) (value ARAPZoneAccess) {
	value, _ = ARAPZoneAccess_Lookup(p)
	return
}

func ARAPZoneAccess_Gets(p *radius.Packet) (values []ARAPZoneAccess, err error) {
	var i uint32
	for _, attr := range p.Attributes[ARAPZoneAccess_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ARAPZoneAccess(i))
	}
	return
}

func ARAPZoneAccess_Lookup(p *radius.Packet) (value ARAPZoneAccess, err error) {
	a, ok := p.Lookup(ARAPZoneAccess_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ARAPZoneAccess(i)
	return
}

func ARAPZoneAccess_Set(p *radius.Packet, value ARAPZoneAccess) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(ARAPZoneAccess_Type, a)
	return nil
}

type ARAPSecurity uint32

var ARAPSecurity_Strings = map[ARAPSecurity]string{}

func (a ARAPSecurity) String() string {
	if str, ok := ARAPSecurity_Strings[a]; ok {
		return str
	}
	return "ARAPSecurity(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func ARAPSecurity_Add(p *radius.Packet, value ARAPSecurity) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(ARAPSecurity_Type, a)
	return nil
}

func ARAPSecurity_Get(p *radius.Packet) (value ARAPSecurity) {
	value, _ = ARAPSecurity_Lookup(p)
	return
}

func ARAPSecurity_Gets(p *radius.Packet) (values []ARAPSecurity, err error) {
	var i uint32
	for _, attr := range p.Attributes[ARAPSecurity_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, ARAPSecurity(i))
	}
	return
}

func ARAPSecurity_Lookup(p *radius.Packet) (value ARAPSecurity, err error) {
	a, ok := p.Lookup(ARAPSecurity_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = ARAPSecurity(i)
	return
}

func ARAPSecurity_Set(p *radius.Packet, value ARAPSecurity) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(ARAPSecurity_Type, a)
	return nil
}

func ARAPSecurityData_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(ARAPSecurityData_Type, a)
	return nil
}

func ARAPSecurityData_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(ARAPSecurityData_Type, a)
	return nil
}

func ARAPSecurityData_Get(p *radius.Packet) (value []byte) {
	value, _ = ARAPSecurityData_Lookup(p)
	return
}

func ARAPSecurityData_GetString(p *radius.Packet) (value string) {
	return string(ARAPSecurityData_Get(p))
}

func ARAPSecurityData_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[ARAPSecurityData_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ARAPSecurityData_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[ARAPSecurityData_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ARAPSecurityData_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(ARAPSecurityData_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ARAPSecurityData_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(ARAPSecurityData_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ARAPSecurityData_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(ARAPSecurityData_Type, a)
	return
}

func ARAPSecurityData_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(ARAPSecurityData_Type, a)
	return
}

type PasswordRetry uint32

var PasswordRetry_Strings = map[PasswordRetry]string{}

func (a PasswordRetry) String() string {
	if str, ok := PasswordRetry_Strings[a]; ok {
		return str
	}
	return "PasswordRetry(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func PasswordRetry_Add(p *radius.Packet, value PasswordRetry) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(PasswordRetry_Type, a)
	return nil
}

func PasswordRetry_Get(p *radius.Packet) (value PasswordRetry) {
	value, _ = PasswordRetry_Lookup(p)
	return
}

func PasswordRetry_Gets(p *radius.Packet) (values []PasswordRetry, err error) {
	var i uint32
	for _, attr := range p.Attributes[PasswordRetry_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, PasswordRetry(i))
	}
	return
}

func PasswordRetry_Lookup(p *radius.Packet) (value PasswordRetry, err error) {
	a, ok := p.Lookup(PasswordRetry_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = PasswordRetry(i)
	return
}

func PasswordRetry_Set(p *radius.Packet, value PasswordRetry) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(PasswordRetry_Type, a)
	return nil
}

type Prompt uint32

const (
	Prompt_Value_NoEcho Prompt = 0
	Prompt_Value_Echo   Prompt = 1
)

var Prompt_Strings = map[Prompt]string{
	Prompt_Value_NoEcho: "No-Echo",
	Prompt_Value_Echo:   "Echo",
}

func (a Prompt) String() string {
	if str, ok := Prompt_Strings[a]; ok {
		return str
	}
	return "Prompt(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func Prompt_Add(p *radius.Packet, value Prompt) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(Prompt_Type, a)
	return nil
}

func Prompt_Get(p *radius.Packet) (value Prompt) {
	value, _ = Prompt_Lookup(p)
	return
}

func Prompt_Gets(p *radius.Packet) (values []Prompt, err error) {
	var i uint32
	for _, attr := range p.Attributes[Prompt_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, Prompt(i))
	}
	return
}

func Prompt_Lookup(p *radius.Packet) (value Prompt, err error) {
	a, ok := p.Lookup(Prompt_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = Prompt(i)
	return
}

func Prompt_Set(p *radius.Packet, value Prompt) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(Prompt_Type, a)
	return nil
}

func ConnectInfo_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(ConnectInfo_Type, a)
	return nil
}

func ConnectInfo_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(ConnectInfo_Type, a)
	return nil
}

func ConnectInfo_Get(p *radius.Packet) (value []byte) {
	value, _ = ConnectInfo_Lookup(p)
	return
}

func ConnectInfo_GetString(p *radius.Packet) (value string) {
	return string(ConnectInfo_Get(p))
}

func ConnectInfo_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[ConnectInfo_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ConnectInfo_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[ConnectInfo_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ConnectInfo_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(ConnectInfo_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ConnectInfo_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(ConnectInfo_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ConnectInfo_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(ConnectInfo_Type, a)
	return
}

func ConnectInfo_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(ConnectInfo_Type, a)
	return
}

func ConfigurationToken_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(ConfigurationToken_Type, a)
	return nil
}

func ConfigurationToken_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(ConfigurationToken_Type, a)
	return nil
}

func ConfigurationToken_Get(p *radius.Packet) (value []byte) {
	value, _ = ConfigurationToken_Lookup(p)
	return
}

func ConfigurationToken_GetString(p *radius.Packet) (value string) {
	return string(ConfigurationToken_Get(p))
}

func ConfigurationToken_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[ConfigurationToken_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ConfigurationToken_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[ConfigurationToken_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func ConfigurationToken_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(ConfigurationToken_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func ConfigurationToken_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(ConfigurationToken_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func ConfigurationToken_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(ConfigurationToken_Type, a)
	return
}

func ConfigurationToken_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(ConfigurationToken_Type, a)
	return
}

func EAPMessage_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(EAPMessage_Type, a)
	return nil
}

func EAPMessage_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(EAPMessage_Type, a)
	return nil
}

func EAPMessage_Get(p *radius.Packet) (value []byte) {
	value, _ = EAPMessage_Lookup(p)
	return
}

func EAPMessage_GetString(p *radius.Packet) (value string) {
	return string(EAPMessage_Get(p))
}

func EAPMessage_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[EAPMessage_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPMessage_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[EAPMessage_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func EAPMessage_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(EAPMessage_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func EAPMessage_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(EAPMessage_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func EAPMessage_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(EAPMessage_Type, a)
	return
}

func EAPMessage_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(EAPMessage_Type, a)
	return
}

func MessageAuthenticator_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(MessageAuthenticator_Type, a)
	return nil
}

func MessageAuthenticator_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(MessageAuthenticator_Type, a)
	return nil
}

func MessageAuthenticator_Get(p *radius.Packet) (value []byte) {
	value, _ = MessageAuthenticator_Lookup(p)
	return
}

func MessageAuthenticator_GetString(p *radius.Packet) (value string) {
	return string(MessageAuthenticator_Get(p))
}

func MessageAuthenticator_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[MessageAuthenticator_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MessageAuthenticator_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[MessageAuthenticator_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func MessageAuthenticator_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(MessageAuthenticator_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func MessageAuthenticator_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(MessageAuthenticator_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func MessageAuthenticator_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(MessageAuthenticator_Type, a)
	return
}

func MessageAuthenticator_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(MessageAuthenticator_Type, a)
	return
}

type AcctInterimInterval uint32

var AcctInterimInterval_Strings = map[AcctInterimInterval]string{}

func (a AcctInterimInterval) String() string {
	if str, ok := AcctInterimInterval_Strings[a]; ok {
		return str
	}
	return "AcctInterimInterval(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func AcctInterimInterval_Add(p *radius.Packet, value AcctInterimInterval) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Add(AcctInterimInterval_Type, a)
	return nil
}

func AcctInterimInterval_Get(p *radius.Packet) (value AcctInterimInterval) {
	value, _ = AcctInterimInterval_Lookup(p)
	return
}

func AcctInterimInterval_Gets(p *radius.Packet) (values []AcctInterimInterval, err error) {
	var i uint32
	for _, attr := range p.Attributes[AcctInterimInterval_Type] {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, AcctInterimInterval(i))
	}
	return
}

func AcctInterimInterval_Lookup(p *radius.Packet) (value AcctInterimInterval, err error) {
	a, ok := p.Lookup(AcctInterimInterval_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = AcctInterimInterval(i)
	return
}

func AcctInterimInterval_Set(p *radius.Packet, value AcctInterimInterval) (err error) {
	a := radius.NewInteger(uint32(value))
	p.Set(AcctInterimInterval_Type, a)
	return nil
}

func NASPortID_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(NASPortID_Type, a)
	return nil
}

func NASPortID_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(NASPortID_Type, a)
	return nil
}

func NASPortID_Get(p *radius.Packet) (value []byte) {
	value, _ = NASPortID_Lookup(p)
	return
}

func NASPortID_GetString(p *radius.Packet) (value string) {
	return string(NASPortID_Get(p))
}

func NASPortID_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[NASPortID_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NASPortID_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[NASPortID_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func NASPortID_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(NASPortID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func NASPortID_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(NASPortID_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func NASPortID_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(NASPortID_Type, a)
	return
}

func NASPortID_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(NASPortID_Type, a)
	return
}

func FramedPool_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Add(FramedPool_Type, a)
	return nil
}

func FramedPool_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Add(FramedPool_Type, a)
	return nil
}

func FramedPool_Get(p *radius.Packet) (value []byte) {
	value, _ = FramedPool_Lookup(p)
	return
}

func FramedPool_GetString(p *radius.Packet) (value string) {
	return string(FramedPool_Get(p))
}

func FramedPool_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range p.Attributes[FramedPool_Type] {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FramedPool_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range p.Attributes[FramedPool_Type] {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func FramedPool_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := p.Lookup(FramedPool_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func FramedPool_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := p.Lookup(FramedPool_Type)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func FramedPool_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	p.Set(FramedPool_Type, a)
	return
}

func FramedPool_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	p.Set(FramedPool_Type, a)
	return
}
