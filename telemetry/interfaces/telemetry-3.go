package interfaces

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"reflect"
	"testing"
	"time"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	oc "github.com/openconfig/ondatra/telemetry"
	"github.com/openconfig/ygot/ygot"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
)

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath extracts the value of the leaf InForwardedOctets from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_InForwardedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InForwardedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InForwardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath extracts the value of the leaf InForwardedPkts from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_InForwardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InForwardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_InOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_InOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_InOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_InOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_InOctetsPath extracts the value of the leaf InOctets from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_InOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_InPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_InPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_InPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_InPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_InPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/in-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_InPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_InPktsPath extracts the value of the leaf InPkts from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_InPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.InPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-discarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath extracts the value of the leaf OutDiscardedPkts from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_OutDiscardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutDiscardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-error-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutErrorPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath extracts the value of the leaf OutErrorPkts from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_OutErrorPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutErrorPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath extracts the value of the leaf OutForwardedOctets from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutForwardedOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-forwarded-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath extracts the value of the leaf OutForwardedPkts from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_OutForwardedPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutForwardedPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_OutOctetsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_OutOctetsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_OutOctetsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutOctetsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-octets to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutOctetsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_OutOctetsPath extracts the value of the leaf OutOctets from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_OutOctetsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutOctets
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPath) Lookup(t testing.TB) *oc.QualifiedUint64 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Counters", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Counters_OutPktsPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPath) Get(t testing.TB) uint64 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPathAny) Lookup(t testing.TB) []*oc.QualifiedUint64 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint64
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Counters{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Counters_OutPktsPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPathAny) Get(t testing.TB) []uint64 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint64
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Counters_OutPktsPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	w := &oc.Uint64Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Counters{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Counters", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Counters_OutPktsPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint64)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPath) Await(t testing.TB, timeout time.Duration, val uint64) *oc.QualifiedUint64 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint64) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint64 {
	t.Helper()
	c := &oc.CollectionUint64{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint64) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint64) bool) *oc.Uint64Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Counters_OutPktsPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/counters/out-pkts to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Counters_OutPktsPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Counters_OutPktsPath extracts the value of the leaf OutPkts from its parent oc.Interface_RoutedVlan_Ipv4_Counters
