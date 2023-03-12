# go-underattack-gcp
Still on testing. Not finished.
Example of usage:
`./go-underattack-gcp -project=test-gcp-379719 -action=list`

### GoLang app made to activate or deactivate "under attack" mode on your GCP instances.
First of all you might need to compile the go module to have the binary to exacute:
`go build`

Once you got the binary, you must login with your Google Cloud credentials using `gcloud` and later you
will be able to execute it as the first line:
`./go-underattack-gcp -project=test-gcp-379719 -action=list`

You must chnge the `project` value and not leave it as the example. And indicate the action you would like to perform:
| Action | Description |
| ----------- | ----------- |
| list | List the title of all the current rules |
| activate | Activate underattack mode |
| deactivate | Deactivate underattack mode |