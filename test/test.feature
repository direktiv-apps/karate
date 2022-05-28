Feature: greeting end-point

Background:
* url demoBaseUrl

Scenario: test get

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
        "commands": [
            {
            "command": "java -Dtest.server=https://www.direktiv.io -jar /karate.jar test.feature",
            }
        ]
    }
    """
    When method post
    Then status 200

Scenario: test logging

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    {   
        "logging": "DOESNOTEXIST",
        "commands": [
            {
            "command": "grep DOESNOTEXIST logging.xml"
            }
        ]
    }
    """
    When method post
    Then status 200
