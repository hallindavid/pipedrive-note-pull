# Pipedrive Note Pull

This is a [CTO.ai](https://cto.ai) op which pulls notes from Pipedrive.

## What does it do?

It is a bit opinionated to be honest - it will prompt you for a date in `YYYY-MM-DD` format, and then retrieves all the notes for that day 

It will then return them in the format

```
John Doe ( john.doe@gmail.com ) at 2020-06-23 07:40 EST
With: {Organization Name}
Notes: { here are some call notes }  
```

## Setup 
The first thing you should do is find your Pipedrive API key.  
[This link](https://pipedrive.readme.io/docs/how-to-find-the-api-token?ref=api_reference) shows you how to get it.

If you just want to try the op to see how it works, you can just pull it from the registry, and run it, it should prompt you to enter this key, and will run


## How to save the api key

If you don't want to always enter your key before it runs, save the api key as a secret in your ops.

Open your terminal and run this command
```shell script
ops secrets:set
```

When you get to this screen, you want to enter `pipedrive_api_key` as the key, and then ops will open VIM or w/e your preferred text editor is and you can paste your key in

```shell script
Initializing... done

🔑 Add a secret to secret storage for team [your team here] →
 Enter the name of the secret to be stored 
```

## Notes / Considerations

* Pipedrive uses UTC timezone for their dates, so if you have a note which occurred after midnight UTC, it may show up as the following day
* This will automatically convert the date to EST - this can be easily changed in the `getpipedrivenotes.go` file by changing the `toTz` variable to your timezone

## Packages Used
* [grokify/html-strip-tags-go](https://github.com/grokify/html-strip-tags-go) In order to strip the HTML tags out of the notes content we use 

## Support
To say thanks, you can share the project on social media or <br />

<a href="https://www.buymeacoffee.com/tDbQ4kg" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>

## Issues
Please report all issues in the GitHub Issue tracker

## Contributing
Shoot me an email, or DM me on twitter and I am happy to allow other contributors.