package gohaystack

import (
	"encoding/json"
	"errors"
	"fmt"
)

// UnmarshalJSON turns a JSON encoded grid into a Grid object
func (g *Grid) UnmarshalJSON(b []byte) error {
	var temp haystackJSONStructure
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	// find all labels
	labels := make(map[string]*Label, len(temp.Cols))
	for _, v := range temp.Cols {
		labels[v.Name] = &Label{
			Value:   v.Name,
			Display: v.Dis,
		}
	}

	var hasVer bool
	var version string
	if v, ok := temp.Meta["Ver"]; ok {
		hasVer = true
		version = v
	}
	if v, ok := temp.Meta["ver"]; ok {
		hasVer = true
		version = v
	}
	if !hasVer {
		return errors.New("missing version tag")
	}
	if version != "3.0" {
		return errors.New("Unsupported version " + version)
	}
	entities := make([]*Entity, 0)
	for _, e := range temp.Rows {
		entity := &Entity{
			tags: make(map[*Label]*Value, len(e)-1),
		}
		if id, ok := e["id"]; ok {
			if id.kind != HaystackTypeRef {
				return fmt.Errorf("bad type for id %v (expected ref)", id)
			}
			entity.id = id.ref
			delete(e, "id")
		} else {
			return fmt.Errorf("row does not have any id %v", e)
		}
		for k, v := range e {
			// TODO find label
			if label, ok := labels[k]; ok {
				entity.tags[label] = v
			} else {
				return fmt.Errorf("bad input: found tag %v in entity that is undeclared", k)

			}
		}
		entities = append(entities, entity)
	}
	g.Meta = temp.Meta
	g.entities = entities
	return nil
}

// MarshalJSON in haystack compatible format
func (g *Grid) MarshalJSON() ([]byte, error) {
	var hasVer bool
	var version string
	if v, ok := g.Meta["Ver"]; ok {
		hasVer = true
		version = v
	}
	if v, ok := g.Meta["ver"]; ok {
		hasVer = true
		version = v
	}
	if !hasVer {
		return nil, errors.New("Bad formatting, missing version tag")
	}
	if version != "3.0" {
		return nil, errors.New("Unsupported version " + version)
	}
	carrier := haystackJSONStructure{}
	carrier.Meta = g.Meta
	labels := make(map[*Label]struct{}, 0)
	for _, entity := range g.entities {
		for label := range entity.tags {
			labels[label] = struct{}{}
		}
	}
	carrier.Cols = make([]haystackJSONCol, len(labels))

	carrier.Rows = make([]map[string]*Value, len(g.entities))
	i := 0
	for label := range labels {
		carrier.Cols[i] = haystackJSONCol{
			Name: label.Value,
			Dis:  label.Display,
		}
		i++
	}
	for i, entity := range g.entities {
		carrier.Rows[i] = make(map[string]*Value, len(entity.tags)+1) // all tags + id
		carrier.Rows[i]["id"] = &Value{
			kind: HaystackTypeRef,
			ref:  entity.id,
		}
		for k, v := range entity.tags {
			carrier.Rows[i][k.Value] = v
		}
	}

	return json.Marshal(carrier)
}

type haystackJSONCol struct {
	Name string `json:"name"`
	Dis  string `json:"dis,omitempty"`
}
type haystackJSONStructure struct {
	Meta map[string]string   `json:"meta"`
	Cols []haystackJSONCol   `json:"cols"`
	Rows []map[string]*Value `json:"rows"`
}
