package jsonstruct

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
)

var expextedStatus CmsStatus = CmsStatus{
	Cluster: &Cluster{
		ClusterName: "loadtest-cluster",
		LocalNode: &LocalNode{
			Id:                "2tLuZiq-SFix7fPPlpgd3g",
			HostName:          "loadtest-appserver1",
			Master:            "true",
			NumberOfNodesSeen: 5},
		// Members: map[string]interface{}{"DKzIAQZPTOeMq47l0OL-5g": map[string]interface{}{"version": "0.90.5", "master": "false", "name": "loadtest-appserver3", "id": "DKzIAQZPTOeMq47l0OL-5g", "address": "inet[/192.168.4.3:9300]"}, "SS9qzT-sRuSGYoamnJ5AIw": map[string]interface{}{"id": "SS9qzT-sRuSGYoamnJ5AIw", "address": "inet[/192.168.4.2:9300]", "version": "0.90.5", "master": "false", "name": "loadtest-appserver2"}, "2tLuZiq-SFix7fPPlpgd3g": map[string]interface{}{"address": "inet[/192.168.4.1:9300]", "version": "0.90.5", "master": "true", "name": "loadtest-appserver1", "id": "2tLuZiq-SFix7fPPlpgd3g"}, "wu13yY45QBKdu0Ggq1RMJw": map[string]interface{}{"master": "false", "name": "10.0.6.82", "id": "wu13yY45QBKdu0Ggq1RMJw", "address": "inet[/192.168.4.5:9300]", "version": "0.90.5"}, "VigLwoSvTuKl-LF51KVghw": map[string]interface{}{"name": "loadtest-appserver4", "id": "VigLwoSvTuKl-LF51KVghw", "address": "inet[/192.168.4.4:9300]", "version": "0.90.5", "master": "false"}},
	},
	Os:      &Os{Name: "Linux", Version: "3.2.0-34-generic", Cores: 8, LodeAverage: 0},
	Jvm:     &Jvm{Name: "Java HotSpot(TM) 64-Bit Server VM", Vendor: "Oracle Corporation", Version: "24.0-b56", StartTime: 1.380721284659e+12, UpTime: 331855},
	Memory:  &Memory{Heap: &MemoryConsumption{Init: 6.442450944e+09, Max: 6.313607168e+09, Committed: 6.313607168e+09, Used: 6.60364184e+08}, NonHeap: &MemoryConsumption{Init: 2.4313856e+07, Max: 3.18767104e+08, Committed: 8.1133568e+07, Used: 7.9803104e+07}},
	Gc:      &Gc{CollectionTime: 231, CollectionCount: 2},
	Index:   &Index{Status: "GREEN", ActiveShards: 16, ActivePrimaryShards: 4, ActiveReplicas: 12, UnassignedShards: 0, RelocatingShards: 0, InitializingShards: 0, Documents: "0", PrimaryShardsStoreSize: "336b", TotalStoreSize: "1.2kb"},
	Product: &Product{Name: "Enonic CMS", Version: "4.7.13", Edition: "enterprise"}}

func loadCmsJson(t *testing.T) []byte {
	bytes, err := ioutil.ReadFile("../enoniccms4.7.json")
	if err != nil {
		t.Error("Could not read CMS json file")
	}
	return bytes
}

func TestCmsMarshalling(t *testing.T) {
	rawJson := loadCmsJson(t)
	var cmsStatus CmsStatus
	if err := json.Unmarshal(rawJson, &cmsStatus); err != nil {
		t.Error("Could not Unmarshal json to struct")
	}

	if !reflect.DeepEqual(cmsStatus, expextedStatus) {
		t.FailNow()
	}
}

func TestImportantFields(t *testing.T) {
	rawJson := loadCmsJson(t)
	var cmsStatus CmsStatus
	if err := json.Unmarshal(rawJson, &cmsStatus); err != nil {
		t.Logf("Error: %v", err)
		t.Error("Could not Unmarshal json to struct")
	}

	if cmsStatus.Cluster.LocalNode.HostName != "loadtest-appserver1" {
		t.Error("ClusterName not correct")
	}

	if cmsStatus.Index.Status != "GREEN" {
		t.Error("Index.Status not correct")
	}

	if cmsStatus.Cluster.LocalNode.Master != "true" {
		t.Error("Cluster.LocalNode.Master not correct")
	}

	if cmsStatus.Cluster.LocalNode.NumberOfNodesSeen != 5 {
		t.Error("Cluster.LocalNode.NumberOfNodesSeen not correct")
	}

	if cmsStatus.Jvm.UpTime != 331855 {
		t.Error("Jvm.Uptime not correct")
	}

	if cmsStatus.Product.Version != "4.7.13" {
		t.Error("Product.Version not correct")
	}
}
