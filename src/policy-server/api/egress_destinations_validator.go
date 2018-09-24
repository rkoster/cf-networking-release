package api

import (
	"errors"
	"fmt"
)

type EgressDestinationsValidator struct{}

func (v *EgressDestinationsValidator) ValidateEgressDestinations(destinations []EgressDestination) error {
	//if len(policies) == 0 {
	//	return errors.New("missing policies")
	//}
	//
	for _, destination := range destinations {

		if destination.Name == "" {
			return errors.New("missing destination name")
		}

		if len(destination.IPRanges) == 0 {
			return errors.New("missing destination IP range")
		}

		if len(destination.Ports) == 0 {
			return errors.New("missing destination ports")
		}

		if destination.Protocol == "" {
			return errors.New("missing destination protocol")
		}

		if !isValidProtocol(destination.Protocol){
			return fmt.Errorf("invalid destination protocol '%s', specify either tcp, udp, or icmp", destination.Protocol)
		}

		if destination.Protocol == "icmp" && destination.ICMPType != nil {
			
		}

	//	if policy.Source.ID == "" {
	//		return errors.New("missing source id")
	//	}
	//
	//	if policy.Destination.ID == "" {
	//		return errors.New("missing destination id")
	//	}
	//
	//	if policy.Destination.Protocol != "udp" && policy.Destination.Protocol != "tcp" {
	//		return errors.New("invalid destination protocol, specify either udp or tcp")
	//	}
	//
	//	if policy.Destination.Ports.Start > policy.Destination.Ports.End {
	//		return fmt.Errorf("invalid port range %d-%d, start must be less than or equal to end", policy.Destination.Ports.Start, policy.Destination.Ports.End)
	//	}
	//
	//	if policy.Destination.Ports.Start < 0 {
	//		return fmt.Errorf("invalid start port %d, must be in range 1-65535", policy.Destination.Ports.Start)
	//	}
	//
	//	if policy.Destination.Ports.Start == 0 {
	//		return fmt.Errorf("missing start port")
	//	}
	//
	//	if policy.Destination.Ports.End > 65535 {
	//		return fmt.Errorf("invalid end port %d, must be in range 1-65535", policy.Destination.Ports.End)
	//	}
	//
	//	if policy.Source.Tag != "" || policy.Destination.Tag != "" {
	//		return errors.New("tags may not be specified")
	//	}
	}
	return nil
}

func isValidProtocol(protocol string) bool {
	validProtocols := []string{"tcp", "udp", "icmp"}

	for _, validProtocol := range validProtocols {
		if validProtocol == protocol {
			return true
		}
	}
		return false
}