// Code generated by ecs https://github.com/gabstv/ecs; DO NOT EDIT.

package core

import (
    
    "sort"

    "github.com/gabstv/ecs/v2"
)







const uuidDrawableComponent = "E3086C37-F0F5-4BFD-8FEE-F9C451B1E57E"
const capDrawableComponent = 256

type drawerDrawableComponent struct {
    Entity ecs.Entity
    Data   Drawable
}

type slcdrawerDrawableComponent []drawerDrawableComponent
func (a slcdrawerDrawableComponent) Len() int           { return len(a) }
func (a slcdrawerDrawableComponent) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a slcdrawerDrawableComponent) Less(i, j int) bool { return a[i].Entity < a[j].Entity }

// DrawableComponent implements ecs.BaseComponent
type DrawableComponent struct {
    initialized bool
    flag        ecs.Flag
    world       ecs.BaseWorld
    wkey        [4]byte
    data        []drawerDrawableComponent
    
}

// GetDrawableComponent returns the instance of the component in a World
func GetDrawableComponent(w ecs.BaseWorld) *DrawableComponent {
    return w.C(uuidDrawableComponent).(*DrawableComponent)
}

// SetDrawableComponentData updates/adds a Drawable to Entity e
func SetDrawableComponentData(w ecs.BaseWorld, e ecs.Entity, data Drawable) {
    GetDrawableComponent(w).Upsert(e, data)
}

// GetDrawableComponentData gets the *Drawable of Entity e
func GetDrawableComponentData(w ecs.BaseWorld, e ecs.Entity) *Drawable {
    return GetDrawableComponent(w).Data(e)
}

// UUID implements ecs.BaseComponent
func (DrawableComponent) UUID() string {
    return "E3086C37-F0F5-4BFD-8FEE-F9C451B1E57E"
}

// Name implements ecs.BaseComponent
func (DrawableComponent) Name() string {
    return "DrawableComponent"
}

func (c *DrawableComponent) indexof(e ecs.Entity) int {
    i := sort.Search(len(c.data), func(i int) bool { return c.data[i].Entity >= e })
    if i < len(c.data) && c.data[i].Entity == e {
        return i
    }
    return -1
}

// Upsert creates or updates a component data of an entity.
// Not recommended to be used directly. Use SetDrawableComponentData to change component
// data outside of a system loop.
func (c *DrawableComponent) Upsert(e ecs.Entity, data interface{}) {
    v, ok := data.(Drawable)
    if !ok {
        panic("data must be Drawable")
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
    c.data = append(c.data, drawerDrawableComponent{
        Entity: e,
        Data:   v,
    })
    if len(c.data) > 1 {
        if c.data[newindex].Entity < c.data[newindex-1].Entity {
            sort.Sort(slcdrawerDrawableComponent(c.data))
        }
    }
    
    c.world.CAdded(e, c, c.wkey)
    if rsz {
        c.world.CResized(c, c.wkey)
    }
    
}

// Remove a Drawable data from entity e
//
// Warning: DO NOT call remove inside the system entities loop
func (c *DrawableComponent) Remove(e ecs.Entity) {
    
    
    i := c.indexof(e)
    if i == -1 {
        return
    }
    
    //c.data = append(c.data[:i], c.data[i+1:]...)
    c.data = c.data[:i+copy(c.data[i:], c.data[i+1:])]
    c.world.CRemoved(e, c, c.wkey)
    
}

func (c *DrawableComponent) Data(e ecs.Entity) *Drawable {
    
    
    index := c.indexof(e)
    if index > -1 {
        return &c.data[index].Data
    }
    return nil
}

// Flag returns the 
func (c *DrawableComponent) Flag() ecs.Flag {
    return c.flag
}

// Setup is called by ecs.BaseWorld
//
// Do not call this directly
func (c *DrawableComponent) Setup(w ecs.BaseWorld, f ecs.Flag, key [4]byte) {
    if c.initialized {
        panic("DrawableComponent called Setup() more than once")
    }
    c.flag = f
    c.world = w
    c.wkey = key
    c.data = make([]drawerDrawableComponent, 0, 256)
    c.initialized = true
}


func init() {
    ecs.RegisterComponent(func() ecs.BaseComponent {
        return &DrawableComponent{}
    })
}
