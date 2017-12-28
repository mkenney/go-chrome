/*
This set of tests just locks down the structs against expectations
*/
package accessibility

import (
	"reflect"
	"testing"
)

func TestPartialAXTreeParams(t *testing.T) {
	var v PartialAXTreeParams

	if "dom.NodeID" != reflect.TypeOf(v.NodeID).String() {
		t.Errorf("PartialAXTreeParams.NodeID is expected to be of type dom.NodeID, %s found", reflect.TypeOf(v.NodeID).String())
	}
	if 0 != v.NodeID {
		t.Errorf("PartialAXTreeParams.NodeID should have a default value of 0, %v found", v.NodeID)
	}

	if "bool" != reflect.TypeOf(v.FetchRelatives).String() {
		t.Errorf("PartialAXTreeParams.FetchRelatives is expected to be of type bool, %s found", reflect.TypeOf(v.FetchRelatives).String())
	}
	if false != v.FetchRelatives {
		t.Errorf("PartialAXTreeParams.NodeID should have a default value of false, %v found", v.FetchRelatives)
	}
}

func TestPartialAXTreeResult(t *testing.T) {
	var v PartialAXTreeResult

	if "[]accessibility.AXNode" != reflect.TypeOf(v.Nodes).String() {
		t.Errorf("PartialAXTreeResult.Nodes is expected to be of type []accessibility.AXNode, %s found", reflect.TypeOf(v.Nodes).String())
	}
	if nil != v.Nodes {
		t.Errorf("PartialAXTreeResult.Nodes should have a default value of nil, %v found", v.Nodes)
	}
}

