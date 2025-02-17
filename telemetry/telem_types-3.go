package telemetry

// This file contains generated telemetry method augmentations for the
// generated path structs, which makes use of their gNMI paths for making
// ONDATRA telemetry calls.

import (
	"fmt"
	"testing"

	"github.com/openconfig/ondatra/internal/gnmigen/genutil"
	"github.com/openconfig/ygot/ygot"
)

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6LabeledUnicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_Ipv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpnWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnEvpn_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVplsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L2VpnVpls_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv4Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6MulticastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Multicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_L3VpnIpv6Unicast_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4 samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv4_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6 samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6Watcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimit, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceivedWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_SrtePolicyIpv6_PrefixLimitReceived, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_IbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AfiSafi_UseMultiplePaths_Ibgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy is a *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ApplyPolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions is a *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_AsPathOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop is a *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihopWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EbgpMultihop, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd is a *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_EnableBfdWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_EnableBfd, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling is a *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandlingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_ErrorHandling, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart is a *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestartWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_GracefulRestart, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions is a *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptionsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_LoggingOptions, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector is a *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_RouteReflectorWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_RouteReflector, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers is a *NetworkInstance_Protocol_Bgp_PeerGroup_Timers with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_Timers // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Timers sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_Timers {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Timers sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_Timers) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Timers is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_Timers samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Timers struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Timers) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_Timers samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_TimersWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Timers, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport is a *NetworkInstance_Protocol_Bgp_PeerGroup_Transport with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_Transport // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Transport sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_Transport {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_Transport sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_Transport) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Transport is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_Transport samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Transport struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_Transport) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_Transport samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_TransportWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_Transport, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths is a *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePathsWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_EbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ebgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp is a *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) SetVal(v *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp samples.
type CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp struct {
	W    *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher observes a stream of *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp samples.
type NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_IbgpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_PeerGroup_UseMultiplePaths_Ibgp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib is a *NetworkInstance_Protocol_Bgp_Rib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib) *QualifiedNetworkInstance_Protocol_Bgp_Rib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib struct {
	W    *NetworkInstance_Protocol_Bgp_RibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_RibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib samples.
type NetworkInstance_Protocol_Bgp_RibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_RibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafiWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv4Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicyWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6SrtePolicy_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6UnicastWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRibWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_LocRib_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_NeighborWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibInPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPostWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPost_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPreWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_RouteWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute samples.
type NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttributeWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AfiSafi_Ipv6Unicast_Neighbor_AdjRibOutPre_Route_UnknownAttribute, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSetWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSetWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSetWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSetWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AggregatorWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_AggregatorWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_AggregatorWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AggregatorWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_Aggregator, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4SegmentWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4SegmentWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4SegmentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_As4SegmentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_As4Segment, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegmentWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegmentWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegmentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegmentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_AsSegment, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulationWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulationWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulationWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulationWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_TunnelWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_TunnelWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_TunnelWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_TunnelWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_SubtlvWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_SubtlvWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_SubtlvWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_SubtlvWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpointWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpointWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpointWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpointWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_RemoteEndpoint, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentListWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentListWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentListWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentListWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment is a *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment) *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_SegmentWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_SegmentWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment samples.
type NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_SegmentWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_SegmentWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_AttrSet_TunnelEncapsulation_Tunnel_Subtlv_SegmentList_Segment, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_Community is a *NetworkInstance_Protocol_Bgp_Rib_Community with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_Community struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_Community // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_Community sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_Community {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_Community sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_Community) *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_Community is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_Community samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_Community struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_CommunityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_Community
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_Community) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_Community {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_CommunityWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_Community samples.
type NetworkInstance_Protocol_Bgp_Rib_CommunityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_Community
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_CommunityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_Community, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity is a *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity) Val(t testing.TB) *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity sample.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity) SetVal(v *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity) *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Bgp_Rib_ExtCommunity is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity samples.
type CollectionNetworkInstance_Protocol_Bgp_Rib_ExtCommunity struct {
	W    *NetworkInstance_Protocol_Bgp_Rib_ExtCommunityWatcher
	Data []*QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Bgp_Rib_ExtCommunity) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Bgp_Rib_ExtCommunityWatcher observes a stream of *NetworkInstance_Protocol_Bgp_Rib_ExtCommunity samples.
