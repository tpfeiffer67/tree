package tree

func (o *Node) Value(key string) (interface{}, bool) {
	if o.values != nil {
		if v, ok := o.values[key]; ok {
			return v, true
		}
	}
	return nil, false
}

func (o *Node) ValueExists(key string) bool {
	if o.values != nil {
		if _, ok := o.values[key]; ok {
			return true
		}
	}
	return false
}

func (o *Node) SetValue(key string, value interface{}) *Node {
	if o.values == nil {
		o.values = make(map[string]interface{})
	}
	o.values[key] = value
	return o
}

func (o *Node) DeleteValue(key string) {
	delete(o.values, key)
}

func (o *Node) ValueInt(key string, def int) int {
	if v, ok := o.Value(key); ok {
		if i, ok := v.(int); ok {
			return i
		}
	}
	return def
}

func (o *Node) ValueInt64(key string, def int64) int64 {
	if v, ok := o.Value(key); ok {
		if i, ok := v.(int64); ok {
			return i
		}
	}
	return def
}

func (o *Node) ValueBool(key string, def bool) bool {
	if v, ok := o.Value(key); ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return def
}

func (o *Node) ValueString(key string, def string) string {
	if v, ok := o.Value(key); ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return def
}

func (o *Node) Values() map[string]interface{}          { return o.values }
func (o *Node) SetValues(values map[string]interface{}) { o.values = values }

func NewNodeWithValue(key string, value interface{}) *Node {
	o := new(Node)
	o.SetValue(key, value)
	return o
}

func RecursiveSetValue_Index(rootNode *Node, fieldNameOfIndex, fieldNameOfParentIndex string) {
	index := 1
	rootNode.RecursiveProcessing(
		func(node *Node) {
			node.SetValue(fieldNameOfIndex, index)
			parentNode := node.Parent()
			if parentNode != nil {
				node.SetValue(fieldNameOfParentIndex, index)
			}
			index++
		},
		nil)
}

func RecursiveSetValue_HasChildren(rootNode *Node, fieldName string) {
	rootNode.RecursiveProcessing(
		func(node *Node) {
			node.SetValue(fieldName, (len(node.Children()) > 0))
		},
		nil)
}

func RecursiveSetValue_Deepness(rootNode *Node, fieldName string) {
	deepness := 0
	rootNode.RecursiveProcessing(
		func(node *Node) {
			node.SetValue(fieldName, deepness)
			deepness++
		},
		func(node *Node) {
			deepness--
		})
}

func RecursiveDeleteValue(rootNode *Node, key string) {
	rootNode.RecursiveProcessing(
		func(node *Node) {
			node.DeleteValue(key)
		},
		nil)
}
