// Do not edit. Generated from pgtype/int.go.erb
package pgtype

import (
	"database/sql/driver"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"github.com/jackc/pgio"
)

type Int64Scanner interface {
	ScanInt64(v int64, valid bool) error
}

type Int2 struct {
	Int   int16
	Valid bool
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *Int2) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = Int2{}
		return nil
	}

	if n < math.MinInt16 {
		return fmt.Errorf("%d is greater than maximum value for Int2", n)
	}
	if n > math.MaxInt16 {
		return fmt.Errorf("%d is greater than maximum value for Int2", n)
	}
	*dst = Int2{Int: int16(n), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int2) Scan(src interface{}) error {
	if src == nil {
		*dst = Int2{}
		return nil
	}

	var n int64

	switch src := src.(type) {
	case int64:
		n = src
	case string:
		var err error
		n, err = strconv.ParseInt(src, 10, 16)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseInt(string(src), 10, 16)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	if n < math.MinInt16 {
		return fmt.Errorf("%d is greater than maximum value for Int2", n)
	}
	if n > math.MaxInt16 {
		return fmt.Errorf("%d is greater than maximum value for Int2", n)
	}
	*dst = Int2{Int: int16(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Int2) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return int64(src.Int), nil
}

func (src Int2) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(src.Int), 10)), nil
}

func (dst *Int2) UnmarshalJSON(b []byte) error {
	var n *int16
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = Int2{}
	} else {
		*dst = Int2{Int: *n, Valid: true}
	}

	return nil
}

type Int2Codec struct{}

func (Int2Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (Int2Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (Int2Codec) Encode(ci *ConnInfo, oid uint32, format int16, value interface{}, buf []byte) (newBuf []byte, err error) {
	n, valid, err := convertToInt64ForEncode(value)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %v to int2: %v", value, err)
	}
	if !valid {
		return nil, nil
	}

	if n > math.MaxInt16 {
		return nil, fmt.Errorf("%d is greater than maximum value for int2", n)
	}
	if n < math.MinInt16 {
		return nil, fmt.Errorf("%d is less than minimum value for int2", n)
	}

	switch format {
	case BinaryFormatCode:
		return pgio.AppendInt16(buf, int16(n)), nil
	case TextFormatCode:
		return append(buf, strconv.FormatInt(n, 10)...), nil
	default:
		return nil, fmt.Errorf("unknown format code: %v", format)
	}
}

func (Int2Codec) PlanScan(ci *ConnInfo, oid uint32, format int16, target interface{}, actualTarget bool) ScanPlan {

	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *int8:
			return scanPlanBinaryInt2ToInt8{}
		case *int16:
			return scanPlanBinaryInt2ToInt16{}
		case *int32:
			return scanPlanBinaryInt2ToInt32{}
		case *int64:
			return scanPlanBinaryInt2ToInt64{}
		case *int:
			return scanPlanBinaryInt2ToInt{}
		case *uint8:
			return scanPlanBinaryInt2ToUint8{}
		case *uint16:
			return scanPlanBinaryInt2ToUint16{}
		case *uint32:
			return scanPlanBinaryInt2ToUint32{}
		case *uint64:
			return scanPlanBinaryInt2ToUint64{}
		case *uint:
			return scanPlanBinaryInt2ToUint{}
		case Int64Scanner:
			return scanPlanBinaryInt2ToInt64Scanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *int8:
			return scanPlanTextAnyToInt8{}
		case *int16:
			return scanPlanTextAnyToInt16{}
		case *int32:
			return scanPlanTextAnyToInt32{}
		case *int64:
			return scanPlanTextAnyToInt64{}
		case *int:
			return scanPlanTextAnyToInt{}
		case *uint8:
			return scanPlanTextAnyToUint8{}
		case *uint16:
			return scanPlanTextAnyToUint16{}
		case *uint32:
			return scanPlanTextAnyToUint32{}
		case *uint64:
			return scanPlanTextAnyToUint64{}
		case *uint:
			return scanPlanTextAnyToUint{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		}
	}

	return nil
}

