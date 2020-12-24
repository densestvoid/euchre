package main

type Group []int

func (g Group) Equal(other Group) bool {
	for _, member := range g {
		var match bool
		for _, otherMember := range other {
			if member == otherMember {
				match = true
				continue
			}
		}
		if !match {
			return false
		}
	}
	return true
}
