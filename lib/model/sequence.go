package model

import (
	"github.com/dbsteward/dbsteward/lib/util"
)

type Sequence struct {
	Name          string   `xml:"name,attr"`
	Owner         string   `xml:"owner,attr,omitempty"`
	Description   string   `xml:"description,attr,omitempty"`
	Cache         *int     `xml:"cache,attr,omitempty"`
	Start         *int     `xml:"start,attr,omitempty"`
	Min           *int     `xml:"min,attr,omitempty"`
	Max           *int     `xml:"max,attr,omitempty"`
	Increment     *int     `xml:"inc,attr,omitempty"`
	Cycle         bool     `xml:"cycle,attr,omitempty"`
	OwnedByColumn string   `xml:"ownedBy,attr,omitempty"`
	SlonyId       int      `xml:"slonyId,attr,omitempty"`
	SlonySetId    *int     `xml:"slonySetId,attr,omitempty"`
	Grants        []*Grant `xml:"grant"`
}

func (self *Sequence) GetGrantsForRole(role string) []*Grant {
	out := []*Grant{}
	for _, grant := range self.Grants {
		if util.IIndexOfStr(role, grant.Roles) >= 0 {
			out = append(out, grant)
		}
	}
	return out
}

func (self *Sequence) GetGrants() []*Grant {
	return self.Grants
}

func (self *Sequence) AddGrant(grant *Grant) {
	// TODO(feat) sanity check
	self.Grants = append(self.Grants, grant)
}

func (self *Sequence) Merge(overlay *Sequence) {
	if overlay == nil {
		return
	}

	self.Owner = overlay.Owner
	self.Cache = overlay.Cache
	self.Start = overlay.Start
	self.Min = overlay.Min
	self.Max = overlay.Max
	self.Increment = overlay.Increment
	self.Cycle = overlay.Cycle

	for _, overlayGrant := range overlay.Grants {
		self.AddGrant(overlayGrant)
	}
}
