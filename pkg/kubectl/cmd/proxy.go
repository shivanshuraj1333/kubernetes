/*
Copyright 2014 Google Inc. All rights reserved.

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

package cmd

import (
	"io"
	"strings"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/kubectl/cmd/util"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

func (f *Factory) NewCmdProxy(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "Run a proxy to the Kubernetes API server",
		Long:  `Run a proxy to the Kubernetes API server.`,
		Run: func(cmd *cobra.Command, args []string) {
			port := util.GetFlagInt(cmd, "port")
			glog.Infof("Starting to serve on localhost:%d", port)

			clientConfig, err := f.ClientConfig(cmd)
			checkErr(err)

			staticPrefix := util.GetFlagString(cmd, "www-prefix")
			if !strings.HasSuffix(staticPrefix, "/") {
				staticPrefix += "/"
			}
			server, err := kubectl.NewProxyServer(util.GetFlagString(cmd, "www"), staticPrefix, clientConfig)
			checkErr(err)
			glog.Fatal(server.Serve(port))
		},
	}
	cmd.Flags().StringP("www", "w", "", "Also serve static files from the given directory under the specified prefix")
	cmd.Flags().StringP("www-prefix", "P", "/static/", "Prefix to serve static files under, if static file dir is specified")
	cmd.Flags().IntP("port", "p", 8001, "The port on which to run the proxy")
	return cmd
}
