package webos

import (
	"github.com/mitchellh/mapstructure"
)

// Command is the type used by tv.Command to interact with the TV.
type Command string

//https://github.com/supersaiyanmode/PyWebOSTV
const (
	// APIServiceListCommand lists the API services available on the TV.
	APIServiceListCommand Command = "ssap://api/getServiceList"

	// ApplicationManagerForegroundAppCommand returns information about the forgeground app.
	ApplicationManagerForegroundAppCommand Command = "ssap://com.webos.applicationManager/getForegroundAppInfo"

	// AudioGetVolumeCommand returns information about the TV's configured audio output volume.
	AudioGetVolumeCommand Command = "ssap://audio/getStatus"

	// AudioSetVolumeCommand sets the TV's configured audio output volume.
	AudioSetVolumeCommand Command = "ssap://audio/setVolume"

	// AudioVolumeDownCommand decrements the TV's configured audio output volume.
	AudioVolumeDownCommand Command = "ssap://audio/volumeDown"

	// AudioVolumeStatusCommand returns information about the TV's configured audio output volume.
	// Same as AudioGetVolumeCommand.
	AudioVolumeStatusCommand Command = "ssap://audio/getVolume"

	// AudioVolumeUpCommand increments the TV's configured audio output volume.
	AudioVolumeUpCommand Command = "ssap://audio/volumeUp"

	// AudioVolumeSetMuteCommand sets/toggles muting the TV's configured audio output.
	AudioVolumeSetMuteCommand Command = "ssap://audio/setMute"

	// MediaControlFastForwardCommand fast forwards the current media.
	MediaControlFastForwardCommand Command = "ssap://media.controls/fastForward"

	// MediaControlPauseCommand pauses the current media.
	MediaControlPauseCommand Command = "ssap://media.controls/pause"

	// MediaControlPlayCommand plays or resumes the current media.
	MediaControlPlayCommand Command = "ssap://media.controls/play"

	// MediaControlRewindCommand rewinds the current media.
	MediaControlRewindCommand Command = "ssap://media.controls/rewind"

	// MediaControlStopCommand stops the current media.
	MediaControlStopCommand Command = "ssap://media.controls/stop"

	// SystemLauncherCloseCommand closes a given application.
	SystemLauncherCloseCommand Command = "ssap://system.launcher/close"

	// SystemLauncherGetAppStateCommand returns information about the given application state.
	SystemLauncherGetAppStateCommand Command = "ssap://system.launcher/getAppState"

	// SystemLauncherLaunchCommand launches the given application.
	SystemLauncherLaunchCommand Command = "ssap://system.launcher/launch"

	// SystemLauncherOpenCommand opens a previously launched application.
	SystemLauncherOpenCommand Command = "ssap://system.launcher/open"

	// SystemNotificationsCreateToastCommand creates a "toast" notification.
	SystemNotificationsCreateToastCommand Command = "ssap://system.notifications/createToast"

	// SystemTurnOffCommand turns the TV off.
	SystemTurnOffCommand Command = "ssap://system/turnOff"

	// TVChannelDownCommand changes the channel down.
	TVChannelDownCommand Command = "ssap://tv/channelDown"

	// TVChannelListCommand returns information about the available channels.
	TVChannelListCommand Command = "ssap://tv/getChannelList"

	// TVChannelUpCommand changes the channel up.
	TVChannelUpCommand Command = "ssap://tv/channelUp"

	// TVCurrentChannelCommand returns information about the current channel.
	TVCurrentChannelCommand Command = "ssap://tv/getCurrentChannel"

	// TVCurrentChannelProgramCommand returns information about the current program playing on
	// the current channel.
	TVCurrentChannelProgramCommand Command = "ssap://tv/getChannelProgramInfo"

	GetPointerInputSocketCommand Command = "ssap://com.webos.service.networkinput/getPointerInputSocket"

	KeyEnterCommand Command = "ssap://com.webos.service.ime/sendEnterKey"
)

// ServiceList returns information about the available services.
func (tv *TV) ServiceList() (*ServiceList, error) {
	msg, err := tv.Command(APIServiceListCommand, nil)
	if err != nil {
		return nil, err
	}

	sl := &ServiceList{}
	err = mapstructure.Decode(msg.Payload, sl)
	return sl, err
}

// CurrentApp returns information about the current app.
func (tv *TV) CurrentApp() (*App, error) {
	msg, err := tv.Command(ApplicationManagerForegroundAppCommand, nil)
	if err != nil {
		return nil, err
	}

	a := &App{}
	err = mapstructure.Decode(msg.Payload, a)
	return a, err
}

// GetVolume returns information about the audio output volume.
func (tv *TV) GetVolume() (*VolumeStatus, error) {
	msg, err := tv.Command(AudioGetVolumeCommand, nil)
	if err != nil {
		return nil, err
	}

	v := &VolumeStatus{}
	err = mapstructure.Decode(msg.Payload, v)
	return v, err
}

// SetVolume sets the audio output volume to v.
func (tv *TV) SetVolume(v int) error {
	_, err := tv.Command(AudioSetVolumeCommand, Payload{"volume": v})
	return err
}

// VolumeDown decrements the audio output volume.
func (tv *TV) VolumeDown() error {
	_, err := tv.Command(AudioVolumeDownCommand, nil)
	return err
}

// VolumeStatus returns information about the audio output volume.
func (tv *TV) VolumeStatus() (*VolumeStatus, error) {
	msg, err := tv.Command(AudioVolumeStatusCommand, nil)
	if err != nil {
		return nil, err
	}
	v := &VolumeStatus{}
	err = mapstructure.Decode(msg.Payload, v)
	return v, err
}

// VolumeUp increments the audio output volume.
func (tv *TV) VolumeUp() error {
	_, err := tv.Command(AudioVolumeUpCommand, nil)
	return err
}

