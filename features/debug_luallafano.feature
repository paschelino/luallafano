Feature: Debug Luallafano

    As Luallafano developer
    I want to inspect it
    So that I can easily fix issues

    Example: Inspect config and directories
        Given a directory hierarchy exists:
            """
            ~/.luallafano/echo
            ~/.luallafano/bin
            """
        When I type "luallafano inspect"
        Then the output matches:
            """
            Files and directories:
                ~/.luallafano/echo
                ~/.luallafano/bin
            Path variable:
            .*
           """
