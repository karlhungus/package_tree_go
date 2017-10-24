package main

import(
	"sync"
)

type Packager struct {
	nodes map[string][]string
	mux sync.Mutex
}


func NewPackager() Packager {
	return Packager{nodes: make(map[string][]string)}
}

func (p *Packager) Process(msg Message) string {
	switch(msg.cmd) {
	case INDEX:
		return p.index(msg)
	case QUERY:
		return p.query(msg)
	case REMOVE:
		return p.remove(msg)
	default:
		return "ERROR\n"
	}
}

func (p *Packager) index(msg Message) string {
  p.mux.Lock()
	defer p.mux.Unlock()
	for _, dep := range msg.dep {
		if _, present := p.nodes[dep]; !present {
			return "FAIL\n"
		}
	}
	p.nodes[msg.pkg] = msg.dep
	return "OK\n"
}

func (p *Packager) remove(msg Message) string {
	p.mux.Lock()
	defer p.mux.Unlock()
	_, ok := p.nodes[msg.pkg]
	if !ok {
		return "OK\n"
	}
	for _, v := range p.nodes {
		for _, dep := range v {
			if dep == msg.pkg {
				return "FAIL\n"
			}
		}
	}
	delete(p.nodes, msg.pkg)
	return "OK\n"
}

func (p *Packager) query(msg Message) string {
	p.mux.Lock()
	defer p.mux.Unlock()
	if _, ok := p.nodes[msg.pkg]; ok {
		return "OK\n"
	}
	return "FAIL\n"
}