// and combines the update with an existing Metadata to return a *oc.QualifiedUint64.
func convertInterface_RoutedVlan_Ipv4_Counters_OutPktsPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Counters) *oc.QualifiedUint64 {
	t.Helper()
	qv := &oc.QualifiedUint64{
		Metadata: md,
	}
	val := parent.OutPkts
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_DhcpClientPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetDhcpClient())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_DhcpClientPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_DhcpClientPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_DhcpClientPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_DhcpClientPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client to the batch object.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_DhcpClientPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/dhcp-client to the batch object.
func (n *Interface_RoutedVlan_Ipv4_DhcpClientPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_DhcpClientPath extracts the value of the leaf DhcpClient from its parent oc.Interface_RoutedVlan_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_DhcpClientPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.DhcpClient
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_EnabledPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv4_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_EnabledPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv4_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_RoutedVlan_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_MtuPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_MtuPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_MtuPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_MtuPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_MtuPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_MtuPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_MtuPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu to the batch object.
func (n *Interface_RoutedVlan_Ipv4_MtuPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_MtuPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_MtuPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_MtuPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/state/mtu to the batch object.
func (n *Interface_RoutedVlan_Ipv4_MtuPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_MtuPath extracts the value of the leaf Mtu from its parent oc.Interface_RoutedVlan_Ipv4
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv4_MtuPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.Mtu
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_NeighborPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_NeighborPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Neighbor
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_Neighbor)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_NeighborPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor) bool) *oc.Interface_RoutedVlan_Ipv4_NeighborWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_NeighborWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor) bool) *oc.Interface_RoutedVlan_Ipv4_NeighborWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_NeighborPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_Neighbor) *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor to the batch object.
func (n *Interface_RoutedVlan_Ipv4_NeighborPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_NeighborPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Neighbor {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Neighbor{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_NeighborPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Neighbor) bool) *oc.Interface_RoutedVlan_Ipv4_NeighborWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_NeighborPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor to the batch object.
func (n *Interface_RoutedVlan_Ipv4_NeighborPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Neighbor_IpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Neighbor_IpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Neighbor_IpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_IpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath extracts the value of the leaf Ip from its parent oc.Interface_RoutedVlan_Ipv4_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Neighbor_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/link-layer-address to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath extracts the value of the leaf LinkLayerAddress from its parent oc.Interface_RoutedVlan_Ipv4_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Neighbor_LinkLayerAddressPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Neighbor) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.LinkLayerAddress
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Neighbor_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPath) Get(t testing.TB) oc.E_IfIp_NeighborOrigin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfIp_NeighborOrigin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Neighbor_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPathAny) Get(t testing.TB) []oc.E_IfIp_NeighborOrigin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfIp_NeighborOrigin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_NeighborOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_NeighborOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Neighbor_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_NeighborOriginWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Neighbor{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Neighbor", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Neighbor_OriginPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_NeighborOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Neighbor_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfIp_NeighborOrigin) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_NeighborOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_NeighborOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_NeighborOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_NeighborOrigin) bool) *oc.E_IfIp_NeighborOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Neighbor_OriginPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/neighbors/neighbor/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Neighbor_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Neighbor_OriginPath extracts the value of the leaf Origin from its parent oc.Interface_RoutedVlan_Ipv4_Neighbor
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfIp_NeighborOrigin.
func convertInterface_RoutedVlan_Ipv4_Neighbor_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Neighbor) *oc.QualifiedE_IfIp_NeighborOrigin {
	t.Helper()
	qv := &oc.QualifiedE_IfIp_NeighborOrigin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_ProxyArp
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_ProxyArp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_ProxyArp)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_ProxyArpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) bool) *oc.Interface_RoutedVlan_Ipv4_ProxyArpWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_ProxyArpWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_ProxyArp", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) bool) *oc.Interface_RoutedVlan_Ipv4_ProxyArpWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_ProxyArpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_ProxyArp) *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp to the batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_ProxyArp {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_ProxyArp{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_ProxyArp) bool) *oc.Interface_RoutedVlan_Ipv4_ProxyArpWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_ProxyArpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp to the batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Lookup(t testing.TB) *oc.QualifiedE_ProxyArp_Mode {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath(t, md, goStruct)
	}
	return (&oc.QualifiedE_ProxyArp_Mode{
		Metadata: md,
	}).SetVal(goStruct.GetMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Get(t testing.TB) oc.E_ProxyArp_Mode {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePathAny) Lookup(t testing.TB) []*oc.QualifiedE_ProxyArp_Mode {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_ProxyArp_Mode
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_ProxyArp", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePathAny) Get(t testing.TB) []oc.E_ProxyArp_Mode {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_ProxyArp_Mode
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_ProxyArp_Mode {
	t.Helper()
	c := &oc.CollectionE_ProxyArp_Mode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_ProxyArp_Mode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_ProxyArp_ModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_ProxyArp_Mode) bool) *oc.E_ProxyArp_ModeWatcher {
	t.Helper()
	w := &oc.E_ProxyArp_ModeWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_ProxyArp{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_ProxyArp", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_ProxyArp_Mode)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_ProxyArp_Mode) bool) *oc.E_ProxyArp_ModeWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_ProxyArp_ModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Await(t testing.TB, timeout time.Duration, val oc.E_ProxyArp_Mode) *oc.QualifiedE_ProxyArp_Mode {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_ProxyArp_Mode) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode to the batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_ProxyArp_Mode {
	t.Helper()
	c := &oc.CollectionE_ProxyArp_Mode{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_ProxyArp_Mode) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_ProxyArp_Mode) bool) *oc.E_ProxyArp_ModeWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_ProxyArp_ModePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/proxy-arp/state/mode to the batch object.
func (n *Interface_RoutedVlan_Ipv4_ProxyArp_ModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath extracts the value of the leaf Mode from its parent oc.Interface_RoutedVlan_Ipv4_ProxyArp
// and combines the update with an existing Metadata to return a *oc.QualifiedE_ProxyArp_Mode.
func convertInterface_RoutedVlan_Ipv4_ProxyArp_ModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_ProxyArp) *oc.QualifiedE_ProxyArp_Mode {
	t.Helper()
	qv := &oc.QualifiedE_ProxyArp_Mode{
		Metadata: md,
	}
	val := parent.Mode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Unnumbered
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_Unnumbered)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_UnnumberedPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) bool) *oc.Interface_RoutedVlan_Ipv4_UnnumberedWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_UnnumberedWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) bool) *oc.Interface_RoutedVlan_Ipv4_UnnumberedWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_UnnumberedPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered) *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered to the batch object.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered) bool) *oc.Interface_RoutedVlan_Ipv4_UnnumberedWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_UnnumberedPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered to the batch object.
func (n *Interface_RoutedVlan_Ipv4_UnnumberedPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetEnabled())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Unnumbered{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/state/enabled to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_EnabledPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath extracts the value of the leaf Enabled from its parent oc.Interface_RoutedVlan_Ipv4_Unnumbered
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv4_Unnumbered_EnabledPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Unnumbered) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.Enabled
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) bool) *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) bool) *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) bool) *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRefPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/interface to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath extracts the value of the leaf Interface from its parent oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_InterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Interface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Lookup(t testing.TB) *oc.QualifiedUint32 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Get(t testing.TB) uint32 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Lookup(t testing.TB) []*oc.QualifiedUint32 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint32
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Get(t testing.TB) []uint32 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint32
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	w := &oc.Uint32Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint32)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Await(t testing.TB, timeout time.Duration, val uint32) *oc.QualifiedUint32 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint32) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint32 {
	t.Helper()
	c := &oc.CollectionUint32{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint32) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint32) bool) *oc.Uint32Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv4/unnumbered/interface-ref/state/subinterface to the batch object.
