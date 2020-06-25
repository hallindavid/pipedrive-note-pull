package pipedrive

import (
	ctoai "github.com/cto-ai/sdk-go"
)

func Start() {
	client := ctoai.NewClient()
	//Track Workflow Start
	trackInit(&client)

	//Retrieve Pipedrive API Key
	pipedriveApiKey := getPipeDriveSecret(&client)

	//Track API Key Test Event
	trackTestAPIKey(&client)

	//Test API Key
	if err := testApiKey(pipedriveApiKey); err != nil {
		//Track failed api key test
		trackFailedApiKeyTest(&client)

		panic(err)
	}

	//Track Successful API Key Test
	trackSuccessfulApiKeyTest(&client)

	//Tell the user it was successful
	if err := client.Ux.Print("API Key test Successful!"); err != nil {
		panic(err)
	}

	//Track Retrieve Pipedrive Notes
	trackGetPipedriveNotes(&client)
	response := ResponseBody{}
	getPipedriveNotes(&client, pipedriveApiKey, &response)

	printDemoNotes(&response, &client)
}
