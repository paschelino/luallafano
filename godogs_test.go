package main

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

func aDirectoryHierarchyExists(arg1 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}

func aDirectoryHierarchyIsBeingCreated(arg1 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}

func aLuallafanoPromptIsShown() error {
	return godog.ErrPending
}

func aLuallafanoPromptIsShownToEnterTheShortNameForTheCommand() error {
	return godog.ErrPending
}

func iEnterIntoTheLuallafanoPrompt(arg1 string) error {
	return godog.ErrPending
}

func iEnterIntoTheLuallfanoPrompt(arg1 string) error {
	return godog.ErrPending
}

func iEnteredAComplicatedCommandOnTheCLI(arg1 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}

func iType(arg1 string) error {
	return godog.ErrPending
}

func iTypeTheCommandAskingLuallafanoToRememberTheCommand(arg1 string) error {
	return godog.ErrPending
}

func iUseLuallafanoForTheFirstTime() error {
	return godog.ErrPending
}

func luallafanoInformsTheUser(arg1 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}

func theFileIsCreatedWithThisContent(arg1 string, arg2 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}

func theOutputMatches(arg1 *messages.PickleStepArgument_PickleDocString) error {
	return godog.ErrPending
}

func theSymlinkToIsCreated(arg1, arg2 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a directory hierarchy exists:$`, aDirectoryHierarchyExists)
	ctx.Step(`^a directory hierarchy is being created:$`, aDirectoryHierarchyIsBeingCreated)
	ctx.Step(`^a luallafano prompt is shown$`, aLuallafanoPromptIsShown)
	ctx.Step(`^a luallafano prompt is shown to enter the short name for the command$`, aLuallafanoPromptIsShownToEnterTheShortNameForTheCommand)
	ctx.Step(`^I enter "([^"]*)" into the luallafano prompt$`, iEnterIntoTheLuallafanoPrompt)
	ctx.Step(`^I enter "([^"]*)" into the luallfano prompt$`, iEnterIntoTheLuallfanoPrompt)
	ctx.Step(`^I entered a complicated command on the CLI:$`, iEnteredAComplicatedCommandOnTheCLI)
	ctx.Step(`^I type "([^"]*)"$`, iType)
	ctx.Step(`^I type the command "([^"]*)" asking luallafano to remember the command$`, iTypeTheCommandAskingLuallafanoToRememberTheCommand)
	ctx.Step(`^I use Luallafano for the first time$`, iUseLuallafanoForTheFirstTime)
	ctx.Step(`^Luallafano informs the user:$`, luallafanoInformsTheUser)
	ctx.Step(`^the file "([^"]*)" is created with this content:$`, theFileIsCreatedWithThisContent)
	ctx.Step(`^the output matches:$`, theOutputMatches)
	ctx.Step(`^the symlink "([^"]*)" to "([^"]*)" is created$`, theSymlinkToIsCreated)
}
