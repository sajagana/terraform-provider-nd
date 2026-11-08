// Code generated;  DO NOT EDIT.

package resource_fabric_group_vxlan_aci

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"terraform-provider-nd/internal/manage/resource_fabric_common"
)

func TestResourceFabricGroupVxlanAci(t *testing.T) {
	type args struct {
		rscType  string
		rscName  string
		dataFile string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test_resource_fabric_group_vxlan_aci",
			args: args{
				rscType:  "resource",
				rscName:  "fabric_group_vxlan_aci",
				dataFile: "/examples/ndfc_payloads/data_fabric_group_vxlan_aci.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+"model_read", func(t *testing.T) {
			fileName := os.Getenv("GOPATH") + "/src/terraform-provider-ndfc" + tt.args.dataFile
			RsType := tt.args.rscType
			rscName := tt.args.rscName
			dataFromFile, err := os.ReadFile(fileName)
			if err != nil {
				t.Errorf("File read failure %v", err)
				return
			}
			modelData := resource_fabric_common.NDFCFabricCommonModel{}
			v := FabricGroupVxlanAciModel{}

			err = json.Unmarshal(dataFromFile, &modelData)
			if err != nil {
				t.Errorf("Json Unmarshal failed %s_%s: %v", RsType, rscName, err)
			}
			if err := v.SetModelData(&modelData); err != nil {
				t.Errorf("SetModelData failed %s_%s: %v", RsType, rscName, err)
			}
			t.Logf("%s_%s Read and Set ok", RsType, rscName)

			var dataFromModel []byte

			modelDataRead := v.GetModelData()

			dataFromModel, err = json.Marshal(&modelDataRead)
			if err != nil {
				t.Errorf("Json marshal failed %s_%s: %v", RsType, rscName, err)
			}
			t.Logf("%s_%s Marshall ok", RsType, rscName)

			// Normalize only bool_or_string fields to their outgoing JSON representation.
			var expectedPayload map[string]interface{}
			decoder := json.NewDecoder(bytes.NewReader(dataFromFile))
			decoder.UseNumber()
			require.NoError(t, decoder.Decode(&expectedPayload))
			normalizeBoolOrString := func(fieldPath []string) {
				parent := expectedPayload
				for _, name := range fieldPath[:len(fieldPath)-1] {
					nested, ok := parent[name].(map[string]interface{})
					if !ok {
						return
					}
					parent = nested
				}
				name := fieldPath[len(fieldPath)-1]
				switch parent[name] {
				case "true":
					parent[name] = true
				case "false":
					parent[name] = false
				}
			}

			normalizeBoolOrString([]string{"management", "securityGroupTag"})

			expectedJSON, err := json.Marshal(expectedPayload)
			require.NoError(t, err)
			require.JSONEq(t, string(expectedJSON), string(dataFromModel))

		})
	}
}
