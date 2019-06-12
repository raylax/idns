package protocol

type Name string

type Type int16

type Class int16

type TTL int32

const (
	TypeA          Type = 0x0001
	TypeNS         Type = 0x0002
	TypeCNAME      Type = 0x0005
	TypeSOA        Type = 0x0006
	TypePTR        Type = 0x000c
	TypeMX         Type = 0x000f
	TypeTXT        Type = 0x0010
	TypeRP         Type = 0x0011
	TypeAFSDB      Type = 0x0012
	TypeSIG        Type = 0x0018
	TypeKEY        Type = 0x0019
	TypeAAAA       Type = 0x001c
	TypeLOC        Type = 0x001d
	TypeSRV        Type = 0x0021
	TypeNAPTR      Type = 0x0023
	TypeKX         Type = 0x0024
	TypeCERT       Type = 0x0025
	TypeDNAME      Type = 0x0027
	TypeOPT        Type = 0x0029
	TypeAPL        Type = 0x002a
	TypeDS         Type = 0x002b
	TypeSSHFP      Type = 0x002c
	TypeIPSECKEY   Type = 0x002d
	TypeRRSIG      Type = 0x002e
	TypeNSEC       Type = 0x002f
	TypeDNSKEY     Type = 0x0030
	TypeDHCID      Type = 0x0031
	TypeNSEC3      Type = 0x0032
	TypeNSEC3PARAM Type = 0x0033
	TypeTLSA       Type = 0x0034
	TypeHIP        Type = 0x0037
	TypeSPF        Type = 0x0063
	TypeTKEY       Type = 0x00f9
	TypeTSIG       Type = 0x00fa
	TypeIXFR       Type = 0x00fb
	TypeAXFR       Type = 0x00fc
	TypeANY        Type = 0x00ff
	TypeCAA        Type = 0x0101
)

const (
	ClassIN Class = 0x0001
)
