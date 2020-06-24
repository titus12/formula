package lua

import (
	lua "github.com/yuin/gopher-lua"
	"sync"
)

// Global LState pool
var LuaPool = &lStatePool{
	pool: make([]*lua.LState, 0, 10),
}

type lStatePool struct {
	m     sync.Mutex
	pool []*lua.LState
}

func (pl *lStatePool) Get() *lua.LState {
	pl.m.Lock()
	defer pl.m.Unlock()
	n := len(pl.pool)
	if n == 0 {
		return pl.New()
	}
	s := pl.pool[n-1]
	pl.pool = pl.pool[0 : n-1]
	return s
}

func (pl *lStatePool) New() *lua.LState {
	L := lua.NewState()
	// setting the L up here.
	// load scripts, set global variables, share channels, etc...
	return L
}

func (pl *lStatePool) Put(L *lua.LState) {
	pl.m.Lock()
	defer pl.m.Unlock()
	pl.pool = append(pl.pool, L)
}

func (pl *lStatePool) Shutdown() {
	for _, L := range pl.pool {
		L.Close()
	}
}


