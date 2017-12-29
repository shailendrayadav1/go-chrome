package protocol

import (
	"encoding/json"

	heapProfiler "github.com/mkenney/go-chrome/protocol/heap_profiler"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
HeapProfiler provides a namespace for the Chrome HeapProfiler protocol methods. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/
*/
var HeapProfiler = HeapProfilerProtocol{}

/*
HeapProfilerProtocol defines the HeapProfiler protocol methods.
*/
type HeapProfilerProtocol struct{}

/*
AddInspectedHeapObject enables console to refer to the node with given id via $x (see Command Line
API for more details $x functions).

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-addInspectedHeapObject
*/
func (HeapProfilerProtocol) AddInspectedHeapObject(
	socket sock.Socketer,
	params *heapProfiler.AddInspectedHeapObjectParams,
) error {
	command := sock.NewCommand("HeapProfiler.addInspectedHeapObject", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
CollectGarbage EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-collectGarbage
*/
func (HeapProfilerProtocol) CollectGarbage(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("HeapProfiler.collectGarbage", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Disable disables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-disable
*/
func (HeapProfilerProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("HeapProfiler.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables the HeapProfiler.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-enable
*/
func (HeapProfilerProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("HeapProfiler.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetHeapObjectID EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getHeapObjectId
*/
func (HeapProfilerProtocol) GetHeapObjectID(
	socket sock.Socketer,
	params *heapProfiler.GetHeapObjectIDParams,
) (*heapProfiler.GetHeapObjectIDResult, error) {
	command := sock.NewCommand("HeapProfiler.getHeapObjectID", params)
	result := &heapProfiler.GetHeapObjectIDResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetObjectByHeapObjectID EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getObjectByHeapObjectId
*/
func (HeapProfilerProtocol) GetObjectByHeapObjectID(
	socket sock.Socketer,
	params *heapProfiler.GetObjectByHeapObjectIDParams,
) (*heapProfiler.GetObjectByHeapObjectIDResult, error) {
	command := sock.NewCommand("HeapProfiler.getObjectByHeapObjectId", params)
	result := &heapProfiler.GetObjectByHeapObjectIDResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
GetSamplingProfile EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-getSamplingProfile
*/
func (HeapProfilerProtocol) GetSamplingProfile(
	socket sock.Socketer,
	params *heapProfiler.GetSamplingProfileParams,
) error {
	command := sock.NewCommand("HeapProfiler.getSamplingProfile", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StartSampling EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startSampling
*/
func (HeapProfilerProtocol) StartSampling(
	socket sock.Socketer,
	params *heapProfiler.StartSamplingParams,
) error {
	command := sock.NewCommand("HeapProfiler.startSampling", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StartTrackingHeapObjects EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-startTrackingHeapObjects
*/
func (HeapProfilerProtocol) StartTrackingHeapObjects(
	socket sock.Socketer,
	params *heapProfiler.StartTrackingHeapObjectsParams,
) error {
	command := sock.NewCommand("HeapProfiler.startTrackingHeapObjects", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopSampling EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopSampling
*/
func (HeapProfilerProtocol) StopSampling(
	socket sock.Socketer,
	params *heapProfiler.StopSamplingParams,
) error {
	command := sock.NewCommand("HeapProfiler.stopSampling", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopTrackingHeapObjects EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-stopTrackingHeapObjects
*/
func (HeapProfilerProtocol) StopTrackingHeapObjects(
	socket sock.Socketer,
	params *heapProfiler.StopTrackingHeapObjectsParams,
) error {
	command := sock.NewCommand("HeapProfiler.stopTrackingHeapObjects", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
TakeHeapSnapshot EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#method-takeHeapSnapshot
*/
func (HeapProfilerProtocol) TakeHeapSnapshot(
	socket sock.Socketer,
	params *heapProfiler.TakeHeapSnapshotParams,
) error {
	command := sock.NewCommand("HeapProfiler.takeHeapSnapshot", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnAddHeapSnapshotChunk adds a handler to the HeapProfiler.AddHeapSnapshotChunk event. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-addHeapSnapshotChunk
*/
func (HeapProfilerProtocol) OnAddHeapSnapshotChunk(
	socket sock.Socketer,
	callback func(event *heapProfiler.AddHeapSnapshotChunkEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.addHeapSnapshotChunk",
		func(response *sock.Response) {
			event := &heapProfiler.AddHeapSnapshotChunkEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnHeapStatsUpdate adds a handler to the DOM.heapStatsUpdate event. DOM.heapStatsUpdate fires if heap
objects tracking has been started then backend may send update for one or more fragments.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-heapStatsUpdate
*/
func (HeapProfilerProtocol) OnHeapStatsUpdate(
	socket sock.Socketer,
	callback func(event *heapProfiler.HeapStatsUpdateEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.heapStatsUpdate",
		func(response *sock.Response) {
			event := &heapProfiler.HeapStatsUpdateEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnLastSeenObjectID adds a handler to the DOM.LastSeenObjectID event. DOM.LastSeenObjectID fires if
heap objects tracking has been started then backend regularly sends a current value for last seen
object id and corresponding timestamp. If the were changes in the heap since last event then one or
more heapStatsUpdate events will be sent before a new lastSeenObjectId event.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-lastSeenObjectId
*/
func (HeapProfilerProtocol) OnLastSeenObjectID(
	socket sock.Socketer,
	callback func(event *heapProfiler.LastSeenObjectIDEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.lastSeenObjectID",
		func(response *sock.Response) {
			event := &heapProfiler.LastSeenObjectIDEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnReportHeapSnapshotProgress adds a handler to the DOM.ReportHeapSnapshotProgress event.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-reportHeapSnapshotProgress
*/
func (HeapProfilerProtocol) OnReportHeapSnapshotProgress(
	socket sock.Socketer,
	callback func(event *heapProfiler.ReportHeapSnapshotProgressEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.reportHeapSnapshotProgress",
		func(response *sock.Response) {
			event := &heapProfiler.ReportHeapSnapshotProgressEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}

/*
OnResetProfiles adds a handler to the HeapProfiler.ResetProfiles event. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeapProfiler/#event-resetProfiles
*/
func (HeapProfilerProtocol) OnResetProfiles(
	socket sock.Socketer,
	callback func(event *heapProfiler.ResetProfilesEvent),
) {
	handler := sock.NewEventHandler(
		"HeapProfiler.resetProfiles",
		func(response *sock.Response) {
			event := &heapProfiler.ResetProfilesEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
