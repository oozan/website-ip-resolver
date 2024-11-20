package resolver

import (
	"fmt"
	"net"
	errorhandler "website-ip-resolver/internal/handler"
)

type Resolver struct{}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) Resolve(domain string) ([]net.IP, error) {
	if domain == "" {
		return nil, &errorhandler.CustomError{
			Message: "Domain name cannot be empty",
			Code:    400,
		}
	}

	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, &errorhandler.CustomError{
			Message: fmt.Sprintf("Failed to resolve domain %s", domain),
			Code:    500,
		}
	}

	if len(ips) == 0 {
		return nil, &errorhandler.CustomError{
			Message: fmt.Sprintf("No IP addresses found for domain %s", domain),
			Code:    404,
		}
	}

	return ips, nil
}
