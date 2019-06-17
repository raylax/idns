package protocol

type Name string

const Unknown = "UNKNOWN"

type valueName struct {
	Value int
	Name  string
}

type Type valueName

func ParseType(value int) Type {
	for _, t := range Types {
		if t.Value == value {
			return t
		}
	}
	return Type{Value: value, Name: Unknown}
}

type Class valueName

func ParseClass(value int) Class {
	for _, t := range Classes {
		if t.Value == value {
			return t
		}
	}
	return Class{Value: value, Name: Unknown}
}

type QR valueName

func ParseRCode(value int) RCode {
	for _, r := range RCodes {
		if r.Value == value {
			return r
		}
	}
	return RCode{Value: value, Name: Unknown}
}

var (
	QRQuery    QR = QR{Value: 0x00, Name: "Query"}
	QRResponse QR = QR{Value: 0x01, Name: "Response"}
)

type OpCode valueName

func ParseOpCode(value int) OpCode {
	for _, o := range OpCodes {
		if o.Value == value {
			return o
		}
	}
	return OpCode{Value: value, Name: Unknown}
}

var OpCodes = []OpCode{
	OpCodeQuery,
	OpCodeIQuery,
	OpCodeStatus,
	OpCodeNotify,
	OpCodeUpdate,
}

var (
	OpCodeQuery  = OpCode{Value: 0x0000, Name: "Query"}
	OpCodeIQuery = OpCode{Value: 0x0001, Name: "IQuery"}
	OpCodeStatus = OpCode{Value: 0x0002, Name: "Status"}
	OpCodeNotify = OpCode{Value: 0x0004, Name: "Notify"}
	OpCodeUpdate = OpCode{Value: 0x0005, Name: "Update"}
)

type RCode valueName

var RCodes = []RCode{
	RCodeNoError,
	RCodeFormErr,
	RCodeServFail,
	RCodeNXDomain,
	RCodeNotImp,
	RCodeRefused,
	RCodeYXDomain,
	RCodeXRRSet,
	RCodeNotAuth,
	RCodeNotZone,
}

var (
	RCodeNoError  = RCode{Value: 0x0000, Name: "NOERROR"}
	RCodeFormErr  = RCode{Value: 0x0001, Name: "FORMERR"}
	RCodeServFail = RCode{Value: 0x0002, Name: "SERVFAIL"}
	RCodeNXDomain = RCode{Value: 0x0003, Name: "NXDOMAIN"}
	RCodeNotImp   = RCode{Value: 0x0004, Name: "NOTIMP"}
	RCodeRefused  = RCode{Value: 0x0005, Name: "REFUSED"}
	RCodeYXDomain = RCode{Value: 0x0006, Name: "YXDOMAIN"}
	RCodeXRRSet   = RCode{Value: 0x0007, Name: "XRRSET"}
	RCodeNotAuth  = RCode{Value: 0x0008, Name: "NOTAUTH"}
	RCodeNotZone  = RCode{Value: 0x0009, Name: "NOTZONE"}
)

var Types = []Type{
	TypeA,
	TypeNS,
	TypeCNAME,
	TypeSOA,
	TypePTR,
	TypeMX,
	TypeTXT,
	TypeRP,
	TypeAFSDB,
	TypeSIG,
	TypeKEY,
	TypeAAAA,
	TypeLOC,
	TypeSRV,
	TypeNAPTR,
	TypeKX,
	TypeCERT,
	TypeDNAME,
	TypeOPT,
	TypeAPL,
	TypeDS,
	TypeSSHFP,
	TypeIPSECKEY,
	TypeRRSIG,
	TypeNSEC,
	TypeDNSKEY,
	TypeDHCID,
	TypeNSEC3,
	TypeNSEC3PARAM,
	TypeTLSA,
	TypeHIP,
	TypeSPF,
	TypeTKEY,
	TypeTSIG,
	TypeIXFR,
	TypeAXFR,
	TypeANY,
	TypeCAA,
	TypeTA,
	TypeDLV,
}

var (
	TypeA          = Type{Value: 0x0001, Name: "A"}
	TypeNS         = Type{Value: 0x0002, Name: "NS"}
	TypeCNAME      = Type{Value: 0x0005, Name: "CNAME"}
	TypeSOA        = Type{Value: 0x0006, Name: "SOA"}
	TypePTR        = Type{Value: 0x000c, Name: "PTR"}
	TypeMX         = Type{Value: 0x000f, Name: "MX"}
	TypeTXT        = Type{Value: 0x0010, Name: "TXT"}
	TypeRP         = Type{Value: 0x0011, Name: "RP"}
	TypeAFSDB      = Type{Value: 0x0012, Name: "AFSDB"}
	TypeSIG        = Type{Value: 0x0018, Name: "SIG"}
	TypeKEY        = Type{Value: 0x0019, Name: "KEY"}
	TypeAAAA       = Type{Value: 0x001c, Name: "AAAA"}
	TypeLOC        = Type{Value: 0x001d, Name: "LOC"}
	TypeSRV        = Type{Value: 0x0021, Name: "SRV"}
	TypeNAPTR      = Type{Value: 0x0023, Name: "NAPTR"}
	TypeKX         = Type{Value: 0x0024, Name: "KX"}
	TypeCERT       = Type{Value: 0x0025, Name: "CERT"}
	TypeDNAME      = Type{Value: 0x0027, Name: "DNAME"}
	TypeOPT        = Type{Value: 0x0029, Name: "OPT"}
	TypeAPL        = Type{Value: 0x002a, Name: "APL"}
	TypeDS         = Type{Value: 0x002b, Name: "DS"}
	TypeSSHFP      = Type{Value: 0x002c, Name: "SSHFP"}
	TypeIPSECKEY   = Type{Value: 0x002d, Name: "IPSECKEY"}
	TypeRRSIG      = Type{Value: 0x002e, Name: "RRSIG"}
	TypeNSEC       = Type{Value: 0x002f, Name: "NSEC"}
	TypeDNSKEY     = Type{Value: 0x0030, Name: "DNSKEY"}
	TypeDHCID      = Type{Value: 0x0031, Name: "DHCID"}
	TypeNSEC3      = Type{Value: 0x0032, Name: "NSEC3"}
	TypeNSEC3PARAM = Type{Value: 0x0033, Name: "NSEC3PARAM"}
	TypeTLSA       = Type{Value: 0x0034, Name: "TLSA"}
	TypeHIP        = Type{Value: 0x0037, Name: "HIP"}
	TypeSPF        = Type{Value: 0x0063, Name: "SPF"}
	TypeTKEY       = Type{Value: 0x00f9, Name: "TKEY"}
	TypeTSIG       = Type{Value: 0x00fa, Name: "TSIG"}
	TypeIXFR       = Type{Value: 0x00fb, Name: "IXFR"}
	TypeAXFR       = Type{Value: 0x00fc, Name: "AXFR"}
	TypeANY        = Type{Value: 0x00ff, Name: "ANY"}
	TypeCAA        = Type{Value: 0x0101, Name: "CAA"}
	TypeTA         = Type{Value: 0x8000, Name: "TA"}
	TypeDLV        = Type{Value: 0x8001, Name: "DLV"}
)

var Classes = []Class{
	ClassIN,
}

var (
	ClassIN Class = Class{Value: 0x0001, Name: "IN"}
)

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