func (c Int2Codec) DecodeDatabaseSQLValue(ci *ConnInfo, oid uint32, format int16, src []byte) (driver.Value, error) {
	if src == nil {
		return nil, nil
	}

	var n int64
	scanPlan := c.PlanScan(ci, oid, format, &n, true)
	if scanPlan == nil {
		return nil, fmt.Errorf("PlanScan did not find a plan")
	}
	err := scanPlan.Scan(ci, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (c Int2Codec) DecodeValue(ci *ConnInfo, oid uint32, format int16, src []byte) (interface{}, error) {
	if src == nil {
		return nil, nil
	}

	var n int16
	scanPlan := c.PlanScan(ci, oid, format, &n, true)
	if scanPlan == nil {
		return nil, fmt.Errorf("PlanScan did not find a plan")
	}
	err := scanPlan.Scan(ci, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryInt2ToInt8 struct{}

func (scanPlanBinaryInt2ToInt8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for int2: %v", len(src))
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int16(binary.BigEndian.Uint16(src))
	if n < math.MinInt8 {
		return fmt.Errorf("%d is less than minimum value for int8", n)
	} else if n > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for int8", n)
	}

	*p = int8(n)

	return nil
}

type scanPlanBinaryInt2ToUint8 struct{}

func (scanPlanBinaryInt2ToUint8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int16(binary.BigEndian.Uint16(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint8", n)
	}

	if n > math.MaxUint8 {
		return fmt.Errorf("%d is greater than maximum value for uint8", n)
	}

	*p = uint8(n)

	return nil
}

type scanPlanBinaryInt2ToInt16 struct{}

func (scanPlanBinaryInt2ToInt16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for int2: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int16(binary.BigEndian.Uint16(src))

	return nil
}

type scanPlanBinaryInt2ToUint16 struct{}

func (scanPlanBinaryInt2ToUint16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int16(binary.BigEndian.Uint16(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint16", n)
	}

	*p = uint16(n)

	return nil
}

type scanPlanBinaryInt2ToInt32 struct{}

func (scanPlanBinaryInt2ToInt32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for int2: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int32(int16(binary.BigEndian.Uint16(src)))

	return nil
}

type scanPlanBinaryInt2ToUint32 struct{}

func (scanPlanBinaryInt2ToUint32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int16(binary.BigEndian.Uint16(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint32", n)
	}

	*p = uint32(n)

	return nil
}

type scanPlanBinaryInt2ToInt64 struct{}

func (scanPlanBinaryInt2ToInt64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for int2: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int64(int16(binary.BigEndian.Uint16(src)))

	return nil
}

type scanPlanBinaryInt2ToUint64 struct{}

func (scanPlanBinaryInt2ToUint64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int16(binary.BigEndian.Uint16(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint64", n)
	}

	*p = uint64(n)

	return nil
}

type scanPlanBinaryInt2ToInt struct{}

func (scanPlanBinaryInt2ToInt) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for int2: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int(int16(binary.BigEndian.Uint16(src)))

	return nil
}

type scanPlanBinaryInt2ToUint struct{}

func (scanPlanBinaryInt2ToUint) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for uint2: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(int16(binary.BigEndian.Uint16(src)))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint", n)
	}

	*p = uint(n)

	return nil
}

type scanPlanBinaryInt2ToInt64Scanner struct{}

func (scanPlanBinaryInt2ToInt64Scanner) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(0, false)
	}

	if len(src) != 2 {
		return fmt.Errorf("invalid length for int2: %v", len(src))
	}

	n := int64(binary.BigEndian.Uint16(src))

	return s.ScanInt64(n, true)
}

type Int4 struct {
	Int   int32
	Valid bool
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *Int4) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = Int4{}
		return nil
	}

	if n < math.MinInt32 {
		return fmt.Errorf("%d is greater than maximum value for Int4", n)
	}
	if n > math.MaxInt32 {
		return fmt.Errorf("%d is greater than maximum value for Int4", n)
	}
	*dst = Int4{Int: int32(n), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int4) Scan(src interface{}) error {
	if src == nil {
		*dst = Int4{}
		return nil
	}

	var n int64

	switch src := src.(type) {
	case int64:
		n = src
	case string:
		var err error
		n, err = strconv.ParseInt(src, 10, 32)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseInt(string(src), 10, 32)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	if n < math.MinInt32 {
		return fmt.Errorf("%d is greater than maximum value for Int4", n)
	}
	if n > math.MaxInt32 {
		return fmt.Errorf("%d is greater than maximum value for Int4", n)
	}
	*dst = Int4{Int: int32(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Int4) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return int64(src.Int), nil
}

func (src Int4) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(src.Int), 10)), nil
}

