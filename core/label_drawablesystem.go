// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package core

import (
    
    "sort"

    "github.com/gabstv/ecs/v2"
)









const uuidDrawableLabelSystem = "70EC2F13-4C71-4A3F-9F6D-FF11F5DE9384"

type viewDrawableLabelSystem struct {
    entities []VIDrawableLabelSystem
    world ecs.BaseWorld
    
}

type VIDrawableLabelSystem struct {
    Entity ecs.Entity
    
    Drawable *Drawable 
    
    Label *Label 
    
}

type sortedVIDrawableLabelSystems []VIDrawableLabelSystem
func (a sortedVIDrawableLabelSystems) Len() int           { return len(a) }
func (a sortedVIDrawableLabelSystems) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortedVIDrawableLabelSystems) Less(i, j int) bool { return a[i].Entity < a[j].Entity }

func newviewDrawableLabelSystem(w ecs.BaseWorld) *viewDrawableLabelSystem {
    return &viewDrawableLabelSystem{
        entities: make([]VIDrawableLabelSystem, 0),
        world: w,
    }
}

func (v *viewDrawableLabelSystem) Matches() []VIDrawableLabelSystem {
    
    return v.entities
    
}

func (v *viewDrawableLabelSystem) indexof(e ecs.Entity) int {
    i := sort.Search(len(v.entities), func(i int) bool { return v.entities[i].Entity >= e })
    if i < len(v.entities) && v.entities[i].Entity == e {
        return i
    }
    return -1
}

// Fetch a specific entity
func (v *viewDrawableLabelSystem) Fetch(e ecs.Entity) (data VIDrawableLabelSystem, ok bool) {
    
    i := v.indexof(e)
    if i == -1 {
        return VIDrawableLabelSystem{}, false
    }
    return v.entities[i], true
}

func (v *viewDrawableLabelSystem) Add(e ecs.Entity) bool {
    
    
    // MUST NOT add an Entity twice:
    if i := v.indexof(e); i > -1 {
        return false
    }
    v.entities = append(v.entities, VIDrawableLabelSystem{
        Entity: e,
        Drawable: GetDrawableComponent(v.world).Data(e),
Label: GetLabelComponent(v.world).Data(e),

    })
    if len(v.entities) > 1 {
        if v.entities[len(v.entities)-1].Entity < v.entities[len(v.entities)-2].Entity {
            sort.Sort(sortedVIDrawableLabelSystems(v.entities))
        }
    }
    return true
}

func (v *viewDrawableLabelSystem) Remove(e ecs.Entity) bool {
    
    
    if i := v.indexof(e); i != -1 {

        v.entities = append(v.entities[:i], v.entities[i+1:]...)
        return true
    }
    return false
}

func (v *viewDrawableLabelSystem) rescan() {
    
    
    for _, x := range v.entities {
        e := x.Entity
        
        x.Drawable = GetDrawableComponent(v.world).Data(e)
        
        x.Label = GetLabelComponent(v.world).Data(e)
        
        _ = e
        
    }
}

// DrawableLabelSystem implements ecs.BaseSystem
type DrawableLabelSystem struct {
    initialized bool
    world       ecs.BaseWorld
    view        *viewDrawableLabelSystem
    enabled     bool
    
}

// GetDrawableLabelSystem returns the instance of the system in a World
func GetDrawableLabelSystem(w ecs.BaseWorld) *DrawableLabelSystem {
    return w.S(uuidDrawableLabelSystem).(*DrawableLabelSystem)
}

// Enable system
func (s *DrawableLabelSystem) Enable() {
    s.enabled = true
}

// Disable system
func (s *DrawableLabelSystem) Disable() {
    s.enabled = false
}

// Enabled checks if enabled
func (s *DrawableLabelSystem) Enabled() bool {
    return s.enabled
}

// UUID implements ecs.BaseSystem
func (DrawableLabelSystem) UUID() string {
    return "70EC2F13-4C71-4A3F-9F6D-FF11F5DE9384"
}

func (DrawableLabelSystem) Name() string {
    return "DrawableLabelSystem"
}

// ensure matchfn
var _ ecs.MatchFn = matchDrawableLabelSystem

// ensure resizematchfn
var _ ecs.MatchFn = resizematchDrawableLabelSystem

func (s *DrawableLabelSystem) match(eflag ecs.Flag) bool {
    return matchDrawableLabelSystem(eflag, s.world)
}

func (s *DrawableLabelSystem) resizematch(eflag ecs.Flag) bool {
    return resizematchDrawableLabelSystem(eflag, s.world)
}

func (s *DrawableLabelSystem) ComponentAdded(e ecs.Entity, eflag ecs.Flag) {
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

func (s *DrawableLabelSystem) ComponentRemoved(e ecs.Entity, eflag ecs.Flag) {
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

func (s *DrawableLabelSystem) ComponentResized(cflag ecs.Flag) {
    if s.resizematch(cflag) {
        s.view.rescan()
    }
}

func (s *DrawableLabelSystem) V() *viewDrawableLabelSystem {
    return s.view
}

func (*DrawableLabelSystem) Priority() int64 {
    return 10
}

func (s *DrawableLabelSystem) Setup(w ecs.BaseWorld) {
    if s.initialized {
        panic("DrawableLabelSystem called Setup() more than once")
    }
    s.view = newviewDrawableLabelSystem(w)
    s.world = w
    s.enabled = true
    s.initialized = true
    
}


func init() {
    ecs.RegisterSystem(func() ecs.BaseSystem {
        return &DrawableLabelSystem{}
    })
}
