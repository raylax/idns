package protocol

type Protocol struct {
	TransactionId int

	Flags

	QuestionRRs RRs

	AnswerRRs RRs

	AuthorityRRs RRs

	AdditionalRRs RRs

	Questions []Question

	Answers []Answer

	Authorities []Authority

	Additions []Addition
}

type RRs int

const (
	ShiftQR = 0

	ShiftOpCode = ShiftQR + 1

	ShiftAA = ShiftOpCode + 4

	ShiftTC = ShiftAA + 1

	ShiftRD = ShiftTC + 1

	ShiftRA = ShiftRD + 1

	ShiftZero = ShiftRA + 1

	ShiftRCode = ShiftZero + 3
)
