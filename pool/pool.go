package pool

import (
	"fmt"
	"sync"

	. "github.com/pool-beta/pool-server/types"
	. "github.com/pool-beta/pool-server/pool/types"
	. "github.com/pool-beta/pool-server/user/types"
)


// PoolFactory contains the API for managing Pools (singleton; manually managed)
type PoolFactory interface {
	// Creates a new Pool
	CreatePool(string, UserID, PoolType) (Pool, error)
	RetrievePool(PoolID) (Pool, error)
	ReturnPool(PoolID) error
}

type poolFactory struct {
	pools map[PoolID]*poolRef
	mutex sync.Mutex
}

// Interal struct for PoolFactory to keep track of active Pools
// TODO: Flushing mechanism of some sort
type poolRef struct {
	pool Pool
	refCount int
	mutex sync.Mutex
}

// TODO: Database would be inserted into PoolFactory
func InitPoolFactory() (PoolFactory, error) {
	pools := make(map[PoolID]*poolRef)

	return &poolFactory{
		pools: pools,
	}, nil
}

func (pf *poolFactory) CreatePool(poolName string, owner UserID, poolType PoolType) (Pool, error) {
	// TODO: Validate owner

	// Create a new pool id
	pid := NewPoolID()
	
	var pool Pool
	switch (poolType) {
	case POOL:
		pool = newPool(pid, poolName, owner)
	case DRAIN:
		pool = newDrain(pid, poolName, owner)
	case TANK:
		pool = newTank(pid, poolName, owner)
	default:
		return nil, fmt.Errorf("Not a valid PoolType -- poolType: %v", poolType)
	}

	// Create poolRef
	pr := &poolRef{
		pool: pool,
		refCount: 1,
	}

	pf.pools[pid] = pr
	return pool, nil
}

func (pf *poolFactory) RetrievePool(pid PoolID) (Pool, error) {
	pr, ok := pf.pools[pid]
	if !ok {
		return nil, fmt.Errorf("Pool does not exist -- pool: %v", pid)
	}

	pr.mutex.Lock()
	defer pr.mutex.Unlock()

	pr.refCount++
	return pr.pool, nil
}

func (pf *poolFactory) ReturnPool(pid PoolID) error {
	pr, ok := pf.pools[pid]
	if !ok {
		return fmt.Errorf("Pool does not exist -- pool: %v", pid)
	}

	pr.mutex.Lock()
	defer pr.mutex.Unlock()

	pr.refCount--

	// TODO: Flush when refCount == 0
	return nil
}


//--------------------------------------------------------------------------------------------------

type Pool interface {
	// Auth
	PoolAuth

	// Getters
	GetReserve() USDollar

	// Stream Control
	AddPusher(Stream) error
	RemovePusher(Stream) error
	AddPuller(Stream) error
	RemovePuller(Stream) error

	// Flow
	Pull(Flow) error
	Push(Flow) error
	// Reset() error // (?)

	// Drop
	PullDrop(Drop, bool) error
	PushDrop(Drop, bool) error

	// For Tests
	Fund(USDollar) USDollar
}

type pool struct {
	*poolAuth
	id PoolID
	name string
	reserve USDollar
	maxReserve USDollar

	pushers []Stream
	pullers []Stream

	mutex sync.Mutex
}

// Wrapper for initPool for PoolType POOL (default type)
func newPool(pid PoolID, name string, owner UserID) Pool {
	p := initPool(pid, name, owner)
	p.reserve = USDollar(0)

	return p
}

// Used by PoolFactory to init a Pool; error check should be already done
func initPool(id PoolID, name string, owner UserID) *pool {
	// Initialize
	pushers := make([]Stream, 0)
	pullers := make([]Stream, 0)

	// Creat PoolAuth
	auth := NewPoolAuth(id)

	// Add Original Owner
	auth.AddOwner(owner)

	return &pool {
		id: id,
		name: name,
		reserve: USDollar(0), // Always start with zero
		maxReserve: MAXUSDOLLAR,

		pushers: pushers,
		pullers: pullers,

		poolAuth: auth,
	}
}

// Triggers

// Flow
func (p *pool) Push(flow Flow) error {
	// Pull the required money first
	err := p.PullDrop(flow.PullDrop(), true)
	if err != nil {
		flow.Invalid()
		return err
	}

	err = p.PushDrop(flow.PushDrop(), false)
	if err != nil {
		flow.Invalid()
		return err
	}

	flow.Valid()
	return nil
}

func (p *pool) Pull(flow Flow) error {
	err := p.PullDrop(flow.PullDrop(), false)
	if err != nil {
		flow.Invalid()
		return err
	}

	// Initiate Push if necessary
	err = p.PushDrop(flow.PushDrop(), true)
	if err != nil {
		flow.Invalid()
		return err
	}

	flow.Valid()
	return nil
}


// Add to this pool
// useReserve refers to putting the drop into this pool's reserve first
func (p *pool) PushDrop(drop Drop, useReserve bool) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	amount := drop.Amount()
	if useReserve {
		// Change drop amount if needed
		total := amount + p.reserve
		if total > p.maxReserve {
			// add max
			diff := p.maxReserve - p.reserve
			drop.AddWithheld(diff)

			amount = amount - diff
			// Give rest to children
		} else {
			// Easy -- leave all here
			drop.AddWithheld(amount)
			return nil
		}
	}
	
	// TODO: Add logic for further push
	return p.push(drop, amount)
}