func (dst *Int4) UnmarshalJSON(b []byte) error {
	var n *int32
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = Int4{}
	} else {
		*dst = Int4{Int: *n, Valid: true}
	}

	return nil
}

type Int4Codec struct{}

func (Int4Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (Int4Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (Int4Codec) Encode(ci *ConnInfo, oid uint32, format int16, value interface{}, buf []byte) (newBuf []byte, err error) {
	n, valid, err := convertToInt64ForEncode(value)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %v to int4: %v", value, err)
	}
	if !valid {
		return nil, nil
	}

	if n > math.MaxInt32 {
		return nil, fmt.Errorf("%d is greater than maximum value for int4", n)
	}
	if n < math.MinInt32 {
		return nil, fmt.Errorf("%d is less than minimum value for int4", n)
	}

	switch format {
	case BinaryFormatCode:
		return pgio.AppendInt32(buf, int32(n)), nil
	case TextFormatCode:
		return append(buf, strconv.FormatInt(n, 10)...), nil
	default:
		return nil, fmt.Errorf("unknown format code: %v", format)
	}
}

func (Int4Codec) PlanScan(ci *ConnInfo, oid uint32, format int16, target interface{}, actualTarget bool) ScanPlan {

	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *int8:
			return scanPlanBinaryInt4ToInt8{}
		case *int16:
			return scanPlanBinaryInt4ToInt16{}
		case *int32:
			return scanPlanBinaryInt4ToInt32{}
		case *int64:
			return scanPlanBinaryInt4ToInt64{}
		case *int:
			return scanPlanBinaryInt4ToInt{}
		case *uint8:
			return scanPlanBinaryInt4ToUint8{}
		case *uint16:
			return scanPlanBinaryInt4ToUint16{}
		case *uint32:
			return scanPlanBinaryInt4ToUint32{}
		case *uint64:
			return scanPlanBinaryInt4ToUint64{}
		case *uint:
			return scanPlanBinaryInt4ToUint{}
		case Int64Scanner:
			return scanPlanBinaryInt4ToInt64Scanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *int8:
			return scanPlanTextAnyToInt8{}
		case *int16:
			return scanPlanTextAnyToInt16{}
		case *int32:
			return scanPlanTextAnyToInt32{}
		case *int64:
			return scanPlanTextAnyToInt64{}
		case *int:
			return scanPlanTextAnyToInt{}
		case *uint8:
			return scanPlanTextAnyToUint8{}
		case *uint16:
			return scanPlanTextAnyToUint16{}
		case *uint32:
			return scanPlanTextAnyToUint32{}
		case *uint64:
			return scanPlanTextAnyToUint64{}
		case *uint:
			return scanPlanTextAnyToUint{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		}
	}

	return nil
}