func TestAXNode(t *testing.T) {
	var v AXNode

	if "accessibility.AXNodeID" != reflect.TypeOf(v.NodeID).String() {
		t.Errorf("AXNode.Nodes is expected to be of type []accessibility.AXNode, %s found", reflect.TypeOf(v.NodeID).String())
	}
	if "" != v.NodeID {
		t.Errorf("AXNode.Nodes should have a default value of empty string, %v found", v.NodeID)
	}

	if "bool" != reflect.TypeOf(v.Ignored).String() {
		t.Errorf("AXNode.Ignored is expected to be of type bool, %s found", reflect.TypeOf(v.Ignored).String())
	}
	if false != v.Ignored {
		t.Errorf("AXNode.Ignored should have a default value of false, %v found", v.Ignored)
	}

	if "[]*accessibility.AXProperty" != reflect.TypeOf(v.IgnoredReasons).String() {
		t.Errorf("AXNode.IgnoredReasons is expected to be of type []accessibility.AXProperty, %s found", reflect.TypeOf(v.IgnoredReasons).String())
	}
	if nil != v.IgnoredReasons {
		t.Errorf("AXNode.IgnoredReasons should have a default value of nil, %v found", v.IgnoredReasons)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.Role).String() {
		t.Errorf("AXNode.Role is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.Role).String())
	}
	if nil != v.Role {
		t.Errorf("AXNode.Role should have a default value of nil, %v found", v.Role)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.Name).String() {
		t.Errorf("AXNode.Name is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.Name).String())
	}
	if nil != v.Name {
		t.Errorf("AXNode.Name should have a default value of nil, %v found", v.Name)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.Description).String() {
		t.Errorf("AXNode.Description is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.Description).String())
	}
	if nil != v.Description {
		t.Errorf("AXNode.Description should have a default value of nil, %v found", v.Description)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.Value).String() {
		t.Errorf("AXNode.Value is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.Value).String())
	}
	if nil != v.Value {
		t.Errorf("AXNode.Value should have a default value of nil, %v found", v.Value)
	}

	if "[]*accessibility.AXProperty" != reflect.TypeOf(v.Properties).String() {
		t.Errorf("AXNode.Properties is expected to be of type []*accessibility.AXProperty, %s found", reflect.TypeOf(v.Properties).String())
	}
	if nil != v.Properties {
		t.Errorf("AXNode.Properties should have a default value of nil, %v found", v.Properties)
	}

	if "[]*accessibility.AXNodeID" != reflect.TypeOf(v.ChildIDs).String() {
		t.Errorf("AXNode.ChildIDs is expected to be of type []*accessibility.AXNodeID, %s found", reflect.TypeOf(v.ChildIDs).String())
	}
	if nil != v.ChildIDs {
		t.Errorf("AXNode.ChildIDs should have a default value of nil, %v found", v.ChildIDs)
	}

	if "*dom.BackendNodeID" != reflect.TypeOf(v.BackendDOMNodeID).String() {
		t.Errorf("AXNode.BackendDOMNodeID is expected to be of type *dom.BackendNodeID, %s found", reflect.TypeOf(v.BackendDOMNodeID).String())
	}
	if nil != v.BackendDOMNodeID {
		t.Errorf("AXNode.BackendDOMNodeID should have a default value of nil, %v found", v.BackendDOMNodeID)
	}
}

func TestAXNodeID(t *testing.T) {
	var v AXNodeID

	if "string" != reflect.TypeOf(v).Kind().String() {
		t.Errorf("AXNodeID is expected to be of type string, %s found", reflect.TypeOf(v).Kind().String())
	}
	if "" != v {
		t.Errorf("AXNodeID should have a default value of empty string, %v found", v)
	}
}

func TestAXProperty(t *testing.T) {
	var v AXProperty

	if "accessibility.AXPropertyName" != reflect.TypeOf(v.Name).String() {
		t.Errorf("AXProperty.Name is expected to be of type *dom.BackendNodeID, %s found", reflect.TypeOf(v.Name).String())
	}
	if "" != v.Name {
		t.Errorf("AXProperty.Name should have a default value of empty string, %v found", v.Name)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.Value).String() {
		t.Errorf("AXProperty.Value is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.Value).String())
	}
	if nil != v.Value {
		t.Errorf("AXProperty.Value should have a default value of nil, %v found", v.Value)
	}
}

func TestAXPropertyName(t *testing.T) {
	var v AXPropertyName

	if "string" != reflect.TypeOf(v).Kind().String() {
		t.Errorf("AXPropertyName is expected to be of type string, %s found", reflect.TypeOf(v).Kind().String())
	}
	if "" != v {
		t.Errorf("AXPropertyName should have a default value of empty string, %v found", v)
	}
}

func TestAXRelatedNode(t *testing.T) {
	var v AXRelatedNode

	if "dom.BackendNodeID" != reflect.TypeOf(v.BackendDOMNodeID).String() {
		t.Errorf("AXRelatedNode.BackendDOMNodeID is expected to be of type dom.BackendNodeID, %s found", reflect.TypeOf(v.BackendDOMNodeID).String())
	}
	if 0 != v.BackendDOMNodeID {
		t.Errorf("AXRelatedNode.BackendDOMNodeID should have a default value of 0, %v found", v.BackendDOMNodeID)
	}

	if "string" != reflect.TypeOf(v.IDRef).Kind().String() {
		t.Errorf("AXRelatedNode.IDRef is expected to be of type string, %s found", reflect.TypeOf(v.IDRef).String())
	}
	if "" != v.IDRef {
		t.Errorf("AXRelatedNode.IDRef should have a default value of empty string, %v found", v.IDRef)
	}

	if "string" != reflect.TypeOf(v.Text).Kind().String() {
		t.Errorf("AXRelatedNode.Text is expected to be of type string, %s found", reflect.TypeOf(v.Text).String())
	}
	if "" != v.Text {
		t.Errorf("AXRelatedNode.Text should have a default value of empty string, %v found", v.Text)
	}
}

func TestAXValue(t *testing.T) {
	var v AXValue

	if "accessibility.AXValueType" != reflect.TypeOf(v.Type).String() {
		t.Errorf("AXValue.Type is expected to be of type accessibility.AXValueType, %s found", reflect.TypeOf(v.Type).String())
	}
	if "" != v.Type {
		t.Errorf("AXValue.Type should have a default value of 0, %v found", v.Type)
	}

	if nil != reflect.TypeOf(v.Value) {
		t.Errorf("AXValue.Type is expected to be of type interface{}, %v found", reflect.TypeOf(v.Value))
	}
	if nil != v.Value {
		t.Errorf("AXValue.Type should have a default value of nil, %v found", v.Value)
	}

	if "[]*accessibility.AXRelatedNode" != reflect.TypeOf(v.RelatedNodes).String() {
		t.Errorf("AXValue.RelatedNodes is expected to be of type []*accessibility.AXRelatedNode, %v found", reflect.TypeOf(v.RelatedNodes))
	}
	if nil != v.RelatedNodes {
		t.Errorf("AXValue.RelatedNodes should have a default value of nil, %v found", v.RelatedNodes)
	}

	if "[]*accessibility.AXValueSource" != reflect.TypeOf(v.Sources).String() {
		t.Errorf("AXValue.Sources is expected to be of type []*accessibility.AXValueSource, %v found", reflect.TypeOf(v.Sources))
	}
	if nil != v.Sources {
		t.Errorf("AXValue.Sources should have a default value of nil, %v found", v.Sources)
	}
}

func TestAXValueNativeSourceType(t *testing.T) {
	var v AXValueNativeSourceType

	if "string" != reflect.TypeOf(v).Kind().String() {
		t.Errorf("AXValueNativeSourceType is expected to be of type string, %s found", reflect.TypeOf(v).Kind().String())
	}
	if "" != v {
		t.Errorf("AXValueNativeSourceType should have a default value of empty string, %v found", v)
	}
}

func TestAXValueSource(t *testing.T) {
	var v AXValueSource

	if "accessibility.AXValueSourceType" != reflect.TypeOf(v.Type).String() {
		t.Errorf("AXValueSource.Type is expected to be of type accessibility.AXValueSourceType, %s found", reflect.TypeOf(v.Type).String())
	}
	if "" != v.Type {
		t.Errorf("AXValueSource.Type should have a default value of empty string, %v found", v.Type)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.Value).String() {
		t.Errorf("AXValueSource.Value is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.Value).String())
	}
	if nil != v.Value {
		t.Errorf("AXValueSource.Value should have a default value of nil, %v found", v.Value)
	}

	if "string" != reflect.TypeOf(v.Attribute).String() {
		t.Errorf("AXValueSource.Attribute is expected to be of type string, %s found", reflect.TypeOf(v.Attribute).String())
	}
	if "" != v.Attribute {
		t.Errorf("AXValueSource.Attribute should have a default value of empty string, %v found", v.Attribute)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.AttributeValue).String() {
		t.Errorf("AXValueSource.AttributeValue is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.AttributeValue).String())
	}
	if nil != v.AttributeValue {
		t.Errorf("AXValueSource.AttributeValue should have a default value of nil, %v found", v.AttributeValue)
	}

	if "bool" != reflect.TypeOf(v.Superseded).String() {
		t.Errorf("AXValueSource.Superseded is expected to be of type bool, %s found", reflect.TypeOf(v.Superseded).String())
	}
	if false != v.Superseded {
		t.Errorf("AXValueSource.Superseded should have a default value of false, %v found", v.Superseded)
	}

	if "accessibility.AXValueNativeSourceType" != reflect.TypeOf(v.NativeSource).String() {
		t.Errorf("AXValueSource.NativeSource is expected to be of type accessibility.AXValueNativeSourceType, %s found", reflect.TypeOf(v.NativeSource).String())
	}
	if "" != v.NativeSource {
		t.Errorf("AXValueSource.NativeSource should have a default value of empty string, %v found", v.NativeSource)
	}

	if "*accessibility.AXValue" != reflect.TypeOf(v.NativeSourceValue).String() {
		t.Errorf("AXValueSource.NativeSourceValue is expected to be of type *accessibility.AXValue, %s found", reflect.TypeOf(v.NativeSourceValue).String())
	}
	if nil != v.NativeSourceValue {
		t.Errorf("AXValueSource.NativeSourceValue should have a default value of nil, %v found", v.NativeSourceValue)
	}

	if "bool" != reflect.TypeOf(v.Invalid).String() {
		t.Errorf("AXValueSource.Invalid is expected to be of type bool, %s found", reflect.TypeOf(v.Invalid).String())
	}
	if false != v.Invalid {
		t.Errorf("AXValueSource.Invalid should have a default value of false, %v found", v.Invalid)
	}

	if "string" != reflect.TypeOf(v.InvalidReason).String() {
		t.Errorf("AXValueSource.InvalidReason is expected to be of type string, %s found", reflect.TypeOf(v.InvalidReason).String())
	}
	if "" != v.InvalidReason {
		t.Errorf("AXValueSource.InvalidReason should have a default value of empty string, %v found", v.InvalidReason)
	}

}

func TestAXValueSourceType(t *testing.T) {
	var v AXValueSourceType

	if "string" != reflect.TypeOf(v).Kind().String() {
		t.Errorf("AXValueSourceType is expected to be of type string, %s found", reflect.TypeOf(v).Kind().String())
	}
	if "" != v {
		t.Errorf("AXValueSourceType should have a default value of empty string, %v found", v)
	}
}

func TestAXValueType(t *testing.T) {
	var v AXValueType

	if "string" != reflect.TypeOf(v).Kind().String() {
		t.Errorf("AXValueType is expected to be of type string, %s found", reflect.TypeOf(v).Kind().String())
	}
	if "" != v {
		t.Errorf("AXValueType should have a default value of empty string, %v found", v)
	}
}
