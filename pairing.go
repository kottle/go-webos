package webos

// pairPrompt returns a Payload necessary to pair with the TV using
// the PROMPT method.
func pairPrompt() Payload {
	return Payload{
		"forcePairing": false,
		"pairingType":  "PROMPT",
		"manifest": map[string]interface{}{
			"manifestVersion": 1,
			"permissions": []string{
				"LAUNCH",
				"LAUNCH_WEBAPP",
				"APP_TO_APP",
				"CONTROL_AUDIO",
				"CONTROL_INPUT_MEDIA_PLAYBACK",
				"UPDATE_FROM_REMOTE_APP",
				"CONTROL_POWER",
				"READ_INSTALLED_APPS",
				"CONTROL_DISPLAY",
				"CONTROL_INPUT_JOYSTICK",
				"CONTROL_INPUT_MEDIA_RECORDING",
				"CONTROL_INPUT_TV",
				"READ_INPUT_DEVICE_LIST",
				"READ_NETWORK_STATE",
				"READ_TV_CHANNEL_LIST",
				"WRITE_NOTIFICATION_TOAST",
				"CONTROL_BLUETOOTH",
				"CHECK_BLUETOOTH_DEVICE",
				"CONTROL_USER_INFO",
				"CONTROL_TIMER_INFO",
				"READ_SETTINGS",
				"CONTROL_TV_SCREEN",
				"CONTROL_INPUT_TEXT",
				"CONTROL_MOUSE_AND_KEYBOARD",
				"READ_CURRENT_CHANNEL",
				"READ_RUNNING_APPS",
			},
		},
	}
}

// pairPIN returns a Payload necessary to pair with the TV using
// the PIN method.
//  @todo implement PIN pairing.
func pairPIN() Payload {
	panic("not implemented")
}
