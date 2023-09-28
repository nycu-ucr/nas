// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/nycu-ucr/nas/nasType"
)

type PDUSessionAuthenticationCommand struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity
	nasType.EAPMessage
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionAuthenticationCommand(iei uint8) (pDUSessionAuthenticationCommand *PDUSessionAuthenticationCommand) {
	pDUSessionAuthenticationCommand = &PDUSessionAuthenticationCommand{}
	return pDUSessionAuthenticationCommand
}

const (
	PDUSessionAuthenticationCommandExtendedProtocolConfigurationOptionsType uint8 = 0x7B
)

func (a *PDUSessionAuthenticationCommand) EncodePDUSessionAuthenticationCommand(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/PDUSessionID): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/PTI): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.GetLen()); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/EAPMessage): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/EAPMessage): %w", err)
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionAuthenticationCommand/ExtendedProtocolConfigurationOptions): %w", err)
		}
	}
	return nil
}

func (a *PDUSessionAuthenticationCommand) DecodePDUSessionAuthenticationCommand(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/PDUSessionID): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/PTI): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/PDUSESSIONAUTHENTICATIONCOMMANDMessageIdentity): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.EAPMessage.Len); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/EAPMessage): %w", err)
	}
	if a.EAPMessage.Len < 4 || a.EAPMessage.Len > 1500 {
		return fmt.Errorf("invalid ie length (PDUSessionAuthenticationCommand/EAPMessage): %d", a.EAPMessage.Len)
	}
	a.EAPMessage.SetLen(a.EAPMessage.GetLen())
	if err := binary.Read(buffer, binary.BigEndian, a.EAPMessage.Buffer); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/EAPMessage): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionAuthenticationCommandExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/ExtendedProtocolConfigurationOptions): %w", err)
			}
			if a.ExtendedProtocolConfigurationOptions.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionAuthenticationCommand/ExtendedProtocolConfigurationOptions): %d", a.ExtendedProtocolConfigurationOptions.Len)
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionAuthenticationCommand/ExtendedProtocolConfigurationOptions): %w", err)
			}
		default:
		}
	}
	return nil
}
