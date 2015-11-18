# soysos
quick deploy API made in Golang

* you must login  at \<service address\>/login with the following scheme:

  ```
  {
    "username" : "<username>",
    "password" : "<password>"
  }
  ```
  
* the login request will return a session token that you must add to your headers. Use "token" : \<the token you recieved\>
* attempts to use the catfacts api will return Unauthorized errors if you do not have a valid token.
* currently, logging in does not actually check for a valid user, it will return a sessionToken regardless
  