type NetworkInstance_Protocol_Bgp_Rib_ExtCommunityWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Bgp_Rib_ExtCommunityWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Bgp_Rib_ExtCommunity, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Igmp is a *NetworkInstance_Protocol_Igmp with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Igmp struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Igmp // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Igmp) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Igmp sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Igmp) Val(t testing.TB) *NetworkInstance_Protocol_Igmp {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Igmp sample.
func (q *QualifiedNetworkInstance_Protocol_Igmp) SetVal(v *NetworkInstance_Protocol_Igmp) *QualifiedNetworkInstance_Protocol_Igmp {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Igmp) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Igmp is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Igmp samples.
type CollectionNetworkInstance_Protocol_Igmp struct {
	W    *NetworkInstance_Protocol_IgmpWatcher
	Data []*QualifiedNetworkInstance_Protocol_Igmp
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Igmp) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Igmp {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_IgmpWatcher observes a stream of *NetworkInstance_Protocol_Igmp samples.
type NetworkInstance_Protocol_IgmpWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Igmp
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_IgmpWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Igmp, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Igmp_Global is a *NetworkInstance_Protocol_Igmp_Global with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Igmp_Global struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Igmp_Global // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Igmp_Global) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Igmp_Global sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global) Val(t testing.TB) *NetworkInstance_Protocol_Igmp_Global {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Igmp_Global sample.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global) SetVal(v *NetworkInstance_Protocol_Igmp_Global) *QualifiedNetworkInstance_Protocol_Igmp_Global {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Igmp_Global is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Igmp_Global samples.
type CollectionNetworkInstance_Protocol_Igmp_Global struct {
	W    *NetworkInstance_Protocol_Igmp_GlobalWatcher
	Data []*QualifiedNetworkInstance_Protocol_Igmp_Global
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Igmp_Global) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Igmp_Global {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Igmp_GlobalWatcher observes a stream of *NetworkInstance_Protocol_Igmp_Global samples.
type NetworkInstance_Protocol_Igmp_GlobalWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Igmp_Global
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Igmp_GlobalWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Igmp_Global, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm is a *NetworkInstance_Protocol_Igmp_Global_Ssm with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Igmp_Global_Ssm // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Igmp_Global_Ssm sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm) Val(t testing.TB) *NetworkInstance_Protocol_Igmp_Global_Ssm {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Igmp_Global_Ssm sample.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm) SetVal(v *NetworkInstance_Protocol_Igmp_Global_Ssm) *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Igmp_Global_Ssm is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Igmp_Global_Ssm samples.
type CollectionNetworkInstance_Protocol_Igmp_Global_Ssm struct {
	W    *NetworkInstance_Protocol_Igmp_Global_SsmWatcher
	Data []*QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Igmp_Global_Ssm) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Igmp_Global_SsmWatcher observes a stream of *NetworkInstance_Protocol_Igmp_Global_Ssm samples.
type NetworkInstance_Protocol_Igmp_Global_SsmWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Igmp_Global_SsmWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}

// QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping is a *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping with a corresponding timestamp.
type QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping struct {
	*genutil.Metadata
	val     *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping // val is the sample value.
	present bool
}

func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping) String() string {
	return genutil.QualifiedTypeString(q.val, q.Metadata)
}

// Val returns the value of the *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping sample, erroring out if not present.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping) Val(t testing.TB) *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping {
	t.Helper()
	if q == nil {
		t.Fatal("No value present")
	}
	if !q.present {
		pathStr, err := ygot.PathToString(q.Path)
		if err != nil {
			pathStr = fmt.Sprintf("%v", q.Path.GetElem())
		}
		t.Fatalf("No value present at path %s", pathStr)
	}
	return q.val
}

// SetVal sets the value of the *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping sample.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping) SetVal(v *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping) *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping {
	q.val = v
	q.present = true
	return q
}

// IsPresent returns true if the qualified struct contains a value.
func (q *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping) IsPresent() bool {
	return q != nil && q.present
}

// CollectionNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping is a telemetry Collection whose Await method returns a slice of *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping samples.
type CollectionNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping struct {
	W    *NetworkInstance_Protocol_Igmp_Global_Ssm_MappingWatcher
	Data []*QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping
}

// Await blocks until the telemetry collection is complete and returns the slice of values collected.
func (c *CollectionNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping) Await(t testing.TB) []*QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping {
	t.Helper()
	c.W.Await(t)
	return c.Data
}

// NetworkInstance_Protocol_Igmp_Global_Ssm_MappingWatcher observes a stream of *NetworkInstance_Protocol_Igmp_Global_Ssm_Mapping samples.
type NetworkInstance_Protocol_Igmp_Global_Ssm_MappingWatcher struct {
	W       *genutil.Watcher
	LastVal *QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping
}

// Await blocks until the Watch predicate is true or the duration elapses.
// It returns the last value received and a boolean indicating whether it satisfies the predicate.
func (w *NetworkInstance_Protocol_Igmp_Global_Ssm_MappingWatcher) Await(t testing.TB) (*QualifiedNetworkInstance_Protocol_Igmp_Global_Ssm_Mapping, bool) {
	t.Helper()
	return w.LastVal, w.W.Await(t)
}
