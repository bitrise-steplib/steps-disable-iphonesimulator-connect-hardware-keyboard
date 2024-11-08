# Disable iphonesimulator ConnectHardwareKeyboard preference

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/steps-disable-iphonesimulator-connect-hardware-keyboard?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/steps-disable-iphonesimulator-connect-hardware-keyboard/releases)

This Step disables ConnectHardwareKeyboard preference for all iphonesimulator devices.


<details>
<summary>Description</summary>

<nil>
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `iphonesimulator_preferences_pth` | The path of the iphonesimulator preferences file. | required | `~/Library/Preferences/com.apple.iphonesimulator.plist` |
| `verbose` | Print verbose information. | required | `no` |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/steps-disable-iphonesimulator-connect-hardware-keyboard/pulls) and [issues](https://github.com/bitrise-steplib/steps-disable-iphonesimulator-connect-hardware-keyboard/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