// Take from this pool
// useReserve refers to pulling from this pool's reserve first
func (p *pool) PullDrop(drop Drop, useReserve bool) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	amount := drop.Amount()
	if useReserve {
		if amount > p.reserve {
			// Take max from reserve
			r := p.reserve
			p.reserve = 0
	
			drop.AddWithheld(r)
			amount = amount - r
	
			// Go on
		} else {
			// Easy -- take from reserve
			p.reserve -= amount
			drop.AddWithheld(amount)
			return nil
		}
	}
	return p.pull(drop, amount)
}

func (p *pool) Reset() error {
	return nil
}

func (p *pool) AddPusher(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.pushers = append(p.pushers, stream) 
	return nil
}

func (p *pool) RemovePusher(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	pushers, ok := FindAndRemove(p.pushers, stream)
	if !ok {
		return fmt.Errorf("User is not a pusher of this pool -- userID: %v; poolID: %v", stream, p.id)
	}
	p.pushers = pushers
	return nil
}

func (p *pool) AddPuller(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	p.pullers = append(p.pullers, stream) 
	return nil
}

func (p *pool) RemovePuller(stream Stream) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	// TODO: Validate user

	pullers, ok := FindAndRemove(p.pullers, stream)
	if !ok {
		return fmt.Errorf("User is not a puller of this pool -- userID: %v; poolID: %v", stream, p.id)
	}
	p.pullers = pullers
	return nil
}

// Getters
func (p *pool) GetReserve() USDollar {
	return p.reserve
}

// For tesing
func (p *pool) Fund(amount USDollar) USDollar {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.reserve += amount
	return p.reserve
}

// -------------------------------------------------------------------------------------------------------------
/* 
	Helper Functions
*/

// Pull from pushers
func (p *pool) pull(drop Drop, amount USDollar) error {
	// Normalize on Pusher Streams
	amounts, err := normalize(p.pushers, amount)
	if err != nil {
		return err
	}

	// Send Drops according to normalized amounts
	for _, s := range p.pushers {
		d, err := s.Pull(amounts[s.StreamID()])
		if err != nil {
			// Drop not possible
			return err
		}
		drop.AddDroplet(d)
	}
	
	return nil
}

// Push to pullers
func (p *pool) push(drop Drop, amount USDollar) error {
	// TODO: Add push logic
	return nil
}

// Returns a map of Normalized USDollar from given streams; error if isn't possible
func normalize(streams []Stream, amount USDollar) (map[StreamID]USDollar, error) {
	if streams == nil {
		return nil, fmt.Errorf("Streams cannot be nil")
	}
	size := len(streams)

	var share, diff USDollar

	amounts := make(map[StreamID]USDollar, size) // return value
	flexible := make(map[StreamID]Stream)
	flexiblePercents := make(map[StreamID]Percent)
	totalAmount := USDollar(0)

	// Filter Streams that allow overdraft
	for _, stream := range streams {
		// Check if Overdraft is allowed for each stream
		if stream.GetAllowOverdraft() {
			// Check if percent of total is less than stream max
			share = amount.MultiplyPercent(stream.GetPercentOverdraft())
			if share > stream.GetMaxOverdraft() {
				// share is too big; take some out
				diff = share - stream.GetMaxOverdraft()
				share -= diff
			} else {
				// check if flexible
				if stream.GetAllowFlexibleOverdraft() {
					// add to flexible map
					flexible[stream.StreamID()] = stream
					flexiblePercents[stream.StreamID()] = stream.GetPercentOverdraft()
				}
			}
			// add to amounts
			amounts[stream.StreamID()] = share
			// Add to total
			totalAmount += share
		} else {
			// Set to zero since no overdraft
			amounts[stream.StreamID()] = USDollar(0)
		}
	}

	if totalAmount == amount {
		return amounts, nil
	} else if totalAmount > amount {
		return nil, fmt.Errorf("Oh no, you have broken the system -- totalAmount: %v; requestedAmount: %v", totalAmount, amount)
	}

	// Need more money
	// For each iteration: (1) normalize share (2) fix max of each stream

	// loop until diff is one cent
	var normalized map[StreamID]Percent
	diff = amount - totalAmount
	for diff > USDollar(len(flexible)) {
		if len(flexible) == 0 {
			// Can't do anything more
			break;
		}

		// Normalize
		normalized = NormalizePercents(flexiblePercents)
		for sid, s := range flexible {
			// Check Max
			share_add := diff.MultiplyPercent(normalized[sid])
			share = share_add + amounts[sid]

			if share > s.GetMaxOverdraft() {
				// share too big; remove from flexible
				diff = share - s.GetMaxOverdraft()
				share -= diff
				delete(flexible, sid)
			}

			amounts[sid] = share
			totalAmount += share
		}
		diff = amount - totalAmount
	}

	// Check if it was possible
	/*
		if flexible is empty than it means that no more streams can give more
		if diff is zero, all amounts were possible to allocate

		if flexible is not empty but diff is greater than zero, it means there are a couple cents
		we cannot split up
	*/

	if len(flexible) == 0 && diff > USDollar(0) {
		return nil, fmt.Errorf("Cannot fulfil drop -- remaining: %v", diff.String())
	}

	// TODO: Clean up the remaining cents


	return amounts, nil
}

