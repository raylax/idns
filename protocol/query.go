package protocol

type Question struct {

	Name

	Type

	Class

}

type Name string

type Type int16

type Class int16

const (
	TYPE_A			Type = 0x0001
	TYPE_NS			Type = 0x0002
	TYPE_CNAME		Type = 0x0005
	TYPE_SOA		Type = 0x0006
	TYPE_PTR		Type = 0x000c
	TYPE_MX			Type = 0x000f
	TYPE_TXT		Type = 0x0010
	TYPE_RP			Type = 0x0011
	TYPE_AFSDB		Type = 0x0012
	TYPE_SIG		Type = 0x0018
	TYPE_KEY		Type = 0x0019
	TYPE_AAAA		Type = 0x001c
	TYPE_LOC		Type = 0x001d
	TYPE_SRV		Type = 0x0021
	TYPE_NAPTR		Type = 0x0023
	TYPE_KX			Type = 0x0024
	TYPE_CERT		Type = 0x0025
	TYPE_DNAME		Type = 0x0027
	TYPE_OPT		Type = 0x0029
	TYPE_APL		Type = 0x002a
	TYPE_DS			Type = 0x002b
	TYPE_SSHFP		Type = 0x002c
	TYPE_IPSECKEY	Type = 0x002d
	TYPE_RRSIG		Type = 0x002e
	TYPE_NSEC		Type = 0x002f
	TYPE_DNSKEY		Type = 0x0030
	TYPE_DHCID		Type = 0x0031
	TYPE_NSEC3		Type = 0x0032
	TYPE_NSEC3PARAM	Type = 0x0033
	TYPE_TLSA		Type = 0x0034
	TYPE_HIP		Type = 0x0037
	TYPE_SPF		Type = 0x0063
	TYPE_TKEY		Type = 0x00f9
	TYPE_TSIG		Type = 0x00fa
	TYPE_IXFR		Type = 0x00fb
	TYPE_AXFR		Type = 0x00fc
	TYPE_ANY		Type = 0x00ff
	TYPE_CAA		Type = 0x0101
	TYPE_TA			Type = 0x8000
	TYPE_DLV		Type = 0x8001
)

const (
	CLASS_IN		Class = 0x0001
)