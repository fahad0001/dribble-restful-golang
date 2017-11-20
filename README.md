# dribble-restful-golang
The App is developed in GoLang. The App uses Dribble api to get pictures and save it on a gallery

File Structure

configs Folder Contains configuration of the project which includes dribble API access token and path of various data or input sources./

controller Folder Contains 2 Files in which APIHandle file is use to handle the incoming http request to the server whereas the use of driver file is
to handle all other operations like handling db calls, calling dribble API, saving images in a computer e.t.c

dataDir or data Directory Folder contains the files needed for dbInteraction and Storage . dbOperation file is covering all of the aspects of the interaction with 
database. It acts as a bridge to the controller and database.

images Folder is use to store all the images which are retrieved from dribble. 

models Folder contains the schema/model of the Enities which are used to store the data into the db or to convert the incoming IO to json objects.

main File is the kickstart of the project. The basic use of this file is to send config to the related operations of controller.
