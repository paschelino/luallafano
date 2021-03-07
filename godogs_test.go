package main

import (
	"bytes"
	"fmt"
	"strings"

	"os"
	"os/exec"

	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
	"github.com/thanhpk/randstr"

	t "text/template"
)

var luallaFanoSimpleDir = ".luallafano"
var userCommand string
var testHomeParentDir = defineTestHomeParentDir()
var testHomeDir = fmt.Sprintf("%s/%s", testHomeParentDir, randstr.Hex(16))
var luallaFanoDir = fmt.Sprintf("%s/%s", testHomeDir, luallaFanoSimpleDir)

func defineTestHomeParentDir() string {
	dirName, err := os.Getwd()
	if err != nil { panic("Something is utterly wrong. I cannot find the current working directory!") }
	return fmt.Sprintf("%s/test_home_dirs", dirName)
}

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

func iEnteredAComplicatedCommandOnTheCLI(arg1 *messages.PickleStepArgument_PickleDocString) error {
	if arg1.Content == "" {
		return fmt.Errorf("Cannt operate on a command without receiving one!")
	}
	userCommand = arg1.Content
	return nil
}

//cmd := exec.Command(arg1.GetContent())

func iType(arg1 string) error {
	return godog.ErrPending
}

func iTypeTheCommandAskingLuallafanoToRememberTheCommand(lrc string) error {
	shellCommands := struct {
		Command string
		Lrc string
	}{
		userCommand,
		lrc,
	}
	commandSequence, err := t.New("commandSequence").Parse("{{.Command}}\n{{.Lrc}}\n")
	if err != nil { panic(err) }
	var script bytes.Buffer
	err = commandSequence.Execute(&script, shellCommands)
	if err != nil { panic(err) }
	commandWrapper := exec.Command("sh")
	commandWrapper.Stdin = strings.NewReader(script.String())
	var out bytes.Buffer
	var errOut bytes.Buffer
	commandWrapper.Stdout = &out
	commandWrapper.Stderr = &errOut
	err = commandWrapper.Run()
	if err != nil {
		fmt.Println(out.String())
		fmt.Println(errOut.String())
		panic(err) 
	}

	return godog.ErrPending
}

func iUseLuallafanoForTheFirstTime() error {
	_, err := os.Stat(luallaFanoDir)
	if err != nil { return nil }
	return fmt.Errorf("Luallafano was already used. The directory '%s' already exists: %s", luallaFanoDir, err)
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

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.AfterSuite(func() { 
		os.RemoveAll(testHomeParentDir)
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a directory hierarchy exists:$`, aDirectoryHierarchyExists)
	ctx.Step(`^a directory hierarchy is being created:$`, aDirectoryHierarchyIsBeingCreated)
	ctx.Step(`^a luallafano prompt is shown$`, aLuallafanoPromptIsShown)
	ctx.Step(`^a luallafano prompt is shown to enter the short name for the command$`, aLuallafanoPromptIsShownToEnterTheShortNameForTheCommand)
	ctx.Step(`^I enter "([^"]*)" into the luallafano prompt$`, iEnterIntoTheLuallafanoPrompt)
	ctx.Step(`^I entered a complicated command on the CLI:$`, iEnteredAComplicatedCommandOnTheCLI)
	ctx.Step(`^I type "([^"]*)"$`, iType)
	ctx.Step(`^I type the command "([^"]*)" asking luallafano to remember the command$`, iTypeTheCommandAskingLuallafanoToRememberTheCommand)
	ctx.Step(`^I use Luallafano for the first time$`, iUseLuallafanoForTheFirstTime)
	ctx.Step(`^Luallafano informs the user:$`, luallafanoInformsTheUser)
	ctx.Step(`^the file "([^"]*)" is created with this content:$`, theFileIsCreatedWithThisContent)
	ctx.Step(`^the output matches:$`, theOutputMatches)
	ctx.Step(`^the symlink "([^"]*)" to "([^"]*)" is created$`, theSymlinkToIsCreated)

	ctx.BeforeScenario(func(*godog.Scenario) {
		err := os.MkdirAll(testHomeDir, 0755)
		if err != nil { panic(fmt.Sprintf("Could not create the fake home directory '%s'.", testHomeDir)) }
	})
	ctx.AfterScenario(func(sc *godog.Scenario, _ error) {
		os.RemoveAll(testHomeDir)
	})
}
