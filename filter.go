package yum

import (
	"fmt"

	"github.com/cavaliercoder/go-rpm/version"
)

// FilterPackages returns a list of packages filtered according the repo's
// settings.
func FilterPackages(repo *Repo, packages PackageEntries) PackageEntries {
	newest := make(map[string]*PackageEntry, 0)

	// calculate which packages are the latest
	if repo.NewOnly {
		for i, p := range packages {
			// index on name and architecture
			id := fmt.Sprintf("%s.%s", p.Name(), p.Architecture())

			// lookup previous index
			if n, ok := newest[id]; ok {
				// compare version with previous index
				if 1 == version.Compare(&p, n) {
					newest[id] = &packages[i]
				}
			} else {
				// add new index for this package
				newest[id] = &packages[i]
			}
		}

		// replace packages with only the latest packages
		i := 0
		packages = make(PackageEntries, len(newest))
		for _, p := range newest {
			packages[i] = *p
			i++
		}
	}

	// filter the package list
	filtered := make(PackageEntries, 0)
	for _, p := range packages {
		include := true

		// filter by architecture
		if repo.Architecture != "" {
			if p.Architecture() != repo.Architecture {
				include = false
			}
		}

		// filter by minimum build date
		if !repo.MinDate.IsZero() {
			if p.BuildTime().Before(repo.MinDate) {
				include = false
			}
		}

		// filter by maximum build date
		if !repo.MaxDate.IsZero() {
			if p.BuildTime().After(repo.MaxDate) {
				include = false
			}
		}

		// append to output
		if include {
			filtered = append(filtered, p)
		}
	}

	return filtered
}
