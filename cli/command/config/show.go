/*
 *  Copyright (c) 2021 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveAdm
 * Created Date: 2021-10-15
 * Author: Jingli Chen (Wine93)
 */

package config

import (
	"github.com/opencurve/curveadm/cli/cli"
	cliutil "github.com/opencurve/curveadm/internal/utils"
	"github.com/spf13/cobra"
)

type showOptions struct{}

func NewShowCommand(curveadm *cli.CurveAdm) *cobra.Command {
	var options showOptions

	cmd := &cobra.Command{
		Use:   "show",
		Short: "Show cluster topology",
		Args:  cliutil.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runShow(curveadm, options)
		},
		DisableFlagsInUseLine: true,
	}

	return cmd
}

func runShow(curveadm *cli.CurveAdm, options showOptions) error {
	data := curveadm.ClusterTopologyData()
	curveadm.WriteOut("%s", data)
	return nil
}