// Mute mutes the TV audio output.
func (tv *TV) Mute() error {
	_, err := tv.Command(AudioVolumeSetMuteCommand, Payload{"mute": 1})
	return err
}

// Unmute unmutes the TV audio output.
func (tv *TV) Unmute() error {
	_, err := tv.Command(AudioVolumeSetMuteCommand, Payload{"mute": 0})
	return err
}

// FastForward fast forwards the current media.
func (tv *TV) FastForward() error {
	_, err := tv.Command(MediaControlFastForwardCommand, nil)
	return err
}

// Pause pauses the current media.
func (tv *TV) Pause() error {
	_, err := tv.Command(MediaControlPauseCommand, nil)
	return err
}

// Play plays or resumes the current media.
func (tv *TV) Play() error {
	_, err := tv.Command(MediaControlPlayCommand, nil)
	return err
}

// Rewind rewinds the current media.
func (tv *TV) Rewind() error {
	_, err := tv.Command(MediaControlRewindCommand, nil)
	return err
}

// Stop stops the current media.
func (tv *TV) Stop() error {
	_, err := tv.Command(MediaControlStopCommand, nil)
	return err
}

// CloseApp closes the given app.
func (tv *TV) CloseApp(app string) error {
	_, err := tv.Command(SystemLauncherCloseCommand, Payload{"id": app})
	return err
}

// AppStatus returns information about the given app status.
func (tv *TV) AppStatus(app string) (*App, error) {
	msg, err := tv.Command(SystemLauncherGetAppStateCommand, Payload{"id": app})
	if err != nil {
		return nil, err
	}

	a := &App{}
	err = mapstructure.Decode(msg.Payload, a)
	return a, err
}

// LaunchApp launches an app.
func (tv *TV) LaunchApp(app string) error {
	_, err := tv.Command(SystemLauncherLaunchCommand, Payload{"id": app})
	return err
}

// LaunchApp launches an app with content id.
func (tv *TV) LaunchAppWithContentId(app string, contentid string) error {
	_, err := tv.Command(SystemLauncherLaunchCommand, Payload{"id": app, "contentId": contentid})
	return err
}

// OpenApp switches to a previously launched/backgrounded app.
func (tv *TV) OpenApp(app string) error {
	_, err := tv.Command(SystemLauncherOpenCommand, Payload{"id": app})
	return err
}

// Notification creates a "toast" notification.
func (tv *TV) Notification(m string) error {
	_, err := tv.Command(SystemNotificationsCreateToastCommand, Payload{"message": m})
	return err
}

// Shutdown turns the TV off.
func (tv *TV) Shutdown() error {
	_, err := tv.Command(SystemTurnOffCommand, nil)
	return err
}

// ChannelDown decrements the current channel.
func (tv *TV) ChannelDown() error {
	_, err := tv.Command(TVChannelDownCommand, nil)
	return err
}

// ChannelList returns information about available channels.
//  @todo implement a ChannelList type. This doesn't work on my TV.
func (tv *TV) ChannelList() (Message, error) {
	return tv.Command(TVChannelListCommand, nil)
}

// ChannelUp increments the current channel.
func (tv *TV) ChannelUp() error {
	_, err := tv.Command(TVChannelUpCommand, nil)
	return err
}

// CurrentChannel returns information about the current channel.
//  @todo implement a Channel type. This doesn't work on my TV.
func (tv *TV) CurrentChannel() (Message, error) {
	return tv.Command(TVCurrentChannelCommand, nil)
}

// CurrentProgram returns information about the current program
// shown on the CurrentChannel.
//  @todo implement a Program type. This doesn't work on my TV.
func (tv *TV) CurrentProgram() (Message, error) {
	return tv.Command(TVCurrentChannelProgramCommand, nil)
}

func (tv *TV) KeyUp() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	defer input.Close()

	return input.SendButton("UP")
}

func (tv *TV) KeyDown() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	defer input.Close()

	return input.SendButton("DOWN")
}

func (tv *TV) KeyLeft() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	defer input.Close()

	return input.SendButton("LEFT")
}

func (tv *TV) KeyRight() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	defer input.Close()

	return input.SendButton("RIGHT")
}

func (tv *TV) KeyBack() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	defer input.Close()

	return input.SendButton("BACK")

}

func (tv *TV) KeyHome() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("HOME")
}

func (tv *TV) KeyMenu() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("MENU")
}

func (tv *TV) KeyOK() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("ENTER")
}

func (tv *TV) KeyDash() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("DASH")
}

func (tv *TV) KeyInfo() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("INFO")
}

func (tv *TV) KeyNUM_1() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("1")
}

func (tv *TV) KeyNUM_2() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("2")
}

func (tv *TV) KeyNUM_3() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("3")
}

func (tv *TV) KeyNUM_4() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("4")
}

func (tv *TV) KeyNUM_5() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("5")
}
func (tv *TV) KeyNUM_6() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("6")
}

func (tv *TV) KeyNUM_7() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("7")
}

func (tv *TV) KeyNUM_8() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("8")
}

func (tv *TV) KeyNUM_9() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("9")
}

func (tv *TV) KeyNUM_0() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("0")
}

func (tv *TV) KeyASTERISK() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("ASTERISK")
}

func (tv *TV) KeyCC() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("CC")
}

func (tv *TV) KeyExit() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("EXIT")
}

func (tv *TV) KeyRed() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("RED")
}

func (tv *TV) KeyGreen() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("GREEN")
}

func (tv *TV) KeyYellow() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("YELLOW")
}

func (tv *TV) KeyBlue() error {
	input, err := tv.createInput()
	if err != nil {
		return err
	}
	return input.SendButton("BLUE")
}
