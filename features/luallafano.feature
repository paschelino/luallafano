Feature: Remember the previous command

    Scenario: Create a first luallafano, a CLI shortcut with superpowers
        Given I entered a complicated command on the CLI:
            """
            echo "Hello Luallafano!"
            """
        And I use Luallafano for the first time
        
        When I type the command "lrc" asking luallafano to remember the command
        And luallafano informs the user:
            """
            Please give your new command a name you can easily remember:
            """
        And a luallafano prompt is shown

        When I enter "e" into the luallafano prompt
        Then luallafano informs the user:
            """
            Do you want to customize the command (Y/n)?
            """
        And a luallafano prompt is shown

        When I enter "n" into the luallafano prompt
        Then a directory hierarchy is being created:
            """
            ~/.luallafano/echo
            ~/.luallafano/bin
            """
        And the file "~/.luallafano/echo/e" is created with this content:
            """
            echo "Hello Luallafano!"
            """
        And the symlink "~/.luallafano/bin/e" to "~/.luallafano/echo/e" is created
        And luallafano informs the user:
            """
            Hi!
            Thank you for using Luallafano 🌻
            This is the first time you're using it. Luallafano has created a 
            couple of directories and created your first command shortcut:
            
            e
            
            To make your command shortcuts instantly usable, please add the
            following line to your shell profile:

            export PATH=$PATH:$HOME/.luallafano/bin

            Or simply execute:
            
            echo "export PATH=$PATH:$HOME/.luallafano/bin" >> .profile
            """
