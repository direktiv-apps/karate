Feature: simple

Scenario: test get

    Given url "https://www.direktiv.io"
    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
        "commands": [
            {
            "command": "java -jar /karate.jar test.feature",
            }
        ]
    }
    """
    When method get
    Then status 200