func (n *Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath extracts the value of the leaf Subinterface from its parent oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef
// and combines the update with an existing Metadata to return a *oc.QualifiedUint32.
func convertInterface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef_SubinterfacePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv4_Unnumbered_InterfaceRef) *oc.QualifiedUint32 {
	t.Helper()
	qv := &oc.QualifiedUint32{
		Metadata: md,
	}
	val := parent.Subinterface
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6Path) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6Path) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6PathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6PathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6Path) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6 {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6Path(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6) bool) *oc.Interface_RoutedVlan_Ipv6Watcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6Path) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6) bool) *oc.Interface_RoutedVlan_Ipv6Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6Path(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6Path) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6) *oc.QualifiedInterface_RoutedVlan_Ipv6 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 to the batch object.
func (n *Interface_RoutedVlan_Ipv6Path) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6PathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6 {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6PathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6) bool) *oc.Interface_RoutedVlan_Ipv6Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6Path(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6 to the batch object.
func (n *Interface_RoutedVlan_Ipv6PathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Address {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_AddressPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_AddressPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Address {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Address
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Address) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_Address)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_AddressPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address) bool) *oc.Interface_RoutedVlan_Ipv6_AddressWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_AddressWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Address)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address) bool) *oc.Interface_RoutedVlan_Ipv6_AddressWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_AddressPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_Address) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_Address) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address to the batch object.
func (n *Interface_RoutedVlan_Ipv6_AddressPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_AddressPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Address {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Address{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Address) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_AddressPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address) bool) *oc.Interface_RoutedVlan_Ipv6_AddressWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_AddressPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address to the batch object.
func (n *Interface_RoutedVlan_Ipv6_AddressPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Lookup(t testing.TB) *oc.QualifiedString {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_IpPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Get(t testing.TB) string {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPathAny) Lookup(t testing.TB) []*oc.QualifiedString {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedString
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_IpPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPathAny) Get(t testing.TB) []string {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []string
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_IpPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	w := &oc.StringWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_IpPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedString)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_IpPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Await(t testing.TB, timeout time.Duration, val string) *oc.QualifiedString {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedString) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionString {
	t.Helper()
	c := &oc.CollectionString{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedString) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedString) bool) *oc.StringWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_IpPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/ip to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_IpPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_IpPath extracts the value of the leaf Ip from its parent oc.Interface_RoutedVlan_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedString.
func convertInterface_RoutedVlan_Ipv6_Address_IpPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address) *oc.QualifiedString {
	t.Helper()
	qv := &oc.QualifiedString{
		Metadata: md,
	}
	val := parent.Ip
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPath) Lookup(t testing.TB) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_OriginPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPath) Get(t testing.TB) oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPathAny) Lookup(t testing.TB) []*oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_IfIp_IpAddressOrigin
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_OriginPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPathAny) Get(t testing.TB) []oc.E_IfIp_IpAddressOrigin {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_IfIp_IpAddressOrigin
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_OriginPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	w := &oc.E_IfIp_IpAddressOriginWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_OriginPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_IfIp_IpAddressOrigin)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_OriginPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPath) Await(t testing.TB, timeout time.Duration, val oc.E_IfIp_IpAddressOrigin) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_IfIp_IpAddressOrigin {
	t.Helper()
	c := &oc.CollectionE_IfIp_IpAddressOrigin{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_IfIp_IpAddressOrigin) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_IfIp_IpAddressOrigin) bool) *oc.E_IfIp_IpAddressOriginWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_OriginPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/origin to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_OriginPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_OriginPath extracts the value of the leaf Origin from its parent oc.Interface_RoutedVlan_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedE_IfIp_IpAddressOrigin.
