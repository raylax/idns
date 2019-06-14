package protocol

import "encoding/json"

type Protocol struct {
	TransactionId int
	Flags         Flags
	QuestionRRs   int
	AnswerRRs     int
	AuthorityRRs  int
	AdditionalRRs int
	Questions     []Question
	Answers       []Answer
	Authorities   []Authority
	Additions     []Addition
}

func (p *Protocol) String() string {
	j, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(j)
}

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
