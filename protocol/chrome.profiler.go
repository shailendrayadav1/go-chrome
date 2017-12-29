package protocol

import (
	"encoding/json"

	profiler "github.com/mkenney/go-chrome/protocol/profiler"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
Profiler provides a namespace for the Chrome Profiler protocol methods.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/
*/
var Profiler = ProfilerProtocol{}

/*
ProfilerProtocol defines the Profiler protocol methods.
*/
type ProfilerProtocol struct{}

/*
Disable disables profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-disable
*/
func (ProfilerProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Profiler.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-enable
*/
func (ProfilerProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Profiler.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
GetBestEffortCoverage collects coverage data for the current isolate. The coverage data may be
incomplete due to garbage collection.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-getBestEffortCoverage
*/
func (ProfilerProtocol) GetBestEffortCoverage(
	socket sock.Socketer,
) (*profiler.GetBestEffortCoverageResult, error) {
	command := sock.NewCommand("Profiler.getBestEffortCoverage", nil)
	result := &profiler.GetBestEffortCoverageResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
SetSamplingInterval changes CPU profiler sampling interval. Must be called before CPU profiles
recording started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-setSamplingInterval
*/
func (ProfilerProtocol) SetSamplingInterval(
	socket sock.Socketer,
	params *profiler.SetSamplingIntervalParams,
) error {
	command := sock.NewCommand("Profiler.setSamplingInterval", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
Start starts profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-start
*/
func (ProfilerProtocol) Start(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Profiler.start", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StartPreciseCoverage enable precise code coverage. Coverage data for JavaScript executed before
enabling precise code coverage may be incomplete. Enabling prevents running optimized code and
resets execution counters.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startPreciseCoverage
*/
func (ProfilerProtocol) StartPreciseCoverage(
	socket sock.Socketer,
	params *profiler.StartPreciseCoverageParams,
) error {
	command := sock.NewCommand("Profiler.startPreciseCoverage", params)
	socket.SendCommand(command)
	return command.Error()
}

/*
StartTypeProfile enables type profile. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-startTypeProfile
*/
func (ProfilerProtocol) StartTypeProfile(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Profiler.startTypeProfile", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Stop stops profiling.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stop
*/
func (ProfilerProtocol) Stop(
	socket sock.Socketer,
) (*profiler.StopResult, error) {
	command := sock.NewCommand("Profiler.stop", nil)
	result := &profiler.StopResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
StopPreciseCoverage disable precise code coverage. Disabling releases unnecessary execution count
records and allows executing optimized code.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopPreciseCoverage
*/
func (ProfilerProtocol) StopPreciseCoverage(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Profiler.stopPreciseCoverage", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
StopTypeProfile disables type profile. Disabling releases type profile data collected so far.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-stopTypeProfile
*/
func (ProfilerProtocol) StopTypeProfile(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("Profiler.stopTypeProfile", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
TakePreciseCoverage collects coverage data for the current isolate, and resets execution counters.
Precise code coverage needs to have started.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takePreciseCoverage
*/
func (ProfilerProtocol) TakePreciseCoverage(
	socket sock.Socketer,
) (*profiler.TakePreciseCoverageResult, error) {
	command := sock.NewCommand("Profiler.takePreciseCoverage", nil)
	result := &profiler.TakePreciseCoverageResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
TakeTypeProfile collect type profile. EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#method-takeTypeProfile
*/
func (ProfilerProtocol) TakeTypeProfile(
	socket sock.Socketer,
) (*profiler.TakeTypeProfileResult, error) {
	command := sock.NewCommand("Profiler.takeTypeProfile", nil)
	result := &profiler.TakeTypeProfileResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
OnConsoleProfileFinished adds a handler to the Profiler.consoleProfileFinished event.
Profiler.consoleProfileFinished fires when profile recording finishes.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileFinished
*/
func (ProfilerProtocol) OnConsoleProfileFinished(
	socket sock.Socketer,
	callback func(event *profiler.ConsoleProfileFinishedEvent),
) {
	handler := sock.NewEventHandler(
		"Profiler.consoleProfileFinished",
		func(response *sock.Response) {
			event := &profiler.ConsoleProfileFinishedEvent{}
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
OnConsoleProfileStarted adds a handler to the Profiler.consoleProfileStarted event.
Profiler.consoleProfileStarted fires when new profile recording is started using console.profile()
call.

https://chromedevtools.github.io/devtools-protocol/tot/Profiler/#event-consoleProfileStarted
*/
func (ProfilerProtocol) OnConsoleProfileStarted(
	socket sock.Socketer,
	callback func(event *profiler.ConsoleProfileStartedEvent),
) {
	handler := sock.NewEventHandler(
		"Profiler.consoleProfileStarted",
		func(response *sock.Response) {
			event := &profiler.ConsoleProfileStartedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
