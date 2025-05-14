package mutex

import "sync"

type MutextEmbedded struct {
	sync.Mutex // want `sync.Mytex should not be embedded`
}

type MutextNotEmbedded struct {
	mu sync.Mutex
}

type PointerMutextEmbedded struct {
	*sync.Mutex // want `sync.Mytex should not be embedded`
}
