// Code generated by generate.sh, DO NOT EDIT.

package nasMessage

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/nycu-ucr/nas/nasType"
)

type PDUSessionModificationRequest struct {
	nasType.ExtendedProtocolDiscriminator
	nasType.PDUSessionID
	nasType.PTI
	nasType.PDUSESSIONMODIFICATIONREQUESTMessageIdentity
	*nasType.Capability5GSM
	*nasType.Cause5GSM
	*nasType.MaximumNumberOfSupportedPacketFilters
	*nasType.AlwaysonPDUSessionRequested
	*nasType.IntegrityProtectionMaximumDataRate
	*nasType.RequestedQosRules
	*nasType.RequestedQosFlowDescriptions
	*nasType.MappedEPSBearerContexts
	*nasType.ExtendedProtocolConfigurationOptions
}

func NewPDUSessionModificationRequest(iei uint8) (pDUSessionModificationRequest *PDUSessionModificationRequest) {
	pDUSessionModificationRequest = &PDUSessionModificationRequest{}
	return pDUSessionModificationRequest
}

const (
	PDUSessionModificationRequestCapability5GSMType                        uint8 = 0x28
	PDUSessionModificationRequestCause5GSMType                             uint8 = 0x59
	PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType uint8 = 0x55
	PDUSessionModificationRequestAlwaysonPDUSessionRequestedType           uint8 = 0x0B
	PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType    uint8 = 0x13
	PDUSessionModificationRequestRequestedQosRulesType                     uint8 = 0x7A
	PDUSessionModificationRequestRequestedQosFlowDescriptionsType          uint8 = 0x79
	PDUSessionModificationRequestMappedEPSBearerContextsType               uint8 = 0x75
	PDUSessionModificationRequestExtendedProtocolConfigurationOptionsType  uint8 = 0x7B
)

func (a *PDUSessionModificationRequest) EncodePDUSessionModificationRequest(buffer *bytes.Buffer) error {
	if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/PDUSessionID): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/PTI): %w", err)
	}
	if err := binary.Write(buffer, binary.BigEndian, a.PDUSESSIONMODIFICATIONREQUESTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/PDUSESSIONMODIFICATIONREQUESTMessageIdentity): %w", err)
	}
	if a.Capability5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/Capability5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/Capability5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/Capability5GSM): %w", err)
		}
	}
	if a.Cause5GSM != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/Cause5GSM): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.Cause5GSM.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/Cause5GSM): %w", err)
		}
	}
	if a.MaximumNumberOfSupportedPacketFilters != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/MaximumNumberOfSupportedPacketFilters): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/MaximumNumberOfSupportedPacketFilters): %w", err)
		}
	}
	if a.AlwaysonPDUSessionRequested != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.AlwaysonPDUSessionRequested.Octet); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/AlwaysonPDUSessionRequested): %w", err)
		}
	}
	if a.IntegrityProtectionMaximumDataRate != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.IntegrityProtectionMaximumDataRate.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/IntegrityProtectionMaximumDataRate): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.IntegrityProtectionMaximumDataRate.Octet[:]); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/IntegrityProtectionMaximumDataRate): %w", err)
		}
	}
	if a.RequestedQosRules != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosRules.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/RequestedQosRules): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosRules.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/RequestedQosRules): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosRules.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/RequestedQosRules): %w", err)
		}
	}
	if a.RequestedQosFlowDescriptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosFlowDescriptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/RequestedQosFlowDescriptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosFlowDescriptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/RequestedQosFlowDescriptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.RequestedQosFlowDescriptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/RequestedQosFlowDescriptions): %w", err)
		}
	}
	if a.MappedEPSBearerContexts != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/MappedEPSBearerContexts): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/MappedEPSBearerContexts): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/MappedEPSBearerContexts): %w", err)
		}
	}
	if a.ExtendedProtocolConfigurationOptions != nil {
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetIei()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.GetLen()); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/ExtendedProtocolConfigurationOptions): %w", err)
		}
		if err := binary.Write(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
			return fmt.Errorf("NAS encode error (PDUSessionModificationRequest/ExtendedProtocolConfigurationOptions): %w", err)
		}
	}
	return nil
}

