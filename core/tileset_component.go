// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package core

import (
    "sort"
    

    "github.com/gabstv/ecs/v2"
)








const uuidTileSetComponent = "775FFA75-9F2F-423A-A905-D48E4D562AE8"
const capTileSetComponent = 256

type drawerTileSetComponent struct {
    Entity ecs.Entity
    Data   TileSet
}

// WatchTileSet is a helper struct to access a valid pointer of TileSet
type WatchTileSet interface {
    Entity() ecs.Entity
    Data() *TileSet
}

type slcdrawerTileSetComponent []drawerTileSetComponent
func (a slcdrawerTileSetComponent) Len() int           { return len(a) }
func (a slcdrawerTileSetComponent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slcdrawerTileSetComponent) Less(i, j int) bool { return a[i].Entity < a[j].Entity }


type mWatchTileSet struct {
    c *TileSetComponent
    entity ecs.Entity
}

func (w *mWatchTileSet) Entity() ecs.Entity {
    return w.entity
}

func (w *mWatchTileSet) Data() *TileSet {
    
    
    id := w.c.indexof(w.entity)
    if id == -1 {
        return nil
    }
    return &w.c.data[id].Data
}

// TileSetComponent implements ecs.BaseComponent
type TileSetComponent struct {
    initialized bool
    flag        ecs.Flag
    world       ecs.BaseWorld
    wkey        [4]byte
    data        []drawerTileSetComponent
    
}

// GetTileSetComponent returns the instance of the component in a World
func GetTileSetComponent(w ecs.BaseWorld) *TileSetComponent {
    return w.C(uuidTileSetComponent).(*TileSetComponent)
}

// SetTileSetComponentData updates/adds a TileSet to Entity e
func SetTileSetComponentData(w ecs.BaseWorld, e ecs.Entity, data TileSet) {
    GetTileSetComponent(w).Upsert(e, data)
}

// GetTileSetComponentData gets the *TileSet of Entity e
func GetTileSetComponentData(w ecs.BaseWorld, e ecs.Entity) *TileSet {
    return GetTileSetComponent(w).Data(e)
}

// WatchTileSetComponentData gets a pointer getter of an entity's TileSet.
//
// The pointer must not be stored because it may become invalid overtime.
func WatchTileSetComponentData(w ecs.BaseWorld, e ecs.Entity) WatchTileSet {
    return &mWatchTileSet{
        c: GetTileSetComponent(w),
        entity: e,
    }
}

// UUID implements ecs.BaseComponent
func (TileSetComponent) UUID() string {
    return "775FFA75-9F2F-423A-A905-D48E4D562AE8"
}

// Name implements ecs.BaseComponent
func (TileSetComponent) Name() string {
    return "TileSetComponent"
}

func (c *TileSetComponent) indexof(e ecs.Entity) int {
    i := sort.Search(len(c.data), func(i int) bool { return c.data[i].Entity >= e })
    if i < len(c.data) && c.data[i].Entity == e {
        return i
    }
    return -1
}

// Upsert creates or updates a component data of an entity.
// Not recommended to be used directly. Use SetTileSetComponentData to change component
// data outside of a system loop.
func (c *TileSetComponent) Upsert(e ecs.Entity, data interface{}) {
    v, ok := data.(TileSet)
    if !ok {
        panic("data must be TileSet")
    }
    
    id := c.indexof(e)
    
    if id > -1 {
        
        dwr := &c.data[id]
        dwr.Data = v
        
        return
    }
    
    rsz := false
    if cap(c.data) == len(c.data) {
        rsz = true
    }
    newindex := len(c.data)
    c.data = append(c.data, drawerTileSetComponent{
        Entity: e,
        Data:   v,
    })
    if len(c.data) > 1 {
        if c.data[newindex].Entity < c.data[newindex-1].Entity {
            sort.Sort(slcdrawerTileSetComponent(c.data))
        }
    }
    
    if rsz {
        
        c.world.CResized(c, c.wkey)
        c.world.Dispatch(ecs.Event{
            Type: ecs.EvtComponentsResized,
            ComponentName: "TileSetComponent",
            ComponentID: "775FFA75-9F2F-423A-A905-D48E4D562AE8",
        })
    }
    c.onAdd(e)
    c.world.CAdded(e, c, c.wkey)
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentAdded,
        ComponentName: "TileSetComponent",
        ComponentID: "775FFA75-9F2F-423A-A905-D48E4D562AE8",
        Entity: e,
    })
}

// Remove a TileSet data from entity e
//
// Warning: DO NOT call remove inside the system entities loop
func (c *TileSetComponent) Remove(e ecs.Entity) {
    
    
    i := c.indexof(e)
    if i == -1 {
        return
    }
    c.beforeRemove(e)
    //c.data = append(c.data[:i], c.data[i+1:]...)
    c.data = c.data[:i+copy(c.data[i:], c.data[i+1:])]
    c.world.CRemoved(e, c, c.wkey)
    
    c.world.Dispatch(ecs.Event{
        Type: ecs.EvtComponentRemoved,
        ComponentName: "TileSetComponent",
        ComponentID: "775FFA75-9F2F-423A-A905-D48E4D562AE8",
        Entity: e,
    })
}

func (c *TileSetComponent) Data(e ecs.Entity) *TileSet {
    
    
    index := c.indexof(e)
    if index > -1 {
        return &c.data[index].Data
    }
    return nil
}

// Flag returns the 
func (c *TileSetComponent) Flag() ecs.Flag {
    return c.flag
}

// Setup is called by ecs.BaseWorld
//
// Do not call this directly
func (c *TileSetComponent) Setup(w ecs.BaseWorld, f ecs.Flag, key [4]byte) {
    if c.initialized {
        panic("TileSetComponent called Setup() more than once")
    }
    c.flag = f
    c.world = w
    c.wkey = key
    c.data = make([]drawerTileSetComponent, 0, 256)
    c.initialized = true
}


func init() {
    ecs.RegisterComponent(func() ecs.BaseComponent {
        return &TileSetComponent{}
    })
}
