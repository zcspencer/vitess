// Copyright 2014, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package worker

import (
	"fmt"
	"testing"
	"time"

	"golang.org/x/net/context"

	"github.com/youtube/vitess/go/vt/logutil"
	"github.com/youtube/vitess/go/vt/tabletmanager/faketmclient"
	"github.com/youtube/vitess/go/vt/topo"
	"github.com/youtube/vitess/go/vt/wrangler"
	"github.com/youtube/vitess/go/vt/zktopo"
)

func TestExecuteFetchWithRetriesWithError(t *testing.T) {
	ts := zktopo.NewTestServer(t, []string{"cell1", "cell2"})
	wr := wrangler.New(logutil.NewConsoleLogger(), ts, time.Minute, time.Second)

	// fakeDestination := testlib.NewFakeTablet(t, wr, "cell1", 0,
	// 	topo.TYPE_MASTER, testlib.TabletKeyspaceShard(t, "ks", "-80"))
	// ti := topo.NewTabletInfo(&fakeDestination.Tablet, 0)
	ti := topo.NewTabletInfo(
		&topo.Tablet{
			Alias: topo.TabletAlias{
				Cell: "cell1",
				Uid:  123,
			},
			Keyspace: "ks",
			Shard:    "-80",
		},
		0 /* version */)

	ctx := context.Background()
	fakeClient := faketmclient.NewFakeTabletManagerClient()

	// gorpctmclient.GoRpcTabletManagerClient.ExecuteFetch = func(client *GoRpcTabletManagerClient, ctx context.Context, tablet *topo.TabletInfo, query string, maxRows int, wantFields, disableBinlogs bool) (*mproto.QueryResult, error) {
	// 	return nil, fmt.Errorf("fatal: this error is fatal")
	// }

	ti, err := executeFetchWithRetries(ctx, wr, ti, r, "", "fake command")

	fmt.Printf("Done! Error: %v\n", err)

}
