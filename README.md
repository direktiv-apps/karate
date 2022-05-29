
# karate 1.0

Run karatelabs tests suite.

---
- #### Categories: testing
- #### Image: gcr.io/direktiv/apps/karate 
- #### License: [Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0)
- #### Issue Tracking: https://github.com/direktiv-apps/karate/issues
- #### URL: https://github.com/direktiv-apps/karate
- #### Maintainer: [direktiv.io](https://www.direktiv.io)
---

## About karate

This function runs [karate](https://github.com/karatelabs/karate) test scripts in a Direktiv funtion. 
It provides a logging.xml file where the log level can be configured. The reports can be written to the `out` folder in Direktiv
to use them in subsequent states. Alternativley the last command can `cat` the results, e.g. cat app/target/karate-reports/test.test.json

### Example(s)
  #### Function Configuration
  ```yaml
  functions:
  - id: karate
    image: gcr.io/direktiv/apps/karate:1.0
    type: knative-workflow
  ```
   #### Basic
   ```yaml
   - id: req
     type: action
     action:
       function: karate
       input:
        logging: DEBUG
        commands: 
        - java -Dtest.server=https://www.direktiv.io -jar /karate.jar --output out/workflow/ test.feature
   ```

### Request



#### Request Attributes
[PostParamsBody](#post-params-body)

### Response
  nice greeting
#### Reponse Types
    
  

[PostOKBody](#post-o-k-body)
#### Example Reponses
    
```json
{
  "karate": [
    {
      "result": "06:01:34.468 [main]  INFO  com.intuit.karate - Karate version: 1.2.0",
      "success": true
    }
  ]
}
```

### Errors
| Type | Description
|------|---------|
| io.direktiv.command.error | Command execution failed |
| io.direktiv.output.error | Template error for output generation of the service |
| io.direktiv.ri.error | Can not create information object from request |


### Types
#### <span id="post-o-k-body"></span> postOKBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| karate | [][PostOKBodyKarateItems](#post-o-k-body-karate-items)| `[]*PostOKBodyKarateItems` |  | |  |  |


#### <span id="post-o-k-body-karate-items"></span> postOKBodyKarateItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| result | [interface{}](#interface)| `interface{}` | ✓ | |  |  |
| success | boolean| `bool` | ✓ | |  |  |


#### <span id="post-params-body"></span> postParamsBody

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| commands | [][PostParamsBodyCommandsItems](#post-params-body-commands-items)| `[]*PostParamsBodyCommandsItems` |  | | Array of commands. |  |
| files | [][DirektivFile](#direktiv-file)| `[]apps.DirektivFile` |  | | File to create before running commands. |  |
| logging | string| `string` |  | `"WARN"`| Changes log level in logging.xml. Can be used as `-Dlogback.configurationFile=logging.xml` as argument. | `DEBUG` |


#### <span id="post-params-body-commands-items"></span> postParamsBodyCommandsItems

  



**Properties**

| Name | Type | Go type | Required | Default | Description | Example |
|------|------|---------|:--------:| ------- |-------------|---------|
| command | string| `string` |  | | Command to run | `java -Dtest.server=https://www.direktiv.io -jar /karate.jar --output out/workflow/ test.feature` |
| continue | boolean| `bool` |  | | Stops excecution if command fails, otherwise proceeds with next command |  |
| print | boolean| `bool` |  | `true`| If set to false the command will not print the full command with arguments to logs. |  |
| silent | boolean| `bool` |  | | If set to false the command will not print output to logs. |  |

 
