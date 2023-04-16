package netflow

import (
	"fmt"
	"sync"
)

var (
	ErrorTemplateNotFound = fmt.Errorf("Error template not found")
)

type FlowBaseTemplateSet map[uint16]map[uint32]map[uint16]interface{}

// Store interface that allows storing and retrieving template data
type NetFlowTemplateSystem interface {
	GetTemplate(version uint16, obsDomainId uint32, templateId uint16) (interface{}, error)
	AddTemplate(version uint16, obsDomainId uint32, template interface{}) error
}

func (ts *BasicTemplateSystem) GetTemplates() map[uint16]map[uint32]map[uint16]interface{} {
	ts.templateslock.RLock()
	tmp := ts.templates
	ts.templateslock.RUnlock()
	return tmp
}

func (ts *BasicTemplateSystem) AddTemplate(version uint16, obsDomainId uint32, template interface{}) error {
	ts.templateslock.Lock()
	defer ts.templateslock.Unlock()
	_, exists := ts.templates[version]
	if !exists {
		ts.templates[version] = make(map[uint32]map[uint16]interface{})
	}
	_, exists = ts.templates[version][obsDomainId]
	if !exists {
		ts.templates[version][obsDomainId] = make(map[uint16]interface{})
	}
	var templateId uint16
	switch templateIdConv := template.(type) {
	case IPFIXOptionsTemplateRecord:
		templateId = templateIdConv.TemplateId
	case NFv9OptionsTemplateRecord:
		templateId = templateIdConv.TemplateId
	case TemplateRecord:
		templateId = templateIdConv.TemplateId
	}
	ts.templates[version][obsDomainId][templateId] = template
	return nil
}

func (ts *BasicTemplateSystem) GetTemplate(version uint16, obsDomainId uint32, templateId uint16) (interface{}, error) {
	ts.templateslock.RLock()
	defer ts.templateslock.RUnlock()
	templatesVersion, okver := ts.templates[version]
	if okver {
		templatesObsDom, okobs := templatesVersion[obsDomainId]
		if okobs {
			template, okid := templatesObsDom[templateId]
			if okid {
				return template, nil
			}
		}
	}
	return nil, ErrorTemplateNotFound
}

type BasicTemplateSystem struct {
	templates     FlowBaseTemplateSet
	templateslock *sync.RWMutex
}

// Creates a basic store for NetFlow and IPFIX templates.
// Everyting is stored in memory.
func CreateTemplateSystem() NetFlowTemplateSystem {
	ts := &BasicTemplateSystem{
		templates:     make(FlowBaseTemplateSet),
		templateslock: &sync.RWMutex{},
	}
	return ts
}
