package entity

import "sync"

type EntityManager struct {
	lock        sync.Mutex
	generations []uint32 // per-index generation counter
	freeList    []uint32
	entities    map[uint32]Entity
}

func NewEntityManager(capEntites int) *EntityManager {
	return &EntityManager{
		generations: make([]uint32, capEntites),
		freeList:    make([]uint32, 0, capEntites),
		entities:    make(map[uint32]Entity, capEntites),
	}
}

// Create allocates a new Entity, reusing a freed slot if available.
// It is safe for concurrent use.
func (mngr *EntityManager) Create() Entity {
	mngr.lock.Lock()
	defer mngr.lock.Unlock()

	var idx uint32
	if len(mngr.freeList) > 0 {
		// reuse previosly freed index
		idx = mngr.freeList[len(mngr.freeList)-1]
		mngr.freeList = mngr.freeList[:len(mngr.freeList)-1]
	} else {
		idx = uint32(len(mngr.generations))
		mngr.generations = append(mngr.generations, 0)
	}
	ent := Entity{id: EntityId{Index: idx, Generation: mngr.generations[idx]}}
	mngr.entities[idx] = ent

	return ent
}

// Destroy removes the given Entity if it is still valid,
// increments its generation, and makes its slot available for reuse.
// Invalid or already-destroyed Entities are ignored.
func (mngr *EntityManager) Destroy(e Entity) {
	mngr.lock.Lock()
	defer mngr.lock.Unlock()

	idx := e.id.Index

	// early return if has been destroyed already or
	// there is a generation mismatch
	if mngr.generations[idx] != e.id.Generation {
		return
	}

	// increase generation to invalidade old handles
	mngr.generations[idx]++
	delete(mngr.entities, idx)
	mngr.freeList = append(mngr.freeList, idx)
}

// Get retrieves an Entity by its EntityId.
// Returns false if the EntityId is stale or not present.
func (mngr *EntityManager) Get(id EntityId) (Entity, bool) {
	mngr.lock.Lock()
	defer mngr.lock.Unlock()

	if gen := mngr.generations[id.Index]; gen != id.Generation {
		return Entity{}, false
	}

	ent, ok := mngr.entities[id.Index]
	return ent, ok
}

func (mngr *EntityManager) Entites() map[uint32]Entity {
	mngr.lock.Lock()
	defer mngr.lock.Unlock()
	return mngr.entities
}
