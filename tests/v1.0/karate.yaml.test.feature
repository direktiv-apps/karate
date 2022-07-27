
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:

Scenario: karate

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{	
		"files": [
			{
				"name": "test.feature",
				"data": "Feature: simple\n\nScenario: test get\n\n    Given url \"https://www.direktiv.io\"\n    Given path '/'\n    Given header Direktiv-ActionID = 'development'\n    Given header Direktiv-Tempdir = '/tmp'\n    And request \n    \"\"\"\n    { \n        \"commands\": [\n            {\n            \"command\": \"java -jar /karate.jar test.feature\",\n            }\n        ]\n    }\n    \"\"\"\n    When method get\n    Then status 200"
			}
		],
		"commands": [
		{
			"command": "java -jar /karate.jar test.feature",
			"silent": false,
			"print": true,
		}
		]
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	{
	"karate": [
	{
		"result": "#notnull",
		"success": true
	}
	]
	}
	"""

Scenario: testlogging

	Given url karate.properties['testURL']

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

Scenario: version

	Given url karate.properties['testURL']

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    {   
        "commands": [
            {
            "command": "java -jar /karate.jar version"
            }
        ]
    }
    """
    When method post
    Then status 200