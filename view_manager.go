/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package govmomi

import (
	"github.com/vmware/govmomi/vim25/methods"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type ViewManager struct {
	c *Client
}

func (m ViewManager) ViewList() ([]types.ManagedObjectReference, error) {
	var v mo.ViewManager

	err := m.c.Properties(*m.c.ServiceContent.ViewManager, []string{"viewList"}, &v)
	if err != nil {
		return nil, err
	}

	return v.ViewList, nil
}

// Create a ContainerView object for this session.
// This returns a session object with a property that contains all objects in the container.
func (m ViewManager) CreateContainerView(container types.ManagedObjectReference, typeStr []string, recursive bool) (*types.ManagedObjectReference, error) {
	req := types.CreateContainerView{
		This:      *m.c.ServiceContent.ViewManager,
		Container: container,
		Type:      typeStr,
		Recursive: recursive,
	}

	res, err := methods.CreateContainerView(m.c, &req)
	if err != nil {
		return nil, err
	}

	if &res.Returnval == nil {
		return nil, nil
	}
	return &res.Returnval, nil
}

// Create a new InventoryView managed object for this session.
func (m ViewManager) CreateInventoryView() (*types.ManagedObjectReference, error) {
	req := types.CreateInventoryView{
		This: *m.c.ServiceContent.ViewManager,
	}

	res, err := methods.CreateInventoryView(m.c, &req)
	if err != nil {
		return nil, err
	}

	if &res.Returnval == nil {
		return nil, nil
	}
	return &res.Returnval, nil
}

// Create a ListView object for this session.
// This returns a session object with a view property that
// contains the specified objects, and which may be modified on demand.
func (m ViewManager) CreateListView(obj []types.ManagedObjectReference) (*types.ManagedObjectReference, error) {
	req := types.CreateListView{
		This: *m.c.ServiceContent.ViewManager,
		Obj:  obj,
	}

	res, err := methods.CreateListView(m.c, &req)
	if err != nil {
		return nil, err
	}

	if &res.Returnval == nil {
		return nil, nil
	}
	return &res.Returnval, nil
}

// Create a ListView object for this session from an existing view.
// This returns a session object with a view property that
// contains a copy of the set of objects in an existing view.
func (m ViewManager) CreateListViewFromView(view types.ManagedObjectReference) (*types.ManagedObjectReference, error) {
	req := types.CreateListViewFromView{
		This: *m.c.ServiceContent.ViewManager,
		View: view,
	}

	res, err := methods.CreateListViewFromView(m.c, &req)
	if err != nil {
		return nil, err
	}

	if &res.Returnval == nil {
		return nil, nil
	}
	return &res.Returnval, nil
}
