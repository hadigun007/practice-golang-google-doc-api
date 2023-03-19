## Google doc api using golang
### requiremets
- golang^1.9 or up

### Google api used
 - [google docs api](https://developers.google.com/docs/api/reference/rest/v1/documents?hl=id)
 - [google drive api](https://developers.google.com/drive/api/guides/about-sdk?hl=id)

### Note
- delete previous token if modifying usr scope

 ### How to start
 - login to your google workspace
 - in the google cloud console, enable the google docs api and google drive api
 - create desktop credentials, [more](https://developers.google.com/docs/api/quickstart/go?hl=en)
 - run project ``` go run main.go ```