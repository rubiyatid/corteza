package report

import (
	"context"
)

type (
	// DatasourceProvider provides access to system datasources, such as ComposeRecords
	DatasourceProvider interface {
		// Datasource initializes and returns the Datasource we can use
		Datasource(context.Context, *LoadStepDefinition) (Datasource, error)
	}

	// Loader returns the next Frame from the Datasource; returns nil, nil if no more
	Loader func(cap int, processed bool) ([]*Frame, error)
	// Closer closes the Datasource
	Closer func()

	DatasourceSet []Datasource
	Datasource    interface {
		Name() string
		// Closer return argument may be omitted for some datasources
		Load(context.Context, ...*FrameDefinition) (Loader, Closer, error)
		Describe() FrameDescriptionSet
	}

	// GroupableDatasource is able to provide groupped data
	GroupableDatasource interface {
		Datasource
		Group(GroupDefinition, string) (bool, error)
	}

	// @todo make the underlying DB driver determine this alongside the rdbms package.
	//       somethins similar to how we do typecasting should do the trick.
	PartitionableDatasource interface {
		Datasource
		Partition(partitionSize uint, partitionCol string) (bool, error)
	}

	// @todo TransformableDatasource
)

const (
	defaultPageSize = uint(20)
)

// Merge merges the two DatasourceSets and overwrites any duplicates
func (dd DatasourceSet) Merge(mm DatasourceSet) DatasourceSet {
outer:
	for _, m := range mm {
		for i, d := range dd {
			if d.Name() == m.Name() {
				dd[i] = m
				continue outer
			}
		}
		dd = append(dd, m)
	}

	return dd
}

// Find searches for the Datasource by name
func (dd DatasourceSet) Find(name string) Datasource {
	for _, d := range dd {
		if d.Name() == name {
			return d
		}
	}

	return nil
}
