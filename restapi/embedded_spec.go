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
    "description": "Run karatelabs tests suite.",
    "title": "karate",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "testing"
      ],
      "container": "direktiv/karate",
      "issues": "https://github.com/direktiv-apps/karate/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function runs [karate](https://github.com/karatelabs/karate) test scripts in a Direktiv funtion. \nIt provides a logging.xml file where the log level can be configured. The reports can be written to the ` + "`" + `out` + "`" + ` folder in Direktiv\nto use them in subsequent states.",
      "maintainer": "[direktiv.io](https://www.direktiv.io)",
      "url": "https://github.com/direktiv-apps/karate"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
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
            "description": "nice greeting",
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
              "jj": {
                "sss": "ss"
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
              "exec": "bash -c \"sed 's/WARN/{{ default \"WARN\" .Logging }}/g' /log-config.xml \u003e logging.xml\""
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
            "content": "- id: req\n     type: action\n     action:\n       function: karate\n       input:\n        logging: DEBUG\n        commands: \n        - java -Dtest.server=https://www.direktiv.io -jar /karate.jar --output out/workflow/ test.feature",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: karate\n    image: direktiv/karate:1.0\n    type: knative-workflow"
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
    "description": "Run karatelabs tests suite.",
    "title": "karate",
    "version": "1.0",
    "x-direktiv-meta": {
      "categories": [
        "testing"
      ],
      "container": "direktiv/karate",
      "issues": "https://github.com/direktiv-apps/karate/issues",
      "license": "[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)",
      "long-description": "This function runs [karate](https://github.com/karatelabs/karate) test scripts in a Direktiv funtion. \nIt provides a logging.xml file where the log level can be configured. The reports can be written to the ` + "`" + `out` + "`" + ` folder in Direktiv\nto use them in subsequent states.",
      "maintainer": "[direktiv.io](https://www.direktiv.io)",
      "url": "https://github.com/direktiv-apps/karate"
    }
  },
  "paths": {
    "/": {
      "post": {
        "parameters": [
          {
            "type": "string",
            "description": "direktiv action id is an UUID. \nFor development it can be set to 'development'\n",
            "name": "Direktiv-ActionID",
            "in": "header"
          },
          {
            "type": "string",
            "description": "direktiv temp dir is the working directory for that request\nFor development it can be set to e.g. '/tmp'\n",
            "name": "Direktiv-TempDir",
            "in": "header"
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "commands": {
                  "description": "Array of commands.",
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/CommandsItems0"
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
            "description": "nice greeting",
            "schema": {
              "type": "object",
              "properties": {
                "karate": {
                  "type": "array",
                  "items": {
                    "$ref": "#/definitions/KarateItems0"
                  }
                }
              }
            },
            "examples": {
              "jj": {
                "sss": "ss"
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
              "exec": "bash -c \"sed 's/WARN/{{ default \"WARN\" .Logging }}/g' /log-config.xml \u003e logging.xml\""
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
            "content": "- id: req\n     type: action\n     action:\n       function: karate\n       input:\n        logging: DEBUG\n        commands: \n        - java -Dtest.server=https://www.direktiv.io -jar /karate.jar --output out/workflow/ test.feature",
            "title": "Basic"
          }
        ],
        "x-direktiv-function": "functions:\n  - id: karate\n    image: direktiv/karate:1.0\n    type: knative-workflow"
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
    "CommandsItems0": {
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
    },
    "KarateItems0": {
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
    },
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
}
