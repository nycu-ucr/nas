// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/nycu-ucr/nas/nasType"
)

type DeregistrationAcceptUETerminatedDeregistration struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.SpareHalfOctetAndSecurityHeaderType
	nasType.DeregistrationAcceptMessageIdentity
}

func NewDeregistrationAcceptUETerminatedDeregistration(iei uint8) (deregistrationAcceptUETerminatedDeregistration *DeregistrationAcceptUETerminatedDeregistration) {
	deregistrationAcceptUETerminatedDeregistration = &DeregistrationAcceptUETerminatedDeregistration{}
	return deregistrationAcceptUETerminatedDeregistration
}

func (a *DeregistrationAcceptUETerminatedDeregistration) EncodeDeregistrationAcceptUETerminatedDeregistration(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationAcceptUETerminatedDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationAcceptUETerminatedDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.DeregistrationAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (DeregistrationAcceptUETerminatedDeregistration/DeregistrationAcceptMessageIdentity): %w", err)
	}
	return nil
}

func (a *DeregistrationAcceptUETerminatedDeregistration) DecodeDeregistrationAcceptUETerminatedDeregistration(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationAcceptUETerminatedDeregistration/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.SpareHalfOctetAndSecurityHeaderType.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationAcceptUETerminatedDeregistration/SpareHalfOctetAndSecurityHeaderType): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.DeregistrationAcceptMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (DeregistrationAcceptUETerminatedDeregistration/DeregistrationAcceptMessageIdentity): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (DeregistrationAcceptUETerminatedDeregistration/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		default:
		}
	}
	return nil
}