func (c Int4Codec) DecodeDatabaseSQLValue(ci *ConnInfo, oid uint32, format int16, src []byte) (driver.Value, error) {
	if src == nil {
		return nil, nil
	}

	var n int64
	scanPlan := c.PlanScan(ci, oid, format, &n, true)
	if scanPlan == nil {
		return nil, fmt.Errorf("PlanScan did not find a plan")
	}
	err := scanPlan.Scan(ci, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (c Int4Codec) DecodeValue(ci *ConnInfo, oid uint32, format int16, src []byte) (interface{}, error) {
	if src == nil {
		return nil, nil
	}

	var n int32
	scanPlan := c.PlanScan(ci, oid, format, &n, true)
	if scanPlan == nil {
		return nil, fmt.Errorf("PlanScan did not find a plan")
	}
	err := scanPlan.Scan(ci, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryInt4ToInt8 struct{}

func (scanPlanBinaryInt4ToInt8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for int4: %v", len(src))
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int32(binary.BigEndian.Uint32(src))
	if n < math.MinInt8 {
		return fmt.Errorf("%d is less than minimum value for int8", n)
	} else if n > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for int8", n)
	}

	*p = int8(n)

	return nil
}

type scanPlanBinaryInt4ToUint8 struct{}

func (scanPlanBinaryInt4ToUint8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int32(binary.BigEndian.Uint32(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint8", n)
	}

	if n > math.MaxUint8 {
		return fmt.Errorf("%d is greater than maximum value for uint8", n)
	}

	*p = uint8(n)

	return nil
}

type scanPlanBinaryInt4ToInt16 struct{}

func (scanPlanBinaryInt4ToInt16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for int4: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int32(binary.BigEndian.Uint32(src))
	if n < math.MinInt16 {
		return fmt.Errorf("%d is less than minimum value for int16", n)
	} else if n > math.MaxInt16 {
		return fmt.Errorf("%d is greater than maximum value for int16", n)
	}

	*p = int16(n)

	return nil
}

type scanPlanBinaryInt4ToUint16 struct{}

func (scanPlanBinaryInt4ToUint16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int32(binary.BigEndian.Uint32(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint16", n)
	}

	if n > math.MaxUint16 {
		return fmt.Errorf("%d is greater than maximum value for uint16", n)
	}

	*p = uint16(n)

	return nil
}

type scanPlanBinaryInt4ToInt32 struct{}

func (scanPlanBinaryInt4ToInt32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for int4: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int32(binary.BigEndian.Uint32(src))

	return nil
}

type scanPlanBinaryInt4ToUint32 struct{}

func (scanPlanBinaryInt4ToUint32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int32(binary.BigEndian.Uint32(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint32", n)
	}

	*p = uint32(n)

	return nil
}

type scanPlanBinaryInt4ToInt64 struct{}

func (scanPlanBinaryInt4ToInt64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for int4: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int64(int32(binary.BigEndian.Uint32(src)))

	return nil
}

type scanPlanBinaryInt4ToUint64 struct{}

func (scanPlanBinaryInt4ToUint64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int32(binary.BigEndian.Uint32(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint64", n)
	}

	*p = uint64(n)

	return nil
}

type scanPlanBinaryInt4ToInt struct{}

func (scanPlanBinaryInt4ToInt) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for int4: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int(int32(binary.BigEndian.Uint32(src)))

	return nil
}

type scanPlanBinaryInt4ToUint struct{}

func (scanPlanBinaryInt4ToUint) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for uint4: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(int32(binary.BigEndian.Uint32(src)))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint", n)
	}

	*p = uint(n)

	return nil
}

type scanPlanBinaryInt4ToInt64Scanner struct{}

func (scanPlanBinaryInt4ToInt64Scanner) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(0, false)
	}

	if len(src) != 4 {
		return fmt.Errorf("invalid length for int4: %v", len(src))
	}

	n := int64(binary.BigEndian.Uint32(src))

	return s.ScanInt64(n, true)
}

type Int8 struct {
	Int   int64
	Valid bool
}

// ScanInt64 implements the Int64Scanner interface.
func (dst *Int8) ScanInt64(n int64, valid bool) error {
	if !valid {
		*dst = Int8{}
		return nil
	}

	if n < math.MinInt64 {
		return fmt.Errorf("%d is greater than maximum value for Int8", n)
	}
	if n > math.MaxInt64 {
		return fmt.Errorf("%d is greater than maximum value for Int8", n)
	}
	*dst = Int8{Int: int64(n), Valid: true}

	return nil
}

// Scan implements the database/sql Scanner interface.
func (dst *Int8) Scan(src interface{}) error {
	if src == nil {
		*dst = Int8{}
		return nil
	}

	var n int64

	switch src := src.(type) {
	case int64:
		n = src
	case string:
		var err error
		n, err = strconv.ParseInt(src, 10, 64)
		if err != nil {
			return err
		}
	case []byte:
		var err error
		n, err = strconv.ParseInt(string(src), 10, 64)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot scan %T", src)
	}

	if n < math.MinInt64 {
		return fmt.Errorf("%d is greater than maximum value for Int8", n)
	}
	if n > math.MaxInt64 {
		return fmt.Errorf("%d is greater than maximum value for Int8", n)
	}
	*dst = Int8{Int: int64(n), Valid: true}

	return nil
}

// Value implements the database/sql/driver Valuer interface.
func (src Int8) Value() (driver.Value, error) {
	if !src.Valid {
		return nil, nil
	}
	return int64(src.Int), nil
}

func (src Int8) MarshalJSON() ([]byte, error) {
	if !src.Valid {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatInt(int64(src.Int), 10)), nil
}

func (dst *Int8) UnmarshalJSON(b []byte) error {
	var n *int64
	err := json.Unmarshal(b, &n)
	if err != nil {
		return err
	}

	if n == nil {
		*dst = Int8{}
	} else {
		*dst = Int8{Int: *n, Valid: true}
	}

	return nil
}

type Int8Codec struct{}

func (Int8Codec) FormatSupported(format int16) bool {
	return format == TextFormatCode || format == BinaryFormatCode
}

func (Int8Codec) PreferredFormat() int16 {
	return BinaryFormatCode
}

func (Int8Codec) Encode(ci *ConnInfo, oid uint32, format int16, value interface{}, buf []byte) (newBuf []byte, err error) {
	n, valid, err := convertToInt64ForEncode(value)
	if err != nil {
		return nil, fmt.Errorf("cannot convert %v to int8: %v", value, err)
	}
	if !valid {
		return nil, nil
	}

	if n > math.MaxInt64 {
		return nil, fmt.Errorf("%d is greater than maximum value for int8", n)
	}
	if n < math.MinInt64 {
		return nil, fmt.Errorf("%d is less than minimum value for int8", n)
	}

	switch format {
	case BinaryFormatCode:
		return pgio.AppendInt64(buf, int64(n)), nil
	case TextFormatCode:
		return append(buf, strconv.FormatInt(n, 10)...), nil
	default:
		return nil, fmt.Errorf("unknown format code: %v", format)
	}
}

func (Int8Codec) PlanScan(ci *ConnInfo, oid uint32, format int16, target interface{}, actualTarget bool) ScanPlan {

	switch format {
	case BinaryFormatCode:
		switch target.(type) {
		case *int8:
			return scanPlanBinaryInt8ToInt8{}
		case *int16:
			return scanPlanBinaryInt8ToInt16{}
		case *int32:
			return scanPlanBinaryInt8ToInt32{}
		case *int64:
			return scanPlanBinaryInt8ToInt64{}
		case *int:
			return scanPlanBinaryInt8ToInt{}
		case *uint8:
			return scanPlanBinaryInt8ToUint8{}
		case *uint16:
			return scanPlanBinaryInt8ToUint16{}
		case *uint32:
			return scanPlanBinaryInt8ToUint32{}
		case *uint64:
			return scanPlanBinaryInt8ToUint64{}
		case *uint:
			return scanPlanBinaryInt8ToUint{}
		case Int64Scanner:
			return scanPlanBinaryInt8ToInt64Scanner{}
		}
	case TextFormatCode:
		switch target.(type) {
		case *int8:
			return scanPlanTextAnyToInt8{}
		case *int16:
			return scanPlanTextAnyToInt16{}
		case *int32:
			return scanPlanTextAnyToInt32{}
		case *int64:
			return scanPlanTextAnyToInt64{}
		case *int:
			return scanPlanTextAnyToInt{}
		case *uint8:
			return scanPlanTextAnyToUint8{}
		case *uint16:
			return scanPlanTextAnyToUint16{}
		case *uint32:
			return scanPlanTextAnyToUint32{}
		case *uint64:
			return scanPlanTextAnyToUint64{}
		case *uint:
			return scanPlanTextAnyToUint{}
		case Int64Scanner:
			return scanPlanTextAnyToInt64Scanner{}
		}
	}

	return nil
}

func (c Int8Codec) DecodeDatabaseSQLValue(ci *ConnInfo, oid uint32, format int16, src []byte) (driver.Value, error) {
	if src == nil {
		return nil, nil
	}

	var n int64
	scanPlan := c.PlanScan(ci, oid, format, &n, true)
	if scanPlan == nil {
		return nil, fmt.Errorf("PlanScan did not find a plan")
	}
	err := scanPlan.Scan(ci, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

func (c Int8Codec) DecodeValue(ci *ConnInfo, oid uint32, format int16, src []byte) (interface{}, error) {
	if src == nil {
		return nil, nil
	}

	var n int64
	scanPlan := c.PlanScan(ci, oid, format, &n, true)
	if scanPlan == nil {
		return nil, fmt.Errorf("PlanScan did not find a plan")
	}
	err := scanPlan.Scan(ci, oid, format, src, &n)
	if err != nil {
		return nil, err
	}
	return n, nil
}

type scanPlanBinaryInt8ToInt8 struct{}

func (scanPlanBinaryInt8ToInt8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for int8: %v", len(src))
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < math.MinInt8 {
		return fmt.Errorf("%d is less than minimum value for int8", n)
	} else if n > math.MaxInt8 {
		return fmt.Errorf("%d is greater than maximum value for int8", n)
	}

	*p = int8(n)

	return nil
}

type scanPlanBinaryInt8ToUint8 struct{}

func (scanPlanBinaryInt8ToUint8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint8", n)
	}

	if n > math.MaxUint8 {
		return fmt.Errorf("%d is greater than maximum value for uint8", n)
	}

	*p = uint8(n)

	return nil
}

type scanPlanBinaryInt8ToInt16 struct{}

func (scanPlanBinaryInt8ToInt16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for int8: %v", len(src))
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < math.MinInt16 {
		return fmt.Errorf("%d is less than minimum value for int16", n)
	} else if n > math.MaxInt16 {
		return fmt.Errorf("%d is greater than maximum value for int16", n)
	}

	*p = int16(n)

	return nil
}

type scanPlanBinaryInt8ToUint16 struct{}

func (scanPlanBinaryInt8ToUint16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint16", n)
	}

	if n > math.MaxUint16 {
		return fmt.Errorf("%d is greater than maximum value for uint16", n)
	}

	*p = uint16(n)

	return nil
}

type scanPlanBinaryInt8ToInt32 struct{}

func (scanPlanBinaryInt8ToInt32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for int8: %v", len(src))
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < math.MinInt32 {
		return fmt.Errorf("%d is less than minimum value for int32", n)
	} else if n > math.MaxInt32 {
		return fmt.Errorf("%d is greater than maximum value for int32", n)
	}

	*p = int32(n)

	return nil
}

type scanPlanBinaryInt8ToUint32 struct{}

func (scanPlanBinaryInt8ToUint32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint32", n)
	}

	if n > math.MaxUint32 {
		return fmt.Errorf("%d is greater than maximum value for uint32", n)
	}

	*p = uint32(n)

	return nil
}

type scanPlanBinaryInt8ToInt64 struct{}

func (scanPlanBinaryInt8ToInt64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for int8: %v", len(src))
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	*p = int64(binary.BigEndian.Uint64(src))

	return nil
}

type scanPlanBinaryInt8ToUint64 struct{}

func (scanPlanBinaryInt8ToUint64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint64", n)
	}

	*p = uint64(n)

	return nil
}