func convertInterface_RoutedVlan_Ipv6_Address_OriginPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address) *oc.QualifiedE_IfIp_IpAddressOrigin {
	t.Helper()
	qv := &oc.QualifiedE_IfIp_IpAddressOrigin{
		Metadata: md,
	}
	val := parent.Origin
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/prefix-length to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_PrefixLengthPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath extracts the value of the leaf PrefixLength from its parent oc.Interface_RoutedVlan_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv6_Address_PrefixLengthPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PrefixLength
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPath) Lookup(t testing.TB) *oc.QualifiedE_Address_Status {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_StatusPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPath) Get(t testing.TB) oc.E_Address_Status {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPathAny) Lookup(t testing.TB) []*oc.QualifiedE_Address_Status {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedE_Address_Status
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_StatusPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPathAny) Get(t testing.TB) []oc.E_Address_Status {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []oc.E_Address_Status
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Address_Status {
	t.Helper()
	c := &oc.CollectionE_Address_Status{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Address_Status) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_StatusPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedE_Address_Status) bool) *oc.E_Address_StatusWatcher {
	t.Helper()
	w := &oc.E_Address_StatusWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_StatusPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedE_Address_Status)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Address_Status) bool) *oc.E_Address_StatusWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_StatusPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPath) Await(t testing.TB, timeout time.Duration, val oc.E_Address_Status) *oc.QualifiedE_Address_Status {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedE_Address_Status) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionE_Address_Status {
	t.Helper()
	c := &oc.CollectionE_Address_Status{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedE_Address_Status) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedE_Address_Status) bool) *oc.E_Address_StatusWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_StatusPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/state/status to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_StatusPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_StatusPath extracts the value of the leaf Status from its parent oc.Interface_RoutedVlan_Ipv6_Address
