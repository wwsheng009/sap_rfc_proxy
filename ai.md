
## ai prompt for converting the json to sap rfc struct

try to convert the input map to the correct type
the input is interface{},it's converted from the json object

```go
type FieldDescription struct {
	Name      string
	FieldType string
	NucLength uint
	NucOffset uint
	UcLength  uint
	UcOffset  uint
	Decimals  uint
	TypeDesc  TypeDescription
}

type TypeDescription struct {
	Name      string
	NucLength uint
	UcLength  uint
	Fields    []FieldDescription
}

// ParameterDescription type
type ParameterDescription struct {
	Name          string
	ParameterType string    //RFCTYPE_STRUCTURE,RFCTYPE_TABLE
	Direction     string    //"RFC_EXPORT" "RFC_IMPORT" "RFC_TABLES" "RFC_CHANGING"
	NucLength     uint
	UcLength      uint
	Decimals      uint
	DefaultValue  string
	ParameterText string
	Optional      bool
	TypeDesc      TypeDescription
}

// FunctionDescription type
type FunctionDescription struct {
	Name       string
	Parameters []ParameterDescription
}

func convertMapObject(input interface{},funDesc gorfc.FunctionDescription)(output interface{},err error){
    //check and convert the object

    //if the ParameterType of the paramter is RFCTYPE_TABLE, type meta fields is not empty, the Corresponding object is the array of the map object in json,need to check and convert the input array. 

    //if the ParameterType of the paramter is RFCTYPE_STRUCTURE, type meta fields is not empty, the Corresponding object is the map object in json,need to check and convert the input map.

    //if the ParameterType of the paramter is RFCTYPE_STRING or RFCTYPE_CHAR, the Corresponding object is the string in json,need to check and convert the input string.

    //     typedef enum _RFCTYPE
    // {
    //     RFCTYPE_CHAR   = 0,		///< 1-byte or multibyte character, fixed size, blank padded
    //     RFCTYPE_DATE   = 1,		///< Date ( YYYYYMMDD )
    //     RFCTYPE_BCD    = 2,		///< Packed number, any length between 1 and 16 bytes
    //     RFCTYPE_TIME   = 3,		///< Time (HHMMSS) 
    //     RFCTYPE_BYTE   = 4,		///< Raw data, binary, fixed length, zero padded.
    //     RFCTYPE_TABLE   = 5,	///< Internal table
    //     RFCTYPE_NUM    = 6,		///< Digits, fixed size, leading '0' padded.
    //     RFCTYPE_FLOAT  = 7,		///< Floating point, double precision
    //     RFCTYPE_INT    = 8,		///< 4-byte integer
    //     RFCTYPE_INT2   = 9,		///< 2-byte integer. Obsolete, not directly supported by ABAP/4
    //     RFCTYPE_INT1   = 10,	///< 1-byte integer, unsigned. Obsolete, not directly supported by ABAP/4
    //     RFCTYPE_NULL  = 14,		///< Not supported data type.
    //     RFCTYPE_ABAPOBJECT = 16,///< ABAP object.
    //     RFCTYPE_STRUCTURE = 17,	///< ABAP structure
    //     RFCTYPE_DECF16  = 23,	///< IEEE 754r decimal floating point, 8 bytes
    //     RFCTYPE_DECF34  = 24,	///< IEEE 754r decimal floating point, 16 bytes
    //     RFCTYPE_XMLDATA = 28,	///< No longer used!
    //     RFCTYPE_STRING = 29,	///< Variable-length, null-terminated string
    //     RFCTYPE_XSTRING = 30,	///< Variable-length raw string, length in bytes
    //     RFCTYPE_INT8,			///< 8-byte integer
    //     RFCTYPE_UTCLONG ,		///< timestamp/long, 8-byte integer
    //     RFCTYPE_UTCSECOND ,		///< timestamp/second, 8-byte integer
    //     RFCTYPE_UTCMINUTE ,		///< timestamp/minute, 8-byte integer
    //     RFCTYPE_DTDAY ,			///< date/day , 4-byte integer
    //     RFCTYPE_DTWEEK   ,		///< date/week, 4-byte integer
    //     RFCTYPE_DTMONTH  ,		///< date/month, 4-byte integer
    //     RFCTYPE_TSECOND  ,		///< time/second, 4-byte integer
    //     RFCTYPE_TMINUTE  ,		///< time/minute, 2-byte integer
    //     RFCTYPE_CDAY  ,			///< calendar day, 2-byte integer
    //     RFCTYPE_BOX  ,			///< boxed structure, note: not supported by NW RFC lib
    //     RFCTYPE_GENERIC_BOX,	///< boxed client dependent structure, note: not supported by NW RFC lib
    //     _RFCTYPE_max_value		///< the max. value of RFCTYPEs
    // }RFCTYPE;

	return input,nil;
}
```