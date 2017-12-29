package protocol

import (
	"encoding/json"

	headlessExperimental "github.com/mkenney/go-chrome/protocol/headless_experimental"
	sock "github.com/mkenney/go-chrome/socket"

	log "github.com/sirupsen/logrus"
)

/*
HeadlessExperimental provides a namespace for the Chrome HeadlessExperimental protocol methods. The
HeadlessExperimental protocol provides experimental commands only supported in headless mode.
EXPERIMENTAL.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/
*/
var HeadlessExperimental = HeadlessExperimentalProtocol{}

/*
HeadlessExperimentalProtocol defines the HeadlessExperimental protocol methods.
*/
type HeadlessExperimentalProtocol struct{}

/*
BeginFrame sends a BeginFrame to the target and returns when the frame was completed. Optionally
captures a screenshot from the resulting frame. Requires that the target was created with enabled
BeginFrameControl.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-beginFrame
*/
func (HeadlessExperimentalProtocol) BeginFrame(
	socket sock.Socketer,
	params *headlessExperimental.BeginFrameParams,
) (*headlessExperimental.BeginFrameResult, error) {
	command := sock.NewCommand("HeadlessExperimental.beginFrame", params)
	result := &headlessExperimental.BeginFrameResult{}
	socket.SendCommand(command)

	if nil != command.Error() {
		return result, command.Error()
	}

	err := MarshalResult(command, &result)
	return result, err
}

/*
Disable disables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-disable
*/
func (HeadlessExperimentalProtocol) Disable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("HeadlessExperimental.disable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
Enable enables headless events for the target.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#method-enable
*/
func (HeadlessExperimentalProtocol) Enable(
	socket sock.Socketer,
) error {
	command := sock.NewCommand("HeadlessExperimental.enable", nil)
	socket.SendCommand(command)
	return command.Error()
}

/*
OnMainFrameReadyForScreenshots adds a handler to the
HeadlessExperimental.mainFrameReadyForScreenshots event.
HeadlessExperimental.mainFrameReadyForScreenshots fires when the main frame has first submitted a
frame to the browser. May only be fired while a BeginFrame is in flight. Before this event,
screenshotting requests may fail.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-mainFrameReadyForScreenshots
*/
func (HeadlessExperimentalProtocol) OnMainFrameReadyForScreenshots(
	socket sock.Socketer,
	callback func(event *headlessExperimental.MainFrameReadyForScreenshotsEvent),
) {
	handler := sock.NewEventHandler(
		"HeadlessExperimental.mainFrameReadyForScreenshots",
		func(response *sock.Response) {
			event := &headlessExperimental.MainFrameReadyForScreenshotsEvent{}
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
OnNeedsBeginFramesChanged adds a handler to the HeadlessExperimental.needsBeginFramesChanged event.
HeadlessExperimental.needsBeginFramesChanged fires when the target starts or stops needing
BeginFrames.

https://chromedevtools.github.io/devtools-protocol/tot/HeadlessExperimental/#event-needsBeginFramesChanged
*/
func (HeadlessExperimentalProtocol) OnNeedsBeginFramesChanged(
	socket sock.Socketer,
	callback func(event *headlessExperimental.NeedsBeginFramesChangedEvent),
) {
	handler := sock.NewEventHandler(
		"HeadlessExperimental.needsBeginFramesChanged",
		func(response *sock.Response) {
			event := &headlessExperimental.NeedsBeginFramesChangedEvent{}
			if err := json.Unmarshal([]byte(response.Params), event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
}
