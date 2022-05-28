Feature: greeting end-point

Background:
* url demoBaseUrl

Scenario: say my name

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    When method get
    Then status 200