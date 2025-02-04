// Copyright 2022 Linkall Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cs

import (
	// standard libraries.
	"encoding/json"
	"fmt"
	"strconv"

	// third-party libraries.
	"github.com/spf13/cobra"

	// this project.
	"github.com/vanus-labs/vanus/raft/raftpb"
	"github.com/vanus-labs/vanus/vsrepair/meta"
)

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cs BlockID",
		Short: "Get ConfState of a specific Block.",
		Run:   get,
	}
	cmd.Flags().StringVar(&volumePath, "volume", "", "volume path")
	return cmd
}

type getResult struct {
	BlockID   uint64           `json:"BlockID"`
	ConfState raftpb.ConfState `json:"ConfState"`
}

func get(cmd *cobra.Command, args []string) {
	db, err := meta.Open(volumePath, meta.ReadOnly())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for _, arg := range args {
		id, err := strconv.ParseUint(arg, 0, 0)
		if err != nil {
			panic(err)
		}

		cs, err := db.GetConfState(id)
		if err != nil {
			panic(err)
		}

		jsonResult, err := json.MarshalIndent(getResult{
			BlockID:   id,
			ConfState: cs,
		}, "", "  ")
		if err != nil {
			panic(err)
		}

		fmt.Println(string(jsonResult))
	}
}