func (a *PDUSessionModificationRequest) DecodePDUSessionModificationRequest(byteArray *[]byte) error {
	buffer := bytes.NewBuffer(*byteArray)
	if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolDiscriminator.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/ExtendedProtocolDiscriminator): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSessionID.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/PDUSessionID): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PTI.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/PTI): %w", err)
	}
	if err := binary.Read(buffer, binary.BigEndian, &a.PDUSESSIONMODIFICATIONREQUESTMessageIdentity.Octet); err != nil {
		return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/PDUSESSIONMODIFICATIONREQUESTMessageIdentity): %w", err)
	}
	for buffer.Len() > 0 {
		var ieiN uint8
		var tmpIeiN uint8
		if err := binary.Read(buffer, binary.BigEndian, &ieiN); err != nil {
			return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/iei): %w", err)
		}
		// fmt.Println(ieiN)
		if ieiN >= 0x80 {
			tmpIeiN = (ieiN & 0xf0) >> 4
		} else {
			tmpIeiN = ieiN
		}
		// fmt.Println("type", tmpIeiN)
		switch tmpIeiN {
		case PDUSessionModificationRequestCapability5GSMType:
			a.Capability5GSM = nasType.NewCapability5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Capability5GSM.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/Capability5GSM): %w", err)
			}
			if a.Capability5GSM.Len < 1 || a.Capability5GSM.Len > 13 {
				return fmt.Errorf("invalid ie length (PDUSessionModificationRequest/Capability5GSM): %d", a.Capability5GSM.Len)
			}
			a.Capability5GSM.SetLen(a.Capability5GSM.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.Capability5GSM.Octet[:a.Capability5GSM.GetLen()]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/Capability5GSM): %w", err)
			}
		case PDUSessionModificationRequestCause5GSMType:
			a.Cause5GSM = nasType.NewCause5GSM(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.Cause5GSM.Octet); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/Cause5GSM): %w", err)
			}
		case PDUSessionModificationRequestMaximumNumberOfSupportedPacketFiltersType:
			a.MaximumNumberOfSupportedPacketFilters = nasType.NewMaximumNumberOfSupportedPacketFilters(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.MaximumNumberOfSupportedPacketFilters.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/MaximumNumberOfSupportedPacketFilters): %w", err)
			}
		case PDUSessionModificationRequestAlwaysonPDUSessionRequestedType:
			a.AlwaysonPDUSessionRequested = nasType.NewAlwaysonPDUSessionRequested(ieiN)
			a.AlwaysonPDUSessionRequested.Octet = ieiN
		case PDUSessionModificationRequestIntegrityProtectionMaximumDataRateType:
			a.IntegrityProtectionMaximumDataRate = nasType.NewIntegrityProtectionMaximumDataRate(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, a.IntegrityProtectionMaximumDataRate.Octet[:]); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/IntegrityProtectionMaximumDataRate): %w", err)
			}
		case PDUSessionModificationRequestRequestedQosRulesType:
			a.RequestedQosRules = nasType.NewRequestedQosRules(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedQosRules.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/RequestedQosRules): %w", err)
			}
			if a.RequestedQosRules.Len < 4 {
				return fmt.Errorf("invalid ie length (PDUSessionModificationRequest/RequestedQosRules): %d", a.RequestedQosRules.Len)
			}
			a.RequestedQosRules.SetLen(a.RequestedQosRules.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RequestedQosRules.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/RequestedQosRules): %w", err)
			}
		case PDUSessionModificationRequestRequestedQosFlowDescriptionsType:
			a.RequestedQosFlowDescriptions = nasType.NewRequestedQosFlowDescriptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.RequestedQosFlowDescriptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/RequestedQosFlowDescriptions): %w", err)
			}
			if a.RequestedQosFlowDescriptions.Len < 3 {
				return fmt.Errorf("invalid ie length (PDUSessionModificationRequest/RequestedQosFlowDescriptions): %d", a.RequestedQosFlowDescriptions.Len)
			}
			a.RequestedQosFlowDescriptions.SetLen(a.RequestedQosFlowDescriptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.RequestedQosFlowDescriptions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/RequestedQosFlowDescriptions): %w", err)
			}
		case PDUSessionModificationRequestMappedEPSBearerContextsType:
			a.MappedEPSBearerContexts = nasType.NewMappedEPSBearerContexts(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.MappedEPSBearerContexts.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/MappedEPSBearerContexts): %w", err)
			}
			if a.MappedEPSBearerContexts.Len < 4 {
				return fmt.Errorf("invalid ie length (PDUSessionModificationRequest/MappedEPSBearerContexts): %d", a.MappedEPSBearerContexts.Len)
			}
			a.MappedEPSBearerContexts.SetLen(a.MappedEPSBearerContexts.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.MappedEPSBearerContexts.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/MappedEPSBearerContexts): %w", err)
			}
		case PDUSessionModificationRequestExtendedProtocolConfigurationOptionsType:
			a.ExtendedProtocolConfigurationOptions = nasType.NewExtendedProtocolConfigurationOptions(ieiN)
			if err := binary.Read(buffer, binary.BigEndian, &a.ExtendedProtocolConfigurationOptions.Len); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/ExtendedProtocolConfigurationOptions): %w", err)
			}
			if a.ExtendedProtocolConfigurationOptions.Len < 1 {
				return fmt.Errorf("invalid ie length (PDUSessionModificationRequest/ExtendedProtocolConfigurationOptions): %d", a.ExtendedProtocolConfigurationOptions.Len)
			}
			a.ExtendedProtocolConfigurationOptions.SetLen(a.ExtendedProtocolConfigurationOptions.GetLen())
			if err := binary.Read(buffer, binary.BigEndian, a.ExtendedProtocolConfigurationOptions.Buffer); err != nil {
				return fmt.Errorf("NAS decode error (PDUSessionModificationRequest/ExtendedProtocolConfigurationOptions): %w", err)
			}
		default:
		}
	}
	return nil
}