type scanPlanBinaryInt8ToInt struct{}

func (scanPlanBinaryInt8ToInt) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for int8: %v", len(src))
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(binary.BigEndian.Uint64(src))
	if n < math.MinInt {
		return fmt.Errorf("%d is less than minimum value for int", n)
	} else if n > math.MaxInt {
		return fmt.Errorf("%d is greater than maximum value for int", n)
	}

	*p = int(n)

	return nil
}

type scanPlanBinaryInt8ToUint struct{}

func (scanPlanBinaryInt8ToUint) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for uint8: %v", len(src))
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n := int64(int64(binary.BigEndian.Uint64(src)))
	if n < 0 {
		return fmt.Errorf("%d is less than minimum value for uint", n)
	}

	if uint64(n) > math.MaxUint {
		return fmt.Errorf("%d is greater than maximum value for uint", n)
	}

	*p = uint(n)

	return nil
}

type scanPlanBinaryInt8ToInt64Scanner struct{}

func (scanPlanBinaryInt8ToInt64Scanner) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(0, false)
	}

	if len(src) != 8 {
		return fmt.Errorf("invalid length for int8: %v", len(src))
	}

	n := int64(binary.BigEndian.Uint64(src))

	return s.ScanInt64(n, true)
}

type scanPlanTextAnyToInt8 struct{}

