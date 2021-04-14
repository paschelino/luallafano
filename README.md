# luallafano

## Preconditions
To build this application you need to install [go-task](https://taskfile.dev/#/installation). Then you can run
- end2end tests with `task end2end_tests`
- a build with `task build` that outputs the binary to `out/bin/lrc`

Cleaning the build artifacts is done with `task clean`.

## TODO
- [x] set up go module
- [x] set up cucumber for go
- [ ] implement luallafano.feature up until the first prompt is shown
- [ ] enter the new command's shortcut name until the second prompt is shown
- [ ] enter 'n' for no on customizing the new command until the initial usage message is shown
- [ ] set up a CI for the module via github
- [ ] provide rough readme information

