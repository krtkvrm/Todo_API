# Todo_API

## Execution :
  ```sh
  $ git clone https://github.com/vkartik97/Todo_API.git
  ```
  - ### With Docker :
    ```sh
    $ docker build -t username/go-api .
    $ docker container run -d -p 8000:8000 username/go-api
    ```
  - ### Without Docker :
  ```sh
      $ go build .
      $ ./Todo_API
  ```
  
## API Endpoints :
  - ### Creating a List<br>
    Get Request<br>
    `localhost:8000/create?list=listname`
  - ### Get Todos in a List<br>
    Get Request<br>
    `localhost:8000/get?list=new`
  - ### Add Todos to a List<br>
    POST Request <br>
    `localhost:8001/add`<br>
    Body :<br>
    ```javascript
    {
	    "TaskList": "ListName",
	    "TaskName":"TaskName",
	    "TaskDetails": "Task Details"
    }
    ```
  - ### Delete Todos in a List<br>
    Get Request<br>
    `localhost:8000/deletetask?list=Listname&&id=TaskID`
  - ### Update Todos in a List<br>
    Get Request<br>
    `localhost:8000/update?list=Listname&&id=TaskID&&status=New-Status`
  - ### Delete a List<br>
    Get Request<br>
    `llocalhost:8000/deletelist?list=ListName`
