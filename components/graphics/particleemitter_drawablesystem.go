// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package graphics

import (
    
    "sort"

    "github.com/gabstv/ecs/v2"
    
    "github.com/gabstv/primen/components"
    
)









const uuidParticleEmitterSystem = "627C4B36-EE45-40C6-91AE-617D5CFDD8FC"

type viewParticleEmitterSystem struct {
    entities []VIParticleEmitterSystem
    world ecs.BaseWorld
    
}

type VIParticleEmitterSystem struct {
    Entity ecs.Entity
    
    ParticleEmitter *ParticleEmitter 
    
    Transform *components.Transform 
    
}

type sortedVIParticleEmitterSystems []VIParticleEmitterSystem
func (a sortedVIParticleEmitterSystems) Len() int           { return len(a) }
func (a sortedVIParticleEmitterSystems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortedVIParticleEmitterSystems) Less(i, j int) bool { return a[i].Entity < a[j].Entity }

func newviewParticleEmitterSystem(w ecs.BaseWorld) *viewParticleEmitterSystem {
    return &viewParticleEmitterSystem{
        entities: make([]VIParticleEmitterSystem, 0),
        world: w,
    }
}

func (v *viewParticleEmitterSystem) Matches() []VIParticleEmitterSystem {
    
    return v.entities
    
}

func (v *viewParticleEmitterSystem) indexof(e ecs.Entity) int {
    i := sort.Search(len(v.entities), func(i int) bool { return v.entities[i].Entity >= e })
    if i < len(v.entities) && v.entities[i].Entity == e {
        return i
    }
    return -1
}

// Fetch a specific entity
func (v *viewParticleEmitterSystem) Fetch(e ecs.Entity) (data VIParticleEmitterSystem, ok bool) {
    
    i := v.indexof(e)
    if i == -1 {
        return VIParticleEmitterSystem{}, false
    }
    return v.entities[i], true
}

func (v *viewParticleEmitterSystem) Add(e ecs.Entity) bool {
    
    
    // MUST NOT add an Entity twice:
    if i := v.indexof(e); i > -1 {
        return false
    }
    v.entities = append(v.entities, VIParticleEmitterSystem{
        Entity: e,
        ParticleEmitter: GetParticleEmitterComponent(v.world).Data(e),
Transform: components.GetTransformComponentData(v.world, e),

    })
    if len(v.entities) > 1 {
        if v.entities[len(v.entities)-1].Entity < v.entities[len(v.entities)-2].Entity {
            sort.Sort(sortedVIParticleEmitterSystems(v.entities))
        }
    }
    return true
}

func (v *viewParticleEmitterSystem) Remove(e ecs.Entity) bool {
    
    
    if i := v.indexof(e); i != -1 {

        v.entities = append(v.entities[:i], v.entities[i+1:]...)
        return true
    }
    return false
}

func (v *viewParticleEmitterSystem) clearpointers() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].ParticleEmitter = nil
        
        v.entities[i].Transform = nil
        
        _ = e
    }
}

func (v *viewParticleEmitterSystem) rescan() {
    
    
    for i := range v.entities {
        e := v.entities[i].Entity
        
        v.entities[i].ParticleEmitter = GetParticleEmitterComponent(v.world).Data(e)
        
        v.entities[i].Transform = components.GetTransformComponentData(v.world, e)
        
        _ = e
        
    }
}

// ParticleEmitterSystem implements ecs.BaseSystem
type ParticleEmitterSystem struct {
    initialized bool
    world       ecs.BaseWorld
    view        *viewParticleEmitterSystem
    enabled     bool
    
}

// GetParticleEmitterSystem returns the instance of the system in a World
func GetParticleEmitterSystem(w ecs.BaseWorld) *ParticleEmitterSystem {
    return w.S(uuidParticleEmitterSystem).(*ParticleEmitterSystem)
}

// Enable system
func (s *ParticleEmitterSystem) Enable() {
    s.enabled = true
}

// Disable system
func (s *ParticleEmitterSystem) Disable() {
    s.enabled = false
}

// Enabled checks if enabled
func (s *ParticleEmitterSystem) Enabled() bool {
    return s.enabled
}

// UUID implements ecs.BaseSystem
func (ParticleEmitterSystem) UUID() string {
    return "627C4B36-EE45-40C6-91AE-617D5CFDD8FC"
}

func (ParticleEmitterSystem) Name() string {
    return "ParticleEmitterSystem"
}

// ensure matchfn
var _ ecs.MatchFn = matchParticleEmitterSystem

// ensure resizematchfn
var _ ecs.MatchFn = resizematchParticleEmitterSystem

func (s *ParticleEmitterSystem) match(eflag ecs.Flag) bool {
    return matchParticleEmitterSystem(eflag, s.world)
}

func (s *ParticleEmitterSystem) resizematch(eflag ecs.Flag) bool {
    return resizematchParticleEmitterSystem(eflag, s.world)
}

func (s *ParticleEmitterSystem) ComponentAdded(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            s.onEntityAdded(e)
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            s.onEntityRemoved(e)
        }
    }
}

func (s *ParticleEmitterSystem) ComponentRemoved(e ecs.Entity, eflag ecs.Flag) {
    if s.match(eflag) {
        if s.view.Add(e) {
            // TODO: dispatch event that this entity was added to this system
            s.onEntityAdded(e)
        }
    } else {
        if s.view.Remove(e) {
            // TODO: dispatch event that this entity was removed from this system
            s.onEntityRemoved(e)
        }
    }
}

func (s *ParticleEmitterSystem) ComponentResized(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        s.view.rescan()
        s.onResize()
    }
}

func (s *ParticleEmitterSystem) ComponentWillResize(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        s.onWillResize()
        s.view.clearpointers()
    }
}

func (s *ParticleEmitterSystem) V() *viewParticleEmitterSystem {
    return s.view
}

func (*ParticleEmitterSystem) Priority() int64 {
    return 10
}

func (s *ParticleEmitterSystem) Setup(w ecs.BaseWorld) {
    if s.initialized {
        panic("ParticleEmitterSystem called Setup() more than once")
    }
    s.view = newviewParticleEmitterSystem(w)
    s.world = w
    s.enabled = true
    s.initialized = true
    
}


func init() {
    ecs.RegisterSystem(func() ecs.BaseSystem {
        return &ParticleEmitterSystem{}
    })
}