// and combines the update with an existing Metadata to return a *oc.QualifiedE_Address_Status.
func convertInterface_RoutedVlan_Ipv6_Address_StatusPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address) *oc.QualifiedE_Address_Status {
	t.Helper()
	qv := &oc.QualifiedE_Address_Status{
		Metadata: md,
	}
	val := parent.Status
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) bool) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) bool) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup) bool) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroupWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroupPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroupPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Lookup(t testing.TB) *oc.QualifiedBool {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
	}
	return (&oc.QualifiedBool{
		Metadata: md,
	}).SetVal(goStruct.GetAcceptMode())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Get(t testing.TB) bool {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePathAny) Lookup(t testing.TB) []*oc.QualifiedBool {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedBool
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePathAny) Get(t testing.TB) []bool {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []bool
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	w := &oc.BoolWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedBool)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Await(t testing.TB, timeout time.Duration, val bool) *oc.QualifiedBool {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedBool) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionBool {
	t.Helper()
	c := &oc.CollectionBool{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedBool) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedBool) bool) *oc.BoolWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/accept-mode to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath extracts the value of the leaf AcceptMode from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedBool.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AcceptModePath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedBool {
	t.Helper()
	qv := &oc.QualifiedBool{
		Metadata: md,
	}
	val := parent.AcceptMode
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Lookup(t testing.TB) *oc.QualifiedUint16 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint16{
		Metadata: md,
	}).SetVal(goStruct.GetAdvertisementInterval())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Get(t testing.TB) uint16 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Lookup(t testing.TB) []*oc.QualifiedUint16 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint16
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Get(t testing.TB) []uint16 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint16
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	w := &oc.Uint16Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint16)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Await(t testing.TB, timeout time.Duration, val uint16) *oc.QualifiedUint16 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint16) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint16 {
	t.Helper()
	c := &oc.CollectionUint16{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint16) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint16) bool) *oc.Uint16Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/advertisement-interval to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath extracts the value of the leaf AdvertisementInterval from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint16.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_AdvertisementIntervalPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedUint16 {
	t.Helper()
	qv := &oc.QualifiedUint16{
		Metadata: md,
	}
	val := parent.AdvertisementInterval
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, md, goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/state/current-priority to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath extracts the value of the leaf CurrentPriority from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_CurrentPriorityPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.CurrentPriority
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Lookup(t testing.TB) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, false, false)
	if ok {
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
	}
	return nil
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Get(t testing.TB) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Lookup(t testing.TB) []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, false, false)
		if !ok {
			continue
		}
		qv := (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Get(t testing.TB) []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) bool {
		copy, err := ygot.DeepCopy(v.Val(t))
		if err != nil {
			t.Fatal(err)
		}
		c.Data = append(c.Data, (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: v.Metadata,
		}).SetVal(copy.(*oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking)))
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	w := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, false, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", gs, queryPath, false, false)
		return (&oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{
			Metadata: md,
		}).SetVal(gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Await(t testing.TB, timeout time.Duration, val *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking {
	t.Helper()
	c := &oc.CollectionInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) bool) *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingWatcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTrackingPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Lookup fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
// It returns nil if there is no value present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Lookup(t testing.TB) *oc.QualifiedUint8 {
	t.Helper()
	goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	md, ok := oc.Lookup(t, n, "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, true, false)
	if ok {
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
	}
	return (&oc.QualifiedUint8{
		Metadata: md,
	}).SetVal(goStruct.GetPriorityDecrement())
}

// Get fetches the value at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription,
// failing the test fatally is no value is present at the path.
// To avoid a fatal test failure, use the Lookup method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Get(t testing.TB) uint8 {
	t.Helper()
	return n.Lookup(t).Val(t)
}

// Lookup fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
// It returns an empty list if no values are present at the path.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Lookup(t testing.TB) []*oc.QualifiedUint8 {
	t.Helper()
	datapoints, queryPath := genutil.MustGet(t, n)
	datapointGroups, sortedPrefixes := genutil.BundleDatapoints(t, datapoints, uint(len(queryPath.Elem)))

	var data []*oc.QualifiedUint8
	for _, prefix := range sortedPrefixes {
		goStruct := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
		md, ok := genutil.MustUnmarshal(t, datapointGroups[prefix], oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", goStruct, queryPath, true, false)
		if !ok {
			continue
		}
		qv := convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, goStruct)
		data = append(data, qv)
	}
	return data
}

