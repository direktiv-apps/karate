// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "KarateLab's Karate API testing",
    "title": "karate",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "build"
      ],
      "container": "direktiv.azurecr.io/functions/karate",
      "issues": "https://github.com/direktiv-apps/karate/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function runs [karate](https://github.com/karatelabs/karate) test scripts in a function.  It provides a logging.xml file where the log level can be configured. The reports can be written to the ` + "`" + `out` + "`" + ` folder in Direktiv to use them in subsequent states. Alternativley the last command can ` + "`" + `cat` + "`" + ` the results, e.g. cat app/target/karate-reports/test.test.json",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/karate"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "items": {
                    "type": "object",
                    "properties": {
                      "command": {
                        "description": "Command to run",
                        "type": "string",
                        "example": "java -Dtest.server=https://www.direktiv.io -jar /karate.jar --output out/workflow/ test.feature"
                      },
                      "continue": {
                        "description": "Stops excecution if command fails, otherwise proceeds with next command",
                        "type": "boolean"
                      },
                      "print": {
                        "description": "If set to false the command will not print the full command with arguments to logs.",
                        "type": "boolean",
                        "default": true
                      },
                      "silent": {
                        "description": "If set to false the command will not print output to logs.",
                        "type": "boolean",
                        "default": false
                      }
                    }
                  }
                },
                "files": {
                  "description": "File to create before running commands.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/direktivFile"
                  }
                },
                "logging": {
                  "description": "Changes log level in logging.xml. Can be used as ` + "`" + `-Dlogback.configurationFile=logging.xml` + "`" + ` as argument.",
                  "type": "string",
                  "default": "WARN",
                  "example": "DEBUG"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "type": "object",
              "properties": {
                "karate": {
                  "type": "array",
                  "items": {
                    "type": "object",
                    "required": [
                      "success",
                      "result"
                    ],
                    "properties": {
                      "result": {
                        "additionalProperties": false
                      },
                      "success": {
                        "type": "boolean"
                      }
                    }
                  }
                }
              }
            },
            "examples": {
              "examples": null,
              "karate": {
                "karate": [
                  {
                    "result": "06:01:34.468 [main]  INFO  com.intuit.karate - Karate version: 1.2.0",
                    "success": true
                  }
                ]
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "continue": false,
              "exec": "bash -c \"sed 's/WARN/{{ default \"WARN\" .Logging }}/g' /log-config.xml \u003e logging.xml\"",
              "print": false
            },
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"karate\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: karate\n  type: action\n  action:\n    function: karate\n    input:\n      commands:\n      - command: java -jar /karate.jar test.feature\n      - command: cat target/karate-reports/test.karate-json.txt\n      files:\n      - data: |-\n          Feature: simple\n          Scenario: test get\n              Given url \"https://www.direktiv.io\"\n              Given path '/'\n              When method get\n              Then status 200\n        name: test.feature",
            "title": "Basic"
          },
          {
            "content": "- id: karate\n  type: action\n  action:\n    function: karate\n    input:\n      logging: WARN\n      commands:\n      - command: java -Dlogback.configurationFile=logging.xml -jar /karate.jar  test.feature\n      files:\n      - data: |-\n          Feature: simple\n          Scenario: test get\n              Given url \"https://www.direktiv.io\"\n              Given path '/'\n              When method get\n              Then status 200\n        name: test.feature",
            "title": "Logging"
          },
          {
            "content": "- id: karate\n  type: action\n  action:\n    function: karate\n    input:\n      logging: WARN\n      commands:\n      - command: java -jar /karate.jar -f ~html,cucumber:json test.feature\n      - command: cp target/karate-reports/test.json out/workflow/test-result.json\n      files:\n      - data: |-\n          Feature: simple\n          Scenario: test get\n              Given url \"https://www.direktiv.io\"\n              Given path '/'\n              When method get\n              Then status 200\n        name: test.feature",
            "title": "Store in Variable"
          }
        ],
        "x-direktiv-function": "functions:\n- id: karate\n  image: direktiv.azurecr.io/functions/karate:1.0\n  type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "KarateLab's Karate API testing",
    "title": "karate",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "build"
      ],
      "container": "direktiv.azurecr.io/functions/karate",
      "issues": "https://github.com/direktiv-apps/karate/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function runs [karate](https://github.com/karatelabs/karate) test scripts in a function.  It provides a logging.xml file where the log level can be configured. The reports can be written to the ` + "`" + `out` + "`" + ` folder in Direktiv to use them in subsequent states. Alternativley the last command can ` + "`" + `cat` + "`" + ` the results, e.g. cat app/target/karate-reports/test.test.json",
      "maintainer": "[direktiv.io](https://www.direktiv.io) ",
      "url": "https://github.com/direktiv-apps/karate"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "default": "development",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "default": "/tmp",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/postParamsBody"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "List of executed commands.",
            "schema": {
              "$ref": "#/definitions/postOKBody"
            },
            "examples": {
              "examples": null,
              "karate": {
                "karate": [
                  {
                    "result": "06:01:34.468 [main]  INFO  com.intuit.karate - Karate version: 1.2.0",
                    "success": true
                  }
                ]
              }
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            },
            "headers": {
              "Direktiv-ErrorCode": {
                "type": "string"
              },
              "Direktiv-ErrorMessage": {
                "type": "string"
              }
            }
          }
        },
        "x-direktiv": {
          "cmds": [
            {
              "action": "exec",
              "continue": false,
              "exec": "bash -c \"sed 's/WARN/{{ default \"WARN\" .Logging }}/g' /log-config.xml \u003e logging.xml\"",
              "print": false
            },
            {
              "action": "foreach",
              "continue": "{{ .Item.Continue }}",
              "exec": "{{ .Item.Command }}",
              "loop": ".Commands",
              "print": "{{ .Item.Print }}",
              "silent": "{{ .Item.Silent }}"
            }
          ],
          "output": "{\n  \"karate\": {{ index . 1 | toJson }}\n}\n"
        },
        "x-direktiv-errors": {
          "io.direktiv.command.error": "Command execution failed",
          "io.direktiv.output.error": "Template error for output generation of the service",
          "io.direktiv.ri.error": "Can not create information object from request"
        },
        "x-direktiv-examples": [
          {
            "content": "- id: karate\n  type: action\n  action:\n    function: karate\n    input:\n      commands:\n      - command: java -jar /karate.jar test.feature\n      - command: cat target/karate-reports/test.karate-json.txt\n      files:\n      - data: |-\n          Feature: simple\n          Scenario: test get\n              Given url \"https://www.direktiv.io\"\n              Given path '/'\n              When method get\n              Then status 200\n        name: test.feature",
            "title": "Basic"
          },
          {
            "content": "- id: karate\n  type: action\n  action:\n    function: karate\n    input:\n      logging: WARN\n      commands:\n      - command: java -Dlogback.configurationFile=logging.xml -jar /karate.jar  test.feature\n      files:\n      - data: |-\n          Feature: simple\n          Scenario: test get\n              Given url \"https://www.direktiv.io\"\n              Given path '/'\n              When method get\n              Then status 200\n        name: test.feature",
            "title": "Logging"
          },
          {
            "content": "- id: karate\n  type: action\n  action:\n    function: karate\n    input:\n      logging: WARN\n      commands:\n      - command: java -jar /karate.jar -f ~html,cucumber:json test.feature\n      - command: cp target/karate-reports/test.json out/workflow/test-result.json\n      files:\n      - data: |-\n          Feature: simple\n          Scenario: test get\n              Given url \"https://www.direktiv.io\"\n              Given path '/'\n              When method get\n              Then status 200\n        name: test.feature",
            "title": "Store in Variable"
          }
        ],
        "x-direktiv-function": "functions:\n- id: karate\n  image: direktiv.azurecr.io/functions/karate:1.0\n  type: knative-workflow"
      },
      "delete": {
        "parameters": [
          {
            "type": "string",
            "description": "On cancel Direktiv sends a DELETE request to\nthe action with id in the header\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          }
        ],
        "responses": {
          "200": {
            "description": ""
          }
        },
        "x-direktiv": {
          "cancel": "echo 'cancel {{ .DirektivActionID }}'"
        }
      }
    }
  },
  "definitions": {
    "direktivFile": {
      "type": "object",
      "x-go-type": {
        "import": {
          "package": "github.com/direktiv/apps/go/pkg/apps"
        },
        "type": "DirektivFile"
      }
    },
    "error": {
      "type": "object",
      "required": [
        "errorCode",
        "errorMessage"
      ],
      "properties": {
        "errorCode": {
          "type": "string"
        },
        "errorMessage": {
          "type": "string"
        }
      }
    },
    "postOKBody": {
      "type": "object",
      "properties": {
        "karate": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/postOKBodyKarateItems"
          }
        }
      },
      "x-go-gen-location": "operations"
    },
    "postOKBodyKarateItems": {
      "type": "object",
      "required": [
        "success",
        "result"
      ],
      "properties": {
        "result": {
          "additionalProperties": false
        },
        "success": {
          "type": "boolean"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBody": {
      "type": "object",
      "properties": {
        "commands": {
          "description": "Array of commands.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/postParamsBodyCommandsItems"
          }
        },
        "files": {
          "description": "File to create before running commands.",
          "type": "array",
          "items": {
            "$ref": "#/definitions/direktivFile"
          }
        },
        "logging": {
          "description": "Changes log level in logging.xml. Can be used as ` + "`" + `-Dlogback.configurationFile=logging.xml` + "`" + ` as argument.",
          "type": "string",
          "default": "WARN",
          "example": "DEBUG"
        }
      },
      "x-go-gen-location": "operations"
    },
    "postParamsBodyCommandsItems": {
      "type": "object",
      "properties": {
        "command": {
          "description": "Command to run",
          "type": "string",
          "example": "java -Dtest.server=https://www.direktiv.io -jar /karate.jar --output out/workflow/ test.feature"
        },
        "continue": {
          "description": "Stops excecution if command fails, otherwise proceeds with next command",
          "type": "boolean"
        },
        "print": {
          "description": "If set to false the command will not print the full command with arguments to logs.",
          "type": "boolean",
          "default": true
        },
        "silent": {
          "description": "If set to false the command will not print output to logs.",
          "type": "boolean",
          "default": false
        }
      },
      "x-go-gen-location": "operations"
    }
  }
}`))
}