func (scanPlanTextAnyToInt8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*int8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 8)
	if err != nil {
		return err
	}

	*p = int8(n)
	return nil
}

type scanPlanTextAnyToUint8 struct{}

func (scanPlanTextAnyToUint8) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*uint8)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 8)
	if err != nil {
		return err
	}

	*p = uint8(n)
	return nil
}

type scanPlanTextAnyToInt16 struct{}

func (scanPlanTextAnyToInt16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*int16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 16)
	if err != nil {
		return err
	}

	*p = int16(n)
	return nil
}

type scanPlanTextAnyToUint16 struct{}

func (scanPlanTextAnyToUint16) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*uint16)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 16)
	if err != nil {
		return err
	}

	*p = uint16(n)
	return nil
}

type scanPlanTextAnyToInt32 struct{}

func (scanPlanTextAnyToInt32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*int32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 32)
	if err != nil {
		return err
	}

	*p = int32(n)
	return nil
}

type scanPlanTextAnyToUint32 struct{}

func (scanPlanTextAnyToUint32) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*uint32)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 32)
	if err != nil {
		return err
	}

	*p = uint32(n)
	return nil
}

type scanPlanTextAnyToInt64 struct{}

func (scanPlanTextAnyToInt64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*int64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 64)
	if err != nil {
		return err
	}

	*p = int64(n)
	return nil
}