// Get fetches the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a ONCE subscription.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Get(t testing.TB) []uint8 {
	t.Helper()
	fulldata := n.Lookup(t)
	var data []uint8
	for _, full := range fulldata {
		data = append(data, full.Val(t))
	}
	return data
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

func watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, n ygot.PathStruct, duration time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	w := &oc.Uint8Watcher{}
	gs := &oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking{}
	w.W = genutil.MustWatch(t, n, nil, duration, true, func(upd []*genutil.DataPoint, queryPath *gpb.Path) (genutil.QualifiedValue, error) {
		t.Helper()
		md, _ := genutil.MustUnmarshal(t, upd, oc.GetSchema(), "Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking", gs, queryPath, true, false)
		return convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, md, gs), nil
	}, func(qualVal genutil.QualifiedValue) bool {
		val, ok := qualVal.(*oc.QualifiedUint8)
		w.LastVal = val
		return ok && predicate(val)
	})
	return w
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, n, timeout, predicate)
}

// Await observes values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// blocking until a value that is deep equal to the specified val is received
// or failing fatally if the value is not received by the specified timeout.
// To avoid a fatal failure, to wait for a generic predicate, or to make a
// non-blocking call, use the Watch method instead.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Await(t testing.TB, timeout time.Duration, val uint8) *oc.QualifiedUint8 {
	t.Helper()
	got, success := n.Watch(t, timeout, func(data *oc.QualifiedUint8) bool {
		return data.IsPresent() && reflect.DeepEqual(data.Val(t), val)
	}).Await(t)
	if !success {
		t.Fatalf("Await() at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement failed: want %v, last got %v", val, got)
	}
	return got
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// Collect starts an asynchronous collection of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription.
// Calling Await on the return Collection waits for the specified duration to elapse and returns the collected values.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Collect(t testing.TB, duration time.Duration) *oc.CollectionUint8 {
	t.Helper()
	c := &oc.CollectionUint8{}
	c.W = n.Watch(t, duration, func(v *oc.QualifiedUint8) bool {
		c.Data = append(c.Data, v)
		return false
	})
	return c
}

// Watch starts an asynchronous observation of the values at /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement with a STREAM subscription,
// evaluating each observed value with the specified predicate.
// The subscription completes when either the predicate is true or the specified duration elapses.
// Calling Await on the returned Watcher waits for the subscription to complete.
// It returns the last observed value and a boolean that indicates whether that value satisfies the predicate.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Watch(t testing.TB, timeout time.Duration, predicate func(val *oc.QualifiedUint8) bool) *oc.Uint8Watcher {
	t.Helper()
	return watch_Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t, n, timeout, predicate)
}

// Batch adds /openconfig-interfaces/interfaces/interface/routed-vlan/ipv6/addresses/address/vrrp/vrrp-group/interface-tracking/state/priority-decrement to the batch object.
func (n *Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPathAny) Batch(t testing.TB, b *oc.Batch) {
	t.Helper()
	oc.MustAddToBatch(t, b, n)
}

// convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath extracts the value of the leaf PriorityDecrement from its parent oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking
// and combines the update with an existing Metadata to return a *oc.QualifiedUint8.
func convertInterface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking_PriorityDecrementPath(t testing.TB, md *genutil.Metadata, parent *oc.Interface_RoutedVlan_Ipv6_Address_VrrpGroup_InterfaceTracking) *oc.QualifiedUint8 {
	t.Helper()
	qv := &oc.QualifiedUint8{
		Metadata: md,
	}
	val := parent.PriorityDecrement
	if !reflect.ValueOf(val).IsZero() {
		qv.SetVal(*val)
	}
	return qv
}
