// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64
// +build arm64

package client

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type bpfHttpRequestT struct {
	StartTime uint64
	EndTime   uint64
	Method    [10]int8
	Path      [100]int8
	Sc        bpfSpanContext
	Psc       bpfSpanContext
	_         [2]byte
}

type bpfSpanContext struct {
	TraceID [16]uint8
	SpanID  [8]uint8
}

// loadBpf returns the embedded CollectionSpec for bpf.
func loadBpf() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_BpfBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load bpf: %w", err)
	}

	return spec, err
}

// loadBpfObjects loads bpf and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*bpfObjects
//	*bpfPrograms
//	*bpfMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadBpfObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadBpf()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// bpfSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfSpecs struct {
	bpfProgramSpecs
	bpfMapSpecs
}

// bpfSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfProgramSpecs struct {
	UprobeHttpClientDo        *ebpf.ProgramSpec `ebpf:"uprobe_HttpClient_Do"`
	UprobeHttpClientDoReturns *ebpf.ProgramSpec `ebpf:"uprobe_HttpClient_Do_Returns"`
}

// bpfMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type bpfMapSpecs struct {
	AllocMap                  *ebpf.MapSpec `ebpf:"alloc_map"`
	ContextToHttpEvents       *ebpf.MapSpec `ebpf:"context_to_http_events"`
	Events                    *ebpf.MapSpec `ebpf:"events"`
	GolangMapbucketStorageMap *ebpf.MapSpec `ebpf:"golang_mapbucket_storage_map"`
	SpansInProgress           *ebpf.MapSpec `ebpf:"spans_in_progress"`
}

// bpfObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfObjects struct {
	bpfPrograms
	bpfMaps
}

func (o *bpfObjects) Close() error {
	return _BpfClose(
		&o.bpfPrograms,
		&o.bpfMaps,
	)
}

// bpfMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfMaps struct {
	AllocMap                  *ebpf.Map `ebpf:"alloc_map"`
	ContextToHttpEvents       *ebpf.Map `ebpf:"context_to_http_events"`
	Events                    *ebpf.Map `ebpf:"events"`
	GolangMapbucketStorageMap *ebpf.Map `ebpf:"golang_mapbucket_storage_map"`
	SpansInProgress           *ebpf.Map `ebpf:"spans_in_progress"`
}

func (m *bpfMaps) Close() error {
	return _BpfClose(
		m.AllocMap,
		m.ContextToHttpEvents,
		m.Events,
		m.GolangMapbucketStorageMap,
		m.SpansInProgress,
	)
}

// bpfPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadBpfObjects or ebpf.CollectionSpec.LoadAndAssign.
type bpfPrograms struct {
	UprobeHttpClientDo        *ebpf.Program `ebpf:"uprobe_HttpClient_Do"`
	UprobeHttpClientDoReturns *ebpf.Program `ebpf:"uprobe_HttpClient_Do_Returns"`
}

func (p *bpfPrograms) Close() error {
	return _BpfClose(
		p.UprobeHttpClientDo,
		p.UprobeHttpClientDoReturns,
	)
}

func _BpfClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed bpf_bpfel_arm64.o
var _BpfBytes []byte