type scanPlanTextAnyToUint64 struct{}

func (scanPlanTextAnyToUint64) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*uint64)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 64)
	if err != nil {
		return err
	}

	*p = uint64(n)
	return nil
}

type scanPlanTextAnyToInt struct{}

func (scanPlanTextAnyToInt) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*int)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseInt(string(src), 10, 0)
	if err != nil {
		return err
	}

	*p = int(n)
	return nil
}

type scanPlanTextAnyToUint struct{}

func (scanPlanTextAnyToUint) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	if src == nil {
		return fmt.Errorf("cannot scan null into %T", dst)
	}

	p, ok := (dst).(*uint)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	n, err := strconv.ParseUint(string(src), 10, 0)
	if err != nil {
		return err
	}

	*p = uint(n)
	return nil
}

type scanPlanTextAnyToInt64Scanner struct{}

func (scanPlanTextAnyToInt64Scanner) Scan(ci *ConnInfo, oid uint32, formatCode int16, src []byte, dst interface{}) error {
	s, ok := (dst).(Int64Scanner)
	if !ok {
		return ErrScanTargetTypeChanged
	}

	if src == nil {
		return s.ScanInt64(0, false)
	}

	n, err := strconv.ParseInt(string(src), 10, 64)
	if err != nil {
		return err
	}

	err = s.ScanInt64(n, true)
	if err != nil {
		return err
	}

	return nil
}