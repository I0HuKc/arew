package repository

import (
	"context"

	domainDTO "github.com/g10z3r/archx/internal/domain/dto"
)

type ScannerRepository interface {
	Register(ctx context.Context, result *domainDTO.ScanResultDTO) error
	PackageRepo() PackageRepository
